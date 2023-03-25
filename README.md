To run the application, you have to run
```shell
docker compose up
```
- The application was designed to serve HTTP REST API at PORT 8000, but with consideration to pick the driver you need while exposing port (e.g: we can set gRPC driver and expose it as well)
- The application is going to create the database if it doesn't exist and connect to it automatically. Therefore, there is no any prerequisites to bootstrap the app.  

> Create Deck Request
```sh
curl --location -X POST 'localhost:8000/api/decks' \
--header 'Content-Type: application/json' \
--data-raw '{
  "shuffle": true
}'
```
### For Deck Card Creation with filtered Cards

```sh
curl --location -X POST 'localhost:8000/api/decks?cards=AS,KD,AC,2C,KH' \
--header 'Content-Type: application/json' \
--data-raw '{
  "shuffle": false 
}'
```
> Open Deck By ID Request

##### Note: Pass a valid UUID that's retrieved from the previous request

```shell
curl  -X GET 'localhost:8000/api/decks/cb3e373f-a51c-421d-ad05-91aff0df890a' \
-H 'Content-Type: application/json' 

```

>Draw Card Request
```shell
curl  -X PATCH 'localhost:8000/api/decks/cb3e373f-a51c-421d-ad05-91aff0df890a' \
-H 'Content-Type: application/json' 
```
### To Draw multiple cards at once

```sh
curl  -X PATCH 'localhost:8000/api/decks/cb3e373f-a51c-421d-ad05-91aff0df890a?count=1' \
-H 'Content-Type: application/json' 
```
### FOR RUNNING TESTS 
After running docker compose up

to start testing deck suite
> ENV=test go test -v ./tests/... -count=1   
