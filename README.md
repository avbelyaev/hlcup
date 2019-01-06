# HL Cup

Vendor usage
```
go get github.com/kardianos/govendor

cd here
govendor init

# add dependency
govendor fetch github.com/pkg/errors

# remove not used
govendor remove +unused

# add everything from GOPATH
govendor add +external

```