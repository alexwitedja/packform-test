# Script to initialise and populate db with csv files.
import os, csv, json

current_directory = os.path.dirname(os.path.abspath(__file__))
# Begin mongodb insert script
import pymongo

mongo_client = pymongo.MongoClient("mongodb://localhost:27017/") # Replace connection string with your own.
mongo_db = mongo_client["packformtest"]

customer_companies_filepath = os.path.join(current_directory, 'Mongo-customer_companies.csv')
customer_companies_col = mongo_db["customer_companies"]

with open(customer_companies_filepath) as csvFile:
    csvReader = csv.DictReader(csvFile)
    for rows in csvReader:
        rows['_id'] = rows.pop(next(iter(rows))) # Use first key as the unique id.
        customer_companies_col.insert_one(rows)

customers_filepath = os.path.join(current_directory, 'Mongo-customers.csv')
customers_col = mongo_db["customers"]

with open(customers_filepath) as csvFile:
    csvReader = csv.DictReader(csvFile)
    for rows in csvReader:
        rows['_id'] = rows.pop(next(iter(rows)))
        customers_col.insert_one(rows)

orders_filepath = os.path.join(current_directory, 'Mongo-Orders.csv')
orders_col = mongo_db["orders"]

with open(orders_filepath) as csvFile:
    csvReader = csv.DictReader(csvFile)
    for rows in csvReader:
        rows['_id'] = rows.pop(next(iter(rows)))
        orders_col.insert_one(rows)

# Begin postgresql insert script

import psycopg2
conn = psycopg2.connect("host=localhost dbname=postgres \
  user=postgres password=root port=5432") # Replace connection string with your own.
cur = conn.cursor()

# Create tables for the csv.
cur.execute("""CREATE TABLE IF NOT EXISTS orders(
    id INTEGER PRIMARY KEY,
    created_at DATE,
    order_name VARCHAR (100),
    customer_id VARCHAR (100)
  );

  CREATE TABLE IF NOT EXISTS order_items(
    id INTEGER PRIMARY KEY,
    order_id INTEGER,
    price_per_unit REAL,
    quantity INTEGER,
    product VARCHAR (100)
  );

  CREATE TABLE IF NOT EXISTS deliveries(
    id INTEGER PRIMARY KEY,
    order_item_id INTEGER,
    delivered_quantity INTEGER
  );
""")
conn.commit()

orders_filepath = os.path.join(current_directory, 'Postgres-orders.csv')

with open(orders_filepath, 'r') as csvFile:
  csvReader = csv.reader(csvFile)
  next(csvReader) # Skipping header row

  for row in csvReader:
    cur.execute(
      "INSERT INTO orders VALUES(%s, %s, %s, %s)",
      row
    )

order_items_filepath = os.path.join(current_directory, 'Postgres-order_items.csv')

with open(order_items_filepath, 'r') as csvFile:
  csvReader = csv.reader(csvFile)
  next(csvReader) # Skipping header row

  for row in csvReader:
    row = [None if x=='' else x for x in row] # Replace empty values with None (null in postresql)

    cur.execute(
      "INSERT INTO order_items VALUES(%s, %s, %s, %s, %s)",
      row
    )

deliveries_filepath = os.path.join(current_directory, 'Postgres-deliveries.csv')

with open(deliveries_filepath, 'r') as csvFile:
  csvReader = csv.reader(csvFile)
  next(csvReader) # Skipping header row

  for row in csvReader:
    cur.execute(
      "INSERT INTO deliveries VALUES(%s, %s, %s)",
      row
    )
conn.commit()
