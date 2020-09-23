package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"cloud.google.com/go/firestore"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	iter := client.Collection("farmasi").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

////  ORDER /////
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	hotelid := guuid.New().String()
	roomid := guuid.New().String()
	paymentid := guuid.New().String()
	docid := guuid.New().String()

	//post data

	newOrder := new(Order)
	_ = json.NewDecoder(r.Body).Decode(newOrder)

	hotelNameBreakOut := strings.Fields(strings.ToLower(newOrder.HotelName))
	hotelNameBreakOut = append(hotelNameBreakOut, strings.ToLower(newOrder.HotelName))

	newOrder.HotelNameArray = hotelNameBreakOut
	newOrder.HotelID = hotelid
	newOrder.RoomID = roomid
	newOrder.PaymentID = paymentid
	newOrder.HotelNameLC = strings.ToLower(newOrder.HotelName)

	_, err := client.Collection("orders").Doc(docid).Set(ctx, newOrder)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		fmt.Fprintf(w, "An error has occurred: %s", err)
	}
	fmt.Fprintln(w, "Order Created")

}

func searchOrderByHotel(w http.ResponseWriter, r *http.Request) {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	w.Header().Set("Content-Type", "application/json")
	query := new(HotelSearchQuery)
	_ = json.NewDecoder(r.Body).Decode(query)

	hasSinglePercentage, _ := regexp.MatchString("^%$", query.Expression)

	orders := []*Order{}
	queryResult := new(firestore.DocumentIterator)

	if hasSinglePercentage {
		//find all
		queryResult = client.Collection("orders").Documents(ctx)
	} else {
		//find exact
		queryResult = client.Collection("orders").Where("hotelnamearray", "array-contains", query.Expression).Documents(ctx)

	}

	//run search query

	order := new(Order)

	for {
		doc, err := queryResult.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintln(w, "ERROR")
		}

		err = mapstructure.Decode(doc.Data(), order)
		if err != nil {
			// error
		}
		orders = append(orders, order)
	}
	e, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(e))
}

func searchOrderByCustomer(w http.ResponseWriter, r *http.Request) {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	w.Header().Set("Content-Type", "application/json")
	query := new(HotelSearchQuery)
	_ = json.NewDecoder(r.Body).Decode(query)

	hasSinglePercentage, _ := regexp.MatchString("^%$", query.Expression)

	orders := []*Order{}
	queryResult := new(firestore.DocumentIterator)

	if hasSinglePercentage {
		//find all
		queryResult = client.Collection("orders").Documents(ctx)
	} else {
		//find exact
		queryResult = client.Collection("orders").Where(query.Key, "==", query.Expression).Documents(ctx)

	}

	//run search query

	order := new(Order)

	for {
		doc, err := queryResult.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintln(w, "ERROR")
		}

		err = mapstructure.Decode(doc.Data(), order)
		if err != nil {
			// error
		}
		orders = append(orders, order)
	}
	e, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(e))
}

////  PAYMENT  /////
func getPaymentStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getPaymentStatus")
}

func makePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["paymentid"]
	fmt.Fprintln(w, "makePayment:", paymentID)
}
