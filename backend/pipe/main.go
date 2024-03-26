package main

import (
	"backend/lib/envset"
	pipe "backend/pipe/lib"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen).PrintfFunc()

func RunScrap(storeName string, searchType string, brandName string, fileName string) {
	green("\n === Running scrap === \n")
	execCmd("ts-node", "-p", "../admin/tsconfig.js",
		"../admin/src/pipe/index.ts",
		"--store", storeName,
		"--searchType", searchType,
		"--brand", brandName,
		"--fileName", fileName,
	)
	green("\n === Scrap Done!! === \n")
}

func RunPreprocess(storeName string, searchType string, fileName string) {
	green("\n === Running Preprocess === \n")
	p := pipe.NewPreProcessor()
	p.Run(storeName, searchType, fileName)
	green("\n === Preprocess Done!! === \n")

}
func RunUpload(storeName string, searchType string, fileName string) {
	green("\n === Running Upload === \n")
	u := pipe.NewUploader()
	u.Run(storeName, searchType, fileName)
	green("\n === Upload Done!! === \n")
}

func All(storeName string, searchType string, brandName string, fileName string) {
	RunScrap(storeName, searchType, brandName, fileName)
	RunPreprocess(storeName, searchType, fileName)
	RunUpload(storeName, searchType, fileName)
}

func execCmd(args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func setFileName() *string {
	currentTime := time.Now()

	// Format the time according to the layout "060102T150405"
	formattedTime := currentTime.Format("060102T150405")
	fmt.Printf("a file name arg is empty. \n set default flename as %s\n", formattedTime)
	fileName := formattedTime + ".json"
	return &(fileName)
}
func main() {
	var run = flag.String("r", "", "run all|scrap|preprocess|upload")
	var store = flag.String("s", "", "store arg")
	var searchType = flag.String("t", "", "list|page")
	var brand = flag.String("b", "", "brand arg")
	var fileName = flag.String("f", "", "fileName arg")
	var mode = flag.String("m", "", "prod mode or local mode")

	// Parse flags
	flag.Parse()

	if *fileName == "" {
		fileName = setFileName()
	}
	if *mode == "prod" {
		os.Setenv("ProductionLevel", "local-rds")
		fmt.Println("Uploading Product Mode!!!")
	}
	envset.LoadEnv()

	switch *run {
	case "scrap":
		RunScrap(*store, *searchType, *brand, *fileName)
		// fmt.Println(*store, *searchType, *brand, *fileName)
	case "preprocess":
		RunPreprocess(*store, *searchType, *fileName)
		// fmt.Println(*store, *searchType, *fileName)
	case "upload":
		RunUpload(*store, *searchType, *fileName)
		// fmt.Println(*store, *searchType, *fileName)
	case "preprocess-upload":
		RunPreprocess(*store, *searchType, *fileName)
		RunUpload(*store, *searchType, *fileName)
	case "all":
		All(*store, *searchType, *brand, *fileName)
		// fmt.Println(*store, *searchType, *brand, *fileName)
	default:
		fmt.Printf("\n r is unmatched. got %s want all|scrap|preprocess|upload", *run)
	}

}