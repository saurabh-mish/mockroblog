# Project Specification

We’re going to build services for a web application similar to [reddit][1].

## Development

To get started, you will build two microservices for posting and voting functionality, with their automation test suites. After this, you will add two more microservices for account and messaging functionality and their automation test suites.

### Services

**Posting Microservice**

Each post should have a title, text, a community ([subreddit][2]), an optional URL linking to a resource (e.g. a news article or picture), a username, and a date the post was made.

The following operations should be exposed:

+ Create a new post
+ Delete an existing post
+ Retrieve an existing post
+ List the n most recent posts to a particular community
+ List the n most recent posts to any community

When retrieving lists of posts, do not include the text or resource URL for the post.


**Voting Microservice**

Each post maintained by the posting microservice can be voted up or down. This service should maintain the number of upvotes and downvotes for each post. A post’s score can be computed by subtracting the number of downvotes from the number of upvotes.

The following operations should be exposed:

+ Upvote a post
+ Downvote a post
+ Report the number of upvotes and downvotes for a post
+ List the n top-scoring posts to any community
+ Given a list of post identifiers, return the list sorted by score.

Each upvote or downvote should include a unique identifier (e.g., a URL or database key) for the post that can be used to match votes with the posts maintained by the posting microservice.
If this service is implemented with a database separate from the posting microservice, it is not responsible for verifying the existence of a post before recording or reporting votes.


**Account Microservice**

Each user who registers should have the following data associated with them:
+ Username
+ Email
+ Karma

The following operations will be exposed:

+ Create user
+ Update email
+ Increment Karma
+ Decrement Karma
+ Deactivate account

The data for the user can be in the same database or different database as the other services.


**Messaging Microservice**

Users can send and receive messages to each other. Messages will consist of the following data associated with them:

+ Message ID
+ User from
+ User to
+ Message timestamp
+ Message contents
+ Message flag

The following operations will be exposed:

+ Send message
+ Delete message
+ Favorite message

Messaging data can be in the same database as other services or a separate one.

---

### API

**Implementation**

Implement your APIs in Python 3 using [Flask][3]. You are encouraged, but not required, to use [Flask API][4] to obtain additional functionality.
All data, including error messages, should be in JSON format with the Content-Type header field set to application/json.


**Documentation**

Developers will need to create and maintain a specification for the services they create and implement.

For example, a hypothetical service to manage customer information at a bank: A customer is defined as having:

+ Customer_id: a unique id for all customers
+ Email: text field for customer email
+ Address: text field for customer address

| HTTP Method | URI                                                                 | Action                             |
|:-----------:|:-------------------------------------------------------------------:|:----------------------------------:|
| GET         | http://[hostname]/banktec/api/v1.0/customer[customer_id]            | Retrieve a list of customer IDs    |
| POST        | http://[hostname]/banktec/api/v1.0/customer                         | Create a new customer              |
| POST        | http://[hostname]/banktec/api/v1.0/ customer[customer_id, email]    | Update customer’s email            |


---

### Other

**HTTP Status Codes**

Use appropriate HTTP status codes for each operation, with the following guidelines:

+ In general, successful operations other than POST should return HTTP 200 OK.

+ A successful POST should return HTTP 201 Created, with the URL of the newly created object in the Location header field.

+ Attempts to retrieve or modify an existing object should return HTTP 404 Not Found if the specified object does not exist (or no longer exists). Note that this does not apply to objects maintained by other services.

+ Operations which result in a constraint violation such as attempting to INSERT a duplicate value into a column declared UNIQUE or attempting to INSERT a row with a FOREIGN KEY referencing an item that does not exist in another table should return HTTP 409 Conflict.


**Session State**

Requests to each microservice must include all information necessary to complete the request; your APIs must not use the Flask session object to maintain state between requests.


**Database**

Use The Python Standard Library’s [sqlite3][5] module as the database for your Flask application. You may use separate databases for each Flask application, or share a database across microservices.

---
---

## Testing and Automation

### Validation Testing

Each microservice should have an accompanying test script to verify that your newly-defined API endpoints work correctly and to populate the microservices with some sample data. Suitable approaches to scripting include:

+ A shell script that calls [curl][6] commands

  Note that it is complicated to [determine whether a curl command has succeeded][7] programmatically, so you will probably need to read the output carefully.

+ A Python script using the [Requests][8] library

  Use the following command to install Requests on Tuffix:

  ```zsh
  pip3 install --user requests
  ```

+ A YAML script using the [Tavern][9] plugin for [pytest][10]

  Use the following command to install Tavern on Tuffix:

  ```zsh
  pip3 install --user tavern
  ```

---

### System Testing

In addition to basic testing, each group should come up with a framework to test it’s services with load, and in simulated user-scenarios. Multiple users should be able to be simulated concurrently. Each group may pick whatever frameworks and automation tools will be needed to test and simulate this.

Your system test suite must:

+ Test all services
+ Perform a load test simulating 100 users
+ Simulate a real user scenario using the two services
+ Stress the service with: excessive load, bogus data, negative tests


---


[1]: https://www.reddit.com
[2]: https://www.dictionary.com/e/slang/subreddit/
[3]: https://flask.palletsprojects.com/en/2.3.x/
[4]: https://flask.palletsprojects.com/en/1.1.x/#api-reference
[5]: https://docs.python.org/3.10/library/sqlite3.html
[6]: https://alvinalexander.com/web/using-curl-scripts-to-test-restful-web-services/
[7]: https://stackoverflow.com/questions/38905489/how-to-check-if-curl-was-successful-and-print-a-message
[8]: https://realpython.com/api-integration-in-python/
[9]: https://taverntesting.github.io
[10]: https://docs.pytest.org/en/7.3.x/
