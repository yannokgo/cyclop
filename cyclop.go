package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gomarkdown/markdown"
)

func main() {

	title := strings.Join(os.Args[:], ",")
	if title == "./cyclop" {
		screen()
	} else {
		cyclop(title)
	}
}

func cyclop(title string) {
	title = os.Args[1]
	filename := "../" + title + ".md"
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	debut, err := ioutil.ReadFile("debut.html")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fin, err := ioutil.ReadFile("fin.html")
	if err != nil {
		log.Fatal(err)
	}
	html := markdown.ToHTML(content, nil, nil)
	headEnd := "<title>" + title + " | okgo.DEV</title></head><body id='body'>"

	myHTMLstring := string(debut[:]) + headEnd + string(html[:]) + string(fin[:])

	toHTML := "Generated/okgoDEV/static/html/"

	f, err := os.Create(toHTML + title + ".html")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(myHTMLstring)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "HTML bytes written successfully to", toHTML)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	toMARKDOWN := "Generated/markdown/"

	f, err = os.Create(toMARKDOWN + title + ".md")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err = f.WriteString(string(content[:]))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "MARKDOWN bytes written successfully to", toMARKDOWN)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 	err = os.Remove(filename)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	    }
}

func screen() {
	app := "chromium"

	arg0 := "--app=file:///home/yannb/Desktop/cyclop/Generated/okgoDEV/static/index.html"
	arg1 := ""
	arg2 := ""
	arg3 := ""

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
