package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertLegalRegistrationSQL = `insert into legal_registrations
	    (legal_registration_additional_information,
       legal_registration_number,
       legal_registration_type)
  values(
  :legal_registration_additional_information,
  :legal_registration_number,
  :legal_registration_type);`

/*const selectLegalRegistrationsSQL = `select
  id,
  legal_registration_additional_information,
  legal_registration_number,
  legal_registration_type from legal_registrations`*/

func (ps *PartyService) CreateLegalRegistration(ctx context.Context, in *partyproto.CreateLegalRegistrationRequest) (*partyproto.CreateLegalRegistrationResponse, error) {
	legalRegistration, err := ps.ProcessLegalRegistrationRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertLegalRegistration(ctx, insertLegalRegistrationSQL, legalRegistration, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	legalRegistrationResponse := partyproto.CreateLegalRegistrationResponse{}
	legalRegistrationResponse.LegalRegistration = legalRegistration
	return &legalRegistrationResponse, nil
}

// ProcessLegalRegistrationRequest - ProcessLegalRegistrationRequest
func (ps *PartyService) ProcessLegalRegistrationRequest(ctx context.Context, in *partyproto.CreateLegalRegistrationRequest) (*partyproto.LegalRegistration, error) {
	legalRegistration := partyproto.LegalRegistration{}
	legalRegistration.LegalRegistrationAdditionalInformation = in.LegalRegistrationAdditionalInformation
	legalRegistration.LegalRegistrationNumber = in.LegalRegistrationNumber
	legalRegistration.LegalRegistrationType = in.LegalRegistrationType

	return &legalRegistration, nil
}

// insertLegalRegistration - Insert LegalRegistration into database
func (ps *PartyService) insertLegalRegistration(ctx context.Context, insertLegalRegistrationSQL string, legalRegistration *partyproto.LegalRegistration, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLegalRegistrationSQL, legalRegistration)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		legalRegistration.Id = uint32(uID)

		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
