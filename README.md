# go_API_Prac

This is just a sample practice project that I'm doing to try and understand Go Programming

I'll be trying to structure the project in a prod level mvc style and hopefully I'll be able to dockerize the application efficiently.
I'll be tracking progress in this readme as well.

So far the folder structure is as follows:

```bash
go_api_prac
├───app
│   ├───handlers
│   └───models
└───config
```
In this structure

- *app* contains **app.go** which takes care of routing calls using [Gorilla/Mux](https://github.com/gorilla/mux) and two other folders as well
   - *handlers* contains all the functions which the endpoint uses to serve clients. Each function will have its own file to improve modularity. Currently only one endpoint exists
   - *models* contains all the structs as separate files that are used to decode info into and encode from to display to the client/insert data into the server
- *config* contains the connection configuration related to the database connection. I plan to expand it's functionality to include some API configuration details as well.
- The root directory also contains **main.go** which just runs the routing function from **app.go**. Additional functionality might be added later

All the data served by the endpoints is currently being pulled from a MongoDB collection which contains data imported from the [The Complete Pokemon Dataset](https://www.kaggle.com/rounakbanik/pokemon) from Kaggle. The collection is hosted and being connnected to from the localhost through default parameters

The current available endpoints are:

```
* \pokemon -> Returns all the documents in the collection and displays them as a collection of JSON objects

```

More endpoints will be implemented soon.
Comprehensive code comments and documentation are also next on the list of tasks to do




