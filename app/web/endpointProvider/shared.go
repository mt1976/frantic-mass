package endpointProvider

import "encoding/json"

type JSONAPIResponse struct {
	Data Resource `json:"data"`
}

type Resource struct {
	Type          string                 `json:"type"`
	ID            string                 `json:"id"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	Links         map[string]string      `json:"links,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
}

// BuildJSONAPIResponse builds a JSON:API resource response
func BuildJSONAPIResponse(resourceType, id string, attributes map[string]interface{}) ([]byte, error) {
	resp := JSONAPIResponse{
		Data: Resource{
			Type:       resourceType,
			ID:         id,
			Attributes: attributes,
		},
	}
	return json.MarshalIndent(resp, "", "  ")
}
