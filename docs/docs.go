// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/age-rating-categories": {
            "get": {
                "description": "Get List of Age Rating Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get All Age Rating Category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AgeRatingCategory"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Age Rating Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Create Age Rating Category",
                "parameters": [
                    {
                        "description": "the body to create new age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AgeRatingCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            }
        },
        "/age-rating-categories/{id}": {
            "get": {
                "description": "Get one Age Rating category by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get an Age Rating Category by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Eating Category Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Age Rating category by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Delete an Age Rating Category by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Eating Category Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update one Age Rating category by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Update an Age Rating Category by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Eating Category Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to create new age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AgeRatingCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            }
        },
        "/age-rating-categories/{id}/movies": {
            "get": {
                "description": "Get all movies of Age Rating Category by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get movies by Age Rating Category by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Eating Category Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "Get List of Movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get All Movie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Create Movie",
                "parameters": [
                    {
                        "description": "the body to create new Movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MovieInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "description": "Get one Movie by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get an Movie by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Movie by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Delete an Movie by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update one Movie by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Update an Movie by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to create new Movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MovieInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AgeRatingCategoryInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.MovieInput": {
            "type": "object",
            "properties": {
                "age_rating_category_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.AgeRatingCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "age_rating_category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
