# \AccountsApi

All URIs are relative to *https://api.pro.coinbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAccountIdGet**](AccountsApi.md#AccountsAccountIdGet) | **Get** /accounts/{account-id} | Get one account
[**AccountsAccountIdHoldsGet**](AccountsApi.md#AccountsAccountIdHoldsGet) | **Get** /accounts/{account-id}/holds | Get ledger items
[**AccountsAccountIdLedgerGet**](AccountsApi.md#AccountsAccountIdLedgerGet) | **Get** /accounts/{account-id}/ledger | Get ledger items
[**AccountsGet**](AccountsApi.md#AccountsGet) | **Get** /accounts | List accounts


# **AccountsAccountIdGet**
> Account AccountsAccountIdGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, accountId)
Get one account

Information for a single account. Use this endpoint when you know the account_id. API key must belong to the same profile as the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **accountId** | **string**|  | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsAccountIdHoldsGet**
> []Hold AccountsAccountIdHoldsGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, accountId)
Get ledger items

List holds of an account that belong to the same profile as the API key. Holds are placed on an account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If an order is canceled, any remaining hold is removed. For a withdraw, once it is completed, the hold is removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **accountId** | **string**|  | 

### Return type

[**[]Hold**](Hold.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsAccountIdLedgerGet**
> []LedgerItem AccountsAccountIdLedgerGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, accountId)
Get ledger items

List account activity of the API key's profile. Account activity either increases or decreases your account balance. Items are paginated and sorted latest first. See the Pagination section for retrieving additional entries after the first page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **accountId** | **string**|  | 

### Return type

[**[]LedgerItem**](LedgerItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsGet**
> []Account AccountsGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE)
List accounts

Get a list of trading accounts from the profile of the API key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 

### Return type

[**[]Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

