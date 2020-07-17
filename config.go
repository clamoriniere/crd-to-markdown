package main

import (
	"strings"

	flag "github.com/spf13/pflag"
)

// Config use to store documentation generation config
type Config struct {
	TypeFiles         []string
	CRDNames          SetStrings
	HeaderFilePath    string
	FooterFilePath    string
	ExternalLinksFile string
}

// NewConfig return new Config instance
func NewConfig() *Config {
	return &Config{
		CRDNames: SetStrings{
			Values: map[string]bool{},
		},
	}
}

// Parse use to parse configuration
func (c *Config) Parse() {
	flag.StringArrayVarP(&c.TypeFiles, "file", "f", []string{}, "list of types.go files")
	flag.VarP(&c.CRDNames, "crd-name", "n", "root crd stuct")
	flag.StringVar(&c.HeaderFilePath, "header", "header.md", ".md file to be include as header of the generate doc")
	flag.StringVar(&c.FooterFilePath, "footer", "footer.md", ".md file to be include as footer of the generate doc")
	flag.StringVar(&c.ExternalLinksFile, "links", "links.csv", "list of external structs document links")

	flag.Parse()
}

// SetStrings use to store a set value
type SetStrings struct {
	Values map[string]bool
}

// String return the value
func (s *SetStrings) String() string {
	var output []string
	for key := range s.Values {
		output = append(output, key)
	}
	return strings.Join(output, ",")
}

// Set use to store a value in the SetValues
func (s *SetStrings) Set(val string) error {
	s.Values[val] = true
	return nil
}

// Type return the SetValues type
func (s *SetStrings) Type() string {
	return "SetStrings"
}

// Find returns if the string is present in the SetValue
func (s *SetStrings) Find(value string) bool {
	_, ok := s.Values[value]
	return ok
}
