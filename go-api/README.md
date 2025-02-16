# Go API Example with Clean Architecture

This is an example of a Go API implementing clean architecture with separate layers for database, repository, usecase and controllers.

## Project Structure

├── cmd
│ └── main.go
├── controller
│ └── product_controller.go
├── repository
│ └── product_repository.go
├── usecase
│ └── product_usecase.go
├── model
│ └── product.go
├── db
│ └── postgres.go
└── go.mod

## Layer Descriptions

### Database Layer (pkg/db)

Handles database connections and configurations. Example using PostgreSQL:

// pkg/db/postgres.go
package db

import (
"database/sql"
\_ "github.com/lib/pq"
)

func NewPostgresConnection(connStr string) (\*sql.DB, error) {
return sql.Open("postgres", connStr)
}

### Repository Layer (internal/repository)

Handles data persistence operations:

// internal/repository/product_repository.go
package repository

type ProductRepository interface {
GetByID(id int) (*model.Product, error)
Create(product *model.Product) error
Update(product \*model.Product) error
Delete(id int) error
}

### Usecase Layer (internal/usecase)

Contains business logic:

// internal/usecase/product_usecase.go
package usecase

type ProductUsecase interface {
GetProduct(id int) (*model.Product, error)
CreateProduct(product *model.Product) error
UpdateProduct(product \*model.Product) error
DeleteProduct(id int) error
}

### Controller Layer (internal/controller)

Handles HTTP requests and responses:

// internal/controller/product_controller.go
package controller

type ProductController struct {
productUsecase usecase.ProductUsecase
}

func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
// Handle HTTP request
}

## Usage Example

// cmd/main.go
package main

func main() {
// Initialize DB
db, err := postgres.NewPostgresConnection("postgres://user:pass@localhost:5432/dbname")
if err != nil {
log.Fatal(err)
}

    // Initialize layers
    productRepo := repository.NewProductRepository(db)
    productUsecase := usecase.NewProductUsecase(productRepo)
    productController := controller.NewProductController(productUsecase)

    // Setup HTTP routes
    http.HandleFunc("/products", productController.GetProduct)
    http.ListenAndServe(":8080", nil)

}

## Installation

go mod init myapi
go get -u github.com/lib/pq

## Running the Application

go run cmd/main.go

This example demonstrates a clean architecture approach with clear separation of concerns between layers. Each layer has its own responsibility:

- Database: Handles data storage
- Repository: Abstracts data access
- Usecase: Contains business logic
- Controller: Handles HTTP interactions

The layers are independent and communicate through interfaces, making the code modular and testable.
