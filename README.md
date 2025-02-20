# djan-go

**djan-go** is a lightweight Go framework for building REST APIs, inspired by Django. It simplifies defining models, handling database connections, and managing authentication and permissions.

## Installation

```bash
go get github.com/yourusername/djan-go
```

## Setup

Create a `.env` file with your PostgreSQL connection URL:

```
PGURL=your_postgresql_connection_url
```

## Example Usage

1. **Load environment and configure PostgreSQL**:

```go
err := godotenv.Load("test.env")
if err != nil {
    log.Fatal("Error loading .env file")
}
pgurl := os.Getenv("PGURL")
c, err := NewPostgresConfig(pgurl)
if err != nil {
    fmt.Println(err)
}
c.Debug = false
```

2. **Define and register models**:

```go
RegisterDefaultHttpModel[Product](&HttpDataModel[Product]{
    EndPointName: "product",
    GlobalConfig: c,
    Permissions:  roles.Allow(roles.CRUD, "admin"),
    Auth:         true,
})

RegisterDefaultHttpModel[Person](&HttpDataModel[Person]{
    EndPointName: "person",
    GlobalConfig: c,
    Permissions:  roles.Allow(roles.CRUD, "admin"),
})
```

3. **Start the server**:

```go
http.Handle("/", c.Router)
err = http.ListenAndServe(":8000", nil)
if err != nil {
    fmt.Println(err)
}
```

## Authentication & Permissions

Use `roles.Allow(roles.CRUD, "admin")` to grant CRUD permissions to the admin role.

## License

MIT License
