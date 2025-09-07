// A filled rectangle at the screen's top left corner
// width=16 pixels height=RAM[0] pixels
//
// for i=0; i < n; i++ {
//     draw 16 black pixels at the beginning of row i 
// }

    @R0
    D=M
    @n
    M=D

    @SCREEN
    D=A
    @addr
    M=D

    @i
    M=0

(LOOP)
    // if i > n goto END
    @i
    D=M
    @n
    D=D-M
    @END
    D;JGT

    @addr
    A=M
    M=-1 // RAM[addr]=1111111111111111 

    @32
    D=A
    @addr
    M=D+M
    @i
    M=M+1
    @LOOP
    0;JMP


(END)
    @END
    0;JMP
