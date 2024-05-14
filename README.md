### Debug API

Require to update `config/db.go` file and your DB should be ready.

```
go run .\main.go
```

### Run API (only) with docker container

```
docker build -t my-go-app .
docker run -p 8001:8001 my-go-app
```

### Run DB and API with docker container

```
docker-compose up --build
```

#### Site URLs

- Docker container: http://localhost:8080/users
- Debug: http://localhost:8001/users

#### Database

- Docker container: db,5432
- Debug or pgAdmin: 127.0.0.1, 5444
