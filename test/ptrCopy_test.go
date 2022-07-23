package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestCopyPtrData(t *testing.T) {
	t.Skip()
	ptrFolder := "C:\\Program Files (x86)\\World of Warcraft\\_ptr_"
	retailFolder := "C:\\Program Files (x86)\\World of Warcraft\\_retail_\\"

	var folders [2]string
	folders[0] = "Interface"
	folders[1] = "WTF"

	removePtrSubDirs(folders, ptrFolder)

	for _, element := range folders {
		roboCmd := fmt.Sprintf("robocopy \"%s\" \"%s\" /s", retailFolder+element, ptrFolder+"\\"+element)
		fmt.Println(roboCmd)
		fmt.Sprintf("Copying %s to PTR dir", element)
		cmd := exec.Command("powershell", roboCmd)

		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func removePtrSubDirs(folderArray [2]string, dstFolder string) {
	for _, element := range folderArray {
		fmt.Printf("Removing %s\n", element)
		err := os.RemoveAll(dstFolder + "\\" + element)
		if err != nil {
			log.Fatal(err)
		}
	}
}
