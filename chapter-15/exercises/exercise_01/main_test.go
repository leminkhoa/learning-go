package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestParser_ValidInput(t *testing.T) {
	data := []byte("id1\n+\n3\n4\n")
	input, err := parser(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	want := Input{Id: "id1", Op: "+", Val1: 3, Val2: 4}
	if !reflect.DeepEqual(input, want) {
		t.Errorf("got %+v, want %+v", input, want)
	}
}

func TestParser_InvalidInt(t *testing.T) {
	data := []byte("id1\n+\nfoo\n4\n")
	_, err := parser(data)
	if err == nil {
		t.Error("expected error for invalid int, got nil")
	}
}

func TestParser_MissingLines(t *testing.T) {
	data := []byte("id1\n+\n3\n") // only 3 lines
	_, err := parser(data)
	if err == nil {
		t.Error("expected error for missing lines, got nil")
	}
}

func TestDataProcessor_ValidOps(t *testing.T) {
	tests := []struct {
		name   string
		op     string
		v1, v2 int
		want   int
	}{
		{"add", "+", 2, 3, 5},
		{"sub", "-", 5, 2, 3},
		{"mul", "*", 2, 4, 8},
		{"div", "/", 8, 2, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan []byte, 1)
			out := make(chan Result, 1)
			input := Input{Id: "id", Op: tt.op, Val1: tt.v1, Val2: tt.v2}
			data := []byte(strings.Join([]string{input.Id, input.Op, strconv.Itoa(input.Val1), strconv.Itoa(input.Val2)}, "\n") + "\n")
			in <- data
			close(in)
			DataProcessor(in, out)
			res, ok := <-out
			if !ok {
				t.Fatal("expected result, got closed channel")
			}
			if res.Value != tt.want {
				t.Errorf("got %d, want %d", res.Value, tt.want)
			}
		})
	}
}

func TestDataProcessor_InvalidInput(t *testing.T) {
	in := make(chan []byte, 1)
	out := make(chan Result, 1)
	in <- []byte("bad\ndata\n") // not enough lines
	close(in)
	DataProcessor(in, out)
	_, ok := <-out
	if ok {
		t.Error("expected no output for invalid input")
	}
}

func TestDataProcessor_UnknownOp(t *testing.T) {
	in := make(chan []byte, 1)
	out := make(chan Result, 1)
	input := Input{Id: "id", Op: "%", Val1: 1, Val2: 2}
	data := []byte(strings.Join([]string{input.Id, input.Op, strconv.Itoa(input.Val1), strconv.Itoa(input.Val2)}, "\n") + "\n")
	in <- data
	close(in)
	DataProcessor(in, out)
	_, ok := <-out
	if ok {
		t.Error("expected no output for unknown op")
	}
}

func TestWriteData(t *testing.T) {
	in := make(chan Result, 2)
	var buf bytes.Buffer
	in <- Result{Id: "id1", Value: 42}
	in <- Result{Id: "id2", Value: 7}
	close(in)
	WriteData(in, &buf)
	got := buf.String()
	want := "id1:42\nid2:7\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParser_ZeroLines(t *testing.T) {
	_, err := parser([]byte(""))
	if err == nil {
		t.Error("expected error for zero lines, got nil")
	}
}

func TestParser_OneLine(t *testing.T) {
	_, err := parser([]byte("id1\n"))
	if err == nil {
		t.Error("expected error for one line, got nil")
	}
}

func TestParser_TwoLines(t *testing.T) {
	_, err := parser([]byte("id1\n+\n"))
	if err == nil {
		t.Error("expected error for two lines, got nil")
	}
}

// Custom writer that always returns an error

type errorWriter struct{}

func (e *errorWriter) Write(p []byte) (int, error) {
	return 0, io.ErrClosedPipe
}

func TestWriteData_ErrorWriter(t *testing.T) {
	in := make(chan Result, 1)
	in <- Result{Id: "id1", Value: 1}
	close(in)
	errW := &errorWriter{}
	WriteData(in, errW) // Should not panic or error
}

// HTTP handler tests

func TestNewController_Accepted(t *testing.T) {
	ch := make(chan []byte, 1)
	h := NewController(ch)
	req := httptest.NewRequest("POST", "/", io.NopCloser(strings.NewReader("id1\n+\n1\n2\n")))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("got %d, want %d", resp.StatusCode, http.StatusAccepted)
	}
}
