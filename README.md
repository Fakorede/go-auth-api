# Auth API

> A simple golang based api that shows how to implement a REST Auth API using the net/http package


### Setup

```
git clone github.com/fakorede/go-auth-api
cd go-auth-api
```

> Create database and run query in `db.sql` file

Create .env
```
cp .env.example .env
```

Add env variables
```
JWT_SECRET={your-secret}
DATABASE_URL=postgres://{your-username}:{your-password}@localhost/{your-database}?sslmode=disable
```

### Run Application
```
go mod download
go run main.go
```


### Endpoints

|  **Method** |  **Uri** |
|---|---|
|  POST |  /api/signup |
|  POST | /api/login  |
|  GET |  /api/protected |

### Third-party Dependencies

- [Mux router](github.com/gorilla/mux)
- [Godotenv](github.com/joho/godotenv)
- [JWT](github.com/dgrijalva/jwt-go)
- [PG Driver](github.com/lib/pq )

