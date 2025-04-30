package invoiceservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestInvoiceService_CreateInvoiceLineItem(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	invoiceService := NewInvoiceService(log, dbService, redisService, userServiceClient)
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
	invoiceLineItem.InvoiceId = uint32(1)
	invoiceLineItem.TransferOfOwnershipDate = "04/11/2011"
	invoiceLineItem.ActualDeliveryDate = "04/11/2011"
	invoiceLineItem.ServicetimePeriodLineLevelBegin = "04/11/2011"
	invoiceLineItem.ServicetimePeriodLineLevelEnd = "04/11/2011"
	invoiceLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	invoiceLineItem.UserEmail = "sprov300@gmail.com"
	invoiceLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *invoiceproto.CreateInvoiceLineItemRequest
	}
	tests := []struct {
		invs    *InvoiceService
		args    args
		wantErr bool
	}{
		{
			invs: invoiceService,
			args: args{
				ctx: ctx,
				in:  &invoiceLineItem,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		invoiceLineItemResponse, err := tt.invs.CreateInvoiceLineItem(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("InvoiceService.CreateInvoiceLineItem() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		invoiceLineItemResult := invoiceLineItemResponse.InvoiceLineItem
		assert.NotNil(t, invoiceLineItemResult)
		assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.CreditReason, "DAMAGED_GOOD", "they should be equal")
		assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.AiacCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, invoiceLineItemResult.InvoiceLineItemD.CreditLineIndicator, "Credit", "they should be equal")
		assert.NotNil(t, invoiceLineItemResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, invoiceLineItemResult.CrUpdTime.UpdatedAt)
	}
}

func TestInvoiceService_GetInvoiceLineItems(t *testing.T) {
	var err error
	err = test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()
	invoiceService := NewInvoiceService(log, dbService, redisService, userServiceClient)
	invoiceLineItem, err := GetInvoiceLineItem(uint32(3), []byte{67, 160, 154, 58, 126, 148, 70, 68, 170, 98, 243, 219, 119, 46, 203, 84}, "43a09a3a-7e94-4644-aa62-f3db772ecb54", float64(0), "", "", float64(100), "", "EUR", "Credit", "DAMAGED_GOOD", float64(10), "", "", false, "", float64(0), "", "", float64(10), "", "", float64(0), "", "", float64(15), "", "EUR", float64(10), "", "", float64(0), "", "", uint32(1), "", "", uint32(0), float64(0), "", "", float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), "2011-04-11T10:04:26Z", "2011-05-11T10:04:26Z", "2011-04-11T10:04:26Z", "2011-04-12T10:04:26Z", "2011-04-11T10:04:26Z", "2011-04-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
	}
	invoiceLineItems := []*invoiceproto.InvoiceLineItem{}
	invoiceLineItems = append(invoiceLineItems, invoiceLineItem)

	invoiceLineItemsResponse := invoiceproto.GetInvoiceLineItemsResponse{}
	invoiceLineItemsResponse.InvoiceLineItems = invoiceLineItems

	form := invoiceproto.GetInvoiceLineItemsRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "7b71cc13-cc69-44a8-a2b4-4a1cbad97da1"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *invoiceproto.GetInvoiceLineItemsRequest
	}
	tests := []struct {
		invs    *InvoiceService
		args    args
		want    *invoiceproto.GetInvoiceLineItemsResponse
		wantErr bool
	}{
		{
			invs: invoiceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &invoiceLineItemsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		invoiceLineItemsResp, err := tt.invs.GetInvoiceLineItems(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("InvoiceService.GetInvoiceLineItems() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(invoiceLineItemsResp, tt.want) {
			t.Errorf("InvoiceService.GetInvoiceLineItems() = %v, want %v", invoiceLineItemsResp, tt.want)
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
}

func GetInvoiceLineItem(id uint32, uuid4 []byte, idS string, amountExclusiveAllowancesCharges float64, aeacCodeListVersion string, aeacCurrencyCode string, amountInclusiveAllowancesCharges float64, aiacCodeListVersion string, aiacCurrencyCode string, creditLineIndicator string, creditReason string, deliveredQuantity float64, dqMeasurementUnitCode string, dqCodeListVersion string, excludedFromPaymentDiscountIndicator bool, extension string, freeGoodsQuantity float64, fgqMeasurementUnitCode string, fgqCodeListVersion string, invoicedQuantity float64, IqMeasurementUnitCode string, IqCodeListVersion string, itemPriceBaseQuantity float64, ipbqMeasurementUnitCode string, ipbqCodeListVersion string, itemPriceExclusiveAllowancesCharges float64, iPEACCodeListVersion string, iPEACCurrencyCode string, itemPriceInclusiveAllowancesCharges float64, ipiacCodeListVersion string, ipiacCurrencyCode string, legallyFixedRetailPrice float64, lfrpCodeListVersion string, lfrpCurrencyCode string, lineItemNumber uint32, marginSchemeInformation string, owenrshipPriorToPayment string, parentLineItemNumber uint32, recommendedRetailPrice float64, rrpCodeListVersion string, rrpCurrencyCode string, retailPriceExcludingExcise float64, rpeeCodeListVersion string, rpeeCurrencyCode string, totalOrderedQuantity float64, toqMeasurementUnitCode string, toqCodeListVersion string, consumptionReport uint32, contract uint32, deliveryNote uint32, despatchAdvice uint32, energyQuantity uint32, inventoryLocationFrom uint32, inventoryLocationTo uint32, promotionalDeal uint32, purchaseConditions uint32, purchaseOrder uint32, receivingAdvice uint32, returnableAssetIdentification uint32, salesOrder uint32, shipFrom uint32, shipTo uint32, tradeAgreement uint32, invoiceId uint32, transferOfOwnershipDate string, actualDeliveryDate string, servicetimePeriodLineLevelBegin string, servicetimePeriodLineLevelEnd string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*invoiceproto.InvoiceLineItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	transferOfOwnershipDate1, err := common.ConvertTimeToTimestamp(Layout, transferOfOwnershipDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	actualDeliveryDate1, err := common.ConvertTimeToTimestamp(Layout, actualDeliveryDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	servicetimePeriodLineLevelBegin1, err := common.ConvertTimeToTimestamp(Layout, servicetimePeriodLineLevelBegin)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	servicetimePeriodLineLevelEnd1, err := common.ConvertTimeToTimestamp(Layout, servicetimePeriodLineLevelEnd)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	invoiceLineItemD := invoiceproto.InvoiceLineItemD{}
	invoiceLineItemD.Id = id
	invoiceLineItemD.Uuid4 = uuid4
	invoiceLineItemD.IdS = idS
	invoiceLineItemD.AmountExclusiveAllowancesCharges = amountExclusiveAllowancesCharges
	invoiceLineItemD.AeacCodeListVersion = aeacCodeListVersion
	invoiceLineItemD.AeacCurrencyCode = aeacCurrencyCode
	invoiceLineItemD.AmountInclusiveAllowancesCharges = amountInclusiveAllowancesCharges
	invoiceLineItemD.AiacCodeListVersion = aiacCodeListVersion
	invoiceLineItemD.AiacCurrencyCode = aiacCurrencyCode
	invoiceLineItemD.CreditLineIndicator = creditLineIndicator
	invoiceLineItemD.CreditReason = creditReason
	invoiceLineItemD.DeliveredQuantity = deliveredQuantity
	invoiceLineItemD.DqMeasurementUnitCode = dqMeasurementUnitCode
	invoiceLineItemD.DqCodeListVersion = dqCodeListVersion
	invoiceLineItemD.ExcludedFromPaymentDiscountIndicator = excludedFromPaymentDiscountIndicator
	invoiceLineItemD.Extension = extension
	invoiceLineItemD.FreeGoodsQuantity = freeGoodsQuantity
	invoiceLineItemD.FgqMeasurementUnitCode = fgqMeasurementUnitCode
	invoiceLineItemD.FgqCodeListVersion = fgqCodeListVersion
	invoiceLineItemD.InvoicedQuantity = invoicedQuantity
	invoiceLineItemD.IqMeasurementUnitCode = IqMeasurementUnitCode
	invoiceLineItemD.IqCodeListVersion = IqCodeListVersion
	invoiceLineItemD.ItemPriceBaseQuantity = itemPriceBaseQuantity
	invoiceLineItemD.IpbqMeasurementUnitCode = ipbqMeasurementUnitCode
	invoiceLineItemD.IpbqCodeListVersion = ipbqCodeListVersion
	invoiceLineItemD.ItemPriceExclusiveAllowancesCharges = itemPriceExclusiveAllowancesCharges
	invoiceLineItemD.IpeacCodeListVersion = iPEACCodeListVersion
	invoiceLineItemD.IpeacCurrencyCode = iPEACCurrencyCode
	invoiceLineItemD.ItemPriceInclusiveAllowancesCharges = itemPriceInclusiveAllowancesCharges
	invoiceLineItemD.IpiacCodeListVersion = ipiacCodeListVersion
	invoiceLineItemD.IpiacCurrencyCode = ipiacCurrencyCode
	invoiceLineItemD.LegallyFixedRetailPrice = legallyFixedRetailPrice
	invoiceLineItemD.LfrpCodeListVersion = lfrpCodeListVersion
	invoiceLineItemD.LfrpCurrencyCode = lfrpCurrencyCode
	invoiceLineItemD.LineItemNumber = lineItemNumber
	invoiceLineItemD.MarginSchemeInformation = marginSchemeInformation
	invoiceLineItemD.OwenrshipPriorToPayment = owenrshipPriorToPayment
	invoiceLineItemD.ParentLineItemNumber = parentLineItemNumber
	invoiceLineItemD.RecommendedRetailPrice = recommendedRetailPrice
	invoiceLineItemD.RrpCodeListVersion = rrpCodeListVersion
	invoiceLineItemD.RrpCurrencyCode = rrpCurrencyCode
	invoiceLineItemD.RetailPriceExcludingExcise = retailPriceExcludingExcise
	invoiceLineItemD.RpeeCodeListVersion = rpeeCodeListVersion
	invoiceLineItemD.RpeeCurrencyCode = rpeeCurrencyCode
	invoiceLineItemD.TotalOrderedQuantity = totalOrderedQuantity
	invoiceLineItemD.ToqMeasurementUnitCode = toqMeasurementUnitCode
	invoiceLineItemD.ToqCodeListVersion = toqCodeListVersion
	invoiceLineItemD.ConsumptionReport = consumptionReport
	invoiceLineItemD.Contract = contract
	invoiceLineItemD.DeliveryNote = deliveryNote
	invoiceLineItemD.DespatchAdvice = despatchAdvice
	invoiceLineItemD.EnergyQuantity = energyQuantity
	invoiceLineItemD.InventoryLocationFrom = inventoryLocationFrom
	invoiceLineItemD.InventoryLocationTo = inventoryLocationTo
	invoiceLineItemD.PromotionalDeal = promotionalDeal
	invoiceLineItemD.PurchaseConditions = purchaseConditions
	invoiceLineItemD.PurchaseOrder = purchaseOrder
	invoiceLineItemD.ReceivingAdvice = receivingAdvice
	invoiceLineItemD.ReturnableAssetIdentification = returnableAssetIdentification
	invoiceLineItemD.SalesOrder = salesOrder
	invoiceLineItemD.ShipFrom = shipFrom
	invoiceLineItemD.ShipTo = shipTo
	invoiceLineItemD.TradeAgreement = tradeAgreement
	invoiceLineItemD.InvoiceId = invoiceId

	invoiceLineItemT := invoiceproto.InvoiceLineItemT{}
	invoiceLineItemT.TransferOfOwnershipDate = transferOfOwnershipDate1
	invoiceLineItemT.ActualDeliveryDate = actualDeliveryDate1
	invoiceLineItemT.ServicetimePeriodLineLevelBegin = servicetimePeriodLineLevelBegin1
	invoiceLineItemT.ServicetimePeriodLineLevelEnd = servicetimePeriodLineLevelEnd1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	invoiceLineItemResult := invoiceproto.InvoiceLineItem{InvoiceLineItemD: &invoiceLineItemD, InvoiceLineItemT: &invoiceLineItemT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &invoiceLineItemResult, nil
}
