package token

import "testing"

// TestLookupIdent - test for LookupIdent function
func TestLookupIdent(t *testing.T) {
	for keyword, keywordType := range keywords {
		realKeyword := LookupIdent(keyword)
		if realKeyword != keywordType {
			t.Fatalf("keyword type wrong. expected=%q, got=%q",
				keywordType, realKeyword)
		}
	}
	// test identifier
	ident := LookupIdent("simple_identifier")
	if ident != IDENT {
		t.Fatalf("identifier was not recognized correctly")
	}
}
