package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatCuttingDetailSQL = `insert into meat_cutting_details
	  (
uuid4,
meat_profile_code
)
values(
:uuid4,
:meat_profile_code);`

/*const selectMeatCuttingDetailsSQL = `select
  id,
  uuid4,
  meat_profile_code from meat_cutting_details`*/

func (ms *MeatService) CreateMeatCuttingDetail(ctx context.Context, in *meatproto.CreateMeatCuttingDetailRequest) (*meatproto.CreateMeatCuttingDetailResponse, error) {
	meatCuttingDetail, err := ms.ProcessMeatCuttingDetailRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatCuttingDetail(ctx, insertMeatCuttingDetailSQL, meatCuttingDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatCuttingDetailResponse := meatproto.CreateMeatCuttingDetailResponse{}
	meatCuttingDetailResponse.MeatCuttingDetail = meatCuttingDetail

	return &meatCuttingDetailResponse, nil
}

// ProcessMeatCuttingDetailRequest - ProcessMeatCuttingDetailRequest
func (ms *MeatService) ProcessMeatCuttingDetailRequest(ctx context.Context, in *meatproto.CreateMeatCuttingDetailRequest) (*meatproto.MeatCuttingDetail, error) {
	var err error

	meatCuttingDetail := meatproto.MeatCuttingDetail{}
	meatCuttingDetail.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatCuttingDetail.MeatProfileCode = in.MeatProfileCode

	return &meatCuttingDetail, nil
}

// insertMeatCuttingDetail - Insert MeatCuttingDetail into database
func (ms *MeatService) insertMeatCuttingDetail(ctx context.Context, insertMeatCuttingDetailSQL string, meatCuttingDetail *meatproto.MeatCuttingDetail, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatCuttingDetailSQL, meatCuttingDetail)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatCuttingDetail.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatCuttingDetail.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatCuttingDetail.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
