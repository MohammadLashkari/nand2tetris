package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

const (
	maxRAM        = 24_576
	variableIndex = 16
)

func main() {

	var output string
	flag.StringVar(&output, "o", "stdout", "output file name")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] <assembly-file-path>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if output == "" {
		fmt.Fprintf(os.Stderr, "output filename cannot be empty\n")
		os.Exit(1)
	}

	filename := flag.Args()[0]
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "file does not exist %q\n", filename)
		os.Exit(1)

	}

	ext := filepath.Ext(filename)
	if ext != ".asm" {
		fmt.Fprintf(os.Stderr, "invalid file extension %q\n", ext)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open the file %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var writer io.Writer = os.Stdout
	if output != "stdout" {
		writer, err = os.Create(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create output file: %v\n", err)
			os.Exit(1)
		}
	}

	// First pass: extract label symbols
	scanner := bufio.NewScanner(file)
	for instrN := 0; scanner.Scan(); {

		line := strings.TrimSpace(strings.Split(scanner.Text(), "//")[0])
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			symbol := strings.Trim(line, "()")
			labelSymbols[symbol] = instrN
			continue
		}
		instrN++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read the file %v\n", err)
		os.Exit(1)
	}

	// rewind file
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		fmt.Fprintf(os.Stderr, "failed to rewind file: %v\n", err)
		os.Exit(1)
	}

	// Second Pass: parse instructions
	scanner = bufio.NewScanner(file)
	for lineN := 1; scanner.Scan(); lineN++ {

		line := strings.TrimSpace(strings.Split(scanner.Text(), "//")[0])
		if line == "" || strings.HasPrefix(line, "(") {
			continue
		}

		if line[0] == '@' {
			parseAInstruction(lineN, line, writer)
		} else {
			parseCInstruction(lineN, line, writer)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read the file %v\n", err)
		os.Exit(1)
	}
}

// A-instruction
// @value: A non-negative decimal constant OR A symbol referring to a constant
func parseAInstruction(lineN int, instr string, w io.Writer) {
	instr = instr[1:]

	symbol, ok := preDefinedSymbols[instr]
	if ok {
		fmt.Fprintf(w, "%016b\n", symbol)
		return
	}

	symbol, ok = labelSymbols[instr]
	if ok {
		fmt.Fprintf(w, "%016b\n", symbol)
		return
	}

	symbol, ok = variableSymbols[instr]
	if ok {
		fmt.Fprintf(w, "%016b\n", symbol)
		return
	}

	if !unicode.IsDigit(rune(instr[0])) {
		variableAddr := len(variableSymbols) + variableIndex
		variableSymbols[instr] = variableAddr
		fmt.Fprintf(w, "%016b\n", variableAddr)
		return
	}

	value, err := strconv.ParseInt(instr, 10, 16)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[line %d] failed to parse A-instruction %q: %v\n", lineN, instr, err)
		os.Exit(1)
	}

	if value < 0 || value > maxRAM {
		fmt.Fprintf(os.Stderr, "[line %d] invalid A-instruction %q (must be 0–32767)\n", lineN, instr)
		os.Exit(1)
	}

	fmt.Fprintf(w, "%016b\n", value)
}

// C-instruction
// dest = comp; jump -> 111-A-C1C2C3C4C5C6-D1D2D3-J1J2J3
func parseCInstruction(lineN int, instr string, w io.Writer) {
	const (
		aIdx = 3
		dIdx = 10
		jIdx = 13
	)

	parsed := bytes.Repeat([]byte{'0'}, 16)
	parsed[0], parsed[1], parsed[2] = '1', '1', '1'

	dest, compAndJump, hasDest := strings.Cut(instr, "=")
	if hasDest {
		dBits, ok := destMap[dest]
		if !ok {
			fmt.Fprintf(os.Stderr, "[line %d] invalid dest %q in C-instruction %q\n", lineN, dest, instr)
			os.Exit(1)
		}
		for i := range len(dBits) {
			parsed[i+dIdx] = dBits[i]
		}

	} else {
		compAndJump = dest
	}

	comp, jump, hasJump := strings.Cut(compAndJump, ";")
	if hasJump {
		jBits, ok := jumpMap[jump]
		if !ok {
			fmt.Fprintf(os.Stderr, "[line %d] invalid jump %q in C-instruction %q\n", lineN, jump, instr)
			os.Exit(1)
		}
		for i := range len(jBits) {
			parsed[i+jIdx] = jBits[i]
		}
	}

	acBits, ok := compMap[comp]
	if !ok {
		fmt.Fprintf(os.Stderr, "[line %d] invalid comp %q in C-instruction %q\n", lineN, comp, instr)
		os.Exit(1)
	}
	for i := range len(acBits) {
		parsed[i+aIdx] = acBits[i]
	}

	fmt.Fprintln(w, string(parsed))
}

var (
	labelSymbols    = map[string]int{}
	variableSymbols = map[string]int{}
	// Pre-defined symbols are use only in A-instructions @pre-defined
	preDefinedSymbols = map[string]int{
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
	}

	destMap = map[string]string{
		"M":   "001",
		"D":   "010",
		"MD":  "011",
		"A":   "100",
		"AM":  "101",
		"AD":  "110",
		"AMD": "111",
	}

	jumpMap = map[string]string{
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}

	compMap = map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"M":   "1110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"!M":  "1110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"-M":  "1110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"M+1": "1110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"M-1": "1110010",
		"D+A": "0000010",
		"D+M": "1000010",
		"D-A": "0010011",
		"D-M": "1010011",
		"A-D": "0000111",
		"M-D": "1000111",
		"D&A": "0000000",
		"D&M": "1000000",
		"D|A": "0010101",
		"D|M": "1010101",
	}
)
