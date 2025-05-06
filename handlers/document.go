package handlers

import (
	"net/http"

	"github.com/Abdelrahmaan/DocCrud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




func CreateDocument(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context)  {
		var doc models.Document
		if err:= c.ShouldBindJSON(&doc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&doc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		 }
		c.JSON(http.StatusCreated, doc)
	}
}


func ListDocuments(db *gorm.DB) gin.HandlerFunc {
	return func (c *gin.Context)  {
		var docs []models.Document
		if err:=db.Find(&docs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, docs)
	}
}
func GetDocument(db *gorm.DB) gin.HandlerFunc {
	return func (c *gin.Context){
		var doc models.Document
		id := c.Param("id")
		if err:= db.First(&doc, id).Error; err!=nil{
			if err == gorm.ErrRecordNotFound{
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			}else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, doc)
	}
}

func DeleteDocument(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		var doc models.Document
		if err := db.Delete(&doc, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound{
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			}else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusNoContent, doc)
	}
}

func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var doc models.Document
        if err := db.First(&doc, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
            return
        }
        var input models.Document
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        doc.Title = input.Title
        doc.Author = input.Author
        doc.Content = input.Content
        db.Save(&doc)
        c.JSON(http.StatusOK, doc)
    }
}

// func GetDocument(db *gorm.DB) gin.HandlerFunc  {
	
// }

// func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
	
// }

// func DeleteDocument(db *gorm.DB) gin.HandlerFunc {
	
// }


// ## Gin + GORM + PostgreSQL CRUD System

// This is a step-by-step guide to build a Documenting System API in Go using Gin, GORM, and PostgreSQL (in Docker). Each document has: **ID**, **Title**, **Author**, **Content**, **CreatedAt**, **UpdatedAt**.

// ---

// ### 📁 Project Structure

// ```
// gin-gorm-postgres/
// ├── docker-compose.yml
// ├── Dockerfile
// ├── go.mod
// ├── go.sum
// ├── main.go
// └── models/
//     └── document.go
// └── handlers/
//     └── document.go
// └── routes/
//     └── document.go
// ```

// ---

// ### 1. Docker Compose & Dockerfile

// **docker-compose.yml**:

// ```yaml
// version: '3.8'
// services:
//   db:
//     image: postgres:13
//     restart: always
//     environment:
//       POSTGRES_USER: user
//       POSTGRES_PASSWORD: pass
//       POSTGRES_DB: docsdb
//     ports:
//       - "5432:5432"
//     volumes:
//       - pgdata:/var/lib/postgresql/data

//   app:
//     build: .
//     ports:
//       - "8080:8080"
//     depends_on:
//       - db
//     environment:
//       DB_HOST: db
//       DB_USER: user
//       DB_PASS: pass
//       DB_NAME: docsdb
//       DB_PORT: 5432
//       DB_SSLMODE: disable

// volumes:
//   pgdata:
// ```

// **Dockerfile**:

// ```dockerfile
// FROM golang:1.20-alpine AS builder
// WORKDIR /app
// COPY go.mod go.sum ./
// RUN go mod download
// COPY . .
// RUN go build -o server .

// FROM alpine:latest
// WORKDIR /root/
// COPY --from=builder /app/server .
// EXPOSE 8080
// CMD ["./server"]
// ```

// ---

// ### 2. Initialize Go Module

// ```bash
// go mod init github.com/yourusername/gin-gorm-postgres
// go get github.com/gin-gonic/gin
// go get gorm.io/gorm
// go get gorm.io/driver/postgres
// ```

// ---

// ### 3. Define the Model

// **models/document.go**:

// ````go
// package models

// import (
//     "time"
//     "gorm.io/gorm"
// )

// // Document represents a document record
// // GORM tags specify column names and auto timestamps

// ```go
//  type Document struct {
//      ID        uint           `gorm:"primaryKey" json:"id"`
//      Title     string         `gorm:"type:varchar(100)" json:"title"`
//      Author    string         `gorm:"type:varchar(100)" json:"author"`
//      Content   string         `gorm:"type:text" json:"content"`
//      CreatedAt time.Time      `json:"created_at"`
//      UpdatedAt time.Time      `json:"updated_at"`
//      DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
//  }
// ````

// ````

// ---

// ### 4. Database Initialization

// **main.go** (partial):
// ```go
// package main

// import (
//     "fmt"
//     "log"
//     "os"
//     "github.com/gin-gonic/gin"
//     "gorm.io/driver/postgres"
//     "gorm.io/gorm"
//     "github.com/yourusername/gin-gorm-postgres/models"
// )

// func initDB() *gorm.DB {
//     dsn := fmt.Sprintf(
//         "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
//         os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
//         os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"),
//     )
//     db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatal("Failed to connect to database:", err)
//     }
//     // Auto-migrate the Document model
//     db.AutoMigrate(&models.Document{})
//     return db
// }
// ````

// ---

// ### 5. Handlers (CRUD Operations)

// **handlers/document.go**:

// ```go
// package handlers

// import (
//     "net/http"
//     "github.com/gin-gonic/gin"
//     "gorm.io/gorm"
//     "github.com/yourusername/gin-gorm-postgres/models"
// )

// // Create a new document
// func CreateDocument(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         var doc models.Document
//         if err := c.ShouldBindJSON(&doc); err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//             return
//         }
//         if err := db.Create(&doc).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//         c.JSON(http.StatusCreated, doc)
//     }
// }

// // Get document by ID
// func GetDocument(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         id := c.Param("id")
//         var doc models.Document
//         if err := db.First(&doc, id).Error; err != nil {
//             if err == gorm.ErrRecordNotFound {
//                 c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
//             } else {
//                 c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             }
//             return
//         }
//         c.JSON(http.StatusOK, doc)
//     }
// }

// // Update document by ID
// func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         id := c.Param("id")
//         var doc models.Document
//         if err := db.First(&doc, id).Error; err != nil {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
//             return
//         }
//         var input models.Document
//         if err := c.ShouldBindJSON(&input); err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//             return
//         }
//         doc.Title = input.Title
//         doc.Author = input.Author
//         doc.Content = input.Content
//         db.Save(&doc)
//         c.JSON(http.StatusOK, doc)
//     }
// }

// // Delete document by ID
// func DeleteDocument(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         id := c.Param("id")
//         if err := db.Delete(&models.Document{}, id).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//         c.Status(http.StatusNoContent)
//     }
// }

// // Get all documents
// func ListDocuments(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         var docs []models.Document
//         if err := db.Find(&docs).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//         c.JSON(http.StatusOK, docs)
//     }
// }
// ```

// ---

// ### 6. Routes

// **routes/document.go**:

// ```go
// package routes

// import (
//     "github.com/gin-gonic/gin"
//     "gorm.io/gorm"
//     "github.com/yourusername/gin-gorm-postgres/handlers"
// )

// func RegisterDocumentRoutes(r *gin.Engine, db *gorm.DB) {
//     docs := r.Group("/documents")
//     {
//         docs.POST("/", handlers.CreateDocument(db))
//         docs.GET("/", handlers.ListDocuments(db))
//         docs.GET("/:id", handlers.GetDocument(db))
//         docs.PUT("/:id", handlers.UpdateDocument(db))
//         docs.DELETE("/:id", handlers.DeleteDocument(db))
//     }
// }
// ```

// ---

// ### 7. Wire Everything in `main.go`

// ```go
// package main

// import (
//     "github.com/gin-gonic/gin"
//     "github.com/yourusername/gin-gorm-postgres/routes"
// )

// func main() {
//     db := initDB()
//     r := gin.Default()

//     routes.RegisterDocumentRoutes(r, db)

//     r.Run(":8080")
// }
// ```

// ---

// ### 8. Run with Docker Compose

// ```bash
// docker-compose up --build
// ```

// * **App** available at `http://localhost:8080`
// * **Endpoints**:

//   * `POST    /documents/`        → Create
//   * `GET     /documents/`        → List
//   * `GET     /documents/:id`     → Get by ID
//   * `PUT     /documents/:id`     → Update
//   * `DELETE  /documents/:id`     → Delete

// ---

// 🎉 You now have a full CRUD API with Gin, GORM, and PostgreSQL in Docker!
