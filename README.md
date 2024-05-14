### Build DB

```
docker-compose down -v #if needed
docker-compose up -d
```

### Debug API

```
go run .\main.go
```

### Build API

```
docker build -t my-go-app .
docker run -p 8001:8001 my-go-app
```
