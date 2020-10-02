package database

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Subscriptor required information
type Subscriptor struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Topic     string `json:"topic"`
	Country   string `json:"country"`
}

// CreateSubscriptor This function creates a new subscriptor in the database.
func CreateSubscriptor(ctx *gin.Context) {

	var subscriptor Subscriptor
	ctx.BindJSON(&subscriptor)

	firstname := subscriptor.FirstName
	lastname := subscriptor.LastName
	email := subscriptor.Email
	topic := subscriptor.Topic
	country := subscriptor.Country

	subscriptorNew := Subscriptor{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Topic:     topic,
		Country:   country,
	}

	_, err := collection.InsertOne(context.TODO(), subscriptorNew)

	if err != nil {
		log.Printf("[ CreateSubscriptor : InsertionError]: Problems inserting the subscriptor. %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error inserting subscriptor in the database.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusCreated,
		"message": "Subscritpor created successfully.",
	})
	return
}

// GetAllSubscriptors obtains all the subscriptors of the database.
func GetAllSubscriptors(ctx *gin.Context) {
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("[ GetAllSubscriptors : CollectingError]: Error connecting to the database to obtain all the subscriptions. %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error getting the subscriptors.",
		})
		return
	}

	subscriptors := []Subscriptor{}
	for cursor.Next(context.TODO()) {
		var subscriptor Subscriptor
		cursor.Decode(&subscriptor)
		subscriptors = append(subscriptors, subscriptor)

	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   subscriptors,
	})
	return
}

// GetSubscriptor This function obtains a subscriptor from the ID.
func GetSubscriptor(ctx *gin.Context) {
	id := ctx.Param("id")
	var subscriptor Subscriptor
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&subscriptor)

	if err != nil {
		log.Printf("[ GetSubscriptor : CollectingError]: Error finding the subscriptor %s in the database. %v ", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error getting the subscriptor.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   subscriptor,
	})
}

// ModifySubscritor Update the subscriptor information.
func ModifySubscritor(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedInfo Subscriptor
	ctx.BindJSON(&updatedInfo)

	subscriptionInfo := bson.M{
		"$set": bson.M{
			"firstname": updatedInfo.FirstName,
			"lastname":  updatedInfo.LastName,
			"topic":     updatedInfo.Topic,
			"email":     updatedInfo.Email,
			"country":   updatedInfo.Country,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": id}, subscriptionInfo)

	if err != nil {
		log.Printf("[ ModifySubscritor : UpdatingError]: Error updating the subscriptor information. %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error updating the subscriptor information.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "The subscritor information has been updated.",
	})

}

// DeleteSubscritor delete the specified subscritor information.
func DeleteSubscritor(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Printf("[ DeleteSubscritor ]: Error deleting the subscriptor information. %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error deleting the subscriptor information.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "The subscritor information has been removed.",
	})
}

func subscriptorsCollection(db *mongo.Database) {
	collection = db.Collection("subscriptors")
}

// Connection this function helps to connect to mongo
func Connection() {
	// Database connection
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Databse connected...")

	db := client.Database("subscriptions")
	subscriptorsCollection(db)
	return
}
