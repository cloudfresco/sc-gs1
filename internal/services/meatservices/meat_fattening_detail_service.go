package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatFatteningDetailSQL = `insert into meat_fattening_details
	  (
uuid4,
feeding_system_code,
housing_system_code
)
values(
:uuid4,
:feeding_system_code,
:housing_system_code);`

/*const selectMeatFatteningDetailsSQL = `select
  id,
  uuid4,
  feeding_system_code,
  housing_system_code from meat_fattening_details`*/

func (ms *MeatService) CreateMeatFatteningDetail(ctx context.Context, in *meatproto.CreateMeatFatteningDetailRequest) (*meatproto.CreateMeatFatteningDetailResponse, error) {
	meatFatteningDetail, err := ms.ProcessMeatFatteningDetailRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatFatteningDetail(ctx, insertMeatFatteningDetailSQL, meatFatteningDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatFatteningDetailResponse := meatproto.CreateMeatFatteningDetailResponse{}
	meatFatteningDetailResponse.MeatFatteningDetail = meatFatteningDetail

	return &meatFatteningDetailResponse, nil
}

// ProcessMeatFatteningDetailRequest - ProcessMeatFatteningDetailRequest
func (ms *MeatService) ProcessMeatFatteningDetailRequest(ctx context.Context, in *meatproto.CreateMeatFatteningDetailRequest) (*meatproto.MeatFatteningDetail, error) {
	var err error

	meatFatteningDetail := meatproto.MeatFatteningDetail{}
	meatFatteningDetail.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatFatteningDetail.FeedingSystemCode = in.FeedingSystemCode
	meatFatteningDetail.HousingSystemCode = in.HousingSystemCode

	return &meatFatteningDetail, nil
}

// insertMeatFatteningDetail - Insert MeatFatteningDetail into database
func (ms *MeatService) insertMeatFatteningDetail(ctx context.Context, insertMeatFatteningDetailSQL string, meatFatteningDetail *meatproto.MeatFatteningDetail, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatFatteningDetailSQL, meatFatteningDetail)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatFatteningDetail.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatFatteningDetail.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatFatteningDetail.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
