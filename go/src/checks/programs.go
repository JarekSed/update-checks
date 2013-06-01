package checks

import (
    "net/http"
    "code.google.com/p/go.net/html"
    "code.google.com/p/go.net/html/atom"
    "strings"
)

var programs map[string]func()(string, error)

func init() {
    programs = make(map[string]func()(string,error))
    programs["3to2"] =  Get3to2Version
}

// Gets the map of all program names to the function that gets their version
func GetAllPrograms() (map[string]func()(string,  error)) {
    return programs
}

// Gets the latest version # of 3to2.
// Returns a string with the version, and an error (if any)
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
            // we either can't parse the HTML, or we're done
            // In either case, we haven't found a good version
            return "", tokenizer.Err()
        } else if tokenType == html.StartTagToken{
            token := tokenizer.Token()
            // Find the link with the version #
            if token.DataAtom == atom.Lookup([]byte("a")){
                for _,attribute := range token.Attr {
                    // If this link contains 'version=', it has the version #
                    if attribute.Key == "href" && strings.Contains(attribute.Val, "version="){
                        // The version # is everything after the last equals sign
                        index := strings.LastIndex(attribute.Val, "=")
                        version := attribute.Val[index+1:]
                        return version, nil
                    }
                }
            } else{
                continue
            }
        }
    }

}

