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

const insertReturnablePackagingInventoryEventSQL = `insert into returnable_packaging_inventory_events
	    (
event_identifier,
inventory_business_step_code,
inventory_event_reason_code,
inventory_movement_type_code,
inventory_sub_location,
event_date_time,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
:event_identifier,
:inventory_business_step_code,
:inventory_event_reason_code,
:inventory_movement_type_code,
:inventory_sub_location,
:event_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectReturnablePackagingInventoryEventsSQL = `select
  id,
  event_identifier,
  inventory_business_step_code,
  inventory_event_reason_code,
  inventory_movement_type_code,
  inventory_sub_location,
  event_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from returnable_packaging_inventory_events`*/

// CreateReturnablePackagingInventoryEvent - Create ReturnablePackagingInventoryEvent
func (invs *InventoryService) CreateReturnablePackagingInventoryEvent(ctx context.Context, in *inventoryproto.CreateReturnablePackagingInventoryEventRequest) (*inventoryproto.CreateReturnablePackagingInventoryEventResponse, error) {
	returnablePackagingInventoryEvent, err := invs.ProcessReturnablePackagingInventoryEventRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertReturnablePackagingInventoryEvent(ctx, insertReturnablePackagingInventoryEventSQL, returnablePackagingInventoryEvent, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	returnablePackagingInventoryEventResponse := inventoryproto.CreateReturnablePackagingInventoryEventResponse{}
	returnablePackagingInventoryEventResponse.ReturnablePackagingInventoryEvent = returnablePackagingInventoryEvent
	return &returnablePackagingInventoryEventResponse, nil
}

// ProcessReturnablePackagingInventoryEventRequest - ProcessReturnablePackagingInventoryEventRequest
func (invs *InventoryService) ProcessReturnablePackagingInventoryEventRequest(ctx context.Context, in *inventoryproto.CreateReturnablePackagingInventoryEventRequest) (*inventoryproto.ReturnablePackagingInventoryEvent, error) {
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

	returnablePackagingInventoryEventD := inventoryproto.ReturnablePackagingInventoryEventD{}
	returnablePackagingInventoryEventD.EventIdentifier = in.EventIdentifier
	returnablePackagingInventoryEventD.InventoryBusinessStepCode = in.InventoryBusinessStepCode
	returnablePackagingInventoryEventD.InventoryEventReasonCode = in.InventoryEventReasonCode
	returnablePackagingInventoryEventD.InventoryMovementTypeCode = in.InventoryMovementTypeCode
	returnablePackagingInventoryEventD.InventorySubLocationId = in.InventorySubLocationId

	returnablePackagingInventoryEventT := inventoryproto.ReturnablePackagingInventoryEventT{}
	returnablePackagingInventoryEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	returnablePackagingInventoryEvent := inventoryproto.ReturnablePackagingInventoryEvent{ReturnablePackagingInventoryEventD: &returnablePackagingInventoryEventD, ReturnablePackagingInventoryEventT: &returnablePackagingInventoryEventT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &returnablePackagingInventoryEvent, nil
}

// insertReturnablePackagingInventoryEvent - Insert ReturnablePackagingInventoryEvent details into database
func (invs *InventoryService) insertReturnablePackagingInventoryEvent(ctx context.Context, insertReturnablePackagingInventoryEventSQL string, returnablePackagingInventoryEvent *inventoryproto.ReturnablePackagingInventoryEvent, userEmail string, requestID string) error {
	returnablePackagingInventoryEventTmp, err := invs.crReturnablePackagingInventoryEventStruct(ctx, returnablePackagingInventoryEvent, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertReturnablePackagingInventoryEventSQL, returnablePackagingInventoryEventTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		returnablePackagingInventoryEvent.ReturnablePackagingInventoryEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(returnablePackagingInventoryEvent.ReturnablePackagingInventoryEventD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		returnablePackagingInventoryEvent.ReturnablePackagingInventoryEventD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	return nil
}

// crReturnablePackagingInventoryEventStruct - process ReturnablePackagingInventoryEvent details
func (invs *InventoryService) crReturnablePackagingInventoryEventStruct(ctx context.Context, returnablePackagingInventoryEvent *inventoryproto.ReturnablePackagingInventoryEvent, userEmail string, requestID string) (*inventorystruct.ReturnablePackagingInventoryEvent, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(returnablePackagingInventoryEvent.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(returnablePackagingInventoryEvent.CrUpdTime.UpdatedAt)

	returnablePackagingInventoryEventT := new(inventorystruct.ReturnablePackagingInventoryEventT)
	returnablePackagingInventoryEventT.EventDateTime = common.TimestampToTime(returnablePackagingInventoryEvent.ReturnablePackagingInventoryEventT.EventDateTime)

	returnablePackagingInventoryEventTmp := inventorystruct.ReturnablePackagingInventoryEvent{ReturnablePackagingInventoryEventD: returnablePackagingInventoryEvent.ReturnablePackagingInventoryEventD, ReturnablePackagingInventoryEventT: returnablePackagingInventoryEventT, CrUpdUser: returnablePackagingInventoryEvent.CrUpdUser, CrUpdTime: crUpdTime}
	return &returnablePackagingInventoryEventTmp, nil
}
