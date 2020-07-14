package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/alexwitedja/packform-test-api/backend/models"
	"github.com/alexwitedja/packform-test-api/backend/mongohelper"
	"github.com/alexwitedja/packform-test-api/backend/pghelper"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OrderPayload to be sent after request.
type OrderPayload struct {
	ID              string  `json:"_id"`
	CompanyName     string  `json:"company_name"`
	CustomerName    string  `json:"customer_name"`
	OrderDate       string  `json:"order_date"`
	DeliveredAmount float64 `json:"delivered_amount"`
	TotalAmount     float64 `json:"total_amount"`
}

// OrderValue used in getOrderValue function.
type OrderValue struct {
	TotalAmount     float64
	DeliveredAmount float64
}

// Mongo Database instance
var mongodb *mongo.Database

// getOrders runs functions fetching data from db to construct a payload with only the necessary information.
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Create an array of orders
	var payloads []OrderPayload

	collection := mongodb.Collection("orders")
	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		getError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// Create a value into which the single document can be decoded
		var order models.Order
		// & character returns the memory address of the following variable.
		// Decode is similar to deserialize process.
		err := cur.Decode(&order)

		if err != nil {
			log.Fatal(err)
		}

		var customer models.Customer = getCustomer(order.CustomerID)
		var customerCompany models.CustomerCompany = getCustomerCompany(customer.CompanyID)
		orderValue := getOrderValue(order.ID)

		// add item our array
		payload := OrderPayload{
			ID:              order.ID,
			CompanyName:     customerCompany.CompanyName,
			CustomerName:    customer.Name,
			OrderDate:       order.CreatedAt,
			DeliveredAmount: math.Ceil(orderValue.DeliveredAmount*100) / 100,
			TotalAmount:     math.Ceil(orderValue.TotalAmount*100) / 100,
		}

		payloads = append(payloads, payload)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Similar to serialize process
	json.NewEncoder(w).Encode(payloads)
}

// Get customer object from mongodb based on customer ID of the order.
func getCustomer(customerID string) models.Customer {
	collection := mongodb.Collection("customers")

	filter := bson.M{"_id": customerID}

	var customer models.Customer

	err := collection.FindOne(context.TODO(), filter).Decode(&customer)
	if err != nil {
		log.Fatal(err)
	}

	return customer
}

// Get customer company object from mongodb based on customer ID of the order.
func getCustomerCompany(companyID string) models.CustomerCompany {
	collection := mongodb.Collection("customer_companies")

	filter := bson.M{"_id": companyID}

	var company models.CustomerCompany

	err := collection.FindOne(context.TODO(), filter).Decode(&company)
	if err != nil {
		log.Fatal(err)
	}

	return company
}

// pg db instance using gorm
var pgdb *gorm.DB

// Finds out the total amount of the order and the amount delivered.
func getOrderValue(orderID string) *OrderValue {

	var totalAmount float64
	var deliveredAmount float64

	var orderItems []models.OrderItem
	var deliveries []models.Delivery
	pgdb.Where("order_id = ?", orderID).Find(&orderItems)

	for _, orderItem := range orderItems {
		// an order contains more than one order item.
		totalAmount += orderItem.PricePerUnit * float64(orderItem.Quantity)

		// Calculate delivered amount
		pgdb.Where("order_item_id = ?", orderItem.Model.ID).Find(&deliveries)
		for _, delivery := range deliveries {
			// An order item can be delivered multiple times.
			deliveredAmount += orderItem.PricePerUnit * float64(delivery.DeliveredQuantity)
		}
	}

	orderValue := &OrderValue{
		DeliveredAmount: deliveredAmount,
		TotalAmount:     totalAmount,
	}

	return orderValue
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
func getError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

func main() {
	mongodb = mongohelper.ConnectDB()
	pgdb = pghelper.ConnectDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/orders", getOrders).Methods("GET")

	log.Fatal(http.ListenAndServe(":9999", r))
}
