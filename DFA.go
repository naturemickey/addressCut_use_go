package main

type DFA struct {
	startState *DFAState
}

func newDFA(names []string) *DFA {
	dfa := &DFA{startState: createDFAState()}

	for _, name := range names {
		var currentState = dfa.startState
		for _, c := range name {
			currentState = currentState.addPath(c)
		}
		currentState.name = name
	}

	return dfa
}
