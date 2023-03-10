# Sever setup

## Start the API server

```bash
go run main.go
```

## RESTful Routes

### GET /persons - Get all people

Call the API with curl:

```bash
curl -X GET http://localhost:8080/persons
```

Returns a slice of JSON objects:

```json
[
    {
        "id": "1234",
        "name": "Liam Hampton",
        "fact": "Likes F1",
        "location": "London"
    },
    {
        "id": "1234",
        "name": "Adelina Simion",
        "fact": "Likes coding",
        "location": "London"
    }
]
```

### POST /persons - Create a person

Call the API with curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Adelina Simion",
    "fact": "Likes coding",
    "location": "London"
}' localhost:8080/persons
```

### TODO: POST /person/{id} - Delete a person

Call the API with curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "id": "1234"
}' localhost:8080/person/{id}
```

