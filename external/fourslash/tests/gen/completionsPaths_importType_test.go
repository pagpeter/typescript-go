package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestCompletionsPaths_importType(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @moduleResolution: node
// @Filename: /ns.ts
file content not read
// @Filename: /node_modules/package/index.ts
file content not read
// @Filename: /usage.ts
type A = typeof import("p/*1*/");
type B = import(".//*2*/");
// @Filename: /user.js
/** @type {import("/*3*/")} */`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"1", "3"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFolder), Label: "package"}},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFile), Label: "lib"}, &lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFile), Label: "lib.decorators"}, &lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFile), Label: "lib.decorators.legacy"}, &lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFile), Label: "ns"}, &lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFile), Label: "user"}, &lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFolder), Label: "node_modules"}},
		},
	})
}
