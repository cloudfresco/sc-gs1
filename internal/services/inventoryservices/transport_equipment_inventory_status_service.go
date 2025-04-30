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

const insertTransportEquipmentInventoryStatusSQL = `insert into transport_equipment_inventory_statuses
	    (
inventory_disposition_code,
number_of_pieces_of_equipment,
inventory_sub_location_id,
inventory_date_time)
  values(
:inventory_disposition_code,
:number_of_pieces_of_equipment,
:inventory_sub_location_id,
:inventory_date_time
);`

/*const selectTransportEquipmentInventoryStatussSQL = `select
  id,
  inventory_disposition_code,
  number_of_pieces_of_equipment,
  inventory_sub_location_id,
  inventory_date_time from transport_equipment_inventory_statuses`*/

// CreateTransportEquipmentInventoryStatus - Create TransportEquipmentInventoryStatus
func (invs *InventoryService) CreateTransportEquipmentInventoryStatus(ctx context.Context, in *inventoryproto.CreateTransportEquipmentInventoryStatusRequest) (*inventoryproto.CreateTransportEquipmentInventoryStatusResponse, error) {
	transportEquipmentInventoryStatus, err := invs.ProcessTransportEquipmentInventoryStatusRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransportEquipmentInventoryStatus(ctx, insertTransportEquipmentInventoryStatusSQL, transportEquipmentInventoryStatus, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportEquipmentInventoryStatusResponse := inventoryproto.CreateTransportEquipmentInventoryStatusResponse{}
	transportEquipmentInventoryStatusResponse.TransportEquipmentInventoryStatus = transportEquipmentInventoryStatus
	return &transportEquipmentInventoryStatusResponse, nil
}

// ProcessTransportEquipmentInventoryStatusRequest - ProcessTransportEquipmentInventoryStatusRequest
func (invs *InventoryService) ProcessTransportEquipmentInventoryStatusRequest(ctx context.Context, in *inventoryproto.CreateTransportEquipmentInventoryStatusRequest) (*inventoryproto.TransportEquipmentInventoryStatus, error) {
	inventoryDateTime, err := time.Parse(common.Layout, in.InventoryDateTime)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportEquipmentInventoryStatusD := inventoryproto.TransportEquipmentInventoryStatusD{}
	transportEquipmentInventoryStatusD.InventoryDispositionCode = in.InventoryDispositionCode
	transportEquipmentInventoryStatusD.NumberOfPiecesOfEquipment = in.NumberOfPiecesOfEquipment
	transportEquipmentInventoryStatusD.InventorySubLocationId = in.InventorySubLocationId

	transportEquipmentInventoryStatusT := inventoryproto.TransportEquipmentInventoryStatusT{}
	transportEquipmentInventoryStatusT.InventoryDateTime = common.TimeToTimestamp(inventoryDateTime.UTC().Truncate(time.Second))

	transportEquipmentInventoryStatus := inventoryproto.TransportEquipmentInventoryStatus{TransportEquipmentInventoryStatusD: &transportEquipmentInventoryStatusD, TransportEquipmentInventoryStatusT: &transportEquipmentInventoryStatusT}

	return &transportEquipmentInventoryStatus, nil
}

// insertTransportEquipmentInventoryStatus - Insert TransportEquipmentInventoryStatus details into database
func (invs *InventoryService) insertTransportEquipmentInventoryStatus(ctx context.Context, insertTransportEquipmentInventoryStatusSQL string, transportEquipmentInventoryStatus *inventoryproto.TransportEquipmentInventoryStatus, userEmail string, requestID string) error {
	transportEquipmentInventoryStatusTmp, err := invs.crTransportEquipmentInventoryStatusStruct(ctx, transportEquipmentInventoryStatus, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransportEquipmentInventoryStatusSQL, transportEquipmentInventoryStatusTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportEquipmentInventoryStatus.TransportEquipmentInventoryStatusD.Id = uint32(uID)

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTransportEquipmentInventoryStatusStruct - process TransportEquipmentInventoryStatus details
func (invs *InventoryService) crTransportEquipmentInventoryStatusStruct(ctx context.Context, transportEquipmentInventoryStatus *inventoryproto.TransportEquipmentInventoryStatus, userEmail string, requestID string) (*inventorystruct.TransportEquipmentInventoryStatus, error) {
	transportEquipmentInventoryStatusT := new(inventorystruct.TransportEquipmentInventoryStatusT)
	transportEquipmentInventoryStatusT.InventoryDateTime = common.TimestampToTime(transportEquipmentInventoryStatus.TransportEquipmentInventoryStatusT.InventoryDateTime)

	transportEquipmentInventoryStatusTmp := inventorystruct.TransportEquipmentInventoryStatus{TransportEquipmentInventoryStatusD: transportEquipmentInventoryStatus.TransportEquipmentInventoryStatusD, TransportEquipmentInventoryStatusT: transportEquipmentInventoryStatusT}

	return &transportEquipmentInventoryStatusTmp, nil
}
