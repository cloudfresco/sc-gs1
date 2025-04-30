package invoiceservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestInvoiceService_CreateInvoice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	invoiceService := NewInvoiceService(log, dbService, redisService, userServiceClient)
	invoice := invoiceproto.CreateInvoiceRequest{}
	invoice.CountryOfSupplyOfGoods = ""
	invoice.CreditReasonCode = "DAMAGED_GOODS"
	invoice.DiscountAgreementTerms = ""
	invoice.InvoiceCurrencyCode = "EUR"
	invoice.InvoiceType = "CREDIT_NOTE"
	invoice.IsBuyerBasedInEu = false
	invoice.IsFirstSellerBasedInEu = false
	invoice.SupplierAccountReceivable = ""
	invoice.BlanketOrder = uint32(0)
	invoice.Buyer = uint32(13)
	invoice.Contract = uint32(0)
	invoice.DeliveryNote = uint32(0)
	invoice.DespatchAdvice = uint32(0)
	invoice.DisputeNotice = uint32(0)
	invoice.InventoryLocation = uint32(0)
	invoice.InventoryReport = uint32(0)
	invoice.Invoice = uint32(0)
	invoice.InvoiceIdentification = uint32(0)
	invoice.Manifest = uint32(0)
	invoice.OrderResponse = uint32(0)
	invoice.Payee = uint32(0)
	invoice.Payer = uint32(0)
	invoice.PickupFrom = uint32(0)
	invoice.PriceList = uint32(0)
	invoice.PromotionalDeal = uint32(0)
	invoice.PurchaseOrder = uint32(0)
	invoice.ReceivingAdvice = uint32(0)
	invoice.RemitTo = uint32(0)
	invoice.ReturnsNotice = uint32(0)
	invoice.SalesOrder = uint32(0)
	invoice.SalesReport = uint32(0)
	invoice.Seller = uint32(1)
	invoice.ShipFrom = uint32(0)
	invoice.ShipTo = uint32(0)
	invoice.SupplierAgentRepresentative = uint32(0)
	invoice.SupplierCorporateOffice = uint32(0)
	invoice.TaxCurrencyInformation = uint32(0)
	invoice.TaxRepresentative = uint32(0)
	invoice.TradeAgreement = uint32(0)
	invoice.UltimateConsignee = uint32(0)
	invoice.ActualDeliveryDate = "04/18/2023"
	invoice.InvoicingPeriodBegin = "04/18/2023"
	invoice.InvoicingPeriodEnd = "04/18/2023"
	invoice.UserId = "auth0|673ee1a719dd4000cd5a3832"
	invoice.UserEmail = "sprov300@gmail.com"
	invoice.RequestId = "bks1m1g91jau4nkks2f0"

	invoiceLineItem := invoiceproto.CreateInvoiceLineItemRequest{}
	invoiceLineItem.AmountExclusiveAllowancesCharges = float64(0)
	invoiceLineItem.AeacCodeListVersion = ""
	invoiceLineItem.AeacCurrencyCode = ""
	invoiceLineItem.AmountInclusiveAllowancesCharges = float64(100)
	invoiceLineItem.AiacCodeListVersion = ""
	invoiceLineItem.AiacCurrencyCode = "EUR"
	invoiceLineItem.CreditLineIndicator = "Credit"
	invoiceLineItem.CreditReason = "DAMAGED_GOOD"
	invoiceLineItem.DeliveredQuantity = float64(10)
	invoiceLineItem.DqMeasurementUnitCode = ""
	invoiceLineItem.DqCodeListVersion = ""
	invoiceLineItem.ExcludedFromPaymentDiscountIndicator = false
	invoiceLineItem.Extension = ""
	invoiceLineItem.FreeGoodsQuantity = float64(0)
	invoiceLineItem.FgqMeasurementUnitCode = ""
	invoiceLineItem.FgqCodeListVersion = ""
	invoiceLineItem.InvoicedQuantity = float64(10)
	invoiceLineItem.IqMeasurementUnitCode = ""
	invoiceLineItem.IqCodeListVersion = ""
	invoiceLineItem.ItemPriceBaseQuantity = float64(0)
	invoiceLineItem.IpbqMeasurementUnitCode = ""
	invoiceLineItem.IpbqCodeListVersion = ""
	invoiceLineItem.ItemPriceExclusiveAllowancesCharges = float64(15)
	invoiceLineItem.IpeacCodeListVersion = ""
	invoiceLineItem.IpeacCurrencyCode = "EUR"
	invoiceLineItem.ItemPriceInclusiveAllowancesCharges = float64(10)
	invoiceLineItem.IpiacCodeListVersion = ""
	invoiceLineItem.IpiacCurrencyCode = ""
	invoiceLineItem.LegallyFixedRetailPrice = float64(0)
	invoiceLineItem.LfrpCodeListVersion = ""
	invoiceLineItem.LfrpCurrencyCode = ""
	invoiceLineItem.LineItemNumber = uint32(1)
	invoiceLineItem.MarginSchemeInformation = ""
	invoiceLineItem.OwenrshipPriorToPayment = ""
	invoiceLineItem.ParentLineItemNumber = uint32(0)
	invoiceLineItem.RecommendedRetailPrice = float64(0)
	invoiceLineItem.RrpCodeListVersion = ""
	invoiceLineItem.RrpCurrencyCode = ""
	invoiceLineItem.RetailPriceExcludingExcise = float64(0)
	invoiceLineItem.RpeeCodeListVersion = ""
	invoiceLineItem.RpeeCurrencyCode = ""
	invoiceLineItem.TotalOrderedQuantity = float64(15)
	invoiceLineItem.ToqMeasurementUnitCode = ""
	invoiceLineItem.ToqCodeListVersion = ""
	invoiceLineItem.ConsumptionReport = uint32(0)
	invoiceLineItem.Contract = uint32(0)
	invoiceLineItem.DeliveryNote = uint32(0)
	invoiceLineItem.DespatchAdvice = uint32(0)
	invoiceLineItem.EnergyQuantity = uint32(0)
	invoiceLineItem.InventoryLocationFrom = uint32(0)
	invoiceLineItem.InventoryLocationTo = uint32(0)
	invoiceLineItem.PromotionalDeal = uint32(0)
	invoiceLineItem.PurchaseConditions = uint32(0)
	invoiceLineItem.PurchaseOrder = uint32(0)
	invoiceLineItem.ReceivingAdvice = uint32(0)
	invoiceLineItem.ReturnableAssetIdentification = uint32(0)
	invoiceLineItem.SalesOrder = uint32(0)
	invoiceLineItem.ShipFrom = uint32(0)
	invoiceLineItem.ShipTo = uint32(0)
	invoiceLineItem.TradeAgreement = uint32(0)
	invoiceLineItem.TransferOfOwnershipDate = "04/11/2011"
	invoiceLineItem.ActualDeliveryDate = "04/11/2011"
	invoiceLineItem.ServicetimePeriodLineLevelBegin = "04/11/2011"
	invoiceLineItem.ServicetimePeriodLineLevelEnd = "04/11/2011"
	invoiceLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	invoiceLineItem.UserEmail = "sprov300@gmail.com"
	invoiceLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	invoiceLineItems := []*invoiceproto.CreateInvoiceLineItemRequest{}
	invoiceLineItems = append(invoiceLineItems, &invoiceLineItem)
	invoice.InvoiceLineItems = invoiceLineItems
	type args struct {
		ctx context.Context
		in  *invoiceproto.CreateInvoiceRequest
	}
	tests := []struct {
		invs    *InvoiceService
		args    args
		wantErr bool
	}{
		{
			invs: invoiceService,
			args: args{
				ctx: ctx,
				in:  &invoice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		invoiceResponse, err := tt.invs.CreateInvoice(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("InvoiceService.CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		invoiceResult := invoiceResponse.Invoice
		assert.NotNil(t, invoiceResult)
		assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
		assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
	}
}

func TestInvoiceService_GetInvoices(t *testing.T) {
	var err error
	err = test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	invoiceService := NewInvoiceService(log, dbService, redisService, userServiceClient)
	invoices := []*invoiceproto.Invoice{}

	invoice1, err := GetInvoice(uint32(1), []byte{120, 94, 225, 36, 51, 245, 77, 238, 183, 32, 202, 55, 252, 113, 85, 238}, "785ee124-33f5-4dee-b720-ca37fc7155ee", "", "", "", "EUR", "INVOICE", false, false, "", uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), "2011-04-14T10:15:00Z", "2011-04-12T10:15:00Z", "2011-04-12T10:15:00Z", "2011-04-12T10:15:00Z", "2011-04-12T10:15:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	invoice2, err := GetInvoice(uint32(2), []byte{123, 113, 204, 19, 204, 105, 68, 168, 162, 180, 74, 28, 186, 217, 125, 161}, "7b71cc13-cc69-44a8-a2b4-4a1cbad97da1", "", "DAMAGED_GOODS", "", "EUR", "CREDIT_NOTE", false, false, "", uint32(0), uint32(13), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), "2011-04-18T10:15:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	invoices = append(invoices, invoice2, invoice1)

	form := invoiceproto.GetInvoicesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	invoicesResponse := invoiceproto.GetInvoicesResponse{Invoices: invoices, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *invoiceproto.GetInvoicesRequest
	}
	tests := []struct {
		invs    *InvoiceService
		args    args
		want    *invoiceproto.GetInvoicesResponse
		wantErr bool
	}{
		{
			invs: invoiceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &invoicesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		invoiceResponse, err := tt.invs.GetInvoices(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("InvoiceService.GetInvoices() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(invoiceResponse, tt.want) {
			t.Errorf("InvoiceService.GetInvoices() = %v, want %v", invoiceResponse, tt.want)
		}
		invoiceResult := invoiceResponse.Invoices[0]
		assert.NotNil(t, invoiceResult)
		assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
		assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
	}
}

func TestInvoiceService_GetInvoice(t *testing.T) {
	var err error
	err = test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	invoiceService := NewInvoiceService(log, dbService, redisService, userServiceClient)

	invoice, err := GetInvoice(uint32(2), []byte{123, 113, 204, 19, 204, 105, 68, 168, 162, 180, 74, 28, 186, 217, 125, 161}, "7b71cc13-cc69-44a8-a2b4-4a1cbad97da1", "", "DAMAGED_GOODS", "", "EUR", "CREDIT_NOTE", false, false, "", uint32(0), uint32(13), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), "2011-04-18T10:15:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	invoiceResponse := invoiceproto.GetInvoiceResponse{}
	invoiceResponse.Invoice = invoice

	form := invoiceproto.GetInvoiceRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "7b71cc13-cc69-44a8-a2b4-4a1cbad97da1"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *invoiceproto.GetInvoiceRequest
	}
	tests := []struct {
		invs    *InvoiceService
		args    args
		want    *invoiceproto.GetInvoiceResponse
		wantErr bool
	}{
		{
			invs: invoiceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &invoiceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		invoiceResp, err := tt.invs.GetInvoice(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("InvoiceService.GetInvoice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(invoiceResp, tt.want) {
			t.Errorf("InvoiceService.GetInvoice() = %v, want %v", invoiceResp, tt.want)
		}
		invoiceResult := invoiceResp.Invoice
		assert.NotNil(t, invoiceResult)
		assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
		assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
	}
}

func TestInvoiceService_GetInvoiceByPk(t *testing.T) {
	var err error
	err = test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	invoiceService := NewInvoiceService(log, dbService, redisService, userServiceClient)

	invoice, err := GetInvoice(uint32(2), []byte{123, 113, 204, 19, 204, 105, 68, 168, 162, 180, 74, 28, 186, 217, 125, 161}, "7b71cc13-cc69-44a8-a2b4-4a1cbad97da1", "", "DAMAGED_GOODS", "", "EUR", "CREDIT_NOTE", false, false, "", uint32(0), uint32(13), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), "2011-04-18T10:15:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "2011-04-15T10:20:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	invoiceResponse := invoiceproto.GetInvoiceByPkResponse{}
	invoiceResponse.Invoice = invoice

	form := invoiceproto.GetInvoiceByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(2)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *invoiceproto.GetInvoiceByPkRequest
	}
	tests := []struct {
		invs    *InvoiceService
		args    args
		want    *invoiceproto.GetInvoiceByPkResponse
		wantErr bool
	}{
		{
			invs: invoiceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &invoiceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		invoiceResp, err := tt.invs.GetInvoiceByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("InvoiceService.GetInvoiceByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(invoiceResp, tt.want) {
			t.Errorf("InvoiceService.GetInvoiceByPk() = %v, want %v", invoiceResp, tt.want)
		}
		invoiceResult := invoiceResp.Invoice
		assert.NotNil(t, invoiceResult)
		assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
		assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
	}
}

func GetInvoice(id uint32, uuid4 []byte, idS string, countryOfSupplyOfGoods string, creditReasonCode string, discountAgreementTerms string, invoiceCurrencyCode string, invoiceType string, isBuyerBasedInEu bool, isFirstSellerBasedInEu bool, supplierAccountReceivable string, blanketOrder uint32, buyer uint32, contract uint32, deliveryNote uint32, despatchAdvice uint32, disputeNotice uint32, inventoryLocation uint32, inventoryReport uint32, invoice uint32, invoiceIdentification uint32, manifest uint32, orderResponse uint32, payee uint32, payer uint32, pickupFrom uint32, priceList uint32, promotionalDeal uint32, purchaseOrder uint32, receivingAdvice uint32, remitTo uint32, returnsNotice uint32, salesOrder uint32, salesReport uint32, seller uint32, shipFrom uint32, shipTo uint32, supplierAgentRepresentative uint32, supplierCorporateOffice uint32, taxCurrencyInformation uint32, taxRepresentative uint32, tradeAgreement uint32, ultimateConsignee uint32, actualDeliveryDate string, invoicingPeriodBegin string, invoicingPeriodEnd string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*invoiceproto.Invoice, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	actualDeliveryDate1, err := common.ConvertTimeToTimestamp(Layout, actualDeliveryDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	invoicingPeriodBegin1, err := common.ConvertTimeToTimestamp(Layout, invoicingPeriodBegin)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	invoicingPeriodEnd1, err := common.ConvertTimeToTimestamp(Layout, invoicingPeriodEnd)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}
	invoiceD := invoiceproto.InvoiceD{}
	invoiceD.Id = id
	invoiceD.Uuid4 = uuid4
	invoiceD.IdS = idS
	invoiceD.CountryOfSupplyOfGoods = countryOfSupplyOfGoods
	invoiceD.CreditReasonCode = creditReasonCode
	invoiceD.DiscountAgreementTerms = discountAgreementTerms
	invoiceD.InvoiceCurrencyCode = invoiceCurrencyCode
	invoiceD.InvoiceType = invoiceType
	invoiceD.IsBuyerBasedInEu = isBuyerBasedInEu
	invoiceD.IsFirstSellerBasedInEu = isFirstSellerBasedInEu
	invoiceD.SupplierAccountReceivable = supplierAccountReceivable
	invoiceD.BlanketOrder = blanketOrder
	invoiceD.Buyer = buyer
	invoiceD.Contract = contract
	invoiceD.DeliveryNote = deliveryNote
	invoiceD.DespatchAdvice = despatchAdvice
	invoiceD.DisputeNotice = disputeNotice
	invoiceD.InventoryLocation = inventoryLocation
	invoiceD.InventoryReport = inventoryReport
	invoiceD.Invoice = invoice
	invoiceD.InvoiceIdentification = invoiceIdentification
	invoiceD.Manifest = manifest
	invoiceD.OrderResponse = orderResponse
	invoiceD.Payee = payee
	invoiceD.Payer = payer
	invoiceD.PickupFrom = pickupFrom
	invoiceD.PriceList = priceList
	invoiceD.PromotionalDeal = promotionalDeal
	invoiceD.PurchaseOrder = purchaseOrder
	invoiceD.ReceivingAdvice = receivingAdvice
	invoiceD.RemitTo = remitTo
	invoiceD.ReturnsNotice = returnsNotice
	invoiceD.SalesOrder = salesOrder
	invoiceD.SalesReport = salesReport
	invoiceD.Seller = seller
	invoiceD.ShipFrom = shipFrom
	invoiceD.ShipTo = shipTo
	invoiceD.SupplierAgentRepresentative = supplierAgentRepresentative
	invoiceD.SupplierCorporateOffice = supplierCorporateOffice
	invoiceD.TaxCurrencyInformation = taxCurrencyInformation
	invoiceD.TaxRepresentative = taxRepresentative
	invoiceD.TradeAgreement = tradeAgreement
	invoiceD.UltimateConsignee = ultimateConsignee

	invoiceT := invoiceproto.InvoiceT{}
	invoiceT.ActualDeliveryDate = actualDeliveryDate1
	invoiceT.InvoicingPeriodBegin = invoicingPeriodBegin1
	invoiceT.InvoicingPeriodEnd = invoicingPeriodEnd1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	invoiceResult := invoiceproto.Invoice{InvoiceD: &invoiceD, InvoiceT: &invoiceT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &invoiceResult, nil
}
