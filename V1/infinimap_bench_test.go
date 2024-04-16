package V1

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func BenchmarkInt32ToInt32InfiniMapPut(b *testing.B) {
	b.SetParallelism(1)

	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap-*")
	defer deferredCleanup(tempFile)

	imap, _ := Create[uint32, uint32](tempFile.Name(), NewCreateParameters())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = imap.Put(rand.Uint32(), rand.Uint32())
	}
	b.StopTimer()
	_ = imap.Close()
	deferredCleanup(tempFile)
}

func BenchmarkInt32ToInt32GoMapPut(b *testing.B) {
	m := map[uint32]uint32{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[rand.Uint32()] = rand.Uint32()
	}
	b.StopTimer()
}

func BenchmarkInt32ToInt32InfiniMapGet(b *testing.B) {
	b.SetParallelism(1)

	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap-*")
	defer deferredCleanup(tempFile)

	imap, _ := Create[uint32, uint32](tempFile.Name(), NewCreateParameters())
	for i := 0; i < 10_000; i++ {
		_, _, _ = imap.Put(uint32(i), rand.Uint32())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = imap.Get(rand.Uint32() % 10_000)
	}
	b.StopTimer()
	_ = imap.Close()
	deferredCleanup(tempFile)
}

func BenchmarkInt32ToInt32GoMapGet(b *testing.B) {
	m := map[uint32]uint32{}
	for i := 0; i < 10_000; i++ {
		m[uint32(i)] = rand.Uint32()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[rand.Uint32()%10_000]
	}
	b.StopTimer()
}

func BenchmarkInt64ToInt64InfiniMapPut(b *testing.B) {
	b.SetParallelism(1)

	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap-*")
	defer deferredCleanup(tempFile)

	imap, _ := Create[uint64, uint64](tempFile.Name(), NewCreateParameters())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = imap.Put(rand.Uint64(), rand.Uint64())
	}
	b.StopTimer()
	_ = imap.Close()
	deferredCleanup(tempFile)
}

func BenchmarkInt64ToInt64GoMapPut(b *testing.B) {
	m := map[uint64]uint64{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[rand.Uint64()] = rand.Uint64()
	}
	b.StopTimer()
}

func BenchmarkInt64ToInt64InfiniMapGet(b *testing.B) {
	b.SetParallelism(1)

	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap-*")
	defer deferredCleanup(tempFile)

	imap, _ := Create[uint64, uint64](tempFile.Name(), NewCreateParameters())
	for i := 0; i < 10_000; i++ {
		_, _, _ = imap.Put(uint64(i), rand.Uint64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = imap.Get(rand.Uint64() % 10_000)
	}
	b.StopTimer()
	_ = imap.Close()
	deferredCleanup(tempFile)
}

func BenchmarkInt64ToInt64GoMapGet(b *testing.B) {
	m := map[uint64]uint64{}
	for i := 0; i < 10_000; i++ {
		m[uint64(i)] = rand.Uint64()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[rand.Uint64()%10_000]
	}
	b.StopTimer()
}
func BenchmarkStrToStrInfiniMapPut(b *testing.B) {
	b.SetParallelism(1)

	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap-*")
	defer deferredCleanup(tempFile)

	imap, _ := Create[string, string](tempFile.Name(), NewCreateParameters())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := strconv.Itoa(rand.Int())
		_, _, _ = imap.Put(v, v)
	}
	b.StopTimer()
	_ = imap.Close()
	deferredCleanup(tempFile)
}

func BenchmarkStrToStrGoMapPut(b *testing.B) {
	m := map[string]string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := strconv.Itoa(rand.Int())
		m[v] = v
	}
	b.StopTimer()
}

func BenchmarkStrToStrInfiniMapGet(b *testing.B) {
	b.SetParallelism(1)

	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap-*")
	defer deferredCleanup(tempFile)

	imap, _ := Create[string, string](tempFile.Name(), NewCreateParameters())
	for i := 0; i < 10_000; i++ {
		_, _, _ = imap.Put(strconv.Itoa(i), strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = imap.Get(strconv.Itoa(rand.Int() % 10_000))
	}
	b.StopTimer()
	_ = imap.Close()
	deferredCleanup(tempFile)
}

func BenchmarkStrToStrGoMapGet(b *testing.B) {
	m := map[string]string{}
	for i := 0; i < 10_000; i++ {
		m[strconv.Itoa(i)] = strconv.Itoa(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[strconv.Itoa(rand.Int()%10_000)]
	}
	b.StopTimer()
}
