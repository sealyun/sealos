// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://cloud.sealos.io",
        "contact": {
            "email": "bxy4543@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/v1alpha1/costs": {
            "post": {
                "description": "Get user costs within a specified time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Costs"
                ],
                "summary": "Get user costs",
                "parameters": [
                    {
                        "description": "User costs amount request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helper.UserCostsAmountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully retrieved user costs",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "failed to parse user hour costs amount request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "authenticate error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "failed to get user costs",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/account/v1alpha1/costs/consumption": {
            "post": {
                "description": "Get user consumption amount within a specified time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ConsumptionAmount"
                ],
                "summary": "Get user consumption amount",
                "parameters": [
                    {
                        "description": "User consumption amount request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helper.UserCostsAmountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully retrieved user consumption amount",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "failed to parse user consumption amount request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "authenticate error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "failed to get user consumption amount",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/account/v1alpha1/costs/properties": {
            "post": {
                "description": "Get user properties used amount within a specified time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PropertiesUsedAmount"
                ],
                "summary": "Get user properties used amount",
                "parameters": [
                    {
                        "description": "User properties used amount request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helper.UserCostsAmountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully retrieved user properties used amount",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "failed to parse user properties used amount request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "authenticate error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "failed to get user properties used amount",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/account/v1alpha1/costs/recharge": {
            "post": {
                "description": "Get user recharge amount within a specified time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RechargeAmount"
                ],
                "summary": "Get user recharge amount",
                "parameters": [
                    {
                        "description": "User recharge amount request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helper.UserCostsAmountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully retrieved user recharge amount",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "failed to parse user recharge amount request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "authenticate error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "failed to get user recharge amount",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/account/v1alpha1/namespaces": {
            "post": {
                "description": "Get the billing history namespace list from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BillingHistory"
                ],
                "summary": "Get namespace billing history list",
                "parameters": [
                    {
                        "description": "Namespace billing history request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helper.NamespaceBillingHistoryReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully retrieved namespace billing history list",
                        "schema": {
                            "$ref": "#/definitions/helper.NamespaceBillingHistoryRespData"
                        }
                    },
                    "400": {
                        "description": "failed to parse namespace billing history request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "authenticate error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "failed to get namespace billing history list",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/account/v1alpha1/properties": {
            "post": {
                "description": "Get properties from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Properties"
                ],
                "summary": "Get properties",
                "parameters": [
                    {
                        "description": "auth request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helper.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully retrieved properties",
                        "schema": {
                            "$ref": "#/definitions/helper.GetPropertiesResp"
                        }
                    },
                    "401": {
                        "description": "authenticate error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "failed to get properties",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.PropertyQuery": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string",
                    "example": "gpu-tesla-v100"
                },
                "name": {
                    "type": "string",
                    "example": "cpu"
                },
                "unit": {
                    "type": "string",
                    "example": "1m"
                },
                "unit_price": {
                    "type": "number",
                    "example": 10000
                }
            }
        },
        "helper.Auth": {
            "type": "object",
            "required": [
                "kubeConfig",
                "owner"
            ],
            "properties": {
                "kubeConfig": {
                    "type": "string"
                },
                "owner": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "helper.ErrorMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "authentication failure"
                }
            }
        },
        "helper.GetPropertiesResp": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/helper.GetPropertiesRespData"
                },
                "message": {
                    "type": "string",
                    "example": "successfully retrieved properties"
                }
            }
        },
        "helper.GetPropertiesRespData": {
            "type": "object",
            "properties": {
                "properties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.PropertyQuery"
                    }
                }
            }
        },
        "helper.NamespaceBillingHistoryReq": {
            "type": "object",
            "required": [
                "kubeConfig",
                "owner"
            ],
            "properties": {
                "endTime": {
                    "type": "string",
                    "example": "2021-12-01T00:00:00Z"
                },
                "kubeConfig": {
                    "type": "string"
                },
                "owner": {
                    "type": "string",
                    "example": "admin"
                },
                "startTime": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "type": {
                    "description": "@Summary Type of the request (optional)\n@Description Type of the request (optional)\n@JSONSchema",
                    "type": "integer"
                }
            }
        },
        "helper.NamespaceBillingHistoryRespData": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "[\"ns-admin\"",
                        "\"ns-test1\"]"
                    ]
                }
            }
        },
        "helper.UserCostsAmountReq": {
            "type": "object",
            "required": [
                "kubeConfig",
                "owner"
            ],
            "properties": {
                "endTime": {
                    "type": "string",
                    "example": "2021-12-01T00:00:00Z"
                },
                "kubeConfig": {
                    "type": "string"
                },
                "owner": {
                    "type": "string",
                    "example": "admin"
                },
                "startTime": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1alpha1",
	Host:             "localhost:2333",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "sealos account service",
	Description:      "Your API description.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
