// Swap RAM[0] and RAM[1]

    // temp = R0
    @R0
    D=M
    @temp
    M=D

    // R0 = R1
    @R1
    D=M
    @R0
    M=D

    // R1 = temp
    @temp
    D=M
    @R1
    M=D


(END)
    @END
    0;JMP


