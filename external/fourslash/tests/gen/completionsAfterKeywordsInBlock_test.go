package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionsAfterKeywordsInBlock(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C1 {
    method(map: Map<string, string>, key: string, defaultValue: string) {
        try {
            return map.get(key)!;
        }
        catch {
            return default/*1*/
        }
    }
}
class C2 {
    method(map: Map<string, string>, key: string, defaultValue: string) {
        if (map.has(key)) {
            return map.get(key)!;
        }
        else {
            return default/*2*/
        }
    }
}
class C3 {
    method(map: Map<string, string>, key: string, returnValue: string) {
        try {
            return map.get(key)!;
        }
        catch {
            return return/*3*/
        }
    }
}
class C4 {
    method(map: Map<string, string>, key: string, returnValue: string) {
        if (map.has(key)) {
            return map.get(key)!;
        }
        else {
            return return/*4*/
        }
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"1", "2"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextLocationPriority)), Label: "defaultValue"}},
		},
	})
	f.VerifyCompletions(t, []string{"3", "4"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextLocationPriority)), Label: "returnValue"}},
		},
	})
}
