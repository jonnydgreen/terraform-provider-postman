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

// checks if the PANElementCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PANElementCreated{}

// PANElementCreated Information about the Private API Network element.
type PANElementCreated struct {
	// The date and time at which the element was added.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// The user who added the element.
	AddedBy *int32 `json:"addedBy,omitempty"`
	// The date and time at which the element was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user who created the element.
	CreatedBy *int32 `json:"createdBy,omitempty"`
	// The element's description.
	Description *string `json:"description,omitempty"`
	// The element's Postman URL.
	Href *string `json:"href,omitempty"`
	// The element's ID or UID.
	Id *string `json:"id,omitempty"`
	// The element's name.
	Name *string `json:"name,omitempty"`
	// The parent folder's ID.
	ParentFolderId *int32 `json:"parentFolderId,omitempty"`
	// The element's summary.
	Summary *string `json:"summary,omitempty"`
	// The element's type.
	Type *string `json:"type,omitempty"`
	// The date and time at which the element was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The user who last updated the element.
	UpdatedBy *int32 `json:"updatedBy,omitempty"`
}

// NewPANElementCreated instantiates a new PANElementCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPANElementCreated() *PANElementCreated {
	this := PANElementCreated{}
	return &this
}

// NewPANElementCreatedWithDefaults instantiates a new PANElementCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPANElementCreatedWithDefaults() *PANElementCreated {
	this := PANElementCreated{}
	return &this
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *PANElementCreated) GetAddedAt() time.Time {
	if o == nil || isNil(o.AddedAt) {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AddedAt) {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *PANElementCreated) HasAddedAt() bool {
	if o != nil && !isNil(o.AddedAt) {
		return true
	}

	return false
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *PANElementCreated) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *PANElementCreated) GetAddedBy() int32 {
	if o == nil || isNil(o.AddedBy) {
		var ret int32
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetAddedByOk() (*int32, bool) {
	if o == nil || isNil(o.AddedBy) {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *PANElementCreated) HasAddedBy() bool {
	if o != nil && !isNil(o.AddedBy) {
		return true
	}

	return false
}

// SetAddedBy gets a reference to the given int32 and assigns it to the AddedBy field.
func (o *PANElementCreated) SetAddedBy(v int32) {
	o.AddedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PANElementCreated) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PANElementCreated) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PANElementCreated) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PANElementCreated) GetCreatedBy() int32 {
	if o == nil || isNil(o.CreatedBy) {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetCreatedByOk() (*int32, bool) {
	if o == nil || isNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PANElementCreated) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *PANElementCreated) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PANElementCreated) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PANElementCreated) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PANElementCreated) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PANElementCreated) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PANElementCreated) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PANElementCreated) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PANElementCreated) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PANElementCreated) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PANElementCreated) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PANElementCreated) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PANElementCreated) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PANElementCreated) SetName(v string) {
	o.Name = &v
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise.
func (o *PANElementCreated) GetParentFolderId() int32 {
	if o == nil || isNil(o.ParentFolderId) {
		var ret int32
		return ret
	}
	return *o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetParentFolderIdOk() (*int32, bool) {
	if o == nil || isNil(o.ParentFolderId) {
		return nil, false
	}
	return o.ParentFolderId, true
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *PANElementCreated) HasParentFolderId() bool {
	if o != nil && !isNil(o.ParentFolderId) {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given int32 and assigns it to the ParentFolderId field.
func (o *PANElementCreated) SetParentFolderId(v int32) {
	o.ParentFolderId = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PANElementCreated) GetSummary() string {
	if o == nil || isNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetSummaryOk() (*string, bool) {
	if o == nil || isNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PANElementCreated) HasSummary() bool {
	if o != nil && !isNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PANElementCreated) SetSummary(v string) {
	o.Summary = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PANElementCreated) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PANElementCreated) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PANElementCreated) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PANElementCreated) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PANElementCreated) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PANElementCreated) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *PANElementCreated) GetUpdatedBy() int32 {
	if o == nil || isNil(o.UpdatedBy) {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PANElementCreated) GetUpdatedByOk() (*int32, bool) {
	if o == nil || isNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *PANElementCreated) HasUpdatedBy() bool {
	if o != nil && !isNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *PANElementCreated) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

func (o PANElementCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PANElementCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AddedAt) {
		toSerialize["addedAt"] = o.AddedAt
	}
	if !isNil(o.AddedBy) {
		toSerialize["addedBy"] = o.AddedBy
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ParentFolderId) {
		toSerialize["parentFolderId"] = o.ParentFolderId
	}
	if !isNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullablePANElementCreated struct {
	value *PANElementCreated
	isSet bool
}

func (v NullablePANElementCreated) Get() *PANElementCreated {
	return v.value
}

func (v *NullablePANElementCreated) Set(val *PANElementCreated) {
	v.value = val
	v.isSet = true
}

func (v NullablePANElementCreated) IsSet() bool {
	return v.isSet
}

func (v *NullablePANElementCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePANElementCreated(val *PANElementCreated) *NullablePANElementCreated {
	return &NullablePANElementCreated{value: val, isSet: true}
}

func (v NullablePANElementCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePANElementCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


