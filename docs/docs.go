// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-27 18:04:49.325920191 +0800 CST m=+0.057596817

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Kepler server.",
        "title": "Kepler API",
        "contact": {},
        "license": {},
        "version": "v0.6.0"
    },
    "host": "127.0.0.1:8080",
    "basePath": "/",
    "paths": {
        "/key/gen": {
            "get": {
                "description": "公私钥生成，服务端不保存",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "key"
                ],
                "summary": "公私钥生成",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/types.Result"
                        }
                    }
                }
            }
        },
        "/qcp/apply": {
            "get": {
                "description": "联盟链申请查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qcp"
                ],
                "summary": "联盟链申请查询",
                "parameters": [
                    {
                        "minLength": 11,
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/module.ApplyQcp"
                        }
                    }
                }
            },
            "post": {
                "description": "联盟链证书申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qcp"
                ],
                "summary": "联盟链证书申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "联盟链ChainId",
                        "name": "qcpChainId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "公链ChainId",
                        "name": "qosChainId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "QCP公钥",
                        "name": "qcpPub",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minLength": 11,
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "申请说明",
                        "name": "info",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/qcp/apply/{id}": {
            "put": {
                "description": "申请审核",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qcp"
                ],
                "summary": "申请审核",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "申请ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "状态 1发放证书 2申请无效",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/qsc/apply": {
            "get": {
                "description": "联盟链申请查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qsc"
                ],
                "summary": "联盟链申请查询",
                "parameters": [
                    {
                        "minLength": 11,
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/module.ApplyQsc"
                        }
                    }
                }
            },
            "post": {
                "description": "联盟币证书申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qsc"
                ],
                "summary": "联盟币证书申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "联盟币名称",
                        "name": "qscName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "公链ChainId",
                        "name": "qosChainId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "QSC公钥",
                        "name": "qscPub",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用于接收联盟币的账户公钥",
                        "name": "bankerPub",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minLength": 11,
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "申请说明",
                        "name": "info",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/qsc/apply/{id}": {
            "put": {
                "description": "申请审核",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qsc"
                ],
                "summary": "申请审核",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "申请ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "状态 1发放证书 2申请无效",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "module.ApplyQcp": {
            "type": "object",
            "properties": {
                "createTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "qcpChainId": {
                    "type": "string"
                },
                "qcpPub": {
                    "type": "string"
                },
                "qosChainId": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updateTime": {
                    "type": "string"
                }
            }
        },
        "module.ApplyQsc": {
            "type": "object",
            "properties": {
                "bankerPub": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "qosChainId": {
                    "type": "string"
                },
                "qscName": {
                    "type": "string"
                },
                "qscPub": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updateTime": {
                    "type": "string"
                }
            }
        },
        "types.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
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
