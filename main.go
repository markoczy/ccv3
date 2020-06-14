package main

import (
	"fmt"
	"strconv"

	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
	"github.com/markoczy/ccv3/parser/fsm"
	"github.com/markoczy/ccv3/parser/tokenizer"
)

func main() {
	fmt.Println("Test", '0', '9')
	// testRuneQueue()
	// testCreateToken()
	// testCreateTokens()
	testParserFsm()
}

func testCreateToken() {
	stack := tokenizer.RuneQueue("2134")
	token := tokenizer.CreateToken(&stack, parser.NumericToken, 0, parser.IsNumeric)
	fmt.Println("token:", token.Value)
}

func testCreateTokens() {
	tokens, err := tokenizer.CreateTokens("(1+12.5)*3-22")
	if err != nil {
		panic(err)
	}

	for _, v := range tokens {
		fmt.Printf("Token: begin=%d, end=%d, type=%d, val=%s\n", v.Begin, v.End, v.Type, v.Value)
	}
}

func testRuneQueue() {
	queue := tokenizer.RuneQueue("abc")
	fmt.Println("queue:", queue)
	r := queue.Peek()
	fmt.Println("rune:", r)
	fmt.Println("queue:", queue)
	r = queue.Dequeue()
	fmt.Println("rune:", r)
	fmt.Println("queue:", queue)
	r = queue.Dequeue()
	fmt.Println("rune:", r)
	fmt.Println("queue:", queue)
	r = queue.Dequeue()
	fmt.Println("rune:", r)
	fmt.Println("queue:", queue)
}

func testParserFsm() {
	initStateId := 0
	parseNumericStateId := 1
	parseOperatorStateId := 2
	endStateId := 99

	states := map[int]fsm.ParserState{}

	stateInit := fsm.ParserState{
		End:         false,
		Func:        fsm.NoOp,
		Transitions: map[parser.TokenType]int{},
	}
	stateInit.Transitions[parser.NumericToken] = parseNumericStateId

	stateParseNum := fsm.ParserState{
		End: false,
		Func: func(s *fsm.StateParams) error {
			token := s.Tokens.Dequeue()
			f, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return err
			}
			n := common.ValueNode(f)
			s.Cur.AddChild(&n)
			return nil
		},
		Transitions: map[parser.TokenType]int{},
	}
	stateParseNum.Transitions[parser.EndToken] = endStateId
	stateParseNum.Transitions[parser.OperatorToken] = parseOperatorStateId

	stateParseOperator := fsm.ParserState{
		End: false,
		Func: func(s *fsm.StateParams) error {
			token := s.Tokens.Dequeue()
			var op common.Operation
			switch token.Value {
			case "+":
				op = common.Addition
			case "-":
				op = common.Subtraction
			case "/":
				op = common.Division
			case "*":
				op = common.Multiplication
			}
			s.Cur.AddOperation(op)
			return nil
		},
		Transitions: map[parser.TokenType]int{},
	}
	stateParseOperator.Transitions[parser.NumericToken] = parseNumericStateId

	states[initStateId] = stateInit
	states[parseNumericStateId] = stateParseNum
	states[parseOperatorStateId] = stateParseOperator
	states[endStateId] = fsm.EndState

	fsm := fsm.NewParserFsm(states, initStateId)
	eq, err := fsm.Parse("1+2.5-5*3")
	if err != nil {
		panic(err)
	}

	fmt.Println("*** Equation:", eq.String())
}
