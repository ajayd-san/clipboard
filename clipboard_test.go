// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard_test

import (
	"os"
	"testing"

	. "github.com/atotto/clipboard"
)

func TestCopyAndPaste(t *testing.T) {
	InitClipboard(os.Environ())
	expected := "日本語"

	err := WriteAll(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestMultiCopyAndPaste(t *testing.T) {
	InitClipboard(os.Environ())
	expected1 := "French: éèêëàùœç"
	expected2 := "Weird UTF-8: 💩☃"

	err := WriteAll(expected1)
	if err != nil {
		t.Fatal(err)
	}

	actual1, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if actual1 != expected1 {
		t.Errorf("want %s, got %s", expected1, actual1)
	}

	err = WriteAll(expected2)
	if err != nil {
		t.Fatal(err)
	}

	actual2, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if actual2 != expected2 {
		t.Errorf("want %s, got %s", expected2, actual2)
	}
}

func BenchmarkReadAll(b *testing.B) {
	InitClipboard(os.Environ())
	for i := 0; i < b.N; i++ {
		ReadAll()
	}
}

func BenchmarkWriteAll(b *testing.B) {
	InitClipboard(os.Environ())
	text := "いろはにほへと"
	for i := 0; i < b.N; i++ {
		WriteAll(text)
	}
}
