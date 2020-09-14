# \ProductsApi

All URIs are relative to *https://api.pro.coinbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsGet**](ProductsApi.md#ProductsGet) | **Get** /products | Get a list of available currency pairs for trading.
[**ProductsProductIdBookGet**](ProductsApi.md#ProductsProductIdBookGet) | **Get** /products/{product-id}/book | Get a list of open orders for a product. The amount of detail shown can be customized with the level parameter.
[**ProductsProductIdCandlesGet**](ProductsApi.md#ProductsProductIdCandlesGet) | **Get** /products/{product-id}/candles | Get Historic Rates
[**ProductsProductIdGet**](ProductsApi.md#ProductsProductIdGet) | **Get** /products/{product-id} | Get market data for a specific currency pair.
[**ProductsProductIdStatsGet**](ProductsApi.md#ProductsProductIdStatsGet) | **Get** /products/{product-id}/stats | Get 24hr Stats
[**ProductsProductIdTickerGet**](ProductsApi.md#ProductsProductIdTickerGet) | **Get** /products/{product-id}/ticker | Get Product Ticker
[**ProductsProductIdTradesGet**](ProductsApi.md#ProductsProductIdTradesGet) | **Get** /products/{product-id}/trades | Get Trades


# **ProductsGet**
> []Product ProductsGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE)
Get a list of available currency pairs for trading.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 

### Return type

[**[]Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdBookGet**
> Book ProductsProductIdBookGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, productId, optional)
Get a list of open orders for a product. The amount of detail shown can be customized with the level parameter.

By default, only the inside (i.e. best) bid and ask are returned. This is equivalent to a book depth of 1 level. If you would like to see a larger order book, specify the level query parameter. If a level is not aggregated, then all of the orders at each price will be returned. Aggregated levels return only one size for each active price (as if there was only a single order for that size at the level). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **productId** | **string**|  | 
 **optional** | ***ProductsApiProductsProductIdBookGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProductsProductIdBookGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **level** | **optional.String**| 1: Only the best bid and ask 2: Top 50 bids and asks (aggregated) 3: Full order book (non aggregated)  | [default to 1]

### Return type

[**Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdCandlesGet**
> [][]float32 ProductsProductIdCandlesGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, productId, granularity, optional)
Get Historic Rates

Historical rate data may be incomplete. No data is published for intervals where there are no ticks. Historical rates should not be polled frequently. If you need real-time information, use the trade and book endpoints along with the websocket feed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **productId** | **string**|  | 
  **granularity** | **string**| The granularity field must be one of the following values: {60, 300, 900, 3600, 21600, 86400}. Otherwise, your request will be rejected. These values correspond to timeslices representing one minute, five minutes, fifteen minutes, one hour, six hours, and one day, respectively.  | 
 **optional** | ***ProductsApiProductsProductIdCandlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProductsProductIdCandlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **start** | **optional.String**| Start time in ISO 8601 | 
 **end** | **optional.String**| End time in ISO 8601 | 

### Return type

[**[][]float32**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdGet**
> Product ProductsProductIdGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, productId)
Get market data for a specific currency pair.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **productId** | **string**|  | 

### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdStatsGet**
> Stats ProductsProductIdStatsGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, productId)
Get 24hr Stats

Get 24 hr stats for the product. volume is in base currency units. open, high, low are in quote currency units.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **productId** | **string**|  | 

### Return type

[**Stats**](Stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdTickerGet**
> Ticker ProductsProductIdTickerGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, productId)
Get Product Ticker

Snapshot information about the last trade (tick), best bid/ask and 24h volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **productId** | **string**|  | 

### Return type

[**Ticker**](Ticker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdTradesGet**
> []Trade ProductsProductIdTradesGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, productId)
Get Trades

Snapshot information about the last trade (tick), best bid/ask and 24h volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **productId** | **string**|  | 

### Return type

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

