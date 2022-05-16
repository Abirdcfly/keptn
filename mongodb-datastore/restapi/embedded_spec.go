// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/cloudevents+json",
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "mongodb-datastore",
    "version": "develop"
  },
  "basePath": "/",
  "paths": {
    "/event": {
      "get": {
        "tags": [
          "event"
        ],
        "summary": "Gets events from the data store, either keptnContext or project must be specified",
        "operationId": "getEvents",
        "parameters": [
          {
            "type": "string",
            "description": "keptnContext of the events to get",
            "name": "keptnContext",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Type of the keptn cloud event",
            "name": "type",
            "in": "query"
          },
          {
            "type": "string",
            "description": "From time to fetch keptn cloud events",
            "name": "fromTime",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Before time to fetch keptn cloud events",
            "name": "beforeTime",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Set to load only root events",
            "name": "root",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the project",
            "name": "project",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the stage",
            "name": "stage",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the service",
            "name": "service",
            "in": "query"
          },
          {
            "type": "string",
            "description": "EventID",
            "name": "eventID",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the event source",
            "name": "source",
            "in": "query"
          },
          {
            "$ref": "#/parameters/pagesizeParam"
          },
          {
            "$ref": "#/parameters/pageParam"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "object",
              "properties": {
                "events": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/KeptnContextExtendedCE"
                  },
                  "x-isnullable": true,
                  "x-nullable": true
                },
                "nextPageKey": {
                  "description": "Pointer to the next page",
                  "type": "string"
                },
                "pageSize": {
                  "description": "Size of the returned page",
                  "type": "integer"
                },
                "totalCount": {
                  "description": "Total number of events",
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/event/type/{eventType}": {
      "get": {
        "tags": [
          "event"
        ],
        "summary": "Gets events by their type from the mongodb, required filter are either 'data.project:\u003cproject-name\u003e' or 'shkeptncontext:\u003ckeptn-context-id\u003e'",
        "operationId": "getEventsByType",
        "parameters": [
          {
            "type": "string",
            "name": "filter",
            "in": "query",
            "required": true
          },
          {
            "type": "boolean",
            "name": "excludeInvalidated",
            "in": "query"
          },
          {
            "type": "string",
            "name": "fromTime",
            "in": "query"
          },
          {
            "$ref": "#/parameters/limitParam"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "object",
              "properties": {
                "events": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/KeptnContextExtendedCE"
                  }
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "eventType",
          "in": "path",
          "required": true
        }
      ]
    },
    "/health": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "INTERNAL Endpoint: Health endpoint",
        "operationId": "getHealth",
        "deprecated": true,
        "responses": {
          "200": {
            "description": "healthy"
          }
        }
      }
    }
  },
  "definitions": {
    "KeptnContextExtendedCE": {
      "type": "object",
      "x-go-type": {
        "hints": {
          "noValidation": true
        },
        "import": {
          "alias": "keptnapi",
          "package": "github.com/keptn/go-utils/pkg/api/models"
        },
        "type": "KeptnContextExtendedCE"
      },
      "x-nullable": true
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "parameters": {
    "limitParam": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "description": "Page size to be returned",
      "name": "limit",
      "in": "query"
    },
    "pageParam": {
      "type": "string",
      "description": "Key of the page to be returned",
      "name": "nextPageKey",
      "in": "query"
    },
    "pagesizeParam": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "description": "Page size to be returned",
      "name": "pageSize",
      "in": "query"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/cloudevents+json",
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "mongodb-datastore",
    "version": "develop"
  },
  "basePath": "/",
  "paths": {
    "/event": {
      "get": {
        "tags": [
          "event"
        ],
        "summary": "Gets events from the data store, either keptnContext or project must be specified",
        "operationId": "getEvents",
        "parameters": [
          {
            "type": "string",
            "description": "keptnContext of the events to get",
            "name": "keptnContext",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Type of the keptn cloud event",
            "name": "type",
            "in": "query"
          },
          {
            "type": "string",
            "description": "From time to fetch keptn cloud events",
            "name": "fromTime",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Before time to fetch keptn cloud events",
            "name": "beforeTime",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Set to load only root events",
            "name": "root",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the project",
            "name": "project",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the stage",
            "name": "stage",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the service",
            "name": "service",
            "in": "query"
          },
          {
            "type": "string",
            "description": "EventID",
            "name": "eventID",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the event source",
            "name": "source",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 20,
            "description": "Page size to be returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Key of the page to be returned",
            "name": "nextPageKey",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "object",
              "properties": {
                "events": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/KeptnContextExtendedCE"
                  },
                  "x-isnullable": true,
                  "x-nullable": true
                },
                "nextPageKey": {
                  "description": "Pointer to the next page",
                  "type": "string"
                },
                "pageSize": {
                  "description": "Size of the returned page",
                  "type": "integer"
                },
                "totalCount": {
                  "description": "Total number of events",
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/event/type/{eventType}": {
      "get": {
        "tags": [
          "event"
        ],
        "summary": "Gets events by their type from the mongodb, required filter are either 'data.project:\u003cproject-name\u003e' or 'shkeptncontext:\u003ckeptn-context-id\u003e'",
        "operationId": "getEventsByType",
        "parameters": [
          {
            "type": "string",
            "name": "filter",
            "in": "query",
            "required": true
          },
          {
            "type": "boolean",
            "name": "excludeInvalidated",
            "in": "query"
          },
          {
            "type": "string",
            "name": "fromTime",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 20,
            "description": "Page size to be returned",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "object",
              "properties": {
                "events": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/KeptnContextExtendedCE"
                  }
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "eventType",
          "in": "path",
          "required": true
        }
      ]
    },
    "/health": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "INTERNAL Endpoint: Health endpoint",
        "operationId": "getHealth",
        "deprecated": true,
        "responses": {
          "200": {
            "description": "healthy"
          }
        }
      }
    }
  },
  "definitions": {
    "KeptnContextExtendedCE": {
      "type": "object",
      "x-go-type": {
        "hints": {
          "noValidation": true
        },
        "import": {
          "alias": "keptnapi",
          "package": "github.com/keptn/go-utils/pkg/api/models"
        },
        "type": "KeptnContextExtendedCE"
      },
      "x-nullable": true
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "parameters": {
    "limitParam": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "description": "Page size to be returned",
      "name": "limit",
      "in": "query"
    },
    "pageParam": {
      "type": "string",
      "description": "Key of the page to be returned",
      "name": "nextPageKey",
      "in": "query"
    },
    "pagesizeParam": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "description": "Page size to be returned",
      "name": "pageSize",
      "in": "query"
    }
  }
}`))
}
