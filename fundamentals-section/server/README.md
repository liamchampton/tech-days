# Sever setup

## Start the API server

```bash
go run main.go
```

## Routes

### GET /person/getall - Get all people

Call the API with curl:

```bash
curl -X GET http://localhost:8080/person/getall
```

Returns:

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

### POST /person/create - Create a person

Call the API with curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Adelina Simion",
    "fact": "Likes coding",
    "location": "London"
}' localhost:8080/person/create
```

### POST /person/delete - Delete a person

Call the API with curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "id": "6408c132091d45f5fa24f0f5"
}' localhost:8080/person/delete
```

