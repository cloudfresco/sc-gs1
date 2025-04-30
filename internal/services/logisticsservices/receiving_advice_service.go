package logisticsservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	logisticsstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/logistics/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// ReceivingAdviceService - For accessing ReceivingAdvice services
type ReceivingAdviceService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	logisticsproto.UnimplementedReceivingAdviceServiceServer
}

// NewReceivingAdviceService - Create ReceivingAdvice service
func NewReceivingAdviceService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ReceivingAdviceService {
	return &ReceivingAdviceService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

const insertReceivingAdviceSQL = `insert into receiving_advices
	  (
uuid4,
reporting_code,
total_accepted_amount,
taa_code_list_version,
taa_currency_code,
total_deposit_amount,
tda_code_list_version,
tda_currency_code,
total_number_of_lines,
total_on_hold_amount,
toha_code_list_version,
toha_currency_code,
total_rejected_amount,
tra_code_list_version,
tra_currency_code,
receiving_advice_transport_information,
bill_of_lading_number,
buyer,
carrier,
consignment_identification,
delivery_note,
despatch_advice,
inventory_location,
purchase_order,
receiver,
receiving_advice_identification,
seller,
ship_from,
shipment_identification,
shipper,
ship_to,
despatch_advice_delivery_date_time_begin,
despatch_advice_delivery_date_time_end,
payment_date_time_begin,
payment_date_time_end,
receiving_date_time_begin,
receiving_date_time_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
:uuid4,
:reporting_code,
:total_accepted_amount,
:taa_code_list_version,
:taa_currency_code,
:total_deposit_amount,
:tda_code_list_version,
:tda_currency_code,
:total_number_of_lines,
:total_on_hold_amount,
:toha_code_list_version,
:toha_currency_code,
:total_rejected_amount,
:tra_code_list_version,
:tra_currency_code,
:receiving_advice_transport_information,
:bill_of_lading_number,
:buyer,
:carrier,
:consignment_identification,
:delivery_note,
:despatch_advice,
:inventory_location,
:purchase_order,
:receiver,
:receiving_advice_identification,
:seller,
:ship_from,
:shipment_identification,
:shipper,
:ship_to,
:despatch_advice_delivery_date_time_begin,
:despatch_advice_delivery_date_time_end,
:payment_date_time_begin,
:payment_date_time_end,
:receiving_date_time_begin,
:receiving_date_time_end,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectReceivingAdvicesSQL = `select
  id,
  uuid4,
  reporting_code,
  total_accepted_amount,
  taa_code_list_version,
  taa_currency_code,
  total_deposit_amount,
  tda_code_list_version,
  tda_currency_code,
  total_number_of_lines,
  total_on_hold_amount,
  toha_code_list_version,
  toha_currency_code,
  total_rejected_amount,
  tra_code_list_version,
  tra_currency_code,
  receiving_advice_transport_information,
  bill_of_lading_number,
  buyer,
  carrier,
  consignment_identification,
  delivery_note,
  despatch_advice,
  inventory_location,
  purchase_order,
  receiver,
  receiving_advice_identification,
  seller,
  ship_from,
  shipment_identification,
  shipper,
  ship_to,
  despatch_advice_delivery_date_time_begin,
  despatch_advice_delivery_date_time_end,
  payment_date_time_begin,
  payment_date_time_end,
  receiving_date_time_begin,
  receiving_date_time_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from receiving_advices`

// updateReceivingAdviceSQL - update ReceivingAdviceSQL query
const updateReceivingAdviceSQL = `update receiving_advices set 
  reporting_code= ?,
  total_accepted_amount= ?,
  taa_code_list_version= ?,
  taa_currency_code= ?,
  total_deposit_amount= ?,
  tda_code_list_version= ?,
  tda_currency_code= ?,
  total_number_of_lines= ?,
  total_on_hold_amount= ?,
  toha_code_list_version= ?,
  toha_currency_code= ?,
  total_rejected_amount= ?,
  tra_code_list_version= ?,
  tra_currency_code= ?,
  updated_at = ? where uuid4 = ?;`

// CreateReceivingAdvice - CreateReceivingAdvice
func (ras *ReceivingAdviceService) CreateReceivingAdvice(ctx context.Context, in *logisticsproto.CreateReceivingAdviceRequest) (*logisticsproto.CreateReceivingAdviceResponse, error) {
	receivingAdvice, err := ras.ProcessReceivingAdviceRequest(ctx, in)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ras.insertReceivingAdvice(ctx, insertReceivingAdviceSQL, receivingAdvice, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceResponse := logisticsproto.CreateReceivingAdviceResponse{}
	receivingAdviceResponse.ReceivingAdvice = receivingAdvice
	return &receivingAdviceResponse, nil
}

// ProcessReceivingAdviceRequest - ProcessReceivingAdviceRequest
func (ras *ReceivingAdviceService) ProcessReceivingAdviceRequest(ctx context.Context, in *logisticsproto.CreateReceivingAdviceRequest) (*logisticsproto.ReceivingAdvice, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ras.UserServiceClient)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	paymentDateTimeBegin, err := time.Parse(common.Layout, in.PaymentDateTimeBegin)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	paymentDateTimeEnd, err := time.Parse(common.Layout, in.PaymentDateTimeEnd)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceDeliveryDateTimeBegin, err := time.Parse(common.Layout, in.DespatchAdviceDeliveryDateTimeBegin)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceDeliveryDateTimeEnd, err := time.Parse(common.Layout, in.DespatchAdviceDeliveryDateTimeEnd)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingDateTimeBegin, err := time.Parse(common.Layout, in.ReceivingDateTimeBegin)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingDateTimeEnd, err := time.Parse(common.Layout, in.ReceivingDateTimeEnd)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceD := logisticsproto.ReceivingAdviceD{}
	receivingAdviceD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceD.ReportingCode = in.ReportingCode
	receivingAdviceD.TotalAcceptedAmount = in.TotalAcceptedAmount
	receivingAdviceD.TaaCodeListVersion = in.TaaCodeListVersion
	receivingAdviceD.TaaCurrencyCode = in.TaaCurrencyCode
	receivingAdviceD.TotalDepositAmount = in.TotalDepositAmount
	receivingAdviceD.TdaCodeListVersion = in.TdaCodeListVersion
	receivingAdviceD.TdaCurrencyCode = in.TdaCurrencyCode
	receivingAdviceD.TotalNumberOfLines = in.TotalNumberOfLines
	receivingAdviceD.TotalOnHoldAmount = in.TotalOnHoldAmount
	receivingAdviceD.TohaCodeListVersion = in.TohaCodeListVersion
	receivingAdviceD.TohaCurrencyCode = in.TohaCurrencyCode
	receivingAdviceD.TotalRejectedAmount = in.TotalRejectedAmount
	receivingAdviceD.TraCodeListVersion = in.TraCodeListVersion
	receivingAdviceD.TraCurrencyCode = in.TraCurrencyCode
	receivingAdviceD.ReceivingAdviceTransportInformation = in.ReceivingAdviceTransportInformation
	receivingAdviceD.BillOfLadingNumber = in.BillOfLadingNumber
	receivingAdviceD.Buyer = in.Buyer
	receivingAdviceD.Carrier = in.Carrier
	receivingAdviceD.ConsignmentIdentification = in.ConsignmentIdentification
	receivingAdviceD.DeliveryNote = in.DeliveryNote
	receivingAdviceD.DespatchAdvice = in.DespatchAdvice
	receivingAdviceD.InventoryLocation = in.InventoryLocation
	receivingAdviceD.PurchaseOrder = in.PurchaseOrder
	receivingAdviceD.Receiver = in.Receiver
	receivingAdviceD.ReceivingAdviceIdentification = in.ReceivingAdviceIdentification
	receivingAdviceD.Seller = in.Seller
	receivingAdviceD.ShipFrom = in.ShipFrom
	receivingAdviceD.ShipmentIdentification = in.ShipmentIdentification
	receivingAdviceD.Shipper = in.Shipper
	receivingAdviceD.ShipTo = in.ShipTo

	receivingAdviceT := logisticsproto.ReceivingAdviceT{}
	receivingAdviceT.DespatchAdviceDeliveryDateTimeBegin = common.TimeToTimestamp(despatchAdviceDeliveryDateTimeBegin.UTC().Truncate(time.Second))
	receivingAdviceT.DespatchAdviceDeliveryDateTimeEnd = common.TimeToTimestamp(despatchAdviceDeliveryDateTimeEnd.UTC().Truncate(time.Second))
	receivingAdviceT.PaymentDateTimeBegin = common.TimeToTimestamp(paymentDateTimeBegin.UTC().Truncate(time.Second))
	receivingAdviceT.PaymentDateTimeEnd = common.TimeToTimestamp(paymentDateTimeEnd.UTC().Truncate(time.Second))
	receivingAdviceT.ReceivingDateTimeBegin = common.TimeToTimestamp(receivingDateTimeBegin.UTC().Truncate(time.Second))
	receivingAdviceT.ReceivingDateTimeEnd = common.TimeToTimestamp(receivingDateTimeEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	receivingAdvice := logisticsproto.ReceivingAdvice{ReceivingAdviceD: &receivingAdviceD, ReceivingAdviceT: &receivingAdviceT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &receivingAdvice, nil
}

// insertReceivingAdvice - Insert ReceivingAdvice into database
func (ras *ReceivingAdviceService) insertReceivingAdvice(ctx context.Context, insertReceivingAdviceSQL string, receivingAdvice *logisticsproto.ReceivingAdvice, userEmail string, requestID string) error {
	receivingAdviceTmp, err := ras.crReceivingAdviceStruct(ctx, receivingAdvice, userEmail, requestID)
	if err != nil {
		ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ras.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertReceivingAdviceSQL, receivingAdviceTmp)
		if err != nil {
			ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		receivingAdvice.ReceivingAdviceD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(receivingAdvice.ReceivingAdviceD.Uuid4)
		if err != nil {
			ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		receivingAdvice.ReceivingAdviceD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crReceivingAdviceStruct - process ReceivingAdvice details
func (ras *ReceivingAdviceService) crReceivingAdviceStruct(ctx context.Context, receivingAdvice *logisticsproto.ReceivingAdvice, userEmail string, requestID string) (*logisticsstruct.ReceivingAdvice, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(receivingAdvice.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(receivingAdvice.CrUpdTime.UpdatedAt)

	receivingAdviceT := new(logisticsstruct.ReceivingAdviceT)
	receivingAdviceT.DespatchAdviceDeliveryDateTimeBegin = common.TimestampToTime(receivingAdvice.ReceivingAdviceT.DespatchAdviceDeliveryDateTimeBegin)
	receivingAdviceT.DespatchAdviceDeliveryDateTimeEnd = common.TimestampToTime(receivingAdvice.ReceivingAdviceT.DespatchAdviceDeliveryDateTimeEnd)
	receivingAdviceT.PaymentDateTimeBegin = common.TimestampToTime(receivingAdvice.ReceivingAdviceT.PaymentDateTimeBegin)
	receivingAdviceT.PaymentDateTimeEnd = common.TimestampToTime(receivingAdvice.ReceivingAdviceT.PaymentDateTimeEnd)
	receivingAdviceT.ReceivingDateTimeBegin = common.TimestampToTime(receivingAdvice.ReceivingAdviceT.ReceivingDateTimeBegin)
	receivingAdviceT.ReceivingDateTimeEnd = common.TimestampToTime(receivingAdvice.ReceivingAdviceT.ReceivingDateTimeEnd)

	receivingAdviceTmp := logisticsstruct.ReceivingAdvice{ReceivingAdviceD: receivingAdvice.ReceivingAdviceD, ReceivingAdviceT: receivingAdviceT, CrUpdUser: receivingAdvice.CrUpdUser, CrUpdTime: crUpdTime}

	return &receivingAdviceTmp, nil
}

// GetReceivingAdvices - Get ReceivingAdvices
func (ras *ReceivingAdviceService) GetReceivingAdvices(ctx context.Context, in *logisticsproto.GetReceivingAdvicesRequest) (*logisticsproto.GetReceivingAdvicesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ras.DBService.LimitSQLRows
	}

	query := "(status_code = ?)"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	receivingAdvices := []*logisticsproto.ReceivingAdvice{}

	nselectReceivingAdvicesSQL := selectReceivingAdvicesSQL + ` where ` + query

	rows, err := ras.DBService.DB.QueryxContext(ctx, nselectReceivingAdvicesSQL, common.Active)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		receivingAdviceTmp := logisticsstruct.ReceivingAdvice{}
		err = rows.StructScan(&receivingAdviceTmp)
		if err != nil {
			ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		receivingAdvice, err := ras.getReceivingAdviceStruct(ctx, &getRequest, receivingAdviceTmp)
		if err != nil {
			ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		receivingAdvices = append(receivingAdvices, receivingAdvice)

	}

	receivingAdvicesResponse := logisticsproto.GetReceivingAdvicesResponse{}
	if len(receivingAdvices) != 0 {
		next := receivingAdvices[len(receivingAdvices)-1].ReceivingAdviceD.Id
		next--
		nextc := common.EncodeCursor(next)
		receivingAdvicesResponse = logisticsproto.GetReceivingAdvicesResponse{ReceivingAdvices: receivingAdvices, NextCursor: nextc}
	} else {
		receivingAdvicesResponse = logisticsproto.GetReceivingAdvicesResponse{ReceivingAdvices: receivingAdvices, NextCursor: "0"}
	}
	return &receivingAdvicesResponse, nil
}

// GetReceivingAdvice - Get ReceivingAdvice
func (ras *ReceivingAdviceService) GetReceivingAdvice(ctx context.Context, inReq *logisticsproto.GetReceivingAdviceRequest) (*logisticsproto.GetReceivingAdviceResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectReceivingAdvicesSQL := selectReceivingAdvicesSQL + ` where uuid4 = ?;`
	row := ras.DBService.DB.QueryRowxContext(ctx, nselectReceivingAdvicesSQL, uuid4byte)
	receivingAdviceTmp := logisticsstruct.ReceivingAdvice{}
	err = row.StructScan(&receivingAdviceTmp)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdvice, err := ras.getReceivingAdviceStruct(ctx, in, receivingAdviceTmp)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	receivingAdviceResponse := logisticsproto.GetReceivingAdviceResponse{}
	receivingAdviceResponse.ReceivingAdvice = receivingAdvice
	return &receivingAdviceResponse, nil
}

// GetReceivingAdviceByPk - Get ReceivingAdvice By Primary key(Id)
func (ras *ReceivingAdviceService) GetReceivingAdviceByPk(ctx context.Context, inReq *logisticsproto.GetReceivingAdviceByPkRequest) (*logisticsproto.GetReceivingAdviceByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectReceivingAdvicesSQL := selectReceivingAdvicesSQL + ` where id = ?;`
	row := ras.DBService.DB.QueryRowxContext(ctx, nselectReceivingAdvicesSQL, in.Id)
	receivingAdviceTmp := logisticsstruct.ReceivingAdvice{}
	err := row.StructScan(&receivingAdviceTmp)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	receivingAdvice, err := ras.getReceivingAdviceStruct(ctx, &getRequest, receivingAdviceTmp)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	receivingAdviceResponse := logisticsproto.GetReceivingAdviceByPkResponse{}
	receivingAdviceResponse.ReceivingAdvice = receivingAdvice
	return &receivingAdviceResponse, nil
}

// getReceivingAdviceStruct - Get receivingAdvice
func (ras *ReceivingAdviceService) getReceivingAdviceStruct(ctx context.Context, in *commonproto.GetRequest, receivingAdviceTmp logisticsstruct.ReceivingAdvice) (*logisticsproto.ReceivingAdvice, error) {
	receivingAdviceT := new(logisticsproto.ReceivingAdviceT)
	receivingAdviceT.DespatchAdviceDeliveryDateTimeBegin = common.TimeToTimestamp(receivingAdviceTmp.ReceivingAdviceT.DespatchAdviceDeliveryDateTimeBegin)
	receivingAdviceT.DespatchAdviceDeliveryDateTimeEnd = common.TimeToTimestamp(receivingAdviceTmp.ReceivingAdviceT.DespatchAdviceDeliveryDateTimeEnd)
	receivingAdviceT.PaymentDateTimeBegin = common.TimeToTimestamp(receivingAdviceTmp.ReceivingAdviceT.PaymentDateTimeBegin)
	receivingAdviceT.PaymentDateTimeEnd = common.TimeToTimestamp(receivingAdviceTmp.ReceivingAdviceT.PaymentDateTimeEnd)
	receivingAdviceT.ReceivingDateTimeBegin = common.TimeToTimestamp(receivingAdviceTmp.ReceivingAdviceT.ReceivingDateTimeBegin)
	receivingAdviceT.ReceivingDateTimeEnd = common.TimeToTimestamp(receivingAdviceTmp.ReceivingAdviceT.ReceivingDateTimeEnd)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(receivingAdviceTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(receivingAdviceTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(receivingAdviceTmp.ReceivingAdviceD.Uuid4)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceTmp.ReceivingAdviceD.IdS = uuid4Str

	receivingAdvice := logisticsproto.ReceivingAdvice{ReceivingAdviceD: receivingAdviceTmp.ReceivingAdviceD, ReceivingAdviceT: receivingAdviceT, CrUpdUser: receivingAdviceTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &receivingAdvice, nil
}

// UpdateReceivingAdvice - Update ReceivingAdvice
func (ras *ReceivingAdviceService) UpdateReceivingAdvice(ctx context.Context, in *logisticsproto.UpdateReceivingAdviceRequest) (*logisticsproto.UpdateReceivingAdviceResponse, error) {
	db := ras.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateReceivingAdviceSQL)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ras.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.ReportingCode,
			in.TotalAcceptedAmount,
			in.TaaCodeListVersion,
			in.TaaCurrencyCode,
			in.TotalDepositAmount,
			in.TdaCodeListVersion,
			in.TdaCurrencyCode,
			in.TotalNumberOfLines,
			in.TotalOnHoldAmount,
			in.TohaCodeListVersion,
			in.TohaCurrencyCode,
			in.TotalRejectedAmount,
			in.TraCodeListVersion,
			in.TraCurrencyCode,
			tn,
			uuid4byte)
		if err != nil {
			ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &logisticsproto.UpdateReceivingAdviceResponse{}, nil
}
