// Package checks provides utilities for actually checking the version number of programs
package checks


import(
    "net/http"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "strings"
    "os"
)

type OutOfDateProgram struct {
    name string
    aurVersion string
    latestVersion string
}

func min(first int, second int)(int){
    if first < second {
        return first
    } else {
        return second
    }
}

// GetOutOfDatePrograms checks all the programs, and returns a list of names of the programs that are out of date
func GetOutOfDatePrograms() ([]OutOfDateProgram) {
    outOfDatePrograms := []OutOfDateProgram{}
    programs := GetAllPrograms()
    for name, versionFunction := range programs {
        // Get latest upstream version
        latestVersion, err := versionFunction()
        if err != nil {
            fmt.Fprintf(os.Stderr, "%v gave error %v when getting latest version\n",name, err)
            continue
        }
        // Get latest AUR version
        aurVersion, err := GetAurVersion(name)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%v gave error %v when getting AUR version\n",name, err)
            continue
        }
        if (versionCompare(latestVersion,aurVersion) != 0) {
            outOfDate := OutOfDateProgram{name,aurVersion,latestVersion}
            outOfDatePrograms = append(outOfDatePrograms, outOfDate)

        }
    }
    return outOfDatePrograms
}

// versionCompare compares two version numbers
// it returns a negative number if the first is larger , 0 if equal, and positive if the second is larger
func versionCompare(first string, second string) (int) {
    // Split version numbers into pieces
    firstFields := strings.Split(first, ".")
    secondFields := strings.Split(second, ".")


    // If the strings have different # of pieces, we compare based on the # of pieces they both have
    if len(first) != len(second) {
        length := min(len(firstFields), len(secondFields))
        // Re-join the sections into one string
        // We can just take a substring, because we want an equal number of sections, not an equal number of chars
        // for example 1.10.2 and 1.2.3, if we want the first 2 sections, we need 3 chars from the first and 2 from the second
        result := versionCompare(strings.Join(firstFields[:length],".") , strings.Join(secondFields[:length],"."))
        // If the strings are the same up until the length of the shorter one, the longer one should be considered larger (ie, 1.0.1 > 1.0)
        if result == 0 {
            return len(secondFields) - len(firstFields)
        // If the strings differed, use that result
        }else {
            return result
        }
    }

    // If we reach this point, the strings have the same number of sections. So, to compare them we
    // compare each section.
    for i,_ := range firstFields {
        if firstFields[i] < secondFields[i] {
            return 1
        } else if firstFields[i] > secondFields[i] {
            return -1
        }
    }
    // If we reach this point, the strings must be equal
    return 0
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

    //var data versionNumber
    var dataRaw interface{}
    err = json.Unmarshal(body, &dataRaw)
    data := dataRaw.(map[string]interface{})
    version := data["results"].(map[string]interface{})["Version"].(string)
    if (err != nil) {
        return "", err
    }

    index := strings.LastIndex(version, "-")
    version = version[:index]
    return version, nil
}
