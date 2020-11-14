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
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/leaderboard": {
            "get": {
                "description": "Sends back the leaderboard in descending order of level, and for users on the same level, in the ascending order of last successful submission time.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kryptos"
                ],
                "summary": "return leaderboard.",
                "responses": {
                    "200": {
                        "description": "Returns the leaderboard",
                        "schema": {
                            "$ref": "#/definitions/handlers.swagUser"
                        }
                    },
                    "500": {
                        "description": "Could not serialize json",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "get": {
                "description": "Sends \"Test\" back. Use this to check if the server is up.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Kryptos"
                ],
                "summary": "Server health check.",
                "responses": {
                    "200": {
                        "description": "Server is up",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/question": {
            "get": {
                "description": "Sends back the question for the level the user is on. If this is a new user, a user instance is created in the DB and the first question is returned.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kryptos"
                ],
                "summary": "returns the question for the level the user is on.",
                "responses": {
                    "200": {
                        "description": "Returns the question and it's details.",
                        "schema": {
                            "$ref": "#/definitions/handlers.swagQresponse"
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
        "/api/submit": {
            "post": {
                "description": "takes a post request with the answer attempt.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Kryptos"
                ],
                "summary": "takes a post request with the answer attempt.",
                "parameters": [
                    {
                        "description": "Answer format",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.swagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns 'success'",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Returns 'fail'",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.swagQresponse": {
            "type": "object",
            "properties": {
                "hints": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['Hint 1'",
                        " 'Hint 2']"
                    ]
                },
                "image_level": {
                    "type": "boolean",
                    "example": true
                },
                "level_file": {
                    "type": "string",
                    "example": "url_of_image"
                },
                "number": {
                    "type": "integer",
                    "example": 1
                },
                "question": {
                    "type": "string",
                    "example": "What is MEC's techfest?"
                }
            }
        },
        "handlers.swagRequest": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "excel"
                }
            }
        },
        "handlers.swagUser": {
            "type": "object",
            "properties": {
                "curr_level": {
                    "type": "integer",
                    "example": 18
                },
                "name": {
                    "type": "string",
                    "example": "Aswin G"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Excel Play Kryptos API",
	Description: "This is the swagger doc for the API for Excel Play Kryptos.",
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
