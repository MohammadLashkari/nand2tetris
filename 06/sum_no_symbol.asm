// Sum 1 to N
// RAM[1] = 1+2+...+RAM[0]

    @16 // i = 1
    M=1

    @17 // sum = 0
    M=0

    @16 // if i›RAM[0] goto STOP
    D=M
    @0
    D=D-M
    @18
    D;JGT

    @16 // sum += i
    D=M
    @17
    M=D+M

    @16 // i++
    M=M+1
    @4  // goto LOOP
    0;JMP

    @17
    D=M
    @1
    M=D // RAM[1] = sum

    @22
    0;JMP
