# siap

siap database access using sql server

```.env
MSSQLSTRING=server%3Dhost%3Buser%20id%3Duser%3Bpassword%3Dpass%3Bport%3D1433%3Bdatabase%3Ddbname%3B
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
