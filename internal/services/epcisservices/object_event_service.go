package epcisservices

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	epcisstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const insertObjectEventSQL = `insert into object_events
	  (uuid4,
    event_id,
    event_time_zone_offset,
    certification,
    event_time,
    reason,
    declaration_time,
    action,
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
    :action,
    :biz_step,
    :disposition,
    :read_point,
    :biz_location,
    :ilmd,
    :status_code,
    :created_by_user_id,
    :updated_by_user_id);`

const selectObjectEventsSQL = `select
  id,
  uuid4,
  event_id,
  event_time_zone_offset,
  certification,
  event_time,
  reason,
  declaration_time,
  action,
  biz_step,
  disposition,
  read_point,
  biz_location,
  ilmd,
  status_code,
  created_by_user_id,
  updated_by_user_id from object_events`

// EpcisService - For accessing Epcis services
type EpcisService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	epcisproto.UnimplementedEpcisServiceServer
}

// NewEpcisService - Create Epcis service
func NewEpcisService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *EpcisService {
	return &EpcisService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// StartEpcisServer - Start Epcis server
func StartEpcisServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
	common.SetJWTOpt(jwtOpt)

	creds, err := common.GetSrvCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	userCreds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	var srvOpts []grpc.ServerOption

	userConn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(userCreds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srvOpts = append(srvOpts, grpc.Creds(creds))

	srvOpts = append(srvOpts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	uc := partyproto.NewUserServiceClient(userConn)
	epcisService := NewEpcisService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcEpcisServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	epcisproto.RegisterEpcisServiceServer(srv, epcisService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// CreateObjectEvent - Create ObjectEvent
func (es *EpcisService) CreateObjectEvent(ctx context.Context, in *epcisproto.CreateObjectEventRequest) (*epcisproto.CreateObjectEventResponse, error) {
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

	objectEventD := epcisproto.ObjectEventD{}
	objectEventD.Uuid4, err = common.GetUUIDBytes()
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

	objectEventD.Action = in.Action
	objectEventD.BizStep = in.BizStep
	objectEventD.Disposition = in.Disposition
	objectEventD.ReadPoint = in.ReadPoint
	objectEventD.BizLocation = in.BizLocation
	objectEventD.Ilmd = in.Ilmd

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	objectEvent := epcisproto.ObjectEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, ObjectEventD: &objectEventD, CrUpdUser: &crUpdUser}

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

	err = es.insertObjectEvent(ctx, insertObjectEventSQL, &objectEvent, insertPersistentDispositionSQL, persistentDispositions, insertEpcSQL, epcList, insertBizTransactionSQL, bizTransactionList, insertQuantityElementSQL, quantityList, insertSourceSQL, sourceList, insertDestinationSQL, destinationList, insertSensorElementSQL, sensorElementList, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	objectEventResponse := epcisproto.CreateObjectEventResponse{}
	objectEventResponse.ObjectEvent = &objectEvent
	return &objectEventResponse, nil
}

func (es *EpcisService) insertObjectEvent(ctx context.Context, insertObjectEventSQL string, objectEvent *epcisproto.ObjectEvent,
	insertPersistentDispositionSQL string, persistentDispositions []*epcisproto.PersistentDisposition, insertEpcSQL string, epcList []*epcisproto.Epc, insertBizTransactionSQL string, bizTransactionList []*epcisproto.BizTransaction, insertQuantityElementSQL string, quantityList []*epcisproto.QuantityElement, insertSourceSQL string, sourceList []*epcisproto.Source, insertDestinationSQL string, destinationList []*epcisproto.Destination, insertSensorElementSQL string, sensorElementList []*epcisproto.SensorElement, userEmail string, requestID string,
) error {
	objectEventTmp, err := es.crObjectEventStruct(ctx, objectEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertObjectEventSQL, objectEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		objectEvent.ObjectEventD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(objectEvent.ObjectEventD.Uuid4)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		objectEvent.ObjectEventD.IdS = uuid4Str

		for _, persistentDisposition := range persistentDispositions {
			persistentDisposition.EventId = objectEvent.ObjectEventD.Id
			persistentDisposition.TypeOfEvent = "ObjectEvent"
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
			epc.EventId = objectEvent.ObjectEventD.Id
			epc.TypeOfEvent = "ObjectEvent"
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
			bizTransaction.EventId = objectEvent.ObjectEventD.Id
			bizTransaction.TypeOfEvent = "ObjectEvent"
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
			quantityElement.EventId = objectEvent.ObjectEventD.Id
			quantityElement.TypeOfEvent = "ObjectEvent"
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
			source.EventId = objectEvent.ObjectEventD.Id
			source.TypeOfEvent = "ObjectEvent"
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
			destination.EventId = objectEvent.ObjectEventD.Id
			destination.TypeOfEvent = "ObjectEvent"
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
			sensorElement.SensorElementD.EventId = objectEvent.ObjectEventD.Id
			sensorElement.SensorElementD.TypeOfEvent = "ObjectEvent"

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

// crObjectEventStruct - process ObjectEvent details
func (es *EpcisService) crObjectEventStruct(ctx context.Context, objectEvent *epcisproto.ObjectEvent, userEmail string, requestID string) (*epcisstruct.ObjectEvent, error) {
	epcisEventT := new(commonstruct.EpcisEventT)
	epcisEventT.EventTime = common.TimestampToTime(objectEvent.EpcisEventT.EventTime)

	errorDeclarationT := new(commonstruct.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimestampToTime(objectEvent.ErrorDeclarationT.DeclarationTime)

	objectEventTmp := epcisstruct.ObjectEvent{EpcisEventD: objectEvent.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: objectEvent.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, ObjectEventD: objectEvent.ObjectEventD, CrUpdUser: objectEvent.CrUpdUser}

	return &objectEventTmp, nil
}

func (es *EpcisService) GetObjectEvents(ctx context.Context, in *epcisproto.GetObjectEventsRequest) (*epcisproto.GetObjectEventsResponse, error) {
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

	objectEvents := []*epcisproto.ObjectEvent{}

	nselectObjectEventsSQL := selectObjectEventsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectObjectEventsSQL, common.Active)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		objectEventTmp := epcisstruct.ObjectEvent{}
		err = rows.StructScan(&objectEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		objectEvent, err := es.getObjectEventStruct(ctx, &getRequest, objectEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		objectEvents = append(objectEvents, objectEvent)
	}

	objectEventsResponse := epcisproto.GetObjectEventsResponse{}
	if len(objectEvents) != 0 {
		next := objectEvents[len(objectEvents)-1].ObjectEventD.Id
		next--
		nextc := common.EncodeCursor(next)
		objectEventsResponse = epcisproto.GetObjectEventsResponse{ObjectEvents: objectEvents, NextCursor: nextc}
	} else {
		objectEventsResponse = epcisproto.GetObjectEventsResponse{ObjectEvents: objectEvents, NextCursor: "0"}
	}
	return &objectEventsResponse, nil
}

// GetObjectEvent - Get ObjectEvent
func (es *EpcisService) GetObjectEvent(ctx context.Context, inReq *epcisproto.GetObjectEventRequest) (*epcisproto.GetObjectEventResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectObjectEventsSQL := selectObjectEventsSQL + ` where uuid4 = ?;`
	row := es.DBService.DB.QueryRowxContext(ctx, nselectObjectEventsSQL, uuid4byte)
	objectEventTmp := epcisstruct.ObjectEvent{}
	err = row.StructScan(&objectEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	objectEvent, err := es.getObjectEventStruct(ctx, in, objectEventTmp)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	objectEventResponse := epcisproto.GetObjectEventResponse{}
	objectEventResponse.ObjectEvent = objectEvent
	return &objectEventResponse, nil
}

// getObjectEventStruct - Get ObjectEvent
func (es *EpcisService) getObjectEventStruct(ctx context.Context, in *commonproto.GetRequest, objectEventTmp epcisstruct.ObjectEvent) (*epcisproto.ObjectEvent, error) {
	epcisEventT := new(commonproto.EpcisEventT)
	epcisEventT.EventTime = common.TimeToTimestamp(objectEventTmp.EpcisEventT.EventTime)

	errorDeclarationT := new(commonproto.ErrorDeclarationT)
	errorDeclarationT.DeclarationTime = common.TimeToTimestamp(objectEventTmp.ErrorDeclarationT.DeclarationTime)

	uuid4Str, err := common.UUIDBytesToStr(objectEventTmp.ObjectEventD.Uuid4)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	objectEventTmp.ObjectEventD.IdS = uuid4Str

	objectEvent := epcisproto.ObjectEvent{EpcisEventD: objectEventTmp.EpcisEventD, EpcisEventT: epcisEventT, ErrorDeclarationD: objectEventTmp.ErrorDeclarationD, ErrorDeclarationT: errorDeclarationT, ObjectEventD: objectEventTmp.ObjectEventD, CrUpdUser: objectEventTmp.CrUpdUser}

	return &objectEvent, nil
}
