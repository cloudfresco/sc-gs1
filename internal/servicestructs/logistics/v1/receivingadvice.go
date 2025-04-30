package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// ReceivingAdvice - struct ReceivingAdvice
type ReceivingAdvice struct {
	*logisticsproto.ReceivingAdviceD
	*ReceivingAdviceT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ReceivingAdviceT - struct ReceivingAdviceT
type ReceivingAdviceT struct {
	DespatchAdviceDeliveryDateTimeBegin time.Time `protobuf:"bytes,1,opt,name=despatch_advice_delivery_date_time_begin,json=despatchAdviceDeliveryDateTimeBegin,proto3" json:"despatch_advice_delivery_date_time_begin,omitempty"`
	DespatchAdviceDeliveryDateTimeEnd   time.Time `protobuf:"bytes,2,opt,name=despatch_advice_delivery_date_time_end,json=despatchAdviceDeliveryDateTimeEnd,proto3" json:"despatch_advice_delivery_date_time_end,omitempty"`
	PaymentDateTimeBegin                time.Time `protobuf:"bytes,3,opt,name=payment_date_time_begin,json=paymentDateTimeBegin,proto3" json:"payment_date_time_begin,omitempty"`
	PaymentDateTimeEnd                  time.Time `protobuf:"bytes,4,opt,name=payment_date_time_end,json=paymentDateTimeEnd,proto3" json:"payment_date_time_end,omitempty"`
	ReceivingDateTimeBegin              time.Time `protobuf:"bytes,5,opt,name=receiving_date_time_begin,json=receivingDateTimeBegin,proto3" json:"receiving_date_time_begin,omitempty"`
	ReceivingDateTimeEnd                time.Time `protobuf:"bytes,6,opt,name=receiving_date_time_end,json=receivingDateTimeEnd,proto3" json:"receiving_date_time_end,omitempty"`
}

// ReceivingAdviceLineItem - struct ReceivingAdviceLineItem
type ReceivingAdviceLineItem struct {
	*logisticsproto.ReceivingAdviceLineItemD
	*ReceivingAdviceLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ReceivingAdviceLineItemT - struct ReceivingAdviceLineItemT
type ReceivingAdviceLineItemT struct {
	PickUpDateTimeBegin time.Time `protobuf:"bytes,1,opt,name=pick_up_date_time_begin,json=pickUpDateTimeBegin,proto3" json:"pick_up_date_time_begin,omitempty"`
	PickUpDateTimeEnd   time.Time `protobuf:"bytes,2,opt,name=pick_up_date_time_end,json=pickUpDateTimeEnd,proto3" json:"pick_up_date_time_end,omitempty"`
}
