package gwc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Config is gwc app config
type Config struct {
	Filenames []string
	WC        bool
	LC        bool
}

type file struct {
	filename string
	words    uint64
	lines    uint64
	reader   io.Reader
	err      error
}

// GWC is root of this app
type GWC struct {
	lc, wc bool
	files  []*file
}

// New is a function that returns a GWC
func New(config *Config) *GWC {
	gwc := &GWC{
		lc: config.LC,
		wc: config.WC,
	}
	for _, filename := range config.Filenames {
		item := file{
			filename: filename,
		}
		stat, err := os.Stat(filename)
		if errors.Is(err, os.ErrNotExist) {
			item.err = errors.Unwrap(err)
		} else if stat.IsDir() {
			item.err = errors.New("is a directory")
		}
		f, _ := os.Open(filename)
		item.reader = f
		gwc.files = append(gwc.files, &item)
	}
	return gwc
}

// Compute is a method  of GWC that computes files
func (g *GWC) Compute() {
	var wg sync.WaitGroup
	for _, f := range g.files {
		if f.err != nil {
			continue
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, f *file) {
			if g.lc {
				s := bufio.NewScanner(f.reader)
				lineCount := uint64(0)
				for s.Scan() {
					lineCount++
				}
				f.lines = lineCount
				f.reader.(*os.File).Seek(0, 0)
			}
			if g.wc {
				s := bufio.NewScanner(f.reader)
				s.Split(bufio.ScanWords)
				wordCount := uint64(0)
				for s.Scan() {
					wordCount++
				}
				f.words = wordCount
			}
			wg.Done()
		}(&wg, f)
	}
	wg.Wait()
}

func (g *GWC) String() string {
	var b strings.Builder
	for _, f := range g.files {
		if g.lc {
			fmt.Fprintf(&b, "%d\t", f.lines)
		}
		if g.wc {
			fmt.Fprintf(&b, "%d\t", f.words)
		}
		fmt.Fprintf(&b, "%q", f.filename)
		if f.err != nil {
			fmt.Fprintf(&b, ": %v", f.err)
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}
