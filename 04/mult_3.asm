// Solution: repetitive shifted addition with swap (simulate by-hand multiplication)


// At each step n, multiply R1 (bigger number) with n-th LSB of R0, shift left by n, add to R2

// R2(n) = [(R1 * R0-LSB-n) << n ] + R2(n-1)

// For example R1=1010, R0=110
// 
// 0. (1010 * 0) << 0 + 0     == 0000
// 1. (1010 * 1) << 1 + 0000  == 10100
// 2. (1010 * 1) << 2 + 10100 == 111100



