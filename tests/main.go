package main

import (
	"log"
  "time"
  "strconv"
	"github.com/czxcjx/go-hashmap/hashmap"
)

func printTimeElapsed(start time.Time) {
  log.Printf("Took %s\n", time.Since(start))
}

func testSetGetDelete(capacity int, setCount int) {
  defer printTimeElapsed(time.Now())
  log.Printf("Testing %d Sets, Gets, and Deletes for a %d capacity hashmap", setCount, capacity)
  h:= hashmap.New(capacity)
  for i := 0; i < setCount; i++ {
    h.Set(strconv.Itoa(i), i)
  }
  for i := 0; i < setCount; i++ {
    h.Get(strconv.Itoa(i))
  }
  for i := 0; i < setCount; i++ {
    h.Delete(strconv.Itoa(i))
  }
}

func main() {
  testSetGetDelete(100, 20)
  testSetGetDelete(1000, 200)
  testSetGetDelete(10000, 2000)
  testSetGetDelete(100000, 20000)
  testSetGetDelete(1000000, 200000)
  testSetGetDelete(10000000, 2000000)

  testSetGetDelete(100, 50)
  testSetGetDelete(1000, 500)
  testSetGetDelete(10000, 5000)
  testSetGetDelete(100000, 50000)
  testSetGetDelete(1000000, 500000)
  testSetGetDelete(10000000, 5000000)

  testSetGetDelete(100, 90)
  testSetGetDelete(1000, 900)
  testSetGetDelete(10000, 9000)
  testSetGetDelete(100000, 90000)
  testSetGetDelete(1000000, 900000)
  testSetGetDelete(10000000, 9000000)
}
