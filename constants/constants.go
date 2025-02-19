package constants

const (
	CollectionApplicants         = "applicants"
	CollectionClients            = "clients"
	CollectionSecrets            = "secrets"
	CollectionDocuments          = "documents"
	CollectionAuditLogs          = "audit_logs"
	CollectionVerificationLevels = "verification_levels"
)

const (
	// ApplicantStatus
	APPLICANT_STATUS_PENDING   = "pending"
	APPLICANT_STATUS_IN_REVIEW = "in_review"
	APPLICANT_STATUS_VERIFIED  = "verified"
	APPLICANT_STATUS_REJECTED  = "rejected"

	// DocumentStatus
	DOCUMENT_STATUS_UPLOADED       = "uploaded"
	DOCUMENT_STATUS_VERIFIED       = "verified"
	DOCUMENT_STATUS_REJECTED       = "rejected"
	DOCUMENT_STATUS_UPLOAD_PENDING = "uploadpending"

	// Verus DocumentType
	DOCUMENT_TYPE_PASSPORT              = "PASSPORT"
	DOCUMENT_TYPE_DRIVER_LICENSE        = "DRIVER_LICENSE"
	DOCUMENT_TYPE_NATIONAL_ID           = "NATIONAL_ID"
	DOCUMENT_TYPE_UTILITY_BILL          = "UTILITY_BILL"
	DOCUMENT_TYPE_BANK_STATEMENT        = "BANK_STATEMENT"
	DOCUMENT_TYPE_BUSINESS_REGISTRATION = "BUSINESS_REGISTRATION"
	DOCUMENT_TYPE_TAX_DOCUMENT          = "TAX_DOCUMENT"
	DOCUMENT_TYPE_PROOF_OF_ADDRESS      = "PROOF_OF_ADDRESS"
	DOCUMENT_TYPE_FINANCIAL_STATEMENT   = "FINANCIAL_STATEMENT"
	DOCUMENT_TYPE_SELFIE                = "SELFIE"
	DOCUMENT_TYPE_ID_CARD               = "ID_CARD"
	DOCUMENT_TYPE_OTHER                 = "OTHER"
	DOCUMENT_TYPE_VIDEO_SELFIE          = "VIDEO_SELFIE"

	// Sumsub ID Doc Types
	SUMSUB_ID_DOC_TYPE_ID_CARD                          = "ID_CARD"
	SUMSUB_ID_DOC_TYPE_PASSPORT                         = "PASSPORT"
	SUMSUB_ID_DOC_TYPE_DRIVERS                          = "DRIVER_LICENSE"
	SUMSUB_ID_DOC_TYPE_RESIDENCE_PERMIT                 = "RESIDENCE_PERMIT"
	SUMSUB_ID_DOC_TYPE_UTILITY_BILL                     = "UTILITY_BILL"
	SUMSUB_ID_DOC_TYPE_SELFIE                           = "SELFIE"
	SUMSUB_ID_DOC_TYPE_VIDEO_SELFIE                     = "VIDEO_SELFIE"
	SUMSUB_ID_DOC_TYPE_PROFILE_IMAGE                    = "PROFILE_IMAGE"
	SUMSUB_ID_DOC_TYPE_ID_DOC_PHOTO                     = "ID_DOC_PHOTO"
	SUMSUB_ID_DOC_TYPE_AGREEMENT                        = "AGREEMENT"
	SUMSUB_ID_DOC_TYPE_CONTRACT                         = "CONTRACT"
	SUMSUB_ID_DOC_TYPE_DRIVERS_TRANSLATION              = "DRIVERS_TRANSLATION"
	SUMSUB_ID_DOC_TYPE_INVESTER_DOC                     = "INVESTER_DOC"
	SUMSUB_ID_DOC_TYPE_VEHICLE_REGISTRATION_CERTIFICATE = "VEHICLE_REGISTRATION_CERTIFICATE"
	SUMSUB_ID_DOC_TYPE_INCOME_SOURCE                    = "INCOME_SOURCE"
	SUMSUB_ID_DOC_TYPE_PAYMENT_METHOD                   = "PAYMENT_METHOD"
	SUMSUB_ID_DOC_TYPE_BANK_CARD                        = "BANK_CARD"
	SUMSUB_ID_DOC_TYPE_COVID_VACCINATION_FORM           = "COVID_VACCINATION_FORM"
	SUMSUB_ID_DOC_TYPE_ARBITRARY_DOC                    = "ARBITRARY_DOC"
	SUMSUB_ID_DOC_TYPE_OTHER                            = "OTHER"
)
