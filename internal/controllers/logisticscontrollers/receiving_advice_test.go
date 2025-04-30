package logisticscontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	"github.com/cloudfresco/sc-gs1/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateReceivingAdvice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	receivingAdvice := logisticsproto.CreateReceivingAdviceRequest{}
	receivingAdvice.ReportingCode = "FULL_DETAILS"
	receivingAdvice.TotalAcceptedAmount = float64(0)
	receivingAdvice.TaaCodeListVersion = ""
	receivingAdvice.TaaCurrencyCode = ""
	receivingAdvice.TotalDepositAmount = float64(0)
	receivingAdvice.TdaCodeListVersion = ""
	receivingAdvice.TdaCurrencyCode = ""
	receivingAdvice.TotalNumberOfLines = uint32(0)
	receivingAdvice.TotalOnHoldAmount = float64(0)
	receivingAdvice.TohaCodeListVersion = ""
	receivingAdvice.TohaCurrencyCode = ""
	receivingAdvice.TotalRejectedAmount = float64(0)
	receivingAdvice.TraCodeListVersion = ""
	receivingAdvice.TraCurrencyCode = ""
	receivingAdvice.ReceivingAdviceTransportInformation = uint32(0)
	receivingAdvice.BillOfLadingNumber = uint32(0)
	receivingAdvice.Buyer = uint32(0)
	receivingAdvice.Carrier = uint32(0)
	receivingAdvice.ConsignmentIdentification = uint32(0)
	receivingAdvice.DeliveryNote = uint32(0)
	receivingAdvice.DespatchAdvice = uint32(0)
	receivingAdvice.InventoryLocation = uint32(0)
	receivingAdvice.PurchaseOrder = uint32(0)
	receivingAdvice.Receiver = uint32(2)
	receivingAdvice.ReceivingAdviceIdentification = uint32(0)
	receivingAdvice.Seller = uint32(0)
	receivingAdvice.ShipFrom = uint32(0)
	receivingAdvice.ShipmentIdentification = uint32(0)
	receivingAdvice.Shipper = uint32(1)
	receivingAdvice.ShipTo = uint32(3)
	receivingAdvice.DespatchAdviceDeliveryDateTimeBegin = "04/11/2011"
	receivingAdvice.DespatchAdviceDeliveryDateTimeEnd = "04/11/2011"
	receivingAdvice.PaymentDateTimeBegin = "04/11/2011"
	receivingAdvice.PaymentDateTimeEnd = "04/11/2011"
	receivingAdvice.ReceivingDateTimeBegin = "04/11/2011"
	receivingAdvice.ReceivingDateTimeEnd = "04/11/2011"

	data, _ := json.Marshal(&receivingAdvice)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2.3/receiving-advices", bytes.NewBuffer(data))
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

	receivingAdviceResp := &logisticsproto.CreateReceivingAdviceResponse{}
	err = json.Unmarshal(w.Body.Bytes(), receivingAdviceResp)
	if err != nil {
		t.Error(err)
		return
	}

	receivingAdviceResult := receivingAdviceResp.ReceivingAdvice
	assert.NotNil(t, receivingAdviceResult)
	assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
	assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
}

func TestGetReceivingAdvices(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/receiving-advices", bytes.NewBuffer(nil))
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

	receivingAdviceResp := &logisticsproto.GetReceivingAdvicesResponse{}
	err = json.Unmarshal(w.Body.Bytes(), receivingAdviceResp)
	if err != nil {
		t.Error(err)
		return
	}

	receivingAdviceResult := receivingAdviceResp.ReceivingAdvices[0]
	assert.NotNil(t, receivingAdviceResult)
	assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
	assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
}

func TestGetReceivingAdvice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/receiving-advices/a2a9692a-6bf4-41c1-a0b0-e27691b365c9", bytes.NewBuffer(nil))
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

	receivingAdviceResp := &logisticsproto.GetReceivingAdviceResponse{}
	err = json.Unmarshal(w.Body.Bytes(), receivingAdviceResp)
	if err != nil {
		t.Error(err)
		return
	}

	receivingAdviceResult := receivingAdviceResp.ReceivingAdvice
	assert.NotNil(t, receivingAdviceResult)
	assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
	assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
}

func TestGetReceivingAdviceLineItems(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/receiving-advices/a2a9692a-6bf4-41c1-a0b0-e27691b365c9/lines", bytes.NewBuffer(nil))
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

	receivingAdviceLineItemsResp := &logisticsproto.GetReceivingAdviceLineItemsResponse{}
	err = json.Unmarshal(w.Body.Bytes(), receivingAdviceLineItemsResp)
	if err != nil {
		t.Error(err)
		return
	}

	assert.NotNil(t, receivingAdviceLineItemsResp)
	receivingAdviceLineItemResult := receivingAdviceLineItemsResp.ReceivingAdviceLineItems[0]
	assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.LineItemNumber, uint32(1), "they should be equal")
	assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.QuantityAccepted, float64(38), "they should be equal")
	assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.QuantityReceived, float64(48), "they should be equal")
	assert.NotNil(t, receivingAdviceLineItemResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, receivingAdviceLineItemResult.CrUpdTime.UpdatedAt)
}
