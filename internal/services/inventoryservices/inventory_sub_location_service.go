package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInventorySubLocationSQL = `insert into inventory_sub_locations
	    (additional_party_identification,
additional_party_identification_type_code,
code_list_version,
gln,
gln_extension,
inventory_sub_location_function_code,
inventory_sub_location_type_code)
  values(:additional_party_identification,
:additional_party_identification_type_code,
:code_list_version,
:gln,
:gln_extension,
:inventory_sub_location_function_code,
:inventory_sub_location_type_code);`

/*const selectInventorySubLocationsSQL = `select
  id,
  additional_party_identification,
  additional_party_identification_type_code,
  code_list_version,
  gln,
  gln_extension,
  inventory_sub_location_function_code,
  inventory_sub_location_type_code from inventory_sub_locations`*/

// CreateInventorySubLocation - Create InventorySubLocation
func (invs *InventoryService) CreateInventorySubLocation(ctx context.Context, in *inventoryproto.CreateInventorySubLocationRequest) (*inventoryproto.CreateInventorySubLocationResponse, error) {
	inventorySubLocation, err := invs.ProcessInventorySubLocationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInventorySubLocation(ctx, insertInventorySubLocationSQL, inventorySubLocation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventorySubLocationResponse := inventoryproto.CreateInventorySubLocationResponse{}
	inventorySubLocationResponse.InventorySubLocation = inventorySubLocation
	return &inventorySubLocationResponse, nil
}

// ProcessInventorySubLocationRequest - ProcessInventorySubLocationRequest
func (invs *InventoryService) ProcessInventorySubLocationRequest(ctx context.Context, in *inventoryproto.CreateInventorySubLocationRequest) (*inventoryproto.InventorySubLocation, error) {
	inventorySubLocation := inventoryproto.InventorySubLocation{}
	inventorySubLocation.AdditionalPartyIdentification = in.AdditionalPartyIdentification
	inventorySubLocation.AdditionalPartyIdentificationTypeCode = in.AdditionalPartyIdentificationTypeCode
	inventorySubLocation.CodeListVersion = in.CodeListVersion
	inventorySubLocation.Gln = in.Gln
	inventorySubLocation.GlnExtension = in.GlnExtension
	inventorySubLocation.InventorySubLocationFunctionCode = in.InventorySubLocationFunctionCode
	inventorySubLocation.InventorySubLocationTypeCode = in.InventorySubLocationTypeCode

	return &inventorySubLocation, nil
}

// insertInventorySubLocation - Insert InventorySubLocation details into database
func (invs *InventoryService) insertInventorySubLocation(ctx context.Context, insertInventorySubLocationSQL string, inventorySubLocation *inventoryproto.InventorySubLocation, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInventorySubLocationSQL, inventorySubLocation)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventorySubLocation.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
