Testing out https://blog.cloudflare.com/a-deep-dive-into-bpf-lpm-trie-performance-and-optimization/

### Improvements over normal binary trie implimention

- Multibit tries: Instead of 1 rune/bit per step, compare multiple runes per node. This reduces height.

### Future Improvements:

- Path Compression: Merge sparse nodes (e.g., if a node has only one child, skip to the leaf).
- Level Compression : For dense levels, use a single node with many children (e.g., a map or array for all possible next keys)

### Results

```
$ go test ./multibit -bench=BenchmarkTrieLookup -benchmem
goos: darwin
goarch: arm64
pkg: go-trie/multibit
cpu: Apple M2 Pro
BenchmarkTrieLookup-12          25385341                40.50 ns/op            0 B/op          0 allocs/op
PASS
ok      go-trie/multibit        1.402s


$ go test ./binary -bench=BenchmarkTrieLookup -benchmem
goos: darwin
goarch: arm64
pkg: go-trie/binary
cpu: Apple M2 Pro
BenchmarkTrieLookup-12          11645674                91.46 ns/op            0 B/op          0 allocs/op
PASS
ok      go-trie/binary  1.509s
```
