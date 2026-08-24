package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx010SourceContract(t *testing.T) {
    source, err := os.ReadFile("fzf.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := checkIsolatedMode(); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
