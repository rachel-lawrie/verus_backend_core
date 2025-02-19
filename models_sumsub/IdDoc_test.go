package models_sumsub

import (
	"testing"

	"github.com/rachel-lawrie/verus_backend_core/constants"
	"github.com/stretchr/testify/assert"
)

func TestIdDocTypeString(t *testing.T) {
	tests := []struct {
		docType  IdDocType
		expected string
	}{
		{IDCard, constants.SUMSUB_ID_DOC_TYPE_ID_CARD},
		{Passport, constants.SUMSUB_ID_DOC_TYPE_PASSPORT},
		{Drivers, constants.SUMSUB_ID_DOC_TYPE_DRIVERS},
		{ResidencePermit, constants.SUMSUB_ID_DOC_TYPE_RESIDENCE_PERMIT},
		{UtilityBill, constants.SUMSUB_ID_DOC_TYPE_UTILITY_BILL},
		{Selfie, constants.SUMSUB_ID_DOC_TYPE_SELFIE},
		{VideoSelfie, constants.SUMSUB_ID_DOC_TYPE_VIDEO_SELFIE},
		{ProfileImage, constants.SUMSUB_ID_DOC_TYPE_PROFILE_IMAGE},
		{IdDocPhoto, constants.SUMSUB_ID_DOC_TYPE_ID_DOC_PHOTO},
		{Agreement, constants.SUMSUB_ID_DOC_TYPE_AGREEMENT},
		{Contract, constants.SUMSUB_ID_DOC_TYPE_CONTRACT},
		{DriversTranslation, constants.SUMSUB_ID_DOC_TYPE_DRIVERS_TRANSLATION},
		{InvesterDoc, constants.SUMSUB_ID_DOC_TYPE_INVESTER_DOC},
		{VehicleRegistrationCertificate, constants.SUMSUB_ID_DOC_TYPE_VEHICLE_REGISTRATION_CERTIFICATE},
		{IncomeSource, constants.SUMSUB_ID_DOC_TYPE_INCOME_SOURCE},
		{PaymentMethod, constants.SUMSUB_ID_DOC_TYPE_PAYMENT_METHOD},
		{BankCard, constants.SUMSUB_ID_DOC_TYPE_BANK_CARD},
		{CovidVaccinationForm, constants.SUMSUB_ID_DOC_TYPE_COVID_VACCINATION_FORM},
		{ArbitraryDoc, constants.SUMSUB_ID_DOC_TYPE_ARBITRARY_DOC},
		{Other, constants.SUMSUB_ID_DOC_TYPE_OTHER},
		{IdDocType(999), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.docType.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStringToIdDocType(t *testing.T) {
	tests := []struct {
		input    string
		expected IdDocType
		hasError bool
	}{
		{constants.SUMSUB_ID_DOC_TYPE_ID_CARD, IDCard, false},
		{constants.SUMSUB_ID_DOC_TYPE_PASSPORT, Passport, false},
		{constants.SUMSUB_ID_DOC_TYPE_DRIVERS, Drivers, false},
		{constants.SUMSUB_ID_DOC_TYPE_RESIDENCE_PERMIT, ResidencePermit, false},
		{constants.SUMSUB_ID_DOC_TYPE_UTILITY_BILL, UtilityBill, false},
		{constants.SUMSUB_ID_DOC_TYPE_SELFIE, Selfie, false},
		{constants.SUMSUB_ID_DOC_TYPE_VIDEO_SELFIE, VideoSelfie, false},
		{constants.SUMSUB_ID_DOC_TYPE_PROFILE_IMAGE, ProfileImage, false},
		{constants.SUMSUB_ID_DOC_TYPE_ID_DOC_PHOTO, IdDocPhoto, false},
		{constants.SUMSUB_ID_DOC_TYPE_AGREEMENT, Agreement, false},
		{constants.SUMSUB_ID_DOC_TYPE_CONTRACT, Contract, false},
		{constants.SUMSUB_ID_DOC_TYPE_DRIVERS_TRANSLATION, DriversTranslation, false},
		{constants.SUMSUB_ID_DOC_TYPE_INVESTER_DOC, InvesterDoc, false},
		{constants.SUMSUB_ID_DOC_TYPE_VEHICLE_REGISTRATION_CERTIFICATE, VehicleRegistrationCertificate, false},
		{constants.SUMSUB_ID_DOC_TYPE_INCOME_SOURCE, IncomeSource, false},
		{constants.SUMSUB_ID_DOC_TYPE_PAYMENT_METHOD, PaymentMethod, false},
		{constants.SUMSUB_ID_DOC_TYPE_BANK_CARD, BankCard, false},
		{constants.SUMSUB_ID_DOC_TYPE_COVID_VACCINATION_FORM, CovidVaccinationForm, false},
		{constants.SUMSUB_ID_DOC_TYPE_ARBITRARY_DOC, ArbitraryDoc, false},
		{constants.SUMSUB_ID_DOC_TYPE_OTHER, Other, false},
		{"INVALID", IdDocType(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, ok := StringToIdDocType[tt.input]
			if tt.hasError {
				assert.False(t, ok)
			} else {
				assert.True(t, ok)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestDocumentTypeToIdDocType(t *testing.T) {
	tests := []struct {
		input    string
		expected IdDocType
	}{
		{constants.DOCUMENT_TYPE_PASSPORT, Passport},
		{constants.DOCUMENT_TYPE_DRIVER_LICENSE, Drivers},
		{constants.DOCUMENT_TYPE_NATIONAL_ID, IDCard},
		{constants.DOCUMENT_TYPE_UTILITY_BILL, UtilityBill},
		{constants.DOCUMENT_TYPE_BANK_STATEMENT, BankCard},
		{constants.DOCUMENT_TYPE_BUSINESS_REGISTRATION, InvesterDoc},
		{constants.DOCUMENT_TYPE_TAX_DOCUMENT, IncomeSource},
		{constants.DOCUMENT_TYPE_PROOF_OF_ADDRESS, ResidencePermit},
		{constants.DOCUMENT_TYPE_FINANCIAL_STATEMENT, Agreement},
		{constants.DOCUMENT_TYPE_SELFIE, Selfie},
		{constants.DOCUMENT_TYPE_ID_CARD, IDCard},
		{constants.DOCUMENT_TYPE_OTHER, Other},
		{constants.DOCUMENT_TYPE_VIDEO_SELFIE, VideoSelfie},
		{"INVALID", Other},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := DocumentTypeToIdDocType[tt.input]
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIdDoc(t *testing.T) {
	doc := IdDoc{
		IdDocType:    constants.DOCUMENT_TYPE_PASSPORT,
		Country:      "US",
		FirstName:    "John",
		FirstNameEn:  "John",
		MiddleName:   "A",
		MiddleNameEn: "A",
		LastName:     "Doe",
		LastNameEn:   "Doe",
		DateOfBirth:  "1990-01-01",
	}

	assert.Equal(t, constants.DOCUMENT_TYPE_PASSPORT, doc.IdDocType)
	assert.Equal(t, "US", doc.Country)
	assert.Equal(t, "John", doc.FirstName)
	assert.Equal(t, "John", doc.FirstNameEn)
	assert.Equal(t, "A", doc.MiddleName)
	assert.Equal(t, "A", doc.MiddleNameEn)
	assert.Equal(t, "Doe", doc.LastName)
	assert.Equal(t, "Doe", doc.LastNameEn)
	assert.Equal(t, "1990-01-01", doc.DateOfBirth)
}
