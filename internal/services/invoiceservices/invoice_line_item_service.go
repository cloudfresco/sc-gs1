package invoiceservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	invoicestruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInvoiceLineItemSQL = `insert into invoice_line_items
	    (uuid4,
amount_exclusive_allowances_charges,
aeac_code_list_version,
aeac_currency_code,
amount_inclusive_allowances_charges,
aiac_code_list_version,
aiac_currency_code,
credit_line_indicator,
credit_reason,
delivered_quantity,
dq_measurement_unit_code,
dq_code_list_version,
excluded_from_payment_discount_indicator,
extension,
free_goods_quantity,
fgq_measurement_unit_code,
fgq_code_list_version,
invoiced_quantity,
iq_measurement_unit_code,
iq_code_list_version,
item_price_base_quantity,
ipbq_measurement_unit_code,
ipbq_code_list_version,
item_price_exclusive_allowances_charges,
ipeac_code_list_version,
ipeac_currency_code,
item_price_inclusive_allowances_charges,
ipiac_code_list_version,
ipiac_currency_code,
legally_fixed_retail_price,
lfrp_code_list_version,
lfrp_currency_code,
line_item_number,
margin_scheme_information,
owenrship_prior_to_payment,
parent_line_item_number,
recommended_retail_price,
rrp_code_list_version,
rrp_currency_code,
retail_price_excluding_excise,
rpee_code_list_version,
rpee_currency_code,
total_ordered_quantity,
toq_measurement_unit_code,
toq_code_list_version,
consumption_report,
contract,
delivery_note,
despatch_advice,
energy_quantity,
inventory_location_from,
inventory_location_to,
promotional_deal,
purchase_conditions,
purchase_order,
receiving_advice,
returnable_asset_identification,
sales_order,
ship_from,
ship_to,
trade_agreement,
invoice_id,
transfer_of_ownership_date,
actual_delivery_date,
servicetime_period_line_level_begin,
servicetime_period_line_level_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
  :uuid4,
:amount_exclusive_allowances_charges,
:aeac_code_list_version,
:aeac_currency_code,
:amount_inclusive_allowances_charges,
:aiac_code_list_version,
:aiac_currency_code,
:credit_line_indicator,
:credit_reason,
:delivered_quantity,
:dq_measurement_unit_code,
:dq_code_list_version,
:excluded_from_payment_discount_indicator,
:extension,
:free_goods_quantity,
:fgq_measurement_unit_code,
:fgq_code_list_version,
:invoiced_quantity,
:iq_measurement_unit_code,
:iq_code_list_version,
:item_price_base_quantity,
:ipbq_measurement_unit_code,
:ipbq_code_list_version,
:item_price_exclusive_allowances_charges,
:ipeac_code_list_version,
:ipeac_currency_code,
:item_price_inclusive_allowances_charges,
:ipiac_code_list_version,
:ipiac_currency_code,
:legally_fixed_retail_price,
:lfrp_code_list_version,
:lfrp_currency_code,
:line_item_number,
:margin_scheme_information,
:owenrship_prior_to_payment,
:parent_line_item_number,
:recommended_retail_price,
:rrp_code_list_version,
:rrp_currency_code,
:retail_price_excluding_excise,
:rpee_code_list_version,
:rpee_currency_code,
:total_ordered_quantity,
:toq_measurement_unit_code,
:toq_code_list_version,
:consumption_report,
:contract,
:delivery_note,
:despatch_advice,
:energy_quantity,
:inventory_location_from,
:inventory_location_to,
:promotional_deal,
:purchase_conditions,
:purchase_order,
:receiving_advice,
:returnable_asset_identification,
:sales_order,
:ship_from,
:ship_to,
:trade_agreement,
:invoice_id,
:transfer_of_ownership_date,
:actual_delivery_date,
:servicetime_period_line_level_begin,
:servicetime_period_line_level_end,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectInvoiceLineItemsSQL = `select
    id,
    uuid4,
amount_exclusive_allowances_charges,
aeac_code_list_version,
aeac_currency_code,
amount_inclusive_allowances_charges,
aiac_code_list_version,
aiac_currency_code,
credit_line_indicator,
credit_reason,
delivered_quantity,
dq_measurement_unit_code,
dq_code_list_version,
excluded_from_payment_discount_indicator,
extension,
free_goods_quantity,
fgq_measurement_unit_code,
fgq_code_list_version,
invoiced_quantity,
iq_measurement_unit_code,
iq_code_list_version,
item_price_base_quantity,
ipbq_measurement_unit_code,
ipbq_code_list_version,
item_price_exclusive_allowances_charges,
ipeac_code_list_version,
ipeac_currency_code,
item_price_inclusive_allowances_charges,
ipiac_code_list_version,
ipiac_currency_code,
legally_fixed_retail_price,
lfrp_code_list_version,
lfrp_currency_code,
line_item_number,
margin_scheme_information,
owenrship_prior_to_payment,
parent_line_item_number,
recommended_retail_price,
rrp_code_list_version,
rrp_currency_code,
retail_price_excluding_excise,
rpee_code_list_version,
rpee_currency_code,
total_ordered_quantity,
toq_measurement_unit_code,
toq_code_list_version,
consumption_report,
contract,
delivery_note,
despatch_advice,
energy_quantity,
inventory_location_from,
inventory_location_to,
promotional_deal,
purchase_conditions,
purchase_order,
receiving_advice,
returnable_asset_identification,
sales_order,
ship_from,
ship_to,
trade_agreement,
invoice_id,
transfer_of_ownership_date,
actual_delivery_date,
servicetime_period_line_level_begin,
servicetime_period_line_level_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from invoice_line_items`

// CreateInvoiceLineItem - Create InvoiceLineItem
func (invs *InvoiceService) CreateInvoiceLineItem(ctx context.Context, in *invoiceproto.CreateInvoiceLineItemRequest) (*invoiceproto.CreateInvoiceLineItemResponse, error) {
	invoiceLineItem, err := invs.ProcessInvoiceLineItemRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInvoiceLineItem(ctx, insertInvoiceLineItemSQL, invoiceLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceLineItemResponse := invoiceproto.CreateInvoiceLineItemResponse{}
	invoiceLineItemResponse.InvoiceLineItem = invoiceLineItem
	return &invoiceLineItemResponse, nil
}

// ProcessInvoiceLineItemRequest - ProcessInvoiceLineItemRequest
func (invs *InvoiceService) ProcessInvoiceLineItemRequest(ctx context.Context, in *invoiceproto.CreateInvoiceLineItemRequest) (*invoiceproto.InvoiceLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	transferOfOwnershipDate, err := time.Parse(common.Layout, in.TransferOfOwnershipDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	actualDeliveryDate, err := time.Parse(common.Layout, in.ActualDeliveryDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	servicetimePeriodLineLevelBegin, err := time.Parse(common.Layout, in.ServicetimePeriodLineLevelBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	servicetimePeriodLineLevelEnd, err := time.Parse(common.Layout, in.ServicetimePeriodLineLevelEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceLineItemD := invoiceproto.InvoiceLineItemD{}
	invoiceLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	invoiceLineItemD.AmountExclusiveAllowancesCharges = in.AmountExclusiveAllowancesCharges
	invoiceLineItemD.AeacCodeListVersion = in.AeacCodeListVersion
	invoiceLineItemD.AeacCurrencyCode = in.AeacCurrencyCode
	invoiceLineItemD.AmountInclusiveAllowancesCharges = in.AmountInclusiveAllowancesCharges
	invoiceLineItemD.AiacCodeListVersion = in.AiacCodeListVersion
	invoiceLineItemD.AiacCurrencyCode = in.AiacCurrencyCode
	invoiceLineItemD.CreditLineIndicator = in.CreditLineIndicator
	invoiceLineItemD.CreditReason = in.CreditReason
	invoiceLineItemD.DeliveredQuantity = in.DeliveredQuantity
	invoiceLineItemD.DqMeasurementUnitCode = in.DqMeasurementUnitCode
	invoiceLineItemD.DqCodeListVersion = in.DqCodeListVersion
	invoiceLineItemD.ExcludedFromPaymentDiscountIndicator = in.ExcludedFromPaymentDiscountIndicator
	invoiceLineItemD.Extension = in.Extension
	invoiceLineItemD.FreeGoodsQuantity = in.FreeGoodsQuantity
	invoiceLineItemD.FgqMeasurementUnitCode = in.FgqMeasurementUnitCode
	invoiceLineItemD.FgqCodeListVersion = in.FgqCodeListVersion
	invoiceLineItemD.InvoicedQuantity = in.InvoicedQuantity
	invoiceLineItemD.IqMeasurementUnitCode = in.IqMeasurementUnitCode
	invoiceLineItemD.IqCodeListVersion = in.IqCodeListVersion
	invoiceLineItemD.ItemPriceBaseQuantity = in.ItemPriceBaseQuantity
	invoiceLineItemD.IpbqMeasurementUnitCode = in.IpbqMeasurementUnitCode
	invoiceLineItemD.IpbqCodeListVersion = in.IpbqCodeListVersion
	invoiceLineItemD.ItemPriceExclusiveAllowancesCharges = in.ItemPriceExclusiveAllowancesCharges
	invoiceLineItemD.IpeacCodeListVersion = in.IpeacCodeListVersion
	invoiceLineItemD.IpeacCurrencyCode = in.IpeacCurrencyCode
	invoiceLineItemD.ItemPriceInclusiveAllowancesCharges = in.ItemPriceInclusiveAllowancesCharges
	invoiceLineItemD.IpiacCodeListVersion = in.IpiacCodeListVersion
	invoiceLineItemD.IpiacCurrencyCode = in.IpiacCurrencyCode
	invoiceLineItemD.LegallyFixedRetailPrice = in.LegallyFixedRetailPrice
	invoiceLineItemD.LfrpCodeListVersion = in.LfrpCodeListVersion
	invoiceLineItemD.LfrpCurrencyCode = in.LfrpCurrencyCode
	invoiceLineItemD.LineItemNumber = in.LineItemNumber
	invoiceLineItemD.MarginSchemeInformation = in.MarginSchemeInformation
	invoiceLineItemD.OwenrshipPriorToPayment = in.OwenrshipPriorToPayment
	invoiceLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	invoiceLineItemD.RecommendedRetailPrice = in.RecommendedRetailPrice
	invoiceLineItemD.RrpCodeListVersion = in.RrpCodeListVersion
	invoiceLineItemD.RrpCurrencyCode = in.RrpCurrencyCode
	invoiceLineItemD.RetailPriceExcludingExcise = in.RetailPriceExcludingExcise
	invoiceLineItemD.RpeeCodeListVersion = in.RpeeCodeListVersion
	invoiceLineItemD.RpeeCurrencyCode = in.RpeeCurrencyCode
	invoiceLineItemD.TotalOrderedQuantity = in.TotalOrderedQuantity
	invoiceLineItemD.ToqMeasurementUnitCode = in.ToqMeasurementUnitCode
	invoiceLineItemD.ToqCodeListVersion = in.ToqCodeListVersion
	invoiceLineItemD.ConsumptionReport = in.ConsumptionReport
	invoiceLineItemD.Contract = in.Contract
	invoiceLineItemD.DeliveryNote = in.DeliveryNote
	invoiceLineItemD.DespatchAdvice = in.DespatchAdvice
	invoiceLineItemD.EnergyQuantity = in.EnergyQuantity
	invoiceLineItemD.InventoryLocationFrom = in.InventoryLocationFrom
	invoiceLineItemD.InventoryLocationTo = in.InventoryLocationTo
	invoiceLineItemD.PromotionalDeal = in.PromotionalDeal
	invoiceLineItemD.PurchaseConditions = in.PurchaseConditions
	invoiceLineItemD.PurchaseOrder = in.PurchaseOrder
	invoiceLineItemD.ReceivingAdvice = in.ReceivingAdvice
	invoiceLineItemD.ReturnableAssetIdentification = in.ReturnableAssetIdentification
	invoiceLineItemD.SalesOrder = in.SalesOrder
	invoiceLineItemD.ShipFrom = in.ShipFrom
	invoiceLineItemD.ShipTo = in.ShipTo
	invoiceLineItemD.TradeAgreement = in.TradeAgreement
	invoiceLineItemD.InvoiceId = in.InvoiceId

	invoiceLineItemT := invoiceproto.InvoiceLineItemT{}
	invoiceLineItemT.TransferOfOwnershipDate = common.TimeToTimestamp(transferOfOwnershipDate.UTC().Truncate(time.Second))
	invoiceLineItemT.ActualDeliveryDate = common.TimeToTimestamp(actualDeliveryDate.UTC().Truncate(time.Second))
	invoiceLineItemT.ServicetimePeriodLineLevelBegin = common.TimeToTimestamp(servicetimePeriodLineLevelBegin.UTC().Truncate(time.Second))
	invoiceLineItemT.ServicetimePeriodLineLevelEnd = common.TimeToTimestamp(servicetimePeriodLineLevelEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	invoiceLineItem := invoiceproto.InvoiceLineItem{InvoiceLineItemD: &invoiceLineItemD, InvoiceLineItemT: &invoiceLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &invoiceLineItem, nil
}

// insertInvoiceLineItem - Insert InvoiceLineItem details into database
func (invs *InvoiceService) insertInvoiceLineItem(ctx context.Context, insertInvoiceLineItemSQL string, invoiceLineItem *invoiceproto.InvoiceLineItem, userEmail string, requestID string) error {
	invoiceLineItemTmp, err := invs.crInvoiceLineItemStruct(ctx, invoiceLineItem, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInvoiceLineItemSQL, invoiceLineItemTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoiceLineItem.InvoiceLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(invoiceLineItem.InvoiceLineItemD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoiceLineItem.InvoiceLineItemD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crInvoiceLineItemStruct - process InvoiceLineItem details
func (invs *InvoiceService) crInvoiceLineItemStruct(ctx context.Context, invoiceLineItem *invoiceproto.InvoiceLineItem, userEmail string, requestID string) (*invoicestruct.InvoiceLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(invoiceLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(invoiceLineItem.CrUpdTime.UpdatedAt)

	invoiceLineItemT := new(invoicestruct.InvoiceLineItemT)
	invoiceLineItemT.TransferOfOwnershipDate = common.TimestampToTime(invoiceLineItem.InvoiceLineItemT.TransferOfOwnershipDate)
	invoiceLineItemT.ActualDeliveryDate = common.TimestampToTime(invoiceLineItem.InvoiceLineItemT.ActualDeliveryDate)
	invoiceLineItemT.ServicetimePeriodLineLevelBegin = common.TimestampToTime(invoiceLineItem.InvoiceLineItemT.ServicetimePeriodLineLevelBegin)
	invoiceLineItemT.ServicetimePeriodLineLevelEnd = common.TimestampToTime(invoiceLineItem.InvoiceLineItemT.ServicetimePeriodLineLevelEnd)

	invoiceLineItemTmp := invoicestruct.InvoiceLineItem{InvoiceLineItemD: invoiceLineItem.InvoiceLineItemD, InvoiceLineItemT: invoiceLineItemT, CrUpdUser: invoiceLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &invoiceLineItemTmp, nil
}

// GetInvoiceLineItems - GetInvoiceLineItems
func (invs *InvoiceService) GetInvoiceLineItems(ctx context.Context, inReq *invoiceproto.GetInvoiceLineItemsRequest) (*invoiceproto.GetInvoiceLineItemsResponse, error) {
	in := inReq.GetRequest
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := invoiceproto.GetInvoiceRequest{}
	form.GetRequest = &getRequest

	invoiceResponse, err := invs.GetInvoice(ctx, &form)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoice := invoiceResponse.Invoice
	invoiceLineItems := []*invoiceproto.InvoiceLineItem{}

	nselectInvoiceLineItemsSQL := selectInvoiceLineItemsSQL + ` where invoice_id = ?;`
	rows, err := invs.DBService.DB.QueryxContext(ctx, nselectInvoiceLineItemsSQL, invoice.InvoiceD.Id)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		invoiceLineItemTmp := invoicestruct.InvoiceLineItem{}
		err = rows.StructScan(&invoiceLineItemTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		invoiceLineItemT := new(invoiceproto.InvoiceLineItemT)
		invoiceLineItemT.TransferOfOwnershipDate = common.TimeToTimestamp(invoiceLineItemTmp.InvoiceLineItemT.TransferOfOwnershipDate)
		invoiceLineItemT.ActualDeliveryDate = common.TimeToTimestamp(invoiceLineItemTmp.InvoiceLineItemT.ActualDeliveryDate)
		invoiceLineItemT.ServicetimePeriodLineLevelBegin = common.TimeToTimestamp(invoiceLineItemTmp.InvoiceLineItemT.ServicetimePeriodLineLevelBegin)
		invoiceLineItemT.ServicetimePeriodLineLevelEnd = common.TimeToTimestamp(invoiceLineItemTmp.InvoiceLineItemT.ServicetimePeriodLineLevelEnd)

		uuid4Str, err := common.UUIDBytesToStr(invoiceLineItemTmp.InvoiceLineItemD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		invoiceLineItemTmp.InvoiceLineItemD.IdS = uuid4Str

		crUpdTime := new(commonproto.CrUpdTime)
		crUpdTime.CreatedAt = common.TimeToTimestamp(invoiceLineItemTmp.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimeToTimestamp(invoiceLineItemTmp.CrUpdTime.UpdatedAt)

		invoiceLineItem := invoiceproto.InvoiceLineItem{InvoiceLineItemD: invoiceLineItemTmp.InvoiceLineItemD, InvoiceLineItemT: invoiceLineItemT, CrUpdUser: invoiceLineItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

		invoiceLineItems = append(invoiceLineItems, &invoiceLineItem)
	}
	invoiceLineItemsResponse := invoiceproto.GetInvoiceLineItemsResponse{}
	invoiceLineItemsResponse.InvoiceLineItems = invoiceLineItems
	return &invoiceLineItemsResponse, nil
}
