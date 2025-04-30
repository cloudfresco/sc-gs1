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

const insertTradeItemInventoryEventSQL = `insert into trade_item_inventory_events
	    (event_date_time,
event_identifier,
inventory_business_step_code,
inventory_disposition_code,
inventory_event_reason_code,
inventory_movement_type_code,
inventory_sub_location_id,
liable_party,
logistics_inventory_report_id
)
  values(:event_date_time,
:event_identifier,
:inventory_business_step_code,
:inventory_disposition_code,
:inventory_event_reason_code,
:inventory_movement_type_code,
:inventory_sub_location_id,
:liable_party,
logistics_inventory_report_id
);`

/*const selectTradeItemInventoryEventsSQL = `select
  id,
  event_date_time,
  event_identifier,
  inventory_business_step_code,
  inventory_disposition_code,
  inventory_event_reason_code,
  inventory_movement_type_code,
  inventory_sub_location_id,
  liable_party,
  logistics_inventory_report_id from trade_item_inventory_events`*/

// CreateTradeItemInventoryEvent - Create TradeItemInventoryEvent
func (invs *InventoryService) CreateTradeItemInventoryEvent(ctx context.Context, in *inventoryproto.CreateTradeItemInventoryEventRequest) (*inventoryproto.CreateTradeItemInventoryEventResponse, error) {
	tradeItemInventoryEvent, err := invs.ProcessTradeItemInventoryEventRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTradeItemInventoryEvent(ctx, insertTradeItemInventoryEventSQL, tradeItemInventoryEvent, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	tradeItemInventoryEventResponse := inventoryproto.CreateTradeItemInventoryEventResponse{}
	tradeItemInventoryEventResponse.TradeItemInventoryEvent = tradeItemInventoryEvent
	return &tradeItemInventoryEventResponse, nil
}

// ProcessTradeItemInventoryEventRequest - ProcessTradeItemInventoryEventRequest
func (invs *InventoryService) ProcessTradeItemInventoryEventRequest(ctx context.Context, in *inventoryproto.CreateTradeItemInventoryEventRequest) (*inventoryproto.TradeItemInventoryEvent, error) {
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

	tradeItemInventoryEventD := inventoryproto.TradeItemInventoryEventD{}
	tradeItemInventoryEventD.EventIdentifier = in.EventIdentifier
	tradeItemInventoryEventD.InventoryBusinessStepCode = in.InventoryBusinessStepCode
	tradeItemInventoryEventD.InventoryDispositionCode = in.InventoryDispositionCode
	tradeItemInventoryEventD.InventoryEventReasonCode = in.InventoryEventReasonCode
	tradeItemInventoryEventD.InventoryMovementTypeCode = in.InventoryMovementTypeCode
	tradeItemInventoryEventD.InventorySubLocationId = in.InventorySubLocationId
	tradeItemInventoryEventD.LiableParty = in.LiableParty
	tradeItemInventoryEventD.LogisticsInventoryReportId = in.LogisticsInventoryReportId

	tradeItemInventoryEventT := inventoryproto.TradeItemInventoryEventT{}
	tradeItemInventoryEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	tradeItemInventoryEvent := inventoryproto.TradeItemInventoryEvent{TradeItemInventoryEventD: &tradeItemInventoryEventD, TradeItemInventoryEventT: &tradeItemInventoryEventT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &tradeItemInventoryEvent, nil
}

// insertTradeItemInventoryEvent - Insert TradeItemInventoryEvent details into database
func (invs *InventoryService) insertTradeItemInventoryEvent(ctx context.Context, insertTradeItemInventoryEventSQL string, tradeItemInventoryEvent *inventoryproto.TradeItemInventoryEvent, userEmail string, requestID string) error {
	tradeItemInventoryEventTmp, err := invs.crTradeItemInventoryEventStruct(ctx, tradeItemInventoryEvent, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTradeItemInventoryEventSQL, tradeItemInventoryEventTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		tradeItemInventoryEvent.TradeItemInventoryEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(tradeItemInventoryEvent.TradeItemInventoryEventD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		tradeItemInventoryEvent.TradeItemInventoryEventD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTradeItemInventoryEventStruct - process TradeItemInventoryEvent details
func (invs *InventoryService) crTradeItemInventoryEventStruct(ctx context.Context, tradeItemInventoryEvent *inventoryproto.TradeItemInventoryEvent, userEmail string, requestID string) (*inventorystruct.TradeItemInventoryEvent, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(tradeItemInventoryEvent.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(tradeItemInventoryEvent.CrUpdTime.UpdatedAt)

	tradeItemInventoryEventT := new(inventorystruct.TradeItemInventoryEventT)
	tradeItemInventoryEventT.EventDateTime = common.TimestampToTime(tradeItemInventoryEvent.TradeItemInventoryEventT.EventDateTime)

	tradeItemInventoryEventTmp := inventorystruct.TradeItemInventoryEvent{TradeItemInventoryEventD: tradeItemInventoryEvent.TradeItemInventoryEventD, TradeItemInventoryEventT: tradeItemInventoryEventT, CrUpdUser: tradeItemInventoryEvent.CrUpdUser, CrUpdTime: crUpdTime}

	return &tradeItemInventoryEventTmp, nil
}
