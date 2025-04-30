package v1

import (
	"time"
)

// CrUpdTime - struct CrUpdTime
type CrUpdTime struct {
	CreatedAt time.Time `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt time.Time `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

// EpcisEventT - struct EpcisEventT
type EpcisEventT struct {
	EventTime time.Time `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
}

// ErrorDeclarationT - struct ErrorDeclarationT
type ErrorDeclarationT struct {
	DeclarationTime time.Time `protobuf:"bytes,1,opt,name=declaration_time,json=declarationTime,proto3" json:"declaration_time,omitempty"`
}

// SensorMetadataT - struct SensorMetadataT
type SensorMetadataT struct {
	SensorTime time.Time `protobuf:"bytes,1,opt,name=sensor_time,json=sensorTime,proto3" json:"sensor_time,omitempty"`
	StartTime  time.Time `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    time.Time `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}
