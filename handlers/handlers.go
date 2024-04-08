package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/practice_golang/database"
	"github.com/kasyap1234/practice_golang/models"
)

func GetBooksHandler(c *gin.Context) {
	books, err := database.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)

}
func GetBooksByIDHandler(c *gin.Context) {
	id := c.Param("id")
	book, err := database.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not found"})

		return
	}
	c.JSON(http.StatusOK, book)
}
func CreateBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
		return
	}
	if err := database.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H("error", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, book)

}
func UpdateBookHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.UpdateBook(id, &updatedBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book UPdated successfully "})
	return

}
func DeleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	if err := database.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book Deleted Successfully "})
}
