package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func showBanner() {
	fmt.Printf(`
88888888888 .d8888b.  888    d8P  888b    888                            
    888    d88P  Y88b 888   d8P   8888b   888                            
    888    888    888 888  d8P    88888b  888                            
    888    888        888d88K     888Y88b 888  .d88b.   .d88b.  88888b.  
    888    888        8888888b    888 Y88b888 d88P"88b d8P  Y8b 888 "88b 
    888    888    888 888  Y88b   888  Y88888 888  888 88888888 888  888 
    888    Y88b  d88P 888   Y88b  888   Y8888 Y88b 888 Y8b.     888  888 
    888     "Y8888P"  888    Y88b 888    Y888  "Y88888  "Y8888  888  888 
                                                   888                   
                                              Y8b d88P			     
                                               "Y88P"			
									Built by "Catakan" with hate.
									https://github.com/Armageddon0x00
									https://twitter.com/0x00Armageddon
`)
}

func showHelp() {
	fmt.Printf(`
Generate:
tckn generate 5
tckn generate 5 nobanner
tckn generate endless nobanner

Validate:
tckn validate 12345678901
tckn validate exampleNumbers.txt nobanner
tckn validate exampleNumbers.txt valid nobanner
tckn validate exampleNumbers.txt invalid nobanner

Write:
tckn generate 100 nobanner > example100.txt
tckn generate endless nobanner > endlessTCKN.txt

Misc:
time (tckn generate 50000 nobanner > qq && tckn validate qq nobanner && rm qq)
`)
}

func validateTckn(TCKN string) bool {
	patternTckn := regexp.MustCompile(`^[0-9]{11}$`)

	if patternTckn.MatchString(TCKN) {
		if TCKN[0] == '0' {
			return false
		}

		oddNumbers := int(TCKN[0]-'0') + int(TCKN[2]-'0') + int(TCKN[4]-'0') + int(TCKN[6]-'0') + int(TCKN[8]-'0')
		evenNumbers := int(TCKN[1]-'0') + int(TCKN[3]-'0') + int(TCKN[5]-'0') + int(TCKN[7]-'0')
		tenthNumber := (oddNumbers*7 - evenNumbers) % 10
		generalSum := (oddNumbers + evenNumbers + int(TCKN[9]-'0')) % 10

		if tenthNumber != int(TCKN[9]-'0') {
			return false
		}
		if generalSum != int(TCKN[10]-'0') {
			return false
		}
		return true
	}
	return false
}

func generateTCKN() string {

	base := rand.Intn(899999999) + 100000000
	baseStr := fmt.Sprintf("%d", base)

	var oddSum, evenSum int
	for i := 0; i < 9; i++ {
		digit := int(baseStr[i] - '0')
		if i%2 == 0 {
			oddSum += digit
		} else {
			evenSum += digit
		}
	}

	tenthNumber := (oddSum*7 - evenSum) % 10
	generalSum := (oddSum + evenSum + tenthNumber) % 10

	tckn := baseStr + fmt.Sprintf("%d%d", tenthNumber, generalSum)

	return tckn
}

func readTCKN(fileName string) ([]string, []string) {
	var fileLines []string
	var validTCKN []string
	var invalidTCKN []string

	lineTCKN, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(lineTCKN)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	lineTCKN.Close()

	for _, tcnumber := range fileLines {
		if validateTckn(strings.TrimSpace(tcnumber)) {
			validTCKN = append(validTCKN, tcnumber)
		} else {
			invalidTCKN = append(invalidTCKN, tcnumber)
		}
	}

	return validTCKN, invalidTCKN
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// This is bad LMAO
// RTFM for flags was a better option i guess.
func main() {

	if os.Args[len(os.Args)-1] != "nobanner" {
		showBanner()
	}

	if len(os.Args) < 2 {
		fmt.Println("Please enter an option.")
		showHelp()
	} else if os.Args[1] == "validate" {
		validateFile := os.Args[2]

		if fileExists(validateFile) {
			validArr, invalidArr := readTCKN(validateFile)
			if os.Args[3] == "valid" {
				for _, validTC := range validArr {
					fmt.Println(validTC)
				}
			} else if os.Args[3] == "invalid" {
				for _, invalidTC := range invalidArr {
					fmt.Println(invalidTC)
				}
			} else {
				for _, validTC := range validArr {
					fmt.Println("[+] TCKN valid.", validTC)
				}
				for _, invalidTC := range invalidArr {
					fmt.Println("[-] TCKN not valid.", invalidTC)
				}
			}
		} else {
			if validateTckn(os.Args[2]) {
				fmt.Println("[+] TCKN valid.", os.Args[2])
			} else {
				fmt.Println("[-] TCKN not valid.", os.Args[2])
			}
		}
	} else if os.Args[1] == "generate" {
		if os.Args[2] != "endless" {
			generateNumber, err := strconv.Atoi(os.Args[2])

			if err != nil {
				fmt.Println(err)
			}

			for i := 0; i < generateNumber; i++ {
				fmt.Println(generateTCKN())
			}
		} else {
			for {
				// This is broken. Endless wacks shit.
				fmt.Println(generateTCKN())
			}
		}
	} else {
		fmt.Println("Invalid options.")
		showHelp()
	}
}
