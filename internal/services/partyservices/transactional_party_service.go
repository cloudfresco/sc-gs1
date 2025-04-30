package partyservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const insertTransactionalPartySQL = `insert into transactional_parties
	    (uuid4,
	     gln)
  values(
  :uuid4,
  :gln);`

/*const selectTransactionalPartiesSQL = `select
    id,
    uuid4,
	gln from transactional_parties`*/

// PartyService - For accessing Party services
type PartyService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	partyproto.UnimplementedPartyServiceServer
}

// NewPartyService - Create Party service
func NewPartyService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *PartyService {
	return &PartyService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// StartPartyServer - Start Party server
func StartPartyServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	partyService := NewPartyService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcPartyServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	partyproto.RegisterPartyServiceServer(srv, partyService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

func (ps *PartyService) CreateTransactionalParty(ctx context.Context, in *partyproto.CreateTransactionalPartyRequest) (*partyproto.CreateTransactionalPartyResponse, error) {
	transactionalParty, err := ps.ProcessTransactionalPartyRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertTransactionalParty(ctx, insertTransactionalPartySQL, transactionalParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	transactionalPartyResponse := partyproto.CreateTransactionalPartyResponse{}
	transactionalPartyResponse.TransactionalParty = transactionalParty
	return &transactionalPartyResponse, nil
}

// ProcessTransactionalPartyRequest - ProcessTransactionalPartyRequest
func (ps *PartyService) ProcessTransactionalPartyRequest(ctx context.Context, in *partyproto.CreateTransactionalPartyRequest) (*partyproto.TransactionalParty, error) {
	var err error
	transactionalParty := partyproto.TransactionalParty{}
	transactionalParty.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	transactionalParty.Gln = in.Gln
	return &transactionalParty, nil
}

// insertTransactionalParty - Insert TransactionalParty into database
func (ps *PartyService) insertTransactionalParty(ctx context.Context, insertTransactionalPartySQL string, transactionalParty *partyproto.TransactionalParty, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalPartySQL, transactionalParty)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uuid4Str, err := common.UUIDBytesToStr(transactionalParty.Uuid4)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalParty.IdS = uuid4Str
		transactionalParty.Id = uint32(uID)

		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
