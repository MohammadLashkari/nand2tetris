// Sum 1 to N
// RAM[1] = 1+2+...+RAM[0]

    @R0
    D=M
    @n
    M=D

    @i
    M=1

    @sum
    M=0

(LOOP)
    // if i > n goto STOP
    @i
    D=M
    @n
    D=D-M
    @STOP
    D;JGT

    // sum += i
    @sum
    D=M
    @i
    D=D+M
    @sum
    M=D
    // i++
    @i
    M=M+1

    @LOOP
    0;JMP


(STOP)
    @sum
    D=M
    @R1
    M=D

(END)
    @END
    0;JMP
