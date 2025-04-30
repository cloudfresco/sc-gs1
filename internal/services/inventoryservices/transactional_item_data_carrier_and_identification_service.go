package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertTransactionalItemDataCarrierAndIdentificationSQL = `insert into transactional_item_data_carrier_identifications
	    (data_carrier,
gs1_transactional_item_identification_key,
transactional_item_data_id)
  values(
 :data_carrier,
 :gs1_transactional_item_identification_key,
 :transactional_item_data_id
);`

/*const selectTransactionalItemDataCarrierAndIdentificationsSQL = `select
  id,
  data_carrier,
  gs1_transactional_item_identification_key,
  transactional_item_data_id from transactional_item_data_carrier_identifications`*/

// CreateTransactionalItemDataCarrierAndIdentification - Create TransactionalItemDataCarrierAndIdentification
func (invs *InventoryService) CreateTransactionalItemDataCarrierAndIdentification(ctx context.Context, in *inventoryproto.CreateTransactionalItemDataCarrierAndIdentificationRequest) (*inventoryproto.CreateTransactionalItemDataCarrierAndIdentificationResponse, error) {
	transactionalItemDataCarrierAndIdentification, err := invs.ProcessTransactionalItemDataCarrierAndIdentificationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransactionalItemDataCarrierAndIdentification(ctx, insertTransactionalItemDataCarrierAndIdentificationSQL, transactionalItemDataCarrierAndIdentification, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalItemDataCarrierAndIdentificationResponse := inventoryproto.CreateTransactionalItemDataCarrierAndIdentificationResponse{}
	transactionalItemDataCarrierAndIdentificationResponse.TransactionalItemDataCarrierAndIdentification = transactionalItemDataCarrierAndIdentification
	return &transactionalItemDataCarrierAndIdentificationResponse, nil
}

// ProcessTransactionalItemDataCarrierAndIdentificationRequest - ProcessTransactionalItemDataCarrierAndIdentificationRequest
func (invs *InventoryService) ProcessTransactionalItemDataCarrierAndIdentificationRequest(ctx context.Context, in *inventoryproto.CreateTransactionalItemDataCarrierAndIdentificationRequest) (*inventoryproto.TransactionalItemDataCarrierAndIdentification, error) {
	transactionalItemDataCarrierAndIdentification := inventoryproto.TransactionalItemDataCarrierAndIdentification{}
	transactionalItemDataCarrierAndIdentification.DataCarrier = in.DataCarrier
	transactionalItemDataCarrierAndIdentification.Gs1TransactionalItemIdentificationKey = in.Gs1TransactionalItemIdentificationKey
	transactionalItemDataCarrierAndIdentification.TransactionalItemDataId = in.TransactionalItemDataId
	return &transactionalItemDataCarrierAndIdentification, nil
}

// insertTransactionalItemDataCarrierAndIdentification - Insert TransactionalItemDataCarrierAndIdentification details into database
func (invs *InventoryService) insertTransactionalItemDataCarrierAndIdentification(ctx context.Context, insertTransactionalItemDataCarrierAndIdentificationSQL string, transactionalItemDataCarrierAndIdentification *inventoryproto.TransactionalItemDataCarrierAndIdentification, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalItemDataCarrierAndIdentificationSQL, transactionalItemDataCarrierAndIdentification)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemDataCarrierAndIdentification.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
