package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertAdditionalPartyIdentificationSQL = `insert into additional_party_identifications
	    (
	    additional_party_identification,
      additional_party_identification_type_code,
      code_list_version,
      gln,
      transactional_party_id
	    )
  values(
    :additional_party_identification,
    :additional_party_identification_type_code,
    :code_list_version,
    :gln,
    :transactional_party_id
   );`

/*const selectAdditionalPartyIdentificationsSQL = `select
  id,
  additional_party_identification,
  additional_party_identification_type_code,
  code_list_version,
  gln,
  transactional_party_id from additional_party_identifications`*/

func (ps *PartyService) CreateAdditionalPartyIdentification(ctx context.Context, in *partyproto.CreateAdditionalPartyIdentificationRequest) (*partyproto.CreateAdditionalPartyIdentificationResponse, error) {
	additionalPartyIdentification, err := ps.ProcessAdditionalPartyIdentificationRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertAdditionalPartyIdentification(ctx, insertAdditionalPartyIdentificationSQL, additionalPartyIdentification, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	additionalPartyIdentificationResponse := partyproto.CreateAdditionalPartyIdentificationResponse{}
	additionalPartyIdentificationResponse.AdditionalPartyIdentification = additionalPartyIdentification
	return &additionalPartyIdentificationResponse, nil
}

// ProcessAdditionalPartyIdentificationRequest - ProcessAdditionalPartyIdentificationRequest
func (ps *PartyService) ProcessAdditionalPartyIdentificationRequest(ctx context.Context, in *partyproto.CreateAdditionalPartyIdentificationRequest) (*partyproto.AdditionalPartyIdentification, error) {
	additionalPartyIdentification := partyproto.AdditionalPartyIdentification{}
	additionalPartyIdentification.AdditionalPartyIdentification = in.AdditionalPartyIdentification
	additionalPartyIdentification.AdditionalPartyIdentificationTypeCode = in.AdditionalPartyIdentificationTypeCode
	additionalPartyIdentification.CodeListVersion = in.CodeListVersion
	additionalPartyIdentification.Gln = in.Gln
	additionalPartyIdentification.TransactionalPartyId = in.TransactionalPartyId
	return &additionalPartyIdentification, nil
}

// insertAdditionalPartyIdentification - Insert AdditionalPartyIdentification into database
func (ps *PartyService) insertAdditionalPartyIdentification(ctx context.Context, insertAdditionalPartyIdentificationSQL string, additionalPartyIdentification *partyproto.AdditionalPartyIdentification, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertAdditionalPartyIdentificationSQL, additionalPartyIdentification)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		additionalPartyIdentification.Id = uint32(uID)
		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
