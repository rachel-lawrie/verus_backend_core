package models

import (
	"errors"
	"time"

	"github.com/rachel-lawrie/verus_backend_core/constants"
	models_persona "github.com/rachel-lawrie/verus_backend_core/models_persona"
	sumsub "github.com/rachel-lawrie/verus_backend_core/models_sumsub"
)

// Applicant represents an applicant associated with a client
type Applicant struct {
	ApplicantID       string           `bson:"applicant_id" json:"applicant_id"` // Unique ID for the applicant
	FirstName         string           `bson:"first_name" json:"first_name"`     // First name of the applicant
	MiddleName        string           `bson:"middle_name" json:"middle_name"`   // Middle name of the applicant
	LastName          string           `bson:"last_name" json:"last_name"`
	Email             string           `bson:"email" json:"email"`                           // Applicant's email address
	Phone             string           `bson:"phone" json:"phone"`                           // Applicant's phone number
	ClientID          string           `bson:"client_id" json:"client_id"`                   // ID of the associated client (foreign key)
	VerificationLevel string           `bson:"verification_level" json:"verification_level"` // Name of the associated verification level (foreign key)
	ExternalUserId    string           `bson:"external_user_id" json:"external_user_id"`     // External user ID
	Status            ApplicantStatus  `bson:"status" json:"status"`                         // PENDING, IN_REVIEW, VERIFIED, REJECTED
	CreatedAt         time.Time        `bson:"created_at" json:"created_at"`                 // Creation timestamp
	UpdatedAt         time.Time        `bson:"updated_at" json:"updated_at"`                 // Last update timestamp
	Deleted           bool             `bson:"deleted" json:"deleted"`                       // Soft delete flag
	DeletedAt         *time.Time       `bson:"deleted_at" json:"deleted_at"`                 // Soft delete timestamp
	DeletedBy         *string          `bson:"deleted_by" json:"deleted_by"`                 // User/system that deleted the applicant
	EncryptedData     EncryptedData    `bson:"encrypted_data" json:"encrypted_data"`         // Encrypted fields (DOB, Address, and key)
	Documents         []Document       `bson:"documents" json:"documents"`                   // List of documents associated with the applicant
	SumsubApplicant   sumsub.Applicant `bson:"sumsub_applicant" json:"sumsub_applicant"`     // Sumsub applicant object
	Payload           Payload          `bson:"payload" json:"payload"`                       // Payload object
}

// Payload represents the payload associated with an applicant
type Payload struct {
	Persona models_persona.PersonaEvent `bson:"persona" json:"persona"` // Persona object
}

// EncryptedField represents an encrypted field with ciphertext and nonce
type EncryptedField struct {
	Ciphertext []byte `bson:"ciphertext" json:"ciphertext"` // Encrypted data
	Nonce      []byte `bson:"nonce" json:"nonce"`           // Nonce used for encryption
}

// EncryptedData contains the encrypted sensitive fields (DOB and Address) and the encrypted key
type EncryptedData struct {
	DOB          EncryptedField   `bson:"dob" json:"dob"`                     // Encrypted DOB
	Address      EncryptedAddress `bson:"address" json:"address"`             // Encrypted Address
	EncryptedKey []byte           `bson:"encrypted_key" json:"encrypted_key"` // Encrypted data encryption key (DEK)
}

// Address represents the address fields
type RawAddress struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

// EncryptedAddress represents the encrypted address fields
type EncryptedAddress struct {
	Line1      EncryptedField `bson:"line1" json:"line1"`
	Line2      EncryptedField `bson:"line2" json:"line2"`
	City       EncryptedField `bson:"city" json:"city"`
	Region     EncryptedField `bson:"region" json:"region"`
	PostalCode EncryptedField `bson:"postal_code" json:"postal_code"`
	Country    EncryptedField `bson:"country" json:"country"`
}

type ApplicantStatus int

const (
	ApplicantStatusPending  ApplicantStatus = iota //Applicant created but no docs uploaded.
	ApplicantStatusInReview                        //Applicant has uploaded all required docs and is under review.
	ApplicantStatusVerified                        //Applicant has been verified.
	ApplicantStatusRejected                        //Applicant has been rejected.
)

var applicantStatusToString = map[ApplicantStatus]string{
	ApplicantStatusPending:  constants.APPLICANT_STATUS_PENDING,
	ApplicantStatusInReview: constants.APPLICANT_STATUS_IN_REVIEW,
	ApplicantStatusVerified: constants.APPLICANT_STATUS_VERIFIED,
	ApplicantStatusRejected: constants.APPLICANT_STATUS_REJECTED,
}

var stringToApplicantStatus = map[string]ApplicantStatus{
	constants.APPLICANT_STATUS_PENDING:   ApplicantStatusPending,
	constants.APPLICANT_STATUS_IN_REVIEW: ApplicantStatusInReview,
	constants.APPLICANT_STATUS_VERIFIED:  ApplicantStatusVerified,
	constants.APPLICANT_STATUS_REJECTED:  ApplicantStatusRejected,
}

func (s ApplicantStatus) String() string {
	if val, ok := applicantStatusToString[s]; ok {
		return val
	}
	return "Unknown"
}

func ParseApplicantStatus(s string) (ApplicantStatus, error) {
	if val, ok := stringToApplicantStatus[s]; ok {
		return val, nil
	}
	return ApplicantStatusPending, errors.New("Unknown applicant status")
}
