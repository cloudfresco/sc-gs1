package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertTransactionalReferenceSQL = `insert into transactional_references
	    (transactional_reference_type_code,
ecom_document_reference)
  values(:transactional_reference_type_code,
:ecom_document_reference);`

/*const selectTransactionalReferencesSQL = `select
  id,
  transactional_reference_type_code,
  ecom_document_reference from transactional_references`*/

// CreateTransactionalReference - Create TransactionalReference
func (invs *InventoryService) CreateTransactionalReference(ctx context.Context, in *inventoryproto.CreateTransactionalReferenceRequest) (*inventoryproto.CreateTransactionalReferenceResponse, error) {
	transactionalReference, err := invs.ProcessTransactionalReferenceRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransactionalReference(ctx, insertTransactionalReferenceSQL, transactionalReference, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalReferenceResponse := inventoryproto.CreateTransactionalReferenceResponse{}
	transactionalReferenceResponse.TransactionalReference = transactionalReference
	return &transactionalReferenceResponse, nil
}

// ProcessTransactionalReferenceRequest - ProcessTransactionalReferenceRequest
func (invs *InventoryService) ProcessTransactionalReferenceRequest(ctx context.Context, in *inventoryproto.CreateTransactionalReferenceRequest) (*inventoryproto.TransactionalReference, error) {
	transactionalReference := inventoryproto.TransactionalReference{}
	transactionalReference.TransactionalReferenceTypeCode = in.TransactionalReferenceTypeCode
	transactionalReference.EcomDocumentReference = in.EcomDocumentReference

	return &transactionalReference, nil
}

// insertTransactionalReference - Insert TransactionalReference details into database
func (invs *InventoryService) insertTransactionalReference(ctx context.Context, insertTransactionalReferenceSQL string, transactionalReference *inventoryproto.TransactionalReference, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalReferenceSQL, transactionalReference)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalReference.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
