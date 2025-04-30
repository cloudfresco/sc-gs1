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

const insertLogisticUnitInventoryStatusSQL = `insert into logistics_unit_inventory_statuses
	    (
inventory_disposition_code,
inventory_sub_location_id,
inventory_date_time)
  values(:inventory_date_time,
:inventory_disposition_code,
:inventory_sub_location_id,
:inventory_date_time
);`

/*const selectLogisticUnitInventoryStatussSQL = `select
  id,
  inventory_disposition_code,
  inventory_sub_location_id,
  inventory_date_time from logistics_unit_inventory_statuses`*/

// CreateLogisticUnitInventoryStatus - Create LogisticUnitInventoryStatus
func (invs *InventoryService) CreateLogisticUnitInventoryStatus(ctx context.Context, in *inventoryproto.CreateLogisticUnitInventoryStatusRequest) (*inventoryproto.CreateLogisticUnitInventoryStatusResponse, error) {
	logisticUnitInventoryStatus, err := invs.ProcessLogisticUnitInventoryStatusRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertLogisticUnitInventoryStatus(ctx, insertLogisticUnitInventoryStatusSQL, logisticUnitInventoryStatus, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticUnitInventoryStatusResponse := inventoryproto.CreateLogisticUnitInventoryStatusResponse{}
	logisticUnitInventoryStatusResponse.LogisticUnitInventoryStatus = logisticUnitInventoryStatus
	return &logisticUnitInventoryStatusResponse, nil
}

// ProcessLogisticUnitInventoryStatusRequest - ProcessLogisticUnitInventoryStatusRequest
func (invs *InventoryService) ProcessLogisticUnitInventoryStatusRequest(ctx context.Context, in *inventoryproto.CreateLogisticUnitInventoryStatusRequest) (*inventoryproto.LogisticUnitInventoryStatus, error) {
	inventoryDateTime, err := time.Parse(common.Layout, in.InventoryDateTime)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticUnitInventoryStatusD := inventoryproto.LogisticUnitInventoryStatusD{}
	logisticUnitInventoryStatusD.InventoryDispositionCode = in.InventoryDispositionCode
	logisticUnitInventoryStatusD.InventorySubLocationId = in.InventorySubLocationId

	logisticUnitInventoryStatusT := inventoryproto.LogisticUnitInventoryStatusT{}
	logisticUnitInventoryStatusT.InventoryDateTime = common.TimeToTimestamp(inventoryDateTime.UTC().Truncate(time.Second))

	logisticUnitInventoryStatus := inventoryproto.LogisticUnitInventoryStatus{LogisticUnitInventoryStatusD: &logisticUnitInventoryStatusD, LogisticUnitInventoryStatusT: &logisticUnitInventoryStatusT}

	return &logisticUnitInventoryStatus, nil
}

// insertLogisticUnitInventoryStatus - Insert LogisticUnitInventoryStatus details into database
func (invs *InventoryService) insertLogisticUnitInventoryStatus(ctx context.Context, insertLogisticUnitInventoryStatusSQL string, logisticUnitInventoryStatus *inventoryproto.LogisticUnitInventoryStatus, userEmail string, requestID string) error {
	logisticUnitInventoryStatusTmp, err := invs.crLogisticUnitInventoryStatusStruct(ctx, logisticUnitInventoryStatus, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLogisticUnitInventoryStatusSQL, logisticUnitInventoryStatusTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticUnitInventoryStatus.LogisticUnitInventoryStatusD.Id = uint32(uID)

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crLogisticUnitInventoryStatusStruct - process LogisticUnitInventoryStatus details
func (invs *InventoryService) crLogisticUnitInventoryStatusStruct(ctx context.Context, logisticUnitInventoryStatus *inventoryproto.LogisticUnitInventoryStatus, userEmail string, requestID string) (*inventorystruct.LogisticUnitInventoryStatus, error) {
	logisticUnitInventoryStatusT := new(inventorystruct.LogisticUnitInventoryStatusT)
	logisticUnitInventoryStatusT.InventoryDateTime = common.TimestampToTime(logisticUnitInventoryStatus.LogisticUnitInventoryStatusT.InventoryDateTime)

	logisticUnitInventoryStatusTmp := inventorystruct.LogisticUnitInventoryStatus{LogisticUnitInventoryStatusD: logisticUnitInventoryStatus.LogisticUnitInventoryStatusD, LogisticUnitInventoryStatusT: logisticUnitInventoryStatusT}

	return &logisticUnitInventoryStatusTmp, nil
}
