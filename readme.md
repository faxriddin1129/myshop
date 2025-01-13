# GOLANG BACKEND ````GOFAX````

### USE PACKAGES
```
go get -u github.com/gorilla/mux
```
```
go get -u gorm.io/gorm
```
```
go get -u gorm.io/driver/postgres
```
```
go get github.com/joho/godotenv
```
```
go get github.com/go-playground/validator/v10
```

### DATABASE `POSTGRESSQL`
Supports the latest version of Postgres!
Uses `gorm` package for ORM.

### CACHE OPERATIONS
This document outlines the basic operations for working with cache in Golang. The cache can be used to store, retrieve, and delete data in a backend application.
Uses `.bin` files!
### Cache Set
Use the `Set` method to store data in the cache. You can set a specific expiration time for the cached data.
```
cache.Set("cache_key", data, 1*time.Hour)
```
### Cache GET
Use the `Set` method to store data in the cache. You can set a specific expiration time for the cached data.
```
cache.Get("cache_key", data, 1*time.Hour)
```
### Cache Delete
Use the `Set` method to store data in the cache. You can set a specific expiration time for the cached data.

```
cache.Delete("cache_key", data, 1*time.Hour)
```