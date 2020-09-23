package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
)

// @title AVA API
// @version 1.0
// @description Serice for managing hotel orders
// @termsOfService http://swagger.io/terms/
// @contact.name Zarul Zakuan
// @contact.email zarulzakuan@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	router := NewRouter()
	log.Println("Running server...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "avab011"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
