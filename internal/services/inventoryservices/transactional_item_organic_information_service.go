package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertTransactionalItemOrganicInformationSQL = `insert into transactional_item_organic_informations
	    (is_trade_item_organic,
transactional_item_data_id)
  values(
 :is_trade_item_organic,
 :transactional_item_data_id);`

/*const selectTransactionalItemOrganicInformationsSQL = `select
  id,
  is_trade_item_organic,
  transactional_item_data_id from transactional_item_organic_informations`*/

// CreateTransactionalItemOrganicInformation - Create TransactionalItemOrganicInformation
func (invs *InventoryService) CreateTransactionalItemOrganicInformation(ctx context.Context, in *inventoryproto.CreateTransactionalItemOrganicInformationRequest) (*inventoryproto.CreateTransactionalItemOrganicInformationResponse, error) {
	transactionalItemOrganicInformation, err := invs.ProcessTransactionalItemOrganicInformationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransactionalItemOrganicInformation(ctx, insertTransactionalItemOrganicInformationSQL, transactionalItemOrganicInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalItemOrganicInformationResponse := inventoryproto.CreateTransactionalItemOrganicInformationResponse{}
	transactionalItemOrganicInformationResponse.TransactionalItemOrganicInformation = transactionalItemOrganicInformation
	return &transactionalItemOrganicInformationResponse, nil
}

// ProcessTransactionalItemOrganicInformationRequest - ProcessTransactionalItemOrganicInformationRequest
func (invs *InventoryService) ProcessTransactionalItemOrganicInformationRequest(ctx context.Context, in *inventoryproto.CreateTransactionalItemOrganicInformationRequest) (*inventoryproto.TransactionalItemOrganicInformation, error) {
	transactionalItemOrganicInformation := inventoryproto.TransactionalItemOrganicInformation{}
	transactionalItemOrganicInformation.IsTradeItemOrganic = in.IsTradeItemOrganic
	transactionalItemOrganicInformation.TransactionalItemDataId = in.TransactionalItemDataId
	return &transactionalItemOrganicInformation, nil
}

// insertTransactionalItemOrganicInformation - Insert TransactionalItemOrganicInformation details into database
func (invs *InventoryService) insertTransactionalItemOrganicInformation(ctx context.Context, insertTransactionalItemOrganicInformationSQL string, transactionalItemOrganicInformation *inventoryproto.TransactionalItemOrganicInformation, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalItemOrganicInformationSQL, transactionalItemOrganicInformation)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemOrganicInformation.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
