package ffmt

import (
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	b := []struct {
		Na string
		Ba string
	}{
		{"1111", "123123"},
		{"1", "1231233231"},
		{"aaaa", "1231231"},
	}

	out := `Na    Ba         
1111  123123     
1     1231233231 
aaaa  1231231    
aaa3a            `

	if strings.Join(FmtTable(ToTable(b[0], b[0], b[1], b[2], map[string]string{
		"Na": "aaa3a",
	})), "\n") != out {
		t.Fail()
	}
}

func TestTableText(t *testing.T) {
	tableData := [][2]string{
		{`
A AA
BBBB B
`, `
A     AA 
BBBB  B
`},
	}

	for _, v := range tableData {
		if strings.TrimSpace(TableText(v[0], "", " ")) != strings.TrimSpace(v[1]) {
			t.Fail()
		}
	}
}
