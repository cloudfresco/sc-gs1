package logisticsservices

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	partyservices "github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	"github.com/cloudfresco/sc-gs1/test"
	_ "github.com/go-sql-driver/mysql"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var (
	dbService         *common.DBService
	redisService      *common.RedisService
	userServiceClient partyproto.UserServiceClient
	mailerService     common.MailerIntf
	jwtOpt            *config.JWTOptions
	userTestOpt       *config.UserTestOptions
	redisOpt          *config.RedisOptions
	mailerOpt         *config.MailerOptions
	serverOpt         *config.ServerOptions
	grpcServerOpt     *config.GrpcServerOptions
	oauthOpt          *config.OauthOptions
	userOpt           *config.UserOptions
	uptraceOpt        *config.UptraceOptions
	log               *zap.Logger
	logLogistics      *zap.Logger
	logUser           *zap.Logger
	Layout            string
)

func TestMain(m *testing.M) {
	var err error

	v, err := config.GetViper()
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	logOpt, err := config.GetLogConfig(v)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	log = config.SetUpLogging(logOpt.Path)
	logUser = config.SetUpLogging(logOpt.UserPath)
	logLogistics = config.SetUpLogging(logOpt.LogisticsPath)
	Layout = "2006-01-02T15:04:05Z"

	dbOpt, err := config.GetDbConfig(log, v, true, "SC_GS1_DB", "SC_GS1_DBHOST", "SC_GS1_DBPORT", "SC_GS1_DBUSER_TEST", "SC_GS1_DBPASS_TEST", "SC_GS1_DBNAME_TEST", "SC_GS1_DBSQL_MYSQL_TEST", "SC_GS1_DBSQL_MYSQL_SCHEMA", "SC_GS1_DBSQL_MYSQL_TRUNCATE", "SC_GS1_DBSQL_PGSQL_TEST", "SC_GS1_DBSQL_PGSQL_SCHEMA", "SC_GS1_DBSQL_PGSQL_TRUNCATE")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	jwtOpt, err = config.GetJWTConfig(log, v, true, "SC_GS1_JWT_KEY_TEST", "SC_GS1_JWT_DURATION_TEST")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	userTestOpt, err = config.GetUserTestConfig(log, v)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	redisOpt, mailerOpt, serverOpt, grpcServerOpt, oauthOpt, userOpt, uptraceOpt = config.GetConfigOpt(log, v)

	dbService, redisService, _ = common.GetServices(log, true, dbOpt, redisOpt, jwtOpt, mailerOpt)

	mailerService, err = test.CreateMailerServiceTest(log)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	pwd, _ := os.Getwd()
	go partyservices.StartUserServer(logUser, true, pwd, dbOpt, redisOpt, mailerOpt, serverOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)
	go StartLogisticsServer(logLogistics, true, pwd, dbOpt, redisOpt, mailerOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)

	keyPath := filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCaCertPath))
	creds, err := credentials.NewClientTLSFromFile(keyPath, "localhost")
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	userconn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	userServiceClient = partyproto.NewUserServiceClient(userconn)

	os.Exit(m.Run())
}

func LoginUser() context.Context {
	md := metadata.Pairs("authorization", "Bearer "+userTestOpt.Tokenstring)
	ctxNew := metadata.NewIncomingContext(context.Background(), md)
	return ctxNew
}
