package meatservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	meatstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/meat/v1"
)

const insertMeatBreedingDetailSQL = `insert into meat_breeding_details
	  (
uuid4,
animal_type_code,
breed_code,
breed_of_father_code,
breed_of_mother_code,
cross_breed_indicator,
feeding_system_code,
gender_code,
housing_system_code,
breeding_dna_test,
date_of_birth,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
)
values(
:uuid4,
:animal_type_code,
:breed_code,
:breed_of_father_code,
:breed_of_mother_code,
:cross_breed_indicator,
:feeding_system_code,
:gender_code,
:housing_system_code,
:breeding_dna_test,
:date_of_birth,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectBreedingDetailsSQL = `select
 id,
uuid4,
animal_type_code,
breed_code,
breed_of_father_code,
breed_of_mother_code,
cross_breed_indicator,
feeding_system_code,
gender_code,
housing_system_code,
breeding_dna_test,
date_of_birth,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from meat_breeding_details`*/

func (ms *MeatService) CreateMeatBreedingDetail(ctx context.Context, in *meatproto.CreateMeatBreedingDetailRequest) (*meatproto.CreateMeatBreedingDetailResponse, error) {
	meatBreedingDetail, err := ms.ProcessMeatBreedingDetailRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatBreedingDetail(ctx, insertMeatBreedingDetailSQL, meatBreedingDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatBreedingDetailResponse := meatproto.CreateMeatBreedingDetailResponse{}
	meatBreedingDetailResponse.MeatBreedingDetail = meatBreedingDetail

	return &meatBreedingDetailResponse, nil
}

// ProcessMeatBreedingDetailRequest - ProcessMeatBreedingDetailRequest
func (ms *MeatService) ProcessMeatBreedingDetailRequest(ctx context.Context, in *meatproto.CreateMeatBreedingDetailRequest) (*meatproto.MeatBreedingDetail, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ms.UserServiceClient)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	dateOfBirth, err := time.Parse(common.Layout, in.DateOfBirth)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatBreedingDetailD := meatproto.MeatBreedingDetailD{}
	meatBreedingDetailD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatBreedingDetailD.AnimalTypeCode = in.AnimalTypeCode
	meatBreedingDetailD.BreedCode = in.BreedCode
	meatBreedingDetailD.BreedOfFatherCode = in.BreedOfFatherCode
	meatBreedingDetailD.BreedOfMotherCode = in.BreedOfMotherCode
	meatBreedingDetailD.CrossBreedIndicator = in.CrossBreedIndicator
	meatBreedingDetailD.FeedingSystemCode = in.FeedingSystemCode
	meatBreedingDetailD.GenderCode = in.GenderCode
	meatBreedingDetailD.HousingSystemCode = in.HousingSystemCode
	meatBreedingDetailD.BreedingDNATest = in.BreedingDNATest

	meatBreedingDetailT := meatproto.MeatBreedingDetailT{}
	meatBreedingDetailT.DateOfBirth = common.TimeToTimestamp(dateOfBirth.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	meatBreedingDetail := meatproto.MeatBreedingDetail{MeatBreedingDetailD: &meatBreedingDetailD, MeatBreedingDetailT: &meatBreedingDetailT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &meatBreedingDetail, nil
}

// insertMeatBreedingDetail - Insert MeatBreedingDetail into database
func (ms *MeatService) insertMeatBreedingDetail(ctx context.Context, insertMeatBreedingDetailSQL string, meatBreedingDetail *meatproto.MeatBreedingDetail, userEmail string, requestID string) error {
	meatBreedingDetailTmp, err := ms.crMeatBreedingDetailStruct(ctx, meatBreedingDetail, userEmail, requestID)
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatBreedingDetailSQL, meatBreedingDetailTmp)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatBreedingDetail.MeatBreedingDetailD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatBreedingDetail.MeatBreedingDetailD.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatBreedingDetail.MeatBreedingDetailD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crMeatBreedingDetailStruct - process MeatBreedingDetail details
func (ms *MeatService) crMeatBreedingDetailStruct(ctx context.Context, meatBreedingDetail *meatproto.MeatBreedingDetail, userEmail string, requestID string) (*meatstruct.MeatBreedingDetail, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(meatBreedingDetail.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(meatBreedingDetail.CrUpdTime.UpdatedAt)

	meatBreedingDetailT := new(meatstruct.MeatBreedingDetailT)
	meatBreedingDetailT.DateOfBirth = common.TimestampToTime(meatBreedingDetail.MeatBreedingDetailT.DateOfBirth)

	meatBreedingDetailTmp := meatstruct.MeatBreedingDetail{MeatBreedingDetailD: meatBreedingDetail.MeatBreedingDetailD, MeatBreedingDetailT: meatBreedingDetailT, CrUpdUser: meatBreedingDetail.CrUpdUser, CrUpdTime: crUpdTime}

	return &meatBreedingDetailTmp, nil
}
