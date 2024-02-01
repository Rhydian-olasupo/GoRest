package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sender struct {
	ID         interface{} `json:"id" bson:"_id,omitempty"`
	First_Name string      `json:"first_name" bson:"first_name"`
	Last_Name  string      `json:"last_name" bson:"last_name"`
	Address    Address     `json:"address" bson:"address"`
	Phone      string      `json:"phone" bson:"phone"`
}

type Address struct {
	Type    string `json:"type" bson:"type"`
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Pincode uint64 `json:"pincode" bson:"pincode"`
	Country string `json:"country" bson:"country"`
}

type Receiver struct {
	ID         interface{} `json:"id" bson:"_id,omitempty"`
	First_Name string      `json:"first_name" bson:"first_name"`
	Last_Name  string      `json:"last_name" bson:"last_name"`
	Address    Address     `json:"address" bson:"address"`
	Phone      string      `json:"phone" bson:"phone"`
}

type Package struct {
	ID         interface{} `json:"id" bson:"_id,omitempty"`
	Dimensions Dimensions  `json:"dimensions" bson:"dimensions"`
	Weight     int         `json:"weight" bson:"weight"`
	Damaged    bool        `json:"is_damaged" bson:"is_damaged"`
	Status     string      `json:"status" bson:"status"`
}

type Dimensions struct {
	Width  int `json:"width" bson:"width"`
	Height int `json:"height" bson:"width"`
}

type Payment struct {
	ID             interface{}    `json:"id" bson:"_id,omitempty"`
	InitiatedOn    time.Time      `json:"initiated_on" bson:"initiated_on"`
	SuccessfulOn   time.Time      `json:"successfull_on" bson:"successful_on"`
	MerchantID     int            `json:"merchant_id" bson:"merchant_id"`
	ModeOfPayment  string         `json:"mode_of_payment" bson:"mode_of_payment"`
	PaymentDetails PaymentDetails `json:"payment_details" bson:"payment_details"`
}

type PaymentDetails struct {
	TransactionToken string `json:"transaction_token" bson:"transaction_token"`
}

type Carrier struct {
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	CarrierCode int         `json:"carrier_code" bson:"carrier_code"`
	IsPartner   bool        `json:"is_partner" bson:"is_partner"`
}

type Shipment struct {
	ID         interface{} `json:"id" bson:"_id,omitempty"`
	Sender     Sender      `json:"sender" bson:"sender"`
	Receiver   Receiver    `json:"receiver" bson:"receiver"`
	Package    Package     `json:"package" bson:"package"`
	Payment    Payment     `json:"payment" bson:"payment"`
	Carrier    Carrier     `json:"carrier" bson:"carrier"`
	PromisedOn time.Time   `json:"promised_on" bson:"promised_on"`
}

type DB struct {
	collection *mongo.Collection
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	collection := client.Database("appDB").Collection("logistics")
	db := &DB{collection: collection}

	r := mux.NewRouter()

	r.HandleFunc("/v1/logistics/user/sender/{id:[a-zA-Z0-9]*}", db.GetSender).Methods("GET")
	r.HandleFunc("/v1/logistics/user/sender", db.PostSender).Methods("POST")
	r.HandleFunc("/v1/logistics/user/receiver/{id:[a-zA-Z0-9]*}", db.GetReceiver).Methods("GET")
	r.HandleFunc("/v1/logistics/user/receiver", db.PostReceiver).Methods("POST")
	//r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	//r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

//GetSender fetches the details of the Sender with a given ID

func (db *DB) GetSender(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var sender Sender

	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}
	err := db.collection.FindOne(context.TODO(), filter).Decode(&sender)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(sender)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

//GetReceiver fetches the details of the Receiver with a given ID

func (db *DB) GetReceiver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var receiver Receiver

	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}
	err := db.collection.FindOne(context.TODO(), filter).Decode(&receiver)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(receiver)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// PostSender Create the Sender details in our database
func (db *DB) PostSender(w http.ResponseWriter, r *http.Request) {
	var sender Sender
	postBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(postBody, &sender)
	result, err := db.collection.InsertOne(context.TODO(), sender)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (db *DB) PostReceiver(w http.ResponseWriter, r *http.Request) {
	var receiver Receiver
	postBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(postBody, &receiver)
	result, err := db.collection.InsertOne(context.TODO(), receiver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
