package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// DespatchAdvice - struct DespatchAdvice
type DespatchAdvice struct {
	*logisticsproto.DespatchAdviceD
	*DespatchAdviceT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// DespatchAdviceT - struct DespatchAdviceT
type DespatchAdviceT struct {
	DespatchAdviceDeliveryDateTimeBegin time.Time `protobuf:"bytes,1,opt,name=despatch_advice_delivery_date_time_begin,json=despatchAdviceDeliveryDateTimeBegin,proto3" json:"despatch_advice_delivery_date_time_begin,omitempty"`
	DespatchAdviceDeliveryDateTimeEnd   time.Time `protobuf:"bytes,2,opt,name=despatch_advice_delivery_date_time_end,json=despatchAdviceDeliveryDateTimeEnd,proto3" json:"despatch_advice_delivery_date_time_end,omitempty"`
	PaymentDateTimeBegin                time.Time `protobuf:"bytes,3,opt,name=payment_date_time_begin,json=paymentDateTimeBegin,proto3" json:"payment_date_time_begin,omitempty"`
	PaymentDateTimeEnd                  time.Time `protobuf:"bytes,4,opt,name=payment_date_time_end,json=paymentDateTimeEnd,proto3" json:"payment_date_time_end,omitempty"`
	ReceivingDateTimeBegin              time.Time `protobuf:"bytes,5,opt,name=receiving_date_time_begin,json=receivingDateTimeBegin,proto3" json:"receiving_date_time_begin,omitempty"`
	ReceivingDateTimeEnd                time.Time `protobuf:"bytes,6,opt,name=receiving_date_time_end,json=receivingDateTimeEnd,proto3" json:"receiving_date_time_end,omitempty"`
}

// DespatchAdviceLineItem - struct DespatchAdviceLineItem
type DespatchAdviceLineItem struct {
	*logisticsproto.DespatchAdviceLineItemD
	*DespatchAdviceLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// DespatchAdviceLineItemT - struct DespatchAdviceLineItemT
type DespatchAdviceLineItemT struct {
	FirstInFirstOutDateTime time.Time `protobuf:"bytes,1,opt,name=first_in_first_out_date_time,json=firstInFirstOutDateTime,proto3" json:"first_in_first_out_date_time,omitempty"`
	PickUpDateTime          time.Time `protobuf:"bytes,2,opt,name=pick_up_date_time,json=pickUpDateTime,proto3" json:"pick_up_date_time,omitempty"`
}

// DespatchAdviceLogisticUnit - struct DespatchAdviceLogisticUnit
type DespatchAdviceLogisticUnit struct {
	*logisticsproto.DespatchAdviceLogisticUnitD
	*DespatchAdviceLogisticUnitT
}

// DespatchAdviceLogisticUnitT - struct DespatchAdviceLogisticUnitT
type DespatchAdviceLogisticUnitT struct {
	EstimatedDeliveryDateTimeAtUltimateConsignee time.Time `protobuf:"bytes,1,opt,name=estimated_delivery_date_time_at_ultimate_consignee,json=estimatedDeliveryDateTimeAtUltimateConsignee,proto3" json:"estimated_delivery_date_time_at_ultimate_consignee,omitempty"`
}

// DespatchAdviceQuantityVariance - struct DespatchAdviceQuantityVariance
type DespatchAdviceQuantityVariance struct {
	*logisticsproto.DespatchAdviceQuantityVarianceD
	*DespatchAdviceQuantityVarianceT
}

// DespatchAdviceQuantityVarianceT - struct DespatchAdviceQuantityVarianceT
type DespatchAdviceQuantityVarianceT struct {
	DeliveryDateVariance time.Time `protobuf:"bytes,1,opt,name=delivery_date_variance,json=deliveryDateVariance,proto3" json:"delivery_date_variance,omitempty"`
}

// DespatchInformation - struct DespatchInformation
type DespatchInformation struct {
	*logisticsproto.DespatchInformationD
	*DespatchInformationT
}

// DespatchInformationT - struct DespatchInformationT
type DespatchInformationT struct {
	ActualShipDateTime                           time.Time `protobuf:"bytes,1,opt,name=actual_ship_date_time,json=actualShipDateTime,proto3" json:"actual_ship_date_time,omitempty"`
	DespatchDateTime                             time.Time `protobuf:"bytes,2,opt,name=despatch_date_time,json=despatchDateTime,proto3" json:"despatch_date_time,omitempty"`
	EstimatedDeliveryDateTime                    time.Time `protobuf:"bytes,3,opt,name=estimated_delivery_date_time,json=estimatedDeliveryDateTime,proto3" json:"estimated_delivery_date_time,omitempty"`
	EstimatedDeliveryDateTimeAtUltimateConsignee time.Time `protobuf:"bytes,4,opt,name=estimated_delivery_date_time_at_ultimate_consignee,json=estimatedDeliveryDateTimeAtUltimateConsignee,proto3" json:"estimated_delivery_date_time_at_ultimate_consignee,omitempty"`
	LoadingDateTime                              time.Time `protobuf:"bytes,5,opt,name=loading_date_time,json=loadingDateTime,proto3" json:"loading_date_time,omitempty"`
	PickUpDateTime                               time.Time `protobuf:"bytes,6,opt,name=pick_up_date_time,json=pickUpDateTime,proto3" json:"pick_up_date_time,omitempty"`
	ReleaseDateTimeOfSupplier                    time.Time `protobuf:"bytes,7,opt,name=release_date_time_of_supplier,json=releaseDateTimeOfSupplier,proto3" json:"release_date_time_of_supplier,omitempty"`
	EstimatedDeliveryPeriodBegin                 time.Time `protobuf:"bytes,8,opt,name=estimated_delivery_period_begin,json=estimatedDeliveryPeriodBegin,proto3" json:"estimated_delivery_period_begin,omitempty"`
	EstimatedDeliveryPeriodEnd                   time.Time `protobuf:"bytes,9,opt,name=estimated_delivery_period_end,json=estimatedDeliveryPeriodEnd,proto3" json:"estimated_delivery_period_end,omitempty"`
}
