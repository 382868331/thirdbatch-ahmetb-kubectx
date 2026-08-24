package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx019SourceContract(t *testing.T) {
    source, err := os.ReadFile("current.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer kc.Close()") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "kc.Close()") {
        t.Fatalf("mutated source contract is still present")
    }
}
