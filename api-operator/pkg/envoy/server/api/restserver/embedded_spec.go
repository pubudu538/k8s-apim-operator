// Code generated by go-swagger; DO NOT EDIT.

package restserver

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
    "application/json"
  ],
  "produces": [
    "application/zip"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This document specifies a **RESTful API** for allowing you to access internal data .\nPlease see [full swagger definition](https://raw.githubusercontent.com/wso2/carbon-apimgt/master/components/apimgt/org.wso2.carbon.apimgt.internal.service/src/main/resources/api.yaml) of the API which is written using [swagger 2.0](http://swagger.io/) specification.\n",
    "title": "Internal Utility API",
    "contact": {
      "name": "WSO2",
      "url": "http://wso2.com/products/api-manager/",
      "email": "architecture@wso2.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "apis.wso2.com",
  "basePath": "/operator/2.0",
  "paths": {
    "/apis": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to get all the APIs deployed in Kubernetes.\n",
        "produces": [
          "application/zip"
        ],
        "tags": [
          "APIs (All)"
        ],
        "summary": "Get all apis in a zip file",
        "responses": {
          "200": {
            "description": "Sent.\nAPIs sent Successfully.\n",
            "schema": {
              "type": "string",
              "format": "binary"
            }
          },
          "403": {
            "description": "Forbidden\nNot Authorized to send.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not Found.\nRequested APIs to send not found.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error.\nError in sending APIs.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "title": "Error object returned with 4XX HTTP status",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "integer",
          "format": "int64"
        },
        "description": {
          "description": "A detail description about the error message.\n",
          "type": "string"
        },
        "error": {
          "description": "If there are more than one error list them out.\nFor example, list out validation errors by each field.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ErrorListItem"
          }
        },
        "message": {
          "description": "Error message.",
          "type": "string"
        },
        "moreInfo": {
          "description": "Preferably an url with more details about the error.\n",
          "type": "string"
        }
      }
    },
    "ErrorListItem": {
      "title": "Description of individual errors that may have occurred during a request.",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "message": {
          "description": "Description about individual errors occurred\n",
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "tenant": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BasicAuth": {
      "type": "basic"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/zip"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This document specifies a **RESTful API** for allowing you to access internal data .\nPlease see [full swagger definition](https://raw.githubusercontent.com/wso2/carbon-apimgt/master/components/apimgt/org.wso2.carbon.apimgt.internal.service/src/main/resources/api.yaml) of the API which is written using [swagger 2.0](http://swagger.io/) specification.\n",
    "title": "Internal Utility API",
    "contact": {
      "name": "WSO2",
      "url": "http://wso2.com/products/api-manager/",
      "email": "architecture@wso2.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "apis.wso2.com",
  "basePath": "/operator/2.0",
  "paths": {
    "/apis": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to get all the APIs deployed in Kubernetes.\n",
        "produces": [
          "application/zip"
        ],
        "tags": [
          "APIs (All)"
        ],
        "summary": "Get all apis in a zip file",
        "responses": {
          "200": {
            "description": "Sent.\nAPIs sent Successfully.\n",
            "schema": {
              "type": "string",
              "format": "binary"
            }
          },
          "403": {
            "description": "Forbidden\nNot Authorized to send.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not Found.\nRequested APIs to send not found.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error.\nError in sending APIs.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "title": "Error object returned with 4XX HTTP status",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "integer",
          "format": "int64"
        },
        "description": {
          "description": "A detail description about the error message.\n",
          "type": "string"
        },
        "error": {
          "description": "If there are more than one error list them out.\nFor example, list out validation errors by each field.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ErrorListItem"
          }
        },
        "message": {
          "description": "Error message.",
          "type": "string"
        },
        "moreInfo": {
          "description": "Preferably an url with more details about the error.\n",
          "type": "string"
        }
      }
    },
    "ErrorListItem": {
      "title": "Description of individual errors that may have occurred during a request.",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "message": {
          "description": "Description about individual errors occurred\n",
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "tenant": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BasicAuth": {
      "type": "basic"
    }
  }
}`))
}
