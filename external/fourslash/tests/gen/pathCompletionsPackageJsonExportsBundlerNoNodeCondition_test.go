package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestPathCompletionsPackageJsonExportsBundlerNoNodeCondition(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @moduleResolution: bundler
// @Filename: /node_modules/foo/package.json
{
  "name": "foo",
  "exports": {
    "./only-for-node": {
      "node": "./something.js"
    },
    "./for-everywhere": "./other.js"
  }
}
// @Filename: /node_modules/foo/something.d.ts
export const index = 0;
// @Filename: /node_modules/foo/other.d.ts
export const index = 0;
// @Filename: /index.ts
import { } from "foo//**/";`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{&lsproto.CompletionItem{Kind: ptrTo(lsproto.CompletionItemKindFile), Label: "for-everywhere"}},
		},
	})
}
