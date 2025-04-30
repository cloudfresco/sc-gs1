package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// MeatActivityHistory - struct MeatActivityHistory
type MeatActivityHistory struct {
	*meatproto.MeatActivityHistoryD
	*MeatActivityHistoryT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// MeatActivityHistoryT - struct MeatActivityHistoryT
type MeatActivityHistoryT struct {
	DateOfArrival   time.Time `protobuf:"bytes,1,opt,name=date_of_arrival,json=dateOfArrival,proto3" json:"date_of_arrival,omitempty"`
	DateOfDeparture time.Time `protobuf:"bytes,2,opt,name=date_of_departure,json=dateOfDeparture,proto3" json:"date_of_departure,omitempty"`
}

// MeatBreedingDetail - struct MeatBreedingDetail
type MeatBreedingDetail struct {
	*meatproto.MeatBreedingDetailD
	*MeatBreedingDetailT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// MeatBreedingDetailT - struct MeatBreedingDetailT
type MeatBreedingDetailT struct {
	DateOfBirth time.Time `protobuf:"bytes,1,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
}

// MeatSlaughteringDetail - struct MeatSlaughteringDetail
type MeatSlaughteringDetail struct {
	*meatproto.MeatSlaughteringDetailD
	*MeatSlaughteringDetailT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// MeatSlaughteringDetailT - struct MeatSlaughteringDetailT
type MeatSlaughteringDetailT struct {
	DateOfSlaughtering    time.Time `protobuf:"bytes,1,opt,name=date_of_slaughtering,json=dateOfSlaughtering,proto3" json:"date_of_slaughtering,omitempty"`
	OptimumMaturationDate time.Time `protobuf:"bytes,2,opt,name=optimum_maturation_date,json=optimumMaturationDate,proto3" json:"optimum_maturation_date,omitempty"`
}

// MeatWorkItemIdentification - struct MeatWorkItemIdentification
type MeatWorkItemIdentification struct {
	*meatproto.MeatWorkItemIdentificationD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
