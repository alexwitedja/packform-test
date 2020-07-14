## Backend setup

# Endpoints

/api/orders which returns json:
```
{
    "_id": "1",
    "company_name": "Roga & Kopyta",
    "customer_name": "Ivan Ivanovich",
    "order_date": "2020-01-02T15:34:12Z",
    "delivered_amount": 6.73,
    "total_amount": 918.13
}
```

# Setup

1. Install all dependencies:
```
go get -u github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get go.mongodb.org/mongo-driver
go get github.com/lib/pq
```

2. go to pghelper/pghelper.go change the connection string to suit your environment
3. go to mongohelper/mongohelper.go and do the same.
4. run the server `go run backend/main.go`.
5. endpoint will be served in localhost:9999

## Frontend setup