package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// ConsumptionReportLineItem - struct ConsumptionReportLineItem
type ConsumptionReportLineItem struct {
	*inventoryproto.ConsumptionReportLineItemD
	*ConsumptionReportLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ConsumptionReportLineItemT - struct ConsumptionReportLineItemT
type ConsumptionReportLineItemT struct {
	ConsumptionPeriodBegin time.Time `protobuf:"bytes,1,opt,name=consumption_period_begin,json=consumptionPeriodBegin,proto3" json:"consumption_period_begin,omitempty"`
	ConsumptionPeriodEnd   time.Time `protobuf:"bytes,2,opt,name=consumption_period_end,json=consumptionPeriodEnd,proto3" json:"consumption_period_end,omitempty"`
}

// ConsumptionReport - struct ConsumptionReport
type ConsumptionReport struct {
	*inventoryproto.ConsumptionReportD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// InventoryActivityLineItem - struct InventoryActivityLineItem
type InventoryActivityLineItem struct {
	*inventoryproto.InventoryActivityLineItemD
	*InventoryActivityLineItemT
}

// InventoryActivityLineItemT - struct InventoryActivityLineItemT
type InventoryActivityLineItemT struct {
	ReportingPeriodBegin time.Time `protobuf:"bytes,1,opt,name=reporting_period_begin,json=reportingPeriodBegin,proto3" json:"reporting_period_begin,omitempty"`
	ReportingPeriodEnd   time.Time `protobuf:"bytes,2,opt,name=reporting_period_end,json=reportingPeriodEnd,proto3" json:"reporting_period_end,omitempty"`
}

// InventoryReport - struct InventoryReport
type InventoryReport struct {
	*inventoryproto.InventoryReportD
	*InventoryReportT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// InventoryReportT - struct InventoryReportT
type InventoryReportT struct {
	ReportingPeriodBegin time.Time `protobuf:"bytes,1,opt,name=reporting_period_begin,json=reportingPeriodBegin,proto3" json:"reporting_period_begin,omitempty"`
	ReportingPeriodEnd   time.Time `protobuf:"bytes,2,opt,name=reporting_period_end,json=reportingPeriodEnd,proto3" json:"reporting_period_end,omitempty"`
}

// InventoryStatusLineItem - struct InventoryStatusLineItem
type InventoryStatusLineItem struct {
	*inventoryproto.InventoryStatusLineItemD
	*InventoryStatusLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// InventoryStatusLineItemT - struct InventoryStatusLineItemT
type InventoryStatusLineItemT struct {
	FirstInFirstOutDateTimeBegin time.Time `protobuf:"bytes,1,opt,name=first_in_first_out_date_time_begin,json=firstInFirstOutDateTimeBegin,proto3" json:"first_in_first_out_date_time_begin,omitempty"`
	FirstInFirstOutDateTimeEnd   time.Time `protobuf:"bytes,2,opt,name=first_in_first_out_date_time_end,json=firstInFirstOutDateTimeEnd,proto3" json:"first_in_first_out_date_time_end,omitempty"`
	InventoryDateTimeBegin       time.Time `protobuf:"bytes,3,opt,name=inventory_date_time_begin,json=inventoryDateTimeBegin,proto3" json:"inventory_date_time_begin,omitempty"`
	InventoryDateTimeEnd         time.Time `protobuf:"bytes,4,opt,name=inventory_date_time_end,json=inventoryDateTimeEnd,proto3" json:"inventory_date_time_end,omitempty"`
	ReportingPeriodBegin         time.Time `protobuf:"bytes,5,opt,name=reporting_period_begin,json=reportingPeriodBegin,proto3" json:"reporting_period_begin,omitempty"`
	ReportingPeriodEnd           time.Time `protobuf:"bytes,6,opt,name=reporting_period_end,json=reportingPeriodEnd,proto3" json:"reporting_period_end,omitempty"`
}

// LogisticUnitInventoryEvent - struct LogisticUnitInventoryEvent
type LogisticUnitInventoryEvent struct {
	*inventoryproto.LogisticUnitInventoryEventD
	*LogisticUnitInventoryEventT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// LogisticUnitInventoryEventT - struct LogisticUnitInventoryEventT
type LogisticUnitInventoryEventT struct {
	EventDateTime time.Time `protobuf:"bytes,1,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}

// LogisticsInventoryReport - struct LogisticsInventoryReport
type LogisticsInventoryReport struct {
	*inventoryproto.LogisticsInventoryReportD
	*LogisticsInventoryReportT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// LogisticsInventoryReportT - struct LogisticsInventoryReportT
type LogisticsInventoryReportT struct {
	ReportingPeriodBegin time.Time `protobuf:"bytes,1,opt,name=reporting_period_begin,json=reportingPeriodBegin,proto3" json:"reporting_period_begin,omitempty"`
	ReportingPeriodEnd   time.Time `protobuf:"bytes,2,opt,name=reporting_period_end,json=reportingPeriodEnd,proto3" json:"reporting_period_end,omitempty"`
}

// LogisticUnitInventoryStatus - struct LogisticUnitInventoryStatus
type LogisticUnitInventoryStatus struct {
	*inventoryproto.LogisticUnitInventoryStatusD
	*LogisticUnitInventoryStatusT
}

// LogisticUnitInventoryStatusT - struct LogisticUnitInventoryStatusT
type LogisticUnitInventoryStatusT struct {
	InventoryDateTime time.Time `protobuf:"bytes,1,opt,name=inventory_date_time,json=inventoryDateTime,proto3" json:"inventory_date_time,omitempty"`
}

// ReturnablePackagingInventoryEvent - struct ReturnablePackagingInventoryEvent
type ReturnablePackagingInventoryEvent struct {
	*inventoryproto.ReturnablePackagingInventoryEventD
	*ReturnablePackagingInventoryEventT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ReturnablePackagingInventoryEventT - struct ReturnablePackagingInventoryEventT
type ReturnablePackagingInventoryEventT struct {
	EventDateTime time.Time `protobuf:"bytes,1,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}

// TradeItemInventoryEvent - struct TradeItemInventoryEvent
type TradeItemInventoryEvent struct {
	*inventoryproto.TradeItemInventoryEventD
	*TradeItemInventoryEventT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// TradeItemInventoryEventT - struct TradeItemInventoryEventT
type TradeItemInventoryEventT struct {
	EventDateTime time.Time `protobuf:"bytes,1,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}

// TradeItemInventoryStatus  - struct TradeItemInventoryStatus
type TradeItemInventoryStatus struct {
	*inventoryproto.TradeItemInventoryStatusD
	*TradeItemInventoryStatusT
}

// TradeItemInventoryStatusT  - struct TradeItemInventoryStatusT
type TradeItemInventoryStatusT struct {
	InventoryDateTime time.Time `protobuf:"bytes,1,opt,name=inventory_date_time,json=inventoryDateTime,proto3" json:"inventory_date_time,omitempty"`
}

// TransactionalItemData - struct TransactionalItemData
type TransactionalItemData struct {
	*inventoryproto.TransactionalItemDataD
	*TransactionalItemDataT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// TransactionalItemDataT - struct TransactionalItemDataT
type TransactionalItemDataT struct {
	AvailableForSaleDate time.Time `protobuf:"bytes,1,opt,name=available_for_sale_date,json=availableForSaleDate,proto3" json:"available_for_sale_date,omitempty"`
	BestBeforeDate       time.Time `protobuf:"bytes,2,opt,name=best_before_date,json=bestBeforeDate,proto3" json:"best_before_date,omitempty"`
	ItemExpirationDate   time.Time `protobuf:"bytes,3,opt,name=item_expiration_date,json=itemExpirationDate,proto3" json:"item_expiration_date,omitempty"`
	PackagingDate        time.Time `protobuf:"bytes,4,opt,name=packaging_date,json=packagingDate,proto3" json:"packaging_date,omitempty"`
	ProductionDate       time.Time `protobuf:"bytes,5,opt,name=production_date,json=productionDate,proto3" json:"production_date,omitempty"`
	SellByDate           time.Time `protobuf:"bytes,6,opt,name=sell_by_date,json=sellByDate,proto3" json:"sell_by_date,omitempty"`
}

// TransactionalItemLogisticUnitInformation - struct TransactionalItemLogisticUnitInformation
type TransactionalItemLogisticUnitInformation struct {
	*inventoryproto.TransactionalItemLogisticUnitInformationD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// TransportEquipmentInventoryEvent - struct TransportEquipmentInventoryEvent
type TransportEquipmentInventoryEvent struct {
	*inventoryproto.TransportEquipmentInventoryEventD
	*TransportEquipmentInventoryEventT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// TransportEquipmentInventoryEventT - struct TransportEquipmentInventoryEventT
type TransportEquipmentInventoryEventT struct {
	EventDateTime time.Time `protobuf:"bytes,1,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}

// TransportEquipmentInventoryStatus - struct TransportEquipmentInventoryStatus
type TransportEquipmentInventoryStatus struct {
	*inventoryproto.TransportEquipmentInventoryStatusD
	*TransportEquipmentInventoryStatusT
}

// TransportEquipmentInventoryStatusT - struct TransportEquipmentInventoryStatusT
type TransportEquipmentInventoryStatusT struct {
	InventoryDateTime time.Time `protobuf:"bytes,1,opt,name=inventory_date_time,json=inventoryDateTime,proto3" json:"inventory_date_time,omitempty"`
}
