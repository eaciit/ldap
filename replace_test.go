package ldap

import (
	"strings"
	"testing"
)

var dntests = []string{
`unescaped`,			`unescaped`,
`James "Jim" Smith, III`,	`James \"Jim\" Smith\, III`,
` front space`,			`\ front space`,
`tail space `,			`tail space\ `,
` both space `,			`\ both space\ `,
`#hashtag`,			`\#hashtag`,
`gathsah#`,			`gathsah\#`,
`#bothtag#`,			`\#bothtag\#`,
`#hashnspace `,			`\#hashnspace\ `,
` spacenhash#`,			`\ spacenhash\#`,
``,``}

var filtertests = []string{
`unescaped`,		`unescaped`,
`\o/`,				`\5co/`,
`*_*`,				`\2a_\2a`,
`sunn o)))`,		`sunn o\29\29\29`,
`(((o nnus`,		`\28\28\28o nnus`,
"\000\\*()",		`\00\5c\2a\28\29`,
``,``}

func TestDnReplacer(t *testing.T) {
	for i := 0; i < len(dntests); i += 2 {
		out := DnReplace(dntests[i])
		if strings.Compare(dntests[i+1], out) != 0  {
			t.Errorf("Expected: %v Got: %v\n", dntests[i+1], out)
		}
	}
}

func TestFilterReplacer(t *testing.T) {
	for i := 0; i < len(filtertests); i += 2 {
		out := FilterReplace(filtertests[i])
		if strings.Compare(filtertests[i+1], out) != 0  {
			t.Errorf("Expected: %v Got: %v\n", filtertests[i+1], out)
		}
	}
}
