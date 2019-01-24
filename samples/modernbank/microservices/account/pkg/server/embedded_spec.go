// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Account Microservice, responsible for managing accounts and their balances in Modern Bank.",
    "title": "Account",
    "version": "1.0.0"
  },
  "host": "accounts",
  "basePath": "/v1",
  "paths": {
    "/users/{owner}/accounts": {
      "get": {
        "description": "Lists all accounts for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Lists all accounts for a given customer",
        "operationId": "listAccounts",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the accounts",
            "name": "owner",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Account"
              }
            }
          },
          "404": {
            "description": "Owner not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "description": "Creates a new account for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Create a new account for a customer",
        "operationId": "createAccount",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the account",
            "name": "owner",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Account"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/users/{owner}/accounts/{number}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Get account details",
        "operationId": "getAccountByNumber",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the account",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "The number of the account that is to be fetched.",
            "name": "number",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/Account"
            }
          },
          "404": {
            "description": "Account not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "delete": {
        "description": "Delete account by account number.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Delete account by account number",
        "operationId": "deleteAccount",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the account",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "The number of the account that is to be deleted.",
            "name": "number",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Aaaaand it's gone!"
          },
          "404": {
            "description": "Account not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "Account": {
      "type": "object",
      "properties": {
        "balance": {
          "type": "number",
          "format": "currency"
        },
        "number": {
          "type": "integer",
          "format": "int64"
        },
        "owner": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about Accounts",
      "name": "account"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Account Microservice, responsible for managing accounts and their balances in Modern Bank.",
    "title": "Account",
    "version": "1.0.0"
  },
  "host": "accounts",
  "basePath": "/v1",
  "paths": {
    "/users/{owner}/accounts": {
      "get": {
        "description": "Lists all accounts for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Lists all accounts for a given customer",
        "operationId": "listAccounts",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the accounts",
            "name": "owner",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Account"
              }
            }
          },
          "404": {
            "description": "Owner not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "description": "Creates a new account for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Create a new account for a customer",
        "operationId": "createAccount",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the account",
            "name": "owner",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Account"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/users/{owner}/accounts/{number}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Get account details",
        "operationId": "getAccountByNumber",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the account",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "The number of the account that is to be fetched.",
            "name": "number",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/Account"
            }
          },
          "404": {
            "description": "Account not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "delete": {
        "description": "Delete account by account number.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Delete account by account number",
        "operationId": "deleteAccount",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of the account",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "The number of the account that is to be deleted.",
            "name": "number",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Aaaaand it's gone!"
          },
          "404": {
            "description": "Account not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "Account": {
      "type": "object",
      "properties": {
        "balance": {
          "type": "number",
          "format": "currency"
        },
        "number": {
          "type": "integer",
          "format": "int64"
        },
        "owner": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about Accounts",
      "name": "account"
    }
  ]
}`))
}
