package invoiceservices

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	invoicestruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/invoice/v1"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// InvoiceService - For accessing Invoice services
type InvoiceService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
  CurrencyService   *common.CurrencyService
	invoiceproto.UnimplementedInvoiceServiceServer
}

// NewInvoiceService - Create Invoice service
func NewInvoiceService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient, currency *common.CurrencyService) *InvoiceService {
	return &InvoiceService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
    CurrencyService:   currency,
	}
}

// StartInvoiceServer - Start Invoice server
func StartInvoiceServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	invoiceService := NewInvoiceService(log, dbService, redisService, uc)
	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcInvoiceServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	invoiceproto.RegisterInvoiceServiceServer(srv, invoiceService)
	invoiceproto.RegisterDebitCreditAdviceServiceServer(srv, debitCreditAdviceService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

const insertInvoiceSQL = `insert into invoices
	    (uuid4,
country_of_supply_of_goods,
credit_reason_code,
discount_agreement_terms,
invoice_currency_code,
invoice_type,
is_buyer_based_in_eu,
is_first_seller_based_in_eu,
supplier_account_receivable,
blanket_order,
buyer,
contract,
delivery_note,
despatch_advice,
dispute_notice,
inventory_location,
inventory_report,
invoice,
invoice_identification,
manifest,
order_response,
payee,
payer,
pickup_from,
price_list,
promotional_deal,
purchase_order,
receiving_advice,
remit_to,
returns_notice,
sales_order,
sales_report,
seller,
ship_from,
ship_to,
supplier_agent_representative,
supplier_corporate_office,
tax_currency_information,
tax_representative,
trade_agreement,
ultimate_consignee,
actual_delivery_date,
invoicing_period_begin,
invoicing_period_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(:uuid4,
:country_of_supply_of_goods,
:credit_reason_code,
:discount_agreement_terms,
:invoice_currency_code,
:invoice_type,
:is_buyer_based_in_eu,
:is_first_seller_based_in_eu,
:supplier_account_receivable,
:blanket_order,
:buyer,
:contract,
:delivery_note,
:despatch_advice,
:dispute_notice,
:inventory_location,
:inventory_report,
:invoice,
:invoice_identification,
:manifest,
:order_response,
:payee,
:payer,
:pickup_from,
:price_list,
:promotional_deal,
:purchase_order,
:receiving_advice,
:remit_to,
:returns_notice,
:sales_order,
:sales_report,
:seller,
:ship_from,
:ship_to,
:supplier_agent_representative,
:supplier_corporate_office,
:tax_currency_information,
:tax_representative,
:trade_agreement,
:ultimate_consignee,
:actual_delivery_date,
:invoicing_period_begin,
:invoicing_period_end,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectInvoicesSQL = `select
    id,
uuid4,
country_of_supply_of_goods,
credit_reason_code,
discount_agreement_terms,
invoice_currency_code,
invoice_type,
is_buyer_based_in_eu,
is_first_seller_based_in_eu,
supplier_account_receivable,
blanket_order,
buyer,
contract,
delivery_note,
despatch_advice,
dispute_notice,
inventory_location,
inventory_report,
invoice,
invoice_identification,
manifest,
order_response,
payee,
payer,
pickup_from,
price_list,
promotional_deal,
purchase_order,
receiving_advice,
remit_to,
returns_notice,
sales_order,
sales_report,
seller,
ship_from,
ship_to,
supplier_agent_representative,
supplier_corporate_office,
tax_currency_information,
tax_representative,
trade_agreement,
ultimate_consignee,
actual_delivery_date,
invoicing_period_begin,
invoicing_period_end, 
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from invoices`

// updateInvoiceSQL - update InvoiceSQL query
const updateInvoiceSQL = `update invoices set 
  country_of_supply_of_goods = ?, 
  credit_reason_code = ?, 
  discount_agreement_terms = ?, 
  invoice_currency_code = ?,
  invoice_type = ?,
  supplier_account_receivable = ?,
  updated_at = ? where uuid4 = ?;`

// CreateInvoice - Create Invoice
func (invs *InvoiceService) CreateInvoice(ctx context.Context, in *invoiceproto.CreateInvoiceRequest) (*invoiceproto.CreateInvoiceResponse, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	invoicingPeriodBegin, err := time.Parse(common.Layout, in.InvoicingPeriodBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoicingPeriodEnd, err := time.Parse(common.Layout, in.InvoicingPeriodEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	actualDeliveryDate, err := time.Parse(common.Layout, in.ActualDeliveryDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceD := invoiceproto.InvoiceD{}
	invoiceD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceD.CountryOfSupplyOfGoods = in.CountryOfSupplyOfGoods
	invoiceD.CreditReasonCode = in.CreditReasonCode
	invoiceD.DiscountAgreementTerms = in.DiscountAgreementTerms
	invoiceD.InvoiceCurrencyCode = in.InvoiceCurrencyCode
	invoiceD.InvoiceType = in.InvoiceType
	invoiceD.IsBuyerBasedInEu = in.IsBuyerBasedInEu
	invoiceD.IsFirstSellerBasedInEu = in.IsFirstSellerBasedInEu
	invoiceD.SupplierAccountReceivable = in.SupplierAccountReceivable
	invoiceD.BlanketOrder = in.BlanketOrder
	invoiceD.Buyer = in.Buyer
	invoiceD.Contract = in.Contract
	invoiceD.DeliveryNote = in.DeliveryNote
	invoiceD.DespatchAdvice = in.DespatchAdvice
	invoiceD.DisputeNotice = in.DisputeNotice
	invoiceD.InventoryLocation = in.InventoryLocation
	invoiceD.InventoryReport = in.InventoryReport
	invoiceD.Invoice = in.Invoice
	invoiceD.InvoiceIdentification = in.InvoiceIdentification
	invoiceD.Manifest = in.Manifest
	invoiceD.OrderResponse = in.OrderResponse
	invoiceD.Payee = in.Payee
	invoiceD.Payer = in.Payer
	invoiceD.PickupFrom = in.PickupFrom
	invoiceD.PriceList = in.PriceList
	invoiceD.PromotionalDeal = in.PromotionalDeal
	invoiceD.PurchaseOrder = in.PurchaseOrder
	invoiceD.ReceivingAdvice = in.ReceivingAdvice
	invoiceD.RemitTo = in.RemitTo
	invoiceD.ReturnsNotice = in.ReturnsNotice
	invoiceD.SalesOrder = in.SalesOrder
	invoiceD.SalesReport = in.SalesReport
	invoiceD.Seller = in.Seller
	invoiceD.ShipFrom = in.ShipFrom
	invoiceD.ShipTo = in.ShipTo
	invoiceD.SupplierAgentRepresentative = in.SupplierAgentRepresentative
	invoiceD.SupplierCorporateOffice = in.SupplierCorporateOffice
	invoiceD.TaxCurrencyInformation = in.TaxCurrencyInformation
	invoiceD.TaxRepresentative = in.TaxRepresentative
	invoiceD.TradeAgreement = in.TradeAgreement
	invoiceD.UltimateConsignee = in.UltimateConsignee

	invoiceT := invoiceproto.InvoiceT{}
	invoiceT.ActualDeliveryDate = common.TimeToTimestamp(actualDeliveryDate.UTC().Truncate(time.Second))
	invoiceT.InvoicingPeriodBegin = common.TimeToTimestamp(invoicingPeriodBegin.UTC().Truncate(time.Second))
	invoiceT.InvoicingPeriodEnd = common.TimeToTimestamp(invoicingPeriodEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	invoice := invoiceproto.Invoice{InvoiceD: &invoiceD, InvoiceT: &invoiceT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	invoiceLineItems := []*invoiceproto.InvoiceLineItem{}
	// we will do for loop on lines which is comes from client form
	for _, line := range in.InvoiceLineItems {
		line.UserId = in.UserId
		line.UserEmail = in.UserEmail
		line.RequestId = in.RequestId
		// we wl call CreateInvoiceLineItem function which wl populate form values to invoiceline struct
		invoiceLineItem, err := invs.ProcessInvoiceLineItemRequest(ctx, line)
		if err != nil {
			invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		invoiceLineItems = append(invoiceLineItems, invoiceLineItem)
	}

	err = invs.insertInvoice(ctx, insertInvoiceSQL, &invoice, insertInvoiceLineItemSQL, invoiceLineItems, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceResponse := invoiceproto.CreateInvoiceResponse{}
	invoiceResponse.Invoice = &invoice
	return &invoiceResponse, nil
}

func (invs *InvoiceService) insertInvoice(ctx context.Context, insertInvoiceSQL string, invoice *invoiceproto.Invoice, insertInvoiceLineItemSQL string, invoiceLineItems []*invoiceproto.InvoiceLineItem, userEmail string, requestID string) error {
	invoiceTmp, err := invs.crInvoiceStruct(ctx, invoice, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		// header creation
		res, err := tx.NamedExecContext(ctx, insertInvoiceSQL, invoiceTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoice.InvoiceD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(invoice.InvoiceD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoice.InvoiceD.IdS = uuid4Str

		for _, invoiceLineItem := range invoiceLineItems {
			invoiceLineItem.InvoiceLineItemD.InvoiceId = invoice.InvoiceD.Id
			invoiceLineItemTmp, err := invs.crInvoiceLineItemStruct(ctx, invoiceLineItem, userEmail, requestID)
			if err != nil {
				invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertInvoiceLineItemSQL, invoiceLineItemTmp)
			if err != nil {
				invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			if err != nil {
				invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}
		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crInvoiceStruct - process Invoice details
func (invs *InvoiceService) crInvoiceStruct(ctx context.Context, invoice *invoiceproto.Invoice, userEmail string, requestID string) (*invoicestruct.Invoice, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(invoice.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(invoice.CrUpdTime.UpdatedAt)

	invoiceT := new(invoicestruct.InvoiceT)
	invoiceT.ActualDeliveryDate = common.TimestampToTime(invoice.InvoiceT.ActualDeliveryDate)
	invoiceT.InvoicingPeriodBegin = common.TimestampToTime(invoice.InvoiceT.InvoicingPeriodBegin)
	invoiceT.InvoicingPeriodEnd = common.TimestampToTime(invoice.InvoiceT.InvoicingPeriodEnd)

	invoiceTmp := invoicestruct.Invoice{InvoiceD: invoice.InvoiceD, InvoiceT: invoiceT, CrUpdUser: invoice.CrUpdUser, CrUpdTime: crUpdTime}

	return &invoiceTmp, nil
}

// GetInvoices - Get Invoices
func (invs *InvoiceService) GetInvoices(ctx context.Context, in *invoiceproto.GetInvoicesRequest) (*invoiceproto.GetInvoicesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = invs.DBService.LimitSQLRows
	}

	query := "(status_code = ?)"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	invoices := []*invoiceproto.Invoice{}

	nselectInvoicesSQL := selectInvoicesSQL + ` where ` + query
	rows, err := invs.DBService.DB.QueryxContext(ctx, nselectInvoicesSQL, common.Active)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		invoiceTmp := invoicestruct.Invoice{}
		err = rows.StructScan(&invoiceTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		invoice, err := invs.getInvoiceStruct(ctx, &getRequest, invoiceTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		invoices = append(invoices, invoice)

	}

	invoicesResponse := invoiceproto.GetInvoicesResponse{}
	if len(invoices) != 0 {
		next := invoices[len(invoices)-1].InvoiceD.Id
		next--
		nextc := common.EncodeCursor(next)
		invoicesResponse = invoiceproto.GetInvoicesResponse{Invoices: invoices, NextCursor: nextc}
	} else {
		invoicesResponse = invoiceproto.GetInvoicesResponse{Invoices: invoices, NextCursor: "0"}
	}
	return &invoicesResponse, nil
}

// GetInvoice - Get Invoice
func (invs *InvoiceService) GetInvoice(ctx context.Context, inReq *invoiceproto.GetInvoiceRequest) (*invoiceproto.GetInvoiceResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectInvoicesSQL := selectInvoicesSQL + ` where uuid4 = ?;`
	row := invs.DBService.DB.QueryRowxContext(ctx, nselectInvoicesSQL, uuid4byte)
	invoiceTmp := invoicestruct.Invoice{}
	err = row.StructScan(&invoiceTmp)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoice, err := invs.getInvoiceStruct(ctx, in, invoiceTmp)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	invoiceResponse := invoiceproto.GetInvoiceResponse{}
	invoiceResponse.Invoice = invoice
	return &invoiceResponse, nil
}

// GetInvoiceByPk - Get Invoice By Primary key(Id)
func (invs *InvoiceService) GetInvoiceByPk(ctx context.Context, inReq *invoiceproto.GetInvoiceByPkRequest) (*invoiceproto.GetInvoiceByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectInvoicesSQL := selectInvoicesSQL + ` where id = ?;`
	row := invs.DBService.DB.QueryRowxContext(ctx, nselectInvoicesSQL, in.Id)
	invoiceTmp := invoicestruct.Invoice{}
	err := row.StructScan(&invoiceTmp)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	invoice, err := invs.getInvoiceStruct(ctx, &getRequest, invoiceTmp)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	invoiceResponse := invoiceproto.GetInvoiceByPkResponse{}
	invoiceResponse.Invoice = invoice
	return &invoiceResponse, nil
}

// getInvoiceStruct - Get invoice
func (invs *InvoiceService) getInvoiceStruct(ctx context.Context, in *commonproto.GetRequest, invoiceTmp invoicestruct.Invoice) (*invoiceproto.Invoice, error) {
	invoiceT := new(invoiceproto.InvoiceT)
	invoiceT.ActualDeliveryDate = common.TimeToTimestamp(invoiceTmp.InvoiceT.ActualDeliveryDate)
	invoiceT.InvoicingPeriodBegin = common.TimeToTimestamp(invoiceTmp.InvoiceT.InvoicingPeriodBegin)
	invoiceT.InvoicingPeriodEnd = common.TimeToTimestamp(invoiceTmp.InvoiceT.InvoicingPeriodEnd)

	uuid4Str, err := common.UUIDBytesToStr(invoiceTmp.InvoiceD.Uuid4)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceTmp.InvoiceD.IdS = uuid4Str

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(invoiceTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(invoiceTmp.CrUpdTime.UpdatedAt)

	invoice := invoiceproto.Invoice{InvoiceD: invoiceTmp.InvoiceD, InvoiceT: invoiceT, CrUpdUser: invoiceTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &invoice, nil
}

// UpdateInvoice - Update Invoice
func (invs *InvoiceService) UpdateInvoice(ctx context.Context, in *invoiceproto.UpdateInvoiceRequest) (*invoiceproto.UpdateInvoiceResponse, error) {
	db := invs.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateInvoiceSQL)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.CountryOfSupplyOfGoods,
			in.CreditReasonCode,
			in.DiscountAgreementTerms,
			in.InvoiceCurrencyCode,
			in.InvoiceType,
			in.SupplierAccountReceivable,
			tn,
			uuid4byte)
		if err != nil {
			invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &invoiceproto.UpdateInvoiceResponse{}, nil
}
