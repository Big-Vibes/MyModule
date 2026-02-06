package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	model "github.com/Big-vibes/mongoapi/Model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson" // Fixed: removed /v2 for consistency
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"         // Fixed: removed /v2 for consistency
	"go.mongodb.org/mongo-driver/mongo/options" // Fixed: removed /v2 for consistency
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const dbname = "netflix"
const colname = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB
var connectionString string

func init() {

	connectionString = os.Getenv("MONGODB_URI")
	if connectionString == "" {
		log.Fatal("MONGODB_URI env variable not set")
	}

	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbname).Collection(colname)

	//COLLECTION INSTANCE
	fmt.Println("Collection instance is ready")

}

// This is where MONGODB Helpers - file

// insert 1 record
func insert1Movies(movie model.NetflixEX) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// Update 1 Record
func update1movie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete 1 record
func delete1movie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie Got Deleted with Delete Count: ", deletecount)
}

// delete all record from mongoDB
func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Numbers of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// To Get All Movies in Mongodb
func getAllMovies() []bson.M { // Fixed: changed from []primitive.M to []bson.M to match appended items
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []bson.M // Fixed: changed from []primitive.M to []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err = cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

//you can do other type of functions

// Actual controllers - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.NetflixEX
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insert1Movies(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	update1movie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	delete1movie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
