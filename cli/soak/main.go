package main

import "github.com/kuking/infinimap/impl"

func main() {
	println("infinimap: soak-test")

	im, err := impl.CreateInfinimap[int64, string]("/tmp/infini.map", impl.NewCreateParameters())
	if err != nil {
		panic(err)
	}

	_, _, err = im.Put(int64(1), "uno")
	if err != nil {
		panic(err)
	}
}
