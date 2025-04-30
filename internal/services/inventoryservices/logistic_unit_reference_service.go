package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertLogisticUnitReferenceSQL = `insert into logistic_unit_references
	    (trade_item_quantity,
q_measurement_unit_code,
q_code_list_version)
  values(:trade_item_quantity,
:q_measurement_unit_code,
:q_code_list_version);`

/*const selectLogisticUnitReferencesSQL = `select
  id,
  trade_item_quantity,
  q_measurement_unit_code,
  q_code_list_version from logistic_unit_references`*/

// CreateLogisticUnitReference - Create LogisticUnitReference
func (invs *InventoryService) CreateLogisticUnitReference(ctx context.Context, in *inventoryproto.CreateLogisticUnitReferenceRequest) (*inventoryproto.CreateLogisticUnitReferenceResponse, error) {
	logisticUnitReference, err := invs.ProcessLogisticUnitReferenceRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertLogisticUnitReference(ctx, insertLogisticUnitReferenceSQL, logisticUnitReference, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticUnitReferenceResponse := inventoryproto.CreateLogisticUnitReferenceResponse{}
	logisticUnitReferenceResponse.LogisticUnitReference = logisticUnitReference
	return &logisticUnitReferenceResponse, nil
}

// ProcessLogisticUnitReferenceRequest - ProcessLogisticUnitReferenceRequest
func (invs *InventoryService) ProcessLogisticUnitReferenceRequest(ctx context.Context, in *inventoryproto.CreateLogisticUnitReferenceRequest) (*inventoryproto.LogisticUnitReference, error) {
	logisticUnitReference := inventoryproto.LogisticUnitReference{}
	logisticUnitReference.TradeItemQuantity = in.TradeItemQuantity
	logisticUnitReference.QMeasurementUnitCode = in.QMeasurementUnitCode
	logisticUnitReference.QCodeListVersion = in.QCodeListVersion

	return &logisticUnitReference, nil
}

// insertLogisticUnitReference - Insert LogisticUnitReference details into database
func (invs *InventoryService) insertLogisticUnitReference(ctx context.Context, insertLogisticUnitReferenceSQL string, logisticUnitReference *inventoryproto.LogisticUnitReference, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLogisticUnitReferenceSQL, logisticUnitReference)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticUnitReference.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
