package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatMincingDetailSQL = `insert into meat_mincing_details
	  (
uuid4,
fat_content_percent,
mincing_type_code
)
values(
:uuid4,
:fat_content_percent,
:mincing_type_code);`

/*const selectMeatMincingDetailsSQL = `select
  id,
  uuid4,
  fat_content_percent,
  mincing_type_code from meat_mincing_details`*/

func (ms *MeatService) CreateMeatMincingDetail(ctx context.Context, in *meatproto.CreateMeatMincingDetailRequest) (*meatproto.CreateMeatMincingDetailResponse, error) {
	meatMincingDetail, err := ms.ProcessMeatMincingDetailRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatMincingDetail(ctx, insertMeatMincingDetailSQL, meatMincingDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatMincingDetailResponse := meatproto.CreateMeatMincingDetailResponse{}
	meatMincingDetailResponse.MeatMincingDetail = meatMincingDetail

	return &meatMincingDetailResponse, nil
}

// ProcessMeatMincingDetailRequest - ProcessMeatMincingDetailRequest
func (ms *MeatService) ProcessMeatMincingDetailRequest(ctx context.Context, in *meatproto.CreateMeatMincingDetailRequest) (*meatproto.MeatMincingDetail, error) {
	var err error

	meatMincingDetail := meatproto.MeatMincingDetail{}
	meatMincingDetail.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatMincingDetail.FatContentPercent = in.FatContentPercent
	meatMincingDetail.MincingTypeCode = in.MincingTypeCode

	return &meatMincingDetail, nil
}

// insertMeatMincingDetail - Insert MeatMincingDetail into database
func (ms *MeatService) insertMeatMincingDetail(ctx context.Context, insertMeatMincingDetailSQL string, meatMincingDetail *meatproto.MeatMincingDetail, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatMincingDetailSQL, meatMincingDetail)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatMincingDetail.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatMincingDetail.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatMincingDetail.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
