#### Some results

```json
[
    were:     92
    that:     126
    from:     68
    with:     128
    scarlett: 146 
    about:    59 
    said:     200
    going:    51
    melanie:  63 
    went:     62 
]
```

#### Updated results

```json
ID  Word      Count  Order
1   were      92     10
2   that      126    13
3   from      68     23
4   with      128    25
5   scarlett  146    31
6   about     59     60
7   said      200    69
8   going     51     73
9   melanie   63     124
10  went      62     145   
```

#### Benchmarks

```json
BenchmarkInsert-4                4746104               262 ns/op
BenchmarkGetFrequentUses-4      12489802               100 ns/op
```

#### Updated

```json
pkg: tasks/task3/pkg/hardMapper
BenchmarkInsert-4                4777845               277 ns/op
BenchmarkGetFrequentUses-4      11041116               103 ns/op
PASS
ok      tasks/task3/pkg/hardMapper      2.830s
go test -bench=. ./pkg/simpleMapper
goos: linux
goarch: amd64
pkg: tasks/task3/pkg/simpleMapper
BenchmarkInsert-4               41022320                27.8 ns/op
BenchmarkGetFrequentUses-4      31736749                38.0 ns/op
PASS
ok      tasks/task3/pkg/simpleMapper    3.387s

```