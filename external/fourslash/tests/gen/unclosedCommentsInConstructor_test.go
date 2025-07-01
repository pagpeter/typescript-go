package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestUnclosedCommentsInConstructor(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Foo {
    constructor(/* /**/) { }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", nil)
}
