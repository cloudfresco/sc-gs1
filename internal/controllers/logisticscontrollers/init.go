package logisticscontrollers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
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

// Init the order controllers
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

// InitTest the logistics controllers
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

	logisticsconn, err := grpc.NewClient(grpcServerOpt.GrpcLogisticsServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	ra := logisticsproto.NewReceivingAdviceServiceClient(logisticsconn)
	da := logisticsproto.NewDespatchAdviceServiceClient(logisticsconn)

	initReceivingAdvice(mux, serverOpt, log, u, ra, h, workflowClient)
	initDespatchAdvice(mux, serverOpt, log, u, da, h, workflowClient)

	return nil
}

func initReceivingAdvice(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, ra logisticsproto.ReceivingAdviceServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	rac := NewReceivingAdviceController(log, u, ra, h, workflowClient, serverOpt)

	mux.Handle("GET /v2.3/receiving-advices", http.HandlerFunc(rac.Index))
	mux.Handle("GET /v2.3/receiving-advices/{id}", http.HandlerFunc(rac.Show))
	mux.Handle("GET /v2.3/receiving-advices/{id}/lines", http.HandlerFunc(rac.GetReceivingAdviceLineItems))
	mux.Handle("POST /v2.3/receiving-advices", http.HandlerFunc(rac.CreateReceivingAdvice))
	mux.Handle("PUT /v2.3/receiving-advices/{id}", http.HandlerFunc(rac.UpdateReceivingAdvice))
}

func initDespatchAdvice(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, da logisticsproto.DespatchAdviceServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	dac := NewDespatchAdviceController(log, u, da, h, workflowClient, serverOpt)

	mux.Handle("GET /v2.3/despatch-advices", http.HandlerFunc(dac.Index))
	mux.Handle("GET /v2.3/despatch-advices/{id}", http.HandlerFunc(dac.Show))
	mux.Handle("GET /v2.3/despatch-advices/{id}/lines", http.HandlerFunc(dac.GetDespatchAdviceLineItems))
	mux.Handle("POST /v2.3/despatch-advices", http.HandlerFunc(dac.CreateDespatchAdvice))
	mux.Handle("PUT /v2.3/despatch-advices/{id}", http.HandlerFunc(dac.UpdateDespatchAdvice))
}
