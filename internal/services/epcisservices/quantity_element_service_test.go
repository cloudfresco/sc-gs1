package epcisservices

import (
	"context"
	"reflect"
	"testing"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestEpcisService_CreateQuantityElement(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	quantityElement := epcisproto.CreateQuantityElementRequest{}
	quantityElement.EpcClass = "urn:epc:class:lgtin:4012345.012345.998889"
	quantityElement.Quantity = float64(200)
	quantityElement.Uom = "KGM"
	quantityElement.EventId = uint32(1)
	quantityElement.TypeOfEvent = "ObjectEvent"
	quantityElement.TypeOfQuantity = "output"
	quantityElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	quantityElement.UserEmail = "sprov300@gmail.com"
	quantityElement.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateQuantityElementRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &quantityElement,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		quantityElementResp, err := tt.es.CreateQuantityElement(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateQuantityElement() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, quantityElementResp)
		quantityElementResult := quantityElementResp.QuantityElement
		assert.Equal(t, quantityElementResult.EpcClass, "urn:epc:class:lgtin:4012345.012345.998889", "they should be equal")
		assert.Equal(t, quantityElementResult.Quantity, float64(200), "they should be equal")
		assert.Equal(t, quantityElementResult.Uom, "KGM", "they should be equal")
	}
}

func TestEpcisService_GetQuantityElements(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	quantityElement, err := GetQuantityElement("urn:epc:class:lgtin:4012345.012345.998877", float64(200), "KGM", uint32(1), "ObjectEvent", "output")
	if err != nil {
		t.Error(err)
	}

	quantityElements := []*epcisproto.QuantityElement{}
	quantityElements = append(quantityElements, quantityElement)

	quantityElementsResponse := epcisproto.GetQuantityElementsResponse{}
	quantityElementsResponse.QuantityElements = quantityElements

	form := epcisproto.GetQuantityElementsRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.TypeOfQuantity = "output"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *epcisproto.GetQuantityElementsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetQuantityElementsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &quantityElementsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		quantityElementResponse, err := tt.es.GetQuantityElements(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetQuantityElements() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(quantityElementResponse, tt.want) {
			t.Errorf("EpcisService.GetQuantityElements() = %v, want %v", quantityElementResponse, tt.want)
		}
		quantityElementResult := quantityElementResponse.QuantityElements[0]
		assert.NotNil(t, quantityElementResult)
		assert.Equal(t, quantityElementResult.EpcClass, "urn:epc:class:lgtin:4012345.012345.998877", "they should be equal")
		assert.Equal(t, quantityElementResult.Quantity, float64(200), "they should be equal")
		assert.Equal(t, quantityElementResult.Uom, "KGM", "they should be equal")
	}
}

func GetQuantityElement(epcClass string, quantity float64, uom string, eventId uint32, typeOfEvent string, typeOfQuantity string) (*epcisproto.QuantityElement, error) {
	quantityElement := epcisproto.QuantityElement{}
	quantityElement.EpcClass = epcClass
	quantityElement.Quantity = quantity
	quantityElement.Uom = uom
	quantityElement.EventId = eventId
	quantityElement.TypeOfEvent = typeOfEvent
	quantityElement.TypeOfQuantity = typeOfQuantity

	return &quantityElement, nil
}
