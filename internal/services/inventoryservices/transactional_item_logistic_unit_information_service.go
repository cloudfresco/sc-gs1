package inventoryservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	inventorystruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertTransactionalItemLogisticUnitInformationSQL = `insert into transactional_item_logistic_unit_informations
	    (uuid4,
	    maximum_stacking_factor,
number_of_layers,
number_of_units_per_layer,
number_of_units_per_pallet,
package_type_code,
packaging_terms,
returnable_package_transport_cost_payment,
transactional_item_data_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(:uuid4,
  :maximum_stacking_factor,
:number_of_layers,
:number_of_units_per_layer,
:number_of_units_per_pallet,
:package_type_code,
:packaging_terms,
:returnable_package_transport_cost_payment,
:transactional_item_data_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectTransactionalItemLogisticUnitInformationsSQL = `select
  id,
  uuid4,
  maximum_stacking_factor,
  number_of_layers,
  number_of_units_per_layer,
  number_of_units_per_pallet,
  package_type_code,
  packaging_terms,
  returnable_package_transport_cost_payment,
  transactional_item_data_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from transactional_item_logistic_unit_informations`*/

// CreateTransactionalItemLogisticUnitInformation - Create TransactionalItemLogisticUnitInformation
func (invs *InventoryService) CreateTransactionalItemLogisticUnitInformation(ctx context.Context, in *inventoryproto.CreateTransactionalItemLogisticUnitInformationRequest) (*inventoryproto.CreateTransactionalItemLogisticUnitInformationResponse, error) {
	transactionalItemLogisticUnitInformation, err := invs.ProcessTransactionalItemLogisticUnitInformationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransactionalItemLogisticUnitInformation(ctx, insertTransactionalItemLogisticUnitInformationSQL, transactionalItemLogisticUnitInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalItemLogisticUnitInformationResponse := inventoryproto.CreateTransactionalItemLogisticUnitInformationResponse{}
	transactionalItemLogisticUnitInformationResponse.TransactionalItemLogisticUnitInformation = transactionalItemLogisticUnitInformation
	return &transactionalItemLogisticUnitInformationResponse, nil
}

// ProcessTransactionalItemLogisticUnitInformationRequest - ProcessTransactionalItemLogisticUnitInformationRequest
func (invs *InventoryService) ProcessTransactionalItemLogisticUnitInformationRequest(ctx context.Context, in *inventoryproto.CreateTransactionalItemLogisticUnitInformationRequest) (*inventoryproto.TransactionalItemLogisticUnitInformation, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	transactionalItemLogisticUnitInformationD := inventoryproto.TransactionalItemLogisticUnitInformationD{}
	transactionalItemLogisticUnitInformationD.MaximumStackingFactor = in.MaximumStackingFactor
	transactionalItemLogisticUnitInformationD.NumberOfLayers = in.NumberOfLayers
	transactionalItemLogisticUnitInformationD.NumberOfUnitsPerLayer = in.NumberOfUnitsPerLayer
	transactionalItemLogisticUnitInformationD.NumberOfUnitsPerPallet = in.NumberOfUnitsPerPallet
	transactionalItemLogisticUnitInformationD.PackageTypeCode = in.PackageTypeCode
	transactionalItemLogisticUnitInformationD.PackagingTerms = in.PackagingTerms
	transactionalItemLogisticUnitInformationD.ReturnablePackageTransportCostPayment = in.ReturnablePackageTransportCostPayment
	transactionalItemLogisticUnitInformationD.TransactionalItemDataId = in.TransactionalItemDataId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transactionalItemLogisticUnitInformation := inventoryproto.TransactionalItemLogisticUnitInformation{TransactionalItemLogisticUnitInformationD: &transactionalItemLogisticUnitInformationD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &transactionalItemLogisticUnitInformation, nil
}

// insertTransactionalItemLogisticUnitInformation - Insert TransactionalItemLogisticUnitInformation details into database
func (invs *InventoryService) insertTransactionalItemLogisticUnitInformation(ctx context.Context, insertTransactionalItemLogisticUnitInformationSQL string, transactionalItemLogisticUnitInformation *inventoryproto.TransactionalItemLogisticUnitInformation, userEmail string, requestID string) error {
	transactionalItemLogisticUnitInformationTmp, err := invs.crTransactionalItemLogisticUnitInformationStruct(ctx, transactionalItemLogisticUnitInformation, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalItemLogisticUnitInformationSQL, transactionalItemLogisticUnitInformationTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemLogisticUnitInformation.TransactionalItemLogisticUnitInformationD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transactionalItemLogisticUnitInformation.TransactionalItemLogisticUnitInformationD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemLogisticUnitInformation.TransactionalItemLogisticUnitInformationD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTransactionalItemLogisticUnitInformationStruct - process TransactionalItemLogisticUnitInformation details
func (invs *InventoryService) crTransactionalItemLogisticUnitInformationStruct(ctx context.Context, transactionalItemLogisticUnitInformation *inventoryproto.TransactionalItemLogisticUnitInformation, userEmail string, requestID string) (*inventorystruct.TransactionalItemLogisticUnitInformation, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transactionalItemLogisticUnitInformation.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transactionalItemLogisticUnitInformation.CrUpdTime.UpdatedAt)

	transactionalItemLogisticUnitInformationTmp := inventorystruct.TransactionalItemLogisticUnitInformation{TransactionalItemLogisticUnitInformationD: transactionalItemLogisticUnitInformation.TransactionalItemLogisticUnitInformationD, CrUpdUser: transactionalItemLogisticUnitInformation.CrUpdUser, CrUpdTime: crUpdTime}

	return &transactionalItemLogisticUnitInformationTmp, nil
}
