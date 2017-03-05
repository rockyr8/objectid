ObjectId package for Go
====
This package provides build a unique object identifier for the high concurrent environment.

The `ObjectId` can instead of GUID,and are stored as 12-bytes.

Reference from : [ObjectId of Mongodb](http://docs.mongodb.org/manual/reference/object-id/)

> ObjectId is a 12-byte BSON type, constructed using:

> a 4-byte value representing the seconds since the Unix epoch,

> a 3-byte machine identifier,

> a 2-byte process id, and

> a 3-byte counter, starting with a random value.

Installation
====
Use the `go` command:
```go
go get github.com/zhengchun/objectid
```

Example
====
```go
package main

import (
	"fmt"
	"github.com/zhengchun/objectid"
)

func main() {
	objid := objectid.New()
	fmt.Printf("ObjectId: %s\n", objid)
	objid = objectid.Parse("55cc60946f29581a606930aa")
	fmt.Printf("%d-%d-%d-%d\n", objid.Timestamp(), objid.Machine(), objid.Pid(), objid.Increment())
}

```