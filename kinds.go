package lottery

import "fmt"

type Kind string

const (
	KindLotto        Kind = "lotto"
	KindEuroMillions Kind = "euromillions"
	KindSetForLife   Kind = "setforlife"
	KindThunderball  Kind = "thunderball"
	KindPowerball    Kind = "powerball"
)

func ParseKind(kind string) (Kind, error) {
	switch kind {
	case "lotto":
		return KindLotto, nil
	case "euromillions":
		return KindEuroMillions, nil
	case "setforlife":
		return KindSetForLife, nil
	case "thunderball":
		return KindThunderball, nil
	case "powerball":
		return KindPowerball, nil
	default:
		return "", fmt.Errorf(
			"invalid lottery kind: %s, must be one of: lotto, euromillions, setforlife, thunderball, powerball",
			kind,
		)
	}
}
