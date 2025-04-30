package partyservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertContactSQL = `insert into contacts
	    (
  uuid4,
  contact_type_code,
  department_name,
  job_title,
  person_name,
  transactional_party_id)
  values(
  :uuid4,
  :contact_type_code,
  :department_name,
  :job_title,
  :person_name,
  :transactional_party_id);`

/*const selectContactsSQL = `select
  id,
  uuid4,
  contact_type_code,
  department_name,
  job_title,
  person_name,
  transactional_party_id from contacts`*/

func (ps *PartyService) CreateContact(ctx context.Context, in *partyproto.CreateContactRequest) (*partyproto.CreateContactResponse, error) {
	contact, err := ps.ProcessContactRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertContact(ctx, insertContactSQL, contact, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	contactResponse := partyproto.CreateContactResponse{}
	contactResponse.Contact = contact
	return &contactResponse, nil
}

// ProcessContactRequest - ProcessContactRequest
func (ps *PartyService) ProcessContactRequest(ctx context.Context, in *partyproto.CreateContactRequest) (*partyproto.Contact, error) {
	var err error
	contact := partyproto.Contact{}
	contact.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	contact.ContactTypeCode = in.ContactTypeCode
	contact.DepartmentName = in.DepartmentName
	contact.JobTitle = in.JobTitle
	contact.PersonName = in.PersonName
	contact.TransactionalPartyId = in.TransactionalPartyId
	return &contact, nil
}

// insertContact - Insert Contact into database
func (ps *PartyService) insertContact(ctx context.Context, insertContactSQL string, contact *partyproto.Contact, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertContactSQL, contact)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uuid4Str, err := common.UUIDBytesToStr(contact.Uuid4)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		contact.IdS = uuid4Str
		contact.Id = uint32(uID)

		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
