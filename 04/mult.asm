// Multiply two numbers
// RAM[2] = RAM[0] * RAM[1]

    @R2
    M=0

    @R0
    D=M

    // if R0 > R1 goto SWAP
    @R1
    D=D-M
    @SWAP
    D;JGT

(SWAP)
    @R0
    D=M
    @temp
    M=D

    @R1
    D=M
    @R0
    M=D
    
    @temp
    D=M
    @R1
    M=D

(LOOP)

    @R0
    D=M
    @END
    D;JEQ
    
    @R1
    D=M
    @R2
    M=D+M


    @R0
    M=M-1
    @LOOP
    0;JMP

(END)
    @END
    0;JMP

