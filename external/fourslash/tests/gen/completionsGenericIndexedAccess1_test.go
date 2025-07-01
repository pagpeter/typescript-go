package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionsGenericIndexedAccess1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Sample {
  addBook: { name: string, year: number }
}

export declare function testIt<T>(method: T[keyof T]): any
testIt<Sample>({ /**/ });`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{Label: "name"}, &lsproto.CompletionItem{Label: "year"}},
		},
	})
}
