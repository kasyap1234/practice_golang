package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/practice_golang/database"
	"github.com/kasyap1234/practice_golang/handlers"
)
func main(){
	database.InitiDatabase()
	r:=gin.New();
	r.GET("/api/books",handlers.GetBooksHandler);
	r.GET("/api/books/:id",handlers.GetBookByIDHandler)
	r.POST("/api/books",handlers.CreateBookHandler);
	r.PUT("/api/books/:id",handlers.UpdateBookHandler);
	r.DELETE("/api/books/:id",handlers.DeleteBookHandler);
	log.Println("starting to run server  on port 3000");
	r.Run(":3000");

}
