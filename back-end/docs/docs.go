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
        "/jogar": {
            "post": {
                "description": "Apos digitar uma equacao e clicar em jogar a equacao e enviada para validacao",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Valida a jogada",
                "parameters": [
                    {
                        "description": "equacao do usuario",
                        "name": "equacao",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.clientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.serverResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "server.clientRequest": {
            "type": "object",
            "properties": {
                "expressao": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "server.serverResponse": {
            "type": "object",
            "properties": {
                "dica": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "expressao": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "match": {
                    "type": "boolean"
                },
                "msg_erro": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
