package inventoryservices

import (
	"context"

	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertTransactionalItemCertificationSQL = `insert into transactional_item_certifications
	    (item_certification_agency,
item_certification_standard,
item_certification_value,
organic_certification_id)
  values(
 :item_certification_agency,
 :item_certification_standard,
 :item_certification_value,
 :organic_certification_id
);`

/*const selectTransactionalItemCertificationsSQL = `select
  id,
  item_certification_agency,
  item_certification_standard,
  item_certification_value,
  organic_certification_id from transactional_item_certifications`*/

// CreateTransactionalItemCertification - Create TransactionalItemCertification
func (invs *InventoryService) CreateTransactionalItemCertification(ctx context.Context, in *inventoryproto.CreateTransactionalItemCertificationRequest) (*inventoryproto.CreateTransactionalItemCertificationResponse, error) {
	transactionalItemCertification, err := invs.ProcessTransactionalItemCertificationRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransactionalItemCertification(ctx, insertTransactionalItemCertificationSQL, transactionalItemCertification, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalItemCertificationResponse := inventoryproto.CreateTransactionalItemCertificationResponse{}
	transactionalItemCertificationResponse.TransactionalItemCertification = transactionalItemCertification
	return &transactionalItemCertificationResponse, nil
}

// ProcessTransactionalItemCertificationRequest - ProcessTransactionalItemCertificationRequest
func (invs *InventoryService) ProcessTransactionalItemCertificationRequest(ctx context.Context, in *inventoryproto.CreateTransactionalItemCertificationRequest) (*inventoryproto.TransactionalItemCertification, error) {
	transactionalItemCertification := inventoryproto.TransactionalItemCertification{}
	transactionalItemCertification.ItemCertificationAgency = in.ItemCertificationAgency
	transactionalItemCertification.ItemCertificationStandard = in.ItemCertificationStandard
	transactionalItemCertification.ItemCertificationValue = in.ItemCertificationValue
	transactionalItemCertification.OrganicCertificationId = in.OrganicCertificationId
	return &transactionalItemCertification, nil
}

// insertTransactionalItemCertification - Insert TransactionalItemCertification details into database
func (invs *InventoryService) insertTransactionalItemCertification(ctx context.Context, insertTransactionalItemCertificationSQL string, transactionalItemCertification *inventoryproto.TransactionalItemCertification, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalItemCertificationSQL, transactionalItemCertification)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemCertification.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
