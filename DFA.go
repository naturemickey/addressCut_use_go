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

func (this *DFA) scan(s string) []string {
	currentState := this.startState
	currentIdx := 0

	var currentAccepted *DFAState = nil
	currentAcceptedIdx := 0

	fromIdx := 0

	return []string{}
}
