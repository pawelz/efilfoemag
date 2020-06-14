package state

type State bool

const (
	Alive State = true
	Dead State = false
)

// Of converts a boolean to the corresponding State.
func Of(isAlive bool) State {
	if isAlive {
		return Alive
	}
	return Dead
}

// IsAlive converts a State to the corresponding boolean.
func (s State) IsAlive() bool {
	return s == Alive
}
