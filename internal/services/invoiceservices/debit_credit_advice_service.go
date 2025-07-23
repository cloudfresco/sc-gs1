package invoiceservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	invoicestruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// DebitCreditAdviceService - For accessing DebitCreditAdvice services
type DebitCreditAdviceService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	CurrencyService   *common.CurrencyService
	invoiceproto.UnimplementedDebitCreditAdviceServiceServer
}

// NewDebitCreditAdviceService - Create DebitCreditAdvice service
func NewDebitCreditAdviceService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient, currency *common.CurrencyService) *DebitCreditAdviceService {
	return &DebitCreditAdviceService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
		CurrencyService:   currency,
	}
}

const insertDebitCreditAdviceSQL = `insert into debit_credit_advices
	    (uuid4,
debit_credit_indicator_code,
total_amount,
total_amount_currency,
ta_code_list_version,
bill_to,
buyer,
carrier,
debit_credit_advice_identification,
seller,
ship_from,
ship_to,
ultimate_consignee,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(:uuid4,
:debit_credit_indicator_code,
:total_amount,
:total_amount_currency,
:ta_code_list_version,
:bill_to,
:buyer,
:carrier,
:debit_credit_advice_identification,
:seller,
:ship_from,
:ship_to,
:ultimate_consignee,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectDebitCreditAdvicesSQL = `select
id,
uuid4,
debit_credit_indicator_code,
total_amount,
total_amount_currency,
ta_code_list_version,
bill_to,
buyer,
carrier,
debit_credit_advice_identification,
seller,
ship_from,
ship_to,
ultimate_consignee,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from debit_credit_advices`

// updateDebitCreditAdviceSQL - update DebitCreditAdviceSQL query
const updateDebitCreditAdviceSQL = `update debit_credit_advices set debit_credit_indicator_code = ?, total_amount = ?, ta_code_list_version = ?, updated_at = ? where uuid4 = ?;`

// CreateDebitCreditAdvice - Create DebitCreditAdvice
func (ds *DebitCreditAdviceService) CreateDebitCreditAdvice(ctx context.Context, in *invoiceproto.CreateDebitCreditAdviceRequest) (*invoiceproto.CreateDebitCreditAdviceResponse, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ds.UserServiceClient)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	debitCreditAdviceD := invoiceproto.DebitCreditAdviceD{}
	debitCreditAdviceD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceD.DebitCreditIndicatorCode = in.DebitCreditIndicatorCode
	debitCreditAdviceD.TaCodeListVersion = in.TaCodeListVersion
	debitCreditAdviceD.BillTo = in.BillTo
	debitCreditAdviceD.Buyer = in.Buyer
	debitCreditAdviceD.Carrier = in.Carrier
	debitCreditAdviceD.DebitCreditAdviceIdentification = in.DebitCreditAdviceIdentification
	debitCreditAdviceD.Seller = in.Seller
	debitCreditAdviceD.ShipFrom = in.ShipFrom
	debitCreditAdviceD.ShipTo = in.ShipTo
	debitCreditAdviceD.UltimateConsignee = in.UltimateConsignee

	totalAmountCurrency, err := ds.CurrencyService.GetCurrency(ctx, in.TotalAmountCurrency)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	totalAmountMinor, err := common.ParseAmountString(in.TotalAmount, totalAmountCurrency)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceD.TotalAmountCurrency = totalAmountCurrency.Code
	debitCreditAdviceD.TotalAmount = totalAmountMinor
	debitCreditAdviceD.TotalAmountString = common.FormatAmountString(totalAmountMinor, totalAmountCurrency)

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	debitCreditAdvice := invoiceproto.DebitCreditAdvice{DebitCreditAdviceD: &debitCreditAdviceD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	debitCreditAdviceLineItems := []*invoiceproto.DebitCreditAdviceLineItem{}
	// we will do for loop on lines which is comes from client form
	for _, line := range in.DebitCreditAdviceLineItems {
		line.UserId = in.UserId
		line.UserEmail = in.UserEmail
		line.RequestId = in.RequestId
		// we wl call CreateDebitCreditAdviceLine function which wl populate form values to debitCreditAdviceline struct
		debitCreditAdviceLineItem, err := ds.ProcessDebitCreditAdviceLineItemRequest(ctx, line)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		debitCreditAdviceLineItems = append(debitCreditAdviceLineItems, debitCreditAdviceLineItem)
	}

	err = ds.insertDebitCreditAdvice(ctx, insertDebitCreditAdviceSQL, &debitCreditAdvice, insertDebitCreditAdviceLineItemSQL, debitCreditAdviceLineItems, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceResponse := invoiceproto.CreateDebitCreditAdviceResponse{}
	debitCreditAdviceResponse.DebitCreditAdvice = &debitCreditAdvice
	return &debitCreditAdviceResponse, nil
}

func (ds *DebitCreditAdviceService) insertDebitCreditAdvice(ctx context.Context, insertDebitCreditAdviceSQL string, debitCreditAdvice *invoiceproto.DebitCreditAdvice, insertDebitCreditAdviceLineItemSQL string, debitCreditAdviceLineItems []*invoiceproto.DebitCreditAdviceLineItem, userEmail string, requestID string) error {
	debitCreditAdviceTmp, err := ds.crDebitCreditAdviceStruct(ctx, debitCreditAdvice, userEmail, requestID)
	if err != nil {
		ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ds.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		// header creation
		res, err := tx.NamedExecContext(ctx, insertDebitCreditAdviceSQL, debitCreditAdviceTmp)
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		debitCreditAdvice.DebitCreditAdviceD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(debitCreditAdvice.DebitCreditAdviceD.Uuid4)
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		debitCreditAdvice.DebitCreditAdviceD.IdS = uuid4Str

		for _, debitCreditAdviceLineItem := range debitCreditAdviceLineItems {
			debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.DebitCreditAdviceId = debitCreditAdvice.DebitCreditAdviceD.Id
			debitCreditAdviceLineItemTmp, err := ds.crDebitCreditAdviceLineItemStruct(ctx, debitCreditAdviceLineItem, userEmail, requestID)
			if err != nil {
				ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertDebitCreditAdviceLineItemSQL, debitCreditAdviceLineItemTmp)
			if err != nil {
				ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}
		return nil
	})
	if err != nil {
		ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crDebitCreditAdviceStruct - process DebitCreditAdvice details
func (ds *DebitCreditAdviceService) crDebitCreditAdviceStruct(ctx context.Context, debitCreditAdvice *invoiceproto.DebitCreditAdvice, userEmail string, requestID string) (*invoicestruct.DebitCreditAdvice, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(debitCreditAdvice.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(debitCreditAdvice.CrUpdTime.UpdatedAt)

	debitCreditAdviceTmp := invoicestruct.DebitCreditAdvice{DebitCreditAdviceD: debitCreditAdvice.DebitCreditAdviceD, CrUpdUser: debitCreditAdvice.CrUpdUser, CrUpdTime: crUpdTime}

	return &debitCreditAdviceTmp, nil
}

// GetDebitCreditAdvices - Get DebitCreditAdvices
func (ds *DebitCreditAdviceService) GetDebitCreditAdvices(ctx context.Context, in *invoiceproto.GetDebitCreditAdvicesRequest) (*invoiceproto.GetDebitCreditAdvicesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ds.DBService.LimitSQLRows
	}
	query := "(status_code = ?)"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	debitCreditAdvices := []*invoiceproto.DebitCreditAdvice{}

	nselectDebitCreditAdvicesSQL := selectDebitCreditAdvicesSQL + ` where ` + query

	rows, err := ds.DBService.DB.QueryxContext(ctx, nselectDebitCreditAdvicesSQL, common.Active)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		debitCreditAdviceTmp := invoicestruct.DebitCreditAdvice{}
		err = rows.StructScan(&debitCreditAdviceTmp)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		debitCreditAdvice, err := ds.getDebitCreditAdviceStruct(ctx, &getRequest, debitCreditAdviceTmp)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		totalAmountCurrency, err := ds.CurrencyService.GetCurrency(ctx, debitCreditAdvice.DebitCreditAdviceD.TotalAmountCurrency)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		debitCreditAdvice.DebitCreditAdviceD.TotalAmountString = common.FormatAmountString(debitCreditAdvice.DebitCreditAdviceD.TotalAmount, totalAmountCurrency)

		debitCreditAdvices = append(debitCreditAdvices, debitCreditAdvice)

	}

	debitCreditAdvicesResponse := invoiceproto.GetDebitCreditAdvicesResponse{}
	if len(debitCreditAdvices) != 0 {
		next := debitCreditAdvices[len(debitCreditAdvices)-1].DebitCreditAdviceD.Id
		next--
		nextc := common.EncodeCursor(next)
		debitCreditAdvicesResponse = invoiceproto.GetDebitCreditAdvicesResponse{DebitCreditAdvices: debitCreditAdvices, NextCursor: nextc}
	} else {
		debitCreditAdvicesResponse = invoiceproto.GetDebitCreditAdvicesResponse{DebitCreditAdvices: debitCreditAdvices, NextCursor: "0"}
	}
	return &debitCreditAdvicesResponse, nil
}

// GetDebitCreditAdvice - Get DebitCreditAdvice
func (ds *DebitCreditAdviceService) GetDebitCreditAdvice(ctx context.Context, inReq *invoiceproto.GetDebitCreditAdviceRequest) (*invoiceproto.GetDebitCreditAdviceResponse, error) {
	in := inReq.GetRequest

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectDebitCreditAdvicesSQL := selectDebitCreditAdvicesSQL + ` where uuid4 = ?;`
	row := ds.DBService.DB.QueryRowxContext(ctx, nselectDebitCreditAdvicesSQL, uuid4byte)
	debitCreditAdviceTmp := invoicestruct.DebitCreditAdvice{}
	err = row.StructScan(&debitCreditAdviceTmp)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdvice, err := ds.getDebitCreditAdviceStruct(ctx, in, debitCreditAdviceTmp)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	totalAmountCurrency, err := ds.CurrencyService.GetCurrency(ctx, debitCreditAdvice.DebitCreditAdviceD.TotalAmountCurrency)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdvice.DebitCreditAdviceD.TotalAmountString = common.FormatAmountString(debitCreditAdvice.DebitCreditAdviceD.TotalAmount, totalAmountCurrency)

	debitCreditAdviceResponse := invoiceproto.GetDebitCreditAdviceResponse{}
	debitCreditAdviceResponse.DebitCreditAdvice = debitCreditAdvice
	return &debitCreditAdviceResponse, nil
}

// GetDebitCreditAdviceByPk - Get DebitCreditAdvice By Primary key(Id)
func (ds *DebitCreditAdviceService) GetDebitCreditAdviceByPk(ctx context.Context, inReq *invoiceproto.GetDebitCreditAdviceByPkRequest) (*invoiceproto.GetDebitCreditAdviceByPkResponse, error) {
	in := inReq.GetByIdRequest

	nselectDebitCreditAdvicesSQL := selectDebitCreditAdvicesSQL + ` where id = ?;`
	row := ds.DBService.DB.QueryRowxContext(ctx, nselectDebitCreditAdvicesSQL, in.Id)
	debitCreditAdviceTmp := invoicestruct.DebitCreditAdvice{}
	err := row.StructScan(&debitCreditAdviceTmp)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	debitCreditAdvice, err := ds.getDebitCreditAdviceStruct(ctx, &getRequest, debitCreditAdviceTmp)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	totalAmountCurrency, err := ds.CurrencyService.GetCurrency(ctx, debitCreditAdvice.DebitCreditAdviceD.TotalAmountCurrency)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdvice.DebitCreditAdviceD.TotalAmountString = common.FormatAmountString(debitCreditAdvice.DebitCreditAdviceD.TotalAmount, totalAmountCurrency)

	debitCreditAdviceResponse := invoiceproto.GetDebitCreditAdviceByPkResponse{}
	debitCreditAdviceResponse.DebitCreditAdvice = debitCreditAdvice
	return &debitCreditAdviceResponse, nil
}

// getDebitCreditAdviceStruct - Get debitCreditAdvice
func (ds *DebitCreditAdviceService) getDebitCreditAdviceStruct(ctx context.Context, in *commonproto.GetRequest, debitCreditAdviceTmp invoicestruct.DebitCreditAdvice) (*invoiceproto.DebitCreditAdvice, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(debitCreditAdviceTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(debitCreditAdviceTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(debitCreditAdviceTmp.DebitCreditAdviceD.Uuid4)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceTmp.DebitCreditAdviceD.IdS = uuid4Str

	debitCreditAdvice := invoiceproto.DebitCreditAdvice{DebitCreditAdviceD: debitCreditAdviceTmp.DebitCreditAdviceD, CrUpdUser: debitCreditAdviceTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &debitCreditAdvice, nil
}

// UpdateDebitCreditAdvice - Update DebitCreditAdvice
func (ds *DebitCreditAdviceService) UpdateDebitCreditAdvice(ctx context.Context, in *invoiceproto.UpdateDebitCreditAdviceRequest) (*invoiceproto.UpdateDebitCreditAdviceResponse, error) {
	db := ds.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateDebitCreditAdviceSQL)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ds.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.DebitCreditIndicatorCode,
			in.TotalAmount,
			in.TaCodeListVersion,
			tn,
			uuid4byte)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &invoiceproto.UpdateDebitCreditAdviceResponse{}, nil
}
