# single-flight

> https://mp.weixin.qq.com/s/TE7zu2t2SjUpGKK-Bji9_g

使用singleflight可以减少对Redis等DB的访问频率，但是和文中得到的结果不同，基准测试发现并没有得到更高的效率..

```
$ go test   -benchtime=1000000x -count=1 -benchmem  -bench=. -cpu=1,2,4
goos: windows
goarch: amd64
pkg: github.com/wangjc/gopractice/singleflight
cpu: 13th Gen Intel(R) Core(TM) i5-13500H
BenchmarkBufferWithPool          1000000              1365 ns/op             145 B/op          2 allocs/op
BenchmarkBufferWithPool-2        1000000               774.2 ns/op            96 B/op          2 allocs/op
BenchmarkBufferWithPool-4        1000000               304.3 ns/op            96 B/op          2 allocs/op
BenchmarkBufferWithPool2         1000000              1366 ns/op              38 B/op          2 allocs/op
ok      github.com/wangjc/gopractice/singleflight       6.063s
```
