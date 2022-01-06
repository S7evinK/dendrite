// Code generated by "stringer -type=PresenceStatus -output=presence_string.go -linecomment"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown - -1]
	_ = x[Online-0]
	_ = x[Offline-1]
	_ = x[Unavailable-2]
}

const _PresenceStatus_name = "unknownonlineofflineunavailable"

var _PresenceStatus_index = [...]uint8{0, 7, 13, 20, 31}

func (i PresenceStatus) String() string {
	i -= -1
	if i < 0 || i >= PresenceStatus(len(_PresenceStatus_index)-1) {
		return "PresenceStatus(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _PresenceStatus_name[_PresenceStatus_index[i]:_PresenceStatus_index[i+1]]
}