package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionListOnAliasedModule(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module M {
    export module N {
        export function foo() { }
        function bar() { }
    }
}
import p = M.N;
p./**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{"foo"},
		},
	})
}
