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

func TestCreateDespatchAdvice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	despatchAdvice := logisticsproto.CreateDespatchAdviceRequest{}
	despatchAdvice.DeliveryTypeCode = ""
	despatchAdvice.RackIdAtPickUpLocation = ""
	despatchAdvice.TotalDepositAmount = float64(0)
	despatchAdvice.TdaCodeListVersion = ""
	despatchAdvice.TdaCurrencyCode = ""
	despatchAdvice.TotalNumberOfLines = uint32(0)
	despatchAdvice.BlanketOrder = uint32(0)
	despatchAdvice.Buyer = uint32(0)
	despatchAdvice.Carrier = uint32(0)
	despatchAdvice.Contract = uint32(0)
	despatchAdvice.CustomerDocumentReference = uint32(0)
	despatchAdvice.DeclarantsCustomsIdentity = uint32(0)
	despatchAdvice.DeliveryNote = uint32(0)
	despatchAdvice.DeliverySchedule = uint32(0)
	despatchAdvice.DespatchAdviceIdentification = uint32(0)
	despatchAdvice.FreightForwarder = uint32(0)
	despatchAdvice.InventoryLocation = uint32(0)
	despatchAdvice.Invoice = uint32(0)
	despatchAdvice.Invoicee = uint32(0)
	despatchAdvice.LogisticServiceProvider = uint32(0)
	despatchAdvice.OrderResponse = uint32(0)
	despatchAdvice.PickUpLocation = uint32(0)
	despatchAdvice.ProductCertification = uint32(0)
	despatchAdvice.PromotionalDeal = uint32(0)
	despatchAdvice.PurchaseConditions = uint32(0)
	despatchAdvice.PurchaseOrder = uint32(0)
	despatchAdvice.Receiver = uint32(2)
	despatchAdvice.ReturnsInstruction = uint32(0)
	despatchAdvice.Seller = uint32(0)
	despatchAdvice.ShipFrom = uint32(0)
	despatchAdvice.Shipper = uint32(1)
	despatchAdvice.ShipTo = uint32(3)
	despatchAdvice.Specification = uint32(0)
	despatchAdvice.TransportInstruction = uint32(0)
	despatchAdvice.UltimateConsignee = uint32(0)

	data, _ := json.Marshal(&despatchAdvice)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2.3/despatch-advices", bytes.NewBuffer(data))
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
}

func TestGetDespatchAdvices(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/despatch-advices", bytes.NewBuffer(nil))
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
}

func TestGetDespatchAdvice(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/despatch-advices/43a938d2-4550-4dc1-80d8-00f5fd134cca", bytes.NewBuffer(nil))
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
}

func TestGetDespatchAdviceLineItems(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", backendServerAddr+"/v2.3/despatch-advices/43a938d2-4550-4dc1-80d8-00f5fd134cca/lines", bytes.NewBuffer(nil))
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
	despatchAdviceLineItemsResp := &logisticsproto.GetDespatchAdviceLineItemsResponse{}
	err = json.Unmarshal(w.Body.Bytes(), despatchAdviceLineItemsResp)
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotNil(t, despatchAdviceLineItemsResp)
	despatchAdviceLineItemResult := despatchAdviceLineItemsResp.DespatchAdviceLineItems[0]
	assert.Equal(t, despatchAdviceLineItemResult.DespatchAdviceLineItemD.LineItemNumber, uint32(1), "they should be equal")
	assert.Equal(t, despatchAdviceLineItemResult.DespatchAdviceLineItemD.DespatchedQuantity, float64(48), "they should be equal")
	assert.Equal(t, despatchAdviceLineItemResult.DespatchAdviceLineItemD.FreeGoodsQuantity, float64(48), "they should be equal")
	assert.Equal(t, despatchAdviceLineItemResult.DespatchAdviceLineItemD.RequestedQuantity, float64(48), "they should be equal")
	assert.NotNil(t, despatchAdviceLineItemResult.CrUpdTime.CreatedAt)
	assert.NotNil(t, despatchAdviceLineItemResult.CrUpdTime.UpdatedAt)
}
