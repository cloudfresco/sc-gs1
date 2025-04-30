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

func TestCreateDebitCreditAdvice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	debitCreditAdvice := invoiceproto.CreateDebitCreditAdviceRequest{}
	debitCreditAdvice.DebitCreditIndicatorCode = "CREDIT"
	debitCreditAdvice.TotalAmount = float64(100)
	debitCreditAdvice.TaCodeListVersion = ""
	debitCreditAdvice.TaCurrencyCode = "EUR"
	debitCreditAdvice.BillTo = uint32(0)
	debitCreditAdvice.Buyer = uint32(2)
	debitCreditAdvice.Carrier = uint32(0)
	debitCreditAdvice.DebitCreditAdviceIdentification = uint32(0)
	debitCreditAdvice.Seller = uint32(1)
	debitCreditAdvice.ShipFrom = uint32(0)
	debitCreditAdvice.ShipTo = uint32(0)
	debitCreditAdvice.UltimateConsignee = uint32(0)

	debitCreditAdviceLineItem := invoiceproto.CreateDebitCreditAdviceLineItemRequest{}
	debitCreditAdviceLineItem.AdjustmentAmount = float64(100)
	debitCreditAdviceLineItem.AaCodeListVersion = ""
	debitCreditAdviceLineItem.AaCurrencyCode = "EUR"
	debitCreditAdviceLineItem.DebitCreditIndicatorCode = "CREDIT"
	debitCreditAdviceLineItem.FinancialAdjustmentReasonCode = "17"
	debitCreditAdviceLineItem.LineItemNumber = uint32(1)
	debitCreditAdviceLineItem.ParentLineItemNumber = uint32(0)
	debitCreditAdviceLineItem.DebitCreditAdviceId = uint32(1)

	debitCreditAdviceLineItems := []*invoiceproto.CreateDebitCreditAdviceLineItemRequest{}
	debitCreditAdviceLineItems = append(debitCreditAdviceLineItems, &debitCreditAdviceLineItem)
	debitCreditAdvice.DebitCreditAdviceLineItems = debitCreditAdviceLineItems

	data, _ := json.Marshal(&debitCreditAdvice)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2.3/debit-credit-advices", bytes.NewBuffer(data))
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
	debitCreditAdviceResp := &invoiceproto.CreateDebitCreditAdviceResponse{}
	err = json.Unmarshal(w.Body.Bytes(), debitCreditAdviceResp)
	if err != nil {
		t.Error(err)
		return
	}
	debitCreditAdviceResult := debitCreditAdviceResp.DebitCreditAdvice
	assert.NotNil(t, debitCreditAdviceResult)
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
	assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
}

func TestGetDebitCreditAdvices(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/debit-credit-advices", bytes.NewBuffer(nil))
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
	debitCreditAdviceResp := &invoiceproto.GetDebitCreditAdvicesResponse{}
	err = json.Unmarshal(w.Body.Bytes(), debitCreditAdviceResp)
	if err != nil {
		t.Error(err)
		return
	}
	debitCreditAdviceResult := debitCreditAdviceResp.DebitCreditAdvices[0]
	assert.NotNil(t, debitCreditAdviceResult)
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
	assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
}

func TestGetDebitCreditAdvice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/debit-credit-advices/f7a9f263-a1e1-4307-8990-6f7e4d94df08", bytes.NewBuffer(nil))
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
	debitCreditAdviceResp := &invoiceproto.GetDebitCreditAdviceResponse{}
	err = json.Unmarshal(w.Body.Bytes(), debitCreditAdviceResp)
	if err != nil {
		t.Error(err)
		return
	}
	debitCreditAdviceResult := debitCreditAdviceResp.DebitCreditAdvice
	assert.NotNil(t, debitCreditAdviceResult)
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
	assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
}

func TestGetDebitCreditAdviceLineItems(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/debit-credit-advices/f7a9f263-a1e1-4307-8990-6f7e4d94df08/lines", bytes.NewBuffer(nil))
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
	debitCreditAdviceLineItemResponse := &invoiceproto.GetDebitCreditAdviceLineItemsResponse{}
	err = json.Unmarshal(w.Body.Bytes(), debitCreditAdviceLineItemResponse)
	if err != nil {
		t.Error(err)
		return
	}
	debitCreditAdviceLineItemResult := debitCreditAdviceLineItemResponse.DebitCreditAdviceLineItems[0]
	assert.NotNil(t, debitCreditAdviceLineItemResult)
	assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.AdjustmentAmount, float64(100), "they should be equal")
	assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.AaCurrencyCode, "EUR", "they should be equal")
	assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
	assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.FinancialAdjustmentReasonCode, "17", "they should be equal")
	assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.LineItemNumber, uint32(1), "they should be equal")
	assert.NotNil(t, debitCreditAdviceLineItemResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, debitCreditAdviceLineItemResult.CrUpdTime.UpdatedAt)
}
