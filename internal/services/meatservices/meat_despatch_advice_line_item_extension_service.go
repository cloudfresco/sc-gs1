package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatDespatchAdviceLineItemExtensionSQL = `insert into meat_despatch_advice_line_item_extensions
	  (
uuid4,
animal_identification_id,
slaughter_number_id
)
values(
:uuid4,
:animal_identification_id,
:slaughter_number_id);`

/*const selectMeatDespatchAdviceLineItemExtensionsSQL = `select
  id,
  uuid4,
  animal_identification_id,
  slaughter_number_id from meat_despatch_advice_line_item_extensions`*/

func (ms *MeatService) CreateMeatDespatchAdviceLineItemExtension(ctx context.Context, in *meatproto.CreateMeatDespatchAdviceLineItemExtensionRequest) (*meatproto.CreateMeatDespatchAdviceLineItemExtensionResponse, error) {
	meatDespatchAdviceLineItemExtension, err := ms.ProcessMeatDespatchAdviceLineItemExtensionRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatDespatchAdviceLineItemExtension(ctx, insertMeatDespatchAdviceLineItemExtensionSQL, meatDespatchAdviceLineItemExtension, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatDespatchAdviceLineItemExtensionResponse := meatproto.CreateMeatDespatchAdviceLineItemExtensionResponse{}
	meatDespatchAdviceLineItemExtensionResponse.MeatDespatchAdviceLineItemExtension = meatDespatchAdviceLineItemExtension

	return &meatDespatchAdviceLineItemExtensionResponse, nil
}

// ProcessMeatDespatchAdviceLineItemExtensionRequest - ProcessMeatDespatchAdviceLineItemExtensionRequest
func (ms *MeatService) ProcessMeatDespatchAdviceLineItemExtensionRequest(ctx context.Context, in *meatproto.CreateMeatDespatchAdviceLineItemExtensionRequest) (*meatproto.MeatDespatchAdviceLineItemExtension, error) {
	var err error

	meatDespatchAdviceLineItemExtension := meatproto.MeatDespatchAdviceLineItemExtension{}
	meatDespatchAdviceLineItemExtension.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatDespatchAdviceLineItemExtension.AnimalIdentificationId = in.AnimalIdentificationId
	meatDespatchAdviceLineItemExtension.SlaughterNumberId = in.SlaughterNumberId

	return &meatDespatchAdviceLineItemExtension, nil
}

// insertMeatDespatchAdviceLineItemExtension - Insert MeatDespatchAdviceLineItemExtension into database
func (ms *MeatService) insertMeatDespatchAdviceLineItemExtension(ctx context.Context, insertMeatDespatchAdviceLineItemExtensionSQL string, meatDespatchAdviceLineItemExtension *meatproto.MeatDespatchAdviceLineItemExtension, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatDespatchAdviceLineItemExtensionSQL, meatDespatchAdviceLineItemExtension)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatDespatchAdviceLineItemExtension.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatDespatchAdviceLineItemExtension.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatDespatchAdviceLineItemExtension.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
