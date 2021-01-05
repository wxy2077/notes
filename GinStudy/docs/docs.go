// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/mini/api/article/get/category": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleAPI"
                ],
                "summary": "获取文章分类",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mini/api/article/get/detail": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleAPI"
                ],
                "summary": "获取文章详情",
                "parameters": [
                    {
                        "description": "文章href链接",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ArticleDetail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mini/api/article/get/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleAPI"
                ],
                "summary": "获取文章列表",
                "parameters": [
                    {
                        "description": "分页",
                        "name": "pageInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    },
                    {
                        "description": "文章分类",
                        "name": "cateInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ArticleCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mini/api/article/get/recommend": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleAPI"
                ],
                "summary": "获取推荐文章",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ArticleCategory": {
            "type": "object",
            "properties": {
                "cateId": {
                    "type": "integer"
                }
            }
        },
        "request.ArticleDetail": {
            "type": "object",
            "properties": {
                "href": {
                    "type": "string"
                }
            }
        },
        "request.PageInfo": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
