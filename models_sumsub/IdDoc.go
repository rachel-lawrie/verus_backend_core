package models_sumsub

import (
	"errors"
	"strings"

	"github.com/rachel-lawrie/verus_backend_core/constants"
)

type IdDoc struct {
	IdDocType    string `json:"idDocType,omitempty"`
	Country      string `json:"country,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	FirstNameEn  string `json:"firstNameEn,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	MiddleNameEn string `json:"middleNameEn,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	LastNameEn   string `json:"lastNameEn,omitempty"`
	DateOfBirth  string `json:"dob,omitempty"` // yyyy-mm-dd format
}

type IdDocType int

// Enum values for IdDocType, source: https://developers.sumsub.com/api-reference/#id-documents
const (
	IDCard IdDocType = iota
	Passport
	Drivers
	ResidencePermit
	UtilityBill
	Selfie
	VideoSelfie
	ProfileImage
	IdDocPhoto
	Agreement
	Contract
	DriversTranslation
	InvesterDoc
	VehicleRegistrationCertificate
	IncomeSource
	PaymentMethod
	BankCard
	CovidVaccinationForm
	ArbitraryDoc
	Other
)

// Map enum values to their string representations
var IdDocTypeToString = map[IdDocType]string{
	IDCard:                         constants.SUMSUB_ID_DOC_TYPE_ID_CARD,
	Passport:                       constants.SUMSUB_ID_DOC_TYPE_PASSPORT,
	Drivers:                        constants.SUMSUB_ID_DOC_TYPE_DRIVERS,
	ResidencePermit:                constants.SUMSUB_ID_DOC_TYPE_RESIDENCE_PERMIT,
	UtilityBill:                    constants.SUMSUB_ID_DOC_TYPE_UTILITY_BILL,
	Selfie:                         constants.SUMSUB_ID_DOC_TYPE_SELFIE,
	VideoSelfie:                    constants.SUMSUB_ID_DOC_TYPE_VIDEO_SELFIE,
	ProfileImage:                   constants.SUMSUB_ID_DOC_TYPE_PROFILE_IMAGE,
	IdDocPhoto:                     constants.SUMSUB_ID_DOC_TYPE_ID_DOC_PHOTO,
	Agreement:                      constants.SUMSUB_ID_DOC_TYPE_AGREEMENT,
	Contract:                       constants.SUMSUB_ID_DOC_TYPE_CONTRACT,
	DriversTranslation:             constants.SUMSUB_ID_DOC_TYPE_DRIVERS_TRANSLATION,
	InvesterDoc:                    constants.SUMSUB_ID_DOC_TYPE_INVESTER_DOC,
	VehicleRegistrationCertificate: constants.SUMSUB_ID_DOC_TYPE_VEHICLE_REGISTRATION_CERTIFICATE,
	IncomeSource:                   constants.SUMSUB_ID_DOC_TYPE_INCOME_SOURCE,
	PaymentMethod:                  constants.SUMSUB_ID_DOC_TYPE_PAYMENT_METHOD,
	BankCard:                       constants.SUMSUB_ID_DOC_TYPE_BANK_CARD,
	CovidVaccinationForm:           constants.SUMSUB_ID_DOC_TYPE_COVID_VACCINATION_FORM,
	ArbitraryDoc:                   constants.SUMSUB_ID_DOC_TYPE_ARBITRARY_DOC,
	Other:                          constants.SUMSUB_ID_DOC_TYPE_OTHER,
}

// Map string representations back to enum values
var StringToIdDocType = map[string]IdDocType{
	constants.SUMSUB_ID_DOC_TYPE_ID_CARD:                          IDCard,
	constants.SUMSUB_ID_DOC_TYPE_PASSPORT:                         Passport,
	constants.SUMSUB_ID_DOC_TYPE_DRIVERS:                          Drivers,
	constants.SUMSUB_ID_DOC_TYPE_RESIDENCE_PERMIT:                 ResidencePermit,
	constants.SUMSUB_ID_DOC_TYPE_UTILITY_BILL:                     UtilityBill,
	constants.SUMSUB_ID_DOC_TYPE_SELFIE:                           Selfie,
	constants.SUMSUB_ID_DOC_TYPE_VIDEO_SELFIE:                     VideoSelfie,
	constants.SUMSUB_ID_DOC_TYPE_PROFILE_IMAGE:                    ProfileImage,
	constants.SUMSUB_ID_DOC_TYPE_ID_DOC_PHOTO:                     IdDocPhoto,
	constants.SUMSUB_ID_DOC_TYPE_AGREEMENT:                        Agreement,
	constants.SUMSUB_ID_DOC_TYPE_CONTRACT:                         Contract,
	constants.SUMSUB_ID_DOC_TYPE_DRIVERS_TRANSLATION:              DriversTranslation,
	constants.SUMSUB_ID_DOC_TYPE_INVESTER_DOC:                     InvesterDoc,
	constants.SUMSUB_ID_DOC_TYPE_VEHICLE_REGISTRATION_CERTIFICATE: VehicleRegistrationCertificate,
	constants.SUMSUB_ID_DOC_TYPE_INCOME_SOURCE:                    IncomeSource,
	constants.SUMSUB_ID_DOC_TYPE_PAYMENT_METHOD:                   PaymentMethod,
	constants.SUMSUB_ID_DOC_TYPE_BANK_CARD:                        BankCard,
	constants.SUMSUB_ID_DOC_TYPE_COVID_VACCINATION_FORM:           CovidVaccinationForm,
	constants.SUMSUB_ID_DOC_TYPE_ARBITRARY_DOC:                    ArbitraryDoc,
	constants.SUMSUB_ID_DOC_TYPE_OTHER:                            Other,
}

// String method for Status to get the name
func (s IdDocType) String() string {
	if val, ok := IdDocTypeToString[s]; ok {
		return val
	}
	return "Unknown"
}

// Parse string to Status enum
func ParseIdDocType(s string) (IdDocType, error) {
	if val, ok := StringToIdDocType[strings.Title(strings.ToUpper(s))]; ok {
		return val, nil
	}
	return IdDocType(0), errors.New("invalid type")
}

// map document type to idDocType
var DocumentTypeToIdDocType = map[string]IdDocType{
	constants.DOCUMENT_TYPE_PASSPORT:            Passport,
	constants.DOCUMENT_TYPE_DRIVER_LICENSE:      Drivers,
	constants.DOCUMENT_TYPE_ID_CARD:             IDCard,
	constants.DOCUMENT_TYPE_UTILITY_BILL:        UtilityBill,
	constants.DOCUMENT_TYPE_BANK_STATEMENT:      BankCard,
	constants.DOCUMENT_TYPE_TAX_DOCUMENT:        IncomeSource,
	constants.DOCUMENT_TYPE_PROOF_OF_ADDRESS:    ResidencePermit,
	constants.DOCUMENT_TYPE_FINANCIAL_STATEMENT: Agreement,
	constants.DOCUMENT_TYPE_SELFIE:              Selfie,
	constants.DOCUMENT_TYPE_VIDEO_SELFIE:        VideoSelfie,
	constants.DOCUMENT_TYPE_OTHER:               Other,
}
