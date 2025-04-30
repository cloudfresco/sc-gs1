package logisticsservices

import (
	"context"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDespatchAdviceTransportInformationSQL = `insert into despatch_advice_transport_informations
	  (
route_id,
transport_means_id,
transport_means_name,
transport_means_type,
transport_mode_code,
bill_of_lading_number,
additional_consignment_identification,
additional_consignment_identification_type_code,
code_list_version,
ginc,
driver,
receiver,
receiver_id,
additional_shipment_identification,
additional_shipment_identification_type_code,
gsin,
despatch_advice_id)
  values(
  :route_id,
  :transport_means_id,
  :transport_means_name,
  :transport_means_type,
  :transport_mode_code,
  :bill_of_lading_number,
  :additional_consignment_identification,
  :additional_consignment_identification_type_code,
  :code_list_version,
  :ginc,
  :driver,
  :receiver,
  :receiver_id,
  :additional_shipment_identification,
  :additional_shipment_identification_type_code,
  :gsin,
  despatch_advice_id);`

/*const selectDespatchTransportInformationsSQL = `select
  id,
  route_id,
  transport_means_id,
  transport_means_name,
  transport_means_type,
  transport_mode_code,
  bill_of_lading_number,
  additional_consignment_identification,
  additional_consignment_identification_type_code,
  code_list_version,
  ginc,
  driver,
  receiver,
  receiver_id,
  additional_shipment_identification,
  additional_shipment_identification_type_code,
  gsin,
  despatch_advice_id from despatch_advice_transport_informations`*/

func (das *DespatchAdviceService) CreateDespatchAdviceTransportInformation(ctx context.Context, in *logisticsproto.CreateDespatchAdviceTransportInformationRequest) (*logisticsproto.CreateDespatchAdviceTransportInformationResponse, error) {
	despatchAdviceTransportInformation, err := das.ProcessDespatchAdviceTransportInformationRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdviceTransportInformation(ctx, insertDespatchAdviceTransportInformationSQL, despatchAdviceTransportInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceTransportInformationResponse := logisticsproto.CreateDespatchAdviceTransportInformationResponse{}
	despatchAdviceTransportInformationResponse.DespatchAdviceTransportInformation = despatchAdviceTransportInformation
	return &despatchAdviceTransportInformationResponse, nil
}

// ProcessDespatchAdviceTransportInformationRequest - ProcessDespatchAdviceTransportInformationRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceTransportInformationRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceTransportInformationRequest) (*logisticsproto.DespatchAdviceTransportInformation, error) {
	despatchAdviceTransportInformation := logisticsproto.DespatchAdviceTransportInformation{}
	despatchAdviceTransportInformation.RouteId = in.RouteId
	despatchAdviceTransportInformation.TransportMeansId = in.TransportMeansId
	despatchAdviceTransportInformation.TransportMeansName = in.TransportMeansName
	despatchAdviceTransportInformation.TransportMeansType = in.TransportMeansType
	despatchAdviceTransportInformation.TransportModeCode = in.TransportModeCode
	despatchAdviceTransportInformation.BillOfLadingNumber = in.BillOfLadingNumber
	despatchAdviceTransportInformation.AdditionalConsignmentIdentification = in.AdditionalConsignmentIdentification
	despatchAdviceTransportInformation.AdditionalConsignmentIdentificationTypeCode = in.AdditionalConsignmentIdentificationTypeCode
	despatchAdviceTransportInformation.CodeListVersion = in.CodeListVersion
	despatchAdviceTransportInformation.Ginc = in.Ginc
	despatchAdviceTransportInformation.Driver = in.Driver
	despatchAdviceTransportInformation.DriverId = in.DriverId
	despatchAdviceTransportInformation.Receiver = in.Receiver
	despatchAdviceTransportInformation.ReceiverId = in.ReceiverId
	despatchAdviceTransportInformation.AdditionalShipmentIdentification = in.AdditionalShipmentIdentification
	despatchAdviceTransportInformation.AdditionalShipmentIdentificationTypeCode = in.AdditionalShipmentIdentificationTypeCode
	despatchAdviceTransportInformation.Gsin = in.Gsin
	despatchAdviceTransportInformation.DespatchAdviceId = in.DespatchAdviceId
	return &despatchAdviceTransportInformation, nil
}

// insertDespatchAdviceTransportInformation - Insert DespatchAdviceTransportInformation into database
func (das *DespatchAdviceService) insertDespatchAdviceTransportInformation(ctx context.Context, insertDespatchAdviceTransportInformationSQL string, despatchAdviceTransportInformation *logisticsproto.DespatchAdviceTransportInformation, userEmail string, requestID string) error {
	err := das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertDespatchAdviceTransportInformationSQL, despatchAdviceTransportInformation)
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
