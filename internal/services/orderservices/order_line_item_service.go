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

const insertOrderLineItemSQL = `insert into order_line_items
	    (uuid4,
	    extension,
free_goods_quantity,
fgq_measurement_unit_code,
fgq_code_list_version,
item_price_base_quantity,
ipbq_measurement_unit_code,
ipbq_code_list_version,
item_source_code,
line_item_action_code,
line_item_number,
list_price,
lp_code_list_version,
lp_currency_code,
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
order_instruction_code,
order_line_item_instruction_code,
order_line_item_priority,
parent_line_item_number,
recommended_retail_price,
requested_quantity,
rq_measurement_unit_code,
rq_code_list_version,
return_reason_code,
contract,
customer_document_reference,
delivery_date_according_to_schedule,
despatch_advice,
material_specification,
order_line_item_contact,
preferred_manufacturer,
promotional_deal,
purchase_conditions,
returnable_asset_identification,
order_id,
latest_delivery_date,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
  :uuid4,
  :extension,
:free_goods_quantity,
:fgq_measurement_unit_code,
:fgq_code_list_version,
:item_price_base_quantity,
:ipbq_measurement_unit_code,
:ipbq_code_list_version,
:item_source_code,
:line_item_action_code,
:line_item_number,
:list_price,
:lp_code_list_version,
:lp_currency_code,
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
:order_instruction_code,
:order_line_item_instruction_code,
:order_line_item_priority,
:parent_line_item_number,
:recommended_retail_price,
:requested_quantity,
:rq_measurement_unit_code,
:rq_code_list_version,
:return_reason_code,
:contract,
:customer_document_reference,
:delivery_date_according_to_schedule,
:despatch_advice,
:material_specification,
:order_line_item_contact,
:preferred_manufacturer,
:promotional_deal,
:purchase_conditions,
:returnable_asset_identification,
:order_id,
:latest_delivery_date,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectOrderLineItemsSQL = `select
    id,
    uuid4,
    extension,
free_goods_quantity,
fgq_measurement_unit_code,
fgq_code_list_version,
item_price_base_quantity,
ipbq_measurement_unit_code,
ipbq_code_list_version,
item_source_code,
line_item_action_code,
line_item_number,
list_price,
lp_code_list_version,
lp_currency_code,
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
order_instruction_code,
order_line_item_instruction_code,
order_line_item_priority,
parent_line_item_number,
recommended_retail_price,
requested_quantity,
rq_measurement_unit_code,
rq_code_list_version,
return_reason_code,
contract,
customer_document_reference,
delivery_date_according_to_schedule,
despatch_advice,
material_specification,
order_line_item_contact,
preferred_manufacturer,
promotional_deal,
purchase_conditions,
returnable_asset_identification,
order_id,
latest_delivery_date,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from order_line_items`

// CreateOrderLineItem - Create OrderLineItem
func (o *OrderService) CreateOrderLineItem(ctx context.Context, in *orderproto.CreateOrderLineItemRequest) (*orderproto.CreateOrderLineItemResponse, error) {
	orderLineItem, err := o.ProcessOrderLineItemRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderLineItem(ctx, insertOrderLineItemSQL, orderLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderLineItemResponse := orderproto.CreateOrderLineItemResponse{}
	orderLineItemResponse.OrderLineItem = orderLineItem
	return &orderLineItemResponse, nil
}

// ProcessOrderLineItemRequest - ProcessOrderLineItemRequest
func (o *OrderService) ProcessOrderLineItemRequest(ctx context.Context, in *orderproto.CreateOrderLineItemRequest) (*orderproto.OrderLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, o.UserServiceClient)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	latestDeliveryDate, err := time.Parse(common.Layout, in.LatestDeliveryDate)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderLineItemD := orderproto.OrderLineItemD{}
	orderLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	orderLineItemD.Extension = in.Extension
	orderLineItemD.FreeGoodsQuantity = in.FreeGoodsQuantity
	orderLineItemD.FgqMeasurementUnitCode = in.FgqMeasurementUnitCode
	orderLineItemD.FgqCodeListVersion = in.FgqCodeListVersion
	orderLineItemD.ItemPriceBaseQuantity = in.ItemPriceBaseQuantity
	orderLineItemD.IpbqMeasurementUnitCode = in.IpbqMeasurementUnitCode
	orderLineItemD.IpbqCodeListVersion = in.IpbqCodeListVersion
	orderLineItemD.ItemSourceCode = in.ItemSourceCode
	orderLineItemD.LineItemActionCode = in.LineItemActionCode
	orderLineItemD.LineItemNumber = in.LineItemNumber
	orderLineItemD.ListPrice = in.ListPrice
	orderLineItemD.LpCodeListVersion = in.LpCodeListVersion
	orderLineItemD.LpCurrencyCode = in.LpCurrencyCode
	orderLineItemD.MonetaryAmountExcludingTaxes = in.MonetaryAmountExcludingTaxes
	orderLineItemD.MaetCodeListVersion = in.MaetCodeListVersion
	orderLineItemD.MaetCurrencyCode = in.MaetCurrencyCode
	orderLineItemD.MonetaryAmountIncludingTaxes = in.MonetaryAmountIncludingTaxes
	orderLineItemD.MaitCodeListVersion = in.MaitCodeListVersion
	orderLineItemD.MaitCurrencyCode = in.MaitCurrencyCode
	orderLineItemD.NetAmount = in.NetAmount
	orderLineItemD.NaCodeListVersion = in.NaCodeListVersion
	orderLineItemD.NaCurrencyCode = in.NaCurrencyCode
	orderLineItemD.NetPrice = in.NetPrice
	orderLineItemD.NpCodeListVersion = in.NpCodeListVersion
	orderLineItemD.NpCurrencyCode = in.NpCurrencyCode
	orderLineItemD.OrderInstructionCode = in.OrderInstructionCode
	orderLineItemD.OrderLineItemInstructionCode = in.OrderLineItemInstructionCode
	orderLineItemD.OrderLineItemPriority = in.OrderLineItemPriority
	orderLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	orderLineItemD.RecommendedRetailPrice = in.RecommendedRetailPrice
	orderLineItemD.RequestedQuantity = in.RequestedQuantity
	orderLineItemD.RqMeasurementUnitCode = in.RqMeasurementUnitCode
	orderLineItemD.RqCodeListVersion = in.RqCodeListVersion
	orderLineItemD.ReturnReasonCode = in.ReturnReasonCode
	orderLineItemD.Contract = in.Contract
	orderLineItemD.CustomerDocumentReference = in.CustomerDocumentReference
	orderLineItemD.DeliveryDateAccordingToSchedule = in.DeliveryDateAccordingToSchedule
	orderLineItemD.DespatchAdvice = in.DespatchAdvice
	orderLineItemD.MaterialSpecification = in.MaterialSpecification
	orderLineItemD.OrderLineItemContact = in.OrderLineItemContact
	orderLineItemD.PreferredManufacturer = in.PreferredManufacturer
	orderLineItemD.PromotionalDeal = in.PromotionalDeal
	orderLineItemD.PurchaseConditions = in.PurchaseConditions
	orderLineItemD.ReturnableAssetIdentification = in.ReturnableAssetIdentification
	orderLineItemD.OrderId = in.OrderId

	orderLineItemT := orderproto.OrderLineItemT{}
	orderLineItemT.LatestDeliveryDate = common.TimeToTimestamp(latestDeliveryDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	orderLineItem := orderproto.OrderLineItem{OrderLineItemD: &orderLineItemD, OrderLineItemT: &orderLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &orderLineItem, nil
}

// insertOrderLineItem - Insert OrderLineItem details into database
func (o *OrderService) insertOrderLineItem(ctx context.Context, insertOrderLineItemSQL string, orderLineItem *orderproto.OrderLineItem, userEmail string, requestID string) error {
	orderLineItemTmp, err := o.crOrderLineItemStruct(ctx, orderLineItem, userEmail, requestID)
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertOrderLineItemSQL, orderLineItemTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		orderLineItem.OrderLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(orderLineItem.OrderLineItemD.Uuid4)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		orderLineItem.OrderLineItemD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crOrderLineItemStruct - process OrderLineItem details
func (o *OrderService) crOrderLineItemStruct(ctx context.Context, orderLineItem *orderproto.OrderLineItem, userEmail string, requestID string) (*orderstruct.OrderLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(orderLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(orderLineItem.CrUpdTime.UpdatedAt)

	orderLineItemT := new(orderstruct.OrderLineItemT)
	orderLineItemT.LatestDeliveryDate = common.TimestampToTime(orderLineItem.OrderLineItemT.LatestDeliveryDate)

	orderLineItemTmp := orderstruct.OrderLineItem{OrderLineItemD: orderLineItem.OrderLineItemD, OrderLineItemT: orderLineItemT, CrUpdUser: orderLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &orderLineItemTmp, nil
}

// GetOrderLineItems - GetOrderLineItems
func (o *OrderService) GetOrderLineItems(ctx context.Context, inReq *orderproto.GetOrderLineItemsRequest) (*orderproto.GetOrderLineItemsResponse, error) {
	in := inReq.GetRequest
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := orderproto.GetOrderRequest{}
	form.GetRequest = &getRequest

	orderResponse, err := o.GetOrder(ctx, &form)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	order := orderResponse.Order
	orderLineItems := []*orderproto.OrderLineItem{}

	nselectOrderLineItemsSQL := selectOrderLineItemsSQL + ` where order_id = ?;`
	rows, err := o.DBService.DB.QueryxContext(ctx, nselectOrderLineItemsSQL, order.OrderD.Id)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		orderLineItemTmp := orderstruct.OrderLineItem{}
		err = rows.StructScan(&orderLineItemTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		orderLineItemT := orderproto.OrderLineItemT{}
		orderLineItemT.LatestDeliveryDate = common.TimeToTimestamp(orderLineItemTmp.OrderLineItemT.LatestDeliveryDate)

		uuid4Str, err := common.UUIDBytesToStr(orderLineItemTmp.OrderLineItemD.Uuid4)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		orderLineItemTmp.OrderLineItemD.IdS = uuid4Str

		crUpdTime := new(commonproto.CrUpdTime)
		crUpdTime.CreatedAt = common.TimeToTimestamp(orderLineItemTmp.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimeToTimestamp(orderLineItemTmp.CrUpdTime.UpdatedAt)

		orderLineItem := orderproto.OrderLineItem{OrderLineItemD: orderLineItemTmp.OrderLineItemD, OrderLineItemT: &orderLineItemT, CrUpdUser: orderLineItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

		orderLineItems = append(orderLineItems, &orderLineItem)
	}
	orderLineItemsResponse := orderproto.GetOrderLineItemsResponse{}
	orderLineItemsResponse.OrderLineItems = orderLineItems
	return &orderLineItemsResponse, nil
}
