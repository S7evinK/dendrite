package helper

import (
	"testing"

	"github.com/matrix-org/dendrite/roomserver/proto"
	"github.com/matrix-org/gomatrixserverlib"
)

func TestToMatrixRoomVersion(t *testing.T) {
	tests := []struct {
		name string
		args proto.RoomVersion
		want gomatrixserverlib.RoomVersion
	}{
		{
			name: "Version 1",
			args: proto.RoomVersion_V1,
			want: gomatrixserverlib.RoomVersionV1,
		},
		{
			name: "Version 6",
			args: proto.RoomVersion_V6,
			want: gomatrixserverlib.RoomVersionV6,
		},
		{
			name: "Unknown",
			args: proto.RoomVersion_UNKNOWN,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMatrixRoomVersion(tt.args); got != tt.want {
				t.Errorf("ToMatrixRoomVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToProtoRoomVersion(t *testing.T) {
	tests := []struct {
		name    string
		version gomatrixserverlib.RoomVersion
		want    proto.RoomVersion
	}{
		{
			name:    "Version 1",
			version: gomatrixserverlib.RoomVersionV1,
			want:    proto.RoomVersion_V1,
		},
		{
			name:    "Version 6",
			version: gomatrixserverlib.RoomVersionV6,
			want:    proto.RoomVersion_V6,
		},
		{
			name:    "Unknown",
			version: "does not compute",
			want:    proto.RoomVersion_UNKNOWN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToProtoRoomVersion(tt.version); got != tt.want {
				t.Errorf("ToProtoRoomVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
