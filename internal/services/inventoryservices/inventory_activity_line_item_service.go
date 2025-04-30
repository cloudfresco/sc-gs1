package inventoryservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	inventorystruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInventoryActivityLineItemSQL = `insert into inventory_activity_line_items
	    (line_item_number,
      parent_line_item_number,
      reporting_period_begin,
      reporting_period_end,
      inventory_item_location_information_id,
      inventory_report_id)
  values(
  :line_item_number,
  :parent_line_item_number,
  :inventory_item_location_information_id,
  :inventory_report_id
  :reporting_period_begin,
  :reporting_period_end
);`

/*const selectInventoryActivityLineItemsSQL = `select
  id,
  line_item_number,
  parent_line_item_number,
  inventory_item_location_information_id,
  inventory_report_id
  reporting_period_begin,
  reporting_period_end from inventory_activity_line_items`*/

// CreateInventoryActivityLineItem - Create InventoryActivityLineItem
func (invs *InventoryService) CreateInventoryActivityLineItem(ctx context.Context, in *inventoryproto.CreateInventoryActivityLineItemRequest) (*inventoryproto.CreateInventoryActivityLineItemResponse, error) {
	inventoryActivityLineItem, err := invs.ProcessInventoryActivityLineItemRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInventoryActivityLineItem(ctx, insertInventoryActivityLineItemSQL, inventoryActivityLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryActivityLineItemResponse := inventoryproto.CreateInventoryActivityLineItemResponse{}
	inventoryActivityLineItemResponse.InventoryActivityLineItem = inventoryActivityLineItem
	return &inventoryActivityLineItemResponse, nil
}

// ProcessInventoryActivityLineItemRequest - ProcessInventoryActivityLineItemRequest
func (invs *InventoryService) ProcessInventoryActivityLineItemRequest(ctx context.Context, in *inventoryproto.CreateInventoryActivityLineItemRequest) (*inventoryproto.InventoryActivityLineItem, error) {
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

	inventoryActivityLineItemD := inventoryproto.InventoryActivityLineItemD{}
	inventoryActivityLineItemD.LineItemNumber = in.LineItemNumber
	inventoryActivityLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	inventoryActivityLineItemD.InventoryItemLocationInformationId = in.InventoryItemLocationInformationId
	inventoryActivityLineItemD.InventoryReportId = in.InventoryReportId

	inventoryActivityLineItemT := inventoryproto.InventoryActivityLineItemT{}
	inventoryActivityLineItemT.ReportingPeriodBegin = common.TimeToTimestamp(reportingPeriodBegin.UTC().Truncate(time.Second))
	inventoryActivityLineItemT.ReportingPeriodEnd = common.TimeToTimestamp(reportingPeriodEnd.UTC().Truncate(time.Second))

	inventoryActivityLineItem := inventoryproto.InventoryActivityLineItem{InventoryActivityLineItemD: &inventoryActivityLineItemD, InventoryActivityLineItemT: &inventoryActivityLineItemT}

	return &inventoryActivityLineItem, nil
}

// insertInventoryActivityLineItem - Insert InventoryActivityLineItem details into database
func (invs *InventoryService) insertInventoryActivityLineItem(ctx context.Context, insertInventoryActivityLineItemSQL string, inventoryActivityLineItem *inventoryproto.InventoryActivityLineItem, userEmail string, requestID string) error {
	inventoryActivityLineItemTmp, err := invs.crInventoryActivityLineItemStruct(ctx, inventoryActivityLineItem, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInventoryActivityLineItemSQL, inventoryActivityLineItemTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryActivityLineItem.InventoryActivityLineItemD.Id = uint32(uID)

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crInventoryActivityLineItemStruct - process InventoryActivityLineItem details
func (invs *InventoryService) crInventoryActivityLineItemStruct(ctx context.Context, inventoryActivityLineItem *inventoryproto.InventoryActivityLineItem, userEmail string, requestID string) (*inventorystruct.InventoryActivityLineItem, error) {
	inventoryActivityLineItemT := new(inventorystruct.InventoryActivityLineItemT)
	inventoryActivityLineItemT.ReportingPeriodBegin = common.TimestampToTime(inventoryActivityLineItem.InventoryActivityLineItemT.ReportingPeriodBegin)
	inventoryActivityLineItemT.ReportingPeriodEnd = common.TimestampToTime(inventoryActivityLineItem.InventoryActivityLineItemT.ReportingPeriodEnd)

	inventoryActivityLineItemTmp := inventorystruct.InventoryActivityLineItem{InventoryActivityLineItemD: inventoryActivityLineItem.InventoryActivityLineItemD, InventoryActivityLineItemT: inventoryActivityLineItemT}

	return &inventoryActivityLineItemTmp, nil
}
