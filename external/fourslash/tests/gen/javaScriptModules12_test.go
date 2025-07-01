package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestJavaScriptModules12(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: mod1.js
var x = require('fs');
/*1*/
// @Filename: mod2.js
var y;
if(true) {
    y = require('fs');
}
/*2*/
// @Filename: glob1.js
var a = require;
/*3*/
// @Filename: glob2.js
var b = '';
/*4*/
// @Filename: consumer.js
/*5*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"x", &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "a"}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "b"}},
			Excludes: []string{"y"},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"y", &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "a"}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "b"}},
			Excludes: []string{"x"},
		},
	})
	f.VerifyCompletions(t, "3", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"a", &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "b"}},
			Excludes: []string{"x", "y"},
		},
	})
	f.VerifyCompletions(t, "4", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "a"}, "b"},
			Excludes: []string{"x", "y"},
		},
	})
	f.VerifyCompletions(t, []string{"5"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "a"}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "b"}},
			Excludes: []string{"x", "y"},
		},
	})
}
