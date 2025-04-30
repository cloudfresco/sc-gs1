package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// Order - struct Order
type Order struct {
	*orderproto.OrderD
	*OrderT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// OrderT - struct OrderT
type OrderT struct {
	DeliveryDateAccordingToSchedule time.Time `protobuf:"bytes,1,opt,name=delivery_date_according_to_schedule,json=deliveryDateAccordingToSchedule,proto3" json:"delivery_date_according_to_schedule,omitempty"`
	LatestDeliveryDate              time.Time `protobuf:"bytes,2,opt,name=latest_delivery_date,json=latestDeliveryDate,proto3" json:"latest_delivery_date,omitempty"`
}

// OrderLineItem - struct OrderLineItem
type OrderLineItem struct {
	*orderproto.OrderLineItemD
	*OrderLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// OrderLineItemT - struct OrderLineItemT
type OrderLineItemT struct {
	LatestDeliveryDate time.Time `protobuf:"bytes,1,opt,name=latest_delivery_date,json=latestDeliveryDate,proto3" json:"latest_delivery_date,omitempty"`
}

// OrderLogisticalDateInformation - struct OrderLogisticalDateInformation
type OrderLogisticalDateInformation struct {
	*orderproto.OrderLogisticalDateInformationD
	*OrderLogisticalDateInformationT
}

// OrderLogisticalDateInformationT - struct OrderLogisticalDateInformation
type OrderLogisticalDateInformationT struct {
	RequestedDeliveryDateRangeBegin                    time.Time `protobuf:"bytes,1,opt,name=requested_delivery_date_range_begin,json=requestedDeliveryDateRangeBegin,proto3" json:"requested_delivery_date_range_begin,omitempty"`
	RequestedDeliveryDateRangeEnd                      time.Time `protobuf:"bytes,2,opt,name=requested_delivery_date_range_end,json=requestedDeliveryDateRangeEnd,proto3" json:"requested_delivery_date_range_end,omitempty"`
	RequestedDeliveryDateRangeAtUltimateConsigneeBegin time.Time `protobuf:"bytes,3,opt,name=requested_delivery_date_range_at_ultimate_consignee_begin,json=requestedDeliveryDateRangeAtUltimateConsigneeBegin,proto3" json:"requested_delivery_date_range_at_ultimate_consignee_begin,omitempty"`
	RequestedDeliveryDateRangeAtUltimateConsigneeEnd   time.Time `protobuf:"bytes,4,opt,name=requested_delivery_date_range_at_ultimate_consignee_end,json=requestedDeliveryDateRangeAtUltimateConsigneeEnd,proto3" json:"requested_delivery_date_range_at_ultimate_consignee_end,omitempty"`
	RequestedDeliveryDateTime                          time.Time `protobuf:"bytes,5,opt,name=requested_delivery_date_time,json=requestedDeliveryDateTime,proto3" json:"requested_delivery_date_time,omitempty"`
	RequestedDeliveryDateTimeAtUltimateConsignee       time.Time `protobuf:"bytes,6,opt,name=requested_delivery_date_time_at_ultimate_consignee,json=requestedDeliveryDateTimeAtUltimateConsignee,proto3" json:"requested_delivery_date_time_at_ultimate_consignee,omitempty"`
	RequestedPickUpDateTime                            time.Time `protobuf:"bytes,7,opt,name=requested_pick_up_date_time,json=requestedPickUpDateTime,proto3" json:"requested_pick_up_date_time,omitempty"`
	RequestedShipDateRangeBegin                        time.Time `protobuf:"bytes,8,opt,name=requested_ship_date_range_begin,json=requestedShipDateRangeBegin,proto3" json:"requested_ship_date_range_begin,omitempty"`
	RequestedShipDateRangeEnd                          time.Time `protobuf:"bytes,9,opt,name=requested_ship_date_range_end,json=requestedShipDateRangeEnd,proto3" json:"requested_ship_date_range_end,omitempty"`
	RequestedShipDateTime                              time.Time `protobuf:"bytes,10,opt,name=requested_ship_date_time,json=requestedShipDateTime,proto3" json:"requested_ship_date_time,omitempty"`
}
