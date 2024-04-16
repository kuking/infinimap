# InfiniMap

Scale to Hundreds of Terabytes with This Efficient, Memory-like Speed, Persistent Zero-Copy Map.

## Overview

InfiniMap is a high-performance, disk-persistent map (`InfiniMap[K comparable, V any]`) optimized for storing large datasets efficiently. It leverages many
strengths of the operating system internals, it uses zero-copy techniques and memory-mapped files; in practise you can operate with the speed of in-memory 
databases while maintaining the data in disk, and not having to user any process memory as cache.

There is no IO operations required, but to: Create, Open & Close.

## Usage

InfiniMap API is a generic map i.e.

```go
imap, err := OpenOrCreate[uint64, string]("filename.db", NewCreateParameters())

previous, replaced, err := imap.Put(uint64(123), "123")

if value, found := imap.Get(uint64(123)); found {
// value="123"; found=true
}

deleted := imap.Delete(uint64(123))

err = imap.Each(func (k uint64, v string) bool) {
    fmt.Printf("[%v]=%s\n", k, v)
    return true
})

// etc i.e. Count(), Keys(), Values()
```

## Maintenance

If you do a lot of deletes and updates, you will have to do a `Compact(CompactParameters) error` from time to time to reclaim deleted space (if that is an issue
for you), in order to decide when to reclaim deleted space, you can use the APIs `BytesReclaimable() uint64`, `BytesInUse() uint64` `BytesAvailable() uint64`.

If you have added way too may entries going above the initial Capacity (`WithCapacity(int) CreateParameters`) the index might start to clog, and it will be
detrimental to the overall performance; you can checl `ClogRatio() uint8` and run `Reindex() error`. 

### Files sizes vs occupied space

By default, an infinimap will be created of 16TiB file, which is the limit on the typical Linux using EXT4 with 4KB block size.

- Why so big? because it is needed for the memory mapped file backing the data storage.
- But, will it not fill my hard-disk? NO! all modern filesystems (i.e. ext4, zfs, btrfs, etc.) support sparse files, which differentiate between the "declared"
  file-size, and the actual data being used.

You can read in detail about sparse files here: https://wiki.archlinux.org/title/sparse_file.

## Shipping a map file

Given the map files are going to be big (i.e. 16TiB), if you want to ship a map file, as in a gzipped-tar file, you can `Shrink()` it first.
Once deployed, if you are planning to keep adding data you can `Expand()` to give it space to write. But it is not necessary if it is going to be just a
lookup, read-only map.

`Shrink()` and `Expand()` operations are instant. 

You can still compress a 16TB of 'zeroes' which end up in almost a tiny gzipped tar, but it will take time, and upon decompress, and here you get into the
nitty-gritty details of each filesystem implementation, it might actually allocate space for all those zeros. So, better to do shrink & expand for shipping.

When you `Shrink()` an infinimap, you are effectively removing all available space (`BytesAvailable()` will return `0`). Therefore, you won't be able to add
any new entry or update (Delete is OK). Until you `Expand()` the infinimap again.

## Serializer

Data is ultimately stored in disk, therefore it has to be serialised in a consistent way, i.e. if you take a map from a little endian machine (intel/arm) and
open it in a big endian computer, it should work. All basic data-types are serialised in the file `serializer.go` using little endian (the most common 
endianness).  If you want to implement or extend the default serializer, you can set it the API constructor parameter `SetCustomSerializer(Serializer)`

Types supported, all the basic types: `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `float32`, `float64`, `string`, `bool`.
Slices of fixed sized types are supported (basically everything but array of strings, but a string is OK.

If you want to store complex objects, either write your own custom serializer or store it in a JSON, gob, etc.

## Data Structure

All numbers are stored in little endian.

### High level format

| **Name** |
|:--------:|
|  Header  |
| Buckets  |
| Records  |

### Header

| **Name**         | **bits** | **Description**                                     |
|:-----------------|:--------:|:----------------------------------------------------|
| File Version     |    16    | v1=0x1f17                                           |
| Hashing Algo     |    8     | 1=xx128, 2=cityHash128, ...                         |
| Compression algo |    8     | 0=None, 1=LZ4                                       |
| Buckets          |    32    | Number of buckets                                   |
| Next Free Byte   |    64    | int64, where to store the next value                |
| Inserts          |    64    | Number of Inserts                                   |
| Updates          |    64    | Number of Updates                                   |
| Deletes          |    64    | Number of Deletes                                   |
| Count            |    64    | Number of Elements                                  |
| Clog Ratio       |    8     | 0-0xff as ratio of clogging                         |
|                  |          | 0x00 = bucket found on first stop                   |
|                  |          | 0x80 = 50% clog; i.e. visited 1k buckets out of  2k |                           

### Buckets

| **Name** | **bits** | **Description**                |
|:---------|:--------:|:-------------------------------|
| Hash     |   128    | the key hash                   |
| Offset   |    64    | int64 file offset to the value |

### Records

| **Name**   |    **bits**    | **Description**          |
|:-----------|:--------------:|:-------------------------|
| Hash       |      128       | the key hash (again)     |
| Key Size   |       32       | uint32 size of the key   |
| Value Size |       32       | uint32 size of the value |
| Key        |  key size * 8  | the actual key           |
| Value      | value size * 8 | the actual value         |

## Soak test

There is a soak tests that exercises the infinimap against a reference map (a golang one) and checks for consistency, more details here: [SOAK.md](SOAK.md), the
implementation is in the file `cli/soak/main.go`.

## Benchmarks

There are some benchmark results here: [BENCHMARK.md](BENCHMARK.md).

## TODO

- explain why infinimap is better than https://github.com/peterbourgon/diskv
- compact
- compression
