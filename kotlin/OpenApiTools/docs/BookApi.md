# BookApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](BookApi.md#create) | **POST** /books | 
[**delete**](BookApi.md#delete) | **DELETE** /books/{id} | 
[**get**](BookApi.md#get) | **GET** /books/{id} | 
[**list**](BookApi.md#list) | **GET** /books | 
[**update**](BookApi.md#update) | **PUT** /books/{id} | 


<a name="create"></a>
# **create**
> create(book)



### Example
```kotlin
// Import classes:
//import com.github.csandiego.pocopenapitools.openapitools.infrastructure.*
//import com.github.csandiego.pocopenapitools.openapitools.models.*

val apiInstance = BookApi()
val book : Book =  // Book | 
try {
    apiInstance.create(book)
} catch (e: ClientException) {
    println("4xx response calling BookApi#create")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling BookApi#create")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **book** | [**Book**](Book.md)|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

<a name="delete"></a>
# **delete**
> delete(id)



### Example
```kotlin
// Import classes:
//import com.github.csandiego.pocopenapitools.openapitools.infrastructure.*
//import com.github.csandiego.pocopenapitools.openapitools.models.*

val apiInstance = BookApi()
val id : kotlin.Int = 56 // kotlin.Int | 
try {
    apiInstance.delete(id)
} catch (e: ClientException) {
    println("4xx response calling BookApi#delete")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling BookApi#delete")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **kotlin.Int**|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="get"></a>
# **get**
> Book get(id)



### Example
```kotlin
// Import classes:
//import com.github.csandiego.pocopenapitools.openapitools.infrastructure.*
//import com.github.csandiego.pocopenapitools.openapitools.models.*

val apiInstance = BookApi()
val id : kotlin.Int = 56 // kotlin.Int | 
try {
    val result : Book = apiInstance.get(id)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling BookApi#get")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling BookApi#get")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **kotlin.Int**|  |

### Return type

[**Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="list"></a>
# **list**
> kotlin.Array&lt;Book&gt; list()



### Example
```kotlin
// Import classes:
//import com.github.csandiego.pocopenapitools.openapitools.infrastructure.*
//import com.github.csandiego.pocopenapitools.openapitools.models.*

val apiInstance = BookApi()
try {
    val result : kotlin.Array<Book> = apiInstance.list()
    println(result)
} catch (e: ClientException) {
    println("4xx response calling BookApi#list")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling BookApi#list")
    e.printStackTrace()
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**kotlin.Array&lt;Book&gt;**](Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="update"></a>
# **update**
> update(id, book)



### Example
```kotlin
// Import classes:
//import com.github.csandiego.pocopenapitools.openapitools.infrastructure.*
//import com.github.csandiego.pocopenapitools.openapitools.models.*

val apiInstance = BookApi()
val id : kotlin.Int = 56 // kotlin.Int | 
val book : Book =  // Book | 
try {
    apiInstance.update(id, book)
} catch (e: ClientException) {
    println("4xx response calling BookApi#update")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling BookApi#update")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **kotlin.Int**|  |
 **book** | [**Book**](Book.md)|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

