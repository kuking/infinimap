# InfiniMap

Scale to Hundreds of Terabytes with This Efficient, Persistent Zero-Copy Map.

## Data Structure

### High level format

| **Name** |
|:--------:|
|  Header  |
| Buckets  |
| Records  |

### Header

| **Name**         | **bits** | **Description**                      |
|:-----------------|:--------:|:-------------------------------------|
| File Version     |    16    | v1=0x1f17                            |
| Hashing Algo     |    8     | 1=xx128, 2=cityHash128, ...          |
| Compression algo |    8     | 0=None, 1=LZ4                        |
| Buckets          |    32    | Number of buckets                    |
| Next Free Byte   |    64    | int64, where to store the next value |
| Inserts          |    64    | Number of Inserts                    |
| Deletes          |    64    | Number of Deletes                    |
| Count            |    64    | Number of Elements                   |

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

# TODO

- https://github.com/peterbourgon/diskv