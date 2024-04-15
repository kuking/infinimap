package main

import (
	"fmt"
	"github.com/kuking/infinimap/V1"
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
	log.Printf("running for %v minute(s).\n", mins)
	log.Println("(If you want other value, just call this with a numeric parameter.)")
	tempFile, err := os.CreateTemp(os.TempDir(), "infinimap-*.db")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up: delete the temporary file after the test

	imap, err := V1.Create[uint64, string](tempFile.Name(),
		V1.NewCreateParameters().WithCapacity(25_000_000))
	defer imap.Close() // Ensure the map is closed on program exit

	reference := make(map[uint64]string) // Reference map to validate InfiniMap operations

	soak(imap, reference, time.Duration(mins)*60*time.Second)

	verify(imap, reference)
}

func verify(imap V1.InfiniMap[uint64, string], reference map[uint64]string) {
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

func soak(imap V1.InfiniMap[uint64, string], reference map[uint64]string, duration time.Duration) {
	ops := 0
	insertCount := 0
	startTime := time.Now()
	lastLog := time.Now()
	for time.Since(startTime) < duration {
		ops, insertCount = doRandomOperation(imap, reference, ops, insertCount)
		gets := uint64(ops) - imap.StatsDeletes() - imap.StatsInserts() - imap.StatsUpdates()
		if time.Since(lastLog) > 15*time.Second {
			lastLog = time.Now()
			mill := 1_000_000.0
			gig := 1024.0 * 1024.0 * 1024.0
			log.Printf("[%.f%%] %.2fM ops, %.2fM entries: %.2fM inserts, %.2fM updates, %.2fM deletes, %.2fM gets, %.1f%% clog\n",
				float64(time.Now().Sub(startTime)*100.0/duration), float64(ops)/mill, float64(imap.Count())/mill,
				float64(imap.StatsInserts())/mill, float64(imap.StatsUpdates())/mill, float64(imap.StatsDeletes())/mill, float64(gets)/mill,
				float32(imap.ClogRatio())/255.0)
			log.Printf("  ... disk space: %.1fG allocated, %.1fG in use, %.1fG reclaimable, %.1fG available\n",
				float64(imap.BytesAllocated())/gig, float64(imap.BytesInUse())/gig, float64(imap.BytesReclaimable())/gig, float64(imap.BytesAvailable())/gig)

		}
	}
}

func doRandomOperation(imap V1.InfiniMap[uint64, string], reference map[uint64]string, ops int, insertCount int) (int, int) {
	operation := rand.Intn(4)
	if imap.Count() < 1_000_000 {
		operation = 0 // insert!
	} else if imap.Count() > 10_000_000 {
		operation = 2 // delete!
	}
	switch operation {
	case 0: // Insert
		key := uint64(insertCount)
		if _, exist := reference[key]; !exist {
			value := fmt.Sprintf("Value %d", key)
			reference[key] = value
			_, _, err := imap.Put(key, value)
			if err != nil {
				log.Fatalln(err)
			}
			insertCount++
			ops++
		}
	case 1: // Update
		key := uint64(rand.Intn(insertCount))
		oldRef, found := reference[key]
		if found {
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
			ops++
		}
	case 2: // Delete
		key := uint64(rand.Intn(insertCount))
		if _, found := reference[key]; found {
			delete(reference, key)
			if !imap.Delete(key) {
				log.Println("BUG? it did not delete the key:", key)
			}
			ops++
		}
	case 3: // Get
		key := uint64(rand.Intn(insertCount))
		refValue, refFound := reference[key]
		imapValue, imapFound := imap.Get(key)
		if refFound != imapFound {
			log.Println("BUG! reference has for key,", key, "is", refFound, "but imap", imapFound)
		}
		if refFound && imapFound && refValue != imapValue {
			log.Println("BUG! the expected value was:", refValue, "but we found:", imapValue)
		}
		ops++
	}
	return ops, insertCount
}
