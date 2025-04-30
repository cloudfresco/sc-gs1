package logisticsservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestReceivingAdviceService_CreateReceivingAdviceLineItem(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, userServiceClient)
	receivingAdviceLineItem := logisticsproto.CreateReceivingAdviceLineItemRequest{}
	receivingAdviceLineItem.LineItemNumber = uint32(3)
	receivingAdviceLineItem.ParentLineItemNumber = uint32(0)
	receivingAdviceLineItem.QuantityAccepted = float64(18)
	receivingAdviceLineItem.QaMeasurementUnitCode = ""
	receivingAdviceLineItem.QaCodeListVersion = ""
	receivingAdviceLineItem.QuantityDespatched = float64(0)
	receivingAdviceLineItem.QdMeasurementUnitCode = ""
	receivingAdviceLineItem.QdCodeListVersion = ""
	receivingAdviceLineItem.QuantityReceived = float64(18)
	receivingAdviceLineItem.QrMeasurementUnitCode = ""
	receivingAdviceLineItem.QrCodeListVersion = ""
	receivingAdviceLineItem.TransactionalTradeItem = uint32(0)
	receivingAdviceLineItem.EcomConsignmentIdentification = uint32(0)
	receivingAdviceLineItem.Contract = uint32(0)
	receivingAdviceLineItem.CustomerReference = uint32(0)
	receivingAdviceLineItem.DeliveryNote = uint32(0)
	receivingAdviceLineItem.DespatchAdvice = uint32(0)
	receivingAdviceLineItem.ProductCertification = uint32(0)
	receivingAdviceLineItem.PromotionalDeal = uint32(0)
	receivingAdviceLineItem.PurchaseConditions = uint32(0)
	receivingAdviceLineItem.PurchaseOrder = uint32(0)
	receivingAdviceLineItem.RequestedItemIdentification = uint32(14)
	receivingAdviceLineItem.Specification = uint32(0)
	receivingAdviceLineItem.ReceivingAdviceId = uint32(1)
	receivingAdviceLineItem.PickUpDateTimeBegin = "04/11/2011"
	receivingAdviceLineItem.PickUpDateTimeEnd = "04/12/2011"
	receivingAdviceLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	receivingAdviceLineItem.UserEmail = "sprov300@gmail.com"
	receivingAdviceLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *logisticsproto.CreateReceivingAdviceLineItemRequest
	}
	tests := []struct {
		ras     *ReceivingAdviceService
		args    args
		want    *logisticsproto.CreateReceivingAdviceLineItemResponse
		wantErr bool
	}{
		{
			ras: receivingAdviceService,
			args: args{
				ctx: ctx,
				in:  &receivingAdviceLineItem,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		receivingAdviceLineItemResponse, err := tt.ras.CreateReceivingAdviceLineItem(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReceivingAdviceService.CreateReceivingAdviceLineItem() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		receivingAdviceLineItemResult := receivingAdviceLineItemResponse.ReceivingAdviceLineItem
		assert.NotNil(t, receivingAdviceLineItemResult)
		assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.LineItemNumber, uint32(3), "they should be equal")
		assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.QuantityAccepted, float64(18), "they should be equal")
		assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.QuantityReceived, float64(18), "they should be equal")
		assert.NotNil(t, receivingAdviceLineItemResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, receivingAdviceLineItemResult.CrUpdTime.UpdatedAt)
	}
}

func TestReceivingAdviceService_GetReceivingAdviceLineItems(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, userServiceClient)

	receivingAdviceLineItem1, err := GetReceivingAdviceLineItem(uint32(1), []byte{14, 219, 132, 46, 147, 31, 74, 156, 141, 198, 62, 91, 130, 83, 209, 237}, "0edb842e-931f-4a9c-8dc6-3e5b8253d1ed", uint32(1), uint32(0), float64(38), "", "", float64(0), "", "", float64(48), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(14), uint32(0), uint32(1), "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
	}

	receivingAdviceLineItem2, err := GetReceivingAdviceLineItem(uint32(2), []byte{75, 204, 166, 156, 194, 174, 64, 112, 132, 16, 24, 201, 147, 86, 99, 102}, "4bcca69c-c2ae-4070-8410-18c993566366", uint32(2), uint32(0), float64(24), "", "", float64(0), "", "", float64(24), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(15), uint32(0), uint32(1), "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
	}
	receivingAdviceLineItems := []*logisticsproto.ReceivingAdviceLineItem{}
	receivingAdviceLineItems = append(receivingAdviceLineItems, receivingAdviceLineItem1, receivingAdviceLineItem2)

	receivingAdviceLineItemsResponse := logisticsproto.GetReceivingAdviceLineItemsResponse{}
	receivingAdviceLineItemsResponse.ReceivingAdviceLineItems = receivingAdviceLineItems

	form := logisticsproto.GetReceivingAdviceLineItemsRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "a2a9692a-6bf4-41c1-a0b0-e27691b365c9"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *logisticsproto.GetReceivingAdviceLineItemsRequest
	}
	tests := []struct {
		ras     *ReceivingAdviceService
		args    args
		want    *logisticsproto.GetReceivingAdviceLineItemsResponse
		wantErr bool
	}{
		{
			ras: receivingAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &receivingAdviceLineItemsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		receivingAdviceLineItemsResp, err := tt.ras.GetReceivingAdviceLineItems(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReceivingAdviceService.GetReceivingAdviceLineItems() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(receivingAdviceLineItemsResp, tt.want) {
			t.Errorf("ReceivingAdviceService.GetReceivingAdviceLineItems() = %v, want %v", receivingAdviceLineItemsResp, tt.want)
		}
		assert.NotNil(t, receivingAdviceLineItemsResp)
		receivingAdviceLineItemResult := receivingAdviceLineItemsResp.ReceivingAdviceLineItems[0]
		assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.LineItemNumber, uint32(1), "they should be equal")
		assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.QuantityAccepted, float64(38), "they should be equal")
		assert.Equal(t, receivingAdviceLineItemResult.ReceivingAdviceLineItemD.QuantityReceived, float64(48), "they should be equal")
		assert.NotNil(t, receivingAdviceLineItemResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, receivingAdviceLineItemResult.CrUpdTime.UpdatedAt)
	}
}

func GetReceivingAdviceLineItem(id uint32, uuid4 []byte, idS string, lineItemNumber uint32, parentLineItemNumber uint32, quantityAccepted float64, qaMeasurementUnitCode string, qaCodeListVersion string, quantityDespatched float64, qdMeasurementUnitCode string, qdCodeListVersion string, quantityReceived float64, qrMeasurementUnitCode string, qrCodeListVersion string, transactionalTradeItem uint32, ecomConsignmentIdentification uint32, contract uint32, customerReference uint32, deliveryNote uint32, despatchAdvice uint32, productCertification uint32, promotionalDeal uint32, purchaseConditions uint32, purchaseOrder uint32, requestedItemIdentification uint32, specification uint32, receivingAdviceId uint32, pickUpDateTimeBegin string, pickUpDateTimeEnd string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*logisticsproto.ReceivingAdviceLineItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	pickUpDateTimeBegin1, err := common.ConvertTimeToTimestamp(Layout, pickUpDateTimeBegin)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	pickUpDateTimeEnd1, err := common.ConvertTimeToTimestamp(Layout, pickUpDateTimeEnd)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	receivingAdviceLineItemD := logisticsproto.ReceivingAdviceLineItemD{}
	receivingAdviceLineItemD.Id = id
	receivingAdviceLineItemD.Uuid4 = uuid4
	receivingAdviceLineItemD.IdS = idS
	receivingAdviceLineItemD.LineItemNumber = lineItemNumber
	receivingAdviceLineItemD.ParentLineItemNumber = parentLineItemNumber
	receivingAdviceLineItemD.QuantityAccepted = quantityAccepted
	receivingAdviceLineItemD.QaMeasurementUnitCode = qaMeasurementUnitCode
	receivingAdviceLineItemD.QaCodeListVersion = qaCodeListVersion
	receivingAdviceLineItemD.QuantityDespatched = quantityDespatched
	receivingAdviceLineItemD.QdMeasurementUnitCode = qdMeasurementUnitCode
	receivingAdviceLineItemD.QdCodeListVersion = qdCodeListVersion
	receivingAdviceLineItemD.QuantityReceived = quantityReceived
	receivingAdviceLineItemD.QrMeasurementUnitCode = qrMeasurementUnitCode
	receivingAdviceLineItemD.QrCodeListVersion = qrCodeListVersion
	receivingAdviceLineItemD.TransactionalTradeItem = transactionalTradeItem
	receivingAdviceLineItemD.EcomConsignmentIdentification = ecomConsignmentIdentification
	receivingAdviceLineItemD.Contract = contract
	receivingAdviceLineItemD.CustomerReference = customerReference
	receivingAdviceLineItemD.DeliveryNote = deliveryNote
	receivingAdviceLineItemD.DespatchAdvice = despatchAdvice
	receivingAdviceLineItemD.ProductCertification = productCertification
	receivingAdviceLineItemD.PromotionalDeal = promotionalDeal
	receivingAdviceLineItemD.PurchaseConditions = purchaseConditions
	receivingAdviceLineItemD.PurchaseOrder = purchaseOrder
	receivingAdviceLineItemD.RequestedItemIdentification = requestedItemIdentification
	receivingAdviceLineItemD.Specification = specification
	receivingAdviceLineItemD.ReceivingAdviceId = receivingAdviceId

	receivingAdviceLineItemT := logisticsproto.ReceivingAdviceLineItemT{}
	receivingAdviceLineItemT.PickUpDateTimeBegin = pickUpDateTimeBegin1
	receivingAdviceLineItemT.PickUpDateTimeEnd = pickUpDateTimeEnd1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	receivingAdviceLineItemResult := logisticsproto.ReceivingAdviceLineItem{ReceivingAdviceLineItemD: &receivingAdviceLineItemD, ReceivingAdviceLineItemT: &receivingAdviceLineItemT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &receivingAdviceLineItemResult, nil
}
