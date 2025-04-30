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

const insertTradeItemInventoryStatusSQL = `insert into trade_item_inventory_statuses
	    (inventory_disposition_code,
inventory_sub_location_id,
inventory_date_time,
logistics_inventory_report_id)
  values(:inventory_disposition_code,
:inventory_sub_location_id,
:inventory_date_time,
:logistics_inventory_report_id
);`

/*const selectTradeItemInventoryStatussSQL = `select
  id,
  inventory_disposition_code,
  inventory_sub_location_id,
  inventory_date_time,
  logistics_inventory_report_id from trade_item_inventory_statuses`*/

// CreateTradeItemInventoryStatus - Create TradeItemInventoryStatus
func (invs *InventoryService) CreateTradeItemInventoryStatus(ctx context.Context, in *inventoryproto.CreateTradeItemInventoryStatusRequest) (*inventoryproto.CreateTradeItemInventoryStatusResponse, error) {
	tradeItemInventoryStatus, err := invs.ProcessTradeItemInventoryStatusRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTradeItemInventoryStatus(ctx, insertTradeItemInventoryStatusSQL, tradeItemInventoryStatus, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	tradeItemInventoryStatusResponse := inventoryproto.CreateTradeItemInventoryStatusResponse{}
	tradeItemInventoryStatusResponse.TradeItemInventoryStatus = tradeItemInventoryStatus
	return &tradeItemInventoryStatusResponse, nil
}

// ProcessTradeItemInventoryStatusRequest - ProcessTradeItemInventoryStatusRequest
func (invs *InventoryService) ProcessTradeItemInventoryStatusRequest(ctx context.Context, in *inventoryproto.CreateTradeItemInventoryStatusRequest) (*inventoryproto.TradeItemInventoryStatus, error) {
	inventoryDateTime, err := time.Parse(common.Layout, in.InventoryDateTime)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	tradeItemInventoryStatusD := inventoryproto.TradeItemInventoryStatusD{}
	tradeItemInventoryStatusD.InventoryDispositionCode = in.InventoryDispositionCode
	tradeItemInventoryStatusD.InventorySubLocationId = in.InventorySubLocationId
	tradeItemInventoryStatusD.LogisticsInventoryReportId = in.LogisticsInventoryReportId

	tradeItemInventoryStatusT := inventoryproto.TradeItemInventoryStatusT{}
	tradeItemInventoryStatusT.InventoryDateTime = common.TimeToTimestamp(inventoryDateTime.UTC().Truncate(time.Second))

	tradeItemInventoryStatus := inventoryproto.TradeItemInventoryStatus{TradeItemInventoryStatusD: &tradeItemInventoryStatusD, TradeItemInventoryStatusT: &tradeItemInventoryStatusT}

	return &tradeItemInventoryStatus, nil
}

// insertTradeItemInventoryStatus - Insert TradeItemInventoryStatus details into database
func (invs *InventoryService) insertTradeItemInventoryStatus(ctx context.Context, insertTradeItemInventoryStatusSQL string, tradeItemInventoryStatus *inventoryproto.TradeItemInventoryStatus, userEmail string, requestID string) error {
	tradeItemInventoryStatusTmp, err := invs.crTradeItemInventoryStatusStruct(ctx, tradeItemInventoryStatus, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTradeItemInventoryStatusSQL, tradeItemInventoryStatusTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		tradeItemInventoryStatus.TradeItemInventoryStatusD.Id = uint32(uID)

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTradeItemInventoryStatusStruct - process TradeItemInventoryStatus details
func (invs *InventoryService) crTradeItemInventoryStatusStruct(ctx context.Context, tradeItemInventoryStatus *inventoryproto.TradeItemInventoryStatus, userEmail string, requestID string) (*inventorystruct.TradeItemInventoryStatus, error) {
	tradeItemInventoryStatusT := new(inventorystruct.TradeItemInventoryStatusT)
	tradeItemInventoryStatusT.InventoryDateTime = common.TimestampToTime(tradeItemInventoryStatus.TradeItemInventoryStatusT.InventoryDateTime)

	tradeItemInventoryStatusTmp := inventorystruct.TradeItemInventoryStatus{TradeItemInventoryStatusD: tradeItemInventoryStatus.TradeItemInventoryStatusD, TradeItemInventoryStatusT: tradeItemInventoryStatusT}

	return &tradeItemInventoryStatusTmp, nil
}
