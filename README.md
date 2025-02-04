# DB script setup

1. Install mongodb and postgresql.
2. Install python dependencies `pip3 install pymongo` and `pip3 install psycopg2`
3. Find keywords `connection string` in `db_script/init_db.py` and replace connection string.
4. Run the `init_db.py`

# Backend setup

## Endpoints

`GET` /api/orders which returns array of json:
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
`POST` /api/ordersBetween that ACCEPTS json body:
```
{
    "start": "2019-02-16T13:00:00.000Z",
    "end": "2020-02-23T13:00:00.000Z"
}
```
and in response sends array of json:
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
## Setup

1. Install all dependencies:
```
go get -u github.com/gorilla/mux
go get github.com/gorilla/handlers
go get github.com/jinzhu/gorm
go get go.mongodb.org/mongo-driver
go get github.com/lib/pq
```

2. go to pghelper/pghelper.go change the connection string to suit your environment
3. go to mongohelper/mongohelper.go and do the same.
4. run the server `go run backend/main.go`.
5. endpoint will be served in localhost:9999

# Frontend setup

1. `cd to frontend/packform-test` and do `npm install`
2. `npm run serve` open `localhost:9999`

### Note
There's a bug with the date range picker, filtering this way never
works on the first try but any instances after will work fine.