package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*********** Constants ***********/
const ctemplatePath string = "templates/c/"
const cpptemplatePath string = "templates/cpp/"
const keyword string = "$replaceme$" //this is the string that will be replaced in each template

/*********** Globals ***********/
var programName string

/****************** Template Files Begin ******************/
// Notes Template
//go:embed templates/programNotes.md
var notesFile string

// C++ files
//go:embed templates/cpp/Makefile
var cppMakefile string

//go:embed templates/cpp/source.cpp
var cppSource string

// C files
//go:embed templates/c/Makefile
var cMakefile string

//go:embed templates/c/source.c
var cSource string

/****************** Template Files End ********************/

func copyTempFile(source string, path string) {
	//create a file from a source string (one of the templates)
	//in the path provided, replacing the keyword in each sourceStr
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	source = strings.Replace(source, keyword, programName, -1)

	writer := bufio.NewWriter(file)
	writer.WriteString(source)
	writer.Flush()
}

func printHelp() {
	usageMsg := `Usage: cptemplate [options]
  -l string, case insensitive, defaults to C++
        Language template to use, supports Go, Python, C, and C++, (for C++ provide "cpp" as the arg)
  -n string
        Name of program. This is what $replaceme$ will be replaced with in all files.
  -p string
        Path to write template. Defaults to current dir "./"
  -N    Do not include a markdown notes file
  -d    print debug info

  example: "cptemplate -l c -n example" creates a folder named example
  containing example.c, exampleNotes.md, and a Makefile.` + "\r\n"
	fmt.Print(usageMsg)
}

func main() {

	//parameters
	var language string
	var path string
	var notes bool
	var debug bool

	//change default usage msg
	flag.Usage = printHelp

	flag.StringVar(&language, "l", "cpp", "language template to use")
	flag.StringVar(&path, "p", "./", "path to write template")
	flag.StringVar(&programName, "n", "example", "name of program")
	flag.BoolVar(&notes, "N", true, "include notes or dont")
	flag.BoolVar(&debug, "d", false, "print debug info")

	flag.Parse()

	//fix path if necessary
	if path[len(path)-1] != '/' {
		path = path + "/"
	}

	//no matter what we make the dir
	progFolderPath := (path + programName + "/")
	os.Mkdir(progFolderPath, 0755) //drwxr-xr-x

	if debug {
		log.Printf("Notes files\n %s", notesFile)
	}

	if notes {
		copyTempFile(notesFile, (progFolderPath + programName + "Notes.md"))
	}

	//now we make the source files
	switch strings.ToLower(language) {
	case "cpp":
		//write the makefile and source file
		copyTempFile(cppMakefile, (progFolderPath + "Makefile"))
		copyTempFile(cppSource, (progFolderPath + programName + ".cpp"))
	case "c":
		copyTempFile(cMakefile, (progFolderPath + "Makefile"))
		copyTempFile(cSource, (progFolderPath + programName + ".c"))
	}
}

// s := "some text $replaceme$ and $replaceme$"

// fmt.Println(s)

// s = strings.Replace(s, "$replaceme$", "replaced", -1)

// fmt.Println(s)

// print(sampleFile)
// os.Mkdir("./temp", 0755)

// fmt.Print(cppSource)
// fmt.Print(cppMakefile)

// fmt.Println()
// fmt.Println()

// cppMakefile = strings.Replace(cppMakefile, "$replaceme$", "example", -1)

// fmt.Print(cppMakefile)

// os.Mkdir("example", 0755)
// file, err := os.Create("example/example.cpp")
// if err != nil {
// log.Fatal(err)
// }
// defer file.Close()
//
// writer := bufio.NewWriter(file)
// writer.WriteString(cppSource)
// writer.Flush()
//
// file, err = os.Create("example/Makefile")
// if err != nil {
// log.Fatal(err)
// }
//
// writer = bufio.NewWriter(file)
// writer.WriteString(cppMakefile)
// writer.Flush()
