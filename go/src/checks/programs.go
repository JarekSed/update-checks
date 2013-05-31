package checks

import (
    "net/http"
    "code.google.com/p/go.net/html"
    "code.google.com/p/go.net/html/atom"
    "fmt"
    "strings"
)

var programs map[string]func()(string, error)

func init() {
    programs = make(map[string]func()(string,error))
    programs["3to2"] =  Get3to2Version
}

func GetAllPrograms() (map[string]func()(string,  error)) {
    return programs
}

func Get3to2Version() (string, error) {
    url := "https://pypi.python.org/pypi/3to2"
    resp, err := http.Get(url)
    if (err != nil) {
        return "", err
    }

    defer resp.Body.Close()

    tokenizer := html.NewTokenizer(resp.Body)
    for {
        tokenType := tokenizer.Next()
        if tokenType == html.ErrorToken {
            // we either can't parse the HTML
            return "", tokenizer.Err()
        // Process the current token.
        } else if tokenType == html.StartTagToken{
            token := tokenizer.Token()
            if token.DataAtom == atom.Lookup([]byte("a")){
                for _,attribute := range token.Attr {
                    if attribute.Key == "href" && strings.Contains(attribute.Val, "version"){
                        fmt.Printf("%v=%v\n",attribute.Key, attribute.Val)
                    }
                }
            } else{
                continue
            }
        }
    }

}

