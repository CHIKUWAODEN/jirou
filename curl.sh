#!/bin/sh


# API Version1 Root
echo "GET /"
curl -i -L \
  -X GET \
  -H "Content-Type: application/json" \
  "localhost:8080/"
echo "exit code : $? \n"


# API Version1 Root
echo "GET /v1"
curl -i -L \
  -X GET \
  -H "Content-Type: application/json" \
  "localhost:8080/v1"
echo "exit code : $? \n"

# Index
echo "GET /v1/jirou"
curl -i -L \
  -X GET \
  -H "Content-Type: application/json" \
  "localhost:8080/v1/jirou"
echo "exit code : $? \n"

# Create
echo "POST /v1/jirou"
curl -i -L \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"param1":"foo", "param2":"bar"}' \
  "localhost:8080/v1/jirou"
echo "exit code : $? \n"


# Read
echo "GET /v1/jirou/0"
curl -i -L \
  -X GET \
  "localhost:8080/v1/jirou/0"
echo "exit code : $? \n"


# Update
echo "PUT /v1/jirou/0"
curl -i -L \
  -X PUT \
  -H "Content-Type: application/json" \
  -d '{"param1":"foo", "param2":"bar"}' \
  "localhost:8080/v1/jirou/0"
echo "exit code : $? \n"

# Delete
echo "DELETE /v1/jirou/0"
curl -i -L \
  -X DELETE \
  "localhost:8080/v1/jirou/0"
echo "exit code : $? \n"
