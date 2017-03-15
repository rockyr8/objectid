/*
Package ObjectId provides build a unique object identifier and are stored as 12-bytes.

ObjectId construct format:
a 4-byte value representing the seconds since the Unix epoch,
a 3-byte machine identifier,
a 2-byte process id, and
a 3-byte counter, starting with a random value.

http://docs.mongodb.org/manual/reference/object-id/
*/
package objectid

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

var staticMachine = getMachineHash()
var staticIncrement = getRandomNumber()
var staticPid = int32(os.Getpid())

// A globally unique identifier for objects.
type ObjectId struct {
	timestamp int64
	machine   int32
	pid       int32
	increment int32
}

// News generates new ObjectID with a unique value.
func New() ObjectId {
	timestamp := time.Now().Unix()
	return ObjectId{timestamp, staticMachine, staticPid, atomic.AddInt32(&staticIncrement, 1) & 0xffffff}
}

// Parses a string and creates a new ObjectId.
func Parse(input string) (objid ObjectId, err error) {
	if objid, err = tryParse(input); err == nil {
		return
	}
	return objid, fmt.Errorf("%s is not a valid 24 digit hex string", input)
}

func (objid ObjectId) Timestamp() int64 {
	return objid.timestamp
}

func (objid ObjectId) Machine() int32 {
	return objid.machine
}

func (objid ObjectId) Pid() int32 {
	return objid.pid
}

func (objid ObjectId) Increment() int32 {
	return objid.increment & 0xffffff
}

// String returns the ObjectID id as a 24 byte hex string representation.
func (objid ObjectId) String() string {
	array := []byte{
		byte(objid.timestamp >> 0x18),
		byte(objid.timestamp >> 0x10),
		byte(objid.timestamp >> 8),
		byte(objid.timestamp),
		byte(objid.machine >> 0x10),
		byte(objid.machine >> 8),
		byte(objid.machine),
		byte(objid.pid >> 8),
		byte(objid.pid),
		byte(objid.increment >> 0x10),
		byte(objid.increment >> 8),
		byte(objid.increment),
	}
	return hex.EncodeToString(array)
}

func getMachineHash() int32 {
	machineName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	buf := md5.Sum([]byte(machineName))
	return (int32(buf[0])<<0x10 + int32(buf[1])<<8) + int32(buf[2])
}

func getRandomNumber() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}

func tryParse(input string) (objid ObjectId, err error) {
	if len(input) != 0x18 {
		return objid, errors.New("invalid input length")
	}
	array, err := hex.DecodeString(input)
	if err != nil {
		return objid, err
	}
	return ObjectId{
		timestamp: int64(array[0])<<0x18 + int64(array[1])<<0x10 + int64(array[2])<<8 + int64(array[3]),
		machine:   int32(array[4])<<0x10 + int32(array[5])<<8 + int32(array[6]),
		pid:       int32(array[7])<<8 + int32(array[8]),
		increment: int32(array[9])<<0x10 + (int32(array[10]) << 8) + int32(array[11]),
	}, nil
}
