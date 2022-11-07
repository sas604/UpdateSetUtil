package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/antchfx/xmlquery"
)

var AppFs = afero.NewOsFs()

func main() {
	dirPath := flag.String("p", ".", "Absolute path to the directory")
	flag.Parse()
	count, err := RenameFiles(*dirPath)
	if err != nil {
		fmt.Println(err)
	}
	if count == 1 {
		fmt.Printf("Renamed %d file\n", count)
	} else if count > 1 {
		fmt.Printf("Renamed %d files\n", count)
	} else {
		fmt.Println("No update sets with bad naming were found")
	}

}
func RenameFiles(dirPath string) (n int, err error) {
	count := 0
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return count, err
	}
	for _, file := range files {

		if path.Ext(file.Name()) == ".xml" {
			xmlFile, err := AppFs.Open(filepath.Join(dirPath, file.Name()))
			if err != nil {
				return count, err
			}
			doc, err := xmlquery.Parse(xmlFile)
			if err != nil {
				return count, err
			}
			updateSetName, ok := GetNameFromXML(doc)

			if ok && strings.TrimSuffix(filepath.Base(xmlFile.Name()), filepath.Ext(xmlFile.Name())) != updateSetName {
				newName := GetNewName(dirPath, updateSetName, 0)
				err := AppFs.Rename(xmlFile.Name(), filepath.Join(dirPath, newName)+".xml")
				if err != nil {
					fmt.Printf("Error renaming file : %s \n %s \n", filepath.Base(xmlFile.Name()), err)
				} else {
					fmt.Printf("Renamed: %s to %s \n", filepath.Base(xmlFile.Name()), newName+".xls")
					count++
				}

			}

			defer xmlFile.Close()
		}
	}
	return count, nil
}

func GetNameFromXML(f *xmlquery.Node) (name string, ok bool) {
	n := xmlquery.FindOne(f, "//sys_remote_update_set/name")
	if n != nil {
		return n.InnerText(), true
	}
	return "", false
}

// returns new file name based on passed string.
// if file with the same filename  exsit return name with appended number
func GetNewName(dirPath string, fileName string, i int) string {
	if _, err := AppFs.Stat(filepath.Join(dirPath, fileName) + ".xml"); errors.Is(err, os.ErrNotExist) {
		return fileName
	} else {
		if i > 0 {
			fileName = strings.TrimSuffix(fileName, "-"+fmt.Sprint(i))
		}
		i++
		fileName = fmt.Sprintf("%s-%d", fileName, i)
		return GetNewName(dirPath, fileName, i)
	}

}
