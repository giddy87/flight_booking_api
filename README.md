# Flight Booking API
**Written in Go**
**Database**: Postgresql

## Deploy application

Step 1: Generate Dummy flight data in a flights.csv file (optional)

    go run dummydatagenerator.go 
Step 2: Set env variables. for example,

    export DB_PASSWORD=fli7g1ht2api129 \
    export DB_NAME=flightapi \
    export DB_USER=flightapiuser \
    export DB_HOST=postgres \
    export DB_PORT=5432


**Docker deployment**

    docker-compose up -d
*This comprises the application, postgres and dataloader container

Application is available at <"your ip">:9000 | localhost:9000

## Api documentation

 1. **[/api/flight docs](api-docs/flights.md)**

