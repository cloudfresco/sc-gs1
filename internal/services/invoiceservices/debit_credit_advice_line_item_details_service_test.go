package invoiceservices

import (
	"context"
	"testing"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestDebitCreditAdviceService_CreateDebitCreditAdviceLineItemDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)
	debitCreditAdviceLineItemDetail := invoiceproto.CreateDebitCreditAdviceLineItemDetailRequest{}
	debitCreditAdviceLineItemDetail.AlignedPrice = float64(0)
	debitCreditAdviceLineItemDetail.ApCodeListVersion = ""
	debitCreditAdviceLineItemDetail.ApCurrencyCode = ""
	debitCreditAdviceLineItemDetail.InvoicedPrice = float64(0)
	debitCreditAdviceLineItemDetail.IpCodeListVersion = ""
	debitCreditAdviceLineItemDetail.IpCurrencyCode = ""
	debitCreditAdviceLineItemDetail.Quantity = float64(10)
	debitCreditAdviceLineItemDetail.QMeasurementUnitCode = ""
	debitCreditAdviceLineItemDetail.QCodeListVersion = ""
	debitCreditAdviceLineItemDetail.DebitCreditAdviceId = uint32(1)
	debitCreditAdviceLineItemDetail.DebitCreditAdviceLineItemId = uint32(1)

	type args struct {
		ctx context.Context
		in  *invoiceproto.CreateDebitCreditAdviceLineItemDetailRequest
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
				in:  &debitCreditAdviceLineItemDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceLineItemDetailResult, err := tt.ds.CreateDebitCreditAdviceLineItemDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.CreateDebitCreditAdviceLineItemDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, debitCreditAdviceLineItemDetailResult)
		assert.Equal(t, debitCreditAdviceLineItemDetailResult.DebitCreditAdviceLineItemDetail.Quantity, float64(10), "they should be equal")
	}
}
