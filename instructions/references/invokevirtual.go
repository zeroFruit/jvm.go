package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
	kMethodRef   *heap.ConstantMethodref
	argSlotCount uint
}

func (instr *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	if instr.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		instr.kMethodRef = cp.GetConstant(instr.Index).(*heap.ConstantMethodref)
		instr.argSlotCount = instr.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(instr.argSlotCount)
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	method := instr.kMethodRef.GetVirtualMethod(ref)
	frame.Thread().InvokeMethod(method)
}