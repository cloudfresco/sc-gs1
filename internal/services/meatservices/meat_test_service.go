package meatservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertMeatTestSQL = `insert into meat_tests
	  (
uuid4,
test_method,
test_result
)
values(
:uuid4,
:test_method,
:test_result);`

/*const selectMeatTestsSQL = `select
  id,
  uuid4,
  test_method,
  test_result from meat_tests`*/

func (ms *MeatService) CreateMeatTest(ctx context.Context, in *meatproto.CreateMeatTestRequest) (*meatproto.CreateMeatTestResponse, error) {
	meatTest, err := ms.ProcessMeatTestRequest(ctx, in)
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ms.insertMeatTest(ctx, insertMeatTestSQL, meatTest, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	meatTestResponse := meatproto.CreateMeatTestResponse{}
	meatTestResponse.MeatTest = meatTest

	return &meatTestResponse, nil
}

// ProcessMeatTestRequest - ProcessMeatTestRequest
func (ms *MeatService) ProcessMeatTestRequest(ctx context.Context, in *meatproto.CreateMeatTestRequest) (*meatproto.MeatTest, error) {
	var err error

	meatTest := meatproto.MeatTest{}
	meatTest.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ms.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	meatTest.TestMethod = in.TestMethod
	meatTest.TestResult = in.TestResult

	return &meatTest, nil
}

// insertMeatTest - Insert MeatTest into database
func (ms *MeatService) insertMeatTest(ctx context.Context, insertMeatTestSQL string, meatTest *meatproto.MeatTest, userEmail string, requestID string) error {
	err := ms.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertMeatTestSQL, meatTest)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatTest.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(meatTest.Uuid4)
		if err != nil {
			ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		meatTest.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ms.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
