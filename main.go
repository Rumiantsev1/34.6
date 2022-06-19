package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	//name input data file
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя файла с исходными данными: ")
	inpf, err := reader1.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inpf = strings.TrimSpace(inpf)

	//name result file
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя файла для записи результата: ")
	resf, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	resf = strings.TrimSpace(resf)

	w := ResultByteSlice(inpf)
	WriteResultInFile(w, resf)
}
func ResultByteSlice(s string) []byte {
	var byteSlice []byte

	filename := "./" + s + ".txt"

	f, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileReader := bufio.NewReader(f)

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		if MathsOrNoMaths(string(line)) {
			a, b, i := ArgsForMaths(string(line))
			r := Maths(a, b, i)
			rstr := fmt.Sprintf("%d", r)
			//добавляемая строка с результатом
			add := []byte(string(line) + rstr + "\n")
			byteSlice = append(byteSlice, add...)
		}
	}
	return byteSlice
}
func ArgsForMaths(s string) (arg1 int, arg2 int, znak string) {

	bs := []byte(s)

	for i := 0; i < len(bs); i++ {
		if bs[i] == 42 || bs[i] == 47 || bs[i] == 43 || bs[i] == 45 {
			for j := i; j < len(bs); j++ {
				if bs[j] == 61 {

					b := []byte(s[i+1 : j])
					a := []byte(s[:i])
					Znak := string(s[i])
					Inta, err := strconv.Atoi(string(a))
					if err != nil {
						log.Fatal(err)
					}
					Intb, err := strconv.Atoi(string(b))
					if err != nil {
						log.Fatal(err)
					}
					*&arg1 += Inta
					*&arg2 += Intb
					*&znak += Znak
					return
				}
			}
		}

	}

	return arg1, arg2, znak
}

//проверка строки на математическое выражение
func MathsOrNoMaths(s string) bool {
	s = strings.TrimSpace(s)
	mathsRegex := regexp.MustCompile(`^[0-9]+[\+\-\*\/][0-9]+[\=]$`)
	isMatch := mathsRegex.MatchString(s)

	return isMatch
}

//вычисление результата
func Maths(a int, b int, i string) int {
	var res int
	switch i {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res

}
func WriteResultInFile(result []byte, n string) {
	err := ioutil.WriteFile("./"+n+".txt", result, 0777)
	if err != nil {
		panic(err)
	}
}
