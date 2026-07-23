package lottery

import (
	"fmt"
	"math/rand"
)

// Lottery is an interface that defines the methods that all lottery games must implement.
type Lottery interface {
	Draw()
}

// Lotto represents the Lotto lottery game.
// It draws 6 numbers from a pool of 1 to 59 without replacement.
// There is no bonus ball in this game.
type Lotto struct {
	Numbers [6]int
}

func (l *Lotto) Draw() {
	numbers := drawUnique(len(l.Numbers), 59)
	for i := range l.Numbers {
		l.Numbers[i] = numbers[i]
	}
}

// EuroMillions represents the EuroMillions lottery game.
// It draws 5 numbers from a pool of 1 to 50 without replacement.
// Additionally, it draws 2 Lucky Stars from a pool of 1 to 12 without replacement.
type EuroMillions struct {
	Numbers    [5]int
	LuckyStars [2]int
}

func (e *EuroMillions) Draw() {
	numbers := drawUnique(len(e.Numbers), 50)
	for i := range e.Numbers {
		e.Numbers[i] = numbers[i]
	}

	luckyStars := drawUnique(len(e.LuckyStars), 12)
	for i := range e.LuckyStars {
		e.LuckyStars[i] = luckyStars[i]
	}
}

// SetForLife represents the Set For Life lottery game.
// It draws 5 numbers from a pool of 1 to 47 without replacement.
// Additionally, it draws 1 Life Ball from a pool of 1 to 10 without replacement.
type SetForLife struct {
	Numbers  [5]int
	LifeBall int
}

func (s *SetForLife) Draw() {
	numbers := drawUnique(len(s.Numbers), 47)
	for i := range s.Numbers {
		s.Numbers[i] = numbers[i]
	}
	s.LifeBall = rand.Intn(10) + 1
}

// Thunderball represents the Thunderball lottery game.
// It draws 5 numbers from a pool of 1 to 39 without replacement.
// Additionally, it draws 1 Thunderball from a pool of 1 to 14 without replacement.
type Thunderball struct {
	Numbers     [5]int
	Thunderball int
}

func (t *Thunderball) Draw() {
	numbers := drawUnique(len(t.Numbers), 39)
	for i := range t.Numbers {
		t.Numbers[i] = numbers[i]
	}
	t.Thunderball = rand.Intn(14) + 1
}

// Powerball represents the Powerball lottery game.
// It draws 5 numbers from a pool of 1 to 69 without replacement.
// Additionally, it draws 1 Powerball from a pool of 1 to 26 without replacement.
type Powerball struct {
	Numbers   [5]int
	Powerball int
}

func (p *Powerball) Draw() {
	numbers := drawUnique(len(p.Numbers), 69)
	for i := range p.Numbers {
		p.Numbers[i] = numbers[i]
	}
	p.Powerball = rand.Intn(26) + 1
}

// New returns a new instance of the selected lottery game.
func New(selected Kind) (Lottery, error) {
	switch selected {
	case KindLotto:
		return &Lotto{}, nil
	case KindEuroMillions:
		return &EuroMillions{}, nil
	case KindSetForLife:
		return &SetForLife{}, nil
	case KindThunderball:
		return &Thunderball{}, nil
	case KindPowerball:
		return &Powerball{}, nil
	default:
		return nil, fmt.Errorf("unknown lottery type: %s", string(selected))
	}
}
