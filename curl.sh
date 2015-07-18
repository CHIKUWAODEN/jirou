#!/bin/sh


# API Root
echo "GET /"
curl -s \
  -X GET \
  -H "Content-Type: application/json" \
  "localhost:8080/" | jq "."


# API Version1 Root
echo "GET /v1"
curl -s \
  -X GET \
  -H "Content-Type: application/json" \
  "localhost:8080/v1" | jq "."


# Index
echo "GET /v1/jirou"
curl -s \
  -X GET \
  -H "Content-Type: application/json" \
  "localhost:8080/v1/jirou" | jq "."


# Create
# echo "POST /v1/jirou"
# curl -s \
#   -X POST \
#   -H "Content-Type: application/json" \
#   -d '{"param1":"foo", "param2":"bar"}' \
#   "localhost:8080/v1/jirou" | jq "."


# Read
for id in {0..4}
do
  echo "GET /v1/jirou/${id}"
  curl -s \
    -X GET \
    "localhost:8080/v1/jirou/${id}" | jq "."
done


# Post Report
POST_DATA=`cat <<EOS
{
  "reporter" : "Jhon Smith",
  "comment" : "lorem ipsum dolor sit amet ... (Free comment)",
  "rating" : {
    "noodle" : 5.0,
    "soup" : 5.0
  }
}
EOS
`
for id in {0..4}
do
  echo "POST /v1/jirou/${id}/report"
  echo ${POST_DATA}
  curl -s \
    -X POST \
    -d "${POST_DATA}" \
    "localhost:8080/v1/jirou/${id}/report" | jq "."
done


# Get Report
for id in {0..4}
do
  echo "GET /v1/jirou/${id}/report"
  curl -s \
    -X GET \
    "localhost:8080/v1/jirou/${id}/report" | jq "."
done


# Update
# echo "PUT /v1/jirou/0"
# curl -s \
#   -X PUT \
#   -H "Content-Type: application/json" \
#   -d '{ "param1": "foo", "param2": "bar" }' \
#   "localhost:8080/v1/jirou/0" | jq "."


# Delete
# echo "DELETE /v1/jirou/0"
# curl -s \
#   -X DELETE \
#   "localhost:8080/v1/jirou/0" | jq "."

