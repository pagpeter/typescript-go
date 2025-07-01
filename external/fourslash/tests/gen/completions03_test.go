package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletions03(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = ` interface Foo {
    one: any;
    two: any;
    three: any;
 }

 let x: Foo = {
     get one() { return "" },
     set two(t) {},
     /**/
 }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{"three"},
		},
	})
}
