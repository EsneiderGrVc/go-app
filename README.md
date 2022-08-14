
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
After clone git repository, add Firestore credentials in `./api/serviceAccount.json` file
```bash
git clone https://github.com/EsneiderGrVc/go-app.git
cd go-app/
docker-compose up --build
```
> **Note**: Before running `docker-compose up`
You must have docker and docker compose installed.
Make sure you do not have any services listening on ports 80: and 5434: on your local machine.
Stops the execution of other containers, if any.

## Usage
Visit http://localhost/ or

**Go to** [Swagger Documentation](http://localhost/swagger/index.html) **and try it out.**

## What is missing?
- Swagger Documentation - _In progress_
- Validate object data types before calling services.
- Unit testing.
- Implementing an ORM.
- Automigrate functionality.
- Create a transaction table and create a relationship between deliveries and bots.
- Assign a bot to an order functionality.
- Assign pending orders functionality.
- Apply concurrency to load Firestore collections faster.

**Esneider Granada Valencia** - *Software Engineer* - [Linkedin](https://www.linkedin.com/in/esneider-granada/)