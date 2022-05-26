package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

var MaxChars = 26
var k int = 2

func main() {

	var S string

	fmt.Print("Введите строку из малых латинских букв: ")
	fmt.Scan(&S)

	matched := IsLetter(S)

	if matched == false {
		log.Fatalln("Строка не соотвествует условию!")
	}

	kUniques(S, k)

}

func IsLetter(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

func isValid(count []int, k int) bool {

	val := 0

	for i := 0; i < MaxChars; i++ {
		if count[i] > 0 {
			val = val + 1
		}
	}

	return (k >= val)
}

func kUniques(s string, k int) {

	u := 0 //количество уникальных имволов

	n := utf8.RuneCountInString(s)

	var count []int = make([]int, MaxChars)

	for i := 0; i < n; i++ {
		if count[s[i]-'a'] == 0 {
			u = u + 1
			count[s[i]-'a']++
		}
	}

	// это знаит что все символы в строке одинаковые (по условию k<=2),
	//можно вывести ее длину и завершить программу
	if u < k {
		fmt.Print("Длина наибольшей подстроки: ")
		fmt.Print(utf8.RuneCountInString(s))
		os.Exit(0)
	}

	var (
		curr_start       int = 0
		curr_end         int = 0
		max_window_size  int = 1
		max_window_start int = 0
	)

	// Инициализируем count [] с нуля
	count = make([]int, len(count))
	count[s[0]-'a']++

	for i := 1; i < n; i++ {
		count[s[i]-'a']++
		curr_end++

		for isValid(count, k) == false {
			count[s[curr_start]-'a']--
			curr_start++
		}

		if curr_end-curr_start+1 > max_window_size {
			max_window_size = curr_end - curr_start + 1
			max_window_start = curr_start
		}

	}

	fmt.Print("Длина наибольшей подстроки: ")
	fmt.Println(max_window_size)
	fmt.Print("Наибольшая подстрока: ")
	fmt.Print(s[max_window_start : max_window_start+max_window_size])

}
