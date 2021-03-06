package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
	"github.com/markoczy/ccv3/parser/fsm"
	"github.com/markoczy/ccv3/parser/tokenizer"
)

func main() {
	// Disable log
	//log.SetOutput(ioutil.Discard)
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("[DEBUG] ")

	fmt.Println("Test", '0', '9')
	// testRuneQueue()
	// testCreateToken()
	// testCreateTokens()
	// testParserFsm()
	testFunction()
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
		fmt.Printf("Token: begin=%d, end=%d, type=%s, val=%s\n", v.Begin, v.End, v.Type, v.Value)
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

var parserFsm = fsm.NewParserFsm(
	map[int]fsm.ParserState{
		// Init
		initStateId: fsm.NewParserStateWithExecAndTransitionMap(
			fsm.NoOp,
			map[parser.TokenType]int{
				parser.NumericToken: parseNumericStateId,
			}),
		// Parse Numeric
		parseNumericStateId: fsm.NewParserStateWithExecAndTransitionMap(
			func(s *fsm.CallStack) error {
				state := s.Cur()
				token := s.Tokens.Dequeue()
				f, err := strconv.ParseFloat(token.Value, 64)
				if err != nil {
					return err
				}
				n := common.ValueNode(f)
				state.Equation.AddChild(&n)
				return nil
			},
			map[parser.TokenType]int{
				parser.EndToken:      endStateId,
				parser.OperatorToken: parseOperatorStateId,
			}),
		// Parse Operator
		parseOperatorStateId: fsm.NewParserStateWithExecAndTransitionMap(
			func(s *fsm.CallStack) error {
				state := s.Cur()
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
				case "^":
					op = common.Exponentiation
				default:
					return fmt.Errorf("Unhandled Operation")
					// fmt.Errorf("Undefined operation: %s", token.Value)
				}
				state.Equation.AddOperation(op)
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

const (
	// basic
	initStateId          = 0
	parseNumericStateId  = 1
	parseOperatorStateId = 2
	endStateId           = 99
	// advanced
	parseControlStateId       = 10
	parseUnaryOperatorStateId = 11
	// da reeel stuff
	parseIdentifierStateId = 20
	parseParamStateId      = 21
	postParseParamStateId  = 22
)

var parserFsm2 = fsm.NewParserFsm(
	map[int]fsm.ParserState{
		// Init
		initStateId: fsm.NewParserStateWithExecAndTransitionMap(
			fsm.NoOp,
			map[parser.TokenType]int{
				parser.NumericToken:  parseNumericStateId,
				parser.ControlToken:  parseControlStateId,
				parser.OperatorToken: parseUnaryOperatorStateId,
			}),
		// Parse Numeric
		parseNumericStateId: fsm.NewParserStateWithExecAndTransitionMap(
			func(s *fsm.CallStack) error {
				state := s.Cur()
				token := s.Tokens.Dequeue()
				f, err := strconv.ParseFloat(token.Value, 64)
				if err != nil {
					return err
				}
				if state.Negate {
					state.Negate = false
					f = -f
				}
				n := common.ValueNode(f)
				state.Equation.AddChild(&n)
				return nil
			},
			map[parser.TokenType]int{
				parser.EndToken:      endStateId,
				parser.OperatorToken: parseOperatorStateId,
				parser.ControlToken:  parseControlStateId,
			}),
		// Parse Operator
		parseOperatorStateId: fsm.NewParserStateWithExecAndTransitionMap(
			func(s *fsm.CallStack) error {
				state := s.Cur()
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
				case "^":
					op = common.Exponentiation
				default:
					return fmt.Errorf("Unhandled operation")
				}
				state.Equation.AddOperation(op)
				return nil
			},
			map[parser.TokenType]int{
				parser.NumericToken:  parseNumericStateId,
				parser.ControlToken:  parseControlStateId,
				parser.OperatorToken: parseUnaryOperatorStateId,
			}),
		// Parse Control
		parseControlStateId: fsm.NewParserStateWithExecAndTransitionMap(
			func(s *fsm.CallStack) error {
				token := s.Tokens.Dequeue()

				switch token.Value {
				case "(":
					if err := s.Enter(); err != nil {
						return err
					}
				case ")":
					if err := s.Exit(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("Unhandled control")
				}
				return nil
			},
			map[parser.TokenType]int{
				parser.NumericToken:  parseNumericStateId,
				parser.OperatorToken: parseOperatorStateId,
				parser.EndToken:      endStateId,
				parser.ControlToken:  parseControlStateId,
			}),
		parseUnaryOperatorStateId: fsm.NewParserStateWithExecAndTransitionMap(
			func(s *fsm.CallStack) error {
				state := s.Cur()
				token := s.Tokens.Dequeue()

				switch token.Value {
				case "-":
					state.Negate = !state.Negate
				default:
					return fmt.Errorf("Unhandled unary operator")
				}
				return nil
			},
			map[parser.TokenType]int{
				parser.NumericToken: parseNumericStateId,
				parser.ControlToken: parseControlStateId,
			}),
		parseIdentifierStateId: fsm.NewParserState(func(stack *fsm.CallStack) (next int, err error) {
			token := stack.Tokens.Dequeue()
			fn, err := stack.Functions.Function(token.Value)
			if err != nil {
				return 0, err
			}
			if !fn.Parametrized() {

			}
			// id := stack.Functions.
			return
		}),
		// End
		endStateId: fsm.EndState,
	},
	initStateId,
)

func testParserFsm() {
	// eq, err := parserFsm2.Parse("1+-2.5^(5-2)*3")
	eq, err := parserFsm2.Parse("1+-(2.5^(5-2))*3", common.NewFunctionMap(nil))
	// eq, err := parserFsm2.Parse("1+1")
	if err != nil {
		panic(err)
	}

	fmt.Println("*** Equation:", eq.String())
	res, err := eq.Value()
	if err != nil {
		panic(err)
	}
	fmt.Println("*** Value:", res)
}

func testFunction() {
	fn := common.NewDefaultFunction("sin", 1, func(x ...float64) float64 {
		return math.Sin(x[0])
	})
	x := common.ValueNode(1)
	fmt.Println(fn.AsNode(&x))
}
