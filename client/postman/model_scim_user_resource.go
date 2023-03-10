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
)

// checks if the ScimUserResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimUserResource{}

// ScimUserResource The SCIM user resource object.
type ScimUserResource struct {
	// If true, the team member is active.
	Active *bool `json:"active,omitempty"`
	// The team member's external ID.
	ExternalId *string `json:"externalId,omitempty"`
	// The team member's SCIM ID.
	Id *string `json:"id,omitempty"`
	Meta *ScimUserResourceMeta `json:"meta,omitempty"`
	Name *ScimUserResourceName `json:"name,omitempty"`
	// A list of schema resource URIs.
	Schemas []string `json:"schemas,omitempty"`
	// The team member's SCIM username.
	UserName *string `json:"userName,omitempty"`
}

// NewScimUserResource instantiates a new ScimUserResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimUserResource() *ScimUserResource {
	this := ScimUserResource{}
	return &this
}

// NewScimUserResourceWithDefaults instantiates a new ScimUserResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimUserResourceWithDefaults() *ScimUserResource {
	this := ScimUserResource{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ScimUserResource) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ScimUserResource) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ScimUserResource) SetActive(v bool) {
	o.Active = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ScimUserResource) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ScimUserResource) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ScimUserResource) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScimUserResource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScimUserResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScimUserResource) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimUserResource) GetMeta() ScimUserResourceMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ScimUserResourceMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetMetaOk() (*ScimUserResourceMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimUserResource) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ScimUserResourceMeta and assigns it to the Meta field.
func (o *ScimUserResource) SetMeta(v ScimUserResourceMeta) {
	o.Meta = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimUserResource) GetName() ScimUserResourceName {
	if o == nil || IsNil(o.Name) {
		var ret ScimUserResourceName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetNameOk() (*ScimUserResourceName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScimUserResource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given ScimUserResourceName and assigns it to the Name field.
func (o *ScimUserResource) SetName(v ScimUserResourceName) {
	o.Name = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimUserResource) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimUserResource) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimUserResource) SetSchemas(v []string) {
	o.Schemas = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ScimUserResource) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserResource) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ScimUserResource) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ScimUserResource) SetUserName(v string) {
	o.UserName = &v
}

func (o ScimUserResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimUserResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return toSerialize, nil
}

type NullableScimUserResource struct {
	value *ScimUserResource
	isSet bool
}

func (v NullableScimUserResource) Get() *ScimUserResource {
	return v.value
}

func (v *NullableScimUserResource) Set(val *ScimUserResource) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUserResource) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUserResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUserResource(val *ScimUserResource) *NullableScimUserResource {
	return &NullableScimUserResource{value: val, isSet: true}
}

func (v NullableScimUserResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUserResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


