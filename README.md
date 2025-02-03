# My Numbers API task

## Resources
- [Fun fact API](http://numbersapi.com/#42)
- [Parity numbers](https://en.wikipedia.org/wiki/Parity_(mathematics))

## Task Description
- Create an API that takes a number and returns interesting mathematical properties about it, along with a fun fact.

### Requirements
#### Technology Stack:
1. Use any programming language or framework of your choice (See Sharp (C #), PHP :elephant:, Python :snake:, Go :runner::skin-tone-5:, Java :coffee:, JS/TS :nauseated_face:)
- Must be deployed to a publicly accessible endpoint
- Must handle CORS (Cross-Origin Resource Sharing)
- Must return responses in JSON format

### Version Control:
- Code must be hosted on GitHub
 - Repository must be public
 - Must include a well-structured README.md

### API Specification
 - Endpoint: `GET** <your-domain.com>/api/classify-number?number=371`
 - Required JSON Response Format (200 OK):

### SUCCESS 200 OK

A 200 response is returned from the server when the query passed is a number.

```JSON
{
    "digit_sum": 11,
    "fun_fact": "173 is a square-free number, divisible by no perfect square.",
    "is_perfect": false,
    "is_prime": true,
    "number": 173,
    "properties": [
        "odd"
    ]
}
```

### ERROR 400 BAD REQUEST

A 400 status is returned to client when the user sends a parameter as an alphabet instead of a number.

`curl --location 'localhost:5050/api/classify-number?number=redeye'`

```JSON
{
    "error": "true",
    "number": "alphabet"
}
```