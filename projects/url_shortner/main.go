package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"www.linkedin.com/in/avdhutpatil2309"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

/* mapping
  d9736711 -->{
            	ID : 9736711
          		OriginalURL :"original_link"
				ShortURL : " 9736711"
				CreationDate : time.Now()
            }
*/

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // it converts the originalURL string to a byte slice
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data : ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("EncodeToString : ", hash)
	fmt.Println("final string : ", hash[:8])
	return hash[:8]

}

var urlDB = make(map[string]URL)

func createURl(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL //using short url as a ID for simplicity
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}
func RootPageURL(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Get method")
	fmt.Fprintf(w, "Jai Shree Ram") // not write in output window only but wirtes on writter provided here,w
}

func main() {
	fmt.Println("Starting URL shortner... ")
	OriginalURL := "www.linkedin.com/in/avdhutpatil2309"
	generateShortURL(OriginalURL)

	// register the handler function to handle request to the root URL ("/")
	http.HandleFunc("/", RootPageURL)

	// start the http server on port 3000
	fmt.Println("start the server on port no 3000... ")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("error on starting server : ", err)
		// return
	}
}
