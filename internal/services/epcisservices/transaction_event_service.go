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

const insertTransactionEventSQL = `insert into transaction_events
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

const selectTransactionEventsSQL = `select
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
  updated_by_user_id from transaction_events`

// CreateTransactionEvent - Create TransactionEvent
func (es *EpcisService) CreateTransactionEvent(ctx context.Context, in *epcisproto.CreateTransactionEventRequest) (*epcisproto.CreateTransactionEventResponse, error) {
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

	transactionEventD := epcisproto.TransactionEventD{}
	transactionEventD.Uuid4, err = common.GetUUIDBytes()
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

	transactionEventD.ParentId = in.ParentId
	transactionEventD.Action = in.Action
	transactionEventD.BizStep = in.BizStep
	transactionEventD.Disposition = in.Disposition
	transactionEventD.ReadPoint = in.ReadPoint
	transactionEventD.BizLocation = in.BizLocation

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	transactionEvent := epcisproto.TransactionEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, TransactionEventD: &transactionEventD, CrUpdUser: &crUpdUser}

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

	epcList := []*epcisproto.Epc{}
	for _, epc := range in.EpcList {
		epc.UserId = in.UserId
		epc.UserEmail = in.UserEmail
		epc.RequestId = in.RequestId

		ep, err := es.ProcessEpcRequest(ctx, epc)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		epcList = append(epcList, ep)
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

	quantityList := []*epcisproto.QuantityElement{}
	for _, quantityElement := range in.QuantityList {
		quantityElement.UserId = in.UserId
		quantityElement.UserEmail = in.UserEmail
		quantityElement.RequestId = in.RequestId

		qElement, err := es.ProcessQuantityElementRequest(ctx, quantityElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		quantityList = append(quantityList, qElement)
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

	err = es.insertTransactionEvent(ctx, insertTransactionEventSQL, &transactionEvent, insertPersistentDispositionSQL, persistentDispositions, insertEpcSQL, epcList, insertBizTransactionSQL, bizTransactionList, insertQuantityElementSQL, quantityList, insertSourceSQL, sourceList, insertDestinationSQL, destinationList, insertSensorElementSQL, sensorElementList, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionEventResponse := epcisproto.CreateTransactionEventResponse{}
	transactionEventResponse.TransactionEvent = &transactionEvent
	return &transactionEventResponse, nil
}

func (es *EpcisService) insertTransactionEvent(ctx context.Context, insertTransactionEventSQL string, transactionEvent *epcisproto.TransactionEvent,
	insertPersistentDispositionSQL string, persistentDispositions []*epcisproto.PersistentDisposition, insertEpcSQL string, epcList []*epcisproto.Epc, insertBizTransactionSQL string, bizTransactionList []*epcisproto.BizTransaction, insertQuantityElementSQL string, quantityList []*epcisproto.QuantityElement, insertSourceSQL string, sourceList []*epcisproto.Source, insertDestinationSQL string, destinationList []*epcisproto.Destination, insertSensorElementSQL string, sensorElementList []*epcisproto.SensorElement, userEmail string, requestID string,
) error {
	transactionEventTmp, err := es.crTransactionEventStruct(ctx, transactionEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionEventSQL, transactionEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionEvent.TransactionEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transactionEvent.TransactionEventD.Uuid4)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionEvent.TransactionEventD.IdS = uuid4Str

		for _, persistentDisposition := range persistentDispositions {
			persistentDisposition.EventId = transactionEvent.TransactionEventD.Id
			persistentDisposition.TypeOfEvent = "TransactionEvent"
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

		for _, epc := range epcList {
			epc.EventId = transactionEvent.TransactionEventD.Id
			epc.TypeOfEvent = "TransactionEvent"
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
			bizTransaction.EventId = transactionEvent.TransactionEventD.Id
			bizTransaction.TypeOfEvent = "TransactionEvent"
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

		for _, quantityElement := range quantityList {
			quantityElement.EventId = transactionEvent.TransactionEventD.Id
			quantityElement.TypeOfEvent = "TransactionEvent"
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
			source.EventId = transactionEvent.TransactionEventD.Id
			source.TypeOfEvent = "TransactionEvent"
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
			destination.EventId = transactionEvent.TransactionEventD.Id
			destination.TypeOfEvent = "TransactionEvent"
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
			sensorElement.SensorElementD.EventId = transactionEvent.TransactionEventD.Id
			sensorElement.SensorElementD.TypeOfEvent = "TransactionEvent"

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

// crTransactionEventStruct - process TransactionEvent details
func (es *EpcisService) crTransactionEventStruct(ctx context.Context, transactionEvent *epcisproto.TransactionEvent, userEmail string, requestID string) (*epcisstruct.TransactionEvent, error) {
	epcisEventT := new(commonstruct.EpcisEventT)
	epcisEventT.EventTime = common.TimestampToTime(transactionEvent.EpcisEventT.EventTime)

	errorDeclarationT := new(commonstruct.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimestampToTime(transactionEvent.ErrorDeclarationT.DeclarationTime)

	transactionEventTmp := epcisstruct.TransactionEvent{EpcisEventD: transactionEvent.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: transactionEvent.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, TransactionEventD: transactionEvent.TransactionEventD, CrUpdUser: transactionEvent.CrUpdUser}

	return &transactionEventTmp, nil
}

func (es *EpcisService) GetTransactionEvents(ctx context.Context, in *epcisproto.GetTransactionEventsRequest) (*epcisproto.GetTransactionEventsResponse, error) {
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

	transactionEvents := []*epcisproto.TransactionEvent{}

	nselectTransactionEventsSQL := selectTransactionEventsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectTransactionEventsSQL, common.Active)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transactionEventTmp := epcisstruct.TransactionEvent{}
		err = rows.StructScan(&transactionEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transactionEvent, err := es.getTransactionEventStruct(ctx, &getRequest, transactionEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		transactionEvents = append(transactionEvents, transactionEvent)
	}

	transactionEventsResponse := epcisproto.GetTransactionEventsResponse{}
	if len(transactionEvents) != 0 {
		next := transactionEvents[len(transactionEvents)-1].TransactionEventD.Id
		next--
		nextc := common.EncodeCursor(next)
		transactionEventsResponse = epcisproto.GetTransactionEventsResponse{TransactionEvents: transactionEvents, NextCursor: nextc}
	} else {
		transactionEventsResponse = epcisproto.GetTransactionEventsResponse{TransactionEvents: transactionEvents, NextCursor: "0"}
	}
	return &transactionEventsResponse, nil
}

// GetTransactionEvent - Get TransactionEvent
func (es *EpcisService) GetTransactionEvent(ctx context.Context, inReq *epcisproto.GetTransactionEventRequest) (*epcisproto.GetTransactionEventResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectTransactionEventsSQL := selectTransactionEventsSQL + ` where uuid4 = ?;`
	row := es.DBService.DB.QueryRowxContext(ctx, nselectTransactionEventsSQL, uuid4byte)
	transactionEventTmp := epcisstruct.TransactionEvent{}
	err = row.StructScan(&transactionEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionEvent, err := es.getTransactionEventStruct(ctx, in, transactionEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	transactionEventResponse := epcisproto.GetTransactionEventResponse{}
	transactionEventResponse.TransactionEvent = transactionEvent
	return &transactionEventResponse, nil
}

// getTransactionEventStruct - Get TransactionEvent
func (es *EpcisService) getTransactionEventStruct(ctx context.Context, in *commonproto.GetRequest, transactionEventTmp epcisstruct.TransactionEvent) (*epcisproto.TransactionEvent, error) {
	epcisEventT := new(commonproto.EpcisEventT)
	epcisEventT.EventTime = common.TimeToTimestamp(transactionEventTmp.EpcisEventT.EventTime)

	errorDeclarationT := new(commonproto.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimeToTimestamp(transactionEventTmp.ErrorDeclarationT.DeclarationTime)

	uuid4Str, err := common.UUIDBytesToStr(transactionEventTmp.TransactionEventD.Uuid4)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionEventTmp.TransactionEventD.IdS = uuid4Str

	transactionEvent := epcisproto.TransactionEvent{EpcisEventD: transactionEventTmp.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: transactionEventTmp.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, TransactionEventD: transactionEventTmp.TransactionEventD, CrUpdUser: transactionEventTmp.CrUpdUser}

	return &transactionEvent, nil
}
