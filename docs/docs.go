// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-09-14 10:40:31.3530145 +0800 CST

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/user": {
            "post": {
                "description": "Add a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Add new user to the database",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\"}}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.CreateRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "birthday": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "sex": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
