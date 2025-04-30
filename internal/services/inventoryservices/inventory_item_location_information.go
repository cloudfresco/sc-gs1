package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInventoryItemLocationInformationSQL = `insert into inventory_item_location_informations
(inventory_location_id,
inventory_report_id) 
values(
  :inventory_location_id,
  :inventory_report_id
);`

/*const selectInventoryItemLocationInformationsSQL = `select
  id,
  inventory_location_id,
  inventory_report_id from inventory_item_location_informations`*/

// CreateInventoryItemLocationInformation - Create InventoryItemLocationInformation
func (invs *InventoryService) CreateInventoryItemLocationInformation(ctx context.Context, in *inventoryproto.CreateInventoryItemLocationInformationRequest) (*inventoryproto.CreateInventoryItemLocationInformationResponse, error) {
	inventoryItemLocationInformation, err := invs.ProcessInventoryItemLocationInformationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInventoryItemLocationInformation(ctx, insertInventoryItemLocationInformationSQL, inventoryItemLocationInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryItemLocationInformationResponse := inventoryproto.CreateInventoryItemLocationInformationResponse{}
	inventoryItemLocationInformationResponse.InventoryItemLocationInformation = inventoryItemLocationInformation
	return &inventoryItemLocationInformationResponse, nil
}

// ProcessInventoryItemLocationInformationRequest - ProcessInventoryItemLocationInformationRequest
func (invs *InventoryService) ProcessInventoryItemLocationInformationRequest(ctx context.Context, in *inventoryproto.CreateInventoryItemLocationInformationRequest) (*inventoryproto.InventoryItemLocationInformation, error) {
	inventoryItemLocationInformation := inventoryproto.InventoryItemLocationInformation{}

	inventoryItemLocationInformation.InventoryLocationId = in.InventoryLocationId
	inventoryItemLocationInformation.InventoryReportId = in.InventoryReportId
	return &inventoryItemLocationInformation, nil
}

// insertInventoryItemLocationInformation - Insert InventoryItemLocationInformation details into database
func (invs *InventoryService) insertInventoryItemLocationInformation(ctx context.Context, insertInventoryItemLocationInformationSQL string, inventoryItemLocationInformation *inventoryproto.InventoryItemLocationInformation, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInventoryItemLocationInformationSQL, inventoryItemLocationInformation)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryItemLocationInformation.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
