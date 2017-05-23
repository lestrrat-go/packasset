package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	packasset "github.com/lestrrat/go-packasset"
	"github.com/lestrrat/go-packasset/internal/tty"
	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	var packageName string
	var output string
	var inputfile string
	flag.StringVar(&packageName, "package", "main", "Name of package of the generated file")
	flag.StringVar(&output, "output", "", "Name of file to generate")
	flag.StringVar(&inputfile, "input", "", "File with list of files to process")
	flag.Parse()

	// Input may come from either 1) set of files, or 2) stdin
	var files []string
	if !tty.IsTty(os.Stdin) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			files = append(files, scanner.Text())
		}
	} else if inputfile != "" {
		f, err := os.Open(inputfile)
		if err != nil {
			return errors.Wrapf(err, `packasset: failed to open input file %s`, inputfile)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			files = append(files, scanner.Text())
		}
	}

	g := packasset.NewGenerator(
		packasset.WithPackageName(packageName),
		packasset.WithFiles(files),
	)
	buf, err := g.Generate()
	if err != nil {
		return errors.Wrap(err, `packasset: failed to generate code`)
	}

	var out io.Writer
	if output == "" {
		out = os.Stdout
	} else {
		f, err := os.Create(output)
		if err != nil {
			return errors.Wrapf(err, `packasset: failed to create file %s`, output)
		}
		defer f.Close()
		out = f
	}

	for x := buf; len(x) > 0; {
		n, err := out.Write(x)
		x = x[n:]
		if err != nil && err != io.EOF {
			return errors.Wrapf(err, `packasset: failed to write to file %s`, output)
		}
	}
	return nil
}
