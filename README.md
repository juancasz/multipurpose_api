# multipurpose api

REST API built using Golang with functionalities to manage users and perform some calculations. Basic Authentication implemented.  
   
This project can be run using docker compose in the following way:  
```
    docker-compose up --build -d
```
  
After starting the container, the REST API endpoints will be available at localhost:8888 and a PostgreSQL database will appear at localhost:5432. The following parameters can be used to connect to the created database:
   
```
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=postgres_database
    DB_USER=postgres
    DB_PASSWORD=admin
```
  
The relevant tables in the database are:  
```sql
    SELECT * FROM users u ;
    SELECT * FROM countries c ;
    SELECT * FROM universities u ;
```

You can import the endpoints to your local postman using the file **multipurpose_api.postman_collection.json** provided in this repository.  
  
You first need to get a token using the following endpoint:  
   
```
    POST http://localhost:8888/multipurpose-api/login
```
  
with request body:  
```json
    {
        "username": "juan123",
        "password": "password_test"
    }
```  
  
The other endpoints will require the header:

```
    Authorization : Basic <token>
```