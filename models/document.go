package models

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/rachel-lawrie/verus_backend_core/constants"
)

type Document struct {
	DocumentID   string         `bson:"document_id" json:"document_id"`     // Unique ID for the document
	ApplicantID  string         `bson:"applicant_id" json:"applicant_id"`   // ID of the associated user
	DocumentType DocumentType   `bson:"document_type" json:"document_type"` // Type of document (e.g., "passport", "utility bill")
	Country      string         `bson:"country" json:"country"`             // Country of the document
	FileURL      string         `bson:"file_url" json:"file_url"`           // URL of the uploaded file
	FileSize     int64          `bson:"file_size" json:"file_size"`         // Size of the uploaded file
	Status       DocumentStatus `bson:"status" json:"status"`               // Status of the document (e.g., "DocumentUploaded", "DocumentVerified", "DocumentRejected")
	CreatedAt    time.Time      `bson:"created_at" json:"created_at"`       // Timestamp of when the document was created
	UpdatedAt    time.Time      `bson:"updated_at" json:"updated_at"`       // Timestamp of when the document was last updated
	Deleted      bool           `bson:"deleted" json:"deleted"`             // Soft delete flag
	DeletedAt    *time.Time     `bson:"deleted_at" json:"deleted_at"`       // Timestamp of soft deletion
	DeletedBy    *string        `bson:"deleted_by" json:"deleted_by"`       // User or system who deleted the document
}

type DocumentStatus int

// Uploaded indicates that the document has been successfully uploaded.
const (
	DocumentUploaded DocumentStatus = iota
	DocumentVerified
	DocumentRejected
	DocumentUploadPending
)

// Map enum values to their string representations
var documentStatusToString = map[DocumentStatus]string{
	DocumentUploaded:      constants.DOCUMENT_STATUS_UPLOADED,
	DocumentVerified:      constants.DOCUMENT_STATUS_VERIFIED,
	DocumentRejected:      constants.DOCUMENT_STATUS_REJECTED,
	DocumentUploadPending: constants.DOCUMENT_STATUS_UPLOAD_PENDING,
}

// Map string representations back to enum values
var stringToDocumentStatus = map[string]DocumentStatus{
	constants.DOCUMENT_STATUS_UPLOADED:       DocumentUploaded,
	constants.DOCUMENT_STATUS_VERIFIED:       DocumentVerified,
	constants.DOCUMENT_STATUS_REJECTED:       DocumentRejected,
	constants.DOCUMENT_STATUS_UPLOAD_PENDING: DocumentUploadPending,
}

// String method for Status to get the name
func (s DocumentStatus) String() string {
	if val, ok := documentStatusToString[s]; ok {
		return val
	}
	return "Unknown"
}

// Parse string to Status enum
func ParseDocumentStatus(s string) (DocumentStatus, error) {
	val, ok := stringToDocumentStatus[strings.ToLower(s)]

	if ok {
		log.Printf("Parsing document status: %s", s)
		return val, nil
	}
	return DocumentStatus(0), errors.New("invalid status")
}

type DocumentType int

const (
	DocumentPassport DocumentType = iota
	DocumentDriverLicense
	DocumentNationalID
	DocumentUtilityBill
	DocumentBankStatement
	DocumentBusinessRegistration
	DocumentTaxDocument
	DocumentProofOfAddress
	DocumentFinancialStatement
	DocumentSelfie
	DocumentIDCard
	DocumentOther
	DocumentVideoSelfie
)

// Map enum values to their string representations
var documentTypeToString = map[DocumentType]string{
	DocumentPassport:             constants.DOCUMENT_TYPE_PASSPORT,
	DocumentDriverLicense:        constants.DOCUMENT_TYPE_DRIVER_LICENSE,
	DocumentNationalID:           constants.DOCUMENT_TYPE_NATIONAL_ID,
	DocumentUtilityBill:          constants.DOCUMENT_TYPE_UTILITY_BILL,
	DocumentBankStatement:        constants.DOCUMENT_TYPE_BANK_STATEMENT,
	DocumentBusinessRegistration: constants.DOCUMENT_TYPE_BUSINESS_REGISTRATION,
	DocumentTaxDocument:          constants.DOCUMENT_TYPE_TAX_DOCUMENT,
	DocumentProofOfAddress:       constants.DOCUMENT_TYPE_PROOF_OF_ADDRESS,
	DocumentFinancialStatement:   constants.DOCUMENT_TYPE_FINANCIAL_STATEMENT,
	DocumentSelfie:               constants.DOCUMENT_TYPE_SELFIE,
	DocumentIDCard:               constants.DOCUMENT_TYPE_ID_CARD,
	DocumentOther:                constants.DOCUMENT_TYPE_OTHER,
	DocumentVideoSelfie:          constants.DOCUMENT_TYPE_VIDEO_SELFIE,
}

// Map string representations back to enum values
var stringToDocumentType = map[string]DocumentType{
	constants.DOCUMENT_TYPE_PASSPORT:              DocumentPassport,
	constants.DOCUMENT_TYPE_DRIVER_LICENSE:        DocumentDriverLicense,
	constants.DOCUMENT_TYPE_NATIONAL_ID:           DocumentNationalID,
	constants.DOCUMENT_TYPE_UTILITY_BILL:          DocumentUtilityBill,
	constants.DOCUMENT_TYPE_BANK_STATEMENT:        DocumentBankStatement,
	constants.DOCUMENT_TYPE_BUSINESS_REGISTRATION: DocumentBusinessRegistration,
	constants.DOCUMENT_TYPE_TAX_DOCUMENT:          DocumentTaxDocument,
	constants.DOCUMENT_TYPE_PROOF_OF_ADDRESS:      DocumentProofOfAddress,
	constants.DOCUMENT_TYPE_FINANCIAL_STATEMENT:   DocumentFinancialStatement,
	constants.DOCUMENT_TYPE_SELFIE:                DocumentSelfie,
	constants.DOCUMENT_TYPE_ID_CARD:               DocumentIDCard,
	constants.DOCUMENT_TYPE_OTHER:                 DocumentOther,
	constants.DOCUMENT_TYPE_VIDEO_SELFIE:          DocumentVideoSelfie,
}

// String method for Status to get the name
func (s DocumentType) String() string {
	if val, ok := documentTypeToString[s]; ok {
		return val
	}
	return "Unknown"
}

// Parse string to Status enum
func ParseDocumentType(s string) (DocumentType, error) {
	if val, ok := stringToDocumentType[strings.ToUpper(s)]; ok {
		return val, nil
	}
	return DocumentType(0), errors.New("invalid type")
}
