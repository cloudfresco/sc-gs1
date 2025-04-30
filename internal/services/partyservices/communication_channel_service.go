package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertCommunicationChannelSQL = `insert into communication_channels
	    (communication_channel_code,
communication_channel_name,
communication_value,
contact_id)
  values(
  :communication_channel_code,
  :communication_channel_name,
  :communication_value,
  :contact_id);`

/*const selectCommunicationChannelsSQL = `select
  id,
  communication_channel_code,
  communication_channel_name,
  communication_value,
  contact_id from communication_channels`*/

func (ps *PartyService) CreateCommunicationChannel(ctx context.Context, in *partyproto.CreateCommunicationChannelRequest) (*partyproto.CreateCommunicationChannelResponse, error) {
	communicationChannel, err := ps.ProcessCommunicationChannelRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertCommunicationChannel(ctx, insertCommunicationChannelSQL, communicationChannel, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	communicationChannelResponse := partyproto.CreateCommunicationChannelResponse{}
	communicationChannelResponse.CommunicationChannel = communicationChannel
	return &communicationChannelResponse, nil
}

// ProcessCommunicationChannelRequest - ProcessCommunicationChannelRequest
func (ps *PartyService) ProcessCommunicationChannelRequest(ctx context.Context, in *partyproto.CreateCommunicationChannelRequest) (*partyproto.CommunicationChannel, error) {
	communicationChannel := partyproto.CommunicationChannel{}
	communicationChannel.CommunicationChannelCode = in.CommunicationChannelCode
	communicationChannel.CommunicationChannelName = in.CommunicationChannelName
	communicationChannel.CommunicationValue = in.CommunicationValue
	communicationChannel.ContactId = in.ContactId
	return &communicationChannel, nil
}

// insertCommunicationChannel - Insert CommunicationChannel into database
func (ps *PartyService) insertCommunicationChannel(ctx context.Context, insertCommunicationChannelSQL string, communicationChannel *partyproto.CommunicationChannel, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertCommunicationChannelSQL, communicationChannel)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		communicationChannel.Id = uint32(uID)

		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
