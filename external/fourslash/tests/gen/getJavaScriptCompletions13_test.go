package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestGetJavaScriptCompletions13(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: file1.js
var file1Identifier = 1;
interface Foo { FooProp: number };
// @Filename: file2.js
var file2Identifier1 = 2;
var file2Identifier2 = 2;
/*1*/
file2Identifier2./*2*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"file2Identifier1", "file2Identifier2", &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "file1Identifier"}},
			Excludes: []string{"FooProp"},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextJavascriptIdentifiers)), Label: "file2Identifier1"}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextJavascriptIdentifiers)), Label: "file2Identifier2"}},
			Excludes: []string{"file1Identifier", "FooProp"},
		},
	})
}
