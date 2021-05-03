# {{classname}}

All URIs are relative to *https://cdn-api.co-vin.in/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Districts**](MetadataAPIsApi.md#Districts) | **Get** /v2/admin/location/districts/{state_id} | Get list of districts
[**States**](MetadataAPIsApi.md#States) | **Get** /v2/admin/location/states | Get states

# **Districts**
> InlineResponse2003 Districts(ctx, stateId, optional)
Get list of districts

API to get all the districts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stateId** | **string**|  | 
 **optional** | ***MetadataAPIsApiDistrictsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetadataAPIsApiDistrictsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| The locate code of the preferred language such as en_US. The text data will be returned in the preferred language along with default English text. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, examples

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **States**
> InlineResponse2002 States(ctx, optional)
Get states

API to get all the states in India.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetadataAPIsApiStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetadataAPIsApiStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **optional.String**| The locate code of the preferred language such as en_US. The text data will be returned in the preferred language along with default English text. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, examples

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

