package orderservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	orderstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/order/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrderResponseLineItemSQL = `insert into order_response_line_items
	    (uuid4,
	    confirmed_quantity,
cq_measurement_unit_code,
cq_code_list_version,
line_item_action_code,
line_item_change_indicator,
line_item_number,
monetary_amount_excluding_taxes,
maet_code_list_version,
maet_currency_code,
monetary_amount_including_taxes,
mait_code_list_version,
mait_currency_code,
net_amount,
na_code_list_version,
na_currency_code,
net_price,
np_code_list_version,
np_currency_code,
order_response_reason_code,
original_order_line_item_number,
parent_line_item_number,
order_response_id,
delivery_date_time,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
:uuid4,  
:confirmed_quantity,
:cq_measurement_unit_code,
:cq_code_list_version,
:line_item_action_code,
:line_item_change_indicator,
:line_item_number,
:monetary_amount_excluding_taxes,
:maet_code_list_version,
:maet_currency_code,
:monetary_amount_including_taxes,
:mait_code_list_version,
:mait_currency_code,
:net_amount,
:na_code_list_version,
:na_currency_code,
:net_price,
:np_code_list_version,
:np_currency_code,
:order_response_reason_code,
:original_order_line_item_number,
:parent_line_item_number,
:order_response_id,
:delivery_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectOrderResponseLineItemsSQL = `select
    id,
    uuid4,
confirmed_quantity,
cq_measurement_unit_code,
cq_code_list_version,
line_item_action_code,
line_item_change_indicator,
line_item_number,
monetary_amount_excluding_taxes,
maet_code_list_version,
maet_currency_code,
monetary_amount_including_taxes,
mait_code_list_version,
mait_currency_code,
net_amount,
na_code_list_version,
na_currency_code,
net_price,
np_code_list_version,
np_currency_code,
order_response_reason_code,
original_order_line_item_number,
parent_line_item_number,
order_response_id,
delivery_date_time,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from order_response_line_items`

// CreateOrderResponseLineItem - Create OrderResponseLineItem
func (o *OrderResponseService) CreateOrderResponseLineItem(ctx context.Context, in *orderproto.CreateOrderResponseLineItemRequest) (*orderproto.CreateOrderResponseLineItemResponse, error) {
	orderResponseLineItem, err := o.ProcessOrderResponseLineItemRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderResponseLineItem(ctx, insertOrderResponseLineItemSQL, orderResponseLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	orderRespLineItem := orderproto.CreateOrderResponseLineItemResponse{}
	orderRespLineItem.OrderResponseLineItem = orderResponseLineItem
	return &orderRespLineItem, nil
}

// ProcessOrderResponseLineItemRequest - ProcessOrderResponseLineItemRequest
func (o *OrderResponseService) ProcessOrderResponseLineItemRequest(ctx context.Context, in *orderproto.CreateOrderResponseLineItemRequest) (*orderproto.OrderResponseLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, o.UserServiceClient)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	deliveryDateTime, err := time.Parse(common.Layout, in.DeliveryDateTime)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponseLineItemD := orderproto.OrderResponseLineItemD{}
	orderResponseLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponseLineItemD.ConfirmedQuantity = in.ConfirmedQuantity
	orderResponseLineItemD.CqMeasurementUnitCode = in.CqMeasurementUnitCode
	orderResponseLineItemD.CqCodeListVersion = in.CqCodeListVersion
	orderResponseLineItemD.LineItemActionCode = in.LineItemActionCode
	orderResponseLineItemD.LineItemChangeIndicator = in.LineItemChangeIndicator
	orderResponseLineItemD.LineItemNumber = in.LineItemNumber
	orderResponseLineItemD.MonetaryAmountExcludingTaxes = in.MonetaryAmountExcludingTaxes
	orderResponseLineItemD.MaetCodeListVersion = in.MaetCodeListVersion
	orderResponseLineItemD.MaetCurrencyCode = in.MaetCurrencyCode
	orderResponseLineItemD.MonetaryAmountIncludingTaxes = in.MonetaryAmountIncludingTaxes
	orderResponseLineItemD.MaitCodeListVersion = in.MaitCodeListVersion
	orderResponseLineItemD.MaitCurrencyCode = in.MaitCurrencyCode
	orderResponseLineItemD.NetAmount = in.NetAmount
	orderResponseLineItemD.NaCodeListVersion = in.NaCodeListVersion
	orderResponseLineItemD.NaCurrencyCode = in.NaCurrencyCode
	orderResponseLineItemD.NetPrice = in.NetPrice
	orderResponseLineItemD.NpCodeListVersion = in.NpCodeListVersion
	orderResponseLineItemD.NpCurrencyCode = in.NpCurrencyCode
	orderResponseLineItemD.OrderResponseReasonCode = in.OrderResponseReasonCode
	orderResponseLineItemD.OriginalOrderLineItemNumber = in.OriginalOrderLineItemNumber
	orderResponseLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	orderResponseLineItemD.OrderResponseId = in.OrderResponseId

	orderResponseLineItemT := orderproto.OrderResponseLineItemT{}
	orderResponseLineItemT.DeliveryDateTime = common.TimeToTimestamp(deliveryDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	orderResponseLineItem := orderproto.OrderResponseLineItem{OrderResponseLineItemD: &orderResponseLineItemD, OrderResponseLineItemT: &orderResponseLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &orderResponseLineItem, nil
}

// insertOrderResponseLineItem - Insert OrderResponseLineItem details into database
func (o *OrderResponseService) insertOrderResponseLineItem(ctx context.Context, insertOrderResponseLineItemSQL string, orderResponseLineItem *orderproto.OrderResponseLineItem, userEmail string, requestID string) error {
	orderResponseLineItemTmp, err := o.crOrderResponseLineItemStruct(ctx, orderResponseLineItem, userEmail, requestID)
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertOrderResponseLineItemSQL, orderResponseLineItemTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		orderResponseLineItem.OrderResponseLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(orderResponseLineItem.OrderResponseLineItemD.Uuid4)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		orderResponseLineItem.OrderResponseLineItemD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crOrderResponseLineItemStruct - process OrderResponseLineItem details
func (o *OrderResponseService) crOrderResponseLineItemStruct(ctx context.Context, orderResponseLineItem *orderproto.OrderResponseLineItem, userEmail string, requestID string) (*orderstruct.OrderResponseLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(orderResponseLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(orderResponseLineItem.CrUpdTime.UpdatedAt)

	orderResponseLineItemT := new(orderstruct.OrderResponseLineItemT)
	orderResponseLineItemT.DeliveryDateTime = common.TimestampToTime(orderResponseLineItem.OrderResponseLineItemT.DeliveryDateTime)

	orderResponseLineItemTmp := orderstruct.OrderResponseLineItem{OrderResponseLineItemD: orderResponseLineItem.OrderResponseLineItemD, OrderResponseLineItemT: orderResponseLineItemT, CrUpdUser: orderResponseLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &orderResponseLineItemTmp, nil
}

// GetOrderResponseLineItems - GetOrderResponseLineItems
func (o *OrderResponseService) GetOrderResponseLineItems(ctx context.Context, inReq *orderproto.GetOrderResponseLineItemsRequest) (*orderproto.GetOrderResponseLineItemsResponse, error) {
	in := inReq.GetRequest
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := orderproto.GetOrderResponseRequest{}
	form.GetRequest = &getRequest

	orderRes, err := o.GetOrderResponse(ctx, &form)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponse := orderRes.OrderResponse

	orderResponseLineItems := []*orderproto.OrderResponseLineItem{}

	nselectOrderResponseLineItemsSQL := selectOrderResponseLineItemsSQL + ` where order_response_id = ?;`
	rows, err := o.DBService.DB.QueryxContext(ctx, nselectOrderResponseLineItemsSQL, orderResponse.OrderResponseD.Id)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		orderResponseLineItemTmp := orderstruct.OrderResponseLineItem{}
		err = rows.StructScan(&orderResponseLineItemTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		uuid4Str, err := common.UUIDBytesToStr(orderResponseLineItemTmp.OrderResponseLineItemD.Uuid4)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		orderResponseLineItemTmp.OrderResponseLineItemD.IdS = uuid4Str

		orderResponseLineItemT := orderproto.OrderResponseLineItemT{}
		orderResponseLineItemT.DeliveryDateTime = common.TimeToTimestamp(orderResponseLineItemTmp.OrderResponseLineItemT.DeliveryDateTime)

		crUpdTime := new(commonproto.CrUpdTime)
		crUpdTime.CreatedAt = common.TimeToTimestamp(orderResponseLineItemTmp.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimeToTimestamp(orderResponseLineItemTmp.CrUpdTime.UpdatedAt)

		orderResponseLineItem := orderproto.OrderResponseLineItem{OrderResponseLineItemD: orderResponseLineItemTmp.OrderResponseLineItemD, OrderResponseLineItemT: &orderResponseLineItemT, CrUpdUser: orderResponseLineItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

		orderResponseLineItems = append(orderResponseLineItems, &orderResponseLineItem)
	}

	orderRespLineItems := orderproto.GetOrderResponseLineItemsResponse{}
	orderRespLineItems.OrderResponseLineItems = orderResponseLineItems
	return &orderRespLineItems, nil
}
