package main

import (
	"os"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"github.com/cloudfresco/sc-gs1/internal/services/invoiceservices"
	_ "github.com/go-sql-driver/mysql" // mysql
	"go.uber.org/zap"
)

func main() {
	v, err := config.GetViper()
	if err != nil {
		os.Exit(1)
	}

	logOpt, err := config.GetLogConfig(v)
	if err != nil {
		os.Exit(1)
	}

	log := config.SetUpLogging(logOpt.InvoicePath)

	dbOpt, err := config.GetDbConfig(log, v, false, "SC_GS1_DB", "SC_GS1_DBHOST", "SC_GS1_DBPORT", "SC_GS1_DBUSER", "SC_GS1_DBPASS", "SC_GS1_DBNAME", "", "", "", "", "", "")
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	jwtOpt, err := config.GetJWTConfig(log, v, false, "SC_GS1_JWT_KEY", "SC_GS1_JWT_DURATION")
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	redisOpt, mailerOpt, _, grpcServerOpt, oauthOpt, userOpt, uptraceOpt := config.GetConfigOpt(log, v)

	dbService, redisService, mailerService := common.GetServices(log, false, dbOpt, redisOpt, jwtOpt, mailerOpt)

	pwd, _ := os.Getwd()
	invoiceservices.StartInvoiceServer(log, false, pwd, dbOpt, redisOpt, mailerOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)
}
