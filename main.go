package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dirPath := "./"
	outputFolder := "./512х512"

	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		os.Mkdir(outputFolder, os.ModePerm)
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	picturesNames := make([]string, 0)
	pictures := make([]fs.FileInfo, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		isPicture := filepath.Ext(file.Name()) == ".jpg" || filepath.Ext(file.Name()) == ".jpeg" || filepath.Ext(file.Name()) == ".png"
		if isPicture {
			picturesNames = append(picturesNames, file.Name())
			pictures = append(pictures, file)
			continue
		}
	}
	fmt.Println(picturesNames)
	for _, picture := range pictures {
		fmt.Println("Try to open ", picture.Name())

		imgFile, err := os.Open(picture.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer imgFile.Close()

		imgPng, err := jpeg.Decode(imgFile)
		if err != nil {
			log.Fatal(err)
		}
		imgPng = resize.Resize(512, 512, imgPng, resize.Bicubic)

		name := filepath.Base(picture.Name()[:len(picture.Name())-len(filepath.Ext(picture.Name()))])
		newFileName := fmt.Sprintf("%s %s%s", name, "512x512", ".png")
		outFilePath := filepath.Join(outputFolder, newFileName)
		imgOut, _ := os.Create(outFilePath)
		jpeg.Encode(imgOut, imgPng, nil)
		imgOut.Close()
	}
}
