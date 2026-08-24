package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx003SourceContract(t *testing.T) {
    source, err := os.ReadFile("delete.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := kc.Save(); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
