// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/flights": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Get all flights",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Flight"
                            }
                        }
                    }
                }
            }
        },
        "/sales": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sales"
                ],
                "summary": "Creates a sale",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.saleRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.saleResponseBody"
                        }
                    }
                }
            }
        },
        "/sales/:sale_id/payment": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sales"
                ],
                "summary": "Creates a payment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sale ID",
                        "name": "sale_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.paymentRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.paymentResponseBody"
                        }
                    }
                }
            }
        },
        "/seats": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "summary": "Get all seats",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "origin ID",
                        "name": "origin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "destination ID",
                        "name": "destination",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Flight"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.paymentRequestBody": {
            "type": "object",
            "properties": {
                "card_number": {
                    "type": "integer"
                },
                "expiration_date": {
                    "type": "string"
                },
                "security_number": {
                    "type": "integer"
                }
            }
        },
        "controller.paymentResponseBody": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "controller.saleRequestBody": {
            "type": "object",
            "properties": {
                "dni": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "seat_id": {
                    "type": "integer"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "controller.saleResponseBody": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "passenger": {
                    "$ref": "#/definitions/model.Passenger"
                },
                "price": {
                    "type": "number"
                },
                "reservation_date": {
                    "type": "string"
                },
                "seat_id": {
                    "type": "integer"
                }
            }
        },
        "model.Airport": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "place": {
                    "$ref": "#/definitions/model.Place"
                },
                "placeID": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.Flight": {
            "type": "object",
            "properties": {
                "basePrice": {
                    "type": "number"
                },
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "destination": {
                    "$ref": "#/definitions/model.Airport"
                },
                "destinationID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "origin": {
                    "$ref": "#/definitions/model.Airport"
                },
                "originID": {
                    "type": "integer"
                }
            }
        },
        "model.Passenger": {
            "type": "object",
            "properties": {
                "dni": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "surName": {
                    "type": "string"
                }
            }
        },
        "model.Place": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Airport Application",
	Description:      "Airport Rest API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
