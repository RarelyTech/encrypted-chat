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
        "/group": {
            "post": {
                "description": "create an group (need signature verify)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Group Create",
                "operationId": "GroupCreate",
                "parameters": [
                    {
                        "description": "Group info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GroupCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/app.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.GroupDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/member/nonce/{address}": {
            "get": {
                "description": "Getting nonce string to be signed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Member Signin Nonce",
                "operationId": "MemberNonce",
                "responses": {
                    "200": {
                        "description": "Response success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/app.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.MemberNonceRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/member/{address}": {
            "get": {
                "description": "Getting member's profile using address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Member Profile",
                "operationId": "MemberProfile",
                "responses": {
                    "200": {
                        "description": "Response success",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/memeber": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Member Signin",
                "operationId": "MemberSignin",
                "responses": {
                    "200": {
                        "description": "Response success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/app.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.MemberSigninRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.GroupCreateReq": {
            "type": "object",
            "required": [
                "category",
                "name"
            ],
            "properties": {
                "category": {
                    "description": "group's category",
                    "type": "string"
                },
                "intro": {
                    "description": "group intro",
                    "type": "string"
                },
                "maxMembers": {
                    "description": "group's max members",
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 2
                },
                "public": {
                    "description": "` + "`" + `true` + "`" + ` create public group, ` + "`" + `false` + "`" + ` create private group",
                    "type": "boolean"
                }
            }
        },
        "model.GroupDetail": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "intro": {
                    "type": "string"
                },
                "membersMax": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "description": "group is belongs to current member",
                    "type": "boolean"
                },
                "public": {
                    "description": "group is public or private",
                    "type": "boolean"
                }
            }
        },
        "model.MemberNonceRes": {
            "type": "object",
            "properties": {
                "nonce": {
                    "type": "string"
                }
            }
        },
        "model.MemberProfile": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar": {
                    "description": "optional",
                    "type": "string"
                },
                "intro": {
                    "description": "optional",
                    "type": "string"
                },
                "nickname": {
                    "description": "optional",
                    "type": "string"
                }
            }
        },
        "model.MemberSigninRes": {
            "type": "object",
            "properties": {
                "profile": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.MemberProfile"
                        }
                    ]
                },
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ChatPuppy Api Doc",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}