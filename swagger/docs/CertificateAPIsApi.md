# {{classname}}

All URIs are relative to *https://cdn-api.co-vin.in/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Download**](CertificateAPIsApi.md#Download) | **Get** /v2/registration/certificate/public/download | Download vaccination certificate in PDF format by beneficiary reference id

# **Download**
> *os.File Download(ctx, beneficiaryReferenceId)
Download vaccination certificate in PDF format by beneficiary reference id

API to download vaccination certificate in PDF format by beneficiary reference id. This API requires a <i>Bearer</i> token acquired with user mobile OTP validation as defined in User Authentication APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **beneficiaryReferenceId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

