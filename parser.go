package tldparser

import (
	"bufio"
	"os"
	"strings"
)

type Parser struct {
	m map[string]struct{}
}

func FromFile(path string) (*Parser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m := make(map[string]struct{})

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		text := s.Text()
		if text == "" || text[0] == '#' {
			continue
		}
		m[strings.ToLower(text)] = struct{}{}
	}

	return &Parser{m: m}, s.Err()
}

func (p *Parser) TLD(hostname string) string {
	index := strings.LastIndex(hostname, ".")
	if index != -1 {
		tld := hostname[index+1:]
		if _, ok := p.m[tld]; ok {
			return tld
		}
	}
	return ""
}
