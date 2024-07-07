package parser

import "strings"

const (
	eventLetter = iota
	eventSpace

	eventsCount
)

const (
	stateInitial = iota
	stateLetter
	stateSpace
	stateInvalid

	statesCount
)

type transition struct {
	perform func(byte)
	action  func()
}

type stateMachine struct {
	transitions [statesCount][eventsCount]transition
	curState    int

	wordBuilder strings.Builder
	tokens      []string
}

func newStateMachine() *stateMachine {
	sm := &stateMachine{
		curState: stateInitial,
	}

	sm.transitions = [statesCount][eventsCount]transition{
		stateInitial: {
			eventSpace:  {perform: sm.skipSpaceJump},
			eventLetter: {perform: sm.addLetterJump},
		},

		stateLetter: {
			eventSpace:  {perform: sm.skipSpaceJump, action: sm.addTokenAction},
			eventLetter: {perform: sm.addLetterJump},
		},
		stateSpace: {
			eventSpace:  {perform: sm.skipSpaceJump},
			eventLetter: {perform: sm.addLetterJump},
		},
		stateInvalid: {},
	}

	return sm
}

func (sm *stateMachine) parse(input string) ([]string, error) {
	for i := 0; i < len(input); i++ {
		character := input[i]
		event, err := sm.getEvent(character)
		if err != nil {
			return nil, err
		}

		sm.triggerEvent(event, character)
	}

	// calling to collect the last word, if any
	sm.triggerEvent(eventSpace, '\n')
	return sm.tokens, nil
}

func (sm *stateMachine) getEvent(character byte) (int, error) {
	if isLetter(character) {
		sm.triggerEvent(eventLetter, character)
	} else if isSpace(character) {
		sm.triggerEvent(eventSpace, character)
	}

	return 0, errParserInvalidCharacter
}

func (sm *stateMachine) triggerEvent(event int, character byte) {
	transition := sm.transitions[sm.curState][event]
	transition.perform(character)

	if transition.action != nil {
		transition.action()
	}
}

func (sm *stateMachine) skipSpaceJump(_ byte) {
	sm.curState = stateSpace
}

func (sm *stateMachine) addLetterJump(letter byte) {
	sm.wordBuilder.WriteByte(letter)
	sm.curState = stateLetter
}

func (sm *stateMachine) addTokenAction() {
	sm.tokens = append(sm.tokens, sm.wordBuilder.String())
	sm.wordBuilder.Reset()
}
