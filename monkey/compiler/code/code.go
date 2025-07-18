package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
	OpPop
	// Monkey supports 8 infix operations,
	// with 4 of them being used for arithmetic: + , - , * , /
	OpAdd
	OpSub
	OpMul
	OpDiv
	// Boolean literals should cause VM to load boolean values on to the stack,
	// instead of evaluating to boolean values, as in 'evaluator'
	OpTrue
	OpFalse
	// Now, for the comparison operators, we only actually need 3 opcodes to support all of them
	OpEqual
	OpNotEqual
	OpGreaterThan
	// Prefix expressions are even simpler
	OpMinus
	OpBang
	// Jump opcodes are basically a way to control the flow of execution,
	// so they are in fact conditionals
	OpJumpNotTruthy
	OpJump
	OpNull
	// Binding-related opcodes
	OpGetGlobal
	OpSetGlobal
	// Composite types
	OpArray
	OpHash
	OpIndex
	// Functions
	OpCall
	OpReturnValue // explicit return
	OpReturn      // implicit return - tell VM to return 'nil' - design decision
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	// 'OpPop' needs to be used after every expression, as they don't store value, only return it.
	// That's the whole point of the 'expression vs statement' thing.
	OpPop: {"OpPop", []int{}},
	// Monkey supports 8 infix operations,
	// with 4 of them being used for arithmetic: + , - , * , /
	OpAdd:           {"OpAdd", []int{}},
	OpSub:           {"OpSub", []int{}},
	OpMul:           {"OpMul", []int{}},
	OpDiv:           {"OpDiv", []int{}},
	OpTrue:          {"OpTrue", []int{}},
	OpFalse:         {"OpFalse", []int{}},
	OpEqual:         {"OpEqual", []int{}},
	OpNotEqual:      {"OpNotEqual", []int{}},
	OpGreaterThan:   {"OpGreaterThan", []int{}},
	OpMinus:         {"OpMinus", []int{}},
	OpBang:          {"OpBang", []int{}},
	OpJump:          {"OpJump", []int{2}},
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpNull:          {"OpNull", []int{}},
	// Both have single 2-byte operand to hold the unique number of global binding
	OpGetGlobal:   {"OpGetGlobal", []int{2}},
	OpSetGlobal:   {"OpSetGlobal", []int{2}},
	OpArray:       {"OpArray", []int{2}},
	OpHash:        {"OpHash", []int{2}},
	OpIndex:       {"OpIndex", []int{}},
	OpCall:        {"OpCall", []int{}},
	OpReturnValue: {"OpReturnValue", []int{}},
	OpReturn:      {"OpReturn", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1

	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + read
	}

	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}
	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}
