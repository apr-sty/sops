package infrastructureinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// MetricsSourceType enumerates the values for metrics source type.
type MetricsSourceType string

const (
	// PhysicalNode ...
	PhysicalNode MetricsSourceType = "PhysicalNode"
	// ResourceProvider ...
	ResourceProvider MetricsSourceType = "ResourceProvider"
	// VirtualMachine ...
	VirtualMachine MetricsSourceType = "VirtualMachine"
)

// PossibleMetricsSourceTypeValues returns an array of possible values for the MetricsSourceType const type.
func PossibleMetricsSourceTypeValues() []MetricsSourceType {
	return []MetricsSourceType{PhysicalNode, ResourceProvider, VirtualMachine}
}

// MetricsUnit enumerates the values for metrics unit.
type MetricsUnit string

const (
	// B ...
	B MetricsUnit = "B"
	// GB ...
	GB MetricsUnit = "GB"
	// KB ...
	KB MetricsUnit = "KB"
	// MB ...
	MB MetricsUnit = "MB"
	// One ...
	One MetricsUnit = "One"
	// Percentage ...
	Percentage MetricsUnit = "Percentage"
	// TB ...
	TB MetricsUnit = "TB"
)

// PossibleMetricsUnitValues returns an array of possible values for the MetricsUnit const type.
func PossibleMetricsUnitValues() []MetricsUnit {
	return []MetricsUnit{B, GB, KB, MB, One, Percentage, TB}
}

// Alert this class models an alert resource.
type Alert struct {
	autorest.Response `json:"-"`
	// AlertModel - Holds all information related to Alerts
	*AlertModel `json:"properties,omitempty"`
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Alert.
func (a Alert) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AlertModel != nil {
		objectMap["properties"] = a.AlertModel
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Alert struct.
func (a *Alert) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var alertModel AlertModel
				err = json.Unmarshal(*v, &alertModel)
				if err != nil {
					return err
				}
				a.AlertModel = &alertModel
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		}
	}

	return nil
}

// AlertList a pageable list of Alerts
type AlertList struct {
	autorest.Response `json:"-"`
	// Value - Holds all alerts in this page.
	Value *[]Alert `json:"value,omitempty"`
	// NextLink - Points to the next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// AlertListIterator provides access to a complete listing of Alert values.
type AlertListIterator struct {
	i    int
	page AlertListPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AlertListIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AlertListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AlertListIterator) Response() AlertList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AlertListIterator) Value() Alert {
	if !iter.page.NotDone() {
		return Alert{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (al AlertList) IsEmpty() bool {
	return al.Value == nil || len(*al.Value) == 0
}

// alertListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (al AlertList) alertListPreparer() (*http.Request, error) {
	if al.NextLink == nil || len(to.String(al.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(al.NextLink)))
}

// AlertListPage contains a page of Alert values.
type AlertListPage struct {
	fn func(AlertList) (AlertList, error)
	al AlertList
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AlertListPage) Next() error {
	next, err := page.fn(page.al)
	if err != nil {
		return err
	}
	page.al = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AlertListPage) NotDone() bool {
	return !page.al.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AlertListPage) Response() AlertList {
	return page.al
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AlertListPage) Values() []Alert {
	if page.al.IsEmpty() {
		return nil
	}
	return *page.al.Value
}

// AlertModel holds Alert data
type AlertModel struct {
	// ClosedTimestamp - Gets or sets the closed timestamp of the alert.
	ClosedTimestamp *string `json:"closedTimestamp,omitempty"`
	// CreatedTimestamp - Gets or sets the created timestamp of the alert.
	CreatedTimestamp *string `json:"createdTimestamp,omitempty"`
	// Description - Gets or sets the description of the alert.
	Description *[]map[string]*string `json:"description,omitempty"`
	// FaultID - Gets or sets the fault id of the alert.
	FaultID *string `json:"faultId,omitempty"`
	// AlertID - Gets or sets the id of the alert.
	AlertID *string `json:"alertId,omitempty"`
	// FaultTypeID - Gets or sets the fault type id of the alert.
	FaultTypeID *string `json:"faultTypeId,omitempty"`
	// LastUpdatedTimestamp - Gets or sets last updated timestamp of the alert.
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty"`
	// AlertProperties - Gets or sets properties of the alert.
	AlertProperties map[string]*string `json:"alertProperties"`
	// Remediation - Gets or sets the admin friendly remediation instructions for the alert.
	Remediation *[]map[string]*string `json:"remediation,omitempty"`
	// ResourceRegistrationID - Gets or sets the registration id of the atomic component the alert belongs to.  This is null if not associated with a resource.
	ResourceRegistrationID *string `json:"resourceRegistrationId,omitempty"`
	// ResourceProviderRegistrationID - Gets or sets the registration id of the service the alert belongs to.
	ResourceProviderRegistrationID *string `json:"resourceProviderRegistrationId,omitempty"`
	// Severity - Gets or sets the severity of the alert.
	Severity *string `json:"severity,omitempty"`
	// State - Gets or sets the state of the alert.
	State *string `json:"state,omitempty"`
	// Title - Gets or sets the ResourceId for the impacted item.
	Title *string `json:"title,omitempty"`
	// ImpactedResourceID - Gets or sets the ResourceId for the impacted item.
	ImpactedResourceID *string `json:"impactedResourceId,omitempty"`
	// ImpactedResourceDisplayName - Gets or sets the display name for the impacted item.
	ImpactedResourceDisplayName *string `json:"impactedResourceDisplayName,omitempty"`
	// ClosedByUserAlias - Gets or sets the user alias who closed the alert.
	ClosedByUserAlias *string `json:"closedByUserAlias,omitempty"`
}

// MarshalJSON is the custom marshaler for AlertModel.
func (am AlertModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if am.ClosedTimestamp != nil {
		objectMap["closedTimestamp"] = am.ClosedTimestamp
	}
	if am.CreatedTimestamp != nil {
		objectMap["createdTimestamp"] = am.CreatedTimestamp
	}
	if am.Description != nil {
		objectMap["description"] = am.Description
	}
	if am.FaultID != nil {
		objectMap["faultId"] = am.FaultID
	}
	if am.AlertID != nil {
		objectMap["alertId"] = am.AlertID
	}
	if am.FaultTypeID != nil {
		objectMap["faultTypeId"] = am.FaultTypeID
	}
	if am.LastUpdatedTimestamp != nil {
		objectMap["lastUpdatedTimestamp"] = am.LastUpdatedTimestamp
	}
	if am.AlertProperties != nil {
		objectMap["alertProperties"] = am.AlertProperties
	}
	if am.Remediation != nil {
		objectMap["remediation"] = am.Remediation
	}
	if am.ResourceRegistrationID != nil {
		objectMap["resourceRegistrationId"] = am.ResourceRegistrationID
	}
	if am.ResourceProviderRegistrationID != nil {
		objectMap["resourceProviderRegistrationId"] = am.ResourceProviderRegistrationID
	}
	if am.Severity != nil {
		objectMap["severity"] = am.Severity
	}
	if am.State != nil {
		objectMap["state"] = am.State
	}
	if am.Title != nil {
		objectMap["title"] = am.Title
	}
	if am.ImpactedResourceID != nil {
		objectMap["impactedResourceId"] = am.ImpactedResourceID
	}
	if am.ImpactedResourceDisplayName != nil {
		objectMap["impactedResourceDisplayName"] = am.ImpactedResourceDisplayName
	}
	if am.ClosedByUserAlias != nil {
		objectMap["closedByUserAlias"] = am.ClosedByUserAlias
	}
	return json.Marshal(objectMap)
}

// AlertSummary summary of the alerts.
type AlertSummary struct {
	// CriticalAlertCount - How many critical alerts this service has.
	CriticalAlertCount *int32 `json:"criticalAlertCount,omitempty"`
	// WarningAlertCount - How many warning alerts this service has.
	WarningAlertCount *int32 `json:"warningAlertCount,omitempty"`
}

// BaseHealth models the base properties for health resource.
type BaseHealth struct {
	// AlertSummary - Gets or sets the alert summary.
	AlertSummary *AlertSummary `json:"alertSummary,omitempty"`
	// HealthState - Gets or sets the health status.
	HealthState *string `json:"healthState,omitempty"`
	// Namespace - Gets or sets the name space.
	Namespace *string `json:"namespace,omitempty"`
	// RegistrationID - Gets or sets the registration id.
	RegistrationID *string `json:"registrationId,omitempty"`
	// RoutePrefix - Gets or sets the route prefix.
	RoutePrefix *string `json:"routePrefix,omitempty"`
}

// Metrics metrics for a source.
type Metrics struct {
	// Name - Name of the usage metric.
	Name *string `json:"name,omitempty"`
	// MaCounterName - Name of the usage metric.
	MaCounterName *string `json:"maCounterName,omitempty"`
	// ObservedTimestamp - Name of the usage metric.
	ObservedTimestamp *date.Time `json:"observedTimestamp,omitempty"`
	// SourceType - Name of the usage metric.
	SourceType *string `json:"sourceType,omitempty"`
	// SourceName - Source of the metric. Possible values include: 'PhysicalNode', 'VirtualMachine', 'ResourceProvider'
	SourceName MetricsSourceType `json:"sourceName,omitempty"`
	// Unit - Unit for the metric. Possible values include: 'One', 'Percentage', 'B', 'KB', 'MB', 'GB', 'TB'
	Unit MetricsUnit `json:"unit,omitempty"`
	// Value - Name of the usage metric.
	Value *float64 `json:"value,omitempty"`
}

// RegionHealth contains information related to a regions health.
type RegionHealth struct {
	autorest.Response `json:"-"`
	// RegionHealthModel - Contains information related to a regions health.
	*RegionHealthModel `json:"properties,omitempty"`
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for RegionHealth.
func (rh RegionHealth) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rh.RegionHealthModel != nil {
		objectMap["properties"] = rh.RegionHealthModel
	}
	if rh.ID != nil {
		objectMap["id"] = rh.ID
	}
	if rh.Name != nil {
		objectMap["name"] = rh.Name
	}
	if rh.Type != nil {
		objectMap["type"] = rh.Type
	}
	if rh.Location != nil {
		objectMap["location"] = rh.Location
	}
	if rh.Tags != nil {
		objectMap["tags"] = rh.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for RegionHealth struct.
func (rh *RegionHealth) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var regionHealthModel RegionHealthModel
				err = json.Unmarshal(*v, &regionHealthModel)
				if err != nil {
					return err
				}
				rh.RegionHealthModel = &regionHealthModel
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rh.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rh.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rh.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rh.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rh.Tags = tags
			}
		}
	}

	return nil
}

// RegionHealthList pageable list of region health items.
type RegionHealthList struct {
	autorest.Response `json:"-"`
	// Value - Array of region health items
	Value *[]RegionHealth `json:"value,omitempty"`
	// NextLink - URI to next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// RegionHealthListIterator provides access to a complete listing of RegionHealth values.
type RegionHealthListIterator struct {
	i    int
	page RegionHealthListPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RegionHealthListIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RegionHealthListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RegionHealthListIterator) Response() RegionHealthList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RegionHealthListIterator) Value() RegionHealth {
	if !iter.page.NotDone() {
		return RegionHealth{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (rhl RegionHealthList) IsEmpty() bool {
	return rhl.Value == nil || len(*rhl.Value) == 0
}

// regionHealthListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rhl RegionHealthList) regionHealthListPreparer() (*http.Request, error) {
	if rhl.NextLink == nil || len(to.String(rhl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rhl.NextLink)))
}

// RegionHealthListPage contains a page of RegionHealth values.
type RegionHealthListPage struct {
	fn  func(RegionHealthList) (RegionHealthList, error)
	rhl RegionHealthList
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RegionHealthListPage) Next() error {
	next, err := page.fn(page.rhl)
	if err != nil {
		return err
	}
	page.rhl = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RegionHealthListPage) NotDone() bool {
	return !page.rhl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RegionHealthListPage) Response() RegionHealthList {
	return page.rhl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RegionHealthListPage) Values() []RegionHealth {
	if page.rhl.IsEmpty() {
		return nil
	}
	return *page.rhl.Value
}

// RegionHealthModel contains information related to a regions health.
type RegionHealthModel struct {
	// AlertSummary - Summary of alerts.
	AlertSummary *AlertSummary `json:"alertSummary,omitempty"`
	// UsageMetrics - List of usage metrics for this region.
	UsageMetrics *[]UsageMetrics `json:"usageMetrics,omitempty"`
}

// Resource base Resource Object
type Resource struct {
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceHealth health information related to a resource.
type ResourceHealth struct {
	autorest.Response `json:"-"`
	// ResourceHealthModel - Health information related to a resource.
	*ResourceHealthModel `json:"properties,omitempty"`
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ResourceHealth.
func (rh ResourceHealth) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rh.ResourceHealthModel != nil {
		objectMap["properties"] = rh.ResourceHealthModel
	}
	if rh.ID != nil {
		objectMap["id"] = rh.ID
	}
	if rh.Name != nil {
		objectMap["name"] = rh.Name
	}
	if rh.Type != nil {
		objectMap["type"] = rh.Type
	}
	if rh.Location != nil {
		objectMap["location"] = rh.Location
	}
	if rh.Tags != nil {
		objectMap["tags"] = rh.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceHealth struct.
func (rh *ResourceHealth) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var resourceHealthModel ResourceHealthModel
				err = json.Unmarshal(*v, &resourceHealthModel)
				if err != nil {
					return err
				}
				rh.ResourceHealthModel = &resourceHealthModel
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rh.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rh.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rh.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rh.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rh.Tags = tags
			}
		}
	}

	return nil
}

// ResourceHealthList pageable list of resource healths.
type ResourceHealthList struct {
	autorest.Response `json:"-"`
	// Value - Array of of resource healths
	Value *[]ResourceHealth `json:"value,omitempty"`
	// NextLink - URI to next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceHealthListIterator provides access to a complete listing of ResourceHealth values.
type ResourceHealthListIterator struct {
	i    int
	page ResourceHealthListPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceHealthListIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceHealthListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceHealthListIterator) Response() ResourceHealthList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceHealthListIterator) Value() ResourceHealth {
	if !iter.page.NotDone() {
		return ResourceHealth{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (rhl ResourceHealthList) IsEmpty() bool {
	return rhl.Value == nil || len(*rhl.Value) == 0
}

// resourceHealthListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rhl ResourceHealthList) resourceHealthListPreparer() (*http.Request, error) {
	if rhl.NextLink == nil || len(to.String(rhl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rhl.NextLink)))
}

// ResourceHealthListPage contains a page of ResourceHealth values.
type ResourceHealthListPage struct {
	fn  func(ResourceHealthList) (ResourceHealthList, error)
	rhl ResourceHealthList
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceHealthListPage) Next() error {
	next, err := page.fn(page.rhl)
	if err != nil {
		return err
	}
	page.rhl = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceHealthListPage) NotDone() bool {
	return !page.rhl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceHealthListPage) Response() ResourceHealthList {
	return page.rhl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceHealthListPage) Values() []ResourceHealth {
	if page.rhl.IsEmpty() {
		return nil
	}
	return *page.rhl.Value
}

// ResourceHealthModel health information related to a resource.
type ResourceHealthModel struct {
	// ResourceLocation - Gets or sets the resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceName - Gets or sets the resource name.
	ResourceName *string `json:"resourceName,omitempty"`
	// ResourceDisplayName - Gets or sets the resource display name.
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// ResourceType - Gets or sets the resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceURI - Gets or sets the resource uri.
	ResourceURI *string `json:"resourceURI,omitempty"`
	// RpRegistrationID - Gets or sets the resource provider registration id.
	RpRegistrationID *string `json:"rpRegistrationId,omitempty"`
	// UsageMetrics - Gets or sets the usage metrics.
	UsageMetrics *[]UsageMetrics `json:"usageMetrics,omitempty"`
	// AlertSummary - Gets or sets the alert summary.
	AlertSummary *AlertSummary `json:"alertSummary,omitempty"`
	// HealthState - Gets or sets the health status.
	HealthState *string `json:"healthState,omitempty"`
	// Namespace - Gets or sets the name space.
	Namespace *string `json:"namespace,omitempty"`
	// RegistrationID - Gets or sets the registration id.
	RegistrationID *string `json:"registrationId,omitempty"`
	// RoutePrefix - Gets or sets the route prefix.
	RoutePrefix *string `json:"routePrefix,omitempty"`
}

// ServiceHealth holds information about a services health.
type ServiceHealth struct {
	autorest.Response `json:"-"`
	// ServiceHealthModel - Holds information about a services health.
	*ServiceHealthModel `json:"properties,omitempty"`
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ServiceHealth.
func (sh ServiceHealth) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sh.ServiceHealthModel != nil {
		objectMap["properties"] = sh.ServiceHealthModel
	}
	if sh.ID != nil {
		objectMap["id"] = sh.ID
	}
	if sh.Name != nil {
		objectMap["name"] = sh.Name
	}
	if sh.Type != nil {
		objectMap["type"] = sh.Type
	}
	if sh.Location != nil {
		objectMap["location"] = sh.Location
	}
	if sh.Tags != nil {
		objectMap["tags"] = sh.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ServiceHealth struct.
func (sh *ServiceHealth) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var serviceHealthModel ServiceHealthModel
				err = json.Unmarshal(*v, &serviceHealthModel)
				if err != nil {
					return err
				}
				sh.ServiceHealthModel = &serviceHealthModel
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				sh.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				sh.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				sh.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				sh.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				sh.Tags = tags
			}
		}
	}

	return nil
}

// ServiceHealthList pageable list of service health instances.
type ServiceHealthList struct {
	autorest.Response `json:"-"`
	// Value - Array of service health instances.
	Value *[]ServiceHealth `json:"value,omitempty"`
	// NextLink - URI of the next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// ServiceHealthListIterator provides access to a complete listing of ServiceHealth values.
type ServiceHealthListIterator struct {
	i    int
	page ServiceHealthListPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ServiceHealthListIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ServiceHealthListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ServiceHealthListIterator) Response() ServiceHealthList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ServiceHealthListIterator) Value() ServiceHealth {
	if !iter.page.NotDone() {
		return ServiceHealth{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (shl ServiceHealthList) IsEmpty() bool {
	return shl.Value == nil || len(*shl.Value) == 0
}

// serviceHealthListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (shl ServiceHealthList) serviceHealthListPreparer() (*http.Request, error) {
	if shl.NextLink == nil || len(to.String(shl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(shl.NextLink)))
}

// ServiceHealthListPage contains a page of ServiceHealth values.
type ServiceHealthListPage struct {
	fn  func(ServiceHealthList) (ServiceHealthList, error)
	shl ServiceHealthList
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ServiceHealthListPage) Next() error {
	next, err := page.fn(page.shl)
	if err != nil {
		return err
	}
	page.shl = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ServiceHealthListPage) NotDone() bool {
	return !page.shl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ServiceHealthListPage) Response() ServiceHealthList {
	return page.shl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ServiceHealthListPage) Values() []ServiceHealth {
	if page.shl.IsEmpty() {
		return nil
	}
	return *page.shl.Value
}

// ServiceHealthModel holds information about a services health.
type ServiceHealthModel struct {
	// DisplayName - Name of the alert.
	DisplayName *string `json:"displayName,omitempty"`
	// ServiceLocation - Location of the service.
	ServiceLocation *string `json:"serviceLocation,omitempty"`
	// InfraURI - The route prefix to the alert.
	InfraURI *string `json:"infraURI,omitempty"`
	// AlertSummary - Gets or sets the alert summary.
	AlertSummary *AlertSummary `json:"alertSummary,omitempty"`
	// HealthState - Gets or sets the health status.
	HealthState *string `json:"healthState,omitempty"`
	// Namespace - Gets or sets the name space.
	Namespace *string `json:"namespace,omitempty"`
	// RegistrationID - Gets or sets the registration id.
	RegistrationID *string `json:"registrationId,omitempty"`
	// RoutePrefix - Gets or sets the route prefix.
	RoutePrefix *string `json:"routePrefix,omitempty"`
}

// UsageMetrics metrics of resource usage.
type UsageMetrics struct {
	// Name - Name of the usage metric.
	Name *string `json:"name,omitempty"`
	// MetricsValue - List of usage metrics.
	MetricsValue *[]Metrics `json:"metricsValue,omitempty"`
}
