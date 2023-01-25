# siap

siap database access using sql server

```.env
MSSQLSTRING="server=HOST;user id=USER;password=PASS;port=61782;database=DBNAME;"
```

if u dont like using .env just use var

```go
var server = "host"
var port = 1433
var user = "user"
var password = "pass"
var database = "dbname"
connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
  server, user, password, port, database)
```
