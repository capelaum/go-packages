package titulos

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Feito no curso da cod3r - aula 75-76

// Titulos obtem varios títulos de páginas html através de suas urls
func Titulos(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := ioutil.ReadAll(response.Body)

			// cria um objeto de Regex que pode ser usado para combinar com uma string
			r, _ := regexp.Compile("<title>(.*?)</title>")

			// cria um slice de strings que combinam entre o objeto regex e o parametro html em formato string
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
