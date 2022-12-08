package errorc

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"testing"
)

func TestFormatNew(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		New("error"),
		"%s",
		"error",
	}, {
		New("error"),
		"%v",
		"error",
	}, {
		New("error"),
		"%+v",
		"code: 0, msg: error\n" +
			"github.com/youjiaxing/errorc.TestFormatNew\n" +
			"\t.+/youjiaxing/errorc/format_test.go:26",
	}, {
		New("error"),
		"%q",
		`"error"`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatNewC(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		NewC(1, "error"),
		"%s",
		"error",
	}, {
		NewC(1, "error"),
		"%v",
		"error",
	}, {
		NewC(1, "error"),
		"%+v",
		"code: 1, msg: error\n" +
			"github.com/youjiaxing/errorc.TestFormatNew\n" +
			"\t.+/youjiaxing/errorc/format_test.go:56",
	}, {
		New("error"),
		"%q",
		`"error"`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatWrap(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		Wrap(New("error"), "error2"),
		"%s",
		"error2",
	}, {
		Wrap(New("error"), "error2"),
		"%v",
		"error2",
	}, {
		Wrap(New("error"), "error2"),
		"%+v",
		"code: 0, msg: error\n" +
			"github.com/youjiaxing/errorc.TestFormatWrap\n" +
			"\t.+/youjiaxing/errorc/format_test.go:86",
	}, {
		Wrap(io.EOF, "error"),
		"%s",
		"error",
	}, {
		Wrap(io.EOF, "error"),
		"%v",
		"error",
	}, {
		Wrap(io.EOF, "error"),
		"%+v",
		"EOF\n" +
			"code: 0, msg: error\n" +
			"github.com/youjiaxing/errorc.TestFormatWrap\n" +
			"\t.+/youjiaxing/errorc/format_test.go:100",
	}, {
		Wrap(Wrap(io.EOF, "error1"), "error2"),
		"%+v",
		"EOF\n" +
			"code: 0, msg: error1\n" +
			"github.com/youjiaxing/errorc.TestFormatWrap\n" +
			"\t.+/youjiaxing/errorc/format_test.go:107\n" +
			"testing.tRunner\n" +
			"\t.+/src/testing/testing.go:1446\n" +
			"runtime.goexit\n" +
			"\t.+/src/runtime/asm_amd64.s:1594\n" +
			"code: 0, msg: error2\n" +
			"github.com/youjiaxing/errorc.TestFormatWrap\n" +
			"\t.+/youjiaxing/errorc/format_test.go:107\n" +
			"testing.tRunner\n" +
			"\t.+/src/testing/testing.go:1446\n" +
			"runtime.goexit\n" +
			"\t.+/src/runtime/asm_amd64.s:1594",
	}, {
		Wrap(New("error with space"), "context"),
		"%q",
		`"context"`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatWrapC(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		WrapC(New("error"), 1, "error2"),
		"%s",
		"error2",
	}, {
		WrapC(New("error"), 1, "error2"),
		"%v",
		"error2",
	}, {
		WrapC(New("error"), 1, "error2"),
		"%+v",
		"code: 0, msg: error\n" +
			"github.com/youjiaxing/errorc.TestFormatWrapC\n" +
			"\t.+/youjiaxing/errorc/format_test.go:149\n" +
			"testing.tRunner\n" +
			"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
			"runtime.goexit\n" +
			"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594\n" +
			"code: 1, msg: error2\n" +
			"github.com/youjiaxing/errorc.TestFormatWrapC\n" +
			"\t.+/youjiaxing/errorc/format_test.go:149\n" +
			"testing.tRunner\n" +
			"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
			"runtime.goexit\n" +
			"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
	}, {
		WrapC(io.EOF, 1, "error"),
		"%s",
		"error",
	}, {
		WrapC(io.EOF, 1, "error"),
		"%v",
		"error",
	}, {
		WrapC(io.EOF, 1, "error"),
		"%+v",
		"EOF\n" +
			"code: 1, msg: error\n" +
			"github.com/youjiaxing/errorc.TestFormatWrapC\n" +
			"\t.+/youjiaxing/errorc/format_test.go:174",
	}, {
		WrapC(WrapC(io.EOF, 1, "error1"), 2, "error2"),
		"%+v",
		"EOF\n" +
			"code: 1, msg: error1\n" +
			"github.com/youjiaxing/errorc.TestFormatWrapC\n" +
			"\t.+/youjiaxing/errorc/format_test.go:181\n" +
			"testing.tRunner\n" +
			"\t.+/src/testing/testing.go:1446\n" +
			"runtime.goexit\n" +
			"\t.+/src/runtime/asm_amd64.s:1594\n" +
			"code: 2, msg: error2\n" +
			"github.com/youjiaxing/errorc.TestFormatWrapC\n" +
			"\t.+/youjiaxing/errorc/format_test.go:181\n" +
			"testing.tRunner\n" +
			"\t.+/src/testing/testing.go:1446\n" +
			"runtime.goexit\n" +
			"\t.+/src/runtime/asm_amd64.s:1594",
	}, {
		WrapC(New("error with space"), 1, "context"),
		"%q",
		`"context"`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatWithMessage(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{{
		WithMessage(io.EOF, "with message"),
		"%s",
		[]string{"with message"},
	}, {
		WithMessage(io.EOF, "with message"),
		"%v",
		[]string{"with message"},
	}, {
		WithMessage(io.EOF, "with message"),
		"%+v",
		[]string{"EOF",
			"code: 0, msg: with message",
			"github.com/youjiaxing/errorc.TestFormatWithMessage\n" +
				"\t.+/youjiaxing/errorc/format_test.go:223"},
	}, {
		WithMessage(New("error"), "with message"),
		"%s",
		[]string{"with message"},
	}, {
		WithMessage(New("error"), "with message"),
		"%v",
		[]string{"with message"},
	}, {
		WithMessage(New("error"), "with message"),
		"%+v",
		[]string{"code: 0, msg: error",
			"github.com/youjiaxing/errorc.TestFormatWithMessage\n" +
				"\t.+/youjiaxing/errorc/format_test.go:238\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 0, msg: with message"},
	}, {
		WithMessage(WithMessage(io.EOF, "with message"), "with message2"),
		"%+v",
		[]string{"EOF",
			"code: 0, msg: with message",
			"github.com/youjiaxing/errorc.TestFormatWithMessage\n" +
				"\t.+/youjiaxing/errorc/format_test.go:249\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 0, msg: with message2",
		},
	}, {
		WithMessage(WithMessage(Wrap(io.EOF, "message"), "with message"), "with message2"),
		"%+v",
		[]string{"EOF",
			"code: 0, msg: message",
			"github.com/youjiaxing/errorc.TestFormatWithMessage\n" +
				"\t.+/youjiaxing/errorc/format_test.go:262\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 0, msg: with message",
			"code: 0, msg: with message2",
		},
	}, {
		WithMessage(New("error%d", 1), "with message"),
		"%+v",
		[]string{"code: 0, msg: error1",
			"github.com/youjiaxing/errorc.TestFormatWithMessage\n" +
				"\t.+/youjiaxing/errorc/format_test.go:276\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 0, msg: with message"},
	}}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func TestFormatWithCode(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{{
		WithCode(New("error"), 1, "error2"),
		"%s",
		[]string{"error2"},
	}, {
		WithCode(New("error"), 1, "error2"),
		"%v",
		[]string{"error2"},
	}, {
		WithCode(New("error"), 1, "error2"),
		"%+v",
		[]string{
			"code: 0, msg: error",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:307\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 1, msg: error2",
		},
	}, {
		WithCode(io.EOF, 1, "addition1"),
		"%s",
		[]string{"addition1"},
	}, {
		WithCode(io.EOF, 1, "addition1"),
		"%v",
		[]string{"addition1"},
	}, {
		WithCode(io.EOF, 1, "addition1"),
		"%+v",
		[]string{
			"EOF",
			"code: 1, msg: addition1",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:328\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
		},
	}, {
		WithCode(WithCode(io.EOF, 1, "addition1"), 2, "addition2"),
		"%v",
		[]string{"addition2"},
	}, {
		WithCode(WithCode(io.EOF, 1, "addition1"), 2, "addition2"),
		"%+v",
		[]string{
			"EOF",
			"code: 1, msg: addition1",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:345\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 2, msg: addition2",
		},
	}, {
		Wrap(WithCode(io.EOF, 1, "error1"), "error2"),
		"%+v",
		[]string{
			"EOF",
			"code: 1, msg: error1",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:359\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 1, msg: error2",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:359\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
		},
	}, {
		WithCode(WithMessage(io.EOF, "EOF"), 1, "error"),
		"%+v",
		[]string{
			"EOF",
			"code: 0, msg: EOF",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:379\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 1, msg: error",
		},
	}, {
		WithCode(Wrap(WithMessage(io.EOF, "EOF"), "inside-error"), 1, "outside-error"),
		"%+v",
		[]string{
			"EOF",
			"code: 0, msg: EOF",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:393\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 0, msg: inside-error",
			"github.com/youjiaxing/errorc.TestFormatWithCode\n" +
				"\t.+/youjiaxing/errorc/format_test.go:393\n" +
				"testing.tRunner\n" +
				"\t/usr/local/opt/go/libexec/src/testing/testing.go:1446\n" +
				"runtime.goexit\n" +
				"\t/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1594",
			"code: 1, msg: outside-error"},
	}}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func testFormatRegexp(t *testing.T, n int, arg interface{}, format, want string) {
	t.Helper()
	got := fmt.Sprintf(format, arg)
	gotLines := strings.SplitN(got, "\n", -1)
	wantLines := strings.SplitN(want, "\n", -1)

	if len(wantLines) > len(gotLines) {
		t.Errorf("test %d: wantLines(%d) > gotLines(%d):\n got: %q\nwant: %q", n+1, len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		match, err := regexp.MatchString(w, gotLines[i])
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("test %d: line %d: fmt.Sprintf(%q, err):\n got: %q\nwant: %q", n+1, i+1, format, got, want)
		}
	}
}

var stackLineR = regexp.MustCompile(`\.`)

// parseBlocks parses input into a slice, where:
//   - incase entry contains a newline, its a stacktrace
//   - incase entry contains no newline, its a solo line.
//
// Detecting stack boundaries only works incase the WithStack-calls are
// to be found on the same line, thats why it is optionally here.
//
// Example use:
//
//	for _, e := range blocks {
//	  if strings.ContainsAny(e, "\n") {
//	    // Match as stack
//	  } else {
//	    // Match as line
//	  }
//	}
func parseBlocks(input string, detectStackboundaries bool) ([]string, error) {
	var blocks []string

	stack := ""
	wasStack := false
	lines := map[string]bool{} // already found lines

	for _, l := range strings.Split(input, "\n") {
		isStackLine := stackLineR.MatchString(l)

		switch {
		case !isStackLine && wasStack:
			blocks = append(blocks, stack, l)
			stack = ""
			lines = map[string]bool{}
		case isStackLine:
			if wasStack {
				// Detecting two stacks after another, possible cause lines match in
				// our tests due to WithStack(WithStack(io.EOF)) on same line.
				if detectStackboundaries {
					if lines[l] {
						if len(stack) == 0 {
							return nil, errors.New("len of block must not be zero here")
						}

						blocks = append(blocks, stack)
						stack = l
						lines = map[string]bool{l: true}
						continue
					}
				}

				stack = stack + "\n" + l
			} else {
				stack = l
			}
			lines[l] = true
		case !isStackLine && !wasStack:
			blocks = append(blocks, l)
		default:
			return nil, errors.New("must not happen")
		}

		wasStack = isStackLine
	}

	// Use up stack
	if stack != "" {
		blocks = append(blocks, stack)
	}
	return blocks, nil
}

func testFormatCompleteCompare(t *testing.T, n int, arg interface{}, format string, want []string, detectStackBoundaries bool) {
	gotStr := fmt.Sprintf(format, arg)

	got, err := parseBlocks(gotStr, detectStackBoundaries)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(want) {
		t.Fatalf("test %d: fmt.Sprintf(%s, err) -> wrong number of blocks: got(%d) want(%d)\n got: %s\nwant: %s\ngotStr: %q",
			n+1, format, len(got), len(want), prettyBlocks(got), prettyBlocks(want), gotStr)
	}

	for i := range got {
		if strings.ContainsAny(want[i], "\n") {
			// Match as stack
			match, err := regexp.MatchString(want[i], got[i])
			if err != nil {
				t.Fatal(err)
			}
			if !match {
				t.Fatalf("test %d: block %d: fmt.Sprintf(%q, err):\ngot:\n%q\nwant:\n%q\nall-got:\n%s\nall-want:\n%s\n",
					n+1, i+1, format, got[i], want[i], prettyBlocks(got), prettyBlocks(want))
			}
		} else {
			// Match as message
			if got[i] != want[i] {
				t.Fatalf("test %d: fmt.Sprintf(%s, err) at block %d got != want:\n got: %q\nwant: %q", n+1, format, i+1, got[i], want[i])
			}
		}
	}
}

type wrapper struct {
	wrap func(err error) error
	want []string
}

func prettyBlocks(blocks []string) string {
	var out []string

	for _, b := range blocks {
		out = append(out, fmt.Sprintf("%v", b))
	}

	return "   " + strings.Join(out, "\n   ")
}

func testGenericRecursive(t *testing.T, beforeErr error, beforeWant []string, list []wrapper, maxDepth int) {
	if len(beforeWant) == 0 {
		panic("beforeWant must not be empty")
	}
	for _, w := range list {
		if len(w.want) == 0 {
			panic("want must not be empty")
		}

		err := w.wrap(beforeErr)

		// Copy required cause append(beforeWant, ..) modified beforeWant subtly.
		beforeCopy := make([]string, len(beforeWant))
		copy(beforeCopy, beforeWant)

		beforeWant := beforeCopy
		last := len(beforeWant) - 1
		var want []string

		// Merge two stacks behind each other.
		if strings.ContainsAny(beforeWant[last], "\n") && strings.ContainsAny(w.want[0], "\n") {
			want = append(beforeWant[:last], append([]string{beforeWant[last] + "((?s).*)" + w.want[0]}, w.want[1:]...)...)
		} else {
			want = append(beforeWant, w.want...)
		}

		testFormatCompleteCompare(t, maxDepth, err, "%+v", want, false)
		if maxDepth > 0 {
			testGenericRecursive(t, err, want, list, maxDepth-1)
		}
	}
}
