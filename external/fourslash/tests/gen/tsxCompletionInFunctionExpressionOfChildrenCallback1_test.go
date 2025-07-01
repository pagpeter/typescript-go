package fourslash_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/fourslash"
	"github.com/pagpeter/typescript-go/external/testutil"
)

func TestTsxCompletionInFunctionExpressionOfChildrenCallback1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@module: commonjs
//@jsx: preserve
// @Filename: 1.tsx
 declare module JSX {
     interface Element { }
     interface IntrinsicElements {
     }
     interface ElementAttributesProperty { props; }
     interface ElementChildrenAttribute { children; }
 }
 interface IUser {
     Name: string;
 }
 interface IFetchUserProps {
     children: (user: IUser) => any;
 }
 function FetchUser(props: IFetchUserProps) { return undefined; }
 function UserName() {
     return (
         <FetchUser>
             { user => (
                 <h1>{ user./**/ }</h1>
             )}
         </FetchUser>
     );
 }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{"Name"},
		},
	})
}
