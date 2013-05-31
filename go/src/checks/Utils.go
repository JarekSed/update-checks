// Package checks provides utilities for actually checking the version number of programs
package checks


import(
    "net/http"
    "encoding/json"
    "io/ioutil"
    "container/list"
    "fmt"
)
// Type VersionNumber is a simple struct used to unmarshal JSON.
// All we need to pull out is the version number (stored as a string, because it may have dots, dashes, etc)
type versionNumber struct {
    version string
}

func CheckAllPrograms() (*list.List) {
    out_of_date := list.New()
    programs := GetAllPrograms()
    for name, versionFunction := range programs {
        latestVersion, err := versionFunction()
        if err != nil {
            fmt.Printf("%v gave error %v when getting latest version\n",name, err)
        }else {
            fmt.Printf("%v has version %v\n",name, latestVersion)
        }
    }
    return out_of_date
}

// GetAurVersion gets the current version in the Arch User Repository of a program.
// It returns the version number as a string, and any error encountered
func GetAurVersion(programName string) (string, error){
    url := "https://aur.archlinux.org/rpc.php?type=info&arg=" + programName
    resp, err := http.Get(url)
    if (err != nil){
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if (err != nil){
        return "", err
    }

    var data versionNumber
    err = json.Unmarshal(body, &data)
    if (err != nil) {
        return "", err
    }
    return data.version, nil
}
