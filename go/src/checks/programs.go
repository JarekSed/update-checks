package checks

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"net/http"
	"strings"
)

var programs map[string]func() (string, error)

func init() {
	programs = make(map[string]func() (string, error))
	programs["3to2"] = Get3to2Version
	programs["python-crontab"] = GetPythonCrontabVersion
	programs["pymetar"] = GetPyMetarVersion
	programs["php-mongo"] = GetPhpMongoVersion
	programs["vim-blockcomment"] = GetVimBlockCommentVersion
	programs["vim-commentop"] = GetVimCommentOpVersion
	programs["vim-neocomplcache"] = GetVimNeoComplCacheVersion
	programs["vim-python"] = GetVimPythonVersion
}

// Gets the map of all program names to the function that gets their version
func GetAllPrograms() map[string]func() (string, error) {
	return programs
}

// Gets latest version numbers of vim plugins from vim.org
// Takes one argment, the ID of the script on vim.org
// Returns a string with the version, and an error (if any)
func getVersionFromVimDotOrg(scriptID string) (string, error) {
	url := "http://www.vim.org/scripts/script.php?script_id=" + scriptID
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	tokenizer := html.NewTokenizer(resp.Body)
	// vim.org doesn't annotate their html entities very well,
	// so we use this variable to keep track of which column in the table we are looking at
	// Version #'s are in the second column
	columnInDataTable := 0
	// This loop exits when we find the version, or the tokenizer runs out of input
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			// we either can't parse the HTML, or we're done
			// In either case, we haven't found a good version
			return "", tokenizer.Err()
		case html.StartTagToken:
			token := tokenizer.Token()
			// If this is a table data, it might be part of the data table
			if token.DataAtom == atom.Lookup([]byte("td")) {
				for _, attribute := range token.Attr {
					// If this is annotated with class=rowodd or roweven, this is a field in the data table
					if attribute.Key == "class" &&
						(strings.Contains(attribute.Val, "rowodd") || strings.Contains(attribute.Val, "roweven")) {
						// We have seen one more field in the data table
						columnInDataTable++
					}
				}
			}
			break
		case html.EndTagToken:
			// If this is the end of a table row, we reset the number of data fields seen
			if tokenizer.Token().DataAtom == atom.Lookup([]byte("tr")) {
				columnInDataTable = 0
			}
			break
		case html.TextToken:
			token := tokenizer.Token()
			// If this is the second column in the table, it is the version column.
			// Because vim.org sorts the data table with the most recent version at the top,
			// we can return the first version we find, as it must be the most recent.
			if columnInDataTable == 2 && strings.TrimSpace(token.String()) != "" {
				return token.String(), nil
			}
			break
		}
	}
}

// Gets version numbers from PyPi website
// Returns a string with the version, and an error (if any)
func getVersionFromPyPi(name string) (string, error) {
	url := "https://pypi.python.org/pypi/" + name
	resp, err := http.Get(url)
	if err != nil {
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
		} else if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			// Find the link with the version #
			if token.DataAtom == atom.Lookup([]byte("a")) {
				for _, attribute := range token.Attr {
					// If this link contains 'version=', it has the version #
					if attribute.Key == "href" && strings.Contains(attribute.Val, "version=") {
						// The version # is everything after the last equals sign
						index := strings.LastIndex(attribute.Val, "=")
						version := attribute.Val[index+1:]
						return version, nil
					}
				}
			} else {
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
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	// This loop exits when we find the version, or the tokenizer runs out of input
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			// we either can't parse the HTML, or we're done
			// In either case, we haven't found a good version
			return "", tokenizer.Err()
		case html.StartTagToken:
			token := tokenizer.Token()
			// Find the link with the version #
			if token.DataAtom == atom.Lookup([]byte("a")) {
				for _, attribute := range token.Attr {
					// If this link contains '/package/mongo/', it has the version #
					if attribute.Key == "href" && strings.Contains(attribute.Val, "/package/mongo/") {
						// The version # is everything after the last slash
						index := strings.LastIndex(attribute.Val, "/")
						version := attribute.Val[index+1:]
						return version, nil
					}
				}
			}
			break
		}
	}
}

// Gets the latest version of the vim BlockComment plugin
// Returns a string with the version, and an error (if any)
func GetVimBlockCommentVersion() (string, error) {
	return getVersionFromVimDotOrg("473")

}

// Gets the latest version of the vim python plugin
// Returns a string with the version, and an error (if any)
func GetVimPythonVersion() (string, error) {
	return getVersionFromVimDotOrg("790")

}

// Gets the latest version of the vim neocomplcache plugin
// Returns a string with the version, and an error (if any)
func GetVimNeoComplCacheVersion() (string, error) {
	return getVersionFromVimDotOrg("2620")

}

// Gets the latest version of the vim comment-top plugin
// Returns a string with the version, and an error (if any)
func GetVimCommentOpVersion() (string, error) {
	return getVersionFromVimDotOrg("2708")

}
