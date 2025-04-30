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

func TestOrderService_CreateOrderLineItem(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)

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
	orderLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	orderLineItem.UserEmail = "sprov300@gmail.com"
	orderLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderLineItemRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx: ctx,
				in:  &orderLineItem,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderLineItemResp, err := tt.o.CreateOrderLineItem(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.CreateOrderLineItem() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, orderLineItemResp)
		orderLineItemResult := orderLineItemResp.OrderLineItem
		assert.Equal(t, orderLineItemResult.OrderLineItemD.LineItemNumber, uint32(2), "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NetAmount, float64(4659), "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NetPrice, float64(194.125), "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NpCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.RqMeasurementUnitCode, "EA", "they should be equal")
	}
}

func TestOrderService_GetOrderLineItems(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)

	orderLineItems := []*orderproto.OrderLineItem{}

	orderLineItem1, err := GetOrderLineItem(uint32(3), []byte{106, 43, 202, 207, 78, 230, 71, 35, 189, 155, 180, 87, 101, 103, 12, 215}, "6a2bcacf-4ee6-4723-bd9b-b45765670cd7", "", float64(0), "", "", float64(0), "", "", "", "NOT_AMENDED", uint32(1), float64(0), "", "", float64(0), "", "", float64(0), "", "", float64(8016), "", "EUR", float64(167), "", "EUR", "", "", "", uint32(0), float64(0), float64(48), "EA", "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderLineItem2, err := GetOrderLineItem(uint32(4), []byte{17, 166, 125, 145, 208, 221, 74, 181, 182, 182, 100, 115, 217, 64, 189, 243}, "11a67d91-d0dd-4ab5-b6b6-6473d940bdf3", "", float64(0), "", "", float64(0), "", "", "", "", uint32(2), float64(0), "", "", float64(0), "", "", float64(0), "", "", float64(4659), "", "EUR", float64(194.125), "", "EUR", "", "", "", uint32(0), float64(0), float64(24), "EA", "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderLineItems = append(orderLineItems, orderLineItem1, orderLineItem2)

	orderLineItemsResponse := orderproto.GetOrderLineItemsResponse{}
	orderLineItemsResponse.OrderLineItems = orderLineItems

	form := orderproto.GetOrderLineItemsRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "23e09959-23b2-470a-81ba-0fb8bb22a562"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *orderproto.GetOrderLineItemsRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		want    *orderproto.GetOrderLineItemsResponse
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &orderLineItemsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderLineItemResp, err := tt.o.GetOrderLineItems(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.GetOrderLineItems() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderLineItemResp, tt.want) {
			t.Errorf("OrderService.GetOrderLineItems() = %v, want %v", orderLineItemResp, tt.want)
		}
		assert.NotNil(t, orderLineItemResp)
		orderLineItemResult := orderLineItemResp.OrderLineItems[1]
		assert.Equal(t, orderLineItemResult.OrderLineItemD.LineItemNumber, uint32(2), "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NetAmount, float64(4659), "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NetPrice, float64(194.125), "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.NpCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, orderLineItemResult.OrderLineItemD.RqMeasurementUnitCode, "EA", "they should be equal")

	}
}

func GetOrderLineItem(id uint32, uuid4 []byte, idS string, extension string, freeGoodsQuantity float64, fgqMeasurementUnitCode string, fgqCodeListVersion string, itemPriceBaseQuantity float64, ipbqMeasurementUnitCode string, ipbqCodeListVersion string, itemSourceCode string, lineItemActionCode string, lineItemNumber uint32, listPrice float64, lpCodeListVersion string, lpCurrencyCode string, monetaryAmountExcludingTaxes float64, maetCodeListVersion string, maetCurrencyCode string, monetaryAmountIncludingTaxes float64, maitCodeListVersion string, maitCurrencyCode string, netAmount float64, naCodeListVersion string, naCurrencyCode string, netPrice float64, npCodeListVersion string, npCurrencyCode string, orderInstructionCode string, orderLineItemInstructionCode string, orderLineItemPriority string, parentLineItemNumber uint32, recommendedRetailPrice float64, requestedQuantity float64, rqMeasurementUnitCode string, rqCodeListVersion string, returnReasonCode string, contract uint32, customerDocumentReference uint32, deliveryDateAccordingToSchedule uint32, despatchAdvice uint32, materialSpecification uint32, orderLineItemContact uint32, preferredManufacturer uint32, promotionalDeal uint32, purchaseConditions uint32, returnableAssetIdentification uint32, orderId uint32, latestDeliveryDate string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*orderproto.OrderLineItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	latestDeliveryDate1, err := common.ConvertTimeToTimestamp(Layout, latestDeliveryDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	orderLineItemD := orderproto.OrderLineItemD{}
	orderLineItemD.Id = id
	orderLineItemD.Uuid4 = uuid4
	orderLineItemD.IdS = idS
	orderLineItemD.Extension = extension
	orderLineItemD.FreeGoodsQuantity = freeGoodsQuantity
	orderLineItemD.FgqMeasurementUnitCode = fgqMeasurementUnitCode
	orderLineItemD.FgqCodeListVersion = fgqCodeListVersion
	orderLineItemD.ItemPriceBaseQuantity = itemPriceBaseQuantity
	orderLineItemD.IpbqMeasurementUnitCode = ipbqMeasurementUnitCode
	orderLineItemD.IpbqCodeListVersion = ipbqCodeListVersion
	orderLineItemD.ItemSourceCode = itemSourceCode
	orderLineItemD.LineItemActionCode = lineItemActionCode
	orderLineItemD.LineItemNumber = lineItemNumber
	orderLineItemD.ListPrice = listPrice
	orderLineItemD.LpCodeListVersion = lpCodeListVersion
	orderLineItemD.LpCurrencyCode = lpCurrencyCode
	orderLineItemD.MonetaryAmountExcludingTaxes = monetaryAmountExcludingTaxes
	orderLineItemD.MaetCodeListVersion = maetCodeListVersion
	orderLineItemD.MaetCurrencyCode = maetCurrencyCode
	orderLineItemD.MonetaryAmountIncludingTaxes = monetaryAmountIncludingTaxes
	orderLineItemD.MaitCodeListVersion = maitCodeListVersion
	orderLineItemD.MaitCurrencyCode = maitCurrencyCode
	orderLineItemD.NetAmount = netAmount
	orderLineItemD.NaCodeListVersion = naCodeListVersion
	orderLineItemD.NaCurrencyCode = naCurrencyCode
	orderLineItemD.NetPrice = netPrice
	orderLineItemD.NpCodeListVersion = npCodeListVersion
	orderLineItemD.NpCurrencyCode = npCurrencyCode
	orderLineItemD.OrderInstructionCode = orderInstructionCode
	orderLineItemD.OrderLineItemInstructionCode = orderLineItemInstructionCode
	orderLineItemD.OrderLineItemPriority = orderLineItemPriority
	orderLineItemD.ParentLineItemNumber = parentLineItemNumber
	orderLineItemD.RecommendedRetailPrice = recommendedRetailPrice
	orderLineItemD.RequestedQuantity = requestedQuantity
	orderLineItemD.RqMeasurementUnitCode = rqMeasurementUnitCode
	orderLineItemD.RqCodeListVersion = rqCodeListVersion
	orderLineItemD.ReturnReasonCode = returnReasonCode
	orderLineItemD.Contract = contract
	orderLineItemD.CustomerDocumentReference = customerDocumentReference
	orderLineItemD.DeliveryDateAccordingToSchedule = deliveryDateAccordingToSchedule
	orderLineItemD.DespatchAdvice = despatchAdvice
	orderLineItemD.MaterialSpecification = materialSpecification
	orderLineItemD.OrderLineItemContact = orderLineItemContact
	orderLineItemD.PreferredManufacturer = preferredManufacturer
	orderLineItemD.PromotionalDeal = promotionalDeal
	orderLineItemD.PurchaseConditions = purchaseConditions
	orderLineItemD.ReturnableAssetIdentification = returnableAssetIdentification
	orderLineItemD.OrderId = orderId

	orderLineItemT := orderproto.OrderLineItemT{}
	orderLineItemT.LatestDeliveryDate = latestDeliveryDate1

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	orderLineItem := orderproto.OrderLineItem{OrderLineItemD: &orderLineItemD, OrderLineItemT: &orderLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &orderLineItem, nil
}
