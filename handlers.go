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
	firebase "firebase.google.com/go"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

////  ORDER /////
// CreateOrder godoc
// @Summary Create New Order
// @Description To receive a new Order and insert it in firestore
// @Tags order
// @Accept  json
// @Param data body Order true "The input todo struct"
// @Produce  json
// @Success 200 {object} HTTPResponse
// @Router /order [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get a Firestore client.
	ctx := context.Background()
	sa := option.WithCredentialsFile("./avab011-firebase-adminsdk-22i2q-02404661c5.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	hotelid := guuid.New().String()
	roomid := guuid.New().String()
	paymentid := guuid.New().String()
	docid := guuid.New().String()

	//post data

	newOrder := new(Order)
	newPayment := new(Payment)
	resp := new(HTTPResponse)
	_ = json.NewDecoder(r.Body).Decode(newOrder)

	hotelNameBreakOut := strings.Fields(strings.ToLower(newOrder.HotelName))
	hotelNameBreakOut = append(hotelNameBreakOut, strings.ToLower(newOrder.HotelName))

	newOrder.HotelNameArray = hotelNameBreakOut
	newOrder.HotelID = hotelid
	newOrder.RoomID = roomid
	newOrder.PaymentID = paymentid
	newOrder.HotelNameLC = strings.ToLower(newOrder.HotelName)

	_, err = client.Collection("orders").Doc(docid).Set(ctx, newOrder)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		fmt.Fprintf(w, "An error has occurred: %s", err)
	}

	newPayment.OrderID = docid
	newPayment.Paid = false

	_, err = client.Collection("payment").Doc(paymentid).Set(ctx, newPayment)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		fmt.Fprintf(w, "An error has occurred: %s", err)
	}

	resp.Success = true
	resp.Msg = docid

	e, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(e))

}

// SearchOrderByHotel godoc
// @Summary Search Order by Hotel Name
// @Description To search orders by hotel name from firestore
// @Tags order
// @Accept  json
// @Param data body HotelSearchQuery true "The input search struct"
// @Produce  json
// @Success 200 {array} Order
// @Router /order/hotel [post]
func searchOrderByHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get a Firestore client.
	ctx := context.Background()
	sa := option.WithCredentialsFile("./avab011-firebase-adminsdk-22i2q-02404661c5.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

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

// SearchOrderByCustomer godoc
// @Summary Search order by Customer Information (tel, email, name)
// @Description To search orders by customer info (tel, email, name) from firestore
// @Tags order
// @Accept  json
// @Param data body HotelSearchQuery true "The input search struct"
// @Produce  json
// @Success 200 {array} Order
// @Router /order/customer [post]
func searchOrderByCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get a Firestore client.
	ctx := context.Background()
	sa := option.WithCredentialsFile("./avab011-firebase-adminsdk-22i2q-02404661c5.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

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
// GetPaymentStatus godoc
// @Summary Get Payment Status of an Order
// @Description To get the status of Payment
// @Tags payment
// @Accept  json
// @Param orderid path string true "Order ID"
// @Produce  json
// @Success 200 {object} Payment
// @Router /paymentstatus/{orderid} [get]
func getPaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderid"]
	payment := new(Payment)

	// Get a Firestore client.
	ctx := context.Background()
	sa := option.WithCredentialsFile("./avab011-firebase-adminsdk-22i2q-02404661c5.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	if len(orderID) > 35 {
		//find payment status
		queryResult := client.Collection("payment").Where("orderid", "==", orderID).Documents(ctx)
		for {
			doc, err := queryResult.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Fprintln(w, "ERROR")
			}

			err = mapstructure.Decode(doc.Data(), payment)
			if err != nil {
				// error
			}
			payment.PaymentID = doc.Ref.ID

		}
		e, err := json.Marshal(payment)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(e))
	} else {
		fmt.Fprintf(w, "")
	}

}

// MakePayment godoc
// @Summary Make Payment of an Order
// @Description To pay the order if not yet paid
// @Tags payment
// @Accept  json
// @Param paymentid path string true "Payment ID"
// @Produce  json
// @Success 200 {object} Order
// @Router /makepayment/{paymentid} [get]
func makePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["paymentid"]
	resp := new(HTTPResponse)

	// Get a Firestore client.
	ctx := context.Background()
	sa := option.WithCredentialsFile("./avab011-firebase-adminsdk-22i2q-02404661c5.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	if len(paymentID) > 35 {
		_, err := client.Collection("payment").Doc(paymentID).Set(ctx, map[string]interface{}{
			"paid": true,
		}, firestore.MergeAll)
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}
		resp.Success = true
		resp.Msg = "Payment Success"

	} else {
		resp.Success = false
		resp.Msg = "Payment Failed"
	}

	e, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(e))
}
