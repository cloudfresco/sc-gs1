package orderservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	orderstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/order/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrderResponseSQL = `insert into order_responses
	    (uuid4,
order_response_reason_code,
response_status_code,
total_monetary_amount_excluding_taxes,
tmaet_code_list_version,
tmaet_currency_code,
total_monetary_amount_including_taxes,
tmait_code_list_version,
tmait_currency_code,
total_tax_amount,
tta_code_list_version,
tta_currency_code,
amended_date_time_value,
bill_to,
buyer,
order_response_identification,
original_order,
sales_order,
seller,
ship_to,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
:uuid4,
:order_response_reason_code,
:response_status_code,
:total_monetary_amount_excluding_taxes,
:tmaet_code_list_version,
:tmaet_currency_code,
:total_monetary_amount_including_taxes,
:tmait_code_list_version,
:tmait_currency_code,
:total_tax_amount,
:tta_code_list_version,
:tta_currency_code,
:amended_date_time_value,
:bill_to,
:buyer,
:order_response_identification,
:original_order,
:sales_order,
:seller,
:ship_to,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectOrderResponsesSQL = `select
    id,
uuid4,
order_response_reason_code,
response_status_code,
total_monetary_amount_excluding_taxes,
tmaet_code_list_version,
tmaet_currency_code,
total_monetary_amount_including_taxes,
tmait_code_list_version,
tmait_currency_code,
total_tax_amount,
tta_code_list_version,
tta_currency_code,
amended_date_time_value,
bill_to,
buyer,
order_response_identification,
original_order,
sales_order,
seller,
ship_to,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from order_responses`

// updateOrderResponseSQL - update OrderResponseSQL query
const updateOrderResponseSQL = `update order_responses set 
  order_response_reason_code= ?,
  response_status_code= ?,
  total_monetary_amount_excluding_taxes= ?,
  tmaet_code_list_version= ?,
  tmaet_currency_code= ?,
  total_monetary_amount_including_taxes= ?,
  tmait_code_list_version= ?,
  tmait_currency_code= ?,
  total_tax_amount= ?,
  tta_code_list_version= ?,
  tta_currency_code= ?,
  updated_at = ? where uuid4 = ?;`

// OrderResponseService - For accessing OrderResponse services
type OrderResponseService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	orderproto.UnimplementedOrderResponseServiceServer
}

// NewOrderResponseService - Create OrderResponse service
func NewOrderResponseService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *OrderResponseService {
	return &OrderResponseService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// CreateOrderResponse - Create OrderResponse
func (o *OrderResponseService) CreateOrderResponse(ctx context.Context, in *orderproto.CreateOrderResponseRequest) (*orderproto.CreateOrderResponseResponse, error) {
	orderResponse, err := o.ProcessOrderResponseRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderResponse(ctx, insertOrderResponseSQL, orderResponse, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResp := orderproto.CreateOrderResponseResponse{}
	orderResp.OrderResponse = orderResponse
	return &orderResp, nil
}

// ProcessOrderResponseRequest - ProcessOrderResponseRequest
func (o *OrderResponseService) ProcessOrderResponseRequest(ctx context.Context, in *orderproto.CreateOrderResponseRequest) (*orderproto.OrderResponse, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, o.UserServiceClient)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	orderResponseD := orderproto.OrderResponseD{}
	orderResponseD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponseD.OrderResponseReasonCode = in.OrderResponseReasonCode
	orderResponseD.ResponseStatusCode = in.ResponseStatusCode
	orderResponseD.TotalMonetaryAmountExcludingTaxes = in.TotalMonetaryAmountExcludingTaxes
	orderResponseD.TmaetCodeListVersion = in.TmaetCodeListVersion
	orderResponseD.TmaetCurrencyCode = in.TmaetCurrencyCode
	orderResponseD.TotalMonetaryAmountIncludingTaxes = in.TotalMonetaryAmountIncludingTaxes
	orderResponseD.TmaitCodeListVersion = in.TmaitCodeListVersion
	orderResponseD.TmaitCurrencyCode = in.TmaitCurrencyCode
	orderResponseD.TotalTaxAmount = in.TotalTaxAmount
	orderResponseD.TtaCodeListVersion = in.TtaCodeListVersion
	orderResponseD.TtaCurrencyCode = in.TtaCurrencyCode
	orderResponseD.AmendedDateTimeValue = in.AmendedDateTimeValue
	orderResponseD.BillTo = in.BillTo
	orderResponseD.Buyer = in.Buyer
	orderResponseD.OrderResponseIdentification = in.OrderResponseIdentification
	orderResponseD.OriginalOrder = in.OriginalOrder
	orderResponseD.SalesOrder = in.SalesOrder
	orderResponseD.Seller = in.Seller
	orderResponseD.ShipTo = in.ShipTo

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	orderResponse := orderproto.OrderResponse{OrderResponseD: &orderResponseD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &orderResponse, nil
}

// insertOrderResponse - Insert OrderResponse details into database
func (o *OrderResponseService) insertOrderResponse(ctx context.Context, insertOrderResponseSQL string, orderResponse *orderproto.OrderResponse, userEmail string, requestID string) error {
	orderResponseTmp, err := o.crOrderResponseStruct(ctx, orderResponse, userEmail, requestID)
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertOrderResponseSQL, orderResponseTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		orderResponse.OrderResponseD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(orderResponse.OrderResponseD.Uuid4)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		orderResponse.OrderResponseD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crOrderResponseStruct - process OrderResponse details
func (o *OrderResponseService) crOrderResponseStruct(ctx context.Context, orderResponse *orderproto.OrderResponse, userEmail string, requestID string) (*orderstruct.OrderResponse, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(orderResponse.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(orderResponse.CrUpdTime.UpdatedAt)

	orderResponseTmp := orderstruct.OrderResponse{OrderResponseD: orderResponse.OrderResponseD, CrUpdUser: orderResponse.CrUpdUser, CrUpdTime: crUpdTime}

	return &orderResponseTmp, nil
}

// GetOrderResponses - Get OrderResponses
func (o *OrderResponseService) GetOrderResponses(ctx context.Context, in *orderproto.GetOrderResponsesRequest) (*orderproto.GetOrderResponsesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = o.DBService.LimitSQLRows
	}

	query := "(status_code = ?)"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	orderResponses := []*orderproto.OrderResponse{}

	nselectOrderResponsesSQL := selectOrderResponsesSQL + ` where ` + query

	rows, err := o.DBService.DB.QueryxContext(ctx, nselectOrderResponsesSQL, common.Active)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		orderResponseTmp := orderstruct.OrderResponse{}
		err = rows.StructScan(&orderResponseTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		orderResponse, err := o.getOrderResponseStruct(ctx, &getRequest, orderResponseTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		orderResponses = append(orderResponses, orderResponse)

	}

	ordersResp := orderproto.GetOrderResponsesResponse{}
	if len(orderResponses) != 0 {
		next := orderResponses[len(orderResponses)-1].OrderResponseD.Id
		next--
		nextc := common.EncodeCursor(next)
		ordersResp = orderproto.GetOrderResponsesResponse{OrderResponses: orderResponses, NextCursor: nextc}
	} else {
		ordersResp = orderproto.GetOrderResponsesResponse{OrderResponses: orderResponses, NextCursor: "0"}
	}
	return &ordersResp, nil
}

// GetOrderResponse - Get OrderResponse
func (o *OrderResponseService) GetOrderResponse(ctx context.Context, inReq *orderproto.GetOrderResponseRequest) (*orderproto.GetOrderResponseResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectOrderResponsesSQL := selectOrderResponsesSQL + ` where uuid4 = ?;`
	row := o.DBService.DB.QueryRowxContext(ctx, nselectOrderResponsesSQL, uuid4byte)
	orderResponseTmp := orderstruct.OrderResponse{}
	err = row.StructScan(&orderResponseTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponse, err := o.getOrderResponseStruct(ctx, in, orderResponseTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	orderResp := orderproto.GetOrderResponseResponse{}
	orderResp.OrderResponse = orderResponse
	return &orderResp, nil
}

// GetOrderResponseByPk - Get OrderResponse By Primary key(Id)
func (o *OrderResponseService) GetOrderResponseByPk(ctx context.Context, inReq *orderproto.GetOrderResponseByPkRequest) (*orderproto.GetOrderResponseByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectOrderResponsesSQL := selectOrderResponsesSQL + ` where id = ?;`
	row := o.DBService.DB.QueryRowxContext(ctx, nselectOrderResponsesSQL, in.Id)
	orderResponseTmp := orderstruct.OrderResponse{}
	err := row.StructScan(&orderResponseTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	orderResponse, err := o.getOrderResponseStruct(ctx, &getRequest, orderResponseTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	orderResp := orderproto.GetOrderResponseByPkResponse{}
	orderResp.OrderResponse = orderResponse
	return &orderResp, nil
}

// getOrderResponseStruct - Get orderResponse
func (o *OrderResponseService) getOrderResponseStruct(ctx context.Context, in *commonproto.GetRequest, orderResponseTmp orderstruct.OrderResponse) (*orderproto.OrderResponse, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(orderResponseTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(orderResponseTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(orderResponseTmp.OrderResponseD.Uuid4)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponseTmp.OrderResponseD.IdS = uuid4Str

	orderResponse := orderproto.OrderResponse{OrderResponseD: orderResponseTmp.OrderResponseD, CrUpdUser: orderResponseTmp.CrUpdUser, CrUpdTime: crUpdTime}
	return &orderResponse, nil
}

// UpdateOrderResponse - Update OrderResponse
func (o *OrderResponseService) UpdateOrderResponse(ctx context.Context, in *orderproto.UpdateOrderResponseRequest) (*orderproto.UpdateOrderResponseResponse, error) {
	db := o.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateOrderResponseSQL)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.OrderResponseReasonCode,
			in.ResponseStatusCode,
			in.TotalMonetaryAmountExcludingTaxes,
			in.TmaetCodeListVersion,
			in.TmaetCurrencyCode,
			in.TotalMonetaryAmountIncludingTaxes,
			in.TmaitCodeListVersion,
			in.TmaitCurrencyCode,
			in.TotalTaxAmount,
			in.TtaCodeListVersion,
			in.TtaCurrencyCode,
			tn,
			uuid4byte)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &orderproto.UpdateOrderResponseResponse{}, nil
}
