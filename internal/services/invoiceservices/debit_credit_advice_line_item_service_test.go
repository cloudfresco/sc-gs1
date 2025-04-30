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

func TestDebitCreditAdviceService_CreateDebitCreditAdviceLineItem(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)

	debitCreditAdviceLineItem := invoiceproto.CreateDebitCreditAdviceLineItemRequest{}
	debitCreditAdviceLineItem.AdjustmentAmount = float64(200)
	debitCreditAdviceLineItem.AaCodeListVersion = ""
	debitCreditAdviceLineItem.AaCurrencyCode = "EUR"
	debitCreditAdviceLineItem.DebitCreditIndicatorCode = "CREDIT"
	debitCreditAdviceLineItem.FinancialAdjustmentReasonCode = "18"
	debitCreditAdviceLineItem.LineItemNumber = uint32(1)
	debitCreditAdviceLineItem.ParentLineItemNumber = uint32(0)
	debitCreditAdviceLineItem.DebitCreditAdviceId = uint32(1)
	debitCreditAdviceLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	debitCreditAdviceLineItem.UserEmail = "sprov300@gmail.com"
	debitCreditAdviceLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *invoiceproto.CreateDebitCreditAdviceLineItemRequest
	}
	tests := []struct {
		ds      *DebitCreditAdviceService
		args    args
		wantErr bool
	}{
		{
			ds: debitCreditAdviceService,
			args: args{
				ctx: ctx,
				in:  &debitCreditAdviceLineItem,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceLineItemResponse, err := tt.ds.CreateDebitCreditAdviceLineItem(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.CreateDebitCreditAdviceLineItem() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		debitCreditAdviceLineItemResult := debitCreditAdviceLineItemResponse.DebitCreditAdviceLineItem
		assert.NotNil(t, debitCreditAdviceLineItemResult)
		assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.AdjustmentAmount, float64(200), "they should be equal")
		assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.AaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
		assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.FinancialAdjustmentReasonCode, "18", "they should be equal")
		assert.Equal(t, debitCreditAdviceLineItemResult.DebitCreditAdviceLineItemD.LineItemNumber, uint32(1), "they should be equal")
		assert.NotNil(t, debitCreditAdviceLineItemResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, debitCreditAdviceLineItemResult.CrUpdTime.UpdatedAt)
	}
}

func TestDebitCreditAdviceService_GetDebitCreditAdviceLineItems(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)
	debitCreditAdviceLineItem, err := GetDebitCreditAdviceLineItem(uint32(1), []byte{53, 135, 190, 241, 155, 218, 76, 162, 179, 34, 48, 56, 144, 245, 57, 160}, "3587bef1-9bda-4ca2-b322-303890f539a0", float64(100), "", "EUR", "CREDIT", "17", uint32(1), uint32(0), uint32(1), "2011-04-11T10:04:26Z", "2011-04-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
	}
	debitCreditAdviceLineItems := []*invoiceproto.DebitCreditAdviceLineItem{}
	debitCreditAdviceLineItems = append(debitCreditAdviceLineItems, debitCreditAdviceLineItem)

	debitCreditAdviceLineItemsResponse := invoiceproto.GetDebitCreditAdviceLineItemsResponse{}
	debitCreditAdviceLineItemsResponse.DebitCreditAdviceLineItems = debitCreditAdviceLineItems

	form := invoiceproto.GetDebitCreditAdviceLineItemsRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "f7a9f263-a1e1-4307-8990-6f7e4d94df08"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *invoiceproto.GetDebitCreditAdviceLineItemsRequest
	}
	tests := []struct {
		ds      *DebitCreditAdviceService
		args    args
		want    *invoiceproto.GetDebitCreditAdviceLineItemsResponse
		wantErr bool
	}{
		{
			ds: debitCreditAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &debitCreditAdviceLineItemsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceLineItemResponse, err := tt.ds.GetDebitCreditAdviceLineItems(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdviceLineItems() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(debitCreditAdviceLineItemResponse, tt.want) {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdviceLineItems() = %v, want %v", debitCreditAdviceLineItemResponse, tt.want)
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
}

func GetDebitCreditAdviceLineItem(id uint32, uuid4 []byte, idS string, adjustmentAmount float64, aaCodeListVersion string, aaCurrencyCode string, debitCreditIndicatorCode string, financialAdjustmentReasonCode string, lineItemNumber uint32, parentLineItemNumber uint32, debitCreditAdviceId uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*invoiceproto.DebitCreditAdviceLineItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	debitCreditAdviceLineItemD := invoiceproto.DebitCreditAdviceLineItemD{}
	debitCreditAdviceLineItemD.Id = id
	debitCreditAdviceLineItemD.Uuid4 = uuid4
	debitCreditAdviceLineItemD.IdS = idS
	debitCreditAdviceLineItemD.AdjustmentAmount = adjustmentAmount
	debitCreditAdviceLineItemD.AaCodeListVersion = aaCodeListVersion
	debitCreditAdviceLineItemD.AaCurrencyCode = aaCurrencyCode
	debitCreditAdviceLineItemD.DebitCreditIndicatorCode = debitCreditIndicatorCode
	debitCreditAdviceLineItemD.FinancialAdjustmentReasonCode = financialAdjustmentReasonCode
	debitCreditAdviceLineItemD.LineItemNumber = lineItemNumber
	debitCreditAdviceLineItemD.ParentLineItemNumber = parentLineItemNumber
	debitCreditAdviceLineItemD.DebitCreditAdviceId = debitCreditAdviceId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	debitCreditAdviceLineItemResult := invoiceproto.DebitCreditAdviceLineItem{DebitCreditAdviceLineItemD: &debitCreditAdviceLineItemD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &debitCreditAdviceLineItemResult, nil
}
