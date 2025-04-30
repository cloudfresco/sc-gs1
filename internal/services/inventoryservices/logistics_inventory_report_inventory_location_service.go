package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertLogisticsInventoryReportInventoryLocationSQL = `insert into logistics_inventory_report_inventory_locations
	    (inventory_location_id,
logistics_inventory_report_id)
  values(:inventory_location_id,
:logistics_inventory_report_id);`

/*const selectLogisticsInventoryReportInventoryLocationsSQL = `select
  id,
  inventory_location_id,
  logistics_inventory_report_id from logistics_inventory_report_inventory_locations`*/

// CreateLogisticsInventoryReportInventoryLocation - Create LogisticsInventoryReportInventoryLocation
func (invs *InventoryService) CreateLogisticsInventoryReportInventoryLocation(ctx context.Context, in *inventoryproto.CreateLogisticsInventoryReportInventoryLocationRequest) (*inventoryproto.CreateLogisticsInventoryReportInventoryLocationResponse, error) {
	logisticsInventoryReportInventoryLocation, err := invs.ProcessLogisticsInventoryReportInventoryLocationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertLogisticsInventoryReportInventoryLocation(ctx, insertLogisticsInventoryReportInventoryLocationSQL, logisticsInventoryReportInventoryLocation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticsInventoryReportInventoryLocationResponse := inventoryproto.CreateLogisticsInventoryReportInventoryLocationResponse{}
	logisticsInventoryReportInventoryLocationResponse.LogisticsInventoryReportInventoryLocation = logisticsInventoryReportInventoryLocation
	return &logisticsInventoryReportInventoryLocationResponse, nil
}

// ProcessLogisticsInventoryReportInventoryLocationRequest - ProcessLogisticsInventoryReportInventoryLocationRequest
func (invs *InventoryService) ProcessLogisticsInventoryReportInventoryLocationRequest(ctx context.Context, in *inventoryproto.CreateLogisticsInventoryReportInventoryLocationRequest) (*inventoryproto.LogisticsInventoryReportInventoryLocation, error) {
	logisticsInventoryReportInventoryLocation := inventoryproto.LogisticsInventoryReportInventoryLocation{}
	logisticsInventoryReportInventoryLocation.InventoryLocationId = in.InventoryLocationId
	logisticsInventoryReportInventoryLocation.LogisticsInventoryReportId = in.LogisticsInventoryReportId

	return &logisticsInventoryReportInventoryLocation, nil
}

// insertLogisticsInventoryReportInventoryLocation - Insert LogisticsInventoryReportInventoryLocation details into database
func (invs *InventoryService) insertLogisticsInventoryReportInventoryLocation(ctx context.Context, insertLogisticsInventoryReportInventoryLocationSQL string, logisticsInventoryReportInventoryLocation *inventoryproto.LogisticsInventoryReportInventoryLocation, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLogisticsInventoryReportInventoryLocationSQL, logisticsInventoryReportInventoryLocation)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticsInventoryReportInventoryLocation.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
