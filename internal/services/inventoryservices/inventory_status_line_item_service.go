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

const insertInventoryStatusLineItemSQL = `insert into inventory_status_line_items
	    (uuid4,
handling_unit_type,
inventory_unit_cost,
iuc_code_list_version,
iuc_currency_code,
line_item_number,
parent_line_item_number,
inventory_status_owner,
inventory_sub_location_id,
logistic_unit_identification,
returnable_asset_identification,
inventory_report_type_code,
structure_type_code,
inventory_report_identification,
inventory_reporting_party,
inventory_report_to_party,
inventory_item_location_information_id,
inventory_report_id,
first_in_first_out_date_time_begin,
first_in_first_out_date_time_end,
inventory_date_time_begin,
inventory_date_time_end,
reporting_period_begin,
reporting_period_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
:uuid4,
:handling_unit_type,
:inventory_unit_cost,
:iuc_code_list_version,
:iuc_currency_code,
:line_item_number,
:parent_line_item_number,
:inventory_status_owner,
:inventory_sub_location_id,
:logistic_unit_identification,
:returnable_asset_identification,
:inventory_report_type_code,
:structure_type_code,
:inventory_report_identification,
:inventory_reporting_party,
:inventory_report_to_party,
:inventory_item_location_information_id,
:inventory_report_id,
:first_in_first_out_date_time_begin,
:first_in_first_out_date_time_end,
:inventory_date_time_begin,
:inventory_date_time_end,
:reporting_period_begin,
:reporting_period_end,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectInventoryStatusLineItemsSQL = `select
  id,
  uuid4,
  handling_unit_type,
  inventory_unit_cost,
  iuc_code_list_version,
  iuc_currency_code,
  line_item_number,
  parent_line_item_number,
  inventory_status_owner,
  inventory_sub_location_id,
  logistic_unit_identification,
  returnable_asset_identification,
  inventory_report_type_code,
  structure_type_code,
  inventory_report_identification,
  inventory_reporting_party,
  inventory_report_to_party,
  inventory_item_location_information_id,
  inventory_report_id,
  first_in_first_out_date_time_begin,
  first_in_first_out_date_time_end,
  inventory_date_time_begin,
  inventory_date_time_end,
  reporting_period_begin,
  reporting_period_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from inventory_status_line_items`*/

// CreateInventoryStatusLineItem - Create InventoryStatusLineItem
func (invs *InventoryService) CreateInventoryStatusLineItem(ctx context.Context, in *inventoryproto.CreateInventoryStatusLineItemRequest) (*inventoryproto.CreateInventoryStatusLineItemResponse, error) {
	inventoryStatusLineItem, err := invs.ProcessInventoryStatusLineItemRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInventoryStatusLineItem(ctx, insertInventoryStatusLineItemSQL, inventoryStatusLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryStatusLineItemResponse := inventoryproto.CreateInventoryStatusLineItemResponse{}
	inventoryStatusLineItemResponse.InventoryStatusLineItem = inventoryStatusLineItem
	return &inventoryStatusLineItemResponse, nil
}

// ProcessInventoryStatusLineItemRequest - ProcessInventoryStatusLineItemRequest
func (invs *InventoryService) ProcessInventoryStatusLineItemRequest(ctx context.Context, in *inventoryproto.CreateInventoryStatusLineItemRequest) (*inventoryproto.InventoryStatusLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	firstInFirstOutDateTimeBegin, err := time.Parse(common.Layout, in.FirstInFirstOutDateTimeBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	firstInFirstOutDateTimeEnd, err := time.Parse(common.Layout, in.FirstInFirstOutDateTimeEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryDateTimeBegin, err := time.Parse(common.Layout, in.InventoryDateTimeBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryDateTimeEnd, err := time.Parse(common.Layout, in.InventoryDateTimeEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	reportingPeriodBegin, err := time.Parse(common.Layout, in.ReportingPeriodBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	reportingPeriodEnd, err := time.Parse(common.Layout, in.ReportingPeriodEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryStatusLineItemD := inventoryproto.InventoryStatusLineItemD{}
	inventoryStatusLineItemD.HandlingUnitType = in.HandlingUnitType
	inventoryStatusLineItemD.InventoryUnitCost = in.InventoryUnitCost
	inventoryStatusLineItemD.IUCCodeListVersion = in.IUCCodeListVersion
	inventoryStatusLineItemD.IUCCurrencyCode = in.IUCCurrencyCode
	inventoryStatusLineItemD.LineItemNumber = in.LineItemNumber
	inventoryStatusLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	inventoryStatusLineItemD.InventoryStatusOwner = in.InventoryStatusOwner
	inventoryStatusLineItemD.InventorySubLocationId = in.InventorySubLocationId
	inventoryStatusLineItemD.LogisticUnitIdentification = in.LogisticUnitIdentification
	inventoryStatusLineItemD.ReturnableAssetIdentification = in.ReturnableAssetIdentification
	inventoryStatusLineItemD.InventoryReportTypeCode = in.InventoryReportTypeCode
	inventoryStatusLineItemD.StructureTypeCode = in.StructureTypeCode
	inventoryStatusLineItemD.InventoryReportIdentification = in.InventoryReportIdentification
	inventoryStatusLineItemD.InventoryReportingParty = in.InventoryReportingParty
	inventoryStatusLineItemD.InventoryReportToParty = in.InventoryReportToParty
	inventoryStatusLineItemD.InventoryItemLocationInformationId = in.InventoryItemLocationInformationId
	inventoryStatusLineItemD.InventoryReportId = in.InventoryReportId

	inventoryStatusLineItemT := inventoryproto.InventoryStatusLineItemT{}
	inventoryStatusLineItemT.FirstInFirstOutDateTimeBegin = common.TimeToTimestamp(firstInFirstOutDateTimeBegin.UTC().Truncate(time.Second))
	inventoryStatusLineItemT.FirstInFirstOutDateTimeEnd = common.TimeToTimestamp(firstInFirstOutDateTimeEnd.UTC().Truncate(time.Second))
	inventoryStatusLineItemT.InventoryDateTimeBegin = common.TimeToTimestamp(inventoryDateTimeBegin.UTC().Truncate(time.Second))
	inventoryStatusLineItemT.InventoryDateTimeEnd = common.TimeToTimestamp(inventoryDateTimeEnd.UTC().Truncate(time.Second))
	inventoryStatusLineItemT.ReportingPeriodBegin = common.TimeToTimestamp(reportingPeriodBegin.UTC().Truncate(time.Second))
	inventoryStatusLineItemT.ReportingPeriodEnd = common.TimeToTimestamp(reportingPeriodEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	inventoryStatusLineItem := inventoryproto.InventoryStatusLineItem{InventoryStatusLineItemD: &inventoryStatusLineItemD, InventoryStatusLineItemT: &inventoryStatusLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &inventoryStatusLineItem, nil
}

// insertInventoryStatusLineItem - Insert InventoryStatusLineItem details into database
func (invs *InventoryService) insertInventoryStatusLineItem(ctx context.Context, insertInventoryStatusLineItemSQL string, inventoryStatusLineItem *inventoryproto.InventoryStatusLineItem, userEmail string, requestID string) error {
	inventoryStatusLineItemTmp, err := invs.crInventoryStatusLineItemStruct(ctx, inventoryStatusLineItem, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInventoryStatusLineItemSQL, inventoryStatusLineItemTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryStatusLineItem.InventoryStatusLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(inventoryStatusLineItem.InventoryStatusLineItemD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryStatusLineItem.InventoryStatusLineItemD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crInventoryStatusLineItemStruct - process InventoryStatusLineItem details
func (invs *InventoryService) crInventoryStatusLineItemStruct(ctx context.Context, inventoryStatusLineItem *inventoryproto.InventoryStatusLineItem, userEmail string, requestID string) (*inventorystruct.InventoryStatusLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(inventoryStatusLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(inventoryStatusLineItem.CrUpdTime.UpdatedAt)

	inventoryStatusLineItemT := new(inventorystruct.InventoryStatusLineItemT)
	inventoryStatusLineItemT.FirstInFirstOutDateTimeBegin = common.TimestampToTime(inventoryStatusLineItem.InventoryStatusLineItemT.FirstInFirstOutDateTimeBegin)
	inventoryStatusLineItemT.FirstInFirstOutDateTimeEnd = common.TimestampToTime(inventoryStatusLineItem.InventoryStatusLineItemT.FirstInFirstOutDateTimeEnd)
	inventoryStatusLineItemT.InventoryDateTimeBegin = common.TimestampToTime(inventoryStatusLineItem.InventoryStatusLineItemT.InventoryDateTimeBegin)
	inventoryStatusLineItemT.InventoryDateTimeEnd = common.TimestampToTime(inventoryStatusLineItem.InventoryStatusLineItemT.InventoryDateTimeEnd)
	inventoryStatusLineItemT.ReportingPeriodBegin = common.TimestampToTime(inventoryStatusLineItem.InventoryStatusLineItemT.ReportingPeriodBegin)
	inventoryStatusLineItemT.ReportingPeriodEnd = common.TimestampToTime(inventoryStatusLineItem.InventoryStatusLineItemT.ReportingPeriodEnd)

	inventoryStatusLineItemTmp := inventorystruct.InventoryStatusLineItem{InventoryStatusLineItemD: inventoryStatusLineItem.InventoryStatusLineItemD, InventoryStatusLineItemT: inventoryStatusLineItemT, CrUpdUser: inventoryStatusLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &inventoryStatusLineItemTmp, nil
}
