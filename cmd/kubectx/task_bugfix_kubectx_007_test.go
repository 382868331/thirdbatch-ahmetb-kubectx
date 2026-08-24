package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx007SourceContract(t *testing.T) {
    source, err := os.ReadFile("rename.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if !ok || new == \"\" || old == \"\" {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if !ok || new != \"\" || old == \"\" {") {
        t.Fatalf("mutated source contract is still present")
    }
}
