manvendra@Manvendras-MBP scripts % go build -gcflags=-m go-closure.go
# command-line-arguments
./go-closure.go:7:6: can inline Sphere
./go-closure.go:9:12: can inline Sphere.func1
./go-closure.go:22:16: inlining call to Sphere
./go-closure.go:9:12: can inline main.func1
./go-closure.go:25:42: inlining call to main.func1
./go-closure.go:25:13: inlining call to fmt.Println
./go-closure.go:26:42: inlining call to main.func1
./go-closure.go:26:13: inlining call to fmt.Println
./go-closure.go:8:6: moved to heap: radius
./go-closure.go:9:12: func literal escapes to heap
./go-closure.go:22:16: func literal does not escape
./go-closure.go:25:13: ... argument does not escape
./go-closure.go:25:14: "Volume of Sphere is:" escapes to heap
./go-closure.go:25:42: ~R0 escapes to heap
./go-closure.go:26:13: ... argument does not escape
./go-closure.go:26:14: "Volume of Sphere is:" escapes to heap
./go-closure.go:26:42: ~R0 escapes to heap
manvendra@Manvendras-MBP scripts %
