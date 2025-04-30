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

const insertAggregationEventSQL = `insert into aggregation_events
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

const selectAggregationEventsSQL = `select
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
  updated_by_user_id from aggregation_events`

// CreateAggregationEvent - Create AggregationEvent
func (es *EpcisService) CreateAggregationEvent(ctx context.Context, in *epcisproto.CreateAggregationEventRequest) (*epcisproto.CreateAggregationEventResponse, error) {
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

	aggregationEventD := epcisproto.AggregationEventD{}
	aggregationEventD.Uuid4, err = common.GetUUIDBytes()
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

	aggregationEventD.ParentId = in.ParentId
	aggregationEventD.Action = in.Action
	aggregationEventD.BizStep = in.BizStep
	aggregationEventD.Disposition = in.Disposition
	aggregationEventD.ReadPoint = in.ReadPoint
	aggregationEventD.BizLocation = in.BizLocation

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	aggregationEvent := epcisproto.AggregationEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, AggregationEventD: &aggregationEventD, CrUpdUser: &crUpdUser}

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

	childQuantityList := []*epcisproto.QuantityElement{}
	for _, quantityElement := range in.ChildQuantityList {
		quantityElement.UserId = in.UserId
		quantityElement.UserEmail = in.UserEmail
		quantityElement.RequestId = in.RequestId

		qElement, err := es.ProcessQuantityElementRequest(ctx, quantityElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		childQuantityList = append(childQuantityList, qElement)
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

	err = es.insertAggregationEvent(ctx, insertAggregationEventSQL, &aggregationEvent, insertPersistentDispositionSQL, persistentDispositions, insertEpcSQL, childEpcs, insertBizTransactionSQL, bizTransactionList, insertQuantityElementSQL, childQuantityList, insertSourceSQL, sourceList, insertDestinationSQL, destinationList, insertSensorElementSQL, sensorElementList, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	aggregationEventResponse := epcisproto.CreateAggregationEventResponse{}
	aggregationEventResponse.AggregationEvent = &aggregationEvent
	return &aggregationEventResponse, nil
}

func (es *EpcisService) insertAggregationEvent(ctx context.Context, insertAggregationEventSQL string, aggregationEvent *epcisproto.AggregationEvent, insertPersistentDispositionSQL string, persistentDispositions []*epcisproto.PersistentDisposition, insertEpcSQL string, childEpcs []*epcisproto.Epc, insertBizTransactionSQL string, bizTransactionList []*epcisproto.BizTransaction, insertQuantityElementSQL string, childQuantityList []*epcisproto.QuantityElement, insertSourceSQL string, sourceList []*epcisproto.Source, insertDestinationSQL string, destinationList []*epcisproto.Destination, insertSensorElementSQL string, sensorElementList []*epcisproto.SensorElement, userEmail string, requestID string,
) error {
	aggregationEventTmp, err := es.crAggregationEventStruct(ctx, aggregationEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertAggregationEventSQL, aggregationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		aggregationEvent.AggregationEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(aggregationEvent.AggregationEventD.Uuid4)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		aggregationEvent.AggregationEventD.IdS = uuid4Str

		for _, persistentDisposition := range persistentDispositions {
			persistentDisposition.EventId = aggregationEvent.AggregationEventD.Id
			persistentDisposition.TypeOfEvent = "AggregationEvent"
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
			epc.EventId = aggregationEvent.AggregationEventD.Id
			epc.TypeOfEvent = "AggregationEvent"
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
			bizTransaction.EventId = aggregationEvent.AggregationEventD.Id
			bizTransaction.TypeOfEvent = "AggregationEvent"
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

		for _, quantityElement := range childQuantityList {
			quantityElement.EventId = aggregationEvent.AggregationEventD.Id
			quantityElement.TypeOfEvent = "AggregationEvent"
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
			source.EventId = aggregationEvent.AggregationEventD.Id
			source.TypeOfEvent = "AggregationEvent"
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
			destination.EventId = aggregationEvent.AggregationEventD.Id
			destination.TypeOfEvent = "AggregationEvent"
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
			sensorElement.SensorElementD.EventId = aggregationEvent.AggregationEventD.Id
			sensorElement.SensorElementD.TypeOfEvent = "AggregationEvent"

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

// crAggregationEventStruct - process AggregationEvent details
func (es *EpcisService) crAggregationEventStruct(ctx context.Context, aggregationEvent *epcisproto.AggregationEvent, userEmail string, requestID string) (*epcisstruct.AggregationEvent, error) {
	epcisEventT := new(commonstruct.EpcisEventT)
	epcisEventT.EventTime = common.TimestampToTime(aggregationEvent.EpcisEventT.EventTime)

	errorDeclarationT := new(commonstruct.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimestampToTime(aggregationEvent.ErrorDeclarationT.DeclarationTime)

	aggregationEventTmp := epcisstruct.AggregationEvent{EpcisEventD: aggregationEvent.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: aggregationEvent.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, AggregationEventD: aggregationEvent.AggregationEventD, CrUpdUser: aggregationEvent.CrUpdUser}

	return &aggregationEventTmp, nil
}

func (es *EpcisService) GetAggregationEvents(ctx context.Context, in *epcisproto.GetAggregationEventsRequest) (*epcisproto.GetAggregationEventsResponse, error) {
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

	aggregationEvents := []*epcisproto.AggregationEvent{}

	nselectAggregationEventsSQL := selectAggregationEventsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectAggregationEventsSQL, common.Active)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		aggregationEventTmp := epcisstruct.AggregationEvent{}
		err = rows.StructScan(&aggregationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		aggregationEvent, err := es.getAggregationEventStruct(ctx, &getRequest, aggregationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		aggregationEvents = append(aggregationEvents, aggregationEvent)
	}

	aggregationEventsResponse := epcisproto.GetAggregationEventsResponse{}
	if len(aggregationEvents) != 0 {
		next := aggregationEvents[len(aggregationEvents)-1].AggregationEventD.Id
		next--
		nextc := common.EncodeCursor(next)
		aggregationEventsResponse = epcisproto.GetAggregationEventsResponse{AggregationEvents: aggregationEvents, NextCursor: nextc}
	} else {
		aggregationEventsResponse = epcisproto.GetAggregationEventsResponse{AggregationEvents: aggregationEvents, NextCursor: "0"}
	}
	return &aggregationEventsResponse, nil
}

// GetAggregationEvent - Get AggregationEvent
func (es *EpcisService) GetAggregationEvent(ctx context.Context, inReq *epcisproto.GetAggregationEventRequest) (*epcisproto.GetAggregationEventResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectAggregationEventsSQL := selectAggregationEventsSQL + ` where uuid4 = ?;`
	row := es.DBService.DB.QueryRowxContext(ctx, nselectAggregationEventsSQL, uuid4byte)
	aggregationEventTmp := epcisstruct.AggregationEvent{}
	err = row.StructScan(&aggregationEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	aggregationEvent, err := es.getAggregationEventStruct(ctx, in, aggregationEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	aggregationEventResponse := epcisproto.GetAggregationEventResponse{}
	aggregationEventResponse.AggregationEvent = aggregationEvent
	return &aggregationEventResponse, nil
}

// getAggregationEventStruct - Get AggregationEvent
func (es *EpcisService) getAggregationEventStruct(ctx context.Context, in *commonproto.GetRequest, aggregationEventTmp epcisstruct.AggregationEvent) (*epcisproto.AggregationEvent, error) {
	epcisEventT := new(commonproto.EpcisEventT)
	epcisEventT.EventTime = common.TimeToTimestamp(aggregationEventTmp.EpcisEventT.EventTime)

	errorDeclarationT := new(commonproto.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimeToTimestamp(aggregationEventTmp.ErrorDeclarationT.DeclarationTime)

	uuid4Str, err := common.UUIDBytesToStr(aggregationEventTmp.AggregationEventD.Uuid4)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	aggregationEventTmp.AggregationEventD.IdS = uuid4Str

	aggregationEvent := epcisproto.AggregationEvent{EpcisEventD: aggregationEventTmp.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: aggregationEventTmp.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, AggregationEventD: aggregationEventTmp.AggregationEventD, CrUpdUser: aggregationEventTmp.CrUpdUser}

	return &aggregationEvent, nil
}
