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
        "/api/route": {
            "get": {
                "responses": {}
            }
        },
        "/home/DirectPurchase": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "收件地址",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "2003": {
                        "description": "（直接购买）购买成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4002": {
                        "description": "密码错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4014": {
                        "description": "无对应书籍",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/home/PurchaseFromSC": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电子邮件",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "2000": {
                        "description": "登录成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4008": {
                        "description": "用户余额不足",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4013": {
                        "description": "购物车为空",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/home/ShoppingCarts/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物"
                ],
                "summary": "添加书籍到购物车",
                "parameters": [
                    {
                        "type": "string",
                        "description": "书籍的ISBN编号",
                        "name": "ISBN",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "出售者用户名",
                        "name": "SellerName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "当前用户名",
                        "name": "userName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "2009": {
                        "description": "添加成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4009": {
                        "description": "添加失败",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "5002": {
                        "description": "数据绑定失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/home/ShoppingCarts/remove": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物"
                ],
                "summary": "将书籍从购物车清除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "书籍的ISBN编号",
                        "name": "ISBN",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "出售者用户名",
                        "name": "SellerName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "2010": {
                        "description": "删除成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4010": {
                        "description": "从购物车删除失败",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "5001": {
                        "description": "非法图书ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电子邮件",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "2000": {
                        "description": "登录成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4001": {
                        "description": "该用户不存在",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4002": {
                        "description": "密码错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "注册一个新用户账号",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "summary": "注册新用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "confirmPassword",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电子邮件",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "该用户名已被使用",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "4003": {
                        "description": "重复密码与第一次输入的密码不一致，请重新输入",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "用户名不能为空",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "密码加密错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Your API Title",
	Description:      "This is a sample API for demonstration purposes. It includes multiple endpoints for various operations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
