# ReconDB
![go]

ReconDB is a project written in Go which provides APIs for managing companies, scopes, assets, and outscopes data.

The project uses MongoDB as a backend to store and retrieve data.

---

## Endpoints
The following endpoints are available in the project:

**Scope Endpoint**
- POST /api/scope: Add a new scope

- GET /api/scope/:companyname: Get scopes for a specific company

- GET /api/scope: Get all scopes

- DELETE /api/scope/:companyname: Delete scopes for a specific company

---
**Out of Scope Endpoint**
- POST /api/outscope: Add a new outscope

- GET /api/outscope/:companyname: Get out of scope for a specific company

- GET /api/outscope: Get all out of scopes

- DELETE /api/outscope/:companyname: Delete out of scope for a specific company

---
**Company Endpoint**
- POST /api/company: Add a new company

- GET /api/company/:companyname: Get a specific company

- GET /api/company: Get all companies

- DELETE /api/company/:companyname: Delete a specific company

---
**Asset Endpoint**
- POST /api/asset: Add a new asset

- GET /api/asset/:asset: Get a specific asset

- GET /api/asset: Get all assets

- DELETE /api/asset/:asset: Delete a specific asset


## Configuration

The project requires a configuration file in JSON format to specify the authorization token,
the port number to listen on, the Gin mode, and the MongoDB URI.

Here is an example of the configuration file:

```json
{
  "authorization": "<token>",
  "gin_mode": "debug",
  "port": ":8080",
  "mongo_uri": "mongodb://recondb-mongodb-1:27017"
}

```

- `authorization`: The authorization token to be used for API requests
- `gin_mode`: The mode for the Gin web framework
- `port`: The port number to listen on
- `mongo_uri`: The URI for the MongoDB instance to connect to

For generating token simply use the following command in python3 :
```python
import os
os.urandom(32).hex()
```

## Running the Project
To run the project you have to first compile the program and then run the compiled binary.
```bash
 go build -o ReconDB -ldflags "-s -w -buildid=" -buildvcs=false .
 ./ReconDB
```

## Running the Project With Docker
To run the project with docker , you can use the docker-compose file :
```bash
docker-compose up -d
```
**Note** : make sure to set the MongoDB URI to name of docker name of the mongodb ex:
```bash
{
  ...
  "mongo_uri": "mongodb://recondb-mongodb-1:27017"
}
```


[go]: https://img.shields.io/badge/Go-cyan?logo=go
