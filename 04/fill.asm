// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed the program blackens the screen,
// When no key is pressed, the screen should be cleared.


(LOOP)
    @KBD
    D=M
    @BLACK
    D;JGT
    @WHITE
    0;JMP


    @LOOP
    0;JMP

(BLACK)
    @c
    M=-1
    @COLOR
    0;JMP


(WHITE)
    @c
    M=0
    @COLOR
    0;JMP


(COLOR)
    @KBD
    D=A
    @i
    M=D-1

(LOOP2)
    // if screen colored goto LOOP
    @i
    D=M
    @SCREEN
    D=D-A
    @LOOP
    D;JEQ

    @c
    D=M
    @i
    A=M
    M=D
    @i
    M=M-1
    @LOOP2
    0;JMP


