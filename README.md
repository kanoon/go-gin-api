Build API

```
docker build -t my-go-app .
docker run -p 8001:8001 my-go-app
```

Build DB

```
docker-compose up --build
docker-compose down -v
```
