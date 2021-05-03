# {{classname}}

All URIs are relative to *https://cdn-api.co-vin.in/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalendarByDistrict**](AppointmentAvailabilityAPIsApi.md#CalendarByDistrict) | **Get** /v2/appointment/sessions/public/calendarByDistrict | Get vaccination sessions by district for 7 days
[**CalendarByPin**](AppointmentAvailabilityAPIsApi.md#CalendarByPin) | **Get** /v2/appointment/sessions/public/calendarByPin | Get vaccination sessions by PIN for 7 days
[**FindByDistrict**](AppointmentAvailabilityAPIsApi.md#FindByDistrict) | **Get** /v2/appointment/sessions/public/findByDistrict | Get vaccination sessions by district
[**FindByPin**](AppointmentAvailabilityAPIsApi.md#FindByPin) | **Get** /v2/appointment/sessions/public/findByPin | Get vaccination sessions by PIN

# **CalendarByDistrict**
> InlineResponse2005 CalendarByDistrict(ctx, districtId, date, optional)
Get vaccination sessions by district for 7 days

API to get planned vaccination sessions for 7 days from a specific date in a given district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **districtId** | **string**|  | 
  **date** | **string**|  | 
 **optional** | ***AppointmentAvailabilityAPIsApiCalendarByDistrictOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppointmentAvailabilityAPIsApiCalendarByDistrictOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **optional.String**| The locate code of the preferred language such as en_US. The text data will be returned in the preferred language along with default English text. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CalendarByPin**
> InlineResponse2005 CalendarByPin(ctx, pincode, date, optional)
Get vaccination sessions by PIN for 7 days

API to get planned vaccination sessions for 7 days from a specific date in a given pin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pincode** | **string**|  | 
  **date** | **string**|  | 
 **optional** | ***AppointmentAvailabilityAPIsApiCalendarByPinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppointmentAvailabilityAPIsApiCalendarByPinOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **optional.String**| The locate code of the preferred language such as en_US. The text data will be returned in the preferred language along with default English text. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindByDistrict**
> InlineResponse2004 FindByDistrict(ctx, districtId, date, optional)
Get vaccination sessions by district

API to get planned vaccination sessions on a specific date in a given district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **districtId** | **string**|  | 
  **date** | **string**|  | 
 **optional** | ***AppointmentAvailabilityAPIsApiFindByDistrictOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppointmentAvailabilityAPIsApiFindByDistrictOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **optional.String**| The locate code of the preferred language such as en_US. The text data will be returned in the preferred language along with default English text. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindByPin**
> InlineResponse2004 FindByPin(ctx, pincode, date, optional)
Get vaccination sessions by PIN

API to get planned vaccination sessions on a specific date in a given pin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pincode** | **string**|  | 
  **date** | **string**|  | 
 **optional** | ***AppointmentAvailabilityAPIsApiFindByPinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppointmentAvailabilityAPIsApiFindByPinOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **optional.String**| The locate code of the preferred language such as en_US. The text data will be returned in the preferred language along with default English text. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

