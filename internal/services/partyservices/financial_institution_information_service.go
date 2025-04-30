package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertFinancialInstitutionInformationSQL = `insert into financial_institution_informations
	  (
financial_institution_branch_name,
financial_institution_name,
address,
financial_routing_number,
financial_routing_number_type_code,
financial_account_name,
financial_account_number,
financial_account_number_type_code,
transactional_party_id)
  values(
  :financial_institution_branch_name,
  :financial_institution_name,
  :address,
  :financial_routing_number,
  :financial_routing_number_type_code,
  :financial_account_name,
  :financial_account_number,
  :financial_account_number_type_code,
  :transactional_party_id );`

/*const selectFinancialInstitutionInformationsSQL = `select
  id,
  financial_institution_branch_name,
  financial_institution_name,
  address,
  financial_routing_number,
  financial_routing_number_type_code,
  financial_account_name,
  financial_account_number,
  financial_account_number_type_code,
  transactional_party_id from financial_institution_informations`*/

func (ps *PartyService) CreateFinancialInstitutionInformation(ctx context.Context, in *partyproto.CreateFinancialInstitutionInformationRequest) (*partyproto.CreateFinancialInstitutionInformationResponse, error) {
	financialInstitutionInformation, err := ps.ProcessFinancialInstitutionInformationRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertFinancialInstitutionInformation(ctx, insertFinancialInstitutionInformationSQL, financialInstitutionInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	financialInstitutionInformationResponse := partyproto.CreateFinancialInstitutionInformationResponse{}
	financialInstitutionInformationResponse.FinancialInstitutionInformation = financialInstitutionInformation
	return &financialInstitutionInformationResponse, nil
}

// ProcessFinancialInstitutionInformationRequest - ProcessFinancialInstitutionInformationRequest
func (ps *PartyService) ProcessFinancialInstitutionInformationRequest(ctx context.Context, in *partyproto.CreateFinancialInstitutionInformationRequest) (*partyproto.FinancialInstitutionInformation, error) {
	financialInstitutionInformation := partyproto.FinancialInstitutionInformation{}
	financialInstitutionInformation.FinancialInstitutionBranchName = in.FinancialInstitutionBranchName
	financialInstitutionInformation.FinancialInstitutionName = in.FinancialInstitutionName
	financialInstitutionInformation.Address = in.Address
	financialInstitutionInformation.FinancialRoutingNumber = in.FinancialRoutingNumber
	financialInstitutionInformation.FinancialRoutingNumberTypeCode = in.FinancialRoutingNumberTypeCode
	financialInstitutionInformation.FinancialAccountName = in.FinancialAccountName
	financialInstitutionInformation.FinancialAccountNumber = in.FinancialAccountNumber
	financialInstitutionInformation.FinancialAccountNumberTypeCode = in.FinancialAccountNumberTypeCode
	financialInstitutionInformation.TransactionalPartyId = in.TransactionalPartyId

	return &financialInstitutionInformation, nil
}

// insertFinancialInstitutionInformation - Insert FinancialInstitutionInformation into database
func (ps *PartyService) insertFinancialInstitutionInformation(ctx context.Context, insertFinancialInstitutionInformationSQL string, financialInstitutionInformation *partyproto.FinancialInstitutionInformation, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertFinancialInstitutionInformationSQL, financialInstitutionInformation)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		financialInstitutionInformation.Id = uint32(uID)

		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
