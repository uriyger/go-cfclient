package cfclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// V3ServicePlan represents model of service plan object for v3
type V3ServicePlan struct {

	//Name of the service plan
	Name string `json:"name"`

	//Unique identifier for the service plan
	GUID string `json:"guid"`

	//Denotes the visibility of the plan; can be public, admin, organization, space, see list of visibility types
	VisibilityType string `json:"visibility_type"`

	// Whether the service plan is available
	Available bool `json:"available"`

	// Whether the service plan is free of charge
	Free bool `json:"free"`

	// The cost of the service plan as obtained from the service broker catalog
	Costs []interface{} `json:"costs"`

	// Description of the service plan
	Description string `json:"description"`

	//Information about the version of this service plan
	MaintenanceInfo MaintenanceInfo `json:"maintenance_info"`

	//This object contains information obtained from the service broker catalog
	BrokerCatalog V3ServicePlanBrokerCatalog `json:"broker_catalog"`

	//Schema definitions for service instances and service bindings for the service plan
	Schemas map[string]interface{} `json:"schemas"`

	// Relationships possible keys:
	// "service_offering" to-one relationship. The service offering that this service plan relates to
	// "space" to-one relationship. The space of the service broker, if this service plan is from a space-scoped service broker
	Relationships map[string]V3ToOneRelationship `json:"relationships"`

	// Service plan metadata
	Metadata V3Metadata `json:"metadata"`

	// Links to related resources
	Links map[string]Link `json:"links,omitempty"`

	// The time with zone when the object was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The time with zone when the object was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// V3ServicePlanBrokerCatalog represents "broker_catalog" field from service plan object
type V3ServicePlanBrokerCatalog struct {
	// The identifier that the service broker provided for this service plan
	ID string `json:"id"`

	// Additional information provided by the service broker as specified by OSBAPI
	Metadata map[string]interface{} `json:"metadata"`

	// The maximum number of seconds that Cloud Foundry will wait for an asynchronous service broker operation
	MaximumPollingDuration int `json:"maximum_polling_duration"`

	// Service Plan properties
	Features V3ServicePlanBrokerCatalogFeatures `json:"features"`
}

// V3ServicePlanBrokerCatalogFeatures represents "features" field from broker_catalog object from service plan object
type V3ServicePlanBrokerCatalogFeatures struct {
	// Whether the service plan supports upgrade/downgrade for service plans; when the catalog does not specify a value,
	// it is inherited from the service offering
	PlanUpdateable bool `json:"plan_updateable"`

	// Specifies whether service instances of the service can be bound to applications
	Bindable bool `json:"bindable"`
}

// V3ServiceOffering represents model of service plan object for v3
type V3ServiceOffering struct {
	//Name of the service offering
	Name string `json:"name"`

	//Unique identifier for the service offering
	GUID string `json:"guid"`

	// Description of the service offering
	Description string `json:"description"`

	// Whether the service offering is available
	Available bool `json:"available"`

	// Descriptive tags for the service offering
	Tags []string `json:"tags"`

	// A list of permissions that the user would have to give the service, if they provision it;
	// the only permissions currently supported are syslog_drain, route_forwarding and volume_mount
	Requires []string `json:"requires"`

	// Whether service Instances of this service offering can be shared across organizations and spaces
	Shareable bool `json:"shareable"`

	// Url that points to a documentation page for the service offering, if provided by the service broker as part of the metadata field
	DocumentationURL string `json:"documentation_url"`

	// Relationships possible keys:
	// "service_offering" to-one relationship. The service offering that this service plan relates to
	Relationships map[string]V3ToOneRelationship `json:"relationships"`

	// This object contains information obtained from the service broker Catalog
	BrokerCatalog V3ServiceOfferingBrokerCatalog `json:"broker_catalog"`

	// Service offering metadata
	Metadata V3Metadata `json:"metadata"`

	// Links to related resources
	Links map[string]Link `json:"links,omitempty"`

	// The time with zone when the object was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The time with zone when the object was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// V3ServiceOfferingBrokerCatalog represents "broker_catalog" field from service offering object
type V3ServiceOfferingBrokerCatalog struct {
	// The identifier that the service broker provided for this service offering
	ID string `json:"id"`

	// Additional information provided by the service broker as specified by OSBAPI
	Metadata map[string]interface{} `json:"metadata"`

	// Service Offering properties
	Features V3ServiceOfferingBrokerCatalogFeatures `json:"features"`
}

// V3ServiceOfferingBrokerCatalogFeatures represents "features" field from broker_catalog object from service offering object
type V3ServiceOfferingBrokerCatalogFeatures struct {
	// Whether the service offering supports upgrade/downgrade for service plans by default; service plans can override this field
	PlanUpdateable bool `json:"plan_updateable"`

	// Specifies whether service instances of the service can be bound to applications
	Bindable bool `json:"bindable"`

	// Specifies whether the Fetching a service instance endpoint is supported for all service plans
	InstancesRetrievable bool `json:"instances_retrievable"`

	// Specifies whether the Fetching a service binding endpoint is supported for all service plans
	BindingsRetrievable bool `json:"bindings_retrievable"`

	// Specifies whether service instance updates relating only to context are propagated to the service broker
	AllowContextUpdates bool `json:"allow_context_updates"`
}

// V3PlansIncluded model for included object in service plans response
type V3PlansIncluded struct {
	ServiceOfferings []V3ServiceOffering `json:"service_offerings,omitempty"`
}

type listV3PlansResponse struct {
	Pagination Pagination      `json:"pagination,omitempty"`
	Resources  []V3ServicePlan `json:"resources,omitempty"`
	Included   V3PlansIncluded `json:"included,omitempty"`
}

// ListV3ServicePlansForSpace retrieves available service plans and offerings for specific space and broker names
func (c *Client) ListV3ServicePlansForSpace(spaceGUIDs, brokerNames string) ([]V3ServicePlan, []V3ServiceOffering, error) {
	query := url.Values{}
	query["available"] = []string{"true"}
	query["include"] = []string{"service_offering"}
	query["space_guids"] = []string{spaceGUIDs}

	if brokerNames != "" {
		query["service_broker_names"] = []string{brokerNames}
	}
	return c.ListV3ServicePlansByQuery(query)
}

// ListV3ServicePlansByQuery retrieves service plans based on query
func (c *Client) ListV3ServicePlansByQuery(query url.Values) ([]V3ServicePlan, []V3ServiceOffering, error) {
	{
		var (
			plans     []V3ServicePlan
			offerings []V3ServiceOffering
		)

		requestURL, err := url.Parse("/v3/service_plans")
		if err != nil {
			return nil, nil, err
		}
		requestURL.RawQuery = query.Encode()

		var data *listV3PlansResponse

		for {
			data, err = c.listV3ServicePlansResponse(requestURL)
			if err != nil {
				return nil, nil, err
			}

			plans = append(plans, data.Resources...)

			offerings = appendUniqOfferings(offerings, data.Included.ServiceOfferings...)

			requestURL, err = url.Parse(data.Pagination.Next.Href)
			if err != nil {
				return nil, nil, errors.Wrap(err, "Error parsing next page URL")
			}
			if requestURL.String() == "" {
				break
			}
		}

		return plans, offerings, nil
	}
}

// appendUniqOfferings adding offering to the list if there are no offerings with such GUID already added
func appendUniqOfferings(offerings []V3ServiceOffering, newOfferings ...V3ServiceOffering) []V3ServiceOffering {
	var exists bool
	for _, no := range newOfferings {

		exists = false
		for _, o := range offerings {
			if o.GUID == no.GUID {
				exists = true
				break
			}
		}

		if !exists {
			offerings = append(offerings, no)
		}
	}

	return offerings
}

func (c *Client) listV3ServicePlansResponse(requestURL *url.URL) (*listV3PlansResponse, error) {
	r := c.NewRequest("GET", fmt.Sprintf("%s?%s", requestURL.Path, requestURL.RawQuery))
	resp, err := c.DoRequest(r)
	if err != nil {
		return nil, errors.Wrap(err, "error requesting v3 service plans")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error listing v3 service plans, response code: %d", resp.StatusCode)
	}

	var data listV3PlansResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.Wrap(err, "error parsing JSON from list v3 service plans")
	}

	return &data, nil
}
