/*
 * Co-WIN Public APIs
 *
 * Co-WIN Public APIs to find appointment availabilty and to download vaccination certificates. These APIs are available for use by all third party applications. <i>[Updated on 30 Apr 2021]</i>
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SessionResponseSchema struct {
	CenterId float64 `json:"center_id"`
	Name     string  `json:"name"`
	// Name in preferred language as specified in Accept-Language header parameter.
	NameL     string `json:"name_l,omitempty"`
	StateName string `json:"state_name"`
	// State name in preferred language as specified in Accept-Language header parameter.
	StateNameL   string `json:"state_name_l,omitempty"`
	DistrictName string `json:"district_name"`
	// District name in preferred language as specified in Accept-Language header parameter.
	DistrictNameL string `json:"district_name_l,omitempty"`
	BlockName     string `json:"block_name"`
	// Block name in preferred language as specified in Accept-Language header parameter.
	BlockNameL string  `json:"block_name_l,omitempty"`
	Pincode    string  `json:"pincode"`
	Lat        float32 `json:"lat,omitempty"`
	Long       float32 `json:"long,omitempty"`
	From       string  `json:"from"`
	To         string  `json:"to"`
	// Fee charged for vaccination
	FeeType           string  `json:"fee_type"`
	Fee               string  `json:"fee"`
	SessionId         string  `json:"session_id"`
	Date              string  `json:"date"`
	AvailableCapacity float64 `json:"available_capacity"`
	MinAgeLimit       float64 `json:"min_age_limit"`
	Vaccine           string  `json:"vaccine"`
	// Array of slot names
	Slots []string `json:"slots"`
}
