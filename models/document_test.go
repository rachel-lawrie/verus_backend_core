package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDocumentStatus(t *testing.T) {
	tests := []struct {
		input    string
		expected DocumentStatus
		hasError bool
	}{
		{"uploaded", DocumentUploaded, false},
		{"verified", DocumentVerified, false},
		{"rejected", DocumentRejected, false},
		{"uploadpending", DocumentUploadPending, false},
		{"invalid", DocumentStatus(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDocumentStatus(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestParseDocumentType(t *testing.T) {
	tests := []struct {
		input    string
		expected DocumentType
		hasError bool
	}{
		{"PASSPORT", DocumentPassport, false},
		{"DRIVERLICENSE", DocumentDriverLicense, false},
		{"NATIONALID", DocumentNationalID, false},
		{"UTILITYBILL", DocumentUtilityBill, false},
		{"BANKSTATEMENT", DocumentBankStatement, false},
		{"BUSINESSREGISTRATION", DocumentBusinessRegistration, false},
		{"TAXDOCUMENT", DocumentTaxDocument, false},
		{"PROOFOFADDRESS", DocumentProofOfAddress, false},
		{"FINANCIALSTATEMENT", DocumentFinancialStatement, false},
		{"invalid", DocumentType(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDocumentType(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestDocumentStatusString(t *testing.T) {
	tests := []struct {
		status   DocumentStatus
		expected string
	}{
		{DocumentUploaded, "uploaded"},
		{DocumentVerified, "verified"},
		{DocumentRejected, "rejected"},
		{DocumentUploadPending, "uploadpending"},
		{DocumentStatus(999), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.status.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDocumentTypeString(t *testing.T) {
	tests := []struct {
		docType  DocumentType
		expected string
	}{
		{DocumentPassport, "PASSPORT"},
		{DocumentDriverLicense, "DRIVERLICENSE"},
		{DocumentNationalID, "NATIONALID"},
		{DocumentUtilityBill, "UTILITYBILL"},
		{DocumentBankStatement, "BANKSTATEMENT"},
		{DocumentBusinessRegistration, "BUSINESSREGISTRATION"},
		{DocumentTaxDocument, "TAXDOCUMENT"},
		{DocumentProofOfAddress, "PROOFOFADDRESS"},
		{DocumentFinancialStatement, "FINANCIALSTATEMENT"},
		{DocumentType(999), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.docType.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDocument(t *testing.T) {
	now := time.Now()
	doc := Document{
		DocumentID:   "doc12345",
		ApplicantID:  "applicant123",
		DocumentType: DocumentPassport,
		Country:      "US",
		FileURL:      "https://example.com/doc12345.pdf",
		FileSize:     2048,
		Status:       DocumentUploaded,
		CreatedAt:    now,
		UpdatedAt:    now,
		Deleted:      false,
	}

	assert.Equal(t, "doc12345", doc.DocumentID)
	assert.Equal(t, "applicant123", doc.ApplicantID)
	assert.Equal(t, DocumentPassport, doc.DocumentType)
	assert.Equal(t, "US", doc.Country)
	assert.Equal(t, "https://example.com/doc12345.pdf", doc.FileURL)
	assert.Equal(t, int64(2048), doc.FileSize)
	assert.Equal(t, DocumentUploaded, doc.Status)
	assert.Equal(t, now, doc.CreatedAt)
	assert.Equal(t, now, doc.UpdatedAt)
	assert.False(t, doc.Deleted)
}
