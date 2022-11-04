package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/antchfx/xmlquery"
)

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
			xmlFile, err := os.Open(filepath.Join(dirPath, file.Name()))
			if err != nil {
				return count, err
			}
			doc, err := xmlquery.Parse(xmlFile)
			if err != nil {
				return count, err
			}
			updateSetName, ok := GetNameFromXML(doc)
			if ok && filepath.Base(xmlFile.Name()) != updateSetName+".xml" {
				fmt.Printf("Renamed: %s to %s \n", filepath.Base(xmlFile.Name()), updateSetName+".xls")
				os.Rename(xmlFile.Name(), filepath.Join(dirPath, updateSetName)+".xml")
				count++
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
