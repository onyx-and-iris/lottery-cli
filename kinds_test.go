package lottery

import (
	"strings"
	"testing"
)

func TestAllKindsReturnsExpectedOrder(t *testing.T) {
	t.Parallel()

	got := AllKinds()
	want := []Kind{
		KindLotto,
		KindEuroMillions,
		KindSetForLife,
		KindThunderball,
		KindPowerball,
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d kinds, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("kind at index %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}

func TestAllKindsReturnsCopy(t *testing.T) {
	t.Parallel()

	first := AllKinds()
	first[0] = "mutated"

	again := AllKinds()
	if again[0] != KindLotto {
		t.Fatalf("expected AllKinds to return a defensive copy")
	}
}

func TestParseKindAcceptsAllKinds(t *testing.T) {
	t.Parallel()

	for _, kind := range AllKinds() {
		parsed, err := ParseKind(string(kind))
		if err != nil {
			t.Fatalf("expected no error for %q, got %v", kind, err)
		}
		if parsed != kind {
			t.Fatalf("expected %q, got %q", kind, parsed)
		}
	}
}

func TestParseKindIncludesAllKindsInError(t *testing.T) {
	t.Parallel()

	_, err := ParseKind("invalid")
	if err == nil {
		t.Fatalf("expected an error for invalid kind")
	}

	message := err.Error()
	for _, kind := range AllKinds() {
		if !strings.Contains(message, string(kind)) {
			t.Fatalf("expected error message to contain kind %q: %s", kind, message)
		}
	}
}
