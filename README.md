cch
===

CouchCB CLI tool.

## Binaries

[darwin-386](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-darwin-386)
[darwin-amd64](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-darwin-amd64)
[freebsd-386](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-freebsd-386)
[freebsd-amd64](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-freebsd-amd64)
[freebsd-arm](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-freebsd-arm)
[linux-386](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-linux-386)
[linux-amd64](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-linux-amd64)
[linux-arm](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-linux-arm)
[windows-386](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-windows-386)
[windows-amd64](https://s3.amazonaws.com/labs.healpay.com/cch/bin/-windows-amd64)

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