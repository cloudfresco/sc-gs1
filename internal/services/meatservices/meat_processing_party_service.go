package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatProcessingPartySQL = `insert into meat_processing_parties
	  (
uuid4,
approval_number,
meat_processing_party_identification_type_code,
meat_processing_party_type_code,
transactional_party_id
)
values(
:uuid4,
:approval_number,
:meat_processing_party_identification_type_code,
:meat_processing_party_type_code,
:transactional_party_id);`

/*const selectMeatProcessingPartiesSQL = `select
  id,
  uuid4,
  approval_number,
  meat_processing_party_identification_type_code,
  meat_processing_party_type_code,
  transactional_party_id from meat_processing_parties`*/

func (ms *MeatService) CreateMeatProcessingParty(ctx context.Context, in *meatproto.CreateMeatProcessingPartyRequest) (*meatproto.CreateMeatProcessingPartyResponse, error) {
	meatProcessingParty, err := ms.ProcessMeatProcessingPartyRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatProcessingParty(ctx, insertMeatProcessingPartySQL, meatProcessingParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatProcessingPartyResponse := meatproto.CreateMeatProcessingPartyResponse{}
	meatProcessingPartyResponse.MeatProcessingParty = meatProcessingParty

	return &meatProcessingPartyResponse, nil
}

// ProcessMeatProcessingPartyRequest - ProcessMeatProcessingPartyRequest
func (ms *MeatService) ProcessMeatProcessingPartyRequest(ctx context.Context, in *meatproto.CreateMeatProcessingPartyRequest) (*meatproto.MeatProcessingParty, error) {
	var err error

	meatProcessingParty := meatproto.MeatProcessingParty{}
	meatProcessingParty.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatProcessingParty.ApprovalNumber = in.ApprovalNumber
	meatProcessingParty.MeatProcessingPartyIdentificationTypeCode = in.MeatProcessingPartyIdentificationTypeCode
	meatProcessingParty.MeatProcessingPartyTypeCode = in.MeatProcessingPartyTypeCode
	meatProcessingParty.TransactionalPartyId = in.TransactionalPartyId

	return &meatProcessingParty, nil
}

// insertMeatProcessingParty - Insert MeatProcessingParty into database
func (ms *MeatService) insertMeatProcessingParty(ctx context.Context, insertMeatProcessingPartySQL string, meatProcessingParty *meatproto.MeatProcessingParty, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatProcessingPartySQL, meatProcessingParty)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatProcessingParty.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatProcessingParty.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatProcessingParty.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
