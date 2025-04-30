package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// OrderResponse - struct OrderResponse
type OrderResponse struct {
	*orderproto.OrderResponseD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// OrderResponseLineItem - struct OrderResponseLineItem
type OrderResponseLineItem struct {
	*orderproto.OrderResponseLineItemD
	*OrderResponseLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// OrderResponseLineItemT - struct OrderResponseLineItemT
type OrderResponseLineItemT struct {
	DeliveryDateTime time.Time `protobuf:"bytes,1,opt,name=delivery_date_time,json=deliveryDateTime,proto3" json:"delivery_date_time,omitempty"`
}
