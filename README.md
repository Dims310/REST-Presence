# REST-Presence
This is the first REST from what I've learned so far, I can make a REST API about the "presence" system.

There are 5 endpoints with each method in this REST:

- ```/dasboard (GET)```; at that endpoint, the user can run queries and see all of the data that has been entered into the database. 
- ```/presence (POST)```, this endpoint creates a database record and stores the id, name, and time of attendance. 
- ```/absence (PUT)```, this endpoint modifies data (absence column) via the search record "id" and returns a boolean to the "already_absence" column to indicate whether or not absence was made.
- ```/dashboard/:id (GET)```, the same as the "/dashboard" endpoint but with the additional id parameter behind the endpoint (query parameter), which aims to retrieve only one data record.
- ```/delete/:id (DELETE)```, this endpoint works to delete record data via the additional id parameter behind the endpoint (query parameter).

Tested with POSTMAN

Using https://github.com/julienschmidt/httprouter for making all route.
