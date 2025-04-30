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

func TestCreateOrderResponse(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	orderResponse := orderproto.CreateOrderResponseRequest{}
	orderResponse.OrderResponseReasonCode = "DISCONTINUED_LINE"
	orderResponse.ResponseStatusCode = "REJECTED"
	orderResponse.TotalMonetaryAmountExcludingTaxes = float64(1257)
	orderResponse.TmaetCodeListVersion = ""
	orderResponse.TmaetCurrencyCode = ""
	orderResponse.TotalMonetaryAmountIncludingTaxes = float64(0)
	orderResponse.TmaitCodeListVersion = ""
	orderResponse.TmaitCurrencyCode = ""
	orderResponse.TotalTaxAmount = float64(450)
	orderResponse.TtaCodeListVersion = ""
	orderResponse.TtaCurrencyCode = ""
	orderResponse.AmendedDateTimeValue = uint32(0)
	orderResponse.BillTo = uint32(0)
	orderResponse.Buyer = uint32(2)
	orderResponse.OrderResponseIdentification = uint32(4)
	orderResponse.OriginalOrder = uint32(0)
	orderResponse.SalesOrder = uint32(0)
	orderResponse.Seller = uint32(1)

	data, _ := json.Marshal(&orderResponse)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2.3/order-responses", bytes.NewBuffer(data))
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
	orderResp := &orderproto.CreateOrderResponseResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderRespResult := orderResp.OrderResponse
	assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
	assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")
	assert.Equal(t, orderRespResult.OrderResponseD.TotalMonetaryAmountExcludingTaxes, float64(1257), "they should be equal")
	assert.Equal(t, orderRespResult.OrderResponseD.TotalTaxAmount, float64(450), "they should be equal")
}

func TestOrderResponses(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/order-responses", nil)
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
	orderResp := &orderproto.GetOrderResponsesResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderRespResult := orderResp.OrderResponses[1]
	assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
	assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")
}

func TestOrderResponse(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/order-responses/0542056b-7fb4-4a79-9e0e-e65bdb07a2a2", nil)
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
	orderResp := &orderproto.GetOrderResponseResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResp)
	if err != nil {
		t.Error(err)
		return
	}
	orderRespResult := orderResp.OrderResponse
	assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
	assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")
}

func TestGetOrderResponseLineItems(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/order-responses/a9cf4bc4-bf4b-4013-95d3-d3bb4c7af952/lines", nil)
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
	orderResponseLineItemResp := &orderproto.GetOrderResponseLineItemsResponse{}
	err = json.Unmarshal(w.Body.Bytes(), orderResponseLineItemResp)
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotNil(t, orderResponseLineItemResp)
	orderResponseLineItemResult := orderResponseLineItemResp.OrderResponseLineItems[1]
	assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.LineItemNumber, uint32(2), "they should be equal")
	assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.ConfirmedQuantity, float64(24), "they should be equal")
	assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.LineItemChangeIndicator, "MODIFIED", "they should be equal")
	assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.OrderResponseReasonCode, "PRODUCT_OUT_OF_STOCK", "they should be equal")
	assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.OriginalOrderLineItemNumber, uint32(2), "they should be equal")
}
