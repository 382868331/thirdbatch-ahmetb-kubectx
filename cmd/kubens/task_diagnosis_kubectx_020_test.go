package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKubectx020SourceContract(t *testing.T) {
    source, err := os.ReadFile("statefile.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return true") {
        t.Fatalf("expected source contract is missing")
    }
}
