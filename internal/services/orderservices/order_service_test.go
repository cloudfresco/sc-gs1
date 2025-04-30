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

func TestOrderService_CreateOrder(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)
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
	order.UserId = "auth0|673ee1a719dd4000cd5a3832"
	order.UserEmail = "sprov300@gmail.com"
	order.RequestId = "bks1m1g91jau4nkks2f0"

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

	orderLineItems := []*orderproto.CreateOrderLineItemRequest{}
	orderLineItems = append(orderLineItems, &orderLineItem)
	order.OrderLineItems = orderLineItems

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		want    *orderproto.CreateOrderResponse
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx: ctx,
				in:  &order,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		orderResp, err := tt.o.CreateOrder(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		assert.NotNil(t, orderResp)
		orderResult := orderResp.Order
		assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
		assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")
	}
}

func TestOrderService_GetOrders(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)
	orders := []*orderproto.Order{}

	order1, err := GetOrder(uint32(1), []byte{0, 232, 112, 16, 155, 84, 79, 179, 180, 68, 0, 40, 13, 164, 77, 137}, "00e87010-9b54-4fb3-b444-00280da44d89", true, false, "", "", "PARTIAL_DELIVERY_ALLOWED", "", "220", float64(0), "", "", float64(0), "", "", float64(0), "", "", uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), "ASP0002NET", uint32(1), "56895632", "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	order2, err := GetOrder(uint32(2), []byte{35, 224, 153, 89, 35, 178, 71, 10, 129, 186, 15, 184, 187, 34, 165, 98}, "23e09959-23b2-470a-81ba-0fb8bb22a562", true, false, "", "", "PARTIAL_DELIVERY_ALLOWED", "", "402", float64(12675), "", "", float64(0), "", "", float64(2661.75), "", "", uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), "ASP0002NET", uint32(1), "56895632", "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orders = append(orders, order2, order1)

	form := orderproto.GetOrdersRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	ordersResponse := orderproto.GetOrdersResponse{Orders: orders, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *orderproto.GetOrdersRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		want    *orderproto.GetOrdersResponse
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &ordersResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		orderResp, err := tt.o.GetOrders(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.GetOrders() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderResp, tt.want) {
			t.Errorf("OrderService.GetOrders() = %v, want %v", orderResp, tt.want)
		}
		assert.NotNil(t, orderResp)
		orderResult := orderResp.Orders[0]
		assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
		assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")
	}
}

func TestOrderService_GetOrder(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)
	order, err := GetOrder(uint32(2), []byte{35, 224, 153, 89, 35, 178, 71, 10, 129, 186, 15, 184, 187, 34, 165, 98}, "23e09959-23b2-470a-81ba-0fb8bb22a562", true, false, "", "", "PARTIAL_DELIVERY_ALLOWED", "", "402", float64(12675), "", "", float64(0), "", "", float64(2661.75), "", "", uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), "ASP0002NET", uint32(1), "56895632", "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResponse := orderproto.GetOrderResponse{}
	orderResponse.Order = order

	form := orderproto.GetOrderRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "23e09959-23b2-470a-81ba-0fb8bb22a562"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *orderproto.GetOrderRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		want    *orderproto.GetOrderResponse
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &orderResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderResp, err := tt.o.GetOrder(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.GetOrder() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderResp, tt.want) {
			t.Errorf("OrderService.GetOrder() = %v, want %v", orderResp, tt.want)
		}
		assert.NotNil(t, orderResp)
		orderResult := orderResp.Order
		assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
		assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")

	}
}

func TestOrderService_GetOrderByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)
	order, err := GetOrder(uint32(2), []byte{35, 224, 153, 89, 35, 178, 71, 10, 129, 186, 15, 184, 187, 34, 165, 98}, "23e09959-23b2-470a-81ba-0fb8bb22a562", true, false, "", "", "PARTIAL_DELIVERY_ALLOWED", "", "402", float64(12675), "", "", float64(0), "", "", float64(2661.75), "", "", uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), "ASP0002NET", uint32(1), "56895632", "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResponse := orderproto.GetOrderByPkResponse{}
	orderResponse.Order = order

	form := orderproto.GetOrderByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(2)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *orderproto.GetOrderByPkRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		want    *orderproto.GetOrderByPkResponse
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &orderResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderResp, err := tt.o.GetOrderByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.GetOrderByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderResp, tt.want) {
			t.Errorf("OrderService.GetOrderByPk() = %v, want %v", orderResp, tt.want)
		}
		assert.NotNil(t, orderResp)
		orderResult := orderResp.Order
		assert.Equal(t, orderResult.OrderD.OrderTypeCode, "402", "they should be equal")
		assert.Equal(t, orderResult.OrderD.OrderInstructionCode, "PARTIAL_DELIVERY_ALLOWED", "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalMonetaryAmountExcludingTaxes, float64(12675), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TotalTaxAmount, float64(2661.75), "they should be equal")
		assert.Equal(t, orderResult.OrderD.TradeAgreement, "56895632", "they should be equal")
	}
}

func GetOrder(id uint32, uuid4 []byte, idS string, isApplicationReceiptAcknowledgementRequired bool, isOrderFreeOfExciseTaxDuty bool, orderChangeReasonCode string, orderEntryType string, orderInstructionCode string, orderPriority string, orderTypeCode string, totalMonetaryAmountExcludingTaxes float64, tmaetCodeListVersion string, tmaetCurrencyCode string, totalMonetaryAmountIncludingTaxes float64, tmaitCodeListVersion string, tmaitCurrencyCode string, totalTaxAmount float64, ttaCodeListVersion string, ttaCurrencyCode string, billTo uint32, buyer uint32, contract uint32, customerDocumentReference uint32, customsBroker uint32, orderIdentification uint32, pickupFrom uint32, promotionalDeal uint32, quoteNumber string, seller uint32, tradeAgreement string, deliveryDateAccordingToSchedule string, latestDeliveryDate string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*orderproto.Order, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	deliveryDateAccordingToSchedule1, err := common.ConvertTimeToTimestamp(Layout, deliveryDateAccordingToSchedule)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	latestDeliveryDate1, err := common.ConvertTimeToTimestamp(Layout, latestDeliveryDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	orderD := orderproto.OrderD{}
	orderD.Id = id
	orderD.Uuid4 = uuid4
	orderD.IdS = idS
	orderD.IsApplicationReceiptAcknowledgementRequired = isApplicationReceiptAcknowledgementRequired
	orderD.IsOrderFreeOfExciseTaxDuty = isOrderFreeOfExciseTaxDuty
	orderD.OrderChangeReasonCode = orderChangeReasonCode
	orderD.OrderEntryType = orderEntryType
	orderD.OrderInstructionCode = orderInstructionCode
	orderD.OrderPriority = orderPriority
	orderD.OrderTypeCode = orderTypeCode
	orderD.TotalMonetaryAmountExcludingTaxes = totalMonetaryAmountExcludingTaxes
	orderD.TmaetCodeListVersion = tmaetCodeListVersion
	orderD.TmaetCurrencyCode = tmaetCurrencyCode
	orderD.TotalMonetaryAmountIncludingTaxes = totalMonetaryAmountIncludingTaxes
	orderD.TmaitCodeListVersion = tmaitCodeListVersion
	orderD.TmaitCurrencyCode = tmaitCurrencyCode
	orderD.TotalTaxAmount = totalTaxAmount
	orderD.TtaCodeListVersion = ttaCodeListVersion
	orderD.TtaCurrencyCode = ttaCurrencyCode
	orderD.BillTo = billTo
	orderD.Buyer = buyer
	orderD.Contract = contract
	orderD.CustomerDocumentReference = customerDocumentReference
	orderD.CustomsBroker = customsBroker
	orderD.OrderIdentification = orderIdentification
	orderD.PickupFrom = pickupFrom
	orderD.PromotionalDeal = promotionalDeal
	orderD.QuoteNumber = quoteNumber
	orderD.Seller = seller
	orderD.TradeAgreement = tradeAgreement

	orderT := orderproto.OrderT{}
	orderT.DeliveryDateAccordingToSchedule = deliveryDateAccordingToSchedule1
	orderT.LatestDeliveryDate = latestDeliveryDate1

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	order := orderproto.Order{OrderD: &orderD, OrderT: &orderT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &order, nil
}
