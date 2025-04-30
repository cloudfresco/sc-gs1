package epcisservices

import (
	"context"
	"reflect"
	"testing"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestEpcisService_CreateEpc(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	epc := epcisproto.CreateEpcRequest{}
	epc.EpcValue = "urn:epc:id:sgtin:0614141.107346.2018"
	epc.EventId = uint32(1)
	epc.TypeOfEvent = "ObjectEvent"
	epc.TypeOfEpc = "output"
	epc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	epc.UserEmail = "sprov300@gmail.com"
	epc.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateEpcRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.CreateEpcResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &epc,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		epcResp, err := tt.es.CreateEpc(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateEpc() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, epcResp)
		epcResult := epcResp.Epc
		assert.Equal(t, epcResult.EpcValue, "urn:epc:id:sgtin:0614141.107346.2018", "they should be equal")
	}
}

func TestEpcisService_GetEpcs(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	epc, err := GetEpc("urn:epc:id:sgtin:0614141.107346.2018", uint32(1), "ObjectEvent", "output")
	if err != nil {
		t.Error(err)
	}

	epcs := []*epcisproto.Epc{}
	epcs = append(epcs, epc)

	epcsResponse := epcisproto.GetEpcsResponse{}
	epcsResponse.Epcs = epcs

	form := epcisproto.GetEpcsRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.TypeOfEpc = "output"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.GetEpcsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetEpcsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &epcsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		epcResponse, err := tt.es.GetEpcs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetEpcs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(epcResponse, tt.want) {
			t.Errorf("EpcisService.GetEpcs() = %v, want %v", epcResponse, tt.want)
		}
		epcResult := epcResponse.Epcs[0]
		assert.NotNil(t, epcResult)
		assert.Equal(t, epcResult.EpcValue, "urn:epc:id:sgtin:0614141.107346.2018", "they should be equal")
	}
}

func GetEpc(epcValue string, eventId uint32, typeOfEvent string, typeOfEpc string) (*epcisproto.Epc, error) {
	epc := epcisproto.Epc{}
	epc.EpcValue = epcValue
	epc.EventId = eventId
	epc.TypeOfEvent = typeOfEvent
	epc.TypeOfEpc = typeOfEpc

	return &epc, nil
}
