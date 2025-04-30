package invoiceworkers

import (
	"os"

	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-gs1/internal/common"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"

	invoiceworkflows "github.com/cloudfresco/sc-gs1/internal/workflows/invoiceworkflows"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.WfHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, invoiceworkflows.ApplicationName, workerOptions)
}

func StartInvoiceWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	invoiceconn, err := grpc.NewClient(grpcServerOpt.GrpcInvoiceServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	invoiceServiceClient := invoiceproto.NewInvoiceServiceClient(invoiceconn)
	invoiceActivities := &invoiceworkflows.InvoiceActivities{InvoiceServiceClient: invoiceServiceClient}

	debitCreditAdviceServiceClient := invoiceproto.NewDebitCreditAdviceServiceClient(invoiceconn)
	debitCreditAdviceActivities := &invoiceworkflows.DebitCreditAdviceActivities{DebitCreditAdviceServiceClient: debitCreditAdviceServiceClient}

	h.RegisterWorkflow(invoiceworkflows.CreateInvoiceWorkflow)
	h.RegisterWorkflow(invoiceworkflows.UpdateInvoiceWorkflow)
	h.RegisterWorkflow(invoiceworkflows.CreateDebitCreditAdviceWorkflow)
	h.RegisterWorkflow(invoiceworkflows.UpdateDebitCreditAdviceWorkflow)
	h.RegisterActivity(invoiceActivities)
	h.RegisterActivity(debitCreditAdviceActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
