package main

import (
	"fmt"

	pdfcpuAPI "github.com/hhrutter/pdfcpu/pkg/api"
	pdfcpuConfig "github.com/hhrutter/pdfcpu/pkg/pdfcpu"
)

func main() {
	config := pdfcpuConfig.NewDefaultConfiguration()

	config.UserPW = "123"
	config.OwnerPW = "123"

	cmd := pdfcpuAPI.EncryptCommand("./data/test.pdf", "./data/testEncrypt.pdf", config)
	_, err := pdfcpuAPI.Process(cmd)

	if err != nil {
		fmt.Println(err)
	}
}
