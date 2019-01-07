package main

import (
	"fmt"
	"math/big"
	"sort"
)

/*
	Explore the space of Turing Machines
	of n states and 2 symbols.
	Print the results per machine.
*/

// States of the turing machines (n)
const STATES = 2

// MAX_NUMBER_OF_STEPS will halt the machine, discarding its output
const MAX_NUMBER_OF_STEPS = 108

// We will be always dealing with two symbols, and no more than 7 states
var SYMBOLPICK = []string{"0", "1"}
var STATEPICK = []string{"A", "B", "C", "D", "E", "F", "G"}

// Useful for sanity checks
var MAX_MACHINES = new(big.Int).Exp(big.NewInt(4*STATES+2), big.NewInt(2*STATES), nil)

// TuringMachine of n states and 2 symbols
type TuringMachine struct {
	states          int
	transitionTable map[TTKey]TTValue
	tape            *Tape
	currentState    string
}

// TTKey is a key of the transition table
type TTKey struct {
	state  string
	symbol string
}

// TTValue is a value of the transition table
type TTValue struct {
	symbol    string
	state     string
	direction string
}

// Tape is the Turing Machine's tape
type Tape struct {
	values       []string
	headPosition int
}

// newTuringMachine returns a (n,2) turing machine with a Tape initialized
// at `0`, and a currentState of `A`
func newTuringMachine(states int, transitionIndex *big.Int) *TuringMachine {
	tm := &TuringMachine{}

	tm.states = states

	tm.updateTransitionTableByIndex(transitionIndex)

	tm.tape = &Tape{
		values:       []string{"0"},
		headPosition: 0,
	}

	tm.currentState = "A"

	return tm
}

// getTransitionTableByIndex updates the tm.transitionTable
// using an index the following way:
// - We convert the index number into a number base `4n + 2`,
//   that is, `2(2n + 1)`, since each state can move left or right,
//   the halting state won't move, and either `0` or `1` symbols
//   can be printed.
// - The resulting `4n + 2` base number is _zero padded_ to `2n`.
// - Each pair `(state, symbol)` in the table is assigned with one of
//   the base `4n + 2` digits of the index number, subsequently, these
//   digits are mapped to a tuple `(state, symbol, direction)`.
// - Note that the tuples containing the halting state, will not
//   get a direction value assigned.
func (tm *TuringMachine) updateTransitionTableByIndex(index *big.Int) {
	// sanity check
	if index.Cmp(MAX_MACHINES) > 0 {
		panic("max index of TM reached!")
	}

	// there are `2n` keys in the transition table
	transitionKeys := 2 * tm.states
	// there are `4n + 2` available transitions
	availableTransitions := 4*tm.states + 2

	// convert the number into a base `4n + 2` representation
	indexStr := index.Text(availableTransitions)
	padLength := transitionKeys - len(indexStr)
	for i := 0; i < padLength; i++ {
		indexStr = "0" + indexStr
	}

	// take the above representation into array for convenience
	indexArr := make([]int, transitionKeys)
	for i, _ := range indexStr {
		if indexStr[i] >= '0' && indexStr[i] <= '9' {
			indexArr[i] = int(indexStr[i]) - 48
		} else if indexStr[i] >= 'a' && indexStr[i] <= 'z' {
			indexArr[i] = int(indexStr[i]) - 96 + 9
		} else {
			panic("non valid parsing value for indexStr")
		}
	}

	// the shiny new transition table
	tm.transitionTable = make(map[TTKey]TTValue, transitionKeys)

	for i := 0; i < transitionKeys; i++ {
		ttKey := TTKey{
			state:  STATEPICK[i/2],
			symbol: SYMBOLPICK[i%2],
		}

		// we have `4n + 2` = `2(2n + 1)` available transitions
		// let's start with the new symbol
		newSymbol := SYMBOLPICK[indexArr[i]/(availableTransitions/2)]

		// clean up this variable, i.e. consider only state and direction (`2m + 1`)
		indexArr[i] = indexArr[i] % (availableTransitions / 2)

		var newState, newDirection string
		// and let's find out the new state
		if indexArr[i] == (availableTransitions/2)-1 {
			// the last one (of ``2n + 1`) should be the halting state
			newState = "H"
			// since the machine halted, there is no new direction
			// let's be explicit about that
			newDirection = ""
		} else {
			// this is a non-halting state, and they go two-by-two
			// (i.e. AL, AR, BL, BR, etc)
			newState = STATEPICK[indexArr[i]/2]
			// with the new direction, a case of odd/even
			newDirection = []string{"L", "R"}[indexArr[i]%2]
		}

		// we finally update the value on the transition table
		tm.transitionTable[ttKey] = TTValue{
			symbol:    newSymbol,
			state:     newState,
			direction: newDirection,
		}
	}
}

// getCurrentSymbol is a convenience method
func (tm *TuringMachine) getCurrentSymbol() string {
	return tm.tape.values[tm.tape.headPosition]
}

// setNewSymbol is a convenience method
func (tm *TuringMachine) setNewSymbol(input string) {
	tm.tape.values[tm.tape.headPosition] = input
}

// resetTape is a convenience method
func (tm *TuringMachine) resetTape() {
	tm.tape.values = []string{""}
}

// run performs the execution of the Turing Machine without input
// on an empty tape. If it succesfully halts, its output can be retrieved
// with the function getResultingString()
func (tm *TuringMachine) run() {
	for i := 0; true; i++ {
		if i == MAX_NUMBER_OF_STEPS {
			// we assume the machine won't halt
			// so, we delete the tape and return
			tm.resetTape()
			return
		}

		// reads current symbol on tape andmap the transition
		// (state, symbol) -> (symbol, state, direction)
		currentStateSymbol := TTKey{
			symbol: tm.getCurrentSymbol(),
			state:  tm.currentState,
		}
		transition := tm.transitionTable[currentStateSymbol]

		// writes the new symbol
		tm.setNewSymbol(transition.symbol)

		// updates the state
		tm.currentState = transition.state

		// if the new state is `Halt`, we return
		if tm.currentState == "H" {
			return
		}

		// otherwise, move the head across the tape
		tm.moveHead(transition.direction)
	}
}

// moveHead moves the head of the turing machine in the
// specified direction, creating new cells in the tape
// when needed, filling them with a _zero_
func (tm *TuringMachine) moveHead(direction string) {
	if direction == "L" {
		if tm.tape.headPosition == 0 {
			// we are at the leftmost space
			// add a zero cell at the left, keep the index at zero
			tm.tape.values = append([]string{"0"}, tm.tape.values...)
		} else {
			tm.tape.headPosition--
		}
	} else {
		if tm.tape.headPosition == len(tm.tape.values)-1 {
			// we are at the rightmost space
			// add a zero cell to the right
			tm.tape.values = append(tm.tape.values, "0")
		}

		// always move the head position
		tm.tape.headPosition++
	}
}

// getResultingString is a convenience method
func (tm *TuringMachine) getResultingString() string {
	output := ""
	for _, cell := range tm.tape.values {
		output += cell
	}

	return output
}

// printResultMap outputs the key in a sorted fashion
func printResultMap(m map[string]int) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Printf("Explored space of (%d, 2) TMs:\n\n", STATES)
	for _, k := range keys {
		fmt.Printf("\t%s\t%d\n", k, m[k])
	}
}

// entry point, will run the space of (n,2) based on certain logic
func main() {
	// used for the iteration
	index := new(big.Int).SetInt64(0)
	one := big.NewInt(1)

	// contains our obtained results
	resultMap := make(map[string]int)

	for {
		if index.Cmp(MAX_MACHINES) >= 0 {
			break
		}

		// create the turing machine
		tm := newTuringMachine(STATES, index)

		// run turing machine
		tm.run()

		// verify output
		result := tm.getResultingString()
		if result == "" {
			// if empty, it means the machine didn't halted at MAX_NUMBER_OF_STEPS
			// we increment our HALTED key anyways
			result = "HALTED"
		}
		if i, ok := resultMap[result]; !ok {
			resultMap[result] = 1
		} else {
			resultMap[result] = i + 1
		}

		// increment the counter
		index.Add(index, one)
	}

	// print the result map
	printResultMap(resultMap)
}
