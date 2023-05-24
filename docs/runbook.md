## Project Dependencies

Install project dependencies:

+ `go`
+ `curl`
+ `jq`

---

## Execute Tests

Run the below commands from the project root

+ Run tests

  ```zsh
  make exectests
  ```

+ Record coverage for dedicated test package

  ```zsh
  make coverage
  ```

+ Generate HTML report from coverage data

  ```zsh
  make report
  ```

+ View the HTML file `coverage.html` in a browser

---

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
