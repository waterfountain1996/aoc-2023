package day7

import (
	"bufio"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const CamelCards = "AKQJT98765432"
const JokerCamelCards = "AKQT98765432J"

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func cardStrength(card string) int {
	return strings.LastIndex(CamelCards, card)
}

func jokerCardStrength(card string) int {
	return strings.LastIndex(JokerCamelCards, card)
}

type Hand struct {
	Cards string
	Bid   int
}

func handFromString(s string) Hand {
	tokens := strings.Split(s, " ")
	bid, _ := strconv.Atoi(tokens[1])
	return Hand{
		Cards: tokens[0],
		Bid:   bid,
	}
}

func (h *Hand) String() string {
	var typeString string
	switch h.JokerType() {
	case HighCard:
		typeString = "High Card"
	case OnePair:
		typeString = "One Pair"
	case TwoPair:
		typeString = "Two Pair"
	case ThreeOfAKind:
		typeString = "Three of a Kind"
	case FullHouse:
		typeString = "Full House"
	case FourOfAKind:
		typeString = "Four of a Kind"
	case FiveOfAKind:
		typeString = "Five of a Kind"
	}

	return fmt.Sprintf("%s (%s)", h.Cards, typeString)
}

func (h *Hand) Type() HandType {
	charCount := make(map[rune]int)
	for _, ch := range h.Cards {
		charCount[ch]++
	}

	if len(charCount) == 1 {
		return FiveOfAKind
	} else if len(charCount) == 2 {
		for _, v := range charCount {
			if v == 1 || v == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	} else if len(charCount) == 3 {
		for _, v := range charCount {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	}

	for _, v := range charCount {
		if v == 2 {
			return OnePair
		}
	}

	return HighCard
}

func (h *Hand) JokerType() HandType {
	hasJokers := false
	charCount := make(map[rune]int)
	for _, ch := range h.Cards {
		charCount[ch]++
		if ch == 'J' {
			hasJokers = true
		}
	}

	if !hasJokers {
		return h.Type()
	}

	r := HighCard

	for card := range charCount {
		newHand := Hand{
			Cards: strings.ReplaceAll(h.Cards, "J", string(card)),
			Bid:   h.Bid,
		}
		if newHand.Type() > r {
			r = newHand.Type()
		}
	}

	return r
}

func (h *Hand) Cmp(other Hand) int {
	if h.Type() > other.Type() {
		return 1
	} else if h.Type() == other.Type() {
		for idx := 0; idx < len(h.Cards); idx++ {
			r := cmp.Compare(
				cardStrength(string(other.Cards[idx])),
				cardStrength(string(h.Cards[idx])),
			)

			if r != 0 {
				return r
			}
		}

	} else {
		return -1
	}

	return 0
}

func (h *Hand) CmpJoker(other Hand) int {
	if h.JokerType() > other.JokerType() {
		return 1
	} else if h.JokerType() == other.JokerType() {
		for idx := 0; idx < len(h.Cards); idx++ {
			r := cmp.Compare(
				jokerCardStrength(string(other.Cards[idx])),
				jokerCardStrength(string(h.Cards[idx])),
			)

			if r != 0 {
				return r
			}
		}

	} else {
		return -1
	}

	return 0
}

func PartA(s *bufio.Scanner) string {
	hands := []Hand{}

	for s.Scan() {
		hands = append(hands, handFromString(s.Text()))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Cmp(b)
	})

	r := 0
	for rank, hand := range hands {
		r += (rank + 1) * hand.Bid
	}

	return strconv.Itoa(r)
}

func PartB(s *bufio.Scanner) string {
	hands := []Hand{}

	for s.Scan() {
		hands = append(hands, handFromString(s.Text()))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.CmpJoker(b)
	})

	r := 0
	for rank, hand := range hands {
		fmt.Printf("%s\n", hand.String())
		r += (rank + 1) * hand.Bid
	}

	return strconv.Itoa(r)
}
