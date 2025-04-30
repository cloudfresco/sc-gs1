package epcisservices

import (
	"context"
	"reflect"
	"testing"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestEpcisService_CreateDestination(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	destination := epcisproto.CreateDestinationRequest{}
	destination.DestType = "location"
	destination.Destination = "urn:epc:id:sgln:0614141.00777"
	destination.EventId = uint32(1)
	destination.TypeOfEvent = "ObjectEvent"
	destination.UserId = "auth0|673ee1a719dd4000cd5a3832"
	destination.UserEmail = "sprov300@gmail.com"
	destination.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateDestinationRequest
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
				in:  &destination,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		destinationResp, err := tt.es.CreateDestination(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateDestination() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, destinationResp)
		destinationResult := destinationResp.Destination
		assert.Equal(t, destinationResult.DestType, "location", "they should be equal")
		assert.Equal(t, destinationResult.Destination, "urn:epc:id:sgln:0614141.00777", "they should be equal")
	}
}

func TestEpcisService_GetDestinations(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	destination3, err := GetDestination("location", "urn:epc:id:sgln:0614141.00777.0", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	destination2, err := GetDestination("possessing_party", "urn:epc:id:pgln:0614141.00777", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	destination1, err := GetDestination("owning_party", "urn:epc:id:pgln:0614141.00777", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	destinations := []*epcisproto.Destination{}
	destinations = append(destinations, destination3, destination2, destination1)

	destinationsResponse := epcisproto.GetDestinationsResponse{}
	destinationsResponse.Destinations = destinations

	form := epcisproto.GetDestinationsRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *epcisproto.GetDestinationsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetDestinationsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &destinationsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		destinationResponse, err := tt.es.GetDestinations(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetDestinations() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(destinationResponse, tt.want) {
			t.Errorf("EpcisService.GetDestinations() = %v, want %v", destinationResponse, tt.want)
		}
		destinationResult := destinationResponse.Destinations[1]
		assert.NotNil(t, destinationResult)
		assert.Equal(t, destinationResult.DestType, "possessing_party", "they should be equal")
		assert.Equal(t, destinationResult.Destination, "urn:epc:id:pgln:0614141.00777", "they should be equal")
	}
}

func GetDestination(destType string, src string, eventId uint32, typeOfEvent string) (*epcisproto.Destination, error) {
	destination := epcisproto.Destination{}
	destination.DestType = destType
	destination.Destination = src
	destination.EventId = eventId
	destination.TypeOfEvent = typeOfEvent

	return &destination, nil
}
