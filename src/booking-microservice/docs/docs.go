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
        "description": "List APIs of Booking Service",
        "title": "Booking Service API Document",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "localhost:5001",
    "basePath": "/api/v1",
    "paths": {
        "/Job/AddJob": {
            "post": {
                "description": "Post a new job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "Add a new job",
                "parameters": [
                    {
                        "type": "string",
                        "description": "job",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddJob"
                        },
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/Job/RemoveJob": {
            "post": {
                "description": "Remove a job from the list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "Remove a job from the list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/Job/UpdateJob": {
            "post": {
                "description": "Update a current job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "Update a current job",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "job",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddJob"
                        },
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "models.AddJob": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2024-01-22T02:54:15.225Z"
                },
                "service": {
                    "type": "string",
                    "example": "Clean up the air-condition"
                },
                "note": {
                    "type": "string",
                }
            }
        },
        "models.UpdateJob": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2024-01-22T02:54:15.225Z"
                },
                "service": {
                    "type": "string",
                    "example": "Clean up the air-condition"
                },
                "note": {
                    "type": "string",
                }
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
        }
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