package errorc_test

import (
	"errors"
	"fmt"
	"github.com/youjiaxing/errorc"
)

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

func fn() error {
	e1 := errorc.New("error")
	e2 := errorc.Wrap(e1, "inner")
	e3 := errorc.WrapC(e2, 1, "middle")
	return errorc.WrapC(e3, 2, "outer")
}

func ExampleCause() {
	err := fn()
	fmt.Println(err)
	fmt.Println(errorc.Cause(err))

	// Output:
	// outer
	// error
}

func ExampleUnwrap() {
	err := fn()
	for err != nil {
		fmt.Println(err)
		err = errors.Unwrap(err)
	}

	// Output:
	// outer
	// middle
	// inner
	// error
}

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

func ExampleStackTrace() {
	type stackTracer interface {
		StackTrace() errorc.StackTrace
	}

	err, ok := errorc.Cause(fn()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v", st[0:2]) // top two frames

	// Output:
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:69
	// github.com/youjiaxing/errorc_test.ExampleStackTrace
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:309
}

func ExampleNewC_extended_minus_v() {
	err := errorc.NewC(1, "whoops: %s", "foo")
	fmt.Printf("%-v", err)

	// Output:
	// code: 1, msg: whoops: foo
	// github.com/youjiaxing/errorc_test.ExampleNewC_extended_minus_v
	//	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:325
}

func ExampleWrap_extended_minus_v() {
	err := fn()
	fmt.Printf("%-v\n", err)

	// Output:
	// error
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:69
	// inner
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:70
	// code: 1, msg: middle
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:71
	// code: 2, msg: outer
	// github.com/youjiaxing/errorc_test.fn
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:72
}

func ExampleWithMessage_extend_minus_v() {
	e1 := errors.New("raw error")
	e2 := errorc.Wrap(e1, "wrap error")
	e3 := errorc.WithMessage(e2, "with message error")
	e4 := errorc.WithCode(e3, 100, "with code")
	e5 := errorc.WithMessage(e4, "with message")
	fmt.Printf("%-v", e5)

	// Output:
	// raw error
	// wrap error
	// github.com/youjiaxing/errorc_test.ExampleWithMessage_extend_minus_v
	// 	/Users/youjiaxing/Documents/GitHub/youjiaxing/errorc/example_test.go:355
	// with message error
	// code: 100, msg: with code
	// code: 100, msg: with message
}
