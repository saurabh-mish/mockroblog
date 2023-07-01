echo "---- Hitting the 'Hello, World!' endpoint ----"
curl http://localhost:8080/api/v1/hello

echo "---- Retrieving all users ----"
curl http://localhost:8080/api/v1/users | jq

echo "---- Retrieving user with ID ----"
curl http://localhost:8080/api/v1/user/560

echo "---- Retrieving all posts ----"
curl http://localhost:8080/api/v1/posts | jq

#echo "---- Create user ----"
#curl --request POST "http://localhost:8080/api/v1/user?username=testuser&email=testuser@domain.com&password=p@ssword"

echo "---- Create user ----"
curl --request POST --url http://localhost:8080/api/v1/user \
	--header "Content-Type: application/json" \
	--data '{"username":"testuser","email":"testuser@domain.com","password":"p@ssword"}'
