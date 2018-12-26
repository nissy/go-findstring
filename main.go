package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
)

var (
	isRecursive = flag.Bool("r", false, "Search directory recursive")
	isHelp      = flag.Bool("h", false, "This help")
)

func main() {
	os.Exit(exitcode(run()))
}

func exitcode(err error) int {
	if err != nil {
		if _, perr := fmt.Fprintf(os.Stderr, "Error: %s\n", err); perr != nil {
			fmt.Println(err)
			panic(perr)
		}
		return 1
	}

	return 0
}

func run() (err error) {
	flag.Parse()

	if *isHelp {
		_, err := fmt.Fprintf(os.Stderr, "Usage: %s [options] directory\n", os.Args[0])
		flag.PrintDefaults()
		return err
	}

	var dir string
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	} else {
		return nil
	}

	if *isRecursive {
		err = filepath.Walk(dir, func(filename string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if err = printStringPosition(filename); err != nil {
					return err
				}
			}
			return nil
		})
	} else {
		err = printStringPosition(dir)
	}

	return err
}

func printStringPosition(dir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path.Clean(dir), nil, 0)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, v := range f.Decls {
				if ident, ok := v.(*ast.GenDecl); ok {
					if ident.Tok == token.IMPORT {
						continue
					}
				}
				ast.Inspect(v, func(n ast.Node) bool {
					if ident, ok := n.(*ast.BasicLit); ok {
						if ident.Kind == token.STRING {
							fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), ident.Value)
						}
					}
					return true
				})
			}
		}
	}

	return nil
}
