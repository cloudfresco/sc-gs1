package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatAciditySQL = `insert into meat_acidities
	  (
uuid4,
acidity_measurement_time,
acidity_of_meat,
meat_slaughtering_detail_id
)
values(
:uuid4,
:acidity_measurement_time,
:acidity_of_meat,
:meat_slaughtering_detail_id);`

/*const selectMeatAciditiesSQL = `select
  id,
  uuid4,
  acidity_measurement_time,
  acidity_of_meat,
  meat_slaughtering_detail_id from meat_acidities`*/

func (ms *MeatService) CreateMeatAcidity(ctx context.Context, in *meatproto.CreateMeatAcidityRequest) (*meatproto.CreateMeatAcidityResponse, error) {
	meatAcidity, err := ms.ProcessMeatAcidityRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatAcidity(ctx, insertMeatAciditySQL, meatAcidity, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatAcidityResponse := meatproto.CreateMeatAcidityResponse{}
	meatAcidityResponse.MeatAcidity = meatAcidity

	return &meatAcidityResponse, nil
}

// ProcessMeatAcidityRequest - ProcessMeatAcidityRequest
func (ms *MeatService) ProcessMeatAcidityRequest(ctx context.Context, in *meatproto.CreateMeatAcidityRequest) (*meatproto.MeatAcidity, error) {
	var err error

	meatAcidity := meatproto.MeatAcidity{}
	meatAcidity.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatAcidity.AcidityMeasurementTime = in.AcidityMeasurementTime
	meatAcidity.AcidityOfMeat = in.AcidityOfMeat
	meatAcidity.MeatSlaughteringDetailId = in.MeatSlaughteringDetailId

	return &meatAcidity, nil
}

// insertMeatAcidity - Insert MeatAcidity into database
func (ms *MeatService) insertMeatAcidity(ctx context.Context, insertMeatAciditySQL string, meatAcidity *meatproto.MeatAcidity, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatAciditySQL, meatAcidity)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatAcidity.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatAcidity.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatAcidity.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
