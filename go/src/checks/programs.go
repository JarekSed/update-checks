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
    programs["python-crontab"] = GetPythonCrontabVersion
    programs["pymetar"] = GetPyMetarVersion
    programs["php-mongo"] = GetPhpMongoVersion
}

// Gets the map of all program names to the function that gets their version
func GetAllPrograms() (map[string]func()(string,  error)) {
    return programs
}

// Gets version numbers from PyPi website
// Returns a string with the version, and an error (if any)
func getVersionFromPyPi(name string) (string, error) {
    url := "https://pypi.python.org/pypi/" + name
    resp, err := http.Get(url)
    if (err != nil) {
        return "", err
    }

    defer resp.Body.Close()

    tokenizer := html.NewTokenizer(resp.Body)
    // This loop exits when we find the version, or the tokenizer runs out of input
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

// Gets the latest version # of 3to2.
// Returns a string with the version, and an error (if any)
func Get3to2Version() (string, error) {
    return getVersionFromPyPi("3to2")
}

// Gets the latest version # of pymetar.
// Returns a string with the version, and an error (if any)
func GetPyMetarVersion() (string, error) {
    return getVersionFromPyPi("pymetar")
}

// Gets the latest version # of python-crontab.
// Returns a string with the version, and an error (if any)
func GetPythonCrontabVersion() (string, error) {
    return getVersionFromPyPi("python-crontab")
}



// Gets the latest version of the php mongodb drivers from pecl
// Returns a string with the version, and an error (if any)
func GetPhpMongoVersion() (string, error) {
    url := "http://pecl.php.net/package/mongo"
    resp, err := http.Get(url)
    if (err != nil) {
        return "", err
    }

    defer resp.Body.Close()

    tokenizer := html.NewTokenizer(resp.Body)
    // This loop exits when we find the version, or the tokenizer runs out of input
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
                    // If this link contains '/package/mongo/', it has the version #
                    if attribute.Key == "href" && strings.Contains(attribute.Val, "/package/mongo/"){
                        // The version # is everything after the last slash
                        index := strings.LastIndex(attribute.Val, "/")
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


