# {{classname}}

All URIs are relative to *https://cdn-api.co-vin.in/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmOTP**](UserAuthenticationAPIsApi.md#ConfirmOTP) | **Post** /v2/auth/public/confirmOTP | Confirm mobile OTP for authentication
[**GenerateOTP**](UserAuthenticationAPIsApi.md#GenerateOTP) | **Post** /v2/auth/public/generateOTP | Authenticate a beneficiary by Mobile/OTP

# **ConfirmOTP**
> InlineResponse2001 ConfirmOTP(ctx, optional)
Confirm mobile OTP for authentication

API to confirm the OTP for authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserAuthenticationAPIsApiConfirmOTPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserAuthenticationAPIsApiConfirmOTPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body1**](Body1.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateOTP**
> InlineResponse200 GenerateOTP(ctx, optional)
Authenticate a beneficiary by Mobile/OTP

Initiate beneficiary authentication using mobile and OTP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserAuthenticationAPIsApiGenerateOTPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserAuthenticationAPIsApiGenerateOTPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

