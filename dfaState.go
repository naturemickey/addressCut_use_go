package main

type DFAState struct {
	name string
	path map[rune]*DFAState
}

func createDFAState() *DFAState {
	return &DFAState{name: "", path: make(map[rune]*DFAState)}
}

func (s *DFAState) isAccepted() bool {
	return s.name != ""
}

func (s *DFAState) addPath(c rune) *DFAState {
	if state, ok := s.path[c]; ok {
		return state
	} else {
		state = createDFAState()
		s.path[c] = state
		return state
	}
}
