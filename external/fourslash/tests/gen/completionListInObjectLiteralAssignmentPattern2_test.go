package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionListInObjectLiteralAssignmentPattern2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `let x = { a: 1, b: 2 };
let y = ({ a, /**/ } = x, 1);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{"b"},
		},
	})
}
