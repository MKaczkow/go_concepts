# Pointer Performance Test Results
* `cpu`: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz

| Benchmark                     | Operations | Time per Operation |
|-------------------------------|------------|-------------------|
| BenchmarkPointer10In-8       | 1000000000 | 1.025 ns/op       |
| BenchmarkValue10In-8         | 945092488  | 1.192 ns/op       |
| BenchmarkPointer10Out-8      | 65293329   | 16.98 ns/op       |
| BenchmarkValue10Out-8        | 191989155  | 6.189 ns/op       |
| BenchmarkPointer100In-8      | 1000000000 | 1.006 ns/op       |
| BenchmarkValue100In-8        | 350708606  | 3.343 ns/op       |
| BenchmarkPointer100Out-8     | 34208500   | 33.21 ns/op       |
| BenchmarkValue100Out-8       | 100000000  | 10.28 ns/op       |
| BenchmarkPointer1_000In-8    | 953362309  | 1.205 ns/op       |
| BenchmarkValue1_000In-8      | 53018286   | 20.16 ns/op       |
| BenchmarkPointer1_000Out-8   | 6488534    | 177.1 ns/op       |
| BenchmarkValue1_000Out-8     | 17232514   | 60.88 ns/op       |
| BenchmarkPointer100_000In-8  | 1000000000 | 1.216 ns/op       |
| BenchmarkValue100_000In-8    | 526632     | 2213 ns/op        |
| BenchmarkPointer100_000Out-8 | 132204     | 9156 ns/op        |
| BenchmarkPointer10_000_000In-8 | 1000000000 | 1.070 ns/op     |
| BenchmarkValue10_000_000In-8 | 1080       | 1074963 ns/op     |
| BenchmarkPointer10_000_000Out-8 | 969       | 1089116 ns/op   |  
| BenchmarkValue10_000_000Out-8 | 572        | 1976812 ns/op    |

* `result`: PASS