package lottery

import (
	"fmt"
	"slices"
	"strings"
)

type Kind string

const (
	KindLotto        Kind = "lotto"
	KindEuroMillions Kind = "euromillions"
	KindSetForLife   Kind = "setforlife"
	KindThunderball  Kind = "thunderball"
	KindPowerball    Kind = "powerball"
)

var allKinds = []Kind{
	KindLotto,
	KindEuroMillions,
	KindSetForLife,
	KindThunderball,
	KindPowerball,
}

var allKindsText = []string{
	string(KindLotto),
	string(KindEuroMillions),
	string(KindSetForLife),
	string(KindThunderball),
	string(KindPowerball),
}

var allKindsCSV = strings.Join(allKindsText, ", ")

// AllKinds returns the complete, ordered list of supported lottery kinds.
//
// The returned slice is a copy and can be modified safely by callers.
func AllKinds() []Kind {
	kinds := make([]Kind, len(allKinds))
	copy(kinds, allKinds)
	return kinds
}

// ParseKind parses a string into a Kind, returning an error if the string is not a valid kind.
func ParseKind(kind string) (Kind, error) {
	parsed := Kind(kind)
	if slices.Contains(allKinds, parsed) {
		return parsed, nil
	}

	return "", fmt.Errorf(
		"invalid lottery kind: %s, must be one of: %s",
		kind,
		allKindsCSV,
	)
}
