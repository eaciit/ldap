// generated by stringer -type=Scope; DO NOT EDIT

package ldap

import "fmt"

const _Scope_name = "ScopeBaseObjectScopeSingleLevelScopeWholeSubtree"

var _Scope_index = [...]uint8{0, 15, 31, 48}

func (i Scope) String() string {
	if i >= Scope(len(_Scope_index)-1) {
		return fmt.Sprintf("Scope(%d)", i)
	}
	return _Scope_name[_Scope_index[i]:_Scope_index[i+1]]
}