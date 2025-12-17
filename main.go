package main

import (
	"bufio"
	"fmt"
	"os"
)

// Функция MaskHTTPLinks принимает строку и маскирует все ссылки с "http://"
func MaskHTTPLinks(msg string) string {
	src := []byte(msg)            // исходный текст в байтах
	buf := make([]byte, len(src)) // создаём новый буфер такого же размера
	copy(buf, src)                // копируем исходные байты в буфер

	httpPrefix := []byte("http://") // префикс, который ищем
	plen := len(httpPrefix)

	i := 0
	for i < len(buf) {
		// проверяем, совпадает ли текущая позиция с "http://"
		if i+plen <= len(buf) {
			match := true
			for j := 0; j < plen; j++ {
				if buf[i+j] != httpPrefix[j] {
					match = false
					break
				}
			}

			if match {
				// нашли "http://", теперь маскируем все символы до пробела или конца строки
				k := i + plen
				for k < len(buf) && buf[k] != ' ' {
					buf[k] = '*'
					k++
				}
				i = k
				continue
			}
		}
		i++
	}

	return string(buf) // превращаем буфер обратно в строку
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')

	result := MaskHTTPLinks(text) // вызываем функцию маскировки
	fmt.Println("Результат:")
	fmt.Println(result) // выводим результат
} // Feature branch test
