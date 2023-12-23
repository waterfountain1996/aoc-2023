package day20

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type PulseTransfer struct {
	Pulse    bool
	From, To string
}

func NewPulseTransfer(pulse bool, from, to string) *PulseTransfer {
	return &PulseTransfer{
		Pulse: pulse,
		From:  from,
		To:    to,
	}
}

func (p *PulseTransfer) String() string {
	var pulse string
	if p.Pulse {
		pulse = "high"
	} else {
		pulse = "low"
	}

	return fmt.Sprintf("%s -%s-> %s", p.From, pulse, p.To)
}

type FlipFlop bool

func (f FlipFlop) Flip() FlipFlop {
	return !f
}

func (f FlipFlop) Output() bool {
	return bool(f)
}

type Conjunction map[string]bool

func (con Conjunction) Output() bool {
	for _, input := range con {
		// Low pulse only if all inputs were high
		if !input {
			return true
		}
	}
	return false
}

func pressButton(
	modules map[string]interface{},
	inputs, outputs map[string][]string,
	pulses chan<- *PulseTransfer,
) {
	q := []*PulseTransfer{NewPulseTransfer(false, "button" /* from */, "broadcaster" /* to */)}
	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		pulses <- item

		var out bool

		switch m := modules[item.To]; m.(type) {
		case FlipFlop:
			// Flip-flop ignores high pulses
			if item.Pulse {
				continue
			}
			ff := m.(FlipFlop).Flip()
			modules[item.To] = ff
			out = ff.Output()
		case Conjunction:
			con := m.(Conjunction)
			con[item.From] = item.Pulse
			out = con.Output()
		default:
			out = item.Pulse
		}

		for _, dst := range outputs[item.To] {
			q = append(q, NewPulseTransfer(out, item.To /* from */, dst /* to */))
		}
	}
	close(pulses)
}

func PartA(s *bufio.Scanner) string {
	modules := make(map[string]interface{})
	outputs := make(map[string][]string)
	inputs := make(map[string][]string)

	for s.Scan() {
		source, destinations, _ := strings.Cut(s.Text(), " -> ")

		switch source[0] {
		case '%':
			// Flip-flop
			source = source[1:]
			modules[source] = FlipFlop(false)
		case '&':
			// Conjunction
			source = source[1:]
			modules[source] = make(Conjunction)
		default:
			// Broadcaster
			modules[source] = source
		}

		for _, dst := range strings.Split(destinations, ", ") {
			outputs[source] = append(outputs[source], dst)
			inputs[dst] = append(inputs[dst], source)
		}
	}

	for key, values := range inputs {
		switch m := modules[key]; m.(type) {
		case Conjunction:
			con := m.(Conjunction)
			for _, value := range values {
				con[value] = false
			}
		}
	}

	numLow, numHigh := 0, 0

	for i := 0; i < 1000; i++ {
		ch := make(chan *PulseTransfer, 1)

		go pressButton(modules, inputs, outputs, ch)

		for pulse := range ch {
			if pulse.Pulse {
				numHigh++
			} else {
				numLow++
			}
		}
	}

	return strconv.Itoa(numLow * numHigh)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func PartB(s *bufio.Scanner) string {
	modules := make(map[string]interface{})
	outputs := make(map[string][]string)
	inputs := make(map[string][]string)

	for s.Scan() {
		source, destinations, _ := strings.Cut(s.Text(), " -> ")

		switch source[0] {
		case '%':
			// Flip-flop
			source = source[1:]
			modules[source] = FlipFlop(false)
		case '&':
			// Conjunction
			source = source[1:]
			modules[source] = make(Conjunction)
		default:
			// Broadcaster
			modules[source] = source
		}

		for _, dst := range strings.Split(destinations, ", ") {
			outputs[source] = append(outputs[source], dst)
			inputs[dst] = append(inputs[dst], source)
		}
	}

	for key, values := range inputs {
		switch m := modules[key]; m.(type) {
		case Conjunction:
			con := m.(Conjunction)
			for _, value := range values {
				con[value] = false
			}
		}
	}

	var target string

	for _, src := range inputs["rx"] {
		switch modules[src].(type) {
		case Conjunction:
			target = src
			break
		}
	}

	if target == "" {
		return "No conjunction connected to rx"
	}

	cycleStarts := make(map[string]int)
	cycles := make(map[string]int)

	for _, src := range inputs[target] {
		cycleStarts[src] = -1
		cycles[src] = -1
	}

	numPresses := 0

	for {
		numPresses++
		ch := make(chan *PulseTransfer)
		go pressButton(modules, inputs, outputs, ch)
		for pulse := range ch {
			if pulse.Pulse && pulse.To == target {

				if start := cycleStarts[pulse.From]; start != -1 {
					cycles[pulse.From] = numPresses - start
				} else {
					cycleStarts[pulse.From] = numPresses
				}

				result := 1

				for _, num := range cycles {
					if num == -1 {
						result = 0
						break
					}

					result = lcm(result, num)
				}

				if result != 0 {
					return strconv.Itoa(result)
				}
			}
		}
	}
}
