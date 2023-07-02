# Getting started
echo "---- 'Hello, World!' endpoint ----"
curl http://localhost:8080/api/v1/hello

##############################
#### Account Microservice ####
##############################

echo "---- Retrieve all users ----"
curl http://localhost:8080/api/v1/users | jq

echo "---- Create new user ----"
curl --request POST --url http://localhost:8080/api/v1/user \
	--header "Content-Type: application/json" \
	--data '{"username":"testuser","email":"testuser@domain.com","password":"p@ssword"}'

echo "---- Retrieve existing user with ID ----"
curl http://localhost:8080/api/v1/user/0 | jq

###########################
#### Post Microservice ####
###########################

echo "---- Create new post ----"
curl --request POST --url http://localhost:8080/api/v1/post \
	--header "Content-Type: application/json" \
	--data '{"title":"title 1","content":"The quick brown fox jumps over the lazy dog","community":"playground"}'

echo "---- Retrieve existing post with ID ----"
curl http://localhost:8080/api/v1/post/id=14 | jq

echo "---- Delete existing post ----"
curl -X DELETE http://localhost:8080/api/v1/post/id=3

echo "---- Retrieve n most recent posts ----"
curl http://localhost:8080/api/v1/posts/number=1

echo "---- Retrieve n most recent posts from community ----"
curl "http://localhost:8080/api/v1/posts/number=1&community=playground"
