package test

import (
	"fmt"
	"github.com/gonutz/w32/v2"
	"log"
	"testing"
)

func TestGetCurrentAppVersion(t *testing.T) {
	t.Skip()
	executablePath := "D:\\Development\\wowtools\\wowtools.exe"
	size := w32.GetFileVersionInfoSize(executablePath)
	if size <= 0 {
		log.Fatalln("GetFileVersionInfoSize failed")
	}
	info := make([]byte, size)
	getFileInfo := w32.GetFileVersionInfo(executablePath, info)
	if !getFileInfo {
		log.Fatalln("GetFileVersionInfo failed")
	}
	fixed, getFileInfo := w32.VerQueryValueRoot(info)
	if !getFileInfo {
		log.Fatalln("VerQueryValueRoot failed")
	}
	fileVersion := fixed.FileVersion()
	versionString := fmt.Sprintf("%d.%d.%d.%d\n",
		fileVersion&0xFFFF000000000000>>48,
		fileVersion&0x0000FFFF00000000>>32,
		fileVersion&0x00000000FFFF0000>>16,
		fileVersion&0x000000000000FFFF>>0)
	fmt.Println(versionString)
}
