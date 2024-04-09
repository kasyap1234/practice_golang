package database

import (
	"context"
	"log"
	"time"

	"github.com/kasyap1234/practice_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func InitDatabase() {
	clientOptions := options.Client().ApplyURI("")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("library").Collection("books")

}
func GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	// cursor is like a pointer to the result set . initially points before the first document in the result set. the cursor advances tto the  next doument with cursor.next()
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err

	}
	// proper exiting and resource management while exiting the function
	defer cursor.Close(context.Background())

	// the reason why this cursor approach is used is imagine if the result set from 'collection.Find()' were very large  ,directly loading
	// all these documents into memory at once could be memory intensive and inefficient .
	// Instead, By using a loop with 'cursor.Next()'ou can fetch and process documents one by one or in batches, keeping memory usage under control.
	for cursor.Next(context.Background()) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)

	}
	return books, nil

}
func GetBookById(id string) (*models.Book, error) {
	var book models.Book
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err

	}
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil

}
func CreateBook(book *models.Book) error {
	book.ID = primitive.NewObjectID()
	book.PublishedAT = time.Now()

	_, err := collection.InsertOne(context.Background(), book)
	return err

}
func UpdateBook(id string, updatedBook *models.Book) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedBook.Title,
			"author":      updatedBook.Author,
			"publishedAt": time.Now(),
		},
	}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	return err

}
func DeleteBook(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err

}
