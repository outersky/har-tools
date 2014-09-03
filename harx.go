package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Har struct {
	Log HLog
}

type HLog struct {
	//Version string
	Entries []HEntry
}

type HEntry struct {
	//StartedDateTime string
	//Time int
	Request  HRequest
	Response HResponse
}

type HRequest struct {
	Url    string
	Method string
}

type HResponse struct {
	Content HContent
}

type HContent struct {
	Size     int
	MimeType string
	Text     string
}

func (e *HEntry) dump(dir string) {
	u, err := url.Parse(e.Request.Url)
	if err != nil {
		log.Fatal(err)
	}
	scheme, host, path := u.Scheme, u.Host, u.Path // host,path = www.google.com , /search.do
    if scheme=="chrome-extension" {
        return  // ignore all chrome-extension requests.
    }
	if i := strings.LastIndex(host, ":"); i != -1 {
		host = host[0:i] // remove port
	}
	path = dir + "/" + host + path
	if j := strings.LastIndex(path, "/"); j != -1 {
		os.MkdirAll(path[0:j], os.ModePerm)
		e.Response.Content.writeTo(path)
	}
}

func decode(str []byte, fileName string) {
	data, err := base64.StdEncoding.DecodeString(string(str))
	if err != nil {
		log.Fatal(err)
	} else {
		ioutil.WriteFile(fileName, data, os.ModePerm)
	}
}

func (c *HContent) writeTo(f string) {
	if strings.Index(c.MimeType, "text") != -1 || strings.Index(c.MimeType, "javascript") != -1 || strings.Index(c.MimeType, "json") != -1 {
		ioutil.WriteFile(f, []byte(c.Text), os.ModePerm)
	} else {
		decode([]byte(c.Text), f)
	}
}

func (c *HContent) writeToFile(f *os.File) {
	f.WriteString(c.Text)
	f.Close()
}

func handle(r *bufio.Reader) {
	dec := json.NewDecoder(r)
	var har Har
	err := dec.Decode(&har)
	if err != nil {
		log.Fatal(err)
		os.Exit(-2)
	} else {
		for index, entry := range har.Log.Entries {
			output(index, entry)
		}
	}
}

func output(index int, entry HEntry) {
	if list {
		if urlPattern != nil {
			if len(urlPattern.FindString(entry.Request.Url)) > 0 {
				listEntries(index, entry)
			}
		} else {
			listEntries(index, entry)
		}
	} else if extract {
		if index == extractIndex {
			extractOne(entry)
		}
	} else if extractPattern {
		if len(urlPattern.FindString(entry.Request.Url)) > 0 {
			entry.dump(dir)
		}
	} else if extractAll {
		entry.dump(dir)
	}
}

func listEntries(index int, entry HEntry) {
	fmt.Printf("[%3d][%6s][%25s][Size:%8d][URL:%s]\n", index, entry.Request.Method, entry.Response.Content.MimeType, entry.Response.Content.Size, entry.Request.Url)
}

func extractOne(entry HEntry) {
	fmt.Print(entry.Response.Content.Text)
}

var list bool = false

var extract bool = false
var extractIndex int = -1

var extractPattern bool = false
var urlPattern *regexp.Regexp = nil

var extractAll bool = false
var dir string = ""

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
usage: harx [options] har-file
    -l                  List files , lead by [index]
    -lp urlPattern      like -l , but filter with urlPattern
    -a dir              extract All content to [dir]
    -i Index            extract the [index] content , need run with -l first to get [index]
    -p urlPattern dir   like -a , but filter with urlPattern

        `)
		return
	}

	var fileName string
	switch os.Args[1] {
	case "-l":
		list = true
		fileName = os.Args[2]
	case "-lu":
		list = true
		urlPattern = regexp.MustCompile(os.Args[2])
		fileName = os.Args[3]
	case "-i":
		extract = true
		extractIndex, _ = strconv.Atoi(os.Args[2])
		fileName = os.Args[3]
	case "-u":
		extractPattern = true
		urlPattern = regexp.MustCompile(os.Args[2])
		dir = os.Args[3]
		fileName = os.Args[4]
	case "-a":
		extractAll = true
		dir = os.Args[2]
		fileName = os.Args[3]
	}

	file, err := os.Open(fileName)

	if err == nil {
		handle(bufio.NewReader(file))
	} else {
		fmt.Printf("Cannot open file : %s\n", fileName)
		log.Fatal(err)
		os.Exit(-1)
	}

	if extractPattern || extractAll {
		os.MkdirAll(dir, os.ModePerm)
		//os.Chdir(dir)
	}
}
