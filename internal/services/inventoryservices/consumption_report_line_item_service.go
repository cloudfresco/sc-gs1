package inventoryservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	inventorystruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertConsumptionReportLineItemSQL = `insert into consumption_report_line_items
	    (uuid4,
	    consumed_quantity,
cq_measurement_unit_code,
cq_code_list_version,
line_item_number,
net_consumption_amount,
net_consumption_amount_currency,
ncac_code_list_version,
ncac_currency_code,
net_price,
net_price_currency,
np_code_list_version,
np_currency_code,
parent_line_item_number,
plan_bucket_size_code,
purchase_conditions,
consumption_report_id,
consumption_period_begin,
consumption_period_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
  :uuid4,
  :consumed_quantity,
  :cq_measurement_unit_code,
  :cq_code_list_version,
  :line_item_number,
  :net_consumption_amount,
  :net_consumption_amount_currency,
  :ncac_code_list_version,
  :ncac_currency_code,
  :net_price,
  :net_price_currency,
  :np_code_list_version,
  :np_currency_code,
  :parent_line_item_number,
  :plan_bucket_size_code,
  :purchase_conditions,
  :consumption_report_id,
  :consumption_period_begin,
  :consumption_period_end,
  :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

/*const selectConsumptionReportLineItemsSQL = `select
  id,
  uuid4,
  consumed_quantity,
  cq_measurement_unit_code,
  cq_code_list_version,
  line_item_number,
  net_consumption_amount,
  net_consumption_amount_currency,
  ncac_code_list_version,
  ncac_currency_code,
  net_price,
  net_price_currency,
  np_code_list_version,
  np_currency_code,
  parent_line_item_number,
  plan_bucket_size_code,
  purchase_conditions,
  consumption_report_id,
  consumption_period_begin,
  consumption_period_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from consumption_report_line_items`*/

// CreateConsumptionReportLineItem - Create ConsumptionReportLineItem
func (invs *InventoryService) CreateConsumptionReportLineItem(ctx context.Context, in *inventoryproto.CreateConsumptionReportLineItemRequest) (*inventoryproto.CreateConsumptionReportLineItemResponse, error) {
	consumptionReportLineItem, err := invs.ProcessConsumptionReportLineItemRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertConsumptionReportLineItem(ctx, insertConsumptionReportLineItemSQL, consumptionReportLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionReportLineItemResponse := inventoryproto.CreateConsumptionReportLineItemResponse{}
	consumptionReportLineItemResponse.ConsumptionReportLineItem = consumptionReportLineItem
	return &consumptionReportLineItemResponse, nil
}

// ProcessConsumptionReportLineItemRequest - ProcessConsumptionReportLineItemRequest
func (invs *InventoryService) ProcessConsumptionReportLineItemRequest(ctx context.Context, in *inventoryproto.CreateConsumptionReportLineItemRequest) (*inventoryproto.ConsumptionReportLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	consumptionPeriodBegin, err := time.Parse(common.Layout, in.ConsumptionPeriodBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionPeriodEnd, err := time.Parse(common.Layout, in.ConsumptionPeriodEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionReportLineItemD := inventoryproto.ConsumptionReportLineItemD{}
	consumptionReportLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	consumptionReportLineItemD.ConsumedQuantity = in.ConsumedQuantity
	consumptionReportLineItemD.CQMeasurementUnitCode = in.CQMeasurementUnitCode
	consumptionReportLineItemD.CQCodeListVersion = in.CQCodeListVersion
	consumptionReportLineItemD.LineItemNumber = in.LineItemNumber
	consumptionReportLineItemD.NetConsumptionAmount = in.NetConsumptionAmount
	consumptionReportLineItemD.NCACCodeListVersion = in.NCACCodeListVersion
	consumptionReportLineItemD.NCACCurrencyCode = in.NCACCurrencyCode
	consumptionReportLineItemD.NPCodeListVersion = in.NPCodeListVersion
	consumptionReportLineItemD.NPCurrencyCode = in.NPCurrencyCode
	consumptionReportLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	consumptionReportLineItemD.PlanBucketSizeCode = in.PlanBucketSizeCode
	consumptionReportLineItemD.PurchaseConditions = in.PurchaseConditions
	consumptionReportLineItemD.ConsumptionReportId = in.ConsumptionReportId

	netPriceCurrency, err := invs.CurrencyService.GetCurrency(ctx, in.NetPriceCurrency)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	netPriceMinor, err := common.ParseAmountString(in.NetPrice, netPriceCurrency)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionReportLineItemD.NetPriceCurrency = netPriceCurrency.Code
	consumptionReportLineItemD.NetPrice = netPriceMinor
	consumptionReportLineItemD.NetPriceString = common.FormatAmountString(netPriceMinor, netPriceCurrency)


	consumptionReportLineItemT := inventoryproto.ConsumptionReportLineItemT{}
	consumptionReportLineItemT.ConsumptionPeriodBegin = common.TimeToTimestamp(consumptionPeriodBegin.UTC().Truncate(time.Second))
	consumptionReportLineItemT.ConsumptionPeriodEnd = common.TimeToTimestamp(consumptionPeriodEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	consumptionReportLineItem := inventoryproto.ConsumptionReportLineItem{ConsumptionReportLineItemD: &consumptionReportLineItemD, ConsumptionReportLineItemT: &consumptionReportLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &consumptionReportLineItem, nil
}

// insertConsumptionReportLineItem - Insert ConsumptionReportLineItem details into database
func (invs *InventoryService) insertConsumptionReportLineItem(ctx context.Context, insertConsumptionReportLineItemSQL string, consumptionReportLineItem *inventoryproto.ConsumptionReportLineItem, userEmail string, requestID string) error {
	consumptionReportLineItemTmp, err := invs.crConsumptionReportLineItemStruct(ctx, consumptionReportLineItem, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertConsumptionReportLineItemSQL, consumptionReportLineItemTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consumptionReportLineItem.ConsumptionReportLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(consumptionReportLineItem.ConsumptionReportLineItemD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consumptionReportLineItem.ConsumptionReportLineItemD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crConsumptionReportLineItemStruct - process ConsumptionReportLineItem details
func (invs *InventoryService) crConsumptionReportLineItemStruct(ctx context.Context, consumptionReportLineItem *inventoryproto.ConsumptionReportLineItem, userEmail string, requestID string) (*inventorystruct.ConsumptionReportLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(consumptionReportLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(consumptionReportLineItem.CrUpdTime.UpdatedAt)

	consumptionReportLineItemT := new(inventorystruct.ConsumptionReportLineItemT)
	consumptionReportLineItemT.ConsumptionPeriodBegin = common.TimestampToTime(consumptionReportLineItem.ConsumptionReportLineItemT.ConsumptionPeriodBegin)
	consumptionReportLineItemT.ConsumptionPeriodEnd = common.TimestampToTime(consumptionReportLineItem.ConsumptionReportLineItemT.ConsumptionPeriodEnd)

	consumptionReportLineItemTmp := inventorystruct.ConsumptionReportLineItem{ConsumptionReportLineItemD: consumptionReportLineItem.ConsumptionReportLineItemD, ConsumptionReportLineItemT: consumptionReportLineItemT, CrUpdUser: consumptionReportLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &consumptionReportLineItemTmp, nil
}
