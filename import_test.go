package main

import (
	"io"
	"os"
	"testing"

	"github.com/mjl-/mox/mlog"
)

func TestMboxReader(t *testing.T) {
	createTemp := func(pattern string) (*os.File, error) {
		return os.CreateTemp("", pattern)
	}
	mboxf, err := os.Open("testdata/importtest.mbox")
	if err != nil {
		t.Fatalf("open mbox: %v", err)
	}
	defer mboxf.Close()

	mr := newMboxReader(false, false, createTemp, mboxf, mlog.New("mboxreader"))
	_, mf0, _, err := mr.Next()
	if err != nil {
		t.Fatalf("next mbox message: %v", err)
	}
	defer mf0.Close()
	defer os.Remove(mf0.Name())

	_, mf1, _, err := mr.Next()
	if err != nil {
		t.Fatalf("next mbox message: %v", err)
	}
	defer mf1.Close()
	defer os.Remove(mf1.Name())

	_, _, _, err = mr.Next()
	if err != io.EOF {
		t.Fatalf("got err %v, expected eof for next mbox message", err)
	}
}

func TestMaildirReader(t *testing.T) {
	createTemp := func(pattern string) (*os.File, error) {
		return os.CreateTemp("", pattern)
	}
	// todo: rename 1642966915.1.mox to "1642966915.1.mox:2,"? cannot have that name in the git repo because go module (or the proxy) doesn't like it. could also add some flags and test they survive the import.
	newf, err := os.Open("testdata/importtest.maildir/new")
	if err != nil {
		t.Fatalf("open maildir new: %v", err)
	}
	defer newf.Close()

	curf, err := os.Open("testdata/importtest.maildir/cur")
	if err != nil {
		t.Fatalf("open maildir cur: %v", err)
	}
	defer curf.Close()

	mr := newMaildirReader(false, false, createTemp, newf, curf, mlog.New("maildirreader"))
	_, mf0, _, err := mr.Next()
	if err != nil {
		t.Fatalf("next maildir message: %v", err)
	}
	defer mf0.Close()
	defer os.Remove(mf0.Name())

	_, mf1, _, err := mr.Next()
	if err != nil {
		t.Fatalf("next maildir message: %v", err)
	}
	defer mf1.Close()
	defer os.Remove(mf1.Name())

	_, _, _, err = mr.Next()
	if err != io.EOF {
		t.Fatalf("got err %v, expected eof for next maildir message", err)
	}
}
