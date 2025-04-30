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

const insertLogisticUnitInventoryEventSQL = `insert into logistics_unit_inventory_events
	    (uuid4,
event_identifier,
inventory_business_step_code,
inventory_event_reason_code,
inventory_movement_type_code,
inventory_sub_location_id,
event_date_time,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
  :uuid4,
  :event_date_time,
:event_identifier,
:inventory_business_step_code,
:inventory_event_reason_code,
:inventory_movement_type_code,
:inventory_sub_location_id,
:event_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectLogisticUnitInventoryEventsSQL = `select
  id,
  uuid4,
  event_date_time,
  event_identifier,
  inventory_business_step_code,
  inventory_event_reason_code,
  inventory_movement_type_code,
  inventory_sub_location_id,
  event_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from logistics_unit_inventory_events`*/

// CreateLogisticUnitInventoryEvent - Create LogisticUnitInventoryEvent
func (invs *InventoryService) CreateLogisticUnitInventoryEvent(ctx context.Context, in *inventoryproto.CreateLogisticUnitInventoryEventRequest) (*inventoryproto.CreateLogisticUnitInventoryEventResponse, error) {
	logisticUnitInventoryEvent, err := invs.ProcessLogisticUnitInventoryEventRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertLogisticUnitInventoryEvent(ctx, insertLogisticUnitInventoryEventSQL, logisticUnitInventoryEvent, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticUnitInventoryEventResponse := inventoryproto.CreateLogisticUnitInventoryEventResponse{}
	logisticUnitInventoryEventResponse.LogisticUnitInventoryEvent = logisticUnitInventoryEvent
	return &logisticUnitInventoryEventResponse, nil
}

// ProcessLogisticUnitInventoryEventRequest - ProcessLogisticUnitInventoryEventRequest
func (invs *InventoryService) ProcessLogisticUnitInventoryEventRequest(ctx context.Context, in *inventoryproto.CreateLogisticUnitInventoryEventRequest) (*inventoryproto.LogisticUnitInventoryEvent, error) {
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

	logisticUnitInventoryEventD := inventoryproto.LogisticUnitInventoryEventD{}
	logisticUnitInventoryEventD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticUnitInventoryEventD.EventIdentifier = in.EventIdentifier
	logisticUnitInventoryEventD.InventoryBusinessStepCode = in.InventoryBusinessStepCode
	logisticUnitInventoryEventD.InventoryEventReasonCode = in.InventoryEventReasonCode
	logisticUnitInventoryEventD.InventoryMovementTypeCode = in.InventoryMovementTypeCode
	logisticUnitInventoryEventD.InventorySubLocationId = in.InventorySubLocationId

	logisticUnitInventoryEventT := inventoryproto.LogisticUnitInventoryEventT{}
	logisticUnitInventoryEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	logisticUnitInventoryEvent := inventoryproto.LogisticUnitInventoryEvent{LogisticUnitInventoryEventD: &logisticUnitInventoryEventD, LogisticUnitInventoryEventT: &logisticUnitInventoryEventT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &logisticUnitInventoryEvent, nil
}

// insertLogisticUnitInventoryEvent - Insert LogisticUnitInventoryEvent details into database
func (invs *InventoryService) insertLogisticUnitInventoryEvent(ctx context.Context, insertLogisticUnitInventoryEventSQL string, logisticUnitInventoryEvent *inventoryproto.LogisticUnitInventoryEvent, userEmail string, requestID string) error {
	logisticUnitInventoryEventTmp, err := invs.crLogisticUnitInventoryEventStruct(ctx, logisticUnitInventoryEvent, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLogisticUnitInventoryEventSQL, logisticUnitInventoryEventTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticUnitInventoryEvent.LogisticUnitInventoryEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(logisticUnitInventoryEvent.LogisticUnitInventoryEventD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticUnitInventoryEvent.LogisticUnitInventoryEventD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crLogisticUnitInventoryEventStruct - process LogisticUnitInventoryEvent details
func (invs *InventoryService) crLogisticUnitInventoryEventStruct(ctx context.Context, logisticUnitInventoryEvent *inventoryproto.LogisticUnitInventoryEvent, userEmail string, requestID string) (*inventorystruct.LogisticUnitInventoryEvent, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(logisticUnitInventoryEvent.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(logisticUnitInventoryEvent.CrUpdTime.UpdatedAt)

	logisticUnitInventoryEventT := new(inventorystruct.LogisticUnitInventoryEventT)
	logisticUnitInventoryEventT.EventDateTime = common.TimestampToTime(logisticUnitInventoryEvent.LogisticUnitInventoryEventT.EventDateTime)

	logisticUnitInventoryEventTmp := inventorystruct.LogisticUnitInventoryEvent{LogisticUnitInventoryEventD: logisticUnitInventoryEvent.LogisticUnitInventoryEventD, LogisticUnitInventoryEventT: logisticUnitInventoryEventT, CrUpdUser: logisticUnitInventoryEvent.CrUpdUser, CrUpdTime: crUpdTime}

	return &logisticUnitInventoryEventTmp, nil
}
