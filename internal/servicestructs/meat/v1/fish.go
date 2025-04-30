package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
)

// FishDespatchAdviceLineItemExtension - struct FishDespatchAdviceLineItemExtension
type FishDespatchAdviceLineItemExtension struct {
	*meatproto.FishDespatchAdviceLineItemExtensionD
	*FishDespatchAdviceLineItemExtensionT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// FishDespatchAdviceLineItemExtensionT - struct FishDespatchAdviceLineItemExtensionT
type FishDespatchAdviceLineItemExtensionT struct {
	DateOfLanding   time.Time `protobuf:"bytes,1,opt,name=date_of_landing,json=dateOfLanding,proto3" json:"date_of_landing,omitempty"`
	DateOfSlaughter time.Time `protobuf:"bytes,2,opt,name=date_of_slaughter,json=dateOfSlaughter,proto3" json:"date_of_slaughter,omitempty"`
}

// FishCatchOrProductionDate - struct FishCatchOrProductionDate
type FishCatchOrProductionDate struct {
	*meatproto.FishCatchOrProductionDateD
	*FishCatchOrProductionDateT
}

// FishCatchOrProductionDateT - struct FishCatchOrProductionDateT
type FishCatchOrProductionDateT struct {
	CatchEndDate    time.Time `protobuf:"bytes,1,opt,name=catch_end_date,json=catchEndDate,proto3" json:"catch_end_date,omitempty"`
	CatchStartDate  time.Time `protobuf:"bytes,2,opt,name=catch_start_date,json=catchStartDate,proto3" json:"catch_start_date,omitempty"`
	FirstFreezeDate time.Time `protobuf:"bytes,3,opt,name=first_freeze_date,json=firstFreezeDate,proto3" json:"first_freeze_date,omitempty"`
	CatchDateTime   time.Time `protobuf:"bytes,4,opt,name=catch_date_time,json=catchDateTime,proto3" json:"catch_date_time,omitempty"`
}
