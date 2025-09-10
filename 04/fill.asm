// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed the program blackens the screen,
// When no key is pressed, the screen should be cleared.


(LOOP)
    // reset i
    @KBD
    D=A
    @i
    M=D-1

    @KBD
    D=M

    @BLACK
    D;JGT
    @WHITE
    0;JMP


    @LOOP
    0;JMP

(BLACK)
    @color
    M=-1
    @FILL
    0;JMP


(WHITE)
    @color
    M=0
    @FILL
    0;JMP


(FILL)
    // if i < SCREEN goto LOOP
    @i
    D=M
    @SCREEN
    D=D-A
    @LOOP
    D;JLT

    @color
    D=M
    @i
    A=M
    M=D

    @i
    M=M-1
    @FILL
    0;JMP

