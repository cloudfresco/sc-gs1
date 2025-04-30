package meatservices

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	meatstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// MeatService - For accessing MeatActivityHistory services
type MeatService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	meatproto.UnimplementedMeatServiceServer
}

// NewMeatService - Create MeatActivityHistory service
func NewMeatService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *MeatService {
	return &MeatService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

const insertMeatActivityHistorySQL = `insert into meat_activity_histories
	  (
uuid4,
activity_sub_step_identification,
country_of_activity_code,
current_step_identification,
meat_processing_activity_type_code,
movement_reason_code,
next_step_identification,
meat_mincing_detail_id,
meat_fattening_detail_id,
meat_cutting_detail_id,
meat_breeding_detail_id,
meat_processing_party_id,
meat_work_item_identification_id,
meat_slaughtering_detail_id,
meat_despatch_advice_line_item_extension_id,
date_of_arrival,
date_of_departure,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
)
values(
:uuid4,
:activity_sub_step_identification,
:country_of_activity_code,
:current_step_identification,
:meat_processing_activity_type_code,
:movement_reason_code,
:next_step_identification,
:meat_mincing_detail_id,
:meat_fattening_detail_id,
:meat_cutting_detail_id,
:meat_breeding_detail_id,
:meat_processing_party_id,
:meat_work_item_identification_id,
:meat_slaughtering_detail_id,
:meat_despatch_advice_line_item_extension_id,
:date_of_arrival,
:date_of_departure,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectMeatActivityHistoriesSQL = `select
 id,
uuid4,
activity_sub_step_identification,
country_of_activity_code,
current_step_identification,
meat_processing_activity_type_code,
movement_reason_code,
next_step_identification,
meat_mincing_detail_id,
meat_fattening_detail_id,
meat_cutting_detail_id,
meat_breeding_detail_id,
meat_processing_party_id,
meat_work_item_identification_id,
meat_slaughtering_detail_id,
meat_despatch_advice_line_item_extension_id,
date_of_arrival,
date_of_departure,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from meat_activity_histories`*/

// StartMeatServer - Start Meat server
func StartMeatServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	fishService := NewFishService(log, dbService, redisService, uc)
	meatService := NewMeatService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcMeatServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	meatproto.RegisterFishServiceServer(srv, fishService)
	meatproto.RegisterMeatServiceServer(srv, meatService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

func (ms *MeatService) CreateMeatActivityHistory(ctx context.Context, in *meatproto.CreateMeatActivityHistoryRequest) (*meatproto.CreateMeatActivityHistoryResponse, error) {
	meatActivityHistory, err := ms.ProcessMeatActivityHistoryRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatActivityHistory(ctx, insertMeatActivityHistorySQL, meatActivityHistory, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatActivityHistoryResponse := meatproto.CreateMeatActivityHistoryResponse{}
	meatActivityHistoryResponse.MeatActivityHistory = meatActivityHistory

	return &meatActivityHistoryResponse, nil
}

// ProcessMeatActivityHistoryRequest - ProcessMeatActivityHistoryRequest
func (ms *MeatService) ProcessMeatActivityHistoryRequest(ctx context.Context, in *meatproto.CreateMeatActivityHistoryRequest) (*meatproto.MeatActivityHistory, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ms.UserServiceClient)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	dateOfArrival, err := time.Parse(common.Layout, in.DateOfArrival)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	dateOfDeparture, err := time.Parse(common.Layout, in.DateOfDeparture)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatActivityHistoryD := meatproto.MeatActivityHistoryD{}
	meatActivityHistoryD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatActivityHistoryD.ActivitySubStepIdentification = in.ActivitySubStepIdentification
	meatActivityHistoryD.CountryOfActivityCode = in.CountryOfActivityCode
	meatActivityHistoryD.CurrentStepIdentification = in.CurrentStepIdentification
	meatActivityHistoryD.MeatProcessingActivityTypeCode = in.MeatProcessingActivityTypeCode
	meatActivityHistoryD.MovementReasonCode = in.MovementReasonCode
	meatActivityHistoryD.NextStepIdentification = in.NextStepIdentification
	meatActivityHistoryD.MeatMincingDetailId = in.MeatMincingDetailId
	meatActivityHistoryD.MeatFatteningDetailId = in.MeatFatteningDetailId
	meatActivityHistoryD.MeatCuttingDetailId = in.MeatCuttingDetailId
	meatActivityHistoryD.MeatBreedingDetailId = in.MeatBreedingDetailId
	meatActivityHistoryD.MeatProcessingPartyId = in.MeatProcessingPartyId
	meatActivityHistoryD.MeatWorkItemIdentificationId = in.MeatWorkItemIdentificationId
	meatActivityHistoryD.MeatSlaughteringDetailId = in.MeatSlaughteringDetailId
	meatActivityHistoryD.MeatDespatchAdviceLineItemExtensionId = in.MeatDespatchAdviceLineItemExtensionId

	meatActivityHistoryT := meatproto.MeatActivityHistoryT{}
	meatActivityHistoryT.DateOfArrival = common.TimeToTimestamp(dateOfArrival.UTC().Truncate(time.Second))
	meatActivityHistoryT.DateOfDeparture = common.TimeToTimestamp(dateOfDeparture.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	meatActivityHistory := meatproto.MeatActivityHistory{MeatActivityHistoryD: &meatActivityHistoryD, MeatActivityHistoryT: &meatActivityHistoryT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &meatActivityHistory, nil
}

// insertMeatActivityHistory - Insert MeatActivityHistory into database
func (ms *MeatService) insertMeatActivityHistory(ctx context.Context, insertMeatActivityHistorySQL string, meatActivityHistory *meatproto.MeatActivityHistory, userEmail string, requestID string) error {
	meatActivityHistoryTmp, err := ms.crMeatActivityHistoryStruct(ctx, meatActivityHistory, userEmail, requestID)
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatActivityHistorySQL, meatActivityHistoryTmp)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatActivityHistory.MeatActivityHistoryD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatActivityHistory.MeatActivityHistoryD.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatActivityHistory.MeatActivityHistoryD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crMeatActivityHistoryStruct - process MeatActivityHistory details
func (ms *MeatService) crMeatActivityHistoryStruct(ctx context.Context, meatActivityHistory *meatproto.MeatActivityHistory, userEmail string, requestID string) (*meatstruct.MeatActivityHistory, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(meatActivityHistory.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(meatActivityHistory.CrUpdTime.UpdatedAt)

	meatActivityHistoryT := new(meatstruct.MeatActivityHistoryT)
	meatActivityHistoryT.DateOfArrival = common.TimestampToTime(meatActivityHistory.MeatActivityHistoryT.DateOfArrival)
	meatActivityHistoryT.DateOfDeparture = common.TimestampToTime(meatActivityHistory.MeatActivityHistoryT.DateOfDeparture)

	meatActivityHistoryTmp := meatstruct.MeatActivityHistory{MeatActivityHistoryD: meatActivityHistory.MeatActivityHistoryD, MeatActivityHistoryT: meatActivityHistoryT, CrUpdUser: meatActivityHistory.CrUpdUser, CrUpdTime: crUpdTime}

	return &meatActivityHistoryTmp, nil
}
