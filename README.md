cch
===

CouchCB CLI tool.

## REST

### GET

```
cch get http://localhost:5984
{"couchdb":"Welcome","version":"1.2.1"}
```

### POST

```
cch post http://localhost:5984/test -d '{foo: bar}'
{"ok":true,"id":"415f519538c5382d33781819c3000457","rev":"1-967a00dff5e02add41819138abb3284d"}
```

### PUT

```
cch put http://localhost:5984/test/test -d '{foo: bar}'
{"ok":true,"id":"test","rev":"1-967a00dff5e02add41819138abb3284d"}
```

### DELETE