package rest

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestToBool(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected bool
		ok       bool
	}{
		{
			test:     "true, ok",
			input:    "true",
			expected: true,
			ok:       true,
		},
		{
			test:     "false, ok",
			input:    "false",
			expected: false,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "someting else",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToBool(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %t - want: %t", actual, tc.expected)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected float32
		ok       bool
	}{
		{
			test:     "-0.1, ok",
			input:    "-0.1",
			expected: -0.1,
			ok:       true,
		},
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "0.1, ok",
			input:    "0.1",
			expected: 0.1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "-03.01, ok",
			input:    "-03.01",
			expected: -3.01,
			ok:       true,
		},
		{
			test:     "03.01, ok",
			input:    "03.01",
			expected: 3.01,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "someting else",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToFloat32(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %f - want: %f", actual, tc.expected)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected float64
		ok       bool
	}{
		{
			test:     "-0.1, ok",
			input:    "-0.1",
			expected: -0.1,
			ok:       true,
		},
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "0.1, ok",
			input:    "0.1",
			expected: 0.1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "-03.01, ok",
			input:    "-03.01",
			expected: -3.01,
			ok:       true,
		},
		{
			test:     "03.01, ok",
			input:    "03.01",
			expected: 3.01,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "someting else",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToFloat64(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %f - want: %f", actual, tc.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected string
		ok       bool
	}{
		{
			test:     "test, ok",
			input:    "test",
			expected: "test",
			ok:       true,
		},
		{
			test:     "not ok",
			input:    "",
			expected: "",
			ok:       false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToString(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %s - want: %s", actual, tc.expected)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected int
		ok       bool
	}{
		{
			test:     "-1, ok",
			input:    "-1",
			expected: -1,
			ok:       true,
		},
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "1, ok",
			input:    "1",
			expected: 1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "-03, ok",
			input:    "-03",
			expected: -3,
			ok:       true,
		},
		{
			test:     "03, ok",
			input:    "03",
			expected: 3,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "2.1",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToInt(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %d - want: %d", actual, tc.expected)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected int32
		ok       bool
	}{
		{
			test:     "-1, ok",
			input:    "-1",
			expected: -1,
			ok:       true,
		},
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "1, ok",
			input:    "1",
			expected: 1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "-03, ok",
			input:    "-03",
			expected: -3,
			ok:       true,
		},
		{
			test:     "03, ok",
			input:    "03",
			expected: 3,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "2.1",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToInt32(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %d - want: %d", actual, tc.expected)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected int64
		ok       bool
	}{
		{
			test:     "-1, ok",
			input:    "-1",
			expected: -1,
			ok:       true,
		},
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "1, ok",
			input:    "1",
			expected: 1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "-03, ok",
			input:    "-03",
			expected: -3,
			ok:       true,
		},
		{
			test:     "03, ok",
			input:    "03",
			expected: 3,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "2.1",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToInt64(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %d - want: %d", actual, tc.expected)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected uint
		ok       bool
	}{
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "1, ok",
			input:    "1",
			expected: 1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "02, ok",
			input:    "02",
			expected: 2,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "-1",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToUint(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %d- want: %d", actual, tc.expected)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected uint32
		ok       bool
	}{
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "1, ok",
			input:    "1",
			expected: 1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "02, ok",
			input:    "02",
			expected: 2,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "-1",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToUint32(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %d - want: %d", actual, tc.expected)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected uint64
		ok       bool
	}{
		{
			test:     "0, ok",
			input:    "0",
			expected: 0,
			ok:       true,
		},
		{
			test:     "1, ok",
			input:    "1",
			expected: 1,
			ok:       true,
		},
		{
			test:     "2, ok",
			input:    "2",
			expected: 2,
			ok:       true,
		},
		{
			test:     "02, ok",
			input:    "02",
			expected: 2,
			ok:       true,
		},
		{
			test:  "not ok",
			input: "-1",
			ok:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToUint64(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if actual != tc.expected {
				t.Fatalf("Got: %d - want: %d", actual, tc.expected)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	testcases := []struct {
		test     string
		input    string
		expected []byte
		ok       bool
	}{
		{
			test:     "test, ok",
			input:    base64.URLEncoding.EncodeToString([]byte("test")),
			expected: []byte("test"),
			ok:       true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			actual, ok := ToBytes(tc.input)

			if ok != tc.ok {
				t.Fatalf("Got: %t - want: %t", ok, tc.ok)
			}

			if bytes.Compare(actual, tc.expected) != 0 {
				t.Fatalf("Got: %v - want: %v", actual, tc.expected)
			}
		})
	}
}
