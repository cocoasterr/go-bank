# Go Bank API

## Setup your Database
- insert .env following format .env copy
```bash
PG_DB_DSN = "postgres://<postgres>:<password>@<host>:<port>/"
```

## Init Your Go Module
```bash
go mod init github.com/cocoasterr/go_bank
```
## Download Module and create Go binary file
```bash
make build
```
## Docker Build
```bash
docker build -t go_bank:tag .
```
Example like this:
 ```bash 
 - docker build -t warung-app:1.0 .
```
## Docker run
```bash
docker run -p outer-port:inner-port nama-image:tag
```
Notes:
- Inner-port using for golang app
- and outer-port using for browser to acces the golang app

example like this:
 ```bash 
docker run -p 3000:3000 no_image
```
- my golang app using port 3000
