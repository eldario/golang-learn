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

#### Updated results without exclude first or last words

```json
ID  Word    Count  Order  
1   that    36     13     
2   from    58     23     
3   like    35     58     
4   told    29     78     
5   looked  39     92     
6   went    62     146    
7   love    32     188    
8   want    29     196    
9   into    29     204    
10  down    28     207    

```

#### Update result with exclude list

```json
ID  Word    Count  Order  
1   like    35     56     
2   told    29     76     
3   looked  39     90     
4   marry   21     133    
5   went    62     139    
6   love    32     175    
7   want    29     182    
8   into    29     188    
9   took    20     218    
10  cant    18     335    
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

pkg: tasks/task3/pkg/simpleMapper
BenchmarkInsert-4               41022320                27.8 ns/op
BenchmarkGetFrequentUses-4      31736749                38.0 ns/op
PASS
ok      tasks/task3/pkg/simpleMapper    3.387s

pkg: tasks/task3/pkg/mapper/simple
BenchmarkInsert-4               38757088                26.9 ns/op
BenchmarkRemove-4               196112100                5.95 ns/op
BenchmarkGetResults-4           19765978                58.9 ns/op
PASS
ok      tasks/task3/pkg/mapper/simple   4.094s

```