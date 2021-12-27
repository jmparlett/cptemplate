package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/*********** Constants ***********/
const ctemplatePath string = "templates/c/"
const cpptemplatePath string = "templates/cpp/"
const keyword string = "$replaceme$" //this is the string that will be replaced in each template
const dateWord string = "$date$"     //this will be replaced with the current date in any template file

/*********** Globals ***********/
var programName string
var date string
var programStartTime time.Time = time.Now()

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

// Go files
//go:embed templates/go/Makefile
var goMakefile string

//go:embed templates/go/source.go
var goSource string

// Python files
//go:embed templates/python/Makefile
var pyMakefile string

//go:embed templates/python/source.py
var pySource string

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
	source = strings.Replace(source, dateWord, date, -1)

	writer := bufio.NewWriter(file)
	writer.WriteString(source)
	writer.Flush()
}

func printHelp() {
	usageMsg := `Usage: cptemplate [programName] [language] [options]
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

func permuteArgs(args []string) int {
	//rearrange the args array so that named arguments come first, this allows us to use positionals and named args
	//we return the index of start of non flag args
	args = args[1:] //remember changing the slice changes the underlying array
	// optind := 0
	// for i := range args {
	// if args[i][0] == '-' {
	// tmp := args[i]
	// args[i] = args[optind]
	// args[optind] = tmp
	// optind++
	// }
	// }
	j := 0 //our swap index
	i := 0 //our args index
	for i < len(args) {
		if args[i][0] == '-' { //if its a dash we want to send it and its neighbor to the back
			tmp := args[i]
			args[i] = args[j]
			args[j] = tmp
			j++
			i++

			//now the neighbor
			if i < len(args) {
				tmp := args[i]
				args[i] = args[j]
				args[j] = tmp
				j++
			}
		}
		i++
	}
	return j + 1 //pos of named arg
}

func cleanAndDie(path string) {
	file, err := os.Stat(path)
	//don't want to delete somthing we didn't create
	if file.ModTime().After(programStartTime) {
		err = os.RemoveAll(path)
		if err != nil {
			log.Fatal("Could cleanup")
		}
	}
}
func main() {
	//parameters
	var language string
	var path string
	var notes bool
	var debug bool

	//change default usage msg
	flag.Usage = printHelp

	//permute args reg to rearrange named args
	namedArgsPos := permuteArgs(os.Args)

	flag.StringVar(&language, "l", "cpp", "language template to use")
	flag.StringVar(&path, "p", "./", "path to write template")
	flag.StringVar(&programName, "n", "", "name of program")
	flag.BoolVar(&notes, "N", true, "include notes or dont")
	flag.BoolVar(&debug, "d", false, "print debug info")

	flag.Parse()

	if flag.NArg() != 0 { //we may have named args
		if namedArgsPos < len(os.Args) { //first should be prog name
			programName = os.Args[namedArgsPos]
			namedArgsPos++
		}
		if namedArgsPos < len(os.Args) { //second should be lang
			language = os.Args[namedArgsPos]
			namedArgsPos++
		}
	}

	if programName == "" { //if no program name provided print usage and exit
		fmt.Println("Error: A program name is required \"-n <name>\"")
		flag.Usage()
		os.Exit(1)
	}

	//set date time string
	date = fmt.Sprint(programStartTime.Year(), "-", programStartTime.Month(), "-", programStartTime.Day())

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
	case "go":
		copyTempFile(goMakefile, (progFolderPath + "Makefile"))
		copyTempFile(goSource, (progFolderPath + programName + ".go"))
	case "python":
		copyTempFile(pyMakefile, (progFolderPath + "Makefile"))
		copyTempFile(pySource, (progFolderPath + programName + ".py"))
	case "none":
		fmt.Println("Info: \"none\" provided, creating folder with notes only")
	default:
		fmt.Println("Info: language not supported, cleaning up and exiting")
		cleanAndDie(progFolderPath)
		os.Exit(1)
	}
}
