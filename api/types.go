package api

// StudyList represents the top-level JSON object which is a list of studies.
type StudyList struct {
	Studies []Study `json:"studies"`
}

// Top-level struct for the entire JSON object
type Study struct {
	ProtocolSection ProtocolSection `json:"protocolSection"`
	DerivedSection  DerivedSection  `json:"derivedSection"`
	HasResults      bool            `json:"hasResults"`
}

// ProtocolSection represents the protocol details of the study.
type ProtocolSection struct {
	IdentificationModule       IdentificationModule       `json:"identificationModule"`
	StatusModule               StatusModule               `json:"statusModule"`
	SponsorCollaboratorsModule SponsorCollaboratorsModule `json:"sponsorCollaboratorsModule"`
	OversightModule            OversightModule            `json:"oversightModule"`
	DescriptionModule          DescriptionModule          `json:"descriptionModule"`
	ConditionsModule           ConditionsModule           `json:"conditionsModule"`
	DesignModule               DesignModule               `json:"designModule"`
	ArmsInterventionsModule    ArmsInterventionsModule    `json:"armsInterventionsModule"`
	OutcomesModule             OutcomesModule             `json:"outcomesModule"`
	EligibilityModule          EligibilityModule          `json:"eligibilityModule"`
	ContactsLocationsModule    ContactsLocationsModule    `json:"contactsLocationsModule"`
	ReferencesModule           ReferencesModule           `json:"referencesModule"`
}

// IdentificationModule contains identifying information for the study.
type IdentificationModule struct {
	NctID            string            `json:"nctId"`
	OrgStudyIDInfo   OrgStudyIDInfo    `json:"orgStudyIdInfo"`
	SecondaryIdInfos []SecondaryIdInfo `json:"secondaryIdInfos"`
	Organization     Organization      `json:"organization"`
	BriefTitle       string            `json:"briefTitle"`
	OfficialTitle    string            `json:"officialTitle"`
	Acronym          string            `json:"acronym"`
}

// OrgStudyIDInfo holds the organization's study ID.
type OrgStudyIDInfo struct {
	ID string `json:"id"`
}

// SecondaryIdInfo holds secondary identification information.
type SecondaryIdInfo struct {
	ID string `json:"id"`
}

// Organization holds information about the organizing entity.
type Organization struct {
	FullName string `json:"fullName"`
	Class    string `json:"class"`
}

// StatusModule contains information about the study's status and dates.
type StatusModule struct {
	StatusVerifiedDate          string             `json:"statusVerifiedDate"` // Consider time.Time if format is consistent
	OverallStatus               string             `json:"overallStatus"`
	ExpandedAccessInfo          ExpandedAccessInfo `json:"expandedAccessInfo"`
	StartDateStruct             DateStruct         `json:"startDateStruct"`
	PrimaryCompletionDateStruct DateStruct         `json:"primaryCompletionDateStruct"`
	CompletionDateStruct        DateStruct         `json:"completionDateStruct"`
	StudyFirstSubmitDate        string             `json:"studyFirstSubmitDate"`   // Consider time.Time if format is consistent
	StudyFirstSubmitQcDate      string             `json:"studyFirstSubmitQcDate"` // Consider time.Time if format is consistent
	StudyFirstPostDateStruct    DateStruct         `json:"studyFirstPostDateStruct"`
	LastUpdateSubmitDate        string             `json:"lastUpdateSubmitDate"` // Consider time.Time if format is consistent
	LastUpdatePostDateStruct    DateStruct         `json:"lastUpdatePostDateStruct"`
}

// ExpandedAccessInfo holds information about expanded access.
type ExpandedAccessInfo struct {
	HasExpandedAccess bool `json:"hasExpandedAccess"`
}

// DateStruct represents a date with an optional type.
type DateStruct struct {
	Date string `json:"date"` // Consider time.Time if format is consistent
	Type string `json:"type,omitempty"`
}

// SponsorCollaboratorsModule contains information about sponsors and collaborators.
type SponsorCollaboratorsModule struct {
	ResponsibleParty ResponsibleParty `json:"responsibleParty"`
	LeadSponsor      LeadSponsor      `json:"leadSponsor"`
}

// ResponsibleParty holds information about the responsible party.
type ResponsibleParty struct {
	OldNameTitle    string `json:"oldNameTitle"`
	OldOrganization string `json:"oldOrganization"`
}

// LeadSponsor holds information about the lead sponsor.
type LeadSponsor struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

// OversightModule holds information about oversight.
type OversightModule struct {
	OversightHasDmc bool `json:"oversightHasDmc"`
}

// DescriptionModule holds the brief summary of the study.
type DescriptionModule struct {
	BriefSummary string `json:"briefSummary"`
}

// ConditionsModule holds information about the conditions studied.
type ConditionsModule struct {
	Conditions []string `json:"conditions"`
}

// DesignModule holds information about the study design.
type DesignModule struct {
	StudyType      string         `json:"studyType"`
	Phases         []string       `json:"phases"`
	DesignInfo     DesignInfo     `json:"designInfo"`
	EnrollmentInfo EnrollmentInfo `json:"enrollmentInfo"`
}

// DesignInfo holds detailed design information.
type DesignInfo struct {
	Allocation        string      `json:"allocation"`
	InterventionModel string      `json:"interventionModel"`
	PrimaryPurpose    string      `json:"primaryPurpose"`
	MaskingInfo       MaskingInfo `json:"maskingInfo"`
}

// MaskingInfo holds information about the masking (blinding) of the study.
type MaskingInfo struct {
	Masking   string   `json:"masking"`
	WhoMasked []string `json:"whoMasked"`
}

// EnrollmentInfo holds information about the study enrollment.
type EnrollmentInfo struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// ArmsInterventionsModule holds information about arm groups and interventions.
type ArmsInterventionsModule struct {
	ArmGroups     []ArmGroup     `json:"armGroups"`
	Interventions []Intervention `json:"interventions"`
}

// ArmGroup represents a group in the study design.
type ArmGroup struct {
	Label             string   `json:"label"`
	Type              string   `json:"type"`
	Description       string   `json:"description"`
	InterventionNames []string `json:"interventionNames"`
}

// Intervention represents an intervention applied in the study.
type Intervention struct {
	Type           string   `json:"type"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ArmGroupLabels []string `json:"armGroupLabels"`
}

// OutcomesModule holds information about primary and secondary outcomes.
type OutcomesModule struct {
	PrimaryOutcomes   []Outcome `json:"primaryOutcomes"`
	SecondaryOutcomes []Outcome `json:"secondaryOutcomes"`
}

// Outcome represents a study outcome measure.
type Outcome struct {
	Measure   string `json:"measure"`
	TimeFrame string `json:"timeFrame"`
}

// EligibilityModule holds information about eligibility criteria.
type EligibilityModule struct {
	EligibilityCriteria string   `json:"eligibilityCriteria"`
	HealthyVolunteers   bool     `json:"healthyVolunteers"`
	Sex                 string   `json:"sex"`
	MinimumAge          string   `json:"minimumAge"`
	MaximumAge          string   `json:"maximumAge"`
	StdAges             []string `json:"stdAges"`
}

// ContactsLocationsModule holds information about contacts and locations.
type ContactsLocationsModule struct {
	OverallOfficials []OverallOfficial `json:"overallOfficials"`
	Locations        []Location        `json:"locations"`
}

// OverallOfficial represents an overall official for the study.
type OverallOfficial struct {
	Name        string `json:"name"`
	Affiliation string `json:"affiliation"`
	Role        string `json:"role"`
}

// Location represents a study location.
type Location struct {
	Facility string   `json:"facility"`
	City     string   `json:"city"`
	State    string   `json:"state"`
	Zip      string   `json:"zip"`
	Country  string   `json:"country"`
	GeoPoint GeoPoint `json:"geoPoint"`
}

// GeoPoint represents geographic coordinates.
type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// ReferencesModule holds study references.
type ReferencesModule struct {
	References []Reference `json:"references"`
}

// Reference represents a single study reference.
type Reference struct {
	PMID     string `json:"pmid"` // Consider int if always numeric
	Type     string `json:"type"`
	Citation string `json:"citation"`
}

// DerivedSection holds information derived from the protocol section.
type DerivedSection struct {
	MiscInfoModule           MiscInfoModule           `json:"miscInfoModule"`
	ConditionBrowseModule    ConditionBrowseModule    `json:"conditionBrowseModule"`
	InterventionBrowseModule InterventionBrowseModule `json:"interventionBrowseModule"`
}

// MiscInfoModule holds miscellaneous derived information.
type MiscInfoModule struct {
	VersionHolder string `json:"versionHolder"` // Consider time.Time if format is consistent
}

// ConditionBrowseModule holds information about browsable conditions.
type ConditionBrowseModule struct {
	Meshes         []Mesh         `json:"meshes"`
	Ancestors      []BrowseTerm   `json:"ancestors"`
	BrowseLeaves   []BrowseLeaf   `json:"browseLeaves"`
	BrowseBranches []BrowseBranch `json:"browseBranches"`
}

// InterventionBrowseModule holds information about browsable interventions.
type InterventionBrowseModule struct {
	Meshes         []Mesh         `json:"meshes"`
	Ancestors      []BrowseTerm   `json:"ancestors"`
	BrowseLeaves   []BrowseLeaf   `json:"browseLeaves"`
	BrowseBranches []BrowseBranch `json:"browseBranches"`
}

// Mesh represents a Medical Subject Heading.
type Mesh struct {
	ID   string `json:"id"`
	Term string `json:"term"`
}

// BrowseTerm represents a term in a browsing hierarchy.
type BrowseTerm struct {
	ID   string `json:"id"`
	Term string `json:"term"`
}

// BrowseLeaf represents a leaf node in a browsing hierarchy.
type BrowseLeaf struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AsFound   string `json:"asFound,omitempty"`
	Relevance string `json:"relevance"`
}

// BrowseBranch represents a branch in a browsing hierarchy.
type BrowseBranch struct {
	Abbrev string `json:"abbrev"`
	Name   string `json:"name"`
}
