package ordercontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	order := orderproto.CreateOrderRequest{}
	order.IsApplicationReceiptAcknowledgementRequired = true
	order.IsOrderFreeOfExciseTaxDuty = false
	order.OrderChangeReasonCode = ""
	order.OrderEntryType = ""
	order.OrderInstructionCode = "PARTIAL_DELIVERY_ALLOWED"
	order.OrderPriority = ""
	order.OrderTypeCode = "402"
	order.TotalMonetaryAmountExcludingTaxes = float64(12675)
	order.TmaetCodeListVersion = ""
	order.TmaetCurrencyCode = ""
	order.TotalMonetaryAmountIncludingTaxes = float64(0)
	order.TmaitCodeListVersion = ""
	order.TmaitCurrencyCode = ""
	order.TotalTaxAmount = float64(2661.75)
	order.TtaCodeListVersion = ""
	order.TtaCurrencyCode = ""
	order.BillTo = uint32(0)
	order.Buyer = uint32(2)
	order.Contract = uint32(0)
	order.CustomerDocumentReference = uint32(0)
	order.CustomsBroker = uint32(0)
	order.OrderIdentification = uint32(2)
	order.PickupFrom = uint32(0)
	order.PromotionalDeal = uint32(0)
	order.QuoteNumber = "ASP0002NET"
	order.Seller = uint32(1)
	order.TradeAgreement = "56895632"
	order.DeliveryDateAccordingToSchedule = "04/10/2023"
	order.LatestDeliveryDate = "10/10/2023"

	orderLineItem := orderproto.CreateOrderLineItemRequest{}
	orderLineItem.Extension = ""
	orderLineItem.FreeGoodsQuantity = float64(0)
	orderLineItem.FgqMeasurementUnitCode = ""
	orderLineItem.FgqCodeListVersion = ""
	orderLineItem.ItemPriceBaseQuantity = float64(0)
	orderLineItem.IpbqMeasurementUnitCode = ""
	orderLineItem.IpbqCodeListVersion = ""
	orderLineItem.ItemSourceCode = ""
	orderLineItem.LineItemActionCode = ""
	orderLineItem.LineItemNumber = uint32(2)
	orderLineItem.ListPrice = float64(0)
	orderLineItem.LpCodeListVersion = ""
	orderLineItem.LpCurrencyCode = ""
	orderLineItem.MonetaryAmountExcludingTaxes = float64(0)
	orderLineItem.MaetCodeListVersion = ""
	orderLineItem.MaetCurrencyCode = ""
	orderLineItem.MonetaryAmountIncludingTaxes = float64(0)
	orderLineItem.MaitCodeListVersion = ""
	orderLineItem.MaitCurrencyCode = ""
	orderLineItem.NetAmount = float64(4659)
	orderLineItem.NaCodeListVersion = ""
	orderLineItem.NaCurrencyCode = "EUR"
	orderLineItem.NetPrice = float64(194.125)
	orderLineItem.NpCodeListVersion = ""
	orderLineItem.NpCurrencyCode = "EUR"
	orderLineItem.OrderInstructionCode = ""
	orderLineItem.OrderLineItemInstructionCode = ""
	orderLineItem.OrderLineItemPriority = ""
	orderLineItem.ParentLineItemNumber = uint32(0)
	orderLineItem.RecommendedRetailPrice = float64(0)
	orderLineItem.RequestedQuantity = float64(24)
	orderLineItem.RqMeasurementUnitCode = "EA"
	orderLineItem.RqCodeListVersion = ""
	orderLineItem.ReturnReasonCode = ""
	orderLineItem.Contract = uint32(0)
	orderLineItem.CustomerDocumentReference = uint32(0)
	orderLineItem.DeliveryDateAccordingToSchedule = uint32(0)
	orderLineItem.DespatchAdvice = uint32(0)
	orderLineItem.MaterialSpecification = uint32(0)
	orderLineItem.OrderLineItemContact = uint32(0)
	orderLineItem.PreferredManufacturer = uint32(0)
	orderLineItem.PromotionalDeal = uint32(0)
	orderLineItem.PurchaseConditions = uint32(0)
	orderLineItem.ReturnableAssetIdentification = uint32(0)
	orderLineItem.LatestDeliveryDate = "02/19/2023"

	orderLineItems := []*orderproto.CreateOrderLineItemRequest{}
	orderLineItems = append(orderLineItems, &orderLineItem)
	order.OrderLineItems = orderLineItems

	data, _ := json.Marshal(&order)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2.3/orders", bytes.NewBuffer(data))
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

	orderResp := &orderproto.CreateOrderResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderResult := orderResp.Order
	assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
	assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
	assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
	assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
	assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")
}

func TestGetOrder(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/orders/23e09959-23b2-470a-81ba-0fb8bb22a562", nil)
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

	orderResp := &orderproto.GetOrderResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderResult := orderResp.Order
	assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
	assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
	assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
	assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
	assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")
}

func TestGetOrders(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/orders", nil)
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

	orderResp := &orderproto.GetOrdersResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderResult := orderResp.Orders[0]
	assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
	assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
	assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
	assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
	assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")
}

func TestGetOrderLineItems(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/orders/23e09959-23b2-470a-81ba-0fb8bb22a562/lines", nil)
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
	orderLineItemResp := &orderproto.GetOrderLineItemsResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderLineItemResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderLineItemResult := orderLineItemResp.OrderLineItems[1]
	assert.Equal(t, orderLineItemResult.OrderLineItemD.LineItemNumber, uint32(2), "they should be equal")
	assert.Equal(t, orderLineItemResult.OrderLineItemD.NetAmount, float64(4659), "they should be equal")
	assert.Equal(t, orderLineItemResult.OrderLineItemD.NaCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, orderLineItemResult.OrderLineItemD.NetPrice, float64(194.125), "they should be equal")
	assert.Equal(t, orderLineItemResult.OrderLineItemD.NpCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, orderLineItemResult.OrderLineItemD.RqMeasurementUnitCode, "EA", "they should be equal")
}
