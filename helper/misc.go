package helper

import (
	"bufio"
	"io/ioutil"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func StartReading(fileName string, channel chan string) {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer file.Close()
	defer close(channel)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		channel <- scanner.Text()
	}

	err = scanner.Err()
	CheckErr(err)
}

func ReadLines(fileName string) (lines []string) {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	CheckErr(err)

	return
}

func ReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	return string(bytes)
}
