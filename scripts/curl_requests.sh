echo "---- Hitting the 'Hello, World!' endpoint ----"
curl http://localhost:8080/hello

echo "---- Retrieving all users ----"
curl http://localhost:8080/api/v1/users | jq

echo "---- Retrieving all posts ----"
curl http://localhost:8080/api/v1/posts | jq
