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
        "description": "List APIs of Product Service",
        "title": "Product Service API Document",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "localhost:5000",
    "basePath": "/api/v1",
    "paths": {
        "/Job/GetAllServices": {
            "get": {
                "description": "List all existing services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "List all existing services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.GetAllService"
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
        "/Job/GetAllEmployees": {
            "get": {
                "description": "List all existing employees",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "List all existing employees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.GetAllEmployees"
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
        "/Job/GetAllJobsOrdered": {
            "get": {
                "description": "List all existing jobs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "List all existing jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.GetAllJobsOrdered"
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
        "/Job/GetJobsByID": {
            "get": {
                "description": "List all jobs by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "List all jobs by id",
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
                            "$ref": "#/definitions/models.GetJobsByID"
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
        "/Job/AddService": {
            "post": {
                "description": "Add a new service",
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
                        "description": "job",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.PostService"
                        },
                    }
                ],
                "summary": "Add a new service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Message"
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
        "/Job/AddEmployee": {
            "post": {
                "description": "Add a new employee",
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
                        "description": "job",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.PostEmployee"
                        },
                    }
                ],
                "summary": "Add a new employee",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Message"
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
        }
    },
    "definitions": {
        "models.GetAllService": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "name": {
                    "type": "string",
                    "example": "Clean up the air-condition"
                },
                "description": {
                    "type": "string",
                    "example": "Clean up the air-condition"
                },
                "price": {
                    "type": "int",
                    "example": "100"
                }
            }
        },
        "models.GetAllEmployees": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "name": {
                    "type": "string",
                    "example": "Hoang Bao Thinh"
                }
            }
        },
        "models.GetAllJobsOrdered": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "nameEmployee": {
                    "type": "string",
                    "example": "Thinh"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2018-10-24T07:27:48.089020537Z"
                },
                "note": {
                    "type": "string",
                    "example": "Note for the job"
                }
            }
        },
        "models.GetJobsByID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "nameEmployee": {
                    "type": "string",
                    "example": "Hoang Bao Thinh"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2018-10-24T07:27:48.089020537Z"
                },
                "note": {
                    "type": "string",
                    "example": "Note for the job"
                }
            }
        },
        "models.PostService": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "name": {
                    "type": "string",
                    "example": "House Children"
                },
                "description": {
                    "type": "string",
                    "example": ""
                },
                "price": {
                    "type": "interger",
                    "example": 0
                }
            }
        },
        "models.PostEmployee": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "c1aeb2be-eb17-4fc3-9152-c16bd75cfb28"
                },
                "name": {
                    "type": "string",
                    "example": "Hoang Bao Thinh"
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
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "5bb3695b82ebac0f76e1cafa"
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
