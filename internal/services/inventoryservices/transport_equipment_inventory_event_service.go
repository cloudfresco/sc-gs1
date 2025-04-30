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

const insertTransportEquipmentInventoryEventSQL = `insert into transport_equipment_inventory_events
	    (event_date_time,
event_identifier,
inventory_business_step_code,
inventory_disposition_code,
inventory_event_reason_code,
inventory_movement_type_code,
number_of_pieces_of_equipment,
inventory_sub_location_id)
  values(:event_date_time,
:event_identifier,
:inventory_business_step_code,
:inventory_disposition_code,
:inventory_event_reason_code,
:inventory_movement_type_code,
:number_of_pieces_of_equipment,
:inventory_sub_location_id
);`

/*const selectTransportEquipmentInventoryEventsSQL = `select
  id,
  event_date_time,
  event_identifier,
  inventory_business_step_code,
  inventory_disposition_code,
  inventory_event_reason_code,
  inventory_movement_type_code,
  number_of_pieces_of_equipment,
  inventory_sub_location_id from transport_equipment_inventory_events`*/

// CreateTransportEquipmentInventoryEvent - Create TransportEquipmentInventoryEvent
func (invs *InventoryService) CreateTransportEquipmentInventoryEvent(ctx context.Context, in *inventoryproto.CreateTransportEquipmentInventoryEventRequest) (*inventoryproto.CreateTransportEquipmentInventoryEventResponse, error) {
	transportEquipmentInventoryEvent, err := invs.ProcessTransportEquipmentInventoryEventRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransportEquipmentInventoryEvent(ctx, insertTransportEquipmentInventoryEventSQL, transportEquipmentInventoryEvent, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportEquipmentInventoryEventResponse := inventoryproto.CreateTransportEquipmentInventoryEventResponse{}
	transportEquipmentInventoryEventResponse.TransportEquipmentInventoryEvent = transportEquipmentInventoryEvent
	return &transportEquipmentInventoryEventResponse, nil
}

// ProcessTransportEquipmentInventoryEventRequest - ProcessTransportEquipmentInventoryEventRequest
func (invs *InventoryService) ProcessTransportEquipmentInventoryEventRequest(ctx context.Context, in *inventoryproto.CreateTransportEquipmentInventoryEventRequest) (*inventoryproto.TransportEquipmentInventoryEvent, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	eventDateTime, err := time.Parse(common.Layout, in.EventDateTime)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportEquipmentInventoryEventD := inventoryproto.TransportEquipmentInventoryEventD{}
	transportEquipmentInventoryEventD.EventIdentifier = in.EventIdentifier
	transportEquipmentInventoryEventD.InventoryBusinessStepCode = in.InventoryBusinessStepCode
	transportEquipmentInventoryEventD.InventoryDispositionCode = in.InventoryDispositionCode
	transportEquipmentInventoryEventD.InventoryEventReasonCode = in.InventoryEventReasonCode
	transportEquipmentInventoryEventD.InventoryMovementTypeCode = in.InventoryMovementTypeCode
	transportEquipmentInventoryEventD.NumberOfPiecesOfEquipment = in.NumberOfPiecesOfEquipment
	transportEquipmentInventoryEventD.InventorySubLocationId = in.InventorySubLocationId

	transportEquipmentInventoryEventT := inventoryproto.TransportEquipmentInventoryEventT{}
	transportEquipmentInventoryEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transportEquipmentInventoryEvent := inventoryproto.TransportEquipmentInventoryEvent{TransportEquipmentInventoryEventD: &transportEquipmentInventoryEventD, TransportEquipmentInventoryEventT: &transportEquipmentInventoryEventT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &transportEquipmentInventoryEvent, nil
}

// insertTransportEquipmentInventoryEvent - Insert TransportEquipmentInventoryEvent details into database
func (invs *InventoryService) insertTransportEquipmentInventoryEvent(ctx context.Context, insertTransportEquipmentInventoryEventSQL string, transportEquipmentInventoryEvent *inventoryproto.TransportEquipmentInventoryEvent, userEmail string, requestID string) error {
	transportEquipmentInventoryEventTmp, err := invs.crTransportEquipmentInventoryEventStruct(ctx, transportEquipmentInventoryEvent, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransportEquipmentInventoryEventSQL, transportEquipmentInventoryEventTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportEquipmentInventoryEvent.TransportEquipmentInventoryEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transportEquipmentInventoryEvent.TransportEquipmentInventoryEventD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportEquipmentInventoryEvent.TransportEquipmentInventoryEventD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTransportEquipmentInventoryEventStruct - process TransportEquipmentInventoryEvent details
func (invs *InventoryService) crTransportEquipmentInventoryEventStruct(ctx context.Context, transportEquipmentInventoryEvent *inventoryproto.TransportEquipmentInventoryEvent, userEmail string, requestID string) (*inventorystruct.TransportEquipmentInventoryEvent, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transportEquipmentInventoryEvent.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transportEquipmentInventoryEvent.CrUpdTime.UpdatedAt)

	transportEquipmentInventoryEventT := new(inventorystruct.TransportEquipmentInventoryEventT)
	transportEquipmentInventoryEventT.EventDateTime = common.TimestampToTime(transportEquipmentInventoryEvent.TransportEquipmentInventoryEventT.EventDateTime)

	transportEquipmentInventoryEventTmp := inventorystruct.TransportEquipmentInventoryEvent{TransportEquipmentInventoryEventD: transportEquipmentInventoryEvent.TransportEquipmentInventoryEventD, TransportEquipmentInventoryEventT: transportEquipmentInventoryEventT, CrUpdUser: transportEquipmentInventoryEvent.CrUpdUser, CrUpdTime: crUpdTime}

	return &transportEquipmentInventoryEventTmp, nil
}
