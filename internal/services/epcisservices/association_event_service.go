package epcisservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	epcisstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertAssociationEventSQL = `insert into association_events
	  (uuid4,
    event_id,
    event_time_zone_offset,
    certification,
    event_time,
    reason,
    declaration_time,
    parent_id,
    action,
    biz_step,
    disposition,
    read_point,
    biz_location,
    status_code,
    created_by_user_id,
    updated_by_user_id)
      values(
    :uuid4,
    :event_id,
    :event_time_zone_offset,
    :certification,
    :event_time,
    :reason,
    :declaration_time,
    :parent_id,
    :action,
    :biz_step,
    :disposition,
    :read_point,
    :biz_location,
    :status_code,
    :created_by_user_id,
    :updated_by_user_id);`

const selectAssociationEventsSQL = `select
  id,
  uuid4,
  event_id,
  event_time_zone_offset,
  certification,
  event_time,
  reason,
  declaration_time,
  parent_id,
  action,
  biz_step,
  disposition,
  read_point,
  biz_location,
  status_code,
  created_by_user_id,
  updated_by_user_id from association_events`

// CreateAssociationEvent - Create AssociationEvent
func (es *EpcisService) CreateAssociationEvent(ctx context.Context, in *epcisproto.CreateAssociationEventRequest) (*epcisproto.CreateAssociationEventResponse, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, es.UserServiceClient)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eventTime, err := time.Parse(common.Layout, in.EventTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	declarationTime, err := time.Parse(common.Layout, in.DeclarationTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	associationEventD := epcisproto.AssociationEventD{}
	associationEventD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	epcisEventD := commonproto.EpcisEventD{}
	epcisEventD.EventId = in.EventId
	epcisEventD.EventTimeZoneOffset = in.EventTimeZoneOffset
	epcisEventD.Certification = in.Certification

	epcisEventT := commonproto.EpcisEventT{}
	epcisEventT.EventTime = common.TimeToTimestamp(eventTime.UTC().Truncate(time.Second))

	errorDeclarationD := commonproto.ErrorDeclarationD{}
	errorDeclarationD.Reason = in.Reason

	errorDeclarationT := commonproto.ErrorDeclarationT{}
	errorDeclarationT.DeclarationTime = common.TimeToTimestamp(declarationTime.UTC().Truncate(time.Second))

	associationEventD.ParentId = in.ParentId
	associationEventD.Action = in.Action
	associationEventD.BizStep = in.BizStep
	associationEventD.Disposition = in.Disposition
	associationEventD.ReadPoint = in.ReadPoint
	associationEventD.BizLocation = in.BizLocation

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	associationEvent := epcisproto.AssociationEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, AssociationEventD: &associationEventD, CrUpdUser: &crUpdUser}

	persistentDispositions := []*epcisproto.PersistentDisposition{}
	for _, persistentDisposition := range in.PersistentDispositions {
		persistentDisposition.UserId = in.UserId
		persistentDisposition.UserEmail = in.UserEmail
		persistentDisposition.RequestId = in.RequestId

		pDisposition, err := es.ProcessPersistentDispositionRequest(ctx, persistentDisposition)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		persistentDispositions = append(persistentDispositions, pDisposition)
	}

	childEpcs := []*epcisproto.Epc{}
	for _, epc := range in.ChildEpcs {
		epc.UserId = in.UserId
		epc.UserEmail = in.UserEmail
		epc.RequestId = in.RequestId

		ep, err := es.ProcessEpcRequest(ctx, epc)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		childEpcs = append(childEpcs, ep)
	}

	bizTransactionList := []*epcisproto.BizTransaction{}
	for _, bizTransaction := range in.BizTransactionList {
		bizTransaction.UserId = in.UserId
		bizTransaction.UserEmail = in.UserEmail
		bizTransaction.RequestId = in.RequestId

		bTransaction, err := es.ProcessBizTransactionRequest(ctx, bizTransaction)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		bizTransactionList = append(bizTransactionList, bTransaction)
	}

	childChildQuantityList := []*epcisproto.QuantityElement{}
	for _, quantityElement := range in.ChildQuantityList {
		quantityElement.UserId = in.UserId
		quantityElement.UserEmail = in.UserEmail
		quantityElement.RequestId = in.RequestId

		qElement, err := es.ProcessQuantityElementRequest(ctx, quantityElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		childChildQuantityList = append(childChildQuantityList, qElement)
	}

	sourceList := []*epcisproto.Source{}
	for _, source := range in.SourceList {
		source.UserId = in.UserId
		source.UserEmail = in.UserEmail
		source.RequestId = in.RequestId

		src, err := es.ProcessSourceRequest(ctx, source)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		sourceList = append(sourceList, src)
	}

	destinationList := []*epcisproto.Destination{}
	for _, destination := range in.DestinationList {
		destination.UserId = in.UserId
		destination.UserEmail = in.UserEmail
		destination.RequestId = in.RequestId

		dest, err := es.ProcessDestinationRequest(ctx, destination)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		destinationList = append(destinationList, dest)
	}

	sensorElementList := []*epcisproto.SensorElement{}
	for _, sensorElement := range in.SensorElementList {
		sensorElement.UserId = in.UserId
		sensorElement.UserEmail = in.UserEmail
		sensorElement.RequestId = in.RequestId

		sElement, err := es.ProcessSensorElementRequest(ctx, sensorElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		sensorElementList = append(sensorElementList, sElement)
	}

	err = es.insertAssociationEvent(ctx, insertAssociationEventSQL, &associationEvent, insertPersistentDispositionSQL, persistentDispositions, insertEpcSQL, childEpcs, insertBizTransactionSQL, bizTransactionList, insertQuantityElementSQL, childChildQuantityList, insertSourceSQL, sourceList, insertDestinationSQL, destinationList, insertSensorElementSQL, sensorElementList, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	associationEventResponse := epcisproto.CreateAssociationEventResponse{}
	associationEventResponse.AssociationEvent = &associationEvent
	return &associationEventResponse, nil
}

func (es *EpcisService) insertAssociationEvent(ctx context.Context, insertAssociationEventSQL string, associationEvent *epcisproto.AssociationEvent, insertPersistentDispositionSQL string, persistentDispositions []*epcisproto.PersistentDisposition, insertEpcSQL string, childEpcs []*epcisproto.Epc, insertBizTransactionSQL string, bizTransactionList []*epcisproto.BizTransaction, insertQuantityElementSQL string, childChildQuantityList []*epcisproto.QuantityElement, insertSourceSQL string, sourceList []*epcisproto.Source, insertDestinationSQL string, destinationList []*epcisproto.Destination, insertSensorElementSQL string, sensorElementList []*epcisproto.SensorElement, userEmail string, requestID string,
) error {
	associationEventTmp, err := es.crAssociationEventStruct(ctx, associationEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertAssociationEventSQL, associationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		associationEvent.AssociationEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(associationEvent.AssociationEventD.Uuid4)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		associationEvent.AssociationEventD.IdS = uuid4Str

		for _, persistentDisposition := range persistentDispositions {
			persistentDisposition.EventId = associationEvent.AssociationEventD.Id
			persistentDisposition.TypeOfEvent = "AssociationEvent"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertPersistentDispositionSQL, persistentDisposition)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, epc := range childEpcs {
			epc.EventId = associationEvent.AssociationEventD.Id
			epc.TypeOfEvent = "AssociationEvent"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertEpcSQL, epc)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, bizTransaction := range bizTransactionList {
			bizTransaction.EventId = associationEvent.AssociationEventD.Id
			bizTransaction.TypeOfEvent = "AssociationEvent"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertBizTransactionSQL, bizTransaction)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, quantityElement := range childChildQuantityList {
			quantityElement.EventId = associationEvent.AssociationEventD.Id
			quantityElement.TypeOfEvent = "AssociationEvent"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertQuantityElementSQL, quantityElement)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, source := range sourceList {
			source.EventId = associationEvent.AssociationEventD.Id
			source.TypeOfEvent = "AssociationEvent"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertSourceSQL, source)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, destination := range destinationList {
			destination.EventId = associationEvent.AssociationEventD.Id
			destination.TypeOfEvent = "AssociationEvent"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertDestinationSQL, destination)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, sensorElement := range sensorElementList {
			sensorElement.SensorElementD.EventId = associationEvent.AssociationEventD.Id
			sensorElement.SensorElementD.TypeOfEvent = "AssociationEvent"

			sensorElementTmp, err := es.crSensorElementStruct(ctx, sensorElement, userEmail, requestID)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			res, err = tx.NamedExecContext(ctx, insertSensorElementSQL, sensorElementTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			sensorElementId, err := res.LastInsertId()
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

			sensorElement.SensorElementD.Id = uint32(sensorElementId)

			for _, sensorReport := range sensorElement.SensorReports {
				sensorReport.SensorReportD.SensorElementId = sensorElement.SensorElementD.Id

				sensorReportTmp, err := es.crSensorReportStruct(ctx, sensorReport, userEmail, requestID)
				if err != nil {
					es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
					return err
				}
				_, err = tx.NamedExecContext(ctx, insertSensorReportSQL, sensorReportTmp)
				if err != nil {
					es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crAssociationEventStruct - process AssociationEvent details
func (es *EpcisService) crAssociationEventStruct(ctx context.Context, associationEvent *epcisproto.AssociationEvent, userEmail string, requestID string) (*epcisstruct.AssociationEvent, error) {
	epcisEventT := new(commonstruct.EpcisEventT)
	epcisEventT.EventTime = common.TimestampToTime(associationEvent.EpcisEventT.EventTime)

	errorDeclarationT := new(commonstruct.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimestampToTime(associationEvent.ErrorDeclarationT.DeclarationTime)

	associationEventTmp := epcisstruct.AssociationEvent{EpcisEventD: associationEvent.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: associationEvent.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, AssociationEventD: associationEvent.AssociationEventD, CrUpdUser: associationEvent.CrUpdUser}

	return &associationEventTmp, nil
}

func (es *EpcisService) GetAssociationEvents(ctx context.Context, in *epcisproto.GetAssociationEventsRequest) (*epcisproto.GetAssociationEventsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = es.DBService.LimitSQLRows
	}

	query := "status_code = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	associationEvents := []*epcisproto.AssociationEvent{}

	nselectAssociationEventsSQL := selectAssociationEventsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectAssociationEventsSQL, common.Active)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		associationEventTmp := epcisstruct.AssociationEvent{}
		err = rows.StructScan(&associationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		associationEvent, err := es.getAssociationEventStruct(ctx, &getRequest, associationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		associationEvents = append(associationEvents, associationEvent)
	}

	associationEventsResponse := epcisproto.GetAssociationEventsResponse{}
	if len(associationEvents) != 0 {
		next := associationEvents[len(associationEvents)-1].AssociationEventD.Id
		next--
		nextc := common.EncodeCursor(next)
		associationEventsResponse = epcisproto.GetAssociationEventsResponse{AssociationEvents: associationEvents, NextCursor: nextc}
	} else {
		associationEventsResponse = epcisproto.GetAssociationEventsResponse{AssociationEvents: associationEvents, NextCursor: "0"}
	}
	return &associationEventsResponse, nil
}

// GetAssociationEvent - Get AssociationEvent
func (es *EpcisService) GetAssociationEvent(ctx context.Context, inReq *epcisproto.GetAssociationEventRequest) (*epcisproto.GetAssociationEventResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectAssociationEventsSQL := selectAssociationEventsSQL + ` where uuid4 = ?;`
	row := es.DBService.DB.QueryRowxContext(ctx, nselectAssociationEventsSQL, uuid4byte)
	associationEventTmp := epcisstruct.AssociationEvent{}
	err = row.StructScan(&associationEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	associationEvent, err := es.getAssociationEventStruct(ctx, in, associationEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	associationEventResponse := epcisproto.GetAssociationEventResponse{}
	associationEventResponse.AssociationEvent = associationEvent
	return &associationEventResponse, nil
}

// getAssociationEventStruct - Get AssociationEvent
func (es *EpcisService) getAssociationEventStruct(ctx context.Context, in *commonproto.GetRequest, associationEventTmp epcisstruct.AssociationEvent) (*epcisproto.AssociationEvent, error) {
	epcisEventT := new(commonproto.EpcisEventT)
	epcisEventT.EventTime = common.TimeToTimestamp(associationEventTmp.EpcisEventT.EventTime)

	errorDeclarationT := new(commonproto.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimeToTimestamp(associationEventTmp.ErrorDeclarationT.DeclarationTime)

	uuid4Str, err := common.UUIDBytesToStr(associationEventTmp.AssociationEventD.Uuid4)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	associationEventTmp.AssociationEventD.IdS = uuid4Str

	associationEvent := epcisproto.AssociationEvent{EpcisEventD: associationEventTmp.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: associationEventTmp.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, AssociationEventD: associationEventTmp.AssociationEventD, CrUpdUser: associationEventTmp.CrUpdUser}

	return &associationEvent, nil
}
