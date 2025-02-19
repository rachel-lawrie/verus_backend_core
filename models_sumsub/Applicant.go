package models_sumsub

import "time"

type Applicant struct {
	ID                string    `json:"id,omitempty"`
	CreatedAt         string    `json:"createdAt,omitempty"`
	CreatedBy         string    `json:"createdBy,omitempty"`
	Key               string    `json:"sourceKey,omitempty"`
	ClientID          string    `json:"clientId,omitempty"`
	InspectionID      string    `json:"inspectionId,omitempty"`
	ExternalUserID    string    `json:"externalUserId,omitempty"`
	Info              FixedInfo `json:"info,omitempty"`
	FixedInfo         FixedInfo `json:"fixedInfo,omitempty"`
	ApplicantPlatform string    `json:"applicantPlatform,omitempty"`
	RequiredIdDocs    struct {
		DocSets []struct {
			IdDocSetType  string   `json:"idDocSetType,omitempty"`
			Types         []string `json:"idDocType,omitempty"`
			VideoRequired string   `json:"videoRequired,omitempty"`
		} `json:"docSets,omitempty"`
	} `json:"requiredIdDocs,omitempty"`
	Review struct {
		ReviewID              string    `json:"reviewId,omitempty"`
		AttemptID             string    `json:"attemptId,omitempty"`
		AttemptCnt            int       `json:"attemptCnt,omitempty"`
		LevelName             string    `json:"levelName,omitempty"`
		LevelAutoCheckMode    string    `json:"levelAutoCheckMode,omitempty"`
		CreateDate            string    `json:"createDate,omitempty"`
		ReviewStatus          string    `json:"reviewStatus,omitempty"`
		Priority              int       `json:"priority,omitempty"`
		ElapsedSincePendingMs int       `json:"elapsedSincePendingMs,omitempty"`
		ElapsedSinceQueuedMs  int       `json:"elapsedSinceQueuedMs,omitempty"`
		Reprocessing          bool      `json:"reprocessing,omitempty"`
		ReviewDate            time.Time `json:"reviewDate,omitempty"`
		StartDate             string    `json:"startDate,omitempty"`
		ReviewResult          struct {
			ReviewAnswer string `json:"reviewAnswer,omitempty"`
		} `json:"reviewResult,omitempty"`
		NotificationFailureCnt int `json:"notificationFailureCnt,omitempty"`
	} `json:"review,omitempty"`
	Lang            string          `json:"lang,omitempty"`
	Type            string          `json:"type,omitempty"`
	WebhookResponse WebhookResponse `json:"webhookResponse,omitempty"`
}

type WebhookResponse struct {
	ApplicantID    string `json:"applicantId,omitempty"`
	InspectionID   string `json:"inspectionId,omitempty"`
	CorrelationID  string `json:"correlationId,omitempty"`
	LevelName      string `json:"levelName,omitempty"`
	ExternalUserID string `json:"externalUserId,omitempty"`
	Type           string `json:"type,omitempty"`
	SandboxMode    bool   `json:"sandboxMode,omitempty"`
	ReviewStatus   string `json:"reviewStatus,omitempty"`
	CreatedAtMs    string `json:"createdAtMs,omitempty"`
	ClientID       string `json:"clientId,omitempty"`
	ApplicantType  string `json:"applicantType,omitempty"`
	ReviewResult   struct {
		ReviewAnswer     string   `json:"reviewAnswer,omitempty"`
		RejectLabels     []string `json:"rejectLabels,omitempty"`
		ReviewRejectType string   `json:"reviewRejectType,omitempty"`
		ButtonIds        []string `json:"buttonIds,omitempty"`
	} `json:"reviewResult,omitempty"`
	ApplicantActionID         string `json:"applicantActionId,omitempty"`
	ExternalApplicantActionID string `json:"externalApplicantActionId,omitempty"`
}
