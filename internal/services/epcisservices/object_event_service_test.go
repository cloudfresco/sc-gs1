package epcisservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestEpcisService_CreateObjectEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	objectEvent := epcisproto.CreateObjectEventRequest{}
	objectEvent.EventId = "urn:uuid:374d95fc-9457-4a51-bd6a-0bba133845a8"
	objectEvent.EventTimeZoneOffset = "-06:00"
	objectEvent.Certification = "https://accreditation-council.example.org/certificate/ABC12345"
	objectEvent.EventTime = "04/05/2005"
	objectEvent.Reason = "incorrect_data"
	objectEvent.DeclarationTime = "07/11/2005"
	objectEvent.Action = "ADD"
	objectEvent.BizStep = "receiving"
	objectEvent.Disposition = "in_progress"
	objectEvent.ReadPoint = "urn:epc:id:sgln:0012345.11111.400"
	objectEvent.BizLocation = "urn:epc:id:sgln:0012345.11111.0"
	objectEvent.Ilmd = ""
	objectEvent.UserId = "auth0|673ee1a719dd4000cd5a3832"
	objectEvent.UserEmail = "sprov300@gmail.com"
	objectEvent.RequestId = "bks1m1g91jau4nkks2f0"

	persistentDisposition := epcisproto.CreatePersistentDispositionRequest{}
	persistentDisposition.SetDisp = "completeness_verified"
	persistentDisposition.UnsetDisp = "completeness_inferred"
	persistentDisposition.EventId = uint32(1)
	persistentDisposition.TypeOfEvent = "ObjectEvent"
	persistentDisposition.UserId = "auth0|673ee1a719dd4000cd5a3832"
	persistentDisposition.UserEmail = "sprov300@gmail.com"
	persistentDisposition.RequestId = "bks1m1g91jau4nkks2f0"
	persistentDispositions := []*epcisproto.CreatePersistentDispositionRequest{}
	persistentDispositions = append(persistentDispositions, &persistentDisposition)
	objectEvent.PersistentDispositions = persistentDispositions

	epc := epcisproto.CreateEpcRequest{}
	epc.EpcValue = "urn:epc:id:sgtin:0614141.107346.2018"
	epc.EventId = uint32(1)
	epc.TypeOfEvent = "ObjectEvent"
	epc.TypeOfEpc = "output"
	epc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	epc.UserEmail = "sprov300@gmail.com"
	epc.RequestId = "bks1m1g91jau4nkks2f0"

	epcs := []*epcisproto.CreateEpcRequest{}
	epcs = append(epcs, &epc)
	objectEvent.EpcList = epcs

	bizTransaction := epcisproto.CreateBizTransactionRequest{}
	bizTransaction.BizTransactionType = "pedigree"
	bizTransaction.BizTransaction = "urn:epc:id:gsrn:0614141.0000010254"
	bizTransaction.EventId = uint32(1)
	bizTransaction.TypeOfEvent = "ObjectEvent"
	bizTransaction.UserId = "auth0|673ee1a719dd4000cd5a3832"
	bizTransaction.UserEmail = "sprov300@gmail.com"
	bizTransaction.RequestId = "bks1m1g91jau4nkks2f0"

	bizTransactions := []*epcisproto.CreateBizTransactionRequest{}
	bizTransactions = append(bizTransactions, &bizTransaction)
	objectEvent.BizTransactionList = bizTransactions

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

	quantityElements := []*epcisproto.CreateQuantityElementRequest{}
	quantityElements = append(quantityElements, &quantityElement)
	objectEvent.QuantityList = quantityElements

	source := epcisproto.CreateSourceRequest{}
	source.SourceType = "location"
	source.Source = "urn:epc:id:pgln:4012345.00225"
	source.EventId = uint32(1)
	source.TypeOfEvent = "ObjectEvent"
	source.UserId = "auth0|673ee1a719dd4000cd5a3832"
	source.UserEmail = "sprov300@gmail.com"
	source.RequestId = "bks1m1g91jau4nkks2f0"
	sources := []*epcisproto.CreateSourceRequest{}
	sources = append(sources, &source)
	objectEvent.SourceList = sources

	destination := epcisproto.CreateDestinationRequest{}
	destination.DestType = "location"
	destination.Destination = "urn:epc:id:sgln:0614141.00777"
	destination.EventId = uint32(1)
	destination.TypeOfEvent = "ObjectEvent"
	destination.UserId = "auth0|673ee1a719dd4000cd5a3832"
	destination.UserEmail = "sprov300@gmail.com"
	destination.RequestId = "bks1m1g91jau4nkks2f0"

	destinations := []*epcisproto.CreateDestinationRequest{}
	destinations = append(destinations, &destination)
	objectEvent.DestinationList = destinations

	sensorElement := epcisproto.CreateSensorElementRequest{}
	sensorElement.DeviceId = "urn:epc:id:giai:4000001.111"
	sensorElement.DeviceMetadata = "https://id.gs1.org/8004/4000001111"
	sensorElement.RawData = "https://example.org/8004/401234599999"
	sensorElement.DataProcessingMethod = "https://example.com/253/4012345000054987"
	sensorElement.BizRules = "https://example.com/253/4012345000054987"
	sensorElement.SensorTime = "04/11/2023"
	sensorElement.StartTime = "04/11/2023"
	sensorElement.EndTime = "04/11/2023"
	sensorElement.EventId = uint32(1)
	sensorElement.TypeOfEvent = "ObjectEvent"
	sensorElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	sensorElement.UserEmail = "sprov300@gmail.com"
	sensorElement.RequestId = "bks1m1g91jau4nkks2f0"

	sensorReport := epcisproto.CreateSensorReportRequest{}
	sensorReport.SensorReportType = "Temperature"
	sensorReport.DeviceId = "urn:epc:id:giai:4000001.111"
	sensorReport.RawData = "https://example.org/8004/401234599999"
	sensorReport.DataProcessingMethod = "https://example.com/253/4012345000054987"
	sensorReport.Microorganism = "https://www.ncbi.nlm.nih.gov/taxonomy/1126011"
	sensorReport.ChemicalSubstance = "https://identifiers.org/inchikey:CZMRCDWAGMRECN-UGDNZRGBSA-N"
	sensorReport.SensorValue = float64(26)
	sensorReport.Component = "example:x"
	sensorReport.StringValue = "SomeString"
	sensorReport.BooleanValue = true
	sensorReport.HexBinaryValue = "f0f0f0"
	sensorReport.UriValue = "https://id.gs1.org/8004/4000001111"
	sensorReport.MinValue = float64(26)
	sensorReport.MaxValue = float64(26.2)
	sensorReport.MeanValue = float64(13.2)
	sensorReport.PercRank = float64(50)
	sensorReport.PercValue = float64(12.7)
	sensorReport.Uom = "CEL"
	sensorReport.SDev = float64(0.1)
	sensorReport.DeviceMetadata = "https://id.gs1.org/8004/4000001111"
	sensorReport.SensorReportTime = "04/11/2023"
	sensorReport.UserId = "auth0|673ee1a719dd4000cd5a3832"
	sensorReport.UserEmail = "sprov300@gmail.com"
	sensorReport.RequestId = "bks1m1g91jau4nkks2f0"

	sensorReports := []*epcisproto.CreateSensorReportRequest{}
	sensorReports = append(sensorReports, &sensorReport)
	sensorElement.SensorReports = sensorReports
	sensorElements := []*epcisproto.CreateSensorElementRequest{}
	sensorElements = append(sensorElements, &sensorElement)
	objectEvent.SensorElementList = sensorElements

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateObjectEventRequest
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
				in:  &objectEvent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		objectEventResp, err := tt.es.CreateObjectEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateObjectEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, objectEventResp)
		objectEventResult := objectEventResp.ObjectEvent
		assert.Equal(t, objectEventResult.EpcisEventD.EventId, "urn:uuid:374d95fc-9457-4a51-bd6a-0bba133845a8", "they should be equal")
		assert.Equal(t, objectEventResult.EpcisEventD.EventTimeZoneOffset, "-06:00", "they should be equal")
		assert.Equal(t, objectEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.BizStep, "receiving", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.ReadPoint, "urn:epc:id:sgln:0012345.11111.400", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.BizLocation, "urn:epc:id:sgln:0012345.11111.0", "they should be equal")
	}
}

func TestEpcisService_GetObjectEvents(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	objectEvents := []*epcisproto.ObjectEvent{}

	objectEvent, err := GetObjectEvent(uint32(1), []byte{244, 131, 80, 202, 214, 248, 72, 217, 158, 91, 61, 170, 137, 248, 102, 176}, "f48350ca-d6f8-48d9-9e5b-3daa89f866b0", "urn:uuid:374d95fc-9457-4a51-bd6a-0bba133845a8", "-06:00", "https://accreditation-council.example.org/certificate/ABC12345", "2005-04-05T02:33:31Z", "incorrect_data", "2006-11-07T14:00:00Z", "ADD", "receiving", "in_progress", "urn:epc:id:sgln:0012345.11111.400", "urn:epc:id:sgln:0012345.11111.0", "", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	objectEvents = append(objectEvents, objectEvent)

	form := epcisproto.GetObjectEventsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	objectEventsResponse := epcisproto.GetObjectEventsResponse{ObjectEvents: objectEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *epcisproto.GetObjectEventsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetObjectEventsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &objectEventsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		objectEventResp, err := tt.es.GetObjectEvents(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetObjectEvents() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(objectEventResp, tt.want) {
			t.Errorf("EpcisService.GetObjectEvents() = %v, want %v", objectEventResp, tt.want)
		}
		assert.NotNil(t, objectEventResp)
		objectEventResult := objectEventResp.ObjectEvents[0]
		assert.Equal(t, objectEventResult.EpcisEventD.EventId, "urn:uuid:374d95fc-9457-4a51-bd6a-0bba133845a8", "they should be equal")
		assert.Equal(t, objectEventResult.EpcisEventD.EventTimeZoneOffset, "-06:00", "they should be equal")
		assert.Equal(t, objectEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.BizStep, "receiving", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.ReadPoint, "urn:epc:id:sgln:0012345.11111.400", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.BizLocation, "urn:epc:id:sgln:0012345.11111.0", "they should be equal")
	}
}

func TestEpcisService_GetObjectEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	objectEvent, err := GetObjectEvent(uint32(1), []byte{244, 131, 80, 202, 214, 248, 72, 217, 158, 91, 61, 170, 137, 248, 102, 176}, "f48350ca-d6f8-48d9-9e5b-3daa89f866b0", "urn:uuid:374d95fc-9457-4a51-bd6a-0bba133845a8", "-06:00", "https://accreditation-council.example.org/certificate/ABC12345", "2005-04-05T02:33:31Z", "incorrect_data", "2006-11-07T14:00:00Z", "ADD", "receiving", "in_progress", "urn:epc:id:sgln:0012345.11111.400", "urn:epc:id:sgln:0012345.11111.0", "", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	objectEventResponse := epcisproto.GetObjectEventResponse{}
	objectEventResponse.ObjectEvent = objectEvent

	form := epcisproto.GetObjectEventRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "f48350ca-d6f8-48d9-9e5b-3daa89f866b0"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *epcisproto.GetObjectEventRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetObjectEventResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &objectEventResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		objectEventResp, err := tt.es.GetObjectEvent(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetObjectEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(objectEventResp, tt.want) {
			t.Errorf("EpcisService.GetObjectEvent() = %v, want %v", objectEventResp, tt.want)
		}
		assert.NotNil(t, objectEventResp)
		objectEventResult := objectEventResp.ObjectEvent
		assert.Equal(t, objectEventResult.EpcisEventD.EventId, "urn:uuid:374d95fc-9457-4a51-bd6a-0bba133845a8", "they should be equal")
		assert.Equal(t, objectEventResult.EpcisEventD.EventTimeZoneOffset, "-06:00", "they should be equal")
		assert.Equal(t, objectEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.BizStep, "receiving", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.ReadPoint, "urn:epc:id:sgln:0012345.11111.400", "they should be equal")
		assert.Equal(t, objectEventResult.ObjectEventD.BizLocation, "urn:epc:id:sgln:0012345.11111.0", "they should be equal")
	}
}

func GetObjectEvent(id uint32, uuid4 []byte, idS string, eventId string, eventTimeZoneOffset string, certification string, eventTime string, reason string, declarationTime string, action string, bizStep string, disposition string, readPoint string, bizLocation string, ilmd string, statusCode string, createdByUserId string, updatedByUserId string) (*epcisproto.ObjectEvent, error) {
	eventTime1, err := common.ConvertTimeToTimestamp(Layout, eventTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	declarationTime1, err := common.ConvertTimeToTimestamp(Layout, declarationTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	epcisEventD := commonproto.EpcisEventD{}
	epcisEventD.EventId = eventId
	epcisEventD.EventTimeZoneOffset = eventTimeZoneOffset
	epcisEventD.Certification = certification

	epcisEventT := commonproto.EpcisEventT{}
	epcisEventT.EventTime = eventTime1

	errorDeclarationD := commonproto.ErrorDeclarationD{}
	errorDeclarationD.Reason = reason

	errorDeclarationT := commonproto.ErrorDeclarationT{}
	errorDeclarationT.DeclarationTime = declarationTime1

	objectEventD := epcisproto.ObjectEventD{}
	objectEventD.Id = id
	objectEventD.Uuid4 = uuid4
	objectEventD.IdS = idS
	objectEventD.Action = action
	objectEventD.BizStep = bizStep
	objectEventD.Disposition = disposition
	objectEventD.ReadPoint = readPoint
	objectEventD.BizLocation = bizLocation
	objectEventD.Ilmd = ilmd

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = statusCode
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	objectEvent := epcisproto.ObjectEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, ObjectEventD: &objectEventD, CrUpdUser: &crUpdUser}

	return &objectEvent, nil
}
