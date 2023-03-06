/*
Postman API

The Postman API lets you to programmatically access data stored in Postman account with ease.  ## Getting started  The easiest way to get started with the Postman API is to [fork this collection](https://learning.postman.com/docs/collaborating-in-postman/version-control/#creating-a-fork) to your own workspace. You can then use Postman to send requests.  ## Overview  1. You must use a valid API Key to send requests to the API endpoints. You can get your API key from Postman's [integrations dashboard](https://go.postman.co/settings/me/api-keys). 1. The API has access rate limits. 1. The API only responds to HTTPS-secured communications. Any requests sent via HTTP return an HTTP `301` redirect to the corresponding HTTPS resources. 1. The API returns requests responses in [JSON format](https://en.wikipedia.org/wiki/JSON). When an API request returns an error, it is sent in the JSON response as an `\"error\": {}` key. 1. The request method (verb) determines the nature of action you intend to perform. A request made using the `GET` method implies that you want to fetch something from Postman, and `POST` implies you want to save something new to Postman. 1. API calls respond with the appropriate [HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) for all requests. Within the Postman client, when a response is received, the status code is highlighted and is accompanied by a help text that indicates the possible meaning of the response code. A `200 OK` indicates success, while an HTTP `4XX` or HTTP `5XX` response code indicates an error from the requesting client or our API servers, respectively. 1. Individual resources in your Postman account are accessible using its unique ID (`uid`) value. The `uid` is a simple concatenation of the resource owner's user ID and the resource's ID. For example, a collection's `uid` is `{{owner_id}}-{{collection_id}}` value.  ## ID and UID  All items in Postman, such as collections, mocks, workspaces, and APIs, have ID and UIDs:  - An ID is the unique ID assigned to a Postman item. For example, `ec29121c-5203-409f-9e84-e83ffc10f226`. - The UID is the **full** ID of a Postman item. This value is the item's unique ID concatenated with the user ID. For example, in the `12345678-ec29121c-5203-409f-9e84-e83ffc10f226` UID:   - `12345678` is the user's ID.   - `ec29121c-5203-409f-9e84-e83ffc10f226` is the item's ID.  ## Authentication  An API key is required to be sent as part of every request to the Postman API, in the form of an `X-Api-Key` request header. To get a Postman API key, you can generate one in the [**API keys**](https://postman.postman.co/settings/me/api-keys) section in your Postman account settings.  An API key tells the API server that the received request from you. Everything that you have access to in Postman is accessible with your API key.  For ease of use in Postman, you can store your API key as the `postman_api_key` [environment variable](https://www.getpostman.com/docs/environments). The Postman API [collection](https://www.getpostman.com/docs/collections) will automatically use it to make API calls.  ### API Key related error response  If an API key is missing, malformed, or invalid, you will receive an HTTP `401 Unauthorized` response code and the following JSON response:  ```json { \"error\": {     \"name\": \"AuthenticationError\",     \"message\": \"Invalid API Key. Every request requires a valid API Key to be sent.\"   } } ```  ### Using the API key as a query parameter  Each request that accepts API key as `X-Api-Key` request header also accepts the key when it is sent as the `apikey` query parameter.  An API key sent as part of the header has a higher priority when you send the key as both a request header and a query parameter.  ## Rate Limits  API access rate limits apply at a per-API key basis in unit time. Access to the API using an API key is limited to **60 requests per minute**. In addition, every API response is accompanied by the following set of headers to identify the status of your use:  | Header | Description | | ------ | ----------- | | `X-RateLimit-Limit` | The maximum number of requests that the consumer is permitted to make per minute. | | `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window. | | `X-RateLimit-Reset` | The time at which the current rate limit window resets in UTC epoch seconds. |  Once you reach the rate limit you will receive a response similar to the following HTTP `429 Too Many Requests` response:  ```json {   \"error\": {     \"name\": \"rateLimitError\",     \"message\": \"Rate Limit exceeded. Please retry at 1465452702843\"   } } ```  In the event you receive an HTTP `503` response from our servers, it indicates that we have had an unexpected spike in API access traffic. This is usually operational within the next five minutes. If the outage persists or you receive any other form of an HTTP `5XX` error, [contact support](https://support.postman.com/hc/en-us/requests/new/).  ## Support  For help regarding accessing the Postman API, you can:  - Visit [Postman Support](https://support.postman.com/hc/en-us) or our [Community and Support](https://www.postman.com/community/) sites. - Reach out to the [Postman community](https://community.postman.com/). - Submit a help request to [Postman support](https://support.postman.com/hc/en-us/requests/new/).  ## Policies  - [Postman Terms of Service](http://www.postman.com/legal/terms/) - [Postman Privacy Policy](https://www.postman.com/legal/privacy-policy/) 

API version: 1.0
Contact: help@postman.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postman

import (
	"encoding/json"
	"time"
)

// checks if the SingleWorkspace200ResponseWorkspace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleWorkspace200ResponseWorkspace{}

// SingleWorkspace200ResponseWorkspace Information about the workspace.
type SingleWorkspace200ResponseWorkspace struct {
	// The workspace's APIs.
	Apis []SingleWorkspace200ResponseWorkspaceApisInner `json:"apis,omitempty"`
	// The workspace's collections.
	Collections []SingleWorkspace200ResponseWorkspaceCollectionsInner `json:"collections,omitempty"`
	// The date and time at which the workspace was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user ID of the user who created the workspace.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The workspace's description.
	Description *string `json:"description,omitempty"`
	// The workspace's environments.
	Environments []SingleWorkspace200ResponseWorkspaceEnvironmentsInner `json:"environments,omitempty"`
	// The workspace's ID.
	Id *string `json:"id,omitempty"`
	// The workspace's mock servers.
	Mocks []SingleWorkspace200ResponseWorkspaceMocksInner `json:"mocks,omitempty"`
	// The workspace's monitors.
	Monitors []SingleWorkspace200ResponseWorkspaceMonitorsInner `json:"monitors,omitempty"`
	// The workspace's name.
	Name *string `json:"name,omitempty"`
	// The type of workspace.
	Type *string `json:"type,omitempty"`
	// The date and time at which the workspace was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The user ID of the user who last updated the workspace.
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// The workspace's visibility. [Visibility](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/managing-workspaces/#changing-workspace-visibility) determines who can access the workspace:  - `only-me` — Applies to the **My Workspace** workspace. - `personal` — Only you can access the workspace. - `team` — All team members can access the workspace. - `private-team` — Only invited team members can access the workspace. - `public` — Everyone can access the workspace. 
	Visibility *string `json:"visibility,omitempty"`
}

// NewSingleWorkspace200ResponseWorkspace instantiates a new SingleWorkspace200ResponseWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleWorkspace200ResponseWorkspace() *SingleWorkspace200ResponseWorkspace {
	this := SingleWorkspace200ResponseWorkspace{}
	return &this
}

// NewSingleWorkspace200ResponseWorkspaceWithDefaults instantiates a new SingleWorkspace200ResponseWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleWorkspace200ResponseWorkspaceWithDefaults() *SingleWorkspace200ResponseWorkspace {
	this := SingleWorkspace200ResponseWorkspace{}
	return &this
}

// GetApis returns the Apis field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetApis() []SingleWorkspace200ResponseWorkspaceApisInner {
	if o == nil || IsNil(o.Apis) {
		var ret []SingleWorkspace200ResponseWorkspaceApisInner
		return ret
	}
	return o.Apis
}

// GetApisOk returns a tuple with the Apis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetApisOk() ([]SingleWorkspace200ResponseWorkspaceApisInner, bool) {
	if o == nil || IsNil(o.Apis) {
		return nil, false
	}
	return o.Apis, true
}

// HasApis returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasApis() bool {
	if o != nil && !IsNil(o.Apis) {
		return true
	}

	return false
}

// SetApis gets a reference to the given []SingleWorkspace200ResponseWorkspaceApisInner and assigns it to the Apis field.
func (o *SingleWorkspace200ResponseWorkspace) SetApis(v []SingleWorkspace200ResponseWorkspaceApisInner) {
	o.Apis = v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetCollections() []SingleWorkspace200ResponseWorkspaceCollectionsInner {
	if o == nil || IsNil(o.Collections) {
		var ret []SingleWorkspace200ResponseWorkspaceCollectionsInner
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetCollectionsOk() ([]SingleWorkspace200ResponseWorkspaceCollectionsInner, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []SingleWorkspace200ResponseWorkspaceCollectionsInner and assigns it to the Collections field.
func (o *SingleWorkspace200ResponseWorkspace) SetCollections(v []SingleWorkspace200ResponseWorkspaceCollectionsInner) {
	o.Collections = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SingleWorkspace200ResponseWorkspace) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SingleWorkspace200ResponseWorkspace) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SingleWorkspace200ResponseWorkspace) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetEnvironments() []SingleWorkspace200ResponseWorkspaceEnvironmentsInner {
	if o == nil || IsNil(o.Environments) {
		var ret []SingleWorkspace200ResponseWorkspaceEnvironmentsInner
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetEnvironmentsOk() ([]SingleWorkspace200ResponseWorkspaceEnvironmentsInner, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []SingleWorkspace200ResponseWorkspaceEnvironmentsInner and assigns it to the Environments field.
func (o *SingleWorkspace200ResponseWorkspace) SetEnvironments(v []SingleWorkspace200ResponseWorkspaceEnvironmentsInner) {
	o.Environments = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SingleWorkspace200ResponseWorkspace) SetId(v string) {
	o.Id = &v
}

// GetMocks returns the Mocks field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetMocks() []SingleWorkspace200ResponseWorkspaceMocksInner {
	if o == nil || IsNil(o.Mocks) {
		var ret []SingleWorkspace200ResponseWorkspaceMocksInner
		return ret
	}
	return o.Mocks
}

// GetMocksOk returns a tuple with the Mocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetMocksOk() ([]SingleWorkspace200ResponseWorkspaceMocksInner, bool) {
	if o == nil || IsNil(o.Mocks) {
		return nil, false
	}
	return o.Mocks, true
}

// HasMocks returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasMocks() bool {
	if o != nil && !IsNil(o.Mocks) {
		return true
	}

	return false
}

// SetMocks gets a reference to the given []SingleWorkspace200ResponseWorkspaceMocksInner and assigns it to the Mocks field.
func (o *SingleWorkspace200ResponseWorkspace) SetMocks(v []SingleWorkspace200ResponseWorkspaceMocksInner) {
	o.Mocks = v
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetMonitors() []SingleWorkspace200ResponseWorkspaceMonitorsInner {
	if o == nil || IsNil(o.Monitors) {
		var ret []SingleWorkspace200ResponseWorkspaceMonitorsInner
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetMonitorsOk() ([]SingleWorkspace200ResponseWorkspaceMonitorsInner, bool) {
	if o == nil || IsNil(o.Monitors) {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasMonitors() bool {
	if o != nil && !IsNil(o.Monitors) {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []SingleWorkspace200ResponseWorkspaceMonitorsInner and assigns it to the Monitors field.
func (o *SingleWorkspace200ResponseWorkspace) SetMonitors(v []SingleWorkspace200ResponseWorkspaceMonitorsInner) {
	o.Monitors = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SingleWorkspace200ResponseWorkspace) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SingleWorkspace200ResponseWorkspace) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SingleWorkspace200ResponseWorkspace) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SingleWorkspace200ResponseWorkspace) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *SingleWorkspace200ResponseWorkspace) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleWorkspace200ResponseWorkspace) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *SingleWorkspace200ResponseWorkspace) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *SingleWorkspace200ResponseWorkspace) SetVisibility(v string) {
	o.Visibility = &v
}

func (o SingleWorkspace200ResponseWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleWorkspace200ResponseWorkspace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apis) {
		toSerialize["apis"] = o.Apis
	}
	if !IsNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mocks) {
		toSerialize["mocks"] = o.Mocks
	}
	if !IsNil(o.Monitors) {
		toSerialize["monitors"] = o.Monitors
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableSingleWorkspace200ResponseWorkspace struct {
	value *SingleWorkspace200ResponseWorkspace
	isSet bool
}

func (v NullableSingleWorkspace200ResponseWorkspace) Get() *SingleWorkspace200ResponseWorkspace {
	return v.value
}

func (v *NullableSingleWorkspace200ResponseWorkspace) Set(val *SingleWorkspace200ResponseWorkspace) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleWorkspace200ResponseWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleWorkspace200ResponseWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleWorkspace200ResponseWorkspace(val *SingleWorkspace200ResponseWorkspace) *NullableSingleWorkspace200ResponseWorkspace {
	return &NullableSingleWorkspace200ResponseWorkspace{value: val, isSet: true}
}

func (v NullableSingleWorkspace200ResponseWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleWorkspace200ResponseWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


