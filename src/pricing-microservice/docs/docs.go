// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-24 07:27:48.089020537 +0000 UTC m=+0.039932092

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "List APIs of Calculate Price Service",
        "title": "Calculate Price Service API Document",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "localhost:5002",
    "basePath": "/api/v1",
    "paths": {
        "/CalculatePrice/AllJobs": {
            "get": {
                "description": "Calculate price for all jobs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "Calculate price for all jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.CalculatePriceJobs"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/CalculatePrice/JobByTime": {
            "get": {
                "description": "Calculate price for all jobs by time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "endTime",
                        "in": "query",
                        "required": true
                    },
                ],
                "summary": "Calculate price for all jobs by time",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.CalculatePriceJobsByTime"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
    },
    "definitions": {
        "models.CalculatePriceJobs": {
            "type": "object",
            "properties": {
                "totalPriceOfAllJobs": {
                    "type": "integer",
                },
            }
        },
        "models.CalculatePriceJobsByTime": {
            "type": "object",
            "properties": {
                "totalPriceOfAllJobs": {
                    "type": "integer",
                },
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 27
                },
                "message": {
                    "type": "string",
                    "example": "Error message"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}