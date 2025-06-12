package vm

import (
	"fmt"
	"monkey/compiler/code"
	"monkey/compiler/compiler"
	"monkey/compiler/object"
)

const StackSize = 2048

type VM struct {
	// Which parts of VM do we need?
	// * instructions
	// * stack
	// * constants
	constants    []object.Object
	instructions code.Instructions
	stack        []object.Object
	// Always points to the next value, so top of stack is stack[sp-1]
	sp int
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		constants:    bytecode.Constants,
		instructions: bytecode.Instructions,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			// fetch from 'fetch-and-execute' pattern
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2 // Increment by 'number of bytest we read to decode the operands'
			// execute from 'fetch-and-execute' pattern
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}
