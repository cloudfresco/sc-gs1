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

const insertMeatSlaughteringDetailSQL = `insert into meat_slaughtering_details
	  (
uuid4,
age_of_animal,
fat_content_percent,
fat_cover_code,
meat_category_code,
meat_colour_code,
meat_conformation_code,
meat_profile_code,
slaughtering_system_code,
slaughtering_weight,
sw_code_list_version,
sw_measurement_unit_code,
bse_test_id,
slaughtering_dna_test_id,
date_of_slaughtering,
optimum_maturation_date,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
)
values(
:uuid4,
:age_of_animal,
:fat_content_percent,
:fat_cover_code,
:meat_category_code,
:meat_colour_code,
:meat_conformation_code,
:meat_profile_code,
:slaughtering_system_code,
:slaughtering_weight,
:sw_code_list_version,
:sw_measurement_unit_code,
:bse_test_id,
:slaughtering_dna_test_id,
:date_of_slaughtering,
:optimum_maturation_date,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectSlaughteringDetailsSQL = `select
 id,
uuid4,
age_of_animal,
fat_content_percent,
fat_cover_code,
meat_category_code,
meat_colour_code,
meat_conformation_code,
meat_profile_code,
slaughtering_system_code,
slaughtering_weight,
sw_code_list_version,
sw_measurement_unit_code,
bse_test_id,
slaughtering_dna_test_id,
date_of_slaughtering,
optimum_maturation_date,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from meat_slaughtering_details`*/

func (ms *MeatService) CreateMeatSlaughteringDetail(ctx context.Context, in *meatproto.CreateMeatSlaughteringDetailRequest) (*meatproto.CreateMeatSlaughteringDetailResponse, error) {
	meatSlaughteringDetail, err := ms.ProcessMeatSlaughteringDetailRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatSlaughteringDetail(ctx, insertMeatSlaughteringDetailSQL, meatSlaughteringDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatSlaughteringDetailResponse := meatproto.CreateMeatSlaughteringDetailResponse{}
	meatSlaughteringDetailResponse.MeatSlaughteringDetail = meatSlaughteringDetail

	return &meatSlaughteringDetailResponse, nil
}

// ProcessMeatSlaughteringDetailRequest - ProcessMeatSlaughteringDetailRequest
func (ms *MeatService) ProcessMeatSlaughteringDetailRequest(ctx context.Context, in *meatproto.CreateMeatSlaughteringDetailRequest) (*meatproto.MeatSlaughteringDetail, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ms.UserServiceClient)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	dateOfSlaughtering, err := time.Parse(common.Layout, in.DateOfSlaughtering)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	optimumMaturationDate, err := time.Parse(common.Layout, in.OptimumMaturationDate)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatSlaughteringDetailD := meatproto.MeatSlaughteringDetailD{}
	meatSlaughteringDetailD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatSlaughteringDetailD.AgeOfAnimal = in.AgeOfAnimal
	meatSlaughteringDetailD.FatContentPercent = in.FatContentPercent
	meatSlaughteringDetailD.FatCoverCode = in.FatCoverCode
	meatSlaughteringDetailD.MeatCategoryCode = in.MeatCategoryCode
	meatSlaughteringDetailD.MeatColourCode = in.MeatColourCode
	meatSlaughteringDetailD.MeatConformationCode = in.MeatConformationCode
	meatSlaughteringDetailD.MeatProfileCode = in.MeatProfileCode
	meatSlaughteringDetailD.SlaughteringSystemCode = in.SlaughteringSystemCode
	meatSlaughteringDetailD.SlaughteringWeight = in.SlaughteringWeight
	meatSlaughteringDetailD.SWCodeListVersion = in.SWCodeListVersion
	meatSlaughteringDetailD.SWMeasurementUnitCode = in.SWMeasurementUnitCode
	meatSlaughteringDetailD.BseTestId = in.BseTestId
	meatSlaughteringDetailD.SlaughteringDNATestId = in.SlaughteringDNATestId

	meatSlaughteringDetailT := meatproto.MeatSlaughteringDetailT{}
	meatSlaughteringDetailT.DateOfSlaughtering = common.TimeToTimestamp(dateOfSlaughtering.UTC().Truncate(time.Second))
	meatSlaughteringDetailT.OptimumMaturationDate = common.TimeToTimestamp(optimumMaturationDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	meatSlaughteringDetail := meatproto.MeatSlaughteringDetail{MeatSlaughteringDetailD: &meatSlaughteringDetailD, MeatSlaughteringDetailT: &meatSlaughteringDetailT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &meatSlaughteringDetail, nil
}

// insertMeatSlaughteringDetail - Insert MeatSlaughteringDetail into database
func (ms *MeatService) insertMeatSlaughteringDetail(ctx context.Context, insertMeatSlaughteringDetailSQL string, meatSlaughteringDetail *meatproto.MeatSlaughteringDetail, userEmail string, requestID string) error {
	meatSlaughteringDetailTmp, err := ms.crMeatSlaughteringDetailStruct(ctx, meatSlaughteringDetail, userEmail, requestID)
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatSlaughteringDetailSQL, meatSlaughteringDetailTmp)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatSlaughteringDetail.MeatSlaughteringDetailD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatSlaughteringDetail.MeatSlaughteringDetailD.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatSlaughteringDetail.MeatSlaughteringDetailD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crMeatSlaughteringDetailStruct - process MeatSlaughteringDetail details
func (ms *MeatService) crMeatSlaughteringDetailStruct(ctx context.Context, meatSlaughteringDetail *meatproto.MeatSlaughteringDetail, userEmail string, requestID string) (*meatstruct.MeatSlaughteringDetail, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(meatSlaughteringDetail.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(meatSlaughteringDetail.CrUpdTime.UpdatedAt)

	meatSlaughteringDetailT := new(meatstruct.MeatSlaughteringDetailT)
	meatSlaughteringDetailT.DateOfSlaughtering = common.TimestampToTime(meatSlaughteringDetail.MeatSlaughteringDetailT.DateOfSlaughtering)
	meatSlaughteringDetailT.OptimumMaturationDate = common.TimestampToTime(meatSlaughteringDetail.MeatSlaughteringDetailT.OptimumMaturationDate)

	meatSlaughteringDetailTmp := meatstruct.MeatSlaughteringDetail{MeatSlaughteringDetailD: meatSlaughteringDetail.MeatSlaughteringDetailD, MeatSlaughteringDetailT: meatSlaughteringDetailT, CrUpdUser: meatSlaughteringDetail.CrUpdUser, CrUpdTime: crUpdTime}

	return &meatSlaughteringDetailTmp, nil
}
