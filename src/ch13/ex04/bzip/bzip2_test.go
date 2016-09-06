// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata

package bzip_test

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"sync"
	"testing"

	"ch13/ex04/bzip" // writer
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	//+ Exercise 13.4
	w, err := bzip.NewWriter(&compressed)

	if err != nil {
		t.Fatalf("%v", err)
	}
	//- Exercise 13.4

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

func TestBzip2InConcurrently(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	//+ Exercise 13.4
	w, err := bzip.NewWriter(&compressed)

	if err != nil {
		t.Fatalf("%v", err)
	}
	//- Exercise 13.4

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(&uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}

	//+ Exercise 13.3
	tee = io.MultiWriter(w)
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			io.WriteString(tee, "hello") // current writing
			wg.Done()
		}()
	}
	wg.Wait()
	//- Exercise 13.3

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

//- Exercise 13.3
//- Exercise 13.3
