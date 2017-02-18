package main

var dfa *DFA

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

	bl := []rune(s)
	res := []string{}

	for length := len(bl); currentIdx < length; currentIdx++ {
		a := bl[currentIdx]
		currentState = currentState.tran(a)
		if currentState == nil || currentIdx+1 == length {
			if currentState != nil && currentIdx+1 == length && currentState.isAccepted() {
				if !contains(res, currentState.name) {
					res = append(res, currentState.name)
				}
			} else if currentAccepted != nil {
				if !contains(res, currentAccepted.name) {
					res = append(res, currentAccepted.name)
				}
				fromIdx = currentAcceptedIdx + 1
				currentAccepted = nil
				currentIdx = currentAcceptedIdx
			} else {
				currentIdx = fromIdx
				fromIdx = fromIdx + 1
			}
			currentState = this.startState
		} else if currentState.isAccepted() {
			currentAccepted = currentState
			currentAcceptedIdx = currentIdx
		}
	}

	return res
}

func contains(this []string, s string) bool {
	for _, str := range this {
		if str == s {
			return true
		}
	}
	return false
}
