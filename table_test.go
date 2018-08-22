package ffmt

import (
	"strings"
	"testing"
)

func TestTable1(t *testing.T) {
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
aaa3a`

	resu := strings.TrimSpace(strings.Join(FmtTable(ToTable(b[0], b[0], b[1], b[2], map[string]string{
		"Na": "aaa3a",
	})), "\n"))
	if resu != out {
		t.Fatal(resu)
	}
}

func TestTable2(t *testing.T) {
	b := []map[string]interface{}{
		{"Na": "1111", "Ba": "123123"},
		{"Na": "1", "Ba": "1231233231"},
		{"Na": "aaaa", "Ba": "1231231"},
	}

	out := `Ba         Na    
123123     1111  
1231233231 1     
1231231    aaaa  
           aaa3a`

	resu := strings.TrimSpace(strings.Join(FmtTable(ToTable(b[0], b[0], b[1], b[2], map[string]string{
		"Na": "aaa3a",
	})), "\n"))
	if resu != out {
		t.Fatal(resu)
	}
}

func TestTableText1(t *testing.T) {
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
		resu := strings.TrimSpace(TableText(v[0], "", " "))
		if resu != strings.TrimSpace(v[1]) {
			t.Fatal(resu)
		}
	}
}

func TestTableText2(t *testing.T) {
	tableData := [][2]string{
		{`
// A AA
// BBBB B
CC CC
`, `
//  A     AA 
//  BBBB  B  
CC CC
`},
	}

	for _, v := range tableData {
		resu := strings.TrimSpace(TableText(v[0], "//", " "))
		if resu != strings.TrimSpace(v[1]) {
			t.Fatal(resu)
		}
	}
}
