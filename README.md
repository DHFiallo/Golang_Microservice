# Go Microservice

Endpoints Available:
- Endpoint to change data based on user ID
- Endpoint to return a specific user (and all associated information)
- Endpoint to return a list of users created between a date range
- Endpoint to return a list of users based on a specific profession


cURL Requests

PUT Request for Changing Data based on ID
curl -v localhost:9090/105 -X PUT -d "{\"first\":\"rucker\",\"last\":\"roy\",\"email\":\"roy.rucker@gmail.com\",\"profession\":\"engineer\",\"datecreated\":\"2023-01-23\",\"Country\":\"Mexico\",\"City\":\"Cancun\"}"

changes ID 105 with these attributes

GET Request for Specific Person
curl localhost:9090/name/rucker/roy

GET Request for Dates
curl localhost:9090/date/2018-01-01/2020-01-20

GET Request for Specific Job
curl localhost:9090/profession/doctor
