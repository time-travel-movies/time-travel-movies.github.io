package main

import (
	"bufio"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	errTSV = errors.New("need 5 values on a line, found less")
)

func main() {
	ee, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if err := write(os.Stdout, "cmd/gensite/template.html", ee); err != nil {
		log.Fatal(err)
	}
}

type entry struct {
	Title       string
	Year        string
	Image       string
	IMDBLink    string
	Description string
}

func parse(r io.Reader) ([]*entry, error) {
	var ee []*entry
	s := bufio.NewScanner(r)
	for s.Scan() {
		e, err := parseEntry(s.Text())
		if err != nil {
			return nil, err
		}
		ee = append(ee, e)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return ee, nil
}

func parseEntry(s string) (*entry, error) {
	ss := strings.Split(s, "\t")
	if len(ss) < 5 {
		return nil, errTSV
	}
	e := &entry{
		Title:       ss[0],
		Year:        ss[1],
		Image:       ss[2],
		IMDBLink:    ss[3],
		Description: ss[4],
	}
	return e, nil
}

func write(w io.Writer, templateFile string, ee []*entry) error {
	f, err := os.Open(templateFile)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	t, err := template.New("index").Parse(string(b))
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, ee)
}
