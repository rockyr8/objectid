objectid [![GoDoc](https://godoc.org/github.com/zhengchun/objectid?status.svg)](https://godoc.org/github.com/zhengchun/objectid)
====
objectid package provides build a unique object identifier and are stored as 12-bytes.

objectid is a 12-bytes, constructed using:
- a 4-byte value representing the seconds since the Unix epoch
- a 3-byte machine identifier
- a 2-byte process id, and
- a 3-byte counter, starting with a random value.

Reference : [ObjectId](http://docs.mongodb.org/manual/reference/object-id/)

installation
====
> go get github.com/zhengchun/objectid

example
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
	objid, _ = objectid.Parse(objid.String())
	fmt.Printf("%d-%d-%d-%d\n", objid.Timestamp(), objid.Machine(), objid.Pid(), objid.Increment())
}
```