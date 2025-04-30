package orderservices

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	orderstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/order/v1"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// OrderService - For accessing Order services
type OrderService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	orderproto.UnimplementedOrderServiceServer
}

// NewOrderService - Create Order service
func NewOrderService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *OrderService {
	return &OrderService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// StartOrderServer - Start Order server
func StartOrderServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	orderService := NewOrderService(log, dbService, redisService, uc)
	orderResponseService := NewOrderResponseService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcOrderServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)

	orderproto.RegisterOrderServiceServer(srv, orderService)
	orderproto.RegisterOrderResponseServiceServer(srv, orderResponseService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

const insertOrderSQL = `insert into orders
	    (uuid4,
is_application_receipt_acknowledgement_required,
is_order_free_of_excise_tax_duty,
order_change_reason_code,
order_entry_type,
order_instruction_code,
order_priority,
order_type_code,
total_monetary_amount_excluding_taxes,
tmaet_code_list_version,
tmaet_currency_code,
total_monetary_amount_including_taxes,
tmait_code_list_version,
tmait_currency_code,
total_tax_amount,
tta_code_list_version,
tta_currency_code,
bill_to,
buyer,
contract,
customer_document_reference,
customs_broker,
order_identification,
pickup_from,
promotional_deal,
quote_number,
seller,
trade_agreement,
delivery_date_according_to_schedule,
latest_delivery_date,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(:uuid4,
:is_application_receipt_acknowledgement_required,
:is_order_free_of_excise_tax_duty,
:order_change_reason_code,
:order_entry_type,
:order_instruction_code,
:order_priority,
:order_type_code,
:total_monetary_amount_excluding_taxes,
:tmaet_code_list_version,
:tmaet_currency_code,
:total_monetary_amount_including_taxes,
:tmait_code_list_version,
:tmait_currency_code,
:total_tax_amount,
:tta_code_list_version,
:tta_currency_code,
:bill_to,
:buyer,
:contract,
:customer_document_reference,
:customs_broker,
:order_identification,
:pickup_from,
:promotional_deal,
:quote_number,
:seller,
:trade_agreement,
:delivery_date_according_to_schedule,
:latest_delivery_date,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectOrdersSQL = `select
    id,
    uuid4,
is_application_receipt_acknowledgement_required,
is_order_free_of_excise_tax_duty,
order_change_reason_code,
order_entry_type,
order_instruction_code,
order_priority,
order_type_code,
total_monetary_amount_excluding_taxes,
tmaet_code_list_version,
tmaet_currency_code,
total_monetary_amount_including_taxes,
tmait_code_list_version,
tmait_currency_code,
total_tax_amount,
tta_code_list_version,
tta_currency_code,
bill_to,
buyer,
contract,
customer_document_reference,
customs_broker,
order_identification,
pickup_from,
promotional_deal,
quote_number,
seller,
trade_agreement,
delivery_date_according_to_schedule,
latest_delivery_date,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from orders`

// updateOrderSQL - update OrderSQL query
const updateOrderSQL = `update orders set 
  order_change_reason_code= ?,
  order_entry_type= ?,
  order_instruction_code= ?,
  order_priority= ?,
  order_type_code= ?,
  total_monetary_amount_excluding_taxes= ?,
  tmaet_code_list_version= ?,
  tmaet_currency_code= ?,
  total_monetary_amount_including_taxes= ?,
  tmait_code_list_version= ?,
  tmait_currency_code= ?,
  total_tax_amount= ?,
  tta_code_list_version= ?,
  tta_currency_code= ?,
  updated_at = ? where uuid4 = ?;`

// CreateOrder - Create Order
func (o *OrderService) CreateOrder(ctx context.Context, in *orderproto.CreateOrderRequest) (*orderproto.CreateOrderResponse, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, o.UserServiceClient)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	deliveryDateAccordingToSchedule, err := time.Parse(common.Layout, in.DeliveryDateAccordingToSchedule)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	latestDeliveryDate, err := time.Parse(common.Layout, in.LatestDeliveryDate)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderD := orderproto.OrderD{}
	orderD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderD.IsApplicationReceiptAcknowledgementRequired = in.IsApplicationReceiptAcknowledgementRequired
	orderD.IsOrderFreeOfExciseTaxDuty = in.IsOrderFreeOfExciseTaxDuty
	orderD.OrderChangeReasonCode = in.OrderChangeReasonCode
	orderD.OrderEntryType = in.OrderEntryType
	orderD.OrderInstructionCode = in.OrderInstructionCode
	orderD.OrderPriority = in.OrderPriority
	orderD.OrderTypeCode = in.OrderTypeCode
	orderD.TotalMonetaryAmountExcludingTaxes = in.TotalMonetaryAmountExcludingTaxes
	orderD.TmaetCodeListVersion = in.TmaetCodeListVersion
	orderD.TmaetCurrencyCode = in.TmaetCurrencyCode
	orderD.TotalMonetaryAmountIncludingTaxes = in.TotalMonetaryAmountIncludingTaxes
	orderD.TmaitCodeListVersion = in.TmaitCodeListVersion
	orderD.TmaitCurrencyCode = in.TmaitCurrencyCode
	orderD.TotalTaxAmount = in.TotalTaxAmount
	orderD.TtaCodeListVersion = in.TtaCodeListVersion
	orderD.TtaCurrencyCode = in.TtaCurrencyCode
	orderD.BillTo = in.BillTo
	orderD.Buyer = in.Buyer
	orderD.Contract = in.Contract
	orderD.CustomerDocumentReference = in.CustomerDocumentReference
	orderD.CustomsBroker = in.CustomsBroker
	orderD.OrderIdentification = in.OrderIdentification
	orderD.PickupFrom = in.PickupFrom
	orderD.PromotionalDeal = in.PromotionalDeal
	orderD.QuoteNumber = in.QuoteNumber
	orderD.Seller = in.Seller
	orderD.TradeAgreement = in.TradeAgreement

	orderT := orderproto.OrderT{}
	orderT.DeliveryDateAccordingToSchedule = common.TimeToTimestamp(deliveryDateAccordingToSchedule.UTC().Truncate(time.Second))
	orderT.LatestDeliveryDate = common.TimeToTimestamp(latestDeliveryDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	order := orderproto.Order{OrderD: &orderD, OrderT: &orderT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	orderLineItems := []*orderproto.OrderLineItem{}
	// we will do for loop on lines which is comes from client form
	for _, line := range in.OrderLineItems {
		line.UserId = in.UserId
		line.UserEmail = in.UserEmail
		line.RequestId = in.RequestId
		// we wl call CreateOrderLineItem function which wl populate form values to orderline struct
		orderLineItem, err := o.ProcessOrderLineItemRequest(ctx, line)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		orderLineItems = append(orderLineItems, orderLineItem)
	}

	err = o.insertOrder(ctx, insertOrderSQL, &order, insertOrderLineItemSQL, orderLineItems, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponse := orderproto.CreateOrderResponse{}
	orderResponse.Order = &order
	return &orderResponse, nil
}

func (o *OrderService) insertOrder(ctx context.Context, insertOrderSQL string, order *orderproto.Order, insertOrderLineItemSQL string, orderLineItems []*orderproto.OrderLineItem, userEmail string, requestID string) error {
	orderTmp, err := o.crOrderStruct(ctx, order, userEmail, requestID)
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		// header creation
		res, err := tx.NamedExecContext(ctx, insertOrderSQL, orderTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		order.OrderD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(order.OrderD.Uuid4)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		order.OrderD.IdS = uuid4Str

		for _, orderLineItem := range orderLineItems {
			orderLineItem.OrderLineItemD.OrderId = order.OrderD.Id
			orderLineItemTmp, err := o.crOrderLineItemStruct(ctx, orderLineItem, userEmail, requestID)
			if err != nil {
				o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertOrderLineItemSQL, orderLineItemTmp)
			if err != nil {
				o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}
		return nil
	})
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crOrderStruct - process Order details
func (o *OrderService) crOrderStruct(ctx context.Context, order *orderproto.Order, userEmail string, requestID string) (*orderstruct.Order, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(order.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(order.CrUpdTime.UpdatedAt)

	orderT := new(orderstruct.OrderT)
	orderT.DeliveryDateAccordingToSchedule = common.TimestampToTime(order.OrderT.DeliveryDateAccordingToSchedule)
	orderT.LatestDeliveryDate = common.TimestampToTime(order.OrderT.LatestDeliveryDate)

	orderTmp := orderstruct.Order{OrderD: order.OrderD, OrderT: orderT, CrUpdUser: order.CrUpdUser, CrUpdTime: crUpdTime}

	return &orderTmp, nil
}

// GetOrders - Get Orders
func (o *OrderService) GetOrders(ctx context.Context, in *orderproto.GetOrdersRequest) (*orderproto.GetOrdersResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = o.DBService.LimitSQLRows
	}

	query := "(status_code = ?)"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	orders := []*orderproto.Order{}

	nselectOrdersSQL := selectOrdersSQL + ` where ` + query

	rows, err := o.DBService.DB.QueryxContext(ctx, nselectOrdersSQL, common.Active)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		orderTmp := orderstruct.Order{}
		err = rows.StructScan(&orderTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		order, err := o.getOrderStruct(ctx, &getRequest, orderTmp)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		orders = append(orders, order)

	}

	ordersResponse := orderproto.GetOrdersResponse{}
	if len(orders) != 0 {
		next := orders[len(orders)-1].OrderD.Id
		next--
		nextc := common.EncodeCursor(next)
		ordersResponse = orderproto.GetOrdersResponse{Orders: orders, NextCursor: nextc}
	} else {
		ordersResponse = orderproto.GetOrdersResponse{Orders: orders, NextCursor: "0"}
	}
	return &ordersResponse, nil
}

// GetOrder - Get Order
func (o *OrderService) GetOrder(ctx context.Context, inReq *orderproto.GetOrderRequest) (*orderproto.GetOrderResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectOrdersSQL := selectOrdersSQL + ` where uuid4 = ?;`
	row := o.DBService.DB.QueryRowxContext(ctx, nselectOrdersSQL, uuid4byte)
	orderTmp := orderstruct.Order{}
	err = row.StructScan(&orderTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	order, err := o.getOrderStruct(ctx, in, orderTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	orderResponse := orderproto.GetOrderResponse{}
	orderResponse.Order = order
	return &orderResponse, nil
}

// GetOrderByPk - Get Order By Primary key(Id)
func (o *OrderService) GetOrderByPk(ctx context.Context, inReq *orderproto.GetOrderByPkRequest) (*orderproto.GetOrderByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectOrdersSQL := selectOrdersSQL + ` where id = ?;`
	row := o.DBService.DB.QueryRowxContext(ctx, nselectOrdersSQL, in.Id)
	orderTmp := orderstruct.Order{}
	err := row.StructScan(&orderTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	order, err := o.getOrderStruct(ctx, &getRequest, orderTmp)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	orderResponse := orderproto.GetOrderByPkResponse{}
	orderResponse.Order = order
	return &orderResponse, nil
}

// getOrderStruct - Get order
func (o *OrderService) getOrderStruct(ctx context.Context, in *commonproto.GetRequest, orderTmp orderstruct.Order) (*orderproto.Order, error) {
	orderT := new(orderproto.OrderT)
	orderT.DeliveryDateAccordingToSchedule = common.TimeToTimestamp(orderTmp.OrderT.DeliveryDateAccordingToSchedule)
	orderT.LatestDeliveryDate = common.TimeToTimestamp(orderTmp.OrderT.LatestDeliveryDate)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(orderTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(orderTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(orderTmp.OrderD.Uuid4)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderTmp.OrderD.IdS = uuid4Str

	order := orderproto.Order{OrderD: orderTmp.OrderD, OrderT: orderT, CrUpdUser: orderTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &order, nil
}

// UpdateOrder - Update Order
func (o *OrderService) UpdateOrder(ctx context.Context, in *orderproto.UpdateOrderRequest) (*orderproto.UpdateOrderResponse, error) {
	db := o.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateOrderSQL)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.OrderChangeReasonCode,
			in.OrderEntryType,
			in.OrderInstructionCode,
			in.OrderPriority,
			in.OrderTypeCode,
			in.TotalMonetaryAmountExcludingTaxes,
			in.TmaetCodeListVersion,
			in.TmaetCurrencyCode,
			in.TotalMonetaryAmountIncludingTaxes,
			in.TmaitCodeListVersion,
			in.TmaitCurrencyCode,
			in.TotalTaxAmount,
			in.TtaCodeListVersion,
			in.TtaCurrencyCode,
			tn,
			uuid4byte)
		if err != nil {
			o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &orderproto.UpdateOrderResponse{}, nil
}
