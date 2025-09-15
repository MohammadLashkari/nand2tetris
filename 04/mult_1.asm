// Multiply two numbers
// RAM[2] = RAM[0] * RAM[1]  (Assumption: R0 >= 0, R1 >= 0, R2 < 32767)

// 16-bit number -> unsigned max = 65,535 (2^16-1)
//               -> signed max   = 32,767 (2^15-1)


// Solution: repetitively add R1 to R2, R0 amount of times
// if R1=0,1 R0=32767 -> program runs 32767 times


    @R2
    M=0

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

