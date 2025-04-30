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

func TestEpcisService_CreateTransactionEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	transactionEvent := epcisproto.CreateTransactionEventRequest{}
	transactionEvent.EventId = "ni:///sha-256;45a99ca926fdb62b61bb2b29620e1dcdd5b0109613700f7e179881d64d8fabf1?ver=CBV2.0"
	transactionEvent.EventTimeZoneOffset = "-06:00"
	transactionEvent.Certification = "https://accreditation-council.example.org/certificate/ABC12345"
	transactionEvent.EventTime = "04/05/2005"
	transactionEvent.Reason = "incorrect_data"
	transactionEvent.DeclarationTime = "07/11/2005"
	transactionEvent.ParentId = "urn:epc:id:sscc:0614141.1234567890"
	transactionEvent.Action = "ADD"
	transactionEvent.BizStep = "shipping"
	transactionEvent.Disposition = "in_transit"
	transactionEvent.ReadPoint = "urn:epc:id:sgln:0614141.07346.1234"
	transactionEvent.BizLocation = "urn:epc:id:sgln:0614141.00888.0"
	transactionEvent.UserId = "auth0|673ee1a719dd4000cd5a3832"
	transactionEvent.UserEmail = "sprov300@gmail.com"
	transactionEvent.RequestId = "bks1m1g91jau4nkks2f0"

	persistentDisposition := epcisproto.CreatePersistentDispositionRequest{}
	persistentDisposition.SetDisp = "completeness_verified"
	persistentDisposition.UnsetDisp = "completeness_inferred"
	persistentDisposition.EventId = uint32(1)
	persistentDisposition.TypeOfEvent = "TransactionEvent"
	persistentDisposition.UserId = "auth0|673ee1a719dd4000cd5a3832"
	persistentDisposition.UserEmail = "sprov300@gmail.com"
	persistentDisposition.RequestId = "bks1m1g91jau4nkks2f0"
	persistentDispositions := []*epcisproto.CreatePersistentDispositionRequest{}
	persistentDispositions = append(persistentDispositions, &persistentDisposition)
	transactionEvent.PersistentDispositions = persistentDispositions

	epc := epcisproto.CreateEpcRequest{}
	epc.EpcValue = "urn:epc:id:sgtin:0614141.107346.2017"
	epc.EventId = uint32(1)
	epc.TypeOfEvent = "TransactionEvent"
	epc.TypeOfEpc = "output"
	epc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	epc.UserEmail = "sprov300@gmail.com"
	epc.RequestId = "bks1m1g91jau4nkks2f0"

	epcs := []*epcisproto.CreateEpcRequest{}
	epcs = append(epcs, &epc)
	transactionEvent.EpcList = epcs

	bizTransaction := epcisproto.CreateBizTransactionRequest{}
	bizTransaction.BizTransactionType = "po"
	bizTransaction.BizTransaction = "urn:epc:id:gdti:0614141.00001.1618034"
	bizTransaction.EventId = uint32(1)
	bizTransaction.TypeOfEvent = "TransactionEvent"
	bizTransaction.UserId = "auth0|673ee1a719dd4000cd5a3832"
	bizTransaction.UserEmail = "sprov300@gmail.com"
	bizTransaction.RequestId = "bks1m1g91jau4nkks2f0"

	bizTransactions := []*epcisproto.CreateBizTransactionRequest{}
	bizTransactions = append(bizTransactions, &bizTransaction)
	transactionEvent.BizTransactionList = bizTransactions

	quantityElement := epcisproto.CreateQuantityElementRequest{}
	quantityElement.EpcClass = "urn:epc:class:lgtin:4012345.011111.4444"
	quantityElement.Quantity = float64(10)
	quantityElement.Uom = "KGM"
	quantityElement.EventId = uint32(1)
	quantityElement.TypeOfEvent = "TransactionEvent"
	quantityElement.TypeOfQuantity = "output"
	quantityElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	quantityElement.UserEmail = "sprov300@gmail.com"
	quantityElement.RequestId = "bks1m1g91jau4nkks2f0"

	quantityElements := []*epcisproto.CreateQuantityElementRequest{}
	quantityElements = append(quantityElements, &quantityElement)
	transactionEvent.QuantityList = quantityElements

	source := epcisproto.CreateSourceRequest{}
	source.SourceType = "location"
	source.Source = "urn:epc:id:sgln:4012345.00225.0"
	source.EventId = uint32(1)
	source.TypeOfEvent = "TransactionEvent"
	source.UserId = "auth0|673ee1a719dd4000cd5a3832"
	source.UserEmail = "sprov300@gmail.com"
	source.RequestId = "bks1m1g91jau4nkks2f0"
	sources := []*epcisproto.CreateSourceRequest{}
	sources = append(sources, &source)
	transactionEvent.SourceList = sources

	destination := epcisproto.CreateDestinationRequest{}
	destination.DestType = "location"
	destination.Destination = "urn:epc:id:sgln:0614141.00777.0"
	destination.EventId = uint32(1)
	destination.TypeOfEvent = "TransactionEvent"
	destination.UserId = "auth0|673ee1a719dd4000cd5a3832"
	destination.UserEmail = "sprov300@gmail.com"
	destination.RequestId = "bks1m1g91jau4nkks2f0"

	destinations := []*epcisproto.CreateDestinationRequest{}
	destinations = append(destinations, &destination)
	transactionEvent.DestinationList = destinations

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
	sensorElement.TypeOfEvent = "TransactionEvent"
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
	transactionEvent.SensorElementList = sensorElements

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateTransactionEventRequest
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
				in:  &transactionEvent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionEventResp, err := tt.es.CreateTransactionEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateTransactionEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, transactionEventResp)
		transactionEventResult := transactionEventResp.TransactionEvent
		assert.Equal(t, transactionEventResult.EpcisEventD.EventId, "ni:///sha-256;45a99ca926fdb62b61bb2b29620e1dcdd5b0109613700f7e179881d64d8fabf1?ver=CBV2.0", "they should be equal")
		assert.Equal(t, transactionEventResult.EpcisEventD.EventTimeZoneOffset, "-06:00", "they should be equal")
		assert.Equal(t, transactionEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.BizStep, "shipping", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.Disposition, "in_transit", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.ReadPoint, "urn:epc:id:sgln:0614141.07346.1234", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func TestEpcisService_GetTransactionEvents(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	transactionEvents := []*epcisproto.TransactionEvent{}

	transactionEvent, err := GetTransactionEvent(uint32(1), []byte{76, 74, 219, 85, 196, 114, 68, 29, 177, 156, 62, 156, 10, 78, 135, 84}, "4c4adb55-c472-441d-b19c-3e9c0a4e8754", "ni:///sha-256;45a99ca926fdb62b61bb2b29620e1dcdd5b0109613700f7e179881d64d8fabf1?ver=CBV2.0", "-06:00", "https://accreditation-council.example.org/certificate/ABC12345", "2005-04-04T02:33:31Z", "incorrect_data", "2006-11-07T14:00:00Z", "urn:epc:id:sscc:0614141.1234567890", "ADD", "shipping", "in_transit", "urn:epc:id:sgln:0614141.07346.1234", "urn:epc:id:sgln:0614141.00888.0", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	transactionEvents = append(transactionEvents, transactionEvent)

	form := epcisproto.GetTransactionEventsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	transactionEventsResponse := epcisproto.GetTransactionEventsResponse{TransactionEvents: transactionEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *epcisproto.GetTransactionEventsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetTransactionEventsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &transactionEventsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionEventResp, err := tt.es.GetTransactionEvents(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetTransactionEvents() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transactionEventResp, tt.want) {
			t.Errorf("EpcisService.GetTransactionEvents() = %v, want %v", transactionEventResp, tt.want)
		}
		assert.NotNil(t, transactionEventResp)
		transactionEventResult := transactionEventResp.TransactionEvents[0]
		assert.Equal(t, transactionEventResult.EpcisEventD.EventId, "ni:///sha-256;45a99ca926fdb62b61bb2b29620e1dcdd5b0109613700f7e179881d64d8fabf1?ver=CBV2.0", "they should be equal")
		assert.Equal(t, transactionEventResult.EpcisEventD.EventTimeZoneOffset, "-06:00", "they should be equal")
		assert.Equal(t, transactionEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.BizStep, "shipping", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.Disposition, "in_transit", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.ReadPoint, "urn:epc:id:sgln:0614141.07346.1234", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func TestEpcisService_GetTransactionEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	transactionEvent, err := GetTransactionEvent(uint32(1), []byte{76, 74, 219, 85, 196, 114, 68, 29, 177, 156, 62, 156, 10, 78, 135, 84}, "4c4adb55-c472-441d-b19c-3e9c0a4e8754", "ni:///sha-256;45a99ca926fdb62b61bb2b29620e1dcdd5b0109613700f7e179881d64d8fabf1?ver=CBV2.0", "-06:00", "https://accreditation-council.example.org/certificate/ABC12345", "2005-04-04T02:33:31Z", "incorrect_data", "2006-11-07T14:00:00Z", "urn:epc:id:sscc:0614141.1234567890", "ADD", "shipping", "in_transit", "urn:epc:id:sgln:0614141.07346.1234", "urn:epc:id:sgln:0614141.00888.0", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	transactionEventResponse := epcisproto.GetTransactionEventResponse{}
	transactionEventResponse.TransactionEvent = transactionEvent

	form := epcisproto.GetTransactionEventRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "4c4adb55-c472-441d-b19c-3e9c0a4e8754"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *epcisproto.GetTransactionEventRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetTransactionEventResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &transactionEventResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionEventResp, err := tt.es.GetTransactionEvent(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetTransactionEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transactionEventResp, tt.want) {
			t.Errorf("EpcisService.GetTransactionEvent() = %v, want %v", transactionEventResp, tt.want)
		}
		assert.NotNil(t, transactionEventResp)
		transactionEventResult := transactionEventResp.TransactionEvent
		assert.Equal(t, transactionEventResult.EpcisEventD.EventId, "ni:///sha-256;45a99ca926fdb62b61bb2b29620e1dcdd5b0109613700f7e179881d64d8fabf1?ver=CBV2.0", "they should be equal")
		assert.Equal(t, transactionEventResult.EpcisEventD.EventTimeZoneOffset, "-06:00", "they should be equal")
		assert.Equal(t, transactionEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.BizStep, "shipping", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.Disposition, "in_transit", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.ReadPoint, "urn:epc:id:sgln:0614141.07346.1234", "they should be equal")
		assert.Equal(t, transactionEventResult.TransactionEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func GetTransactionEvent(id uint32, uuid4 []byte, idS string, eventId string, eventTimeZoneOffset string, certification string, eventTime string, reason string, declarationTime string, parentId string, action string, bizStep string, disposition string, readPoint string, bizLocation string, statusCode string, createdByUserId string, updatedByUserId string) (*epcisproto.TransactionEvent, error) {
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

	transactionEventD := epcisproto.TransactionEventD{}
	transactionEventD.Id = id
	transactionEventD.Uuid4 = uuid4
	transactionEventD.IdS = idS
	transactionEventD.ParentId = parentId
	transactionEventD.Action = action
	transactionEventD.BizStep = bizStep
	transactionEventD.Disposition = disposition
	transactionEventD.ReadPoint = readPoint
	transactionEventD.BizLocation = bizLocation

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = statusCode
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	transactionEvent := epcisproto.TransactionEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, TransactionEventD: &transactionEventD, CrUpdUser: &crUpdUser}

	return &transactionEvent, nil
}
