package ordercontrollers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
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

// InitTest the order controllers
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

	orderconn, err := grpc.NewClient(grpcServerOpt.GrpcOrderServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	o := orderproto.NewOrderServiceClient(orderconn)
	or := orderproto.NewOrderResponseServiceClient(orderconn)

	initOrder(mux, serverOpt, log, u, o, h, workflowClient)
	initOrderResponse(mux, serverOpt, log, u, or, h, workflowClient)

	return nil
}

func initOrder(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, o orderproto.OrderServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	oc := NewOrderController(log, u, o, h, workflowClient, serverOpt)

	mux.Handle("GET /v2.3/orders", http.HandlerFunc(oc.Index))
	mux.Handle("GET /v2.3/orders/{id}", http.HandlerFunc(oc.Show))
	mux.Handle("GET /v2.3/orders/{id}/lines", http.HandlerFunc(oc.GetOrderLineItems))
	mux.Handle("POST /v2.3/orders", http.HandlerFunc(oc.CreateOrder))
	mux.Handle("PUT /v2.3/orders/{id}", http.HandlerFunc(oc.UpdateOrder))
}

func initOrderResponse(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, or orderproto.OrderResponseServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	orc := NewOrderResponseController(log, u, or, h, workflowClient, serverOpt)

	mux.Handle("GET /v2.3/order-responses", http.HandlerFunc(orc.Index))
	mux.Handle("GET /v2.3/order-responses/{id}", http.HandlerFunc(orc.Show))
	mux.Handle("GET /v2.3/order-responses/{id}/lines", http.HandlerFunc(orc.GetOrderResponseLineItems))
	mux.Handle("POST /v2.3/order-responses", http.HandlerFunc(orc.CreateOrderResponse))
	mux.Handle("PUT /v2.3/order-responses/{id}", http.HandlerFunc(orc.UpdateOrderResponse))
}
