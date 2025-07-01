package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/ls"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionsWithOptionalPropertiesGenericConstructor(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
interface Options {
    someFunction?: () => string
    anotherFunction?: () => string
}

export class Clazz<T extends Options> {
    constructor(public a: T) {}
}

new Clazz({ /*1*/ })`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "someFunction?", InsertText: ptrTo("someFunction"), FilterText: ptrTo("someFunction")}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "anotherFunction?", InsertText: ptrTo("anotherFunction"), FilterText: ptrTo("anotherFunction")}},
		},
	})
}
