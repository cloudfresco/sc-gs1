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

func TestEpcisService_CreateTransformationEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	tranformationEvent := epcisproto.CreateTransformationEventRequest{}
	tranformationEvent.EventId = "ni:///sha-256;0bf4271d60ed65fb687e95f7216c4c0a4c1181c070f657d41385b6fbd93e97ef?ver=CBV2.0"
	tranformationEvent.EventTimeZoneOffset = "+02:00"
	tranformationEvent.Certification = "https://accreditation-council.example.org/certificate/ABC12345"
	tranformationEvent.EventTime = "11/10/2013"
	tranformationEvent.Reason = "incorrect_data"
	tranformationEvent.DeclarationTime = "07/11/2013"
	tranformationEvent.TransformationId = "urn:epc:id:gdti:0614141.12345.400"
	tranformationEvent.BizStep = "commissioning"
	tranformationEvent.Disposition = "in_progress"
	tranformationEvent.ReadPoint = "urn:epc:id:sgln:4012345.00001.0"
	tranformationEvent.BizLocation = "urn:epc:id:sgln:0614141.00888.0"
	tranformationEvent.Ilmd = "20"
	tranformationEvent.UserId = "auth0|673ee1a719dd4000cd5a3832"
	tranformationEvent.UserEmail = "sprov300@gmail.com"
	tranformationEvent.RequestId = "bks1m1g91jau4nkks2f0"

	persistentDisposition := epcisproto.CreatePersistentDispositionRequest{}
	persistentDisposition.SetDisp = "completeness_verified"
	persistentDisposition.UnsetDisp = "completeness_inferred"
	persistentDisposition.EventId = uint32(1)
	persistentDisposition.TypeOfEvent = "TransformationEvent"
	persistentDisposition.UserId = "auth0|673ee1a719dd4000cd5a3832"
	persistentDisposition.UserEmail = "sprov300@gmail.com"
	persistentDisposition.RequestId = "bks1m1g91jau4nkks2f0"
	persistentDispositions := []*epcisproto.CreatePersistentDispositionRequest{}
	persistentDispositions = append(persistentDispositions, &persistentDisposition)
	tranformationEvent.PersistentDispositions = persistentDispositions

	epc := epcisproto.CreateEpcRequest{}
	epc.EpcValue = "urn:epc:id:sgtin:4012345.011122.25"
	epc.EventId = uint32(1)
	epc.TypeOfEvent = "TransformationEvent"
	epc.TypeOfEpc = "input"
	epc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	epc.UserEmail = "sprov300@gmail.com"
	epc.RequestId = "bks1m1g91jau4nkks2f0"

	epcs := []*epcisproto.CreateEpcRequest{}
	epcs = append(epcs, &epc)
	tranformationEvent.InputEpcList = epcs

	outepc := epcisproto.CreateEpcRequest{}
	outepc.EpcValue = "urn:epc:id:sgtin:4012345.077889.25"
	outepc.EventId = uint32(1)
	outepc.TypeOfEvent = "TransformationEvent"
	outepc.TypeOfEpc = "output"
	outepc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	outepc.UserEmail = "sprov300@gmail.com"
	outepc.RequestId = "bks1m1g91jau4nkks2f0"

	outepcs := []*epcisproto.CreateEpcRequest{}
	outepcs = append(outepcs, &outepc)
	tranformationEvent.OutputEpcList = outepcs

	bizTransaction := epcisproto.CreateBizTransactionRequest{}
	bizTransaction.BizTransactionType = "po"
	bizTransaction.BizTransaction = "urn:epc:id:gdti:0614141.00001.1618034"
	bizTransaction.EventId = uint32(1)
	bizTransaction.TypeOfEvent = "TransformationEvent"
	bizTransaction.UserId = "auth0|673ee1a719dd4000cd5a3832"
	bizTransaction.UserEmail = "sprov300@gmail.com"
	bizTransaction.RequestId = "bks1m1g91jau4nkks2f0"

	bizTransactions := []*epcisproto.CreateBizTransactionRequest{}
	bizTransactions = append(bizTransactions, &bizTransaction)
	tranformationEvent.BizTransactionList = bizTransactions

	inQuantityElement := epcisproto.CreateQuantityElementRequest{}
	inQuantityElement.EpcClass = "urn:epc:class:lgtin:4012345.011111.4444"
	inQuantityElement.Quantity = float64(10)
	inQuantityElement.Uom = "KGM"
	inQuantityElement.EventId = uint32(1)
	inQuantityElement.TypeOfEvent = "TransformationEvent"
	inQuantityElement.TypeOfQuantity = "input"
	inQuantityElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	inQuantityElement.UserEmail = "sprov300@gmail.com"
	inQuantityElement.RequestId = "bks1m1g91jau4nkks2f0"

	inQuantityElements := []*epcisproto.CreateQuantityElementRequest{}
	inQuantityElements = append(inQuantityElements, &inQuantityElement)
	tranformationEvent.InputQuantityList = inQuantityElements

	outQuantityElement := epcisproto.CreateQuantityElementRequest{}
	outQuantityElement.EpcClass = "uurn:epc:class:lgtin:0614141.077777.987"
	outQuantityElement.Quantity = float64(30)
	outQuantityElement.Uom = "KGM"
	outQuantityElement.EventId = uint32(1)
	outQuantityElement.TypeOfEvent = "TransformationEvent"
	outQuantityElement.TypeOfQuantity = "output"
	outQuantityElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	outQuantityElement.UserEmail = "sprov300@gmail.com"
	outQuantityElement.RequestId = "bks1m1g91jau4nkks2f0"

	outQuantityElements := []*epcisproto.CreateQuantityElementRequest{}
	outQuantityElements = append(outQuantityElements, &outQuantityElement)
	tranformationEvent.OutputQuantityList = outQuantityElements

	source := epcisproto.CreateSourceRequest{}
	source.SourceType = "location"
	source.Source = "urn:epc:id:sgln:4012345.00225.0"
	source.EventId = uint32(1)
	source.TypeOfEvent = "TransformationEvent"
	source.UserId = "auth0|673ee1a719dd4000cd5a3832"
	source.UserEmail = "sprov300@gmail.com"
	source.RequestId = "bks1m1g91jau4nkks2f0"
	sources := []*epcisproto.CreateSourceRequest{}
	sources = append(sources, &source)
	tranformationEvent.SourceList = sources

	destination := epcisproto.CreateDestinationRequest{}
	destination.DestType = "location"
	destination.Destination = "urn:epc:id:sgln:0614141.00777.0"
	destination.EventId = uint32(1)
	destination.TypeOfEvent = "TransformationEvent"
	destination.UserId = "auth0|673ee1a719dd4000cd5a3832"
	destination.UserEmail = "sprov300@gmail.com"
	destination.RequestId = "bks1m1g91jau4nkks2f0"

	destinations := []*epcisproto.CreateDestinationRequest{}
	destinations = append(destinations, &destination)
	tranformationEvent.DestinationList = destinations

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
	sensorElement.TypeOfEvent = "TransformationEvent"
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
	tranformationEvent.SensorElementList = sensorElements

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateTransformationEventRequest
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
				in:  &tranformationEvent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tranformationEventResp, err := tt.es.CreateTransformationEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateTransformationEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, tranformationEventResp)
		tranformationEventResult := tranformationEventResp.TransformationEvent
		assert.Equal(t, tranformationEventResult.EpcisEventD.EventId, "ni:///sha-256;0bf4271d60ed65fb687e95f7216c4c0a4c1181c070f657d41385b6fbd93e97ef?ver=CBV2.0", "they should be equal")
		assert.Equal(t, tranformationEventResult.EpcisEventD.EventTimeZoneOffset, "+02:00", "they should be equal")
		assert.Equal(t, tranformationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.BizStep, "commissioning", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.ReadPoint, "urn:epc:id:sgln:4012345.00001.0", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func TestEpcisService_GetTransformationEvents(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	tranformationEvents := []*epcisproto.TransformationEvent{}

	tranformationEvent, err := GetTransformationEvent(uint32(1), []byte{84, 42, 209, 102, 123, 104, 65, 228, 158, 194, 77, 23, 111, 115, 88, 130}, "542ad166-7b68-41e4-9ec2-4d176f735882", "ni:///sha-256;0bf4271d60ed65fb687e95f7216c4c0a4c1181c070f657d41385b6fbd93e97ef?ver=CBV2.0", "+02:00", "https://accreditation-council.example.org/certificate/ABC12345", "2013-10-31T14:58:56Z", "incorrect_data", "2013-11-07T14:00:00Z", "urn:epc:id:gdti:0614141.12345.400", "commissioning", "in_progress", "urn:epc:id:sgln:4012345.00001.0", "urn:epc:id:sgln:0614141.00888.0", "20", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	tranformationEvents = append(tranformationEvents, tranformationEvent)

	form := epcisproto.GetTransformationEventsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	tranformationEventsResponse := epcisproto.GetTransformationEventsResponse{TransformationEvents: tranformationEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *epcisproto.GetTransformationEventsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetTransformationEventsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &tranformationEventsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tranformationEventResp, err := tt.es.GetTransformationEvents(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetTransformationEvents() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(tranformationEventResp, tt.want) {
			t.Errorf("EpcisService.GetTransformationEvents() = %v, want %v", tranformationEventResp, tt.want)
		}
		assert.NotNil(t, tranformationEventResp)
		tranformationEventResult := tranformationEventResp.TransformationEvents[0]
		assert.Equal(t, tranformationEventResult.EpcisEventD.EventId, "ni:///sha-256;0bf4271d60ed65fb687e95f7216c4c0a4c1181c070f657d41385b6fbd93e97ef?ver=CBV2.0", "they should be equal")
		assert.Equal(t, tranformationEventResult.EpcisEventD.EventTimeZoneOffset, "+02:00", "they should be equal")
		assert.Equal(t, tranformationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.BizStep, "commissioning", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.ReadPoint, "urn:epc:id:sgln:4012345.00001.0", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func TestEpcisService_GetTransformationEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	tranformationEvent, err := GetTransformationEvent(uint32(1), []byte{84, 42, 209, 102, 123, 104, 65, 228, 158, 194, 77, 23, 111, 115, 88, 130}, "542ad166-7b68-41e4-9ec2-4d176f735882", "ni:///sha-256;0bf4271d60ed65fb687e95f7216c4c0a4c1181c070f657d41385b6fbd93e97ef?ver=CBV2.0", "+02:00", "https://accreditation-council.example.org/certificate/ABC12345", "2013-10-31T14:58:56Z", "incorrect_data", "2013-11-07T14:00:00Z", "urn:epc:id:gdti:0614141.12345.400", "commissioning", "in_progress", "urn:epc:id:sgln:4012345.00001.0", "urn:epc:id:sgln:0614141.00888.0", "20", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	tranformationEventResponse := epcisproto.GetTransformationEventResponse{}
	tranformationEventResponse.TransformationEvent = tranformationEvent

	form := epcisproto.GetTransformationEventRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "542ad166-7b68-41e4-9ec2-4d176f735882"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *epcisproto.GetTransformationEventRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetTransformationEventResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &tranformationEventResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tranformationEventResp, err := tt.es.GetTransformationEvent(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetTransformationEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(tranformationEventResp, tt.want) {
			t.Errorf("EpcisService.GetTransformationEvent() = %v, want %v", tranformationEventResp, tt.want)
		}
		assert.NotNil(t, tranformationEventResp)
		tranformationEventResult := tranformationEventResp.TransformationEvent
		assert.Equal(t, tranformationEventResult.EpcisEventD.EventId, "ni:///sha-256;0bf4271d60ed65fb687e95f7216c4c0a4c1181c070f657d41385b6fbd93e97ef?ver=CBV2.0", "they should be equal")
		assert.Equal(t, tranformationEventResult.EpcisEventD.EventTimeZoneOffset, "+02:00", "they should be equal")
		assert.Equal(t, tranformationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.BizStep, "commissioning", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.ReadPoint, "urn:epc:id:sgln:4012345.00001.0", "they should be equal")
		assert.Equal(t, tranformationEventResult.TransformationEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func GetTransformationEvent(id uint32, uuid4 []byte, idS string, eventId string, eventTimeZoneOffset string, certification string, eventTime string, reason string, declarationTime string, transformationId string, bizStep string, disposition string, readPoint string, bizLocation string, ilmd string, statusCode string, createdByUserId string, updatedByUserId string) (*epcisproto.TransformationEvent, error) {
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

	tranformationEventD := epcisproto.TransformationEventD{}
	tranformationEventD.Id = id
	tranformationEventD.Uuid4 = uuid4
	tranformationEventD.IdS = idS
	tranformationEventD.TransformationId = transformationId
	tranformationEventD.BizStep = bizStep
	tranformationEventD.Disposition = disposition
	tranformationEventD.ReadPoint = readPoint
	tranformationEventD.BizLocation = bizLocation
	tranformationEventD.Ilmd = ilmd

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = statusCode
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	tranformationEvent := epcisproto.TransformationEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, TransformationEventD: &tranformationEventD, CrUpdUser: &crUpdUser}

	return &tranformationEvent, nil
}
