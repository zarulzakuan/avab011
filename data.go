package main

type Order struct {
	HotelID        string   `firestore:"hotelid,omitempty" json:"-"`
	HotelName      string   `firestore:"hotelname,omitempty" json:"hotelname,omitempty"`
	HotelNameLC    string   `firestore:"hotelnamelc,omitempty" json:"-"`
	HotelNameArray []string `firestore:"hotelnamearray,omitempty" json:"-"`
	CheckInDate    string   `firestore:"checkindatetime,omitempty" json:"checkindatetime,omitempty"`
	CheckOutDate   string   `firestore:"checkoutdatetime,omitempty" json:"checkoutdatetime,omitempty"`
	CustomerName   string   `firestore:"customername,omitempty" json:"customername,omitempty"`
	CustomerEmail  string   `firestore:"customeremail,omitempty" json:"customeremail,omitempty"`
	CustomerTel    string   `firestore:"customertel,omitempty" json:"customertel,omitempty"`
	RoomID         string   `firestore:"roomid,omitempty" json:"-"`
	RoomName       string   `firestore:"roomname,omitempty" json:"roomname,omitempty"`
	MaxGuests      int      `firestore:"maxguests,omitempty" json:"maxguests,omitempty"`
	AmountToPay    int      `firestore:"amounttopay,omitempty" json:"amounttopay,omitempty"`
	PaymentID      string   `firestore:"paymentid,omitempty" json:"-"`
}

type HotelSearchQuery struct {
	Expression string `json:"expression,omitempty"`
	Key        string `json:"key,omitempty"`
}

type Payment struct {
	PaymentID string `firestore:"paymentid,omitempty" json:"paymentid,omitempty"`
	OrderID   string `firestore:"orderid,omitempty" json:"orderid,omitempty"`
	Paid      bool   `firestore:"paid,omitempty" json:"paid,omitempty"`
}

type HTTPResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
}
