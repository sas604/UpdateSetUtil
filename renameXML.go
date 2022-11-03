package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/antchfx/xmlquery"
)

func main() {
	flag := 0
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if err != nil {
			fmt.Println(err)
		}
		if path.Ext(file.Name()) == ".xml" {
			xmlFile, err := os.Open(file.Name())
			if err != nil {
				panic(err)
			}
			doc, err := xmlquery.Parse(xmlFile)
			if err != nil {
				panic(err)
			}
			updateSetName := xmlquery.FindOne(doc, "//sys_remote_update_set/name")
			if updateSetName != nil {
				fmt.Printf("Renamed: %s to %s \n", xmlFile.Name(), updateSetName.InnerText()+".xls")
				os.Rename(xmlFile.Name(), updateSetName.InnerText()+".xml")
				flag++
			}

			defer xmlFile.Close()
		}

	}
	fmt.Printf("Renamed %d files\n", flag)

}
