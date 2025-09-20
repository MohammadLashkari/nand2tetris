// Solution: repetitive shifted addition with swap (simulate by-hand multiplication)


// At each step n, multiply R1 (bigger number) with n-th LSB of R0, shift left by n, add to R2
// R2(n) = [(R1 * R0-LSB-n) << n ] + R2(n-1)


// For example R1=1010, R0=110
// 
//    1010 
//    110
// ---------
//    0000
//   1010
//  1010
// ---------
//  111100
//
// 0. (1010 * 0) << 0 + 0     == 0000
// 1. (1010 * 1) << 1 + 0000  == 10100
// 2. (1010 * 1) << 2 + 10100 == 111100


// Shift-left-> append zero to the right == multiply by 10 == multiply by 2 == x+x == x<<1


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


    // R0--
    @R0
    M=M-1
    @LOOP
    0;JMP

(END)
    @END
    0;JMP

