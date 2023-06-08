// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/authors": {
            "get": {
                "description": "Responds with the list of all authors as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Get authors array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    }
                }
            },
            "put": {
                "description": "Takes a author JSON and edit an in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Edit an author",
                "parameters": [
                    {
                        "description": "Author JSON",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a author JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Add a new author",
                "parameters": [
                    {
                        "description": "Author JSON",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/authors/{id}": {
            "get": {
                "description": "Returns the author whose id value matches the id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Get single author by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search author by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove an author from DB by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Delete an author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete author by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "description": "Responds with the list of all books as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get books array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    }
                }
            },
            "put": {
                "description": "Takes a book JSON and edit an in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Edit an book",
                "parameters": [
                    {
                        "description": "Book JSON",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a book JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Add a new book",
                "parameters": [
                    {
                        "description": "Book JSON",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "Returns the book whose id value matches the id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get single book by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search book by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove an book from DB by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete an book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete book by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "description": "Responds with the list of all categories as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get categories array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    }
                }
            },
            "put": {
                "description": "Takes a category JSON and edit an in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Edit an category",
                "parameters": [
                    {
                        "description": "Category JSON",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a category JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Add a new category",
                "parameters": [
                    {
                        "description": "Category JSON",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/categories/{id}": {
            "get": {
                "description": "Returns the category whose id value matches the id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get single category by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search category by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove an category from DB by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Delete an category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete category by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Author": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Book"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/entities.Author"
                },
                "author_id": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Category"
                    }
                },
                "description": {
                    "type": "string"
                },
                "edition": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "isbn": {
                    "description": "Export a field by starting it with an uppercase letter.",
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "published_date": {
                    "type": "string"
                },
                "publishers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Publisher"
                    }
                },
                "reader": {
                    "$ref": "#/definitions/entities.Reader"
                },
                "reader_id": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "entities.Category": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Book"
                    }
                },
                "category_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "entities.Publisher": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Book"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "publisher_name": {
                    "type": "string"
                }
            }
        },
        "entities.Reader": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Book"
                    }
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
