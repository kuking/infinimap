package main

import (
	"fmt"
	"github.com/kuking/infinimap"
	"github.com/kuking/infinimap/impl"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	println("infinimap: soak-test")
	mins := 1
	if len(os.Args) == 2 {
		if m, err := strconv.Atoi(os.Args[1]); err == nil {
			mins = m
		}
	}
	log.Printf("%v: soak test\n", os.Args[0])
	log.Printf("running for %v minutes (if you want other value, just call this with a numeric parameter.)\n", mins)
	tempFile, err := ioutil.TempFile("", "infinimap-*.db")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up: delete the temporary file after the test

	imap, err := impl.CreateInfinimap[uint64, string](tempFile.Name(),
		impl.NewCreateParameters().WithCapacity(5_000_000)) //XXX
	defer imap.Close() // Ensure the map is closed on program exit

	reference := make(map[uint64]string) // Reference map to validate InfiniMap operations
	soak(imap, reference, time.Duration(mins)*60*time.Second)
	verify(imap, reference)
}

func verify(imap infinimap.InfiniMap[uint64, string], reference map[uint64]string) {
	ok := true
	log.Println("Verifying contents ...")
	for k, v := range reference {
		iv, found := imap.Get(k)
		if !found {
			log.Println("Missing key in imap:", k)
			ok = false
		}
		if v != iv {
			log.Println("Incorrect value in imap:", iv, "expected:", v)
			ok = false
		}
	}

	err := imap.Each(func(k uint64, v string) (cont bool) {
		rv, found := reference[k]
		if !found {
			log.Println("Missing key in reference:", k)
			ok = false
		}
		if v != rv {
			log.Println("Incorrect value in reference:", rv, "expected:", v)
			ok = false
		}
		return true
	})
	if err != nil {
		log.Println(err)
	}
	if ok {
		log.Println("SUCCESS! All values matched between reference map and infinimap")
	}
}

func soak(imap infinimap.InfiniMap[uint64, string], reference map[uint64]string, duration time.Duration) {
	ops := 0
	startTime := time.Now()
	lastLog := time.Now()
	for time.Since(startTime) < duration {
		doRandomOperation(imap, reference, ops)
		ops++
		if time.Since(lastLog) > 15*time.Second {
			lastLog = time.Now()
			log.Printf("[%.f%%] %.1fM ops, %d count, %d inserts, %d updates, %d deletes, %.1f%% clog\n",
				float32(time.Now().Sub(startTime)*100.0/duration), float32(ops)/1_000_000.0, imap.Count(),
				imap.StatsInserts(), imap.StatsUpdates(), imap.StatsDeletes(), float32(imap.ClogRatio())/255.0)
		}
	}
}

func doRandomOperation(imap infinimap.InfiniMap[uint64, string], reference map[uint64]string, ops int) {
	operation := rand.Intn(4)
	if len(reference) < 1_000_000 {
		operation = 0
	}
	switch operation {
	case 0: // Insert
		key := uint64(ops)
		if _, exist := reference[key]; !exist {
			value := fmt.Sprintf("Value %d", key)
			reference[key] = value
			_, _, err := imap.Put(key, value)
			if err != nil {
				log.Fatalln(err)
			}
		}
	case 1: // Update
		key := uint64(rand.Intn(ops))
		oldRef, found := reference[key]
		value := fmt.Sprintf("Value %d Updated with rand %d", key, rand.Int())
		reference[key] = value
		oldImap, replaced, err := imap.Put(key, value)
		if replaced != found {
			log.Println("BUG? it did not replaced the key:", key)
		}
		if oldRef != oldImap {
			log.Println("BUG? old value replaced:", oldRef, "do not match with imap old value:", oldImap)
		}
		if err != nil {
			log.Fatalln(err)
		}
	case 2: // Delete
		key := uint64(rand.Intn(ops))
		if _, found := reference[key]; found {
			delete(reference, key)
			if !imap.Delete(key) {
				log.Println("BUG? it did not delete the key:", key)
			}
		}
	case 3: // Get
		key := uint64(rand.Intn(ops))
		refValue, refFound := reference[key]
		imapValue, imapFound := imap.Get(key)
		if refFound != imapFound {
			log.Println("BUG! reference has for key,", key, "is", refFound, "but imap", imapFound)
		}
		if refFound && imapFound && refValue != imapValue {
			log.Println("BUG! the expected value was:", refValue, "but we found:", imapValue)
		}
	}
}
