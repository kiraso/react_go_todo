package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kiraso/react_go_todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/kiraso/react_go_todo/"
)

// 					  / \_______ /|_\             \/
// 					 /          /_/ \__
// 					/             \_/ /
// 				  _|_              |/|_
// 				  _|_  O    _    O  _|_
// 				  _|_      (_)      _|_
// 				   \                 /
// 					_\_____________/_
// 				   /  \/  (___)  \/  \
// 				   \__(  o     o  )__/
var collection *mongo.Collection
func init(){
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error .env")
	}

}

func createDBInstance() {

	connectionString := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(collName)
}
func GetAllTask(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
   w.Header().Set("Access-Control-Allow-Origin", "*")
   payload := getAllTasks()
   //reponse กลับ
   json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	json.NewDecoder(r.Body).Decode(&task)
	//collection.InsertOne(context.Background(), task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TasksComplete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//addType
	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func UndoTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	params := mux.Vars(r)
	undoTasks(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	count :=  deleteAll()
	json.NewEncoder(w).Encode(count) 


}

//function for connect Database 

func getAllTasks() []primitive.M{
	cur,err := collection.Find(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e :=cur.Decode(&result)
		if e != nil {	
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func insertOneTask(task models.ToDoList){
	//gen id key auto 
	insertResult,err := collection.InsertOne(context.Background(),task)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}
func taskComplete(task string){
	//dont care err?
	id,_ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"id":id}
	update := bson.M{"$set": bson.M{"status":true}}
	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("modified count",result.ModifiedCount)


}

func undoTasks(task string){
	id,_ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"id":id}
	update := bson.M{"$set": bson.M{"status":false}}
	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("modified count",result.ModifiedCount)


}

func deleteOneTask(task string){
	id,_ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"id":id}
	d,err := collection.DeleteOne(context.Background(),filter)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("deleted count",d.DeletedCount)
}

func deleteAll() int64{
	d,err := collection.DeleteMany(context.Background(),bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("deleted count",d.DeletedCount)
	return d.DeletedCount
}