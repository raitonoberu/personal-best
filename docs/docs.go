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
            "name": "raitonoberu",
            "url": "http://raitonobe.ru",
            "email": "raitonoberu@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/login": {
            "post": {
                "description": "Login user, return JWT token \u0026 ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuthResponse"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "Register new unverified player\n\"birth_date\" must have format 1889-04-20\n\"phone\" must start with +\n\"telegram\" must start with @",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuthResponse"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete current user",
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Return single user by ID\n\"player\" may not be present (trainer / admin)\nplayer.preparation, player.position may not be present",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AuthResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "model.GetUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "player": {
                    "$ref": "#/definitions/model.Player"
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.Player": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "is_male": {
                    "type": "boolean"
                },
                "is_verified": {
                    "type": "boolean"
                },
                "phone": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "preparation": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                }
            }
        },
        "model.RegisterRequest": {
            "type": "object",
            "required": [
                "birth_date",
                "email",
                "first_name",
                "is_male",
                "last_name",
                "middle_name",
                "password",
                "phone",
                "telegram"
            ],
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "is_male": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "telegram": {
                    "description": "TODO: more validations",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Personal Best API",
	Description:      "neмытьlE yблюdki",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
