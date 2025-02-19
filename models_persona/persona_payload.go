package models_persona

import "time"

type PersonaEvent struct {
	Data      PersonaData `json:"data"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type PersonaData struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes PersonaAttributes `json:"attributes"`
	Included   []PersonaIncluded `json:"included"`
}

type PersonaAttributes struct {
	Name      string         `json:"name"`
	Payload   PersonaPayload `json:"payload"`
	CreatedAt time.Time      `json:"createdAt"`
}

type PersonaPayload struct {
	Data PersonaPayloadData `json:"data"`
}

type PersonaPayloadData struct {
	Type          string                   `json:"type"`
	ID            string                   `json:"id"`
	Attributes    PersonaPayloadAttributes `json:"attributes"`
	Relationships PersonaRelationships     `json:"relationships"`
}

type PersonaPayloadAttributes struct {
	Status                 string           `json:"status"`
	ReferenceID            string           `json:"referenceId"`
	Note                   *string          `json:"note"`
	Behaviors              PersonaBehaviors `json:"behaviors"`
	Tags                   []string         `json:"tags"`
	Creator                string           `json:"creator"`
	ReviewerComment        *string          `json:"reviewerComment"`
	UpdatedAt              time.Time        `json:"updatedAt"`
	CreatedAt              time.Time        `json:"createdAt"`
	StartedAt              time.Time        `json:"startedAt"`
	CompletedAt            time.Time        `json:"completedAt"`
	FailedAt               *time.Time       `json:"failedAt"`
	MarkedForReviewAt      *time.Time       `json:"markedForReviewAt"`
	DecisionedAt           time.Time        `json:"decisionedAt"`
	ExpiredAt              *time.Time       `json:"expiredAt"`
	RedactedAt             *time.Time       `json:"redactedAt"`
	PreviousStepName       string           `json:"previousStepName"`
	NextStepName           string           `json:"nextStepName"`
	NameFirst              string           `json:"nameFirst"`
	NameMiddle             *string          `json:"nameMiddle"`
	NameLast               string           `json:"nameLast"`
	Birthdate              string           `json:"birthdate"`
	AddressStreet1         *string          `json:"addressStreet1"`
	AddressStreet2         *string          `json:"addressStreet2"`
	AddressCity            *string          `json:"addressCity"`
	AddressSubdivision     *string          `json:"addressSubdivision"`
	AddressSubdivisionAbbr *string          `json:"addressSubdivisionAbbr"`
	AddressPostalCode      *string          `json:"addressPostalCode"`
	AddressPostalCodeAbbr  *string          `json:"addressPostalCodeAbbr"`
	SocialSecurityNumber   *string          `json:"socialSecurityNumber"`
	IdentificationNumber   string           `json:"identificationNumber"`
	EmailAddress           *string          `json:"emailAddress"`
	PhoneNumber            *string          `json:"phoneNumber"`
	Fields                 PersonaFields    `json:"fields"`
}

type PersonaBehaviors struct {
	RequestSpoofAttempts   int     `json:"requestSpoofAttempts"`
	UserAgentSpoofAttempts int     `json:"userAgentSpoofAttempts"`
	DistractionEvents      int     `json:"distractionEvents"`
	HesitationBaseline     int     `json:"hesitationBaseline"`
	HesitationCount        int     `json:"hesitationCount"`
	HesitationTime         int     `json:"hesitationTime"`
	ShortcutCopies         int     `json:"shortcutCopies"`
	ShortcutPastes         int     `json:"shortcutPastes"`
	AutofillCancels        int     `json:"autofillCancels"`
	AutofillStarts         int     `json:"autofillStarts"`
	DevtoolsOpen           bool    `json:"devtoolsOpen"`
	CompletionTime         float64 `json:"completionTime"`
	HesitationPercentage   int     `json:"hesitationPercentage"`
	BehaviorThreatLevel    string  `json:"behaviorThreatLevel"`
}

type PersonaFields struct {
	CurrentGovernmentId  PersonaFieldGovernmentId `json:"currentGovernmentId"`
	SelectedCountryCode  PersonaFieldString       `json:"selectedCountryCode"`
	SelectedIdClass      PersonaFieldString       `json:"selectedIdClass"`
	AddressStreet1       PersonaFieldString       `json:"addressStreet1"`
	AddressStreet2       PersonaFieldString       `json:"addressStreet2"`
	AddressCity          PersonaFieldString       `json:"addressCity"`
	AddressSubdivision   PersonaFieldString       `json:"addressSubdivision"`
	AddressPostalCode    PersonaFieldString       `json:"addressPostalCode"`
	AddressCountryCode   PersonaFieldString       `json:"addressCountryCode"`
	Birthdate            PersonaFieldDate         `json:"birthdate"`
	EmailAddress         PersonaFieldString       `json:"emailAddress"`
	IdentificationClass  PersonaFieldString       `json:"identificationClass"`
	IdentificationNumber PersonaFieldString       `json:"identificationNumber"`
	NameFirst            PersonaFieldString       `json:"nameFirst"`
	NameMiddle           PersonaFieldString       `json:"nameMiddle"`
	NameLast             PersonaFieldString       `json:"nameLast"`
	PhoneNumber          PersonaFieldString       `json:"phoneNumber"`
	CurrentSelfie        PersonaFieldSelfie       `json:"currentSelfie"`
}

type PersonaFieldGovernmentId struct {
	Type  string `json:"type"`
	Value struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"value"`
}

type PersonaFieldString struct {
	Type  string  `json:"type"`
	Value *string `json:"value"`
}

type PersonaFieldDate struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PersonaFieldSelfie struct {
	Type  string `json:"type"`
	Value struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"value"`
}

type PersonaRelationships struct {
	Account                PersonaRelationshipDataWrapper      `json:"account"`
	Template               PersonaRelationshipDataWrapper      `json:"template"`
	InquiryTemplate        PersonaRelationshipDataWrapper      `json:"inquiryTemplate"`
	InquiryTemplateVersion PersonaRelationshipDataWrapper      `json:"inquiryTemplateVersion"`
	Transaction            PersonaRelationshipDataWrapper      `json:"transaction"`
	Reviewer               PersonaRelationshipDataWrapper      `json:"reviewer"`
	Reports                PersonaRelationshipDataWrapperSlice `json:"reports"`
	Verifications          PersonaRelationshipDataWrapperSlice `json:"verifications"`
	Sessions               PersonaRelationshipDataWrapperSlice `json:"sessions"`
	Documents              PersonaRelationshipDataWrapperSlice `json:"documents"`
	Selfies                PersonaRelationshipDataWrapperSlice `json:"selfies"`
}

type PersonaRelationshipDataWrapper struct {
	Data PersonaRelationshipData `json:"data"`
}

type PersonaRelationshipDataWrapperSlice struct {
	Data []PersonaRelationshipData `json:"data"`
}

type PersonaRelationshipData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type PersonaIncluded struct {
	Type          string                    `json:"type"`
	ID            string                    `json:"id"`
	Attributes    PersonaIncludedAttributes `json:"attributes"`
	Relationships PersonaRelationships      `json:"relationships"`
}

type PersonaIncludedAttributes struct {
	ReferenceID           string                       `json:"referenceId"`
	CreatedAt             time.Time                    `json:"createdAt"`
	UpdatedAt             time.Time                    `json:"updatedAt"`
	RedactedAt            *time.Time                   `json:"redactedAt"`
	AccountTypeName       string                       `json:"accountTypeName"`
	Fields                PersonaFields                `json:"fields"`
	NameFirst             string                       `json:"nameFirst"`
	NameMiddle            *string                      `json:"nameMiddle"`
	NameLast              string                       `json:"nameLast"`
	SocialSecurityNumber  *string                      `json:"socialSecurityNumber"`
	AddressStreet1        *string                      `json:"addressStreet1"`
	AddressStreet2        *string                      `json:"addressStreet2"`
	AddressCity           *string                      `json:"addressCity"`
	AddressSubdivision    *string                      `json:"addressSubdivision"`
	AddressPostalCode     *string                      `json:"addressPostalCode"`
	CountryCode           string                       `json:"countryCode"`
	Birthdate             string                       `json:"birthdate"`
	PhoneNumber           *string                      `json:"phoneNumber"`
	EmailAddress          *string                      `json:"emailAddress"`
	Tags                  []string                     `json:"tags"`
	IdentificationNumbers PersonaIdentificationNumbers `json:"identificationNumbers"`
}

type PersonaIdentificationNumbers struct {
	PP []struct{} `json:"pp"`
}
