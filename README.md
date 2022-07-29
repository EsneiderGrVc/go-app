# Go App

This application has a docker-compose configuration that allows install and execute it with just a command
Firestore collections (deliveries, bots) are automatically created and loaded with SQL dummy data once docker-compose is executed .

## Pre-requisites
- Docker
- Docker Compose
> If you haven't Docker installed, brew brings an easy way to do it on Linux and MacOs

**Install HomeBrew**
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
**Install Docker**
```bash
brew install --cask docker
brew install docker-compose
```

## Install and Execution
```bash
git clone https://github.com/EsneiderGrVc/go-app.git
cd go-app/
docker-compose up
```
After clone git repository, add Firestore credentials in `./api/serviceAccount.json` file
> **Note**: Before running `docker-compose up`
You must have docker and docker compose installed.
Make sure you do not have any services listening on ports 80: and 5434: on your local machine.
Stops the execution of other containers, if any.

## Usage

|Endpoint|Method|Description               |
|---------------------|----------|-----------------------------|
|`localhost`|GET|Does Nothing.|
|`localhost/deliveries`|GET|Get all the documents inside of **deliveries** collection.|
|`localhost/deliveries/{id}: string`| GET | Get a specific document by its id.|
|`localhost/deliveries`|POST|Create a new delivery with [this structure](#create-delivery-structure).|
|`localhost/bots`|GET| Get all the documents inside of **bots** collection ordered by zone.|
|`localhost/bots`|POST|Create a new bot with [this structure](#create-bot-structure).|

### Create Delivery structure
```json
{
    "id": "342",
    "creation_date": "2022-07-29T10:52:35.401626Z",
    "pickup": {
        "pickup_lat": 38.540644,
        "pickup_lon": 141.064362
    },
    "state": "assigned",
    "dropoff": {
        "dropoff_lat": 18.913308,
        "dropoff_lon": 72.822865
    },
    "zone_id": "004"
}
```
> try: 
```bash
curl -X POST localhost/deliveries -H "Content-Type: application/json" -d '{"id":"654","creation_date":"2022-07-29T10:52:35.401626Z","pickup":{"pickup_lat":38.540644,"pickup_lon":141.064362},"state":"assigned","dropoff":{"dropoff_lat":18.913308,"dropoff_lon":72.822865},"zone_id":"194"}'
```

### Create Bot structure
```json
{
    "id": "10",
    "status": "available",
    "location": {
        "lat": 23.092837,
        "lon": -23.938493
    },
    "zone_id": "13"
}
```
> try: 
```bash
curl -X POST localhost/bots -H "Content-Type: application/json" -d '{"id":"432","status":"available","location":{"lat":23.092837,"lon":-23.938493},"zone_id":"13"}'
```

### What is missing?
- Swagger Documentation
- Validate object data types before calling services.
- Unit testing.
- Implementing an ORM.
- Automigrate functionality.
- Create a transaction table and create a relationship between deliveries and bots.
- Assign a bot to an order functionality.
- Assign pending orders functionality.
- Apply concurrency to load Firestore collections faster.

**Esneider Granada Valencia** - *Software Engineer* - [Linkedin](https://www.linkedin.com/in/esneider-granada/)