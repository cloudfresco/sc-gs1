package invoicecontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

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

	invoiceLineItems := []*invoiceproto.CreateInvoiceLineItemRequest{}
	invoiceLineItems = append(invoiceLineItems, &invoiceLineItem)
	invoice.InvoiceLineItems = invoiceLineItems

	data, _ := json.Marshal(&invoice)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2.3/invoices", bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}
	assert.NotNil(t, w.Body.String())
	invoiceResp := &invoiceproto.CreateInvoiceResponse{}
	err = json.Unmarshal(w.Body.Bytes(), invoiceResp)
	if err != nil {
		t.Error(err)
		return
	}
	invoiceResult := invoiceResp.Invoice
	assert.NotNil(t, invoiceResult)
	assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
	assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
	assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
}

func TestGetInvoices(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/invoices", bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}
	assert.NotNil(t, w.Body.String())
	invoiceResp := &invoiceproto.GetInvoicesResponse{}
	err = json.Unmarshal(w.Body.Bytes(), invoiceResp)
	if err != nil {
		t.Error(err)
		return
	}
	invoiceResult := invoiceResp.Invoices[0]
	assert.NotNil(t, invoiceResult)
	assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
	assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
	assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
}

func TestGetInvoice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/invoices/7b71cc13-cc69-44a8-a2b4-4a1cbad97da1", bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}
	assert.NotNil(t, w.Body.String())
	invoiceResp := &invoiceproto.GetInvoiceResponse{}
	err = json.Unmarshal(w.Body.Bytes(), invoiceResp)
	if err != nil {
		t.Error(err)
		return
	}
	invoiceResult := invoiceResp.Invoice
	assert.NotNil(t, invoiceResult)
	assert.Equal(t, invoiceResult.InvoiceD.CreditReasonCode, "DAMAGED_GOODS", "they should be equal")
	assert.Equal(t, invoiceResult.InvoiceD.InvoiceCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, invoiceResult.InvoiceD.InvoiceType, "CREDIT_NOTE", "they should be equal")
	assert.NotNil(t, invoiceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, invoiceResult.CrUpdTime.UpdatedAt)
}

func TestGetInvoiceLineItems(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/invoices/7b71cc13-cc69-44a8-a2b4-4a1cbad97da1/lines", bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}
	assert.NotNil(t, w.Body.String())
	invoiceLineItemsResp := &invoiceproto.GetInvoiceLineItemsResponse{}
	err = json.Unmarshal(w.Body.Bytes(), invoiceLineItemsResp)
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotNil(t, invoiceLineItemsResp)
	invoiceLineItemResult := invoiceLineItemsResp.InvoiceLineItems[0]
	assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.CreditReason, "DAMAGED_GOOD", "they should be equal")
	assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.AiacCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.CreditLineIndicator, "Credit", "they should be equal")
	assert.NotNil(t, invoiceLineItemResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, invoiceLineItemResult.CrUpdTime.UpdatedAt)
	assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.InvoicedQuantity, float64(10), "they should be equal")
}
