package orderservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestOrderResponseService_CreateOrderResponseLineItem(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)

	orderResponseLineItem := orderproto.CreateOrderResponseLineItemRequest{}
	orderResponseLineItem.ConfirmedQuantity = float64(24)
	orderResponseLineItem.CqMeasurementUnitCode = ""
	orderResponseLineItem.CqCodeListVersion = ""
	orderResponseLineItem.LineItemActionCode = ""
	orderResponseLineItem.LineItemChangeIndicator = "MODIFIED"
	orderResponseLineItem.LineItemNumber = uint32(2)
	orderResponseLineItem.MonetaryAmountExcludingTaxes = float64(0)
	orderResponseLineItem.MaetCodeListVersion = ""
	orderResponseLineItem.MaetCurrencyCode = ""
	orderResponseLineItem.MonetaryAmountIncludingTaxes = float64(0)
	orderResponseLineItem.MaitCodeListVersion = ""
	orderResponseLineItem.MaitCurrencyCode = ""
	orderResponseLineItem.NetAmount = float64(0)
	orderResponseLineItem.NaCodeListVersion = ""
	orderResponseLineItem.NaCurrencyCode = ""
	orderResponseLineItem.NetPrice = float64(0)
	orderResponseLineItem.NpCodeListVersion = ""
	orderResponseLineItem.NpCurrencyCode = ""
	orderResponseLineItem.OrderResponseReasonCode = "PRODUCT_OUT_OF_STOCK"
	orderResponseLineItem.OriginalOrderLineItemNumber = uint32(2)
	orderResponseLineItem.ParentLineItemNumber = uint32(0)
	orderResponseLineItem.DeliveryDateTime = "02/19/2023"
	orderResponseLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	orderResponseLineItem.UserEmail = "sprov300@gmail.com"
	orderResponseLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderResponseLineItemRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.CreateOrderResponseLineItemResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx: ctx,
				in:  &orderResponseLineItem,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderResponseLineItemResp, err := tt.o.CreateOrderResponseLineItem(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.CreateOrderResponseLineItem() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		orderResponseLineItemResult := orderResponseLineItemResp.OrderResponseLineItem
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.LineItemNumber, uint32(2), "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.ConfirmedQuantity, float64(24), "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.LineItemChangeIndicator, "MODIFIED", "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.OrderResponseReasonCode, "PRODUCT_OUT_OF_STOCK", "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.OriginalOrderLineItemNumber, uint32(2), "they should be equal")
	}
}

func TestOrderResponseService_GetOrderResponseLineItems(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)

	orderResponseLineItem1, err := GetOrderResponseLineItem(uint32(1), []byte{52, 57, 14, 231, 202, 5, 71, 220, 151, 64, 57, 41, 81, 29, 187, 239}, "34390ee7-ca05-47dc-9740-3929511dbbef", float64(48), "", "", "", "ACCEPTED", uint32(1), float64(0), "", "", float64(0), "", "", float64(0), "", "", float64(0), "", "", "", uint32(1), uint32(0), uint32(3), "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResponseLineItem2, err := GetOrderResponseLineItem(uint32(2), []byte{55, 46, 23, 120, 232, 83, 71, 69, 148, 112, 22, 157, 0, 147, 83, 21}, "372e1778-e853-4745-9470-169d00935315", float64(24), "", "", "", "MODIFIED", uint32(2), float64(0), "", "", float64(0), "", "", float64(0), "", "", float64(0), "", "", "PRODUCT_OUT_OF_STOCK", uint32(2), uint32(0), uint32(3), "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResponseLineItems := []*orderproto.OrderResponseLineItem{}

	orderResponseLineItems = append(orderResponseLineItems, orderResponseLineItem1, orderResponseLineItem2)

	orderRespLineItems := orderproto.GetOrderResponseLineItemsResponse{}
	orderRespLineItems.OrderResponseLineItems = orderResponseLineItems

	form := orderproto.GetOrderResponseLineItemsRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "a9cf4bc4-bf4b-4013-95d3-d3bb4c7af952"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *orderproto.GetOrderResponseLineItemsRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.GetOrderResponseLineItemsResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &orderRespLineItems,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderResponseLineItemResp, err := tt.o.GetOrderResponseLineItems(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.GetOrderResponseLineItems() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderResponseLineItemResp, tt.want) {
			t.Errorf("OrderResponseService.GetOrderResponseLineItems() = %v, want %v", orderResponseLineItemResp, tt.want)
		}
		assert.NotNil(t, orderResponseLineItemResp)
		orderResponseLineItemResult := orderResponseLineItemResp.OrderResponseLineItems[1]
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.LineItemNumber, uint32(2), "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.ConfirmedQuantity, float64(24), "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.LineItemChangeIndicator, "MODIFIED", "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.OrderResponseReasonCode, "PRODUCT_OUT_OF_STOCK", "they should be equal")
		assert.Equal(t, orderResponseLineItemResult.OrderResponseLineItemD.OriginalOrderLineItemNumber, uint32(2), "they should be equal")
	}
}

func GetOrderResponseLineItem(id uint32, uuid4 []byte, idS string, confirmedQuantity float64, cqMeasurementUnitCode string, cqCodeListVersion string, lineItemActionCode string, lineItemChangeIndicator string, lineItemNumber uint32, monetaryAmountExcludingTaxes float64, maetCodeListVersion string, maetCurrencyCode string, monetaryAmountIncludingTaxes float64, maitCodeListVersion string, maitCurrencyCode string, netAmount float64, naCodeListVersion string, naCurrencyCode string, netPrice float64, npCodeListVersion string, npCurrencyCode string, orderResponseReasonCode string, originalOrderLineItemNumber uint32, parentLineItemNumber uint32, orderResponseId uint32, deliveryDateTime string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*orderproto.OrderResponseLineItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	deliveryDateTime1, err := common.ConvertTimeToTimestamp(Layout, deliveryDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	orderResponseLineItemD := orderproto.OrderResponseLineItemD{}
	orderResponseLineItemD.Id = id
	orderResponseLineItemD.Uuid4 = uuid4
	orderResponseLineItemD.IdS = idS
	orderResponseLineItemD.ConfirmedQuantity = confirmedQuantity
	orderResponseLineItemD.CqMeasurementUnitCode = cqMeasurementUnitCode
	orderResponseLineItemD.CqCodeListVersion = cqCodeListVersion
	orderResponseLineItemD.LineItemActionCode = lineItemActionCode
	orderResponseLineItemD.LineItemChangeIndicator = lineItemChangeIndicator
	orderResponseLineItemD.LineItemNumber = lineItemNumber
	orderResponseLineItemD.MonetaryAmountExcludingTaxes = monetaryAmountExcludingTaxes
	orderResponseLineItemD.MaetCodeListVersion = maetCodeListVersion
	orderResponseLineItemD.MaetCurrencyCode = maetCurrencyCode
	orderResponseLineItemD.MonetaryAmountIncludingTaxes = monetaryAmountIncludingTaxes
	orderResponseLineItemD.MaitCodeListVersion = maitCodeListVersion
	orderResponseLineItemD.MaitCurrencyCode = maitCurrencyCode
	orderResponseLineItemD.NetAmount = netAmount
	orderResponseLineItemD.NaCodeListVersion = naCodeListVersion
	orderResponseLineItemD.NaCurrencyCode = naCurrencyCode
	orderResponseLineItemD.NetPrice = netPrice
	orderResponseLineItemD.NpCodeListVersion = npCodeListVersion
	orderResponseLineItemD.NpCurrencyCode = npCurrencyCode
	orderResponseLineItemD.OrderResponseReasonCode = orderResponseReasonCode
	orderResponseLineItemD.OriginalOrderLineItemNumber = originalOrderLineItemNumber
	orderResponseLineItemD.ParentLineItemNumber = parentLineItemNumber
	orderResponseLineItemD.OrderResponseId = orderResponseId

	orderResponseLineItemT := orderproto.OrderResponseLineItemT{}
	orderResponseLineItemT.DeliveryDateTime = deliveryDateTime1

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	orderResponseLineItem := orderproto.OrderResponseLineItem{OrderResponseLineItemD: &orderResponseLineItemD, OrderResponseLineItemT: &orderResponseLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &orderResponseLineItem, nil
}
