# Simple CRUD API 

This go project illustrates how to create a simple CRUD API using Go.

Data are stored in memory using slices and structs and they represent a collection of movies. 

The available endpoints are:

/movies (GET)=> returns a Json object with the list of movies 
/movies/id (GET)=> returns a Json object which include the movie with id=id
/movies (POST) => create a new movie and returns 
/movies/id (DELETE) => delete the movie with id=id and returns a json with the list of movies
/movies/id (PUT) => update the movie with id=id and returns it as a Json object

TO DO: edge cases are not taken into account - i.e. malformed requests



