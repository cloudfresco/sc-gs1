package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrganisationSQL = `insert into organisations
	    (issued_capital,
ic_code_list_version,
ic_currency_code,
organisation_name,
official_address,
transactional_party_id)
  values(
    :issued_capital,
    :ic_code_list_version,
    :ic_currency_code,
    :organisation_name,
    :official_address,
    :transactional_party_id
   );`

/*const selectOrganisationsSQL = `select
  id,
  issued_capital,
  ic_code_list_version,
  ic_currency_code,
  organisation_name,
  official_address,
  transactional_party_id from organisations`*/

func (ps *PartyService) CreateOrganisation(ctx context.Context, in *partyproto.CreateOrganisationRequest) (*partyproto.CreateOrganisationResponse, error) {
	organisation, err := ps.ProcessOrganisationRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertOrganisation(ctx, insertOrganisationSQL, organisation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	organisationResponse := partyproto.CreateOrganisationResponse{}
	organisationResponse.Organisation = organisation
	return &organisationResponse, nil
}

// ProcessOrganisationRequest - ProcessOrganisationRequest
func (ps *PartyService) ProcessOrganisationRequest(ctx context.Context, in *partyproto.CreateOrganisationRequest) (*partyproto.Organisation, error) {
	organisation := partyproto.Organisation{}
	organisation.IssuedCapital = in.IssuedCapital
	organisation.ICCodeListVersion = in.ICCodeListVersion
	organisation.ICCurrencyCode = in.ICCurrencyCode
	organisation.OrganisationName = in.OrganisationName
	organisation.OfficialAddress = in.OfficialAddress
	organisation.TransactionalPartyId = in.TransactionalPartyId
	return &organisation, nil
}

// insertOrganisation - Insert Organisation into database
func (ps *PartyService) insertOrganisation(ctx context.Context, insertOrganisationSQL string, organisation *partyproto.Organisation, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertOrganisationSQL, organisation)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		organisation.Id = uint32(uID)

		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
