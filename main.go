//-*- coding: utf-8 -*-

package main

import (
	"golang.org/x/text/encoding/charmap"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"time"
)

func toUtf8(inBuf []byte) string {
	buf := inBuf
	buf, _ = charmap.Windows1252.NewDecoder().Bytes(inBuf)
	stringVal := strings.TrimSpace(string(buf))
	return stringVal
}

func copyFileWithoutHeader(sourceFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("Fil saknas: ", sourceFile)
		return err
	}
	defer source.Close()
	
	// Läs header
	buffer := make([]byte, 117)
	_, err = source.Read(buffer)
	if err != nil && err != io.EOF {  
		return err
	}

	header, err := decodeHeader(buffer)
	fmt.Println("Version: '"+header.version+"'")
	fmt.Println("Ursprunglig sökväg och filnamn: '"+header.path+"'")
	fmt.Println("Siffervärde: '"+header.numbervalue+"'")

	parsedTime, err := time.Parse("060102150405", header.created)
	if err != nil {
		return err
	}
        formattedDate := parsedTime.Format("06-01-02 15:04:05")	
	fmt.Println("Skapad tidpunkt:", formattedDate)
	fmt.Println("MDB-filens storlek i bytes: ", header.filesize)

	pathparts := strings.Split(header.path, "\\")
	filename := pathparts[len(pathparts)-1]
	fmt.Println("Återställer till filnamn: '"+filename+"'")

	destination, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer destination.Close()
	
	// Kopiera resten av filen
	_, err = io.Copy(destination, source)
	
	return err
}

type HeaderStruct struct {
	version  string
	path string
	numbervalue string
	created string
	filesize int
}

func decodeHeader(input []byte) (*HeaderStruct, error) {
	if len(input) < 116 {
		return nil, fmt.Errorf("byte-array must be at least 116 bytes long")
	}
	
	firstBytes := input[:7]
	version := toUtf8(firstBytes)
	
	secondBytes := input[7:87]
	path := toUtf8(secondBytes)
	
	secondBytes = input[87:93]
	numbervalue := toUtf8(secondBytes)
	
	secondBytes = input[93:105]
	created := toUtf8(secondBytes)
	
	secondBytes = input[105:]
	filesize, err := strconv.Atoi(toUtf8(secondBytes))
	if err != nil {
		fmt.Println("Trasig filstorlek: ", secondBytes)
	}
	
	return &HeaderStruct{
		version:  version,
		path: path,
		numbervalue: numbervalue,
		created: created,
		filesize: filesize,
	}, nil
}


func main() {
	helpPtr := flag.Bool("help", false, "Skriv ut hjälptext.")
	help2Ptr := flag.Bool("?", false, "Skriv ut hjälptext.")
	
	flag.Parse()
	
	if *helpPtr || *help2Ptr {
		flag.Usage()
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Println("Ange ett filnamn")
		flag.Usage()
		os.Exit(1)
	}
	
	fmt.Println("Angett filnamn: ", os.Args[1])
	if !strings.HasSuffix(os.Args[1], ".HBK") {
		fmt.Println("Argumentet ska vara ett filnamn som slutar på .HBK")
		flag.Usage()
		os.Exit(1)
	}
	
	err := copyFileWithoutHeader(os.Args[1])
	if err != nil {
		panic(err)
	}
	/*
	byteArray := []byte("abcdefghijklmnopqrstuvwxyz1234567890!@#$%^") // Ett exempelbyte-array på 50 bytes
	myStruct, err := byteArrayToStruct(byteArray)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Första strängen:", myStruct.FirstString)
	fmt.Println("Andra strängen:", myStruct.SecondString)
	*/
}
