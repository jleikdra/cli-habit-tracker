package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRunAddAndList(t *testing.T) {
	tmp := t.TempDir()
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	// add habit
	os.Args = []string{"habit", "add", "testhabit"}
	if code := run(); code != 0 {
		t.Fatalf("add returned code %d", code)
	}

	// capture stdout for ls
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"habit", "ls"}
	if code := run(); code != 0 {
		w.Close()
		os.Stdout = oldStdout
		t.Fatalf("ls returned code %d", code)
	}

	w.Close()
	var b strings.Builder
	io.Copy(&b, r)
	os.Stdout = oldStdout

	out := b.String()
	if !strings.Contains(out, "testhabit") {
		t.Fatalf("ls output did not contain habit: %q", out)
	}
}
