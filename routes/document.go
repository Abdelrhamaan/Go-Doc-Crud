package routes

import (
	"github.com/Abdelrahmaan/DocCrud/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




func RegisterDocumentRoutes(r *gin.Engine, db *gorm.DB){
	docs := r.Group("/documents")
	{
		docs.POST("/create", handlers.CreateDocument(db))
		docs.GET("/", handlers.ListDocuments(db))
		docs.GET("/:id", handlers.GetDocument(db))
		docs.PUT("/update/:id", handlers.UpdateDocument(db))
		docs.DELETE("/delete/:id", handlers.DeleteDocument(db))
	}
}