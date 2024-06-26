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
        "/accounts": {
            "post": {
                "description": "Create a new account for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create Account",
                "operationId": "create account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Authorization-Sign",
                        "name": "X-Authorization-Sign",
                        "in": "header"
                    },
                    {
                        "description": "Account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transactions/deposit": {
            "post": {
                "description": "Deposit money into an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Deposit Money",
                "operationId": "deposit money",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Authorization-Sign",
                        "name": "X-Authorization-Sign",
                        "in": "header"
                    },
                    {
                        "description": "Deposit",
                        "name": "deposit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddMoneyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transactions/transfer": {
            "post": {
                "description": "Transfer money between accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Transfer Money",
                "operationId": "transfer money",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Authorization-Sign",
                        "name": "X-Authorization-Sign",
                        "in": "header"
                    },
                    {
                        "description": "Transfer",
                        "name": "transfer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TransferMoneyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transactions/withdraw": {
            "post": {
                "description": "Withdraw money from an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Withdraw Money",
                "operationId": "withdraw money",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Authorization-Sign",
                        "name": "X-Authorization-Sign",
                        "in": "header"
                    },
                    {
                        "description": "Withdraw",
                        "name": "withdraw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddMoneyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "operationId": "create user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Authorization-Sign",
                        "name": "X-Authorization-Sign",
                        "in": "header"
                    },
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{email}/accounts": {
            "get": {
                "description": "Get accounts associated with a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Accounts",
                "operationId": "get accounts",
                "parameters": [
                    {
                        "type": "string",
                        "default": "srecko@gmail.com",
                        "description": "User Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "enums.AccountStatus": {
            "type": "string",
            "enum": [
                "Verified"
            ],
            "x-enum-varnames": [
                "VERIFIED"
            ]
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "acc_number": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/enums.AccountStatus"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Account"
                    }
                },
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.AddMoneyReq": {
            "type": "object",
            "properties": {
                "acc_number": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.TransferMoneyReq": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "from_acc_id": {
                    "type": "string"
                },
                "to_acc_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "e-wallet API",
	Description:      "This is a project for Arringo company.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
