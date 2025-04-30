package v1

import (
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// DebitCreditAdvice - struct DebitCreditAdvice
type DebitCreditAdvice struct {
	*invoiceproto.DebitCreditAdviceD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// DebitCreditAdviceLineItem - struct DebitCreditAdviceLineItem
type DebitCreditAdviceLineItem struct {
	*invoiceproto.DebitCreditAdviceLineItemD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
