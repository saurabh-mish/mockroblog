## Project Dependencies

Install project dependencies using a package manager

Current dependencies include:

+ `go`
+ `curl`
+ `jq`


## Run Application

Run the below commands from the project root

+ Execute the main application in a terminal tab

  ```zsh
  go run ./cmd/app/main.go
  ```

+ Make the shell script executable

  ```zsh
  chmod +x ./scripts/curl_requests.sh
  ```

+ In a new terminal tab, execute `curl` requests for existing endpoints


  ```zsh
  ./scripts/curl_requests.sh
  ```
