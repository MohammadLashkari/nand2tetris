
// Solution: repetitively add R1 to R2, R0 amount of times
//           swap R1, R0 if R1<R1

// Worst case: R0,R1 = 181

    @R2
    M=0

    @R1
    D=M
    @R0
    D=D-M
    @SWAP
    D;JLT
    @LOOP
    0;JMP


(SWAP)
    @R1
    D=M
    @R0
    D=D+M
    M=D-M

    @R1
    M=D-M

(LOOP)
    // if R0 == 0 goto END
    @R0
    D=M
    @END
    D;JEQ

    // R2=R2+R1
    @R1
    D=M
    @R2
    M=D+M

    // R0--
    @R0
    M=M-1
    @LOOP
    0;JMP

(END)
    @END
    0;JMP

