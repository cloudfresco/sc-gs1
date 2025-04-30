package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// ObjectEvent - struct ObjectEvent
type ObjectEvent struct {
	*commonproto.EpcisEventD
	*commonstruct.EpcisEventT
	*commonproto.ErrorDeclarationD
	*commonstruct.ErrorDeclarationT
	*epcisproto.ObjectEventD
	*commonproto.CrUpdUser
}

// AssociationEvent - struct AssociationEvent
type AssociationEvent struct {
	*commonproto.EpcisEventD
	*commonstruct.EpcisEventT
	*commonproto.ErrorDeclarationD
	*commonstruct.ErrorDeclarationT
	*epcisproto.AssociationEventD
	*commonproto.CrUpdUser
}

// AggregationEvent - struct AggregationEvent
type AggregationEvent struct {
	*commonproto.EpcisEventD
	*commonstruct.EpcisEventT
	*commonproto.ErrorDeclarationD
	*commonstruct.ErrorDeclarationT
	*epcisproto.AggregationEventD
	*commonproto.CrUpdUser
}

// TransformationEvent - struct TransformationEvent
type TransformationEvent struct {
	*commonproto.EpcisEventD
	*commonstruct.EpcisEventT
	*commonproto.ErrorDeclarationD
	*commonstruct.ErrorDeclarationT
	*epcisproto.TransformationEventD
	*commonproto.CrUpdUser
}

// TransactionEvent - struct TransactionEvent
type TransactionEvent struct {
	*commonproto.EpcisEventD
	*commonstruct.EpcisEventT
	*commonproto.ErrorDeclarationD
	*commonstruct.ErrorDeclarationT
	*epcisproto.TransactionEventD
	*commonproto.CrUpdUser
}

type SensorElement struct {
	*commonproto.SensorMetadataD
	*commonstruct.SensorMetadataT
	*epcisproto.SensorElementD
}

type SensorReport struct {
	*epcisproto.SensorReportD
	*SensorReportT
}

type SensorReportT struct {
	SensorReportTime time.Time `protobuf:"bytes,1,opt,name=sensor_report_time,json=sensorReportTime,proto3" json:"sensor_report_time,omitempty"`
}
