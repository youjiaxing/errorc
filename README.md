# errorc
error with code package for golang

# apis
## New
```go
func ExampleNew() {
	err := errorc.New("whoops: %s", "foo")
	fmt.Println(err)

	// Output:
	// whoops: foo
}

func ExampleNew_extended() {
	err := errorc.New("whoops: %s", "foo")
	fmt.Printf("%+v", err)

	// Output:
	// code: 0, msg: whoops: foo
	// github.com/youjiaxing/errorc_test.ExampleNew_extended
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:18
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
}
```

## NewC
```go
func ExampleNewC() {
	err := errorc.NewC(1, "whoops: %s", "foo")
	fmt.Println(err)

	// Output: whoops: foo
}

func ExampleNewC_extended() {
	err := errorc.NewC(1, "whoops: %s", "foo")
	fmt.Printf("%+v", err)

	// Output:
	// code: 1, msg: whoops: foo
	// github.com/youjiaxing/errorc_test.ExampleNewC_extended
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:47
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
}

func ExampleNewC_extended_minus_v() {
err := errorc.NewC(1, "whoops: %s", "foo")
fmt.Printf("%-v", err)

// Output:
// code: 1, msg: whoops: foo
// github.com/youjiaxing/errorc_test.ExampleNewC_extended_minus_v
//	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:325
}
```

## Wrap
```go
func ExampleWrap() {
	cause := errorc.New("whoops")
	err := errorc.Wrap(cause, "oh noes #%d", 2)
	fmt.Println(err)

	// Output: oh noes #2
}

func ExampleWrap_extended() {
	err := fn()
	fmt.Printf("%+v\n", err)

	// Output:
	// code: 0, msg: error
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:69
	// github.com/youjiaxing/errorc_test.ExampleWrap_extended
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:108
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
	// code: 0, msg: inner
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:70
	// github.com/youjiaxing/errorc_test.ExampleWrap_extended
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:108
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
	// code: 1, msg: middle
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:71
	// github.com/youjiaxing/errorc_test.ExampleWrap_extended
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:108
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
	// code: 2, msg: outer
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:72
	// github.com/youjiaxing/errorc_test.ExampleWrap_extended
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:108
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
}
```

## WrapC
```go
func ExampleWrapC() {
	cause := errorc.New("whoops")
	err := errorc.WrapC(cause, 1, "oh noes #%d", 2)
	fmt.Println(err)

	// Output: oh noes #2
}

func ExampleWrapC_extend() {
	cause := errorc.New("whoops")
	err := errorc.WrapC(cause, 1, "oh noes #%d", 2)
	fmt.Printf("%+v", err)

	// Output:
	// code: 0, msg: whoops
	// github.com/youjiaxing/errorc_test.ExampleWrapC_extend
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:191
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
	// code: 1, msg: oh noes #2
	// github.com/youjiaxing/errorc_test.ExampleWrapC_extend
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:192
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
}
```

## WithCode
```go
func ExampleWithCode() {
	e1 := errors.New("raw error")
	e2 := errorc.Wrap(e1, "wrap error")
	e3 := errorc.WithCode(e2, 100, "with code error")
	fmt.Println(e3)

	// Output:
	// with code error
}

func ExampleWithCode_extend() {
	e1 := errors.New("raw error")
	e2 := errorc.Wrap(e1, "wrap error")
	e3 := errorc.WithCode(e2, 100, "with code error")
	fmt.Printf("%+v", e3)

	// Output:
	// raw error
	// code: 0, msg: wrap error
	// github.com/youjiaxing/errorc_test.ExampleWithCode_extend
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:240
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
	// code: 100, msg: with code error
}
```

## WithMessage
```go
func ExampleWithMessage() {
	e1 := errors.New("raw error")
	e2 := errorc.Wrap(e1, "wrap error")
	e3 := errorc.WithMessage(e2, "with message error")
	fmt.Println(e3)

	// Output:
	// with message error
}

func ExampleWithMessage_extend() {
	e1 := errors.New("raw error")
	e2 := errorc.Wrap(e1, "wrap error")
	e3 := errorc.WithMessage(e2, "with message error")
	e4 := errorc.WithCode(e3, 100, "with code")
	e5 := errorc.WithMessage(e4, "with message")
	fmt.Printf("%+v", e5)

	// Output:
	// raw error
	// code: 0, msg: wrap error
	// github.com/youjiaxing/errorc_test.ExampleWithMessage_extend
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:276
	// testing.runExample
	// 	/usr/local/opt/go/libexec/src/testing/run_example.go:63
	// testing.runExamples
	// 	/usr/local/opt/go/libexec/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/opt/go/libexec/src/testing/testing.go:1728
	// main.main
	// 	_testmain.go:119
	// runtime.main
	// 	/usr/local/opt/go/libexec/src/runtime/proc.go:250
	// runtime.goexit
	// 	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594
	// code: 0, msg: with message error
	// code: 100, msg: with code
	// code: 100, msg: with message
}
```