package main

import "fmt"
import "encoding/binary"
import "github.com/google/uuid"

func main() {

	u1 := uuid.New()
	//u1 := use_uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	l1 := binary.BigEndian.Uint64(u1[:8])
	l2 := binary.BigEndian.Uint64(u1[8:])

	fmt.Printf("%x %x\n", l1, l2)
	fmt.Printf("%v %v\n", l1, l2)
}
