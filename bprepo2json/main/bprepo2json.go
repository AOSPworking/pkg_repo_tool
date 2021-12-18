package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AOSPworking/pkg_repo_tool/bprepo2json"
	"github.com/AOSPworking/pkg_repo_tool/mod"
	"github.com/AOSPworking/pkg_repo_tool/parser"
	log "github.com/sirupsen/logrus"
)

var (
	// Repositories is repos related
	Repositories map[string]*bprepo2json.Repository
)

func main() {
	flag.Parse()
	Repositories = bprepo2json.NewRepositories(*repoList)

	if flag.NArg() == 0 {
		log.Fatalln("main error: flag.NArg() == 0")
		os.Exit(exitCode)
	}
	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
			os.Exit(exitCode)
		case dir.IsDir():
			walkDir(path)
		default:
			processPackage(path)
		}
	}
	buf, err := json.MarshalIndent(Repositories, "", "  ")
	if err != nil {
		log.Fatalln("main error:", err.Error())
	}
	bprepo2json.WriteJSON(*outputFile, buf)
	os.Exit(exitCode)
}

func walkDir(path string) {
	visitFile := func(path string, f os.FileInfo, err error) (e error) {
		if f.IsDir() {
			return err
		}
		if err != nil {
			report(err)
			return err
		}
		err = processPackage(path)
		return err
	}
	filepath.Walk(path, visitFile)
}

func processPackage(filepath string) error {
	pkg, err := processFile(filepath)
	if err != nil {
		report(err)
		return err
	}
	pkgName := bprepo2json.GetPkgNameByFileName(filepath)
	repoName := bprepo2json.GetRepoNameByPkgName(pkgName, Repositories)
	pkg.Repo = Repositories[repoName]

	Repositories[repoName].Packages = append(Repositories[repoName].Packages, pkg)
	return nil
}

func processFile(filename string) (pkg bprepo2json.Package, err error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Warningln("processFile error:", err.Error())
		return
	}
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	r := bytes.NewBuffer(src)
	file, errs := parser.Parse(filename, r, parser.NewScope(nil))
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err)
		}
		return bprepo2json.Package{}, fmt.Errorf("%d parsing errors", len(errs))
	}

	pkg.Name = bprepo2json.GetPkgNameByFileName(filename)
	for _, def := range file.Defs {
		// only process module, do not process assignment
		if modu, ok := def.(*parser.Module); ok {
			bpModu, err := processModule(modu)
			if err != nil {
				// not allowed type is accepted
				continue
			}
			bpModu.Pkg = &pkg
			pkg.Modules = append(pkg.Modules, bpModu)
			continue
		}
	}
	return pkg, nil
}

// only process module
func processModule(modu *parser.Module) (res bprepo2json.Module, err error) {
	// judge module type whether in FrameworkAllowModuleType
	if ok, contain := mod.FrameworkAllowModuleType[modu.Type]; !ok || !contain {
		log.Debugln("processFileModule info:", modu.Type, "is not allowed")
		return res, errors.New("processFileModule info: type not allowd")
	}

	res = bprepo2json.Module{}
	res.Type = modu.Type
	for _, prop := range modu.Map.Properties {
		// now only process module name
		switch prop.Name {
		case "name":
			processModuleName(prop.Value, &res)
		}
	}
	return
}

func processModuleName(name parser.Expression, modu *bprepo2json.Module) error {
	switch v := name.(type) {
	case *parser.String:
		modu.Name = strings.Trim(strconv.Quote(v.Value), "\"")
	}
	return nil
}
