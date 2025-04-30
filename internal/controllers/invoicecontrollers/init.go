package invoicecontrollers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/throttled/throttled/v2/store/goredisstore"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	h              common.WfHelper
	workflowClient client.Client
)

// Init the invoice controllers
func Init(log *zap.Logger, mux *http.ServeMux, store *goredisstore.GoRedisStore, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, uptraceOpt *config.UptraceOptions, configFilePath string) error {
	pwd, _ := os.Getwd()
	keyPath := pwd + filepath.FromSlash(grpcServerOpt.GrpcCaCertPath)

	err := initSetup(log, mux, keyPath, configFilePath, serverOpt, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
		return err
	}

	return nil
}

// InitTest the invoice controllers
func InitTest(log *zap.Logger, mux *http.ServeMux, store *goredisstore.GoRedisStore, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, uptraceOpt *config.UptraceOptions, configFilePath string) error {
	pwd, _ := os.Getwd()
	keyPath := filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCaCertPath))

	err := initSetup(log, mux, keyPath, configFilePath, serverOpt, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
		return err
	}

	return nil
}

func initSetup(log *zap.Logger, mux *http.ServeMux, keyPath string, configFilePath string, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions) error {
	creds, err := credentials.NewClientTLSFromFile(keyPath, "localhost")
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
	}

	tp, err := config.InitTracerProvider()
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 9108), zap.Error(err))
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Error("Error", zap.Int("msgnum", 9108), zap.Error(err))
		}
	}()

	h.SetupServiceConfig(configFilePath)
	workflowClient, err = h.Builder.BuildCadenceClient()
	if err != nil {
		panic(err)
	}

	userconn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 113), zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)

	invoiceconn, err := grpc.NewClient(grpcServerOpt.GrpcInvoiceServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	i := invoiceproto.NewInvoiceServiceClient(invoiceconn)
	dc := invoiceproto.NewDebitCreditAdviceServiceClient(invoiceconn)

	initInvoice(mux, serverOpt, log, u, i, h, workflowClient)
	initDebitCreditAdvice(mux, serverOpt, log, u, dc, h, workflowClient)

	return nil
}

func initInvoice(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, i invoiceproto.InvoiceServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	ic := NewInvoiceController(log, u, i, h, workflowClient, serverOpt)

	mux.Handle("GET /v2.3/invoices", http.HandlerFunc(ic.Index))
	mux.Handle("GET /v2.3/invoices/{id}", http.HandlerFunc(ic.Show))
	mux.Handle("GET /v2.3/invoices/{id}/lines", http.HandlerFunc(ic.GetInvoiceLineItems))
	mux.Handle("POST /v2.3/invoices", http.HandlerFunc(ic.CreateInvoice))
	mux.Handle("PUT /v2.3/invoices/{id}", http.HandlerFunc(ic.UpdateInvoice))
}

func initDebitCreditAdvice(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, dc invoiceproto.DebitCreditAdviceServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	dcc := NewDebitCreditAdviceController(log, u, dc, h, workflowClient, serverOpt)

	mux.Handle("GET /v2.3/debit-credit-advices", http.HandlerFunc(dcc.Index))
	mux.Handle("GET /v2.3/debit-credit-advices/{id}", http.HandlerFunc(dcc.Show))
	mux.Handle("GET /v2.3/debit-credit-advices/{id}/lines", http.HandlerFunc(dcc.GetDebitCreditAdviceLineItems))
	mux.Handle("POST /v2.3/debit-credit-advices", http.HandlerFunc(dcc.CreateDebitCreditAdvice))
	mux.Handle("PUT /v2.3/debit-credit-advices/{id}", http.HandlerFunc(dcc.UpdateDebitCreditAdvice))
}
