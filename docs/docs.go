// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/avatars": {
            "get": {
                "description": "Get avatar's shop.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "avatar shop"
                ],
                "summary": "Get avatar's shop.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Only admin qdd avatar to shop.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "avatar shop"
                ],
                "summary": "Add avatar to shop",
                "parameters": [
                    {
                        "description": "request body for adding avatar to shop",
                        "name": "avatars",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.AvatarRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/friends": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Friends List.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "List Friends.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/friends/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Remove Friend by Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Remove Friend",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Friend ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/inventories": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get inventory by token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventories"
                ],
                "summary": "Get user's inventory.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add item to user inventories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventories"
                ],
                "summary": "Add item to inventories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "item id that you want to add",
                        "name": "item_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "item type that you want to add [avatar]",
                        "name": "item_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/invitations": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Check Invitation List.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitations"
                ],
                "summary": "List Invitation.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Accept All Invitations.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitations"
                ],
                "summary": "Accept All Invitations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Invite friend by email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitations"
                ],
                "summary": "Invite Friend",
                "parameters": [
                    {
                        "description": "request body invite friend",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.InviteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reject All Invitations.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitations"
                ],
                "summary": "Reject All Invitations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/invitations/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Accept Invitation by Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitations"
                ],
                "summary": "Accept Invitation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Invitation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reject Invitation by Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitations"
                ],
                "summary": "Reject Invitation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Invitation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "request body login",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "return new access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Refresh Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Create user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "request body register",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/themes": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Only admin add theme to shop.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "theme shop"
                ],
                "summary": "Add theme to shop",
                "parameters": [
                    {
                        "description": "request body for adding theme to shop",
                        "name": "avatars",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.ThemeCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "return list users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "User change profile information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Edit profile.",
                "parameters": [
                    {
                        "description": "request body for editing user's profile",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.EditUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/users/avatars/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get avatar by user Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get Avatar.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "User change avatar.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Change Avatar.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.DataResponse"
                        }
                    }
                }
            }
        },
        "/users/coins": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get coin by token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get Coin.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/forgot-password": {
            "put": {
                "description": "Send mail to reset password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "description": "request body to send mail",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.Email"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/reset-password": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Change password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "description": "request body to change password",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.Password"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "interfaces.AvatarRequest": {
            "type": "object",
            "required": [
                "assets",
                "name",
                "preview",
                "price"
            ],
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "preview": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "interfaces.EditUserRequest": {
            "type": "object",
            "properties": {
                "displayName": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "interfaces.Email": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "interfaces.InviteRequest": {
            "type": "object",
            "properties": {
                "targetMailAddress": {
                    "type": "string",
                    "example": "winner@mail.com"
                }
            }
        },
        "interfaces.Login": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "winner@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "winner"
                }
            }
        },
        "interfaces.Password": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "interfaces.RegisterRequest": {
            "type": "object",
            "required": [
                "characterId",
                "email",
                "password"
            ],
            "properties": {
                "birthday": {
                    "type": "string",
                    "example": "2023-08-12"
                },
                "characterId": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string",
                    "example": "winnerkypt"
                },
                "email": {
                    "type": "string",
                    "example": "winner@mail.com"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string",
                    "example": "winner"
                },
                "username": {
                    "type": "string",
                    "example": "winnerkypt"
                }
            }
        },
        "interfaces.ThemeCreateRequest": {
            "type": "object",
            "required": [
                "assets",
                "name",
                "price"
            ],
            "properties": {
                "assets": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "utils.DataResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "utils.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Meet Me API",
	Description:      "This is a API for Meet Me.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
