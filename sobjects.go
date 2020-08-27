package sfdc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

///
/////
/////
///// FIXME: Verify and fix null values handling in mapping structs
/////
/////
/////
/////

type SObjectsService struct {
	Service
}

type DescribeGlobalResponse struct {
	Encoding     string     `json:"encoding"`
	MaxBatchSize int        `json:"maxBatchSize"`
	SObjects     []Describe `json:"sobjects"`
}

type Describe struct {
	Activateable        bool   `json:"activateable"`
	Custom              bool   `json:"custom"`
	CustomSetting       bool   `json:"customSetting"`
	Createable          bool   `json:"createable"`
	Deletable           bool   `json:"deletable"`
	DeprecatedAndHidden bool   `json:"deprecatedAndHidden"`
	FeedEnabled         bool   `json:"feedEnabled"`
	KeyPrefix           string `json:"keyPrefix"`
	Label               string `json:"label"`
	LabelPlural         string `json:"labelPlural"`
	Layoutable          bool   `json:"layoutable"`
	Mergeable           bool   `json:"mergeable"`
	MRUEnabled          bool   `json:"mruEnabled"`
	Name                string `json:"name"`
	Queryable           bool   `json:"queryable"`
	Replicateable       bool   `json:"replicateable"`
	Searchable          bool   `json:"searchable"`
	Triggerable         bool   `json:"triggerable"`
	Undeletable         bool   `json:"undeletable"`
	Updateable          bool   `json:"updateable"`

	URLs map[string]string `json:"urls"`
}

type SObjectBasicInfo struct {
	ObjectDescribe Describe     `json:"objectDescribe"`
	RecentItems    []RecentItem `json:"recentItems"`
}

type RecentItem struct {
	Attributes RecordAttributes `json:"attributes"`
	ID         string           `json:"Id"`
	Name       string           `json:"Name"`
}

type RecordAttributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// SObjectDescribe (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type SObjectDescribe struct {
	ActionOverrides       []ActionOverride    `json:"actionOverrides"`
	Activateable          bool                `json:"activateable"`
	ChildRelationships    []ChildRelationship `json:"childRelationships"`
	CompactLayoutable     bool                `json:"compactLayoutable"`
	Createable            bool                `json:"createable"`
	Custom                bool                `json:"custom"`
	CustomSetting         bool                `json:"customSetting"`
	DeepCloneable         bool                `json:"deepCloneable"`
	DefaultImplementation string              `json:"defaultImplementation"`
	Deletable             bool                `json:"deletable"`
	DeprecatedAndHidden   bool                `json:"deprecatedAndHidden"`
	ExtendedBy            string              `json:"extendedBy"`
	ExtendedInterfaces    string              `json:"extendedInterfaces"`
	FieldEnabled          bool                `json:"fieldEnabled"`
	Fields                []FieldDescribe     `json:"fields"`
	HasSubtypes           bool                `json:"hasSubtypes"`
	ImplementedBy         string              `json:"implementedBy"`
	ImplementsInterfaces  string              `json:"implementsInterfaces"`
	IsInterface           bool                `json:"isInterface"`
	IsSubtype             bool                `json:"isSubtype"`
	KeyPrefix             string              `json:"keyPrefix"`
	Label                 string              `json:"label"`
	LabelPlural           string              `json:"labelPlural"`
	Layoutable            bool                `json:"layoutable"`
	Listviewable          bool                `json:"listviewable"`
	LookupLayoutable      bool                `json:"lookupLayoutable"`
	Mergeable             bool                `json:"mergeable"`
	MRUEnabled            bool                `json:"mruEnabled"`
	Name                  string              `json:"name"`
	NamedLayoutInfos      []NamedLayoutInfo   `json:"namedLayoutInfos"`
	NetworkScopeFieldName string              `json:"networkScopeFieldName"`
	Queryable             bool                `json:"queryable"`
	RecordTypeInfos       []RecordTypeInfo    `json:"recordTypeInfos"`
	Replicateable         bool                `json:"replicateable"`
	Retrieveable          bool                `json:"retrieveable"`
	Searchable            bool                `json:"searchable"`
	SearchLayoutable      bool                `json:"searchLayoutable"`
	SObjectDescribeOption string              `json:"sobjectDescribeOption"`
	SupportedScopes       []ScopeInfo         `json:"supportedScopes"`
	Triggerable           bool                `json:"triggerable"`
	Undeletable           bool                `json:"undeletable"`
	Updateable            bool                `json:"updateable"`
	URLs                  map[string]string   `json:"urls"`
}

type ScopeInfo struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

// RecordTypeInfo (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type RecordTypeInfo struct {
	Active                   bool   `json:"active"`
	Available                bool   `json:"available"`
	DefaultRecordTypeMapping bool   `json:"defaultRecordTypeMapping"`
	DeveloperName            string `json:"developerName"`
	Master                   bool   `json:"master"`
	Name                     string `json:name`
	RecordTypeID             string `json:"recordTypeId"`
}

// NamedLayoutInfo (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type NamedLayoutInfo struct {
	Name string `json:"name"`
}

// ActionOverride (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type ActionOverride struct {
	FormFactor         string `json:"formFactor"`
	IsAvailableInTouch bool   `json:"isAvailableInTouch"`
	Name               string `json:"name"`
	PageID             string `json:"pageId"`
	URL                string `json:"url"`
}

// FieldDescribe (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type FieldDescribe struct {
	Aggregatable                 bool               `json:"aggregatable"`
	AIPredictionField            bool               `json:"aiPredictionField"`
	AutoNumber                   bool               `json:"autoNumnber"`
	ByteLength                   int                `json:"byteLength"`
	Calculated                   bool               `json:"calculated"`
	CalculatedFormula            string             `json:"calculatedFormula"`
	CascadeDelete                bool               `json:"cascadeDelete"`
	CaseSensitive                bool               `json:"caseSensitive"`
	CompoundFieldName            string             `json:"compoundFieldName"`
	ControllerName               string             `json:"controllerName"`
	Createable                   bool               `json:"Createable"`
	Custom                       bool               `json:"custom"`
	DefaultValue                 json.RawMessage    `json:"defaultValue"`
	DefaultValueFormula          string             `json:"defaultValueFormula"`
	DefaultOnCreate              bool               `json:"defaultOnCreate"`
	DependentPicklist            bool               `json:"dependentPicklist"`
	DeprecatedAndHidden          bool               `json:"deprecatedAndHidden"`
	Digits                       int                `json:"digits"`
	DisplayLocationInDecimal     bool               `json:"displayLocationInDecimal"`
	Encrypted                    bool               `json:"encrypted"`
	ExternalID                   bool               `json:externalId`
	ExtraTypeInfo                string             `json:"extraTypeInfo"`
	Filterable                   bool               `json:"filterable"`
	FilteredLookupInfo           FilteredLookupInfo `json:"filteredLookupInfo"`
	FormulaTreatNullNumberAsZero bool               `json:"formulaTreatNullNumberAsZero"`
	Groupable                    bool               `json:"groupable"`
	HighScaleNumber              bool               `json:"highScaleNumber"`
	HtmlFormatted                bool               `json:"htmlFormatted"`
	IDLookup                     bool               `json:"idLookup"`
	InlineHelpText               string             `json:"inlineHelpText"`
	Label                        string             `json:"label"`
	Length                       int                `json:"length"`
	Mask                         string             `json:"mask"`
	MaskType                     string             `json:"maskType"`
	Name                         string             `json:"name"`
	NameField                    bool               `json:"nameField"`
	NamePointing                 bool               `json:"namePointing"`
	Nillable                     bool               `json:"nillable"`
	Permissionable               bool               `json:"permissionable"`
	PicklistValues               []PicklistEntry    `json:"picklistValues"`
	PolymorphicForeignKey        bool               `json:"polymorphicForeignKey"`
	Precision                    int                `json:"precision"`
	QueryByDistance              bool               `json:"queryByDistance"`
	ReferenceTargetField         string             `json:"referenceTargetField"`
	ReferenceTo                  []string           `json:"referenceTo"`
	RelationshipName             string             `json:"relationshipName"`
	RelationshipOrder            *int               `json:"RelationshipOrder"`
	RestrictedDelete             bool               `json:"restrictedDelete"`
	RestrictedPicklist           bool               `json:"restrictedPicklist"`
	Scale                        int                `json:"scale"`
	SearchPrefilterable          bool               `json:"searchPrefilterable"`
	SOAPType                     string             `json:"soapType"`
	Sortable                     bool               `json:"sortable"`
	Type                         string             `json:"type"`
	Unique                       bool               `json:"unique"`
	Updateable                   bool               `json:"updateable"`
	WriteRequiresMAsterRead      bool               `json:"writeRequiresMasterRead"`
}

// PicklistEntry (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type PicklistEntry struct {
	Active       bool   `json:"active"`
	ValidFor     string `json:"validFor"`
	DefaultValue bool   `json:"defaultValus"`
	Label        string `json:"label"`
	Value        string `json:"value"`
}

type FilteredLookupInfo struct {
	ControllingFields []string `json:"controllingFields"`
	Dependent         bool     `json:"dependent"`
	OptionalFilter    bool     `json:"optionalFilter"`
}

// ChildRelationship (see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm)
type ChildRelationship struct {
	Field               string   `json:"field"`
	DeprecatedAndHidden bool     `json:"deprecatedAndHidden"`
	CascadeDelete       bool     `json:"cascadeDelete"`
	ChildSObject        string   `json:"childSObject"`
	JunctionIDListNames []string `json:"junctionIdListNames"`
	JunctionReferenceTo []string `json:"junctionReferenceTo"`
	RelationshipName    string   `json:"relationshipName"`
}

// DescribeGlobal returns the description of the SObjects on the org
func (s *SObjectsService) DescribeGlobal(ctx context.Context) (*DescribeGlobalResponse, error) {
	req, err := s.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	var result DescribeGlobalResponse
	if err = s.client.Do(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BasicInfo returns basic information about the SObject with the given name
func (s *SObjectsService) BasicInfo(ctx context.Context, name string) (*SObjectBasicInfo, error) {
	req, err := s.NewRequest(http.MethodGet, name, nil)
	if err != nil {
		return nil, err
	}

	var info SObjectBasicInfo
	if err = s.client.Do(ctx, req, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// Describe returns the describe info for a specific SObject
func (s *SObjectsService) Describe(ctx context.Context, name string) (*SObjectDescribe, error) {
	req, err := s.NewRequest(http.MethodGet, fmt.Sprintf("%s/describe/", name), nil)
	if err != nil {
		return nil, err
	}

	var desc SObjectDescribe
	if err = s.client.Do(ctx, req, &desc); err != nil {
		return nil, err
	}

	return &desc, nil
}
