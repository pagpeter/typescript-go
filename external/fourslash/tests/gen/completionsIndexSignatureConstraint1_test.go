package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionsIndexSignatureConstraint1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true

repro #9900

interface Test {
  a?: number;
  b?: string;
}

interface TestIndex {
  [key: string]: Test;
}

declare function testFunc<T extends TestIndex>(t: T): void;

testFunc({
  test: {
    /**/
  },
});`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "a?", InsertText: ptrTo("a"), FilterText: ptrTo("a")}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "b?", InsertText: ptrTo("b"), FilterText: ptrTo("b")}},
		},
	})
}
