package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionsWithOptionalPropertiesGenericPartial(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
interface Foo {
    a_a: boolean;
    a_b: boolean;
    a_c: boolean;
    b_a: boolean;
}
function partialFoo<T extends Partial<Foo>>(t: T) {return t}
partialFoo({ /*1*/ });`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "a_a?", InsertText: ptrTo("a_a"), FilterText: ptrTo("a_a")}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "a_b?", InsertText: ptrTo("a_b"), FilterText: ptrTo("a_b")}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "a_c?", InsertText: ptrTo("a_c"), FilterText: ptrTo("a_c")}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "b_a?", InsertText: ptrTo("b_a"), FilterText: ptrTo("b_a")}},
		},
	})
}
