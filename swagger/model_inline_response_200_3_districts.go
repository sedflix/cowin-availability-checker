/*
 * Co-WIN Public APIs
 *
 * Co-WIN Public APIs to find appointment availabilty and to download vaccination certificates. These APIs are available for use by all third party applications. <i>[Updated on 30 Apr 2021]</i>
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse2003Districts struct {
	StateId      float64 `json:"state_id,omitempty"`
	DistrictId   float64 `json:"district_id"`
	DistrictName string  `json:"district_name"`
	// District name in preferred language as specified in Accept-Language header parameter.
	DistrictNameL string `json:"district_name_l,omitempty"`
}