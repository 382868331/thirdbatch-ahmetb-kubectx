package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKubectx012SourceContract(t *testing.T) {
    source, err := os.ReadFile("rename.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if op.Old == \".\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
