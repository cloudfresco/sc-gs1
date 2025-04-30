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

func TestDespatchAdviceService_CreateDespatchAdviceLineItem(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchAdviceLineItem := logisticsproto.CreateDespatchAdviceLineItemRequest{}
	despatchAdviceLineItem.ActualProcessedQuantity = float64(48)
	despatchAdviceLineItem.MeasurementUnitCode = ""
	despatchAdviceLineItem.CodeListVersion = ""
	despatchAdviceLineItem.CountryOfLastProcessing = ""
	despatchAdviceLineItem.CountryOfOrigin = ""
	despatchAdviceLineItem.DespatchedQuantity = float64(48)
	despatchAdviceLineItem.DqMeasurementUnitCode = ""
	despatchAdviceLineItem.DqCodeListVersion = ""
	despatchAdviceLineItem.DutyFeeTaxLiability = ""
	despatchAdviceLineItem.Extension = ""
	despatchAdviceLineItem.FreeGoodsQuantity = float64(48)
	despatchAdviceLineItem.FgqMeasurementUnitCode = ""
	despatchAdviceLineItem.FgqCodeListVersion = ""
	despatchAdviceLineItem.HandlingInstructionCode = ""
	despatchAdviceLineItem.HasItemBeenScannedAtPos = ""
	despatchAdviceLineItem.InventoryStatusType = ""
	despatchAdviceLineItem.LineItemNumber = uint32(1)
	despatchAdviceLineItem.ParentLineItemNumber = uint32(0)
	despatchAdviceLineItem.RequestedQuantity = float64(48)
	despatchAdviceLineItem.RqMeasurementUnitCode = ""
	despatchAdviceLineItem.RqCodeListVersion = ""
	despatchAdviceLineItem.Contract = uint32(0)
	despatchAdviceLineItem.CouponClearingHouse = uint32(0)
	despatchAdviceLineItem.Customer = uint32(0)
	despatchAdviceLineItem.CustomerDocumentReference = uint32(0)
	despatchAdviceLineItem.CustomerReference = uint32(0)
	despatchAdviceLineItem.DeliveryNote = uint32(0)
	despatchAdviceLineItem.ItemOwner = uint32(0)
	despatchAdviceLineItem.OriginalSupplier = uint32(0)
	despatchAdviceLineItem.ProductCertification = uint32(0)
	despatchAdviceLineItem.PromotionalDeal = uint32(0)
	despatchAdviceLineItem.PurchaseConditions = uint32(0)
	despatchAdviceLineItem.PurchaseOrder = uint32(0)
	despatchAdviceLineItem.ReferencedConsignment = uint32(0)
	despatchAdviceLineItem.RequestedItemIdentification = uint32(0)
	despatchAdviceLineItem.Specification = uint32(0)
	despatchAdviceLineItem.DespatchAdviceId = uint32(1)
	despatchAdviceLineItem.FirstInFirstOutDateTime = "04/11/2011"
	despatchAdviceLineItem.PickUpDateTime = "04/11/2011"
	despatchAdviceLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	despatchAdviceLineItem.UserEmail = "sprov300@gmail.com"
	despatchAdviceLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *logisticsproto.CreateDespatchAdviceLineItemRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.CreateDespatchAdviceLineItemResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx: ctx,
				in:  &despatchAdviceLineItem,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdviceLineItemResponse, err := tt.das.CreateDespatchAdviceLineItem(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.CreateDespatchAdviceLineItem() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		despatchAdviceLineItemResult := despatchAdviceLineItemResponse.DespatchAdviceLineItem
		assert.NotNil(t, despatchAdviceLineItemResult)
		assert.NotNil(t, despatchAdviceLineItemResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, despatchAdviceLineItemResult.CrUpdTime.UpdatedAt)
	}
}

func TestDespatchAdviceService_GetDespatchAdviceLineItems(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchAdviceLineItem1, err := GetDespatchAdviceLineItem(uint32(1), []byte{229, 211, 24, 93, 203, 145, 64, 161, 189, 255, 54, 22, 235, 248, 98, 64}, "e5d3185d-cb91-40a1-bdff-3616ebf86240", float64(48), "", "", "", "", float64(48), "", "", "", "", float64(48), "", "", "", "", "", uint32(1), uint32(0), float64(48), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), "2011-04-11T15:00:00Z", "2011-04-11T15:00:00Z", "2011-04-11T15:00:00Z", "2011-04-11T15:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
	}

	despatchAdviceLineItems := []*logisticsproto.DespatchAdviceLineItem{}
	despatchAdviceLineItems = append(despatchAdviceLineItems, despatchAdviceLineItem1)

	despatchAdviceLineItemsResponse := logisticsproto.GetDespatchAdviceLineItemsResponse{}
	despatchAdviceLineItemsResponse.DespatchAdviceLineItems = despatchAdviceLineItems

	form := logisticsproto.GetDespatchAdviceLineItemsRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "43a938d2-4550-4dc1-80d8-00f5fd134cca"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *logisticsproto.GetDespatchAdviceLineItemsRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.GetDespatchAdviceLineItemsResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &despatchAdviceLineItemsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdviceLineItemsResp, err := tt.das.GetDespatchAdviceLineItems(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.GetDespatchAdviceLineItems() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(despatchAdviceLineItemsResp, tt.want) {
			t.Errorf("DespatchAdviceService.GetDespatchAdviceLineItems() = %v, want %v", despatchAdviceLineItemsResp, tt.want)
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
}

func GetDespatchAdviceLineItem(id uint32, uuid4 []byte, idS string, actualProcessedQuantity float64, measurementUnitCode string, codeListVersion string, countryOfLastProcessing string, countryOfOrigin string, despatchedQuantity float64, dqMeasurementUnitCode string, dqCodeListVersion string, dutyFeeTaxLiability string, extension string, freeGoodsQuantity float64, fgqMeasurementUnitCode string, fgqCodeListVersion string, handlingInstructionCode string, hasItemBeenScannedAtPos string, inventoryStatusType string, lineItemNumber uint32, parentLineItemNumber uint32, requestedQuantity float64, rqMeasurementUnitCode string, rqCodeListVersion string, contract uint32, couponClearingHouse uint32, customer uint32, customerDocumentReference uint32, customerReference uint32, deliveryNote uint32, itemOwner uint32, originalSupplier uint32, productCertification uint32, promotionalDeal uint32, purchaseConditions uint32, purchaseOrder uint32, referencedConsignment uint32, requestedItemIdentification uint32, specification uint32, despatchAdviceId uint32, firstInFirstOutDateTime string, pickUpDateTime string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*logisticsproto.DespatchAdviceLineItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	firstInFirstOutDateTime1, err := common.ConvertTimeToTimestamp(Layout, firstInFirstOutDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	pickUpDateTime1, err := common.ConvertTimeToTimestamp(Layout, pickUpDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	despatchAdviceLineItemD := logisticsproto.DespatchAdviceLineItemD{}
	despatchAdviceLineItemD.Id = id
	despatchAdviceLineItemD.Uuid4 = uuid4
	despatchAdviceLineItemD.IdS = idS
	despatchAdviceLineItemD.ActualProcessedQuantity = actualProcessedQuantity
	despatchAdviceLineItemD.MeasurementUnitCode = measurementUnitCode
	despatchAdviceLineItemD.CodeListVersion = codeListVersion
	despatchAdviceLineItemD.CountryOfLastProcessing = countryOfLastProcessing
	despatchAdviceLineItemD.CountryOfOrigin = countryOfOrigin
	despatchAdviceLineItemD.DespatchedQuantity = despatchedQuantity
	despatchAdviceLineItemD.DqMeasurementUnitCode = dqMeasurementUnitCode
	despatchAdviceLineItemD.DqCodeListVersion = dqCodeListVersion
	despatchAdviceLineItemD.DutyFeeTaxLiability = dutyFeeTaxLiability
	despatchAdviceLineItemD.Extension = extension
	despatchAdviceLineItemD.FreeGoodsQuantity = freeGoodsQuantity
	despatchAdviceLineItemD.FgqMeasurementUnitCode = fgqMeasurementUnitCode
	despatchAdviceLineItemD.FgqCodeListVersion = fgqCodeListVersion
	despatchAdviceLineItemD.HandlingInstructionCode = handlingInstructionCode
	despatchAdviceLineItemD.HasItemBeenScannedAtPos = hasItemBeenScannedAtPos
	despatchAdviceLineItemD.InventoryStatusType = inventoryStatusType
	despatchAdviceLineItemD.LineItemNumber = lineItemNumber
	despatchAdviceLineItemD.ParentLineItemNumber = parentLineItemNumber
	despatchAdviceLineItemD.RequestedQuantity = requestedQuantity
	despatchAdviceLineItemD.RqMeasurementUnitCode = rqMeasurementUnitCode
	despatchAdviceLineItemD.RqCodeListVersion = rqCodeListVersion
	despatchAdviceLineItemD.Contract = contract
	despatchAdviceLineItemD.CouponClearingHouse = couponClearingHouse
	despatchAdviceLineItemD.Customer = customer
	despatchAdviceLineItemD.CustomerDocumentReference = customerDocumentReference
	despatchAdviceLineItemD.CustomerReference = customerReference
	despatchAdviceLineItemD.DeliveryNote = deliveryNote
	despatchAdviceLineItemD.ItemOwner = itemOwner
	despatchAdviceLineItemD.OriginalSupplier = originalSupplier
	despatchAdviceLineItemD.ProductCertification = productCertification
	despatchAdviceLineItemD.PromotionalDeal = promotionalDeal
	despatchAdviceLineItemD.PurchaseConditions = purchaseConditions
	despatchAdviceLineItemD.PurchaseOrder = purchaseOrder
	despatchAdviceLineItemD.ReferencedConsignment = referencedConsignment
	despatchAdviceLineItemD.RequestedItemIdentification = requestedItemIdentification
	despatchAdviceLineItemD.Specification = specification
	despatchAdviceLineItemD.DespatchAdviceId = despatchAdviceId

	despatchAdviceLineItemT := logisticsproto.DespatchAdviceLineItemT{}
	despatchAdviceLineItemT.FirstInFirstOutDateTime = firstInFirstOutDateTime1
	despatchAdviceLineItemT.PickUpDateTime = pickUpDateTime1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	despatchAdviceLineItemResult := logisticsproto.DespatchAdviceLineItem{DespatchAdviceLineItemD: &despatchAdviceLineItemD, DespatchAdviceLineItemT: &despatchAdviceLineItemT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &despatchAdviceLineItemResult, nil
}
