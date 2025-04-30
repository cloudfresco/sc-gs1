package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// Invoice - struct Invoice
type Invoice struct {
	*invoiceproto.InvoiceD
	*InvoiceT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// InvoiceT - struct InvoiceT
type InvoiceT struct {
	ActualDeliveryDate   time.Time `protobuf:"bytes,1,opt,name=actual_delivery_date,json=actualDeliveryDate,proto3" json:"actual_delivery_date,omitempty"`
	InvoicingPeriodBegin time.Time `protobuf:"bytes,2,opt,name=invoicing_period_begin,json=invoicingPeriodBegin,proto3" json:"invoicing_period_begin,omitempty"`
	InvoicingPeriodEnd   time.Time `protobuf:"bytes,3,opt,name=invoicing_period_end,json=invoicingPeriodEnd,proto3" json:"invoicing_period_end,omitempty"`
}

// InvoiceLineItem - struct InvoiceLineItem
type InvoiceLineItem struct {
	*invoiceproto.InvoiceLineItemD
	*InvoiceLineItemT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// InvoiceLineItemT - struct InvoiceLineItemT
type InvoiceLineItemT struct {
	TransferOfOwnershipDate         time.Time `protobuf:"bytes,1,opt,name=transfer_of_ownership_date,json=transferOfOwnershipDate,proto3" json:"transfer_of_ownership_date,omitempty"`
	ActualDeliveryDate              time.Time `protobuf:"bytes,2,opt,name=actual_delivery_date,json=actualDeliveryDate,proto3" json:"actual_delivery_date,omitempty"`
	ServicetimePeriodLineLevelBegin time.Time `protobuf:"bytes,3,opt,name=servicetime_period_line_level_begin,json=servicetimePeriodLineLevelBegin,proto3" json:"servicetime_period_line_level_begin,omitempty"`
	ServicetimePeriodLineLevelEnd   time.Time `protobuf:"bytes,4,opt,name=servicetime_period_line_level_end,json=servicetimePeriodLineLevelEnd,proto3" json:"servicetime_period_line_level_end,omitempty"`
}

// InvoiceTotal - struct InvoiceTotal
type InvoiceTotal struct {
	*invoiceproto.InvoiceTotalD
	*InvoiceTotalT
}

// InvoiceTotalT - struct InvoiceTotalT
type InvoiceTotalT struct {
	PrepaidAmountDate time.Time `protobuf:"bytes,1,opt,name=prepaid_amount_date,json=prepaidAmountDate,proto3" json:"prepaid_amount_date,omitempty"`
}

// LeviedDutyFeeTax - struct LeviedDutyFeeTax
type LeviedDutyFeeTax struct {
	*invoiceproto.LeviedDutyFeeTaxD
	*LeviedDutyFeeTaxT
}

// LeviedDutyFeeTaxT - struct LeviedDutyFeeTaxT
type LeviedDutyFeeTaxT struct {
	DutyFeeTaxPointDate time.Time `protobuf:"bytes,1,opt,name=duty_fee_tax_point_date,json=dutyFeeTaxPointDate,proto3" json:"duty_fee_tax_point_date,omitempty"`
}
