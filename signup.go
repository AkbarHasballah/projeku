package projeku

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func initMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var payload SignupPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Gagal mendekode payload JSON", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Gagal mengenkripsi kata sandi: %v", err)
		http.Error(w, "Gagal mengenkripsi kata sandi", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Pendaftaran berhasil"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	collection := client.Database("Isidengan nama database anda").Collection("isi dengan nama collection anda")
	_, err = collection.InsertOne(context.Background(), bson.M{
		"username": payload.Username,
		"password": string(hashedPassword),
	})
	if err != nil {
		log.Printf("Gagal menyimpan data ke MongoDB: %v", err)
		http.Error(w, "Gagal menyimpan data ke MongoDB", http.StatusInternalServerError)
		return
	}
}
