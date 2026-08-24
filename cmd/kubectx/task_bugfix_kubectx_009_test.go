package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx009SourceContract(t *testing.T) {
    source, err := os.ReadFile("shell.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil || choice == \"\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
