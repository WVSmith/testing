package main

import (
	"fmt"
	"github.com/couchbase/go-couchbase"
	//"github.com/davecgh/go-spew/spew"
	"log"
	"strconv"
)

func main() {
	c, err := couchbase.Connect("http://ec2-52-36-89-236.us-west-2.compute.amazonaws.com:8091/")
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	pool, err := c.GetPool("default")
	if err != nil {
		log.Fatalf("Error getting pool:  %v", err)
	}

	TestWrite(pool)
}

func TestRead(pool couchbase.Pool) {
	threads := 50
	bucket, err := pool.GetBucket("beer-sample")
	if err != nil {
		log.Fatalf("Error getting bucket:  %v", err)
	}

	work := make(chan string, 100)
	quit := make(chan struct{})

	var b []byte

	//bucket.Set("someKey", 0, []string{"an", "example", "list"})

	go func() {
		for i := 0; i < 100; i++ {
			work <- string("brasserie_la_choulette-framboise")
		}
		close(work)
	}()

	for i := 0; i < threads; i++ {
		go func() {
			var err error
			for id, ok := <-work; ok; id, ok = <-work {
				b, _, _, err = bucket.GetsRaw(id)
				if err != nil {
					log.Fatalf("Error getting doc: %v", err)
				}
			}
			quit <- struct{}{}
		}()
	}

	select {
	case <-quit:
		log.Println("stoping")
	}

	fmt.Println(string(b))
}

func TestWrite(pool couchbase.Pool) {
	threads := 50
	bucket, err := pool.GetBucket("writetest")
	if err != nil {
		log.Fatalf("Error getting bucket:  %v", err)
	}

	work := make(chan int, 100)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			work <- i
		}
		close(work)
	}()

	for i := 0; i < threads; i++ {
		go func() {
			for id, ok := <-work; ok; id, ok = <-work {
				bucket.Set(strconv.Itoa(id), 0, []string{"an", "example", "list"})
			}
			quit <- struct{}{}
		}()
	}

	for i := 0; i < threads; i++ {
		select {
		case <-quit:
		}
	}

	log.Println("stoping")
}
