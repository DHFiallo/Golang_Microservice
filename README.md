# MagMutualTechnical
Take home technical project from MagMutual

Create a microservice

We would like you to build a Microservice that exposes a couple endpoints to return the
following information. We would also like a simple UI to display the information.

We have provided a CSV of user information; we would like you to build to build the following
endpoints:
- Endpoint to return a specific user (and all associated information)
- Endpoint to return a list of users created between a date range
- Endpoint to return a list of users based on a specific profession
- Custom Endpoint that you design on your own.


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






Questions:

1. What did you think of the project?
Tough. Never truly worked with microservices and APIs before, beyond small examples and guides. Had to watch a lot of videos and pour through a lot of documentation just to get through this. Definitely a rollercoaster of emotions, going from "Wow I'm way too dumb for this" and back to "I'm a genius". Simultaneously not happy with my project but also proud, I'll definitely be working on it some more but I imagine there's a deadline for this and so stopped after about 1-2 days.

2. What didn’t you like about the project?
I'm sure it's different from person to person but wow UIs/GUIs are so hard. I tried several frameworks before ending up on this terminal UI just because I was struggling to figure them out. The microservice part was very confusing but I bashed my head against it, reminded myself that abstraction exists for a reason, and eventually things just started working. Probably not a good thing to admit.

3. How would you change the project or approach?
If I had more time, I'd make a proper GUI. I'm not sure if calling the microservice via cURL is allowed or acceptable but either way I'd like to make it a lot more proper and have like two programs, one running the server and the other calling it. I'd validate info a lot more, add some proper testing to it, and basically just make it nice and neat and more resistant to a bad actor.

4. Anything else you would like to share?
It was fun and frustrating. I'd like to know how it compares to actual business practices and so I'm really excited for the code review even if I'm suspecting what I've done is quite bad. Had to learn a lot of rather new stuff, never done this before and Go is my newest language but super fun.
