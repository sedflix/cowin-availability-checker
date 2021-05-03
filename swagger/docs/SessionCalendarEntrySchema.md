# SessionCalendarEntrySchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CenterId** | **float64** |  | [default to null]
**Name** | **string** |  | [default to null]
**NameL** | **string** | Name in preferred language as specified in Accept-Language header parameter. | [optional] [default to null]
**StateName** | **string** |  | [default to null]
**StateNameL** | **string** | State name in preferred language as specified in Accept-Language header parameter. | [optional] [default to null]
**DistrictName** | **string** |  | [default to null]
**DistrictNameL** | **string** | District name in preferred language as specified in Accept-Language header parameter. | [optional] [default to null]
**BlockName** | **string** |  | [default to null]
**BlockNameL** | **string** | Block name in preferred language as specified in Accept-Language header parameter. | [optional] [default to null]
**Pincode** | **string** |  | [default to null]
**Lat** | **float32** |  | [optional] [default to null]
**Long** | **float32** |  | [optional] [default to null]
**From** | **string** |  | [default to null]
**To** | **string** |  | [default to null]
**FeeType** | **string** | Fee charged for vaccination | [default to null]
**VaccineFees** | [***[]VaccineFeeSchema**](array.md) |  | [optional] [default to null]
**Sessions** | [**[]SessionCalendarEntrySchemaSessions**](SessionCalendarEntrySchema_sessions.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

