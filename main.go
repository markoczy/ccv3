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

const (
	initStateId          = 0
	parseNumericStateId  = 1
	parseOperatorStateId = 2
	endStateId           = 99
)

var parserFsm = fsm.NewParserFsm(
	map[int]fsm.ParserState{
		// Init
		initStateId: *fsm.NewParserState(false, fsm.NoOp, map[parser.TokenType]int{
			parser.NumericToken: parseNumericStateId,
		}),
		// Parse Numeric
		parseNumericStateId: *fsm.NewParserState(false,
			func(s *fsm.StateParams) error {
				token := s.Tokens.Dequeue()
				f, err := strconv.ParseFloat(token.Value, 64)
				if err != nil {
					return err
				}
				n := common.ValueNode(f)
				s.Cur.AddChild(&n)
				return nil
			},
			map[parser.TokenType]int{
				parser.EndToken:      endStateId,
				parser.OperatorToken: parseOperatorStateId,
			}),
		// Parse Operator
		parseOperatorStateId: *fsm.NewParserState(false,
			func(s *fsm.StateParams) error {
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
			map[parser.TokenType]int{
				parser.NumericToken: parseNumericStateId,
			}),
		// End
		endStateId: fsm.EndState,
	},
	initStateId,
)

func testParserFsm() {
	eq, err := parserFsm.Parse("1+2.5-5*3")
	if err != nil {
		panic(err)
	}

	fmt.Println("*** Equation:", eq.String())
}
