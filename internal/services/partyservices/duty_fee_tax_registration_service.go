package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDutyFeeTaxRegistrationSQL = `insert into duty_fee_tax_registrations
	    (duty_fee_tax_agency_name,
duty_fee_tax_description,
duty_fee_tax_registration_type,
duty_fee_tax_type_code,
transactional_party_id)
  values(
  :duty_fee_tax_agency_name,
  :duty_fee_tax_description,
  :duty_fee_tax_registration_type,
  :duty_fee_tax_type_code,
  :transactional_party_id);`

/*const selectDutyFeeTaxRegistrationsSQL = `select
  id,
  duty_fee_tax_agency_name,
  duty_fee_tax_description,
  duty_fee_tax_registration_type,
  duty_fee_tax_type_code,
  transactional_party_id from duty_fee_tax_registrations`*/

func (ps *PartyService) CreateDutyFeeTaxRegistration(ctx context.Context, in *partyproto.CreateDutyFeeTaxRegistrationRequest) (*partyproto.CreateDutyFeeTaxRegistrationResponse, error) {
	dutyFeeTaxRegistration, err := ps.ProcessDutyFeeTaxRegistrationRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertDutyFeeTaxRegistration(ctx, insertDutyFeeTaxRegistrationSQL, dutyFeeTaxRegistration, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	dutyFeeTaxRegistrationResponse := partyproto.CreateDutyFeeTaxRegistrationResponse{}
	dutyFeeTaxRegistrationResponse.DutyFeeTaxRegistration = dutyFeeTaxRegistration
	return &dutyFeeTaxRegistrationResponse, nil
}

// ProcessDutyFeeTaxRegistrationRequest - ProcessDutyFeeTaxRegistrationRequest
func (ps *PartyService) ProcessDutyFeeTaxRegistrationRequest(ctx context.Context, in *partyproto.CreateDutyFeeTaxRegistrationRequest) (*partyproto.DutyFeeTaxRegistration, error) {
	dutyFeeTaxRegistration := partyproto.DutyFeeTaxRegistration{}
	dutyFeeTaxRegistration.DutyFeeTaxAgencyName = in.DutyFeeTaxAgencyName
	dutyFeeTaxRegistration.DutyFeeTaxDescription = in.DutyFeeTaxDescription
	dutyFeeTaxRegistration.DutyFeeTaxRegistrationType = in.DutyFeeTaxRegistrationType
	dutyFeeTaxRegistration.DutyFeeTaxTypeCode = in.DutyFeeTaxTypeCode
	dutyFeeTaxRegistration.TransactionalPartyId = in.TransactionalPartyId
	return &dutyFeeTaxRegistration, nil
}

// insertDutyFeeTaxRegistration - Insert DutyFeeTaxRegistration into database
func (ps *PartyService) insertDutyFeeTaxRegistration(ctx context.Context, insertDutyFeeTaxRegistrationSQL string, dutyFeeTaxRegistration *partyproto.DutyFeeTaxRegistration, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertDutyFeeTaxRegistrationSQL, dutyFeeTaxRegistration)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		dutyFeeTaxRegistration.Id = uint32(uID)
		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
