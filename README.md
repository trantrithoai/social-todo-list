**Social To-Do List API**

This project is a simple RESTful API for a social to-do list application built with Go, Gin, and Gorm. It provides endpoints for creating, listing, retrieving, updating, and deleting to-do items.

**Requirements**

Go 1.16 or higher
MySQL

**Items**

- Create a new item.
POST /v1/items

- List all items.
GET /v1/items

- Get a specific item by ID. 
GET /v1/items/:id

- Update a specific item by ID. 
PATCH /v1/items/:id

- Delete a specific item by ID.
DELETE /v1/items/:id 


**Middleware**

This project uses custom middleware:

Recovery: A middleware for recovering from panics and returning a 500 status code.

**Error Handling**

In case of a panic in the application, the custom recovery middleware will handle the panic and return a 500 status code.
