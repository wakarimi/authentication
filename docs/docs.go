// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Zalimannard",
            "email": "zalimannard@mail.ru"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/register/user": {
            "post": {
                "description": "Register a new user with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Register payload",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.RegisterUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Ok."
                    },
                    "400": {
                        "description": "Error.",
                        "schema": {
                            "$ref": "#/definitions/types.Error"
                        }
                    },
                    "500": {
                        "description": "Error.",
                        "schema": {
                            "$ref": "#/definitions/types.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.RegisterUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Password of the user to be registered.\nRequired: true\nExample: querty01",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                },
                "username": {
                    "description": "Username of the user to be registered.\nRequired: true\nExample: Zalimannard",
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "types.Error": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "description": "The error message.\nExample: User already exists",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0",
	Host:             "localhost:8020",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Wakarimi Authentication API",
	Description:      "This is the authentication service for Wakarimi.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
