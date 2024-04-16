# Benchmarks

Below there are some benchmark results, the comparison in unfair: the golang map is the most efficient map in memory. InfiniMaps are persistent and can store
more than then system memory worth of data, etc. There is roughly a X2 performance cost for them. i.e. for having a 5TB lookup table, in a 64GB system.

```shell
$ make bench                                                                                                                                                                                                                                                                 130 ↵
go test -run=Benchmark -bench=. -benchmem ./...
goos: linux
goarch: amd64
pkg: github.com/kuking/infinimap/V1
cpu: AMD Ryzen 7 3800X 8-Core Processor             
BenchmarkInt32ToInt32InfiniMapPut-16    	 2517903	       429.7 ns/op	      16 B/op	       4 allocs/op
BenchmarkInt32ToInt32GoMapPut-16        	 8370592	       193.1 ns/op	      47 B/op	       0 allocs/op
BenchmarkInt32ToInt32InfiniMapGet-16    	 6621902	       152.8 ns/op	      15 B/op	       3 allocs/op
BenchmarkInt32ToInt32GoMapGet-16        	40999048	        29.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkInt64ToInt64InfiniMapPut-16    	 1600198	       650.2 ns/op	      32 B/op	       4 allocs/op
BenchmarkInt64ToInt64GoMapPut-16        	 5396748	       191.4 ns/op	      66 B/op	       0 allocs/op
BenchmarkInt64ToInt64InfiniMapGet-16    	 7599711	       163.8 ns/op	      31 B/op	       3 allocs/op
BenchmarkInt64ToInt64GoMapGet-16        	39356384	        30.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrToStrInfiniMapPut-16        	 1062675	      1051 ns/op	      95 B/op	       5 allocs/op
BenchmarkStrToStrGoMapPut-16            	 2722738	       379.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkStrToStrInfiniMapGet-16        	 3024544	       456.0 ns/op	      71 B/op	       6 allocs/op
BenchmarkStrToStrGoMapGet-16            	17524879	        80.68 ns/op	       3 B/op	       0 allocs/op
PASS
ok  	github.com/kuking/infinimap/V1	48.784s
PASS
ok  	github.com/kuking/infinimap/cli/soak	0.002s
```