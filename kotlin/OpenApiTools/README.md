# com.github.csandiego.pocopenapitools.openapitools - Kotlin client library for Books API

## Requires

* Kotlin 1.3.41
* Gradle 4.9

## Build

First, create the gradle wrapper script:

```
gradle wrapper
```

Then, run:

```
./gradlew check assemble
```

This runs all tests and packages the library.

## Features/Implementation Notes

* Supports JSON inputs/outputs, File inputs, and Form inputs.
* Supports collection formats for query parameters: csv, tsv, ssv, pipes.
* Some Kotlin and Java types are fully qualified to avoid conflicts with types defined in OpenAPI definitions.
* Implementation of ApiClient is intended to reduce method counts, specifically to benefit Android targets.

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BookApi* | [**create**](docs/BookApi.md#create) | **POST** /books | 
*BookApi* | [**delete**](docs/BookApi.md#delete) | **DELETE** /books/{id} | 
*BookApi* | [**get**](docs/BookApi.md#get) | **GET** /books/{id} | 
*BookApi* | [**list**](docs/BookApi.md#list) | **GET** /books | 
*BookApi* | [**update**](docs/BookApi.md#update) | **PUT** /books/{id} | 


<a name="documentation-for-models"></a>
## Documentation for Models

 - [com.github.csandiego.pocopenapitools.openapitools.models.Book](docs/Book.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
