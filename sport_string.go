// Code generated by "stringer -type=Sport --linecomment"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Football-1]
	_ = x[Basketball-2]
	_ = x[TableTennis-3]
}

const _Sport_name = "足球篮球乒乓球"

var _Sport_index = [...]uint8{0, 6, 12, 21}

func (i Sport) String() string {
	i -= 1
	if i < 0 || i >= Sport(len(_Sport_index)-1) {
		return "Sport(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Sport_name[_Sport_index[i]:_Sport_index[i+1]]
}