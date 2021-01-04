package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const passportDataFile = "04.data.txt"
const passportSampleDataFile = "04.data.sample.txt"

const allFieldsPassportsMsg = "Number of passports with all required fields : %d\n"
const allValidPassportsMsg = "Number of passports with correct values for required fields : %d\n"

func main() {
	passports := readNewPassports(passportDataFile)
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var allFieldsCount, validPassportCount int
	for _, passport := range passports {
		if hasAllFields(passport, validFields) {
			allFieldsCount++
			if checkPassport(passport) {
				validPassportCount++
			}
		}
	}
	fmt.Printf(allFieldsPassportsMsg, allFieldsCount)
	fmt.Printf(allValidPassportsMsg, validPassportCount)
}

func hasAllFields(obj map[string]string, fields []string) bool {
	for _, field := range fields {
		if _, present := obj[field]; !present {
			return false
		}
	}
	return true
}

func checkPassport(passport map[string]string) bool {
	// byr [1920-2002]
	byr, err := strconv.Atoi(passport["byr"])
	helper.CheckErr(err)
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr [2010-2020]
	iyr, err := strconv.Atoi(passport["iyr"])
	helper.CheckErr(err)
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr [2020-2030]
	eyr, err := strconv.Atoi(passport["eyr"])
	helper.CheckErr(err)
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt [150cm-193cm] or [59-76in]
	reHgt := regexp.MustCompile(`(\d+)(cm|in)`)
	matches := reHgt.FindStringSubmatch(passport["hgt"])
	if len(matches) != 3 {
		return false
	}
	hgt, err := strconv.Atoi(matches[1])
	if matches[2] == "cm" {
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else if matches[2] == "in" {
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		helper.CheckErr(errors.New("Invalid value for hgt"))
	}

	// hcl #xxxxxx ; where x=>[0-9a-z]
	reHcl := regexp.MustCompile(`^#[a-z0-9]{6}$`)
	matches = reHcl.FindStringSubmatch(passport["hcl"])
	if len(matches) != 1 {
		return false
	}

	// ecl values in {"amb","blu","brn","gry","grn","hzl","oth"}
	reEcl := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	matches = reEcl.FindStringSubmatch(passport["ecl"])
	if len(matches) != 1 {
		return false
	}

	// pid a 9 digit number including leading zeros
	rePid := regexp.MustCompile(`^\d{9}$`)
	matches = rePid.FindStringSubmatch(passport["pid"])
	if len(matches) != 1 {
		return false
	}

	return true
}

func readNewPassports(fileName string) (passports []map[string]string) {
	content := helper.ReadFile(fileName)
	// Split passport entries in file
	entries := strings.Split(content, "\n\n")
	for _, entry := range entries {
		passport := make(map[string]string)
		// Split by whitespace characters to get list of field:value
		for _, fields := range strings.Fields(entry) {
			keyval := strings.Split(fields, ":")
			passport[keyval[0]] = keyval[1]
		}
		// Add passport
		passports = append(passports, passport)
	}
	return passports
}
