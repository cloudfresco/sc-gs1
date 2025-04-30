package orderworkers

import (
	"os"

	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-gs1/internal/common"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"

	orderworkflows "github.com/cloudfresco/sc-gs1/internal/workflows/orderworkflows"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.WfHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, orderworkflows.ApplicationName, workerOptions)
}

func StartOrderWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	orderconn, err := grpc.NewClient(grpcServerOpt.GrpcOrderServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	orderServiceClient := orderproto.NewOrderServiceClient(orderconn)
	orderActivities := &orderworkflows.OrderActivities{OrderServiceClient: orderServiceClient}

	orderResponseServiceClient := orderproto.NewOrderResponseServiceClient(orderconn)
	orderResponseActivities := &orderworkflows.OrderResponseActivities{OrderResponseServiceClient: orderResponseServiceClient}

	h.RegisterWorkflow(orderworkflows.CreateOrderWorkflow)
	h.RegisterWorkflow(orderworkflows.UpdateOrderWorkflow)
	h.RegisterWorkflow(orderworkflows.CreateOrderResponseWorkflow)
	h.RegisterWorkflow(orderworkflows.UpdateOrderResponseWorkflow)
	h.RegisterActivity(orderActivities)
	h.RegisterActivity(orderResponseActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
