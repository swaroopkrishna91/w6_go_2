# CRUD API for Car Database

This is a simple API that allows you to perform CRUD (Create, Read, Update, Delete) operations on a car database.

You can create, update, delete, and retrieve car data using POST or CURL commands.

## Features

- **Create**: Add a new car to the database.
- **Read**: Retrieve car information from the database.
- **Update**: Modify existing car details.
- **Delete**: Remove a car from the database.

## How to Run

You can run this API using POST or CURL commands.

## API Endpoints

#### Create a New Car (POST)

```
curl -X POST http://localhost:8497/cars/ -d '{"id":1, "make":"Toyota", "model":"Camry", "year":2020, "status":"available"}' -H "Content-Type: application/json"
```

#### Retrieve All Cars (GET)

```
curl -X GET http://localhost:8497/cars/
```

#### Update a Car (PUT)

```
curl -X PUT http://localhost:8497/cars/1 -d '{"make":"Toyota", "model":"Corolla", "year":2021, "status":"sold"}' -H "Content-Type: application/json"
```

#### Delete a Car (DELETE)

```
curl -X DELETE http://localhost:8497/cars/1
```
