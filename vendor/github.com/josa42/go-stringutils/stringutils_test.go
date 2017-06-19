package stringutils

import "testing"

func TestTrimPrefix(t *testing.T) {

	str1 := TrimPrefix("")

	if str1 != "" {
		t.Error("TrimPrefix() should no nothing on empty string")
	}

	str2 := TrimPrefix(`
		Foo
	`)

	if str2 != "Foo" {
		t.Error("TrimPrefix() should strip prefix")
	}

	str3 := TrimPrefix(`
		{
		  "Foo": "bar"
		}
	`)

	if str3 != "{\n  \"Foo\": \"bar\"\n}" {
		t.Error("TrimPrefix() should strip prefix width diffrent prefixes")
	}

}

func TestTrimLeadingTabs(t *testing.T) {
	str1 := TrimLeadingTabs(`
		{
		  "Foo": "bar"
		}
	`)

	if str1 != "{\n  \"Foo\": \"bar\"\n}" {
		t.Error("TrimLeadingTabs() should trim all leading tabs")
	}
}

func TestRemoveSurroundingEmptyLines(t *testing.T) {
	str1 := RemoveSurroundingEmptyLines(`
		Foo
		Bar


	`)

	if str1 != "\t\tFoo\n\t\tBar" {
		t.Error("RemoveSurroundingEmptyLines() should remove all sourrounding empty lines")
	}
}
