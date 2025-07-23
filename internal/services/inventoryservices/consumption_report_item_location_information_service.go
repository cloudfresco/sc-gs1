package inventoryservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// InventoryService - For accessing Inventory services
type InventoryService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	CurrencyService   *common.CurrencyService
	inventoryproto.UnimplementedInventoryServiceServer
}

// NewInventoryService - Create Inventory service
func NewInventoryService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient, currency *common.CurrencyService) *InventoryService {
	return &InventoryService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
		CurrencyService:   currency,
	}
}

// StartInventoryServer - Start Inventory server
func StartInventoryServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	inventoryService := NewInventoryService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcInventoryServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	inventoryproto.RegisterInventoryServiceServer(srv, inventoryService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

const insertConsumptionReportItemLocationInformationSQL = `insert into consumption_report_item_loation_informations
	    (total_consumption_amount,
      tca_code_list_version,
      tcac_currency_code,
      inventory_location,
      ship_to, 
      consumption_report_id)
  values(
  :total_consumption_amount,
  :tca_code_list_version,
  :tcac_currency_code,
  :inventory_location,
  :ship_to, 
  :consumption_report_id
);`

/*const selectConsumptionReportItemLocationInformationsSQL = `select
  id,
  total_consumption_amount,
  tca_code_list_version,
  tcac_currency_code,
  inventory_location,
  ship_to,
  consumption_report_id from consumption_report_item_loation_informations*/

// CreateConsumptionReportItemLocationInformation - Create ConsumptionReportItemLocationInformation
func (invs *InventoryService) CreateConsumptionReportItemLocationInformation(ctx context.Context, in *inventoryproto.CreateConsumptionReportItemLocationInformationRequest) (*inventoryproto.CreateConsumptionReportItemLocationInformationResponse, error) {
	consumptionReportItemLocationInformation, err := invs.ProcessConsumptionReportItemLocationInformationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertConsumptionReportItemLocationInformation(ctx, insertConsumptionReportItemLocationInformationSQL, consumptionReportItemLocationInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionReportItemLocationInformationResponse := inventoryproto.CreateConsumptionReportItemLocationInformationResponse{}
	consumptionReportItemLocationInformationResponse.ConsumptionReportItemLocationInformation = consumptionReportItemLocationInformation
	return &consumptionReportItemLocationInformationResponse, nil
}

// ProcessConsumptionReportItemLocationInformationRequest - ProcessConsumptionReportItemLocationInformationRequest
func (invs *InventoryService) ProcessConsumptionReportItemLocationInformationRequest(ctx context.Context, in *inventoryproto.CreateConsumptionReportItemLocationInformationRequest) (*inventoryproto.ConsumptionReportItemLocationInformation, error) {
	consumptionReportItemLocationInformation := inventoryproto.ConsumptionReportItemLocationInformation{}
	consumptionReportItemLocationInformation.TotalConsumptionAmount = in.TotalConsumptionAmount
	consumptionReportItemLocationInformation.TCACodeListVersion = in.TCACodeListVersion
	consumptionReportItemLocationInformation.TCACCurrencyCode = in.TCACCurrencyCode
	consumptionReportItemLocationInformation.InventoryLocation = in.InventoryLocation
	consumptionReportItemLocationInformation.ShipTo = in.ShipTo
	consumptionReportItemLocationInformation.ConsumptionReportId = in.ConsumptionReportId

	return &consumptionReportItemLocationInformation, nil
}

// insertConsumptionReportItemLocationInformation - Insert ConsumptionReportItemLocationInformation details into database
func (invs *InventoryService) insertConsumptionReportItemLocationInformation(ctx context.Context, insertConsumptionReportItemLocationInformationSQL string, consumptionReportItemLocationInformation *inventoryproto.ConsumptionReportItemLocationInformation, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertConsumptionReportItemLocationInformationSQL, consumptionReportItemLocationInformation)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consumptionReportItemLocationInformation.Id = uint32(uID)
		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
