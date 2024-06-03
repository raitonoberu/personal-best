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
        "/api/admin/users": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create new user with desired params.\nPlayer-related params only required when creating player\n(is_male, phone, telegram, birth_date).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AdminCreateUserRequest"
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
        "/api/admin/users/{id}": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update user.\nPlayer-related params only changed when updating player\n(is_male, phone, telegram, birth_date).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AdminUpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/competitions": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List all competitions from new to old\nFor now there is no way to get start/end but im working on it :)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "competition"
                ],
                "summary": "List competition",
                "parameters": [
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListCompetitionsResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create new competition.\nDays must be different (no same day twice).\nTime must be in format HH:MM.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "competition"
                ],
                "summary": "Create competition",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCompetitionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CreateCompetitionResponse"
                        }
                    }
                }
            }
        },
        "/api/competitions/{comp_id}/matches/{id}": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update match score\nMatch must have players \u0026 score must NOT be set already\nThis will fill next match's players",
                "tags": [
                    "match"
                ],
                "summary": "Update match",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comp_id",
                        "name": "comp_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMatchRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/competitions/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Return competition by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "competition"
                ],
                "summary": "Get competition",
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
                            "$ref": "#/definitions/model.GetCompetitionResponse"
                        }
                    }
                }
            }
        },
        "/api/competitions/{id}/matches": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List all matches with all players",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "List competition matches",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "competition id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListMatchesResponse"
                        }
                    }
                }
            }
        },
        "/api/competitions/{id}/registrations": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List registrations for competition",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "registration"
                ],
                "summary": "List competition registrations",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Registration"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is made for trainers/admins\nHere you can approve or ban players",
                "tags": [
                    "registration"
                ],
                "summary": "Update registration",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Competition must not be closed yet",
                "tags": [
                    "registration"
                ],
                "summary": "Unregister for competition",
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
                        "description": "OK"
                    }
                }
            }
        },
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
        "/api/roles": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List all available roles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "List roles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.RoleResponse"
                            }
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
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update user info.\nPlayer fields will be added later.",
                "tags": [
                    "user"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserRequest"
                        }
                    }
                ],
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
        "model.AdminCreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "middle_name",
                "password",
                "role_id"
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
                "role_id": {
                    "type": "integer"
                },
                "telegram": {
                    "type": "string"
                }
            }
        },
        "model.AdminUpdateUserRequest": {
            "type": "object",
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
                "role_id": {
                    "type": "integer"
                },
                "telegram": {
                    "type": "string"
                }
            }
        },
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
        "model.CompetitionDay": {
            "type": "object",
            "required": [
                "date",
                "end_time",
                "start_time"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "model.CreateCompetitionRequest": {
            "type": "object",
            "required": [
                "age",
                "closes_at",
                "days",
                "description",
                "name",
                "size",
                "tours"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "closes_at": {
                    "type": "string"
                },
                "days": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CompetitionDay"
                    }
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "tours": {
                    "type": "integer"
                }
            }
        },
        "model.CreateCompetitionResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.GetCompetitionResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "closes_at": {
                    "type": "string"
                },
                "days": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CompetitionDay"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "tours": {
                    "type": "integer"
                },
                "trainer": {
                    "$ref": "#/definitions/model.GetUserResponse"
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
                    "$ref": "#/definitions/model.PlayerResponse"
                },
                "role_id": {
                    "type": "integer"
                }
            }
        },
        "model.ListCompetitionsResponse": {
            "type": "object",
            "properties": {
                "competitions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.GetCompetitionResponse"
                    }
                },
                "count": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.ListMatchesResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "matches": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Match"
                    }
                },
                "total": {
                    "type": "integer"
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
        "model.Match": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "left_score": {
                    "type": "integer"
                },
                "left_team": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.GetUserResponse"
                    }
                },
                "right_score": {
                    "type": "integer"
                },
                "right_team": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.GetUserResponse"
                    }
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "model.PlayerResponse": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "is_male": {
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
        },
        "model.Registration": {
            "type": "object",
            "properties": {
                "is_approved": {
                    "type": "boolean"
                },
                "is_dropped": {
                    "type": "boolean"
                },
                "user": {
                    "$ref": "#/definitions/model.GetUserResponse"
                }
            }
        },
        "model.RoleResponse": {
            "type": "object",
            "properties": {
                "can_create": {
                    "type": "boolean"
                },
                "can_participate": {
                    "type": "boolean"
                },
                "can_view": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "is_free": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.UpdateMatchRequest": {
            "type": "object",
            "properties": {
                "left_score": {
                    "type": "integer"
                },
                "right_score": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "password": {
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
