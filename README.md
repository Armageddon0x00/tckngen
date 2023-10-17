# TCKNgen
A simple utility tool for generating and validating Turkish National Identity Numbers (TCKN).

- [TCKNgen](#tckngen)
- [Installation](#installation)
  - [Using `go install`](#using-go-install)
  - [From Source](#from-source)
- [Usage](#usage)
  - [Generation](#generation)
    - [Generate 5 TCKN](#generate-5-tckn)
    - [Generate 5 TCKN Without Banner](#generate-5-tckn-without-banner)
    - [Generate Endless](#generate-endless)
  - [Validation](#validation)
    - [Validate Single TCKN](#validate-single-tckn)
    - [Validate All TCKN in File](#validate-all-tckn-in-file)
    - [Only Show Valid TCKN in a File](#only-show-valid-tckn-in-a-file)
    - [Only Show Invalid TCKN in a File](#only-show-invalid-tckn-in-a-file)
- [Same but in Bash](#same-but-in-bash)
  - [Validate TCKN Using Bash](#validate-tckn-using-bash)
  - [Generate TCKN Using Bash](#generate-tckn-using-bash)
- [TODO](#todo)

# Installation

## Using `go install`
```bash
go install github.com/Armageddon0x00/tckngen@latest
```

## From Source
```bash
git clone https://github.com/Armageddon0x00/tckngen
cd tckngen
go install .
```

# Usage

## Generation

### Generate 5 TCKN
```bash
tckngen generate 5
```

### Generate 5 TCKN Without Banner
```bash
tckngen generate 5 nobanner
```

### Generate Endless
You can pipe this to a file to generate example dataset.
```bash
tckngen generate endless nobanner
```

## Validation

### Validate Single TCKN
```bash
tckngen validate 12345678901
```

### Validate All TCKN in File
```bash
tckngen validate exampleNumbers.txt
```

### Only Show Valid TCKN in a File
```
tckngen validate exampleNumbers.txt valid nobanner
```

### Only Show Invalid TCKN in a File
```
tckngen validate exampleNumbers.txt invalid nobanner
```

# Same but in Bash
Validate and generate TCKN using Bash. Useful for offline and slow needs.

## Validate TCKN Using Bash
```bash
tckncheck() {
  local RED="\e[31m"
  local ENDCOLOR="\e[0m"
  local GREEN="\e[32m"

  local TCKN="$1"
  local Pattern="^[0-9]{11}$"

  if [[ $TCKN =~ $Pattern ]]; then
    if [ "${TCKN:0:1}" -eq 0 ]; then
      echo -e "${RED} [-] TCKN not valid. $TCKN ${ENDCOLOR}"
    fi

    local odd=$(( ${TCKN:0:1} + ${TCKN:2:1} + ${TCKN:4:1} + ${TCKN:6:1} + ${TCKN:8:1} ))
    local even=$(( ${TCKN:1:1} + ${TCKN:3:1} + ${TCKN:5:1} + ${TCKN:7:1} ))
    local 10th_validator=$(( (odd * 7 - even) % 10 ))
    local sum=$(( (odd + even + ${TCKN:9:1}) % 10 ))

    if [ $10th_validator -ne ${TCKN:9:1} ]; then
      echo -e "${RED} [-] TCKN not valid. $TCKN ${ENDCOLOR}"
    fi

    if [ $sum -ne ${TCKN:10:1} ]; then
      echo -e "${RED} [-] TCKN not valid. $TCKN ${ENDCOLOR}"
    else
      echo -e "${GREEN} [+] TCKN is valid. $TCKN ${ENDCOLOR}"
    fi
  else
    echo -e "${RED} [-] TCKN not valid. $TCKN ${ENDCOLOR}"
  fi
}
```

## Generate TCKN Using Bash
TODO

# TODO
- Fix endless generation function (1/50000 error rate?)
- Add more validation options (only show valid - invalid)
- Improve generation algorithm (no duplicates in same generation, more randomness etc.)
- More parsing options
- Usable library functions ?
- KISS and Clean.
- Bash function for TCKN generation.

Sidenote: This was meant to be a simple hourly project. Now I'm invested.