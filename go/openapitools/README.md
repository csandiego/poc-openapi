# Go API client for openapitools

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapitools"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BookApi* | [**Create**](docs/BookApi.md#create) | **Post** /books | 
*BookApi* | [**Delete**](docs/BookApi.md#delete) | **Delete** /books/{id} | 
*BookApi* | [**Get**](docs/BookApi.md#get) | **Get** /books/{id} | 
*BookApi* | [**List**](docs/BookApi.md#list) | **Get** /books | 
*BookApi* | [**Update**](docs/BookApi.md#update) | **Put** /books/{id} | 


## Documentation For Models

 - [Book](docs/Book.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author


