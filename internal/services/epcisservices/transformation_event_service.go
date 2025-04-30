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

const insertTransformationEventSQL = `insert into transformation_events
	  (uuid4,
    event_id,
    event_time_zone_offset,
    certification,
    event_time,
    reason,
    declaration_time,
    transformation_id,
    biz_step,
    disposition,
    read_point,
    biz_location,
    ilmd,
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
    :transformation_id,
    :biz_step,
    :disposition,
    :read_point,
    :biz_location,
    :ilmd,
    :status_code,
    :created_by_user_id,
    :updated_by_user_id);`

const selectTransformationEventsSQL = `select
  id,
  uuid4,
  event_id,
  event_time_zone_offset,
  certification,
  event_time,
  reason,
  declaration_time,
  transformation_id,
  biz_step,
  disposition,
  read_point,
  biz_location,
  ilmd,
  status_code,
  created_by_user_id,
  updated_by_user_id from transformation_events`

// CreateTransformationEvent - Create TransformationEvent
func (es *EpcisService) CreateTransformationEvent(ctx context.Context, in *epcisproto.CreateTransformationEventRequest) (*epcisproto.CreateTransformationEventResponse, error) {
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

	transformationEventD := epcisproto.TransformationEventD{}
	transformationEventD.Uuid4, err = common.GetUUIDBytes()
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

	transformationEventD.TransformationId = in.TransformationId
	transformationEventD.BizStep = in.BizStep
	transformationEventD.Disposition = in.Disposition
	transformationEventD.ReadPoint = in.ReadPoint
	transformationEventD.BizLocation = in.BizLocation
	transformationEventD.Ilmd = in.Ilmd

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	transformationEvent := epcisproto.TransformationEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, TransformationEventD: &transformationEventD, CrUpdUser: &crUpdUser}

	inEpcList := []*epcisproto.Epc{}
	for _, inEpc := range in.InputEpcList {
		inEpc.TypeOfEpc = "input"
		inEpc.UserId = in.UserId
		inEpc.UserEmail = in.UserEmail
		inEpc.RequestId = in.RequestId

		inEp, err := es.ProcessEpcRequest(ctx, inEpc)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		inEpcList = append(inEpcList, inEp)
	}

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

	outEpcList := []*epcisproto.Epc{}
	for _, outEpc := range in.OutputEpcList {
		outEpc.TypeOfEpc = "output"
		outEpc.UserId = in.UserId
		outEpc.UserEmail = in.UserEmail
		outEpc.RequestId = in.RequestId

		outEp, err := es.ProcessEpcRequest(ctx, outEpc)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		outEpcList = append(outEpcList, outEp)
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

	inQuantityList := []*epcisproto.QuantityElement{}
	for _, inQuantityElement := range in.InputQuantityList {
		inQuantityElement.TypeOfQuantity = "input"
		inQuantityElement.UserEmail = in.UserEmail
		inQuantityElement.RequestId = in.RequestId

		qElement, err := es.ProcessQuantityElementRequest(ctx, inQuantityElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		inQuantityList = append(inQuantityList, qElement)
	}

	outQuantityList := []*epcisproto.QuantityElement{}
	for _, outQuantityElement := range in.OutputQuantityList {
		outQuantityElement.TypeOfQuantity = "input"
		outQuantityElement.UserEmail = in.UserEmail
		outQuantityElement.RequestId = in.RequestId

		qElement, err := es.ProcessQuantityElementRequest(ctx, outQuantityElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		outQuantityList = append(outQuantityList, qElement)
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

	err = es.insertTransformationEvent(ctx, insertTransformationEventSQL, &transformationEvent, insertPersistentDispositionSQL, persistentDispositions, insertEpcSQL, inEpcList, outEpcList, insertBizTransactionSQL, bizTransactionList, insertQuantityElementSQL, inQuantityList, outQuantityList, insertSourceSQL, sourceList, insertDestinationSQL, destinationList, insertSensorElementSQL, sensorElementList, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transformationEventResponse := epcisproto.CreateTransformationEventResponse{}
	transformationEventResponse.TransformationEvent = &transformationEvent
	return &transformationEventResponse, nil
}

func (es *EpcisService) insertTransformationEvent(ctx context.Context, insertTransformationEventSQL string, transformationEvent *epcisproto.TransformationEvent, insertPersistentDispositionSQL string, persistentDispositions []*epcisproto.PersistentDisposition, insertEpcSQL string, inEpcList []*epcisproto.Epc, outEpcList []*epcisproto.Epc, insertBizTransactionSQL string, bizTransactionList []*epcisproto.BizTransaction, insertQuantityElementSQL string, inQuantityList []*epcisproto.QuantityElement, outQuantityList []*epcisproto.QuantityElement, insertSourceSQL string, sourceList []*epcisproto.Source, insertDestinationSQL string, destinationList []*epcisproto.Destination, insertSensorElementSQL string, sensorElementList []*epcisproto.SensorElement, userEmail string, requestID string,
) error {
	transformationEventTmp, err := es.crTransformationEventStruct(ctx, transformationEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransformationEventSQL, transformationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transformationEvent.TransformationEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transformationEvent.TransformationEventD.Uuid4)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transformationEvent.TransformationEventD.IdS = uuid4Str

		for _, persistentDisposition := range persistentDispositions {
			persistentDisposition.EventId = transformationEvent.TransformationEventD.Id
			persistentDisposition.TypeOfEvent = "TransformationEvent"
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

		for _, inEpc := range inEpcList {
			inEpc.EventId = transformationEvent.TransformationEventD.Id
			inEpc.TypeOfEvent = "TransformationEvent"
			inEpc.TypeOfEpc = "input"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertEpcSQL, inEpc)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, outEpc := range outEpcList {
			outEpc.EventId = transformationEvent.TransformationEventD.Id
			outEpc.TypeOfEvent = "TransformationEvent"
			outEpc.TypeOfEpc = "output"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertEpcSQL, outEpc)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, bizTransaction := range bizTransactionList {
			bizTransaction.EventId = transformationEvent.TransformationEventD.Id
			bizTransaction.TypeOfEvent = "TransformationEvent"
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

		for _, inQuantityElement := range inQuantityList {
			inQuantityElement.EventId = transformationEvent.TransformationEventD.Id
			inQuantityElement.TypeOfEvent = "TransformationEvent"
			inQuantityElement.TypeOfQuantity = "input"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertQuantityElementSQL, inQuantityElement)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, outQuantityElement := range outQuantityList {
			outQuantityElement.EventId = transformationEvent.TransformationEventD.Id
			outQuantityElement.TypeOfEvent = "TransformationEvent"
			outQuantityElement.TypeOfQuantity = "output"
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertQuantityElementSQL, outQuantityElement)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, source := range sourceList {
			source.EventId = transformationEvent.TransformationEventD.Id
			source.TypeOfEvent = "TransformationEvent"
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
			destination.EventId = transformationEvent.TransformationEventD.Id
			destination.TypeOfEvent = "TransformationEvent"
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
			sensorElement.SensorElementD.EventId = transformationEvent.TransformationEventD.Id
			sensorElement.SensorElementD.TypeOfEvent = "TransformationEvent"

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

// crTransformationEventStruct - process TransformationEvent details
func (es *EpcisService) crTransformationEventStruct(ctx context.Context, transformationEvent *epcisproto.TransformationEvent, userEmail string, requestID string) (*epcisstruct.TransformationEvent, error) {
	epcisEventT := new(commonstruct.EpcisEventT)
	epcisEventT.EventTime = common.TimestampToTime(transformationEvent.EpcisEventT.EventTime)

	errorDeclarationT := new(commonstruct.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimestampToTime(transformationEvent.ErrorDeclarationT.DeclarationTime)

	transformationEventTmp := epcisstruct.TransformationEvent{EpcisEventD: transformationEvent.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: transformationEvent.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, TransformationEventD: transformationEvent.TransformationEventD, CrUpdUser: transformationEvent.CrUpdUser}

	return &transformationEventTmp, nil
}

func (es *EpcisService) GetTransformationEvents(ctx context.Context, in *epcisproto.GetTransformationEventsRequest) (*epcisproto.GetTransformationEventsResponse, error) {
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

	transformationEvents := []*epcisproto.TransformationEvent{}

	nselectTransformationEventsSQL := selectTransformationEventsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectTransformationEventsSQL, common.Active)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transformationEventTmp := epcisstruct.TransformationEvent{}
		err = rows.StructScan(&transformationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transformationEvent, err := es.getTransformationEventStruct(ctx, &getRequest, transformationEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		transformationEvents = append(transformationEvents, transformationEvent)
	}

	transformationEventsResponse := epcisproto.GetTransformationEventsResponse{}
	if len(transformationEvents) != 0 {
		next := transformationEvents[len(transformationEvents)-1].TransformationEventD.Id
		next--
		nextc := common.EncodeCursor(next)
		transformationEventsResponse = epcisproto.GetTransformationEventsResponse{TransformationEvents: transformationEvents, NextCursor: nextc}
	} else {
		transformationEventsResponse = epcisproto.GetTransformationEventsResponse{TransformationEvents: transformationEvents, NextCursor: "0"}
	}
	return &transformationEventsResponse, nil
}

// GetTransformationEvent - Get TransformationEvent
func (es *EpcisService) GetTransformationEvent(ctx context.Context, inReq *epcisproto.GetTransformationEventRequest) (*epcisproto.GetTransformationEventResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectTransformationEventsSQL := selectTransformationEventsSQL + ` where uuid4 = ?;`
	row := es.DBService.DB.QueryRowxContext(ctx, nselectTransformationEventsSQL, uuid4byte)
	transformationEventTmp := epcisstruct.TransformationEvent{}
	err = row.StructScan(&transformationEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transformationEvent, err := es.getTransformationEventStruct(ctx, in, transformationEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	transformationEventResponse := epcisproto.GetTransformationEventResponse{}
	transformationEventResponse.TransformationEvent = transformationEvent
	return &transformationEventResponse, nil
}

// getTransformationEventStruct - Get TransformationEvent
func (es *EpcisService) getTransformationEventStruct(ctx context.Context, in *commonproto.GetRequest, transformationEventTmp epcisstruct.TransformationEvent) (*epcisproto.TransformationEvent, error) {
	epcisEventT := new(commonproto.EpcisEventT)
	epcisEventT.EventTime = common.TimeToTimestamp(transformationEventTmp.EpcisEventT.EventTime)

	errorDeclarationT := new(commonproto.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimeToTimestamp(transformationEventTmp.ErrorDeclarationT.DeclarationTime)

	uuid4Str, err := common.UUIDBytesToStr(transformationEventTmp.TransformationEventD.Uuid4)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transformationEventTmp.TransformationEventD.IdS = uuid4Str

	transformationEvent := epcisproto.TransformationEvent{EpcisEventD: transformationEventTmp.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: transformationEventTmp.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, TransformationEventD: transformationEventTmp.TransformationEventD, CrUpdUser: transformationEventTmp.CrUpdUser}

	return &transformationEvent, nil
}
