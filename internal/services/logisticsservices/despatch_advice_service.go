package logisticsservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	logisticsstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/logistics/v1"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// DespatchAdviceService - For accessing DespatchAdvice services
type DespatchAdviceService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	logisticsproto.UnimplementedDespatchAdviceServiceServer
}

// NewDespatchAdviceService - Create DespatchAdvice service
func NewDespatchAdviceService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *DespatchAdviceService {
	return &DespatchAdviceService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// StartLogisticsServer - Start Logistics server
func StartLogisticsServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, uc)
	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcLogisticsServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	logisticsproto.RegisterDespatchAdviceServiceServer(srv, despatchAdviceService)
	logisticsproto.RegisterReceivingAdviceServiceServer(srv, receivingAdviceService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

const insertDespatchAdviceSQL = `insert into despatch_advices
	  (
	  uuid4,
delivery_type_code,
rack_id_at_pick_up_location,
total_deposit_amount,
tda_code_list_version,
tda_currency_code,
total_number_of_lines,
blanket_order,
buyer,
carrier,
contract,
customer_document_reference,
declarants_customs_identity,
delivery_note,
delivery_schedule,
despatch_advice_identification,
freight_forwarder,
inventory_location,
invoice,
invoicee,
logistic_service_provider,
order_response,
pick_up_location,
product_certification,
promotional_deal,
purchase_conditions,
purchase_order,
receiver,
returns_instruction,
seller,
ship_from,
shipper,
ship_to,
specification,
transport_instruction,
ultimate_consignee,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
  :uuid4,
:delivery_type_code,
:rack_id_at_pick_up_location,
:total_deposit_amount,
:tda_code_list_version,
:tda_currency_code,
:total_number_of_lines,
:blanket_order,
:buyer,
:carrier,
:contract,
:customer_document_reference,
:declarants_customs_identity,
:delivery_note,
:delivery_schedule,
:despatch_advice_identification,
:freight_forwarder,
:inventory_location,
:invoice,
:invoicee,
:logistic_service_provider,
:order_response,
:pick_up_location,
:product_certification,
:promotional_deal,
:purchase_conditions,
:purchase_order,
:receiver,
:returns_instruction,
:seller,
:ship_from,
:shipper,
:ship_to,
:specification,
:transport_instruction,
:ultimate_consignee,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectDespatchAdvicesSQL = `select
    id,
    uuid4,
delivery_type_code,
rack_id_at_pick_up_location,
total_deposit_amount,
tda_code_list_version,
tda_currency_code,
total_number_of_lines,
blanket_order,
buyer,
carrier,
contract,
customer_document_reference,
declarants_customs_identity,
delivery_note,
delivery_schedule,
despatch_advice_identification,
freight_forwarder,
inventory_location,
invoice,
invoicee,
logistic_service_provider,
order_response,
pick_up_location,
product_certification,
promotional_deal,
purchase_conditions,
purchase_order,
receiver,
returns_instruction,
seller,
ship_from,
shipper,
ship_to,
specification,
transport_instruction,
ultimate_consignee,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from despatch_advices`

// updateDespatchAdviceSQL - update DespatchAdviceSQL query
const updateDespatchAdviceSQL = `update despatch_advices set 
  delivery_type_code= ?,
  rack_id_at_pick_up_location= ?,
  total_deposit_amount= ?,
  tda_code_list_version= ?,
  tda_currency_code= ?,
  total_number_of_lines= ?,
  updated_at = ? where uuid4 = ?;`

// CreateDespatchAdvice - CreateDespatchAdvice
func (das *DespatchAdviceService) CreateDespatchAdvice(ctx context.Context, in *logisticsproto.CreateDespatchAdviceRequest) (*logisticsproto.CreateDespatchAdviceResponse, error) {
	despatchAdvice, err := das.ProcessDespatchAdviceRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdvice(ctx, insertDespatchAdviceSQL, despatchAdvice, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceResponse := logisticsproto.CreateDespatchAdviceResponse{}
	despatchAdviceResponse.DespatchAdvice = despatchAdvice
	return &despatchAdviceResponse, nil
}

// ProcessDespatchAdviceRequest - ProcessDespatchAdviceRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceRequest) (*logisticsproto.DespatchAdvice, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, das.UserServiceClient)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	despatchAdviceD := logisticsproto.DespatchAdviceD{}
	despatchAdviceD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceD.DeliveryTypeCode = in.DeliveryTypeCode
	despatchAdviceD.RackIdAtPickUpLocation = in.RackIdAtPickUpLocation
	despatchAdviceD.TotalDepositAmount = in.TotalDepositAmount
	despatchAdviceD.TdaCodeListVersion = in.TdaCodeListVersion
	despatchAdviceD.TdaCurrencyCode = in.TdaCurrencyCode
	despatchAdviceD.TotalNumberOfLines = in.TotalNumberOfLines
	despatchAdviceD.BlanketOrder = in.BlanketOrder
	despatchAdviceD.Buyer = in.Buyer
	despatchAdviceD.Carrier = in.Carrier
	despatchAdviceD.Contract = in.Contract
	despatchAdviceD.CustomerDocumentReference = in.CustomerDocumentReference
	despatchAdviceD.DeclarantsCustomsIdentity = in.DeclarantsCustomsIdentity
	despatchAdviceD.DeliveryNote = in.DeliveryNote
	despatchAdviceD.DeliverySchedule = in.DeliverySchedule
	despatchAdviceD.DespatchAdviceIdentification = in.DespatchAdviceIdentification
	despatchAdviceD.FreightForwarder = in.FreightForwarder
	despatchAdviceD.InventoryLocation = in.InventoryLocation
	despatchAdviceD.Invoice = in.Invoice
	despatchAdviceD.Invoicee = in.Invoicee
	despatchAdviceD.LogisticServiceProvider = in.LogisticServiceProvider
	despatchAdviceD.OrderResponse = in.OrderResponse
	despatchAdviceD.PickUpLocation = in.PickUpLocation
	despatchAdviceD.ProductCertification = in.ProductCertification
	despatchAdviceD.PromotionalDeal = in.PromotionalDeal
	despatchAdviceD.PurchaseConditions = in.PurchaseConditions
	despatchAdviceD.PurchaseOrder = in.PurchaseOrder
	despatchAdviceD.Receiver = in.Receiver
	despatchAdviceD.ReturnsInstruction = in.ReturnsInstruction
	despatchAdviceD.Seller = in.Seller
	despatchAdviceD.ShipFrom = in.ShipFrom
	despatchAdviceD.Shipper = in.Shipper
	despatchAdviceD.ShipTo = in.ShipTo
	despatchAdviceD.Specification = in.Specification
	despatchAdviceD.TransportInstruction = in.TransportInstruction
	despatchAdviceD.UltimateConsignee = in.UltimateConsignee

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	despatchAdvice := logisticsproto.DespatchAdvice{DespatchAdviceD: &despatchAdviceD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &despatchAdvice, nil
}

// insertDespatchAdvice - Insert DespatchAdvice into database
func (das *DespatchAdviceService) insertDespatchAdvice(ctx context.Context, insertDespatchAdviceSQL string, despatchAdvice *logisticsproto.DespatchAdvice, userEmail string, requestID string) error {
	despatchAdviceTmp, err := das.crDespatchAdviceStruct(ctx, despatchAdvice, userEmail, requestID)
	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertDespatchAdviceSQL, despatchAdviceTmp)
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		despatchAdvice.DespatchAdviceD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(despatchAdvice.DespatchAdviceD.Uuid4)
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		despatchAdvice.DespatchAdviceD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crDespatchAdviceStruct - process DespatchAdvice details
func (das *DespatchAdviceService) crDespatchAdviceStruct(ctx context.Context, despatchAdvice *logisticsproto.DespatchAdvice, userEmail string, requestID string) (*logisticsstruct.DespatchAdvice, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(despatchAdvice.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(despatchAdvice.CrUpdTime.UpdatedAt)

	despatchAdviceTmp := logisticsstruct.DespatchAdvice{DespatchAdviceD: despatchAdvice.DespatchAdviceD, CrUpdUser: despatchAdvice.CrUpdUser, CrUpdTime: crUpdTime}

	return &despatchAdviceTmp, nil
}

// GetDespatchAdvices - Get DespatchAdvices
func (das *DespatchAdviceService) GetDespatchAdvices(ctx context.Context, in *logisticsproto.GetDespatchAdvicesRequest) (*logisticsproto.GetDespatchAdvicesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = das.DBService.LimitSQLRows
	}

	query := "(status_code = ?)"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	despatchAdvices := []*logisticsproto.DespatchAdvice{}

	nselectDespatchAdvicesSQL := selectDespatchAdvicesSQL + ` where ` + query

	rows, err := das.DBService.DB.QueryxContext(ctx, nselectDespatchAdvicesSQL, common.Active)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		despatchAdviceTmp := logisticsstruct.DespatchAdvice{}
		err = rows.StructScan(&despatchAdviceTmp)
		if err != nil {
			das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		despatchAdvice, err := das.getDespatchAdviceStruct(ctx, &getRequest, despatchAdviceTmp)
		if err != nil {
			das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		despatchAdvices = append(despatchAdvices, despatchAdvice)

	}

	despatchAdvicesResponse := logisticsproto.GetDespatchAdvicesResponse{}
	if len(despatchAdvices) != 0 {
		next := despatchAdvices[len(despatchAdvices)-1].DespatchAdviceD.Id
		next--
		nextc := common.EncodeCursor(next)
		despatchAdvicesResponse = logisticsproto.GetDespatchAdvicesResponse{DespatchAdvices: despatchAdvices, NextCursor: nextc}
	} else {
		despatchAdvicesResponse = logisticsproto.GetDespatchAdvicesResponse{DespatchAdvices: despatchAdvices, NextCursor: "0"}
	}
	return &despatchAdvicesResponse, nil
}

// GetDespatchAdvice - Get DespatchAdvice
func (das *DespatchAdviceService) GetDespatchAdvice(ctx context.Context, inReq *logisticsproto.GetDespatchAdviceRequest) (*logisticsproto.GetDespatchAdviceResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectDespatchAdvicesSQL := selectDespatchAdvicesSQL + ` where uuid4 = ?;`
	row := das.DBService.DB.QueryRowxContext(ctx, nselectDespatchAdvicesSQL, uuid4byte)
	despatchAdviceTmp := logisticsstruct.DespatchAdvice{}
	err = row.StructScan(&despatchAdviceTmp)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdvice, err := das.getDespatchAdviceStruct(ctx, in, despatchAdviceTmp)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	despatchAdviceResponse := logisticsproto.GetDespatchAdviceResponse{}
	despatchAdviceResponse.DespatchAdvice = despatchAdvice
	return &despatchAdviceResponse, nil
}

// GetDespatchAdviceByPk - Get DespatchAdvice By Primary key(Id)
func (das *DespatchAdviceService) GetDespatchAdviceByPk(ctx context.Context, inReq *logisticsproto.GetDespatchAdviceByPkRequest) (*logisticsproto.GetDespatchAdviceByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectDespatchAdvicesSQL := selectDespatchAdvicesSQL + ` where id = ?;`
	row := das.DBService.DB.QueryRowxContext(ctx, nselectDespatchAdvicesSQL, in.Id)
	despatchAdviceTmp := logisticsstruct.DespatchAdvice{}
	err := row.StructScan(&despatchAdviceTmp)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	despatchAdvice, err := das.getDespatchAdviceStruct(ctx, &getRequest, despatchAdviceTmp)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	despatchAdviceResponse := logisticsproto.GetDespatchAdviceByPkResponse{}
	despatchAdviceResponse.DespatchAdvice = despatchAdvice
	return &despatchAdviceResponse, nil
}

// getDespatchAdviceStruct - Get despatchAdvice
func (das *DespatchAdviceService) getDespatchAdviceStruct(ctx context.Context, in *commonproto.GetRequest, despatchAdviceTmp logisticsstruct.DespatchAdvice) (*logisticsproto.DespatchAdvice, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(despatchAdviceTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(despatchAdviceTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(despatchAdviceTmp.DespatchAdviceD.Uuid4)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceTmp.DespatchAdviceD.IdS = uuid4Str

	despatchAdvice := logisticsproto.DespatchAdvice{DespatchAdviceD: despatchAdviceTmp.DespatchAdviceD, CrUpdUser: despatchAdviceTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &despatchAdvice, nil
}

// UpdateDespatchAdvice - Update DespatchAdvice
func (das *DespatchAdviceService) UpdateDespatchAdvice(ctx context.Context, in *logisticsproto.UpdateDespatchAdviceRequest) (*logisticsproto.UpdateDespatchAdviceResponse, error) {
	db := das.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateDespatchAdviceSQL)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.DeliveryTypeCode,
			in.RackIdAtPickUpLocation,
			in.TotalDepositAmount,
			in.TdaCodeListVersion,
			in.TdaCurrencyCode,
			in.TotalNumberOfLines,
			tn,
			uuid4byte)
		if err != nil {
			das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &logisticsproto.UpdateDespatchAdviceResponse{}, nil
}
