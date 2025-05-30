// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://81.177.220.96/",
            "email": "example@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/machine/{id}": {
            "put": {
                "description": "change machine status to need repair machine:id from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "machine"
                ],
                "summary": "change machine status to need repair",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Machine id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No connection",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "missing id",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/shift/workers/": {
            "post": {
                "description": "create new shift worker in db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shift worker"
                ],
                "summary": "create a shift worker",
                "parameters": [
                    {
                        "description": "Task data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.addWorkerDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.ShiftWorker"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/shifts": {
            "get": {
                "description": "get out shifts that are active.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shifts"
                ],
                "summary": "get out shifts that are active",
                "responses": {
                    "200": {
                        "description": "List of active shifts",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/services.Shift"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/task/": {
            "post": {
                "description": "create new shift in db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shift"
                ],
                "summary": "create a shift",
                "parameters": [
                    {
                        "description": "Shift data",
                        "name": "shift",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.createShiftDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.Shift"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/task/{id}": {
            "delete": {
                "description": "Delete a task:id from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No connection",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "missing id",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "patch": {
                "description": "commands for update task: inProgress, completed, verified, failed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "update task task by command",
                "parameters": [
                    {
                        "description": "Task data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.updateTaskDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/user/": {
            "post": {
                "description": "create new user in db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create a user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.createUserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.User"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "summary": "Get list of users",
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/services.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "description": "Return list of shift workers  by shift id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shift worker"
                ],
                "summary": "Get list of shift workers  by shift id",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "id",
                        "description": "Shift id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of shift workers by shift id",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/services.ShiftWorker"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user:id from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No connection",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "missing id",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/users/{role}": {
            "get": {
                "description": "Return list of users with role.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Get list of users with role",
                "parameters": [
                    {
                        "type": "string",
                        "format": "id",
                        "description": "Users role",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of users with role",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/services.User"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.addWorkerDTO": {
            "type": "object",
            "properties": {
                "shiftId": {
                    "type": "integer"
                },
                "workerid": {
                    "type": "integer"
                }
            }
        },
        "handler.createShiftDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "machineid": {
                    "type": "integer"
                },
                "shiftmaster": {
                    "type": "integer"
                }
            }
        },
        "handler.createUserDTO": {
            "type": "object",
            "properties": {
                "bitrixid": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "handler.updateTaskDTO": {
            "type": "object",
            "properties": {
                "command": {
                    "type": "string"
                },
                "comment": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "services.Shift": {
            "type": "object",
            "properties": {
                "createdat": {
                    "type": "string"
                },
                "deactivatedat": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isactive": {
                    "type": "boolean"
                },
                "machineid": {
                    "type": "integer"
                },
                "shiftMaster": {
                    "type": "integer"
                }
            }
        },
        "services.ShiftTask": {
            "type": "object",
            "properties": {
                "shiftid": {
                    "type": "integer"
                },
                "taskid": {
                    "type": "integer"
                }
            }
        },
        "services.ShiftWorker": {
            "type": "object",
            "properties": {
                "shiftid": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "services.User": {
            "type": "object",
            "properties": {
                "bitrixid": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Shift manager api",
	Description:      "API for managing shifts and adding tasks",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
