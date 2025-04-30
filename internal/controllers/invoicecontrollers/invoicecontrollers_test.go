package invoicecontrollers

import (
	"net/http"
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"github.com/cloudfresco/sc-gs1/internal/controllers/partycontrollers"
	invoiceservices "github.com/cloudfresco/sc-gs1/internal/services/invoiceservices"
	partyservices "github.com/cloudfresco/sc-gs1/internal/services/partyservices"

	invoiceworkers "github.com/cloudfresco/sc-gs1/internal/workers/invoiceworkers"
	"github.com/cloudfresco/sc-gs1/test"

	"github.com/throttled/throttled/v2/store/goredisstore"
)

var (
	dbService         *common.DBService
	redisService      *common.RedisService
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
	mux               *http.ServeMux
	log               *zap.Logger
	logInvoice        *zap.Logger
	logUser           *zap.Logger
	backendServerAddr string
)

func TestMain(m *testing.M) {
	var err error
	v, err := config.GetViper()
	if err != nil {
		os.Exit(1)
	}

	configFilePath := v.GetString("SC_GS1_WORKFLOW_CONFIG_FILE_PATH")

	logOpt, err := config.GetLogConfig(v)
	if err != nil {
		os.Exit(1)
	}

	log = config.SetUpLogging(logOpt.Path)
	logInvoice = config.SetUpLogging(logOpt.InvoicePath)
	logUser = config.SetUpLogging(logOpt.UserPath)

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

	backendServerAddr = serverOpt.BackendServerAddr

	pwd, _ := os.Getwd()
	go invoiceservices.StartInvoiceServer(logInvoice, true, pwd, dbOpt, redisOpt, mailerOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)
	go partyservices.StartUserServer(logUser, true, pwd, dbOpt, redisOpt, mailerOpt, serverOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)
	go invoiceworkers.StartInvoiceWorker(log, true, pwd, grpcServerOpt, configFilePath)

	store, err := goredisstore.New(redisService.RedisClient, "throttled:")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	mux = http.NewServeMux()
	err = InitTest(log, mux, store, serverOpt, grpcServerOpt, uptraceOpt, configFilePath)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}
	err = partycontrollers.InitTest(log, mux, store, serverOpt, grpcServerOpt, uptraceOpt, configFilePath)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}
	os.Exit(m.Run())
}

func LoginUser() (string, string, string) {
	addr := "http://" + backendServerAddr
	return userTestOpt.Tokenstring, userTestOpt.Email, addr
}
