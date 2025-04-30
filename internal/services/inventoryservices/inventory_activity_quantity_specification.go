package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInventoryActivityQuantitySpecificationSQL = `insert into inventory_activity_quantity_specifications
	    (inventory_activity_type_code,
      inventory_movement_type_code,
      quantity_of_units,
      qou_measurement_unit_code,
      qou_code_list_version,
      inventory_status_line_item_id,
      inventory_activity_line_item_id,
      inventory_report_id)
  values(
  :inventory_activity_type_code,
  :inventory_movement_type_code,
  :quantity_of_units,
  :qou_measurement_unit_code,
  :qou_code_list_version,
  :inventory_status_line_item_id,
  :inventory_activity_line_item_id,
  :inventory_report_id);`

/*const selectInventoryActivityQuantitySpecificationsSQL = `select
  id,
  inventory_activity_type_code,
  inventory_movement_type_code,
  quantity_of_units,
  qou_measurement_unit_code,
  qou_code_list_version,
  inventory_status_line_item_id,
  inventory_activity_line_item_id,
  inventory_report_id from inventory_activity_quantity_specifications`*/

// CreateInventoryActivityQuantitySpecification - Create InventoryActivityQuantitySpecification
func (invs *InventoryService) CreateInventoryActivityQuantitySpecification(ctx context.Context, in *inventoryproto.CreateInventoryActivityQuantitySpecificationRequest) (*inventoryproto.CreateInventoryActivityQuantitySpecificationResponse, error) {
	inventoryActivityQuantitySpecification, err := invs.ProcessInventoryActivityQuantitySpecificationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInventoryActivityQuantitySpecification(ctx, insertInventoryActivityQuantitySpecificationSQL, inventoryActivityQuantitySpecification, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryActivityQuantitySpecificationResponse := inventoryproto.CreateInventoryActivityQuantitySpecificationResponse{}
	inventoryActivityQuantitySpecificationResponse.InventoryActivityQuantitySpecification = inventoryActivityQuantitySpecification
	return &inventoryActivityQuantitySpecificationResponse, nil
}

// ProcessInventoryActivityQuantitySpecificationRequest - ProcessInventoryActivityQuantitySpecificationRequest
func (invs *InventoryService) ProcessInventoryActivityQuantitySpecificationRequest(ctx context.Context, in *inventoryproto.CreateInventoryActivityQuantitySpecificationRequest) (*inventoryproto.InventoryActivityQuantitySpecification, error) {
	inventoryActivityQuantitySpecification := inventoryproto.InventoryActivityQuantitySpecification{}
	inventoryActivityQuantitySpecification.InventoryActivityTypeCode = in.InventoryActivityTypeCode
	inventoryActivityQuantitySpecification.InventoryMovementTypeCode = in.InventoryMovementTypeCode
	inventoryActivityQuantitySpecification.QuantityOfUnits = in.QuantityOfUnits
	inventoryActivityQuantitySpecification.QOUMeasurementUnitCode = in.QOUMeasurementUnitCode
	inventoryActivityQuantitySpecification.QOUCodeListVersion = in.QOUCodeListVersion
	inventoryActivityQuantitySpecification.InventoryStatusLineItemId = in.InventoryStatusLineItemId
	inventoryActivityQuantitySpecification.InventoryActivityLineItemId = in.InventoryActivityLineItemId
	inventoryActivityQuantitySpecification.InventoryReportId = in.InventoryReportId
	return &inventoryActivityQuantitySpecification, nil
}

// insertInventoryActivityQuantitySpecification - Insert InventoryActivityQuantitySpecification details into database
func (invs *InventoryService) insertInventoryActivityQuantitySpecification(ctx context.Context, insertInventoryActivityQuantitySpecificationSQL string, inventoryActivityQuantitySpecification *inventoryproto.InventoryActivityQuantitySpecification, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInventoryActivityQuantitySpecificationSQL, inventoryActivityQuantitySpecification)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryActivityQuantitySpecification.Id = uint32(uID)
		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
