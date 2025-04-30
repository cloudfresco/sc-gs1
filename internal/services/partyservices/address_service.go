package partyservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"

	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertAddressSQL = `insert into addresses
	    (
	    uuid4,
      city,
      city_code,
      country_code,
      county_code,
      cross_street,
      currency_of_party_code,
      language_of_the_party_code,
      name,
      po_box_number,
      postal_code,
      province_code,
      state,
      street_address_one,
      street_address_three,
      street_address_two,
      latitude,
      longitude,
      transactional_party_id)
  values(
  :uuid4,
  :city,
  :city_code,
  :country_code,
  :county_code,
  :cross_street,
  :currency_of_party_code,
  :language_of_the_party_code,
  :name,
  :po_box_number,
  :postal_code,
  :province_code,
  :state,
  :street_address_one,
  :street_address_three,
  :street_address_two,
  :latitude,
  :longitude,
  :transactional_party_id);`

/*const selectAddressesSQL = `select
  id,
  uuid4,
  city,
  city_code,
  country_code,
  county_code,
  cross_street,
  currency_of_party_code,
  language_of_the_party_code,
  name,
  po_box_number,
  postal_code,
  province_code,
  state,
  street_address_one,
  street_address_three,
  street_address_two,
  latitude,
  longitude,
  transactional_party_id from addresses`*/

func (ps *PartyService) CreateAddress(ctx context.Context, in *partyproto.CreateAddressRequest) (*partyproto.CreateAddressResponse, error) {
	address, err := ps.ProcessAddressRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertAddress(ctx, insertAddressSQL, address, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	addressResponse := partyproto.CreateAddressResponse{}
	addressResponse.Address = address
	return &addressResponse, nil
}

// ProcessAddressRequest - ProcessAddressRequest
func (ps *PartyService) ProcessAddressRequest(ctx context.Context, in *partyproto.CreateAddressRequest) (*partyproto.Address, error) {
	var err error
	address := partyproto.Address{}
	address.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	address.City = in.City
	address.CityCode = in.CityCode
	address.CountryCode = in.CountryCode
	address.CountyCode = in.CountyCode
	address.CrossStreet = in.CrossStreet
	address.CurrencyOfPartyCode = in.CurrencyOfPartyCode
	address.LanguageOfThePartyCode = in.LanguageOfThePartyCode
	address.Name = in.Name
	address.POBoxNumber = in.POBoxNumber
	address.PostalCode = in.PostalCode
	address.ProvinceCode = in.ProvinceCode
	address.State = in.State
	address.StreetAddressOne = in.StreetAddressOne
	address.StreetAddressThree = in.StreetAddressThree
	address.StreetAddressTwo = in.StreetAddressTwo
	address.Latitude = in.Latitude
	address.Longitude = in.Longitude
	address.TransactionalPartyId = in.TransactionalPartyId
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &address, nil
}

// insertAddress - Insert Address into database
func (ps *PartyService) insertAddress(ctx context.Context, insertAddressSQL string, address *partyproto.Address, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertAddressSQL, address)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		address.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(address.Uuid4)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		address.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
