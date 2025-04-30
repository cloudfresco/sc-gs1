package invoiceservices

import (
	"context"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInvoiceAllowanceChargeSQL = `insert into invoice_allowance_charges
	    (levied_duty_fee_tax,
allowance_charge,
invoice_id)
  values(
:levied_duty_fee_tax,
:allowance_charge,
:invoice_id);`

/*const selectInvoiceAllowanceChargesSQL = `select
    id,
levied_duty_fee_tax,
allowance_charge,
invoice_id from invoice_allowance_charges`*/

// CreateInvoiceAllowanceCharge - Create InvoiceAllowanceCharge
func (invs *InvoiceService) CreateInvoiceAllowanceCharge(ctx context.Context, in *invoiceproto.CreateInvoiceAllowanceChargeRequest) (*invoiceproto.CreateInvoiceAllowanceChargeResponse, error) {
	invoiceAllowanceCharge, err := invs.ProcessInvoiceAllowanceChargeRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInvoiceAllowanceCharge(ctx, insertInvoiceAllowanceChargeSQL, invoiceAllowanceCharge, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceAllowanceChargeResponse := invoiceproto.CreateInvoiceAllowanceChargeResponse{}
	invoiceAllowanceChargeResponse.InvoiceAllowanceCharge = invoiceAllowanceCharge
	return &invoiceAllowanceChargeResponse, nil
}

// ProcessInvoiceAllowanceChargeRequest - ProcessInvoiceAllowanceChargeRequest
func (invs *InvoiceService) ProcessInvoiceAllowanceChargeRequest(ctx context.Context, in *invoiceproto.CreateInvoiceAllowanceChargeRequest) (*invoiceproto.InvoiceAllowanceCharge, error) {
	invoiceAllowanceCharge := invoiceproto.InvoiceAllowanceCharge{}

	invoiceAllowanceCharge.LeviedDutyFeeTax = in.LeviedDutyFeeTax
	invoiceAllowanceCharge.AllowanceCharge = in.AllowanceCharge
	invoiceAllowanceCharge.InvoiceId = in.InvoiceId
	return &invoiceAllowanceCharge, nil
}

// insertInvoiceAllowanceCharge - Insert InvoiceAllowanceCharge details into database
func (invs *InvoiceService) insertInvoiceAllowanceCharge(ctx context.Context, insertInvoiceAllowanceChargeSQL string, invoiceAllowanceCharge *invoiceproto.InvoiceAllowanceCharge, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInvoiceAllowanceChargeSQL, invoiceAllowanceCharge)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoiceAllowanceCharge.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
