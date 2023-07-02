echo "---- 'Hello, World!' endpoint ----"
curl http://localhost:8080/api/v1/hello

##############################
#### Account Microservice ####
##############################

echo "---- Retrieve all users ----"
curl http://localhost:8080/api/v1/users | jq

echo "---- Retrieve non-existing user ----"
curl http://localhost:8080/api/v1/user/560

echo "---- Create user ----"
curl --request POST --url http://localhost:8080/api/v1/user \
	--header "Content-Type: application/json" \
	--data '{"username":"testuser","email":"testuser@domain.com","password":"p@ssword"}'

echo "---- Retrieve existing user with ID ----"
curl http://localhost:8080/api/v1/user/0 | jq

###########################
#### Post Microservice ####
###########################

echo "---- Retrieve all posts ----"
curl http://localhost:8080/api/v1/posts | jq

