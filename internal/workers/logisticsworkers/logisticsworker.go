package logisticsworkers

import (
	"os"

	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-gs1/internal/common"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"

	logisticsworkflows "github.com/cloudfresco/sc-gs1/internal/workflows/logisticsworkflows"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.WfHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, logisticsworkflows.ApplicationName, workerOptions)
}

func StartLogisticsWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	logisticsconn, err := grpc.NewClient(grpcServerOpt.GrpcLogisticsServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	receivingAdviceServiceClient := logisticsproto.NewReceivingAdviceServiceClient(logisticsconn)
	receivingAdviceActivities := &logisticsworkflows.ReceivingAdviceActivities{ReceivingAdviceServiceClient: receivingAdviceServiceClient}

	despatchAdviceServiceClient := logisticsproto.NewDespatchAdviceServiceClient(logisticsconn)
	despatchAdviceActivities := &logisticsworkflows.DespatchAdviceActivities{DespatchAdviceServiceClient: despatchAdviceServiceClient}

	h.RegisterWorkflow(logisticsworkflows.CreateReceivingAdviceWorkflow)
	h.RegisterWorkflow(logisticsworkflows.UpdateReceivingAdviceWorkflow)
	h.RegisterWorkflow(logisticsworkflows.CreateDespatchAdviceWorkflow)
	h.RegisterWorkflow(logisticsworkflows.UpdateDespatchAdviceWorkflow)
	h.RegisterActivity(receivingAdviceActivities)
	h.RegisterActivity(despatchAdviceActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
