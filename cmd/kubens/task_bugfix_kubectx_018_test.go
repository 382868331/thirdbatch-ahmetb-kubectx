package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx018SourceContract(t *testing.T) {
    source, err := os.ReadFile("flags.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "name = argv[1]") {
        t.Fatalf("expected source contract is missing")
    }
}
