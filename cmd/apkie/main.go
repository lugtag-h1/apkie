package main

import (
	apkie "github.com/lugtag-h1/apkie"
	"github.com/lugtag-h1/apkie/helper/files"
	"github.com/docopt/docopt.go"
	"github.com/fatih/color"
	"log"
	"os"
)

const (
	usage = `APK IE - Exported components checker

Usage:
	apkie <filename> <component-name>
	apkie -h | --help
	apkie --version

Options:
	-h --help     Show this screen.
	--version     Show version.
`
	version = "0.1.0"
)

func init() {
	log.SetPrefix("[+] ")
	log.SetFlags(0)
}

func main() {
	var args struct {
		Filename  string `docopt:"<filename>"`
		Component string `docopt:"<component-name>"`
	}

	opts, err := docopt.DefaultParser.ParseArgs(usage, os.Args[1:], version)
	if err != nil {
		return
	}

	opts.Bind(&args)

	log.Printf("Opening '%s'...\n", args.Filename)
	if !files.Exists(args.Filename) {
		log.Fatalln(color.RedString("File does not exist"))
	}

	log.Printf("Checking for valid APK...")
	if !files.IsValidApk(args.Filename) {
		log.Fatalln(color.RedString("The file doesn't appear a valid APK package"))
	}

	log.Println("Parsing AndroidManifest.xml...")
	manifest, err := apkie.ReadAndroidManifest(args.Filename)
	if err != nil {
		log.Fatalln(color.RedString(err.Error()))
	}

	log.Println("Searching for exported components...")
	components := apkie.FindExportedComponents(manifest, args.Component)

	if len(components) == 0 {
		log.Println(color.YellowString("No component named '%s' has been found\n", args.Component))
	} else {
		for _, component := range components {
			if component.IsExported {
				log.Println(color.GreenString("The component '%s' is exported", component.Name))
			} else {
				log.Println(color.RedString("The component '%s' is explicitly not exported", component.Name))
			}
		}
	}
}
