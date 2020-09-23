package main

import (
	"time"
)

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Order struct {
	HotelID        string   `firestore:"hotelid,omitempty" json:"hotelid,omitempty"`
	HotelName      string   `firestore:"hotelname,omitempty" json:"hotelname,omitempty"`
	HotelNameLC    string   `firestore:"hotelnamelc,omitempty" json:"hotelnamelc,omitempty"`
	HotelNameArray []string `firestore:"hotelnamearray,omitempty" json:"hotelnamearray,omitempty"`
	CheckInDate    string   `firestore:"checkindatetime,omitempty" json:"checkindatetime,omitempty"`
	CheckOutDate   string   `firestore:"checkoutdatetime,omitempty" json:"checkoutdatetime,omitempty"`
	CustomerName   string   `firestore:"customername,omitempty" json:"customername,omitempty"`
	CustomerEmail  string   `firestore:"customeremail,omitempty" json:"customeremail,omitempty"`
	CustomerTel    string   `firestore:"customertel,omitempty" json:"customertel,omitempty"`
	RoomID         string   `firestore:"roomid,omitempty" json:"roomid,omitempty"`
	RoomName       string   `firestore:"roomname,omitempty" json:"roomname,omitempty"`
	MaxGuests      int      `firestore:"maxguests,omitempty" json:"maxguests,omitempty"`
	AmountToPay    int      `firestore:"amounttopay,omitempty" json:"amounttopay,omitempty"`
	PaymentID      string   `firestore:"paymentid,omitempty" json:"paymentid,omitempty"`
}

type HotelSearchQuery struct {
	Expression string `json:"expression,omitempty"`
	Key        string `json:"key,omitempty"`
	MaxResult  int    `json:"maxresult,omitempty"`
	Order      string `json:"order,omitempty"`
}

type Todos []Todo

type Orders []Order
