package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	meatstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/meat/v1"
)

const insertMeatWorkItemIdentificationSQL = `insert into meat_breeding_details
	  (
uuid4,
batch_number,
livestock_mob_identifier,
meat_work_item_type_code,
animal_identification_id,
product_identification,
slaughter_number_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
)
values(
:uuid4,
:batch_number,
:livestock_mob_identifier,
:meat_work_item_type_code,
:animal_identification_id,
:product_identification,
:slaughter_number_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectWorkItemIdentificationsSQL = `select
 id,
uuid4,
batch_number,
livestock_mob_identifier,
meat_work_item_type_code,
animal_identification_id,
product_identification,
slaughter_number_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from meat_breeding_details`*/

func (ms *MeatService) CreateMeatWorkItemIdentification(ctx context.Context, in *meatproto.CreateMeatWorkItemIdentificationRequest) (*meatproto.CreateMeatWorkItemIdentificationResponse, error) {
	meatWorkItemIdentification, err := ms.ProcessMeatWorkItemIdentificationRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatWorkItemIdentification(ctx, insertMeatWorkItemIdentificationSQL, meatWorkItemIdentification, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatWorkItemIdentificationResponse := meatproto.CreateMeatWorkItemIdentificationResponse{}
	meatWorkItemIdentificationResponse.MeatWorkItemIdentification = meatWorkItemIdentification

	return &meatWorkItemIdentificationResponse, nil
}

// ProcessMeatWorkItemIdentificationRequest - ProcessMeatWorkItemIdentificationRequest
func (ms *MeatService) ProcessMeatWorkItemIdentificationRequest(ctx context.Context, in *meatproto.CreateMeatWorkItemIdentificationRequest) (*meatproto.MeatWorkItemIdentification, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ms.UserServiceClient)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	meatWorkItemIdentificationD := meatproto.MeatWorkItemIdentificationD{}
	meatWorkItemIdentificationD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatWorkItemIdentificationD.BatchNumber = in.BatchNumber
	meatWorkItemIdentificationD.LivestockMobIdentifier = in.LivestockMobIdentifier
	meatWorkItemIdentificationD.MeatWorkItemTypeCode = in.MeatWorkItemTypeCode
	meatWorkItemIdentificationD.AnimalIdentificationId = in.AnimalIdentificationId
	meatWorkItemIdentificationD.ProductIdentification = in.ProductIdentification
	meatWorkItemIdentificationD.SlaughterNumberId = in.SlaughterNumberId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	meatWorkItemIdentification := meatproto.MeatWorkItemIdentification{MeatWorkItemIdentificationD: &meatWorkItemIdentificationD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &meatWorkItemIdentification, nil
}

// insertMeatWorkItemIdentification - Insert MeatWorkItemIdentification into database
func (ms *MeatService) insertMeatWorkItemIdentification(ctx context.Context, insertMeatWorkItemIdentificationSQL string, meatWorkItemIdentification *meatproto.MeatWorkItemIdentification, userEmail string, requestID string) error {
	meatWorkItemIdentificationTmp, err := ms.crMeatWorkItemIdentificationStruct(ctx, meatWorkItemIdentification, userEmail, requestID)
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatWorkItemIdentificationSQL, meatWorkItemIdentificationTmp)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatWorkItemIdentification.MeatWorkItemIdentificationD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatWorkItemIdentification.MeatWorkItemIdentificationD.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatWorkItemIdentification.MeatWorkItemIdentificationD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crMeatWorkItemIdentificationStruct - process MeatWorkItemIdentification details
func (ms *MeatService) crMeatWorkItemIdentificationStruct(ctx context.Context, meatWorkItemIdentification *meatproto.MeatWorkItemIdentification, userEmail string, requestID string) (*meatstruct.MeatWorkItemIdentification, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(meatWorkItemIdentification.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(meatWorkItemIdentification.CrUpdTime.UpdatedAt)

	meatWorkItemIdentificationTmp := meatstruct.MeatWorkItemIdentification{MeatWorkItemIdentificationD: meatWorkItemIdentification.MeatWorkItemIdentificationD, CrUpdUser: meatWorkItemIdentification.CrUpdUser, CrUpdTime: crUpdTime}

	return &meatWorkItemIdentificationTmp, nil
}
