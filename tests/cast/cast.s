// Code generated by command: go run asm.go -out cast.s -stubs stub.go. DO NOT EDIT.

// func Split(x uint64) (q uint64, l uint32, w uint16, b uint8)
TEXT ·Split(SB), $0-23
	MOVQ	x(FP), AX
	MOVQ	AX, q+8(FP)
	MOVL	AX, l+16(FP)
	MOVW	AX, w+20(FP)
	MOVB	AL, b+22(FP)
	RET
