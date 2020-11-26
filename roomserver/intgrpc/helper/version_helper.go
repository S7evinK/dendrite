package helper

import (
	"strconv"

	"github.com/matrix-org/dendrite/roomserver/proto"
	"github.com/matrix-org/gomatrixserverlib"
)

func ToProtoRoomVersion(version gomatrixserverlib.RoomVersion) proto.RoomVersion {
	i, err := strconv.Atoi(string(version))
	if err != nil {
		i = 0
	}
	return proto.RoomVersion(int32(i))
}

func ToMatrixRoomVersion(v proto.RoomVersion) gomatrixserverlib.RoomVersion {
	return gomatrixserverlib.RoomVersion(strconv.Itoa(int(v.Number())))
}
