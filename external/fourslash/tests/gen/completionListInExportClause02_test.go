package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionListInExportClause02(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare module "M1" {
    export var V;
}
var W;
declare module "M2" {
    export { /**/ } from "M1"
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{"V", &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextGlobalsOrKeywords)), Label: "type"}},
		},
	})
}
