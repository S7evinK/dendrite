package api

import (
	"context"

	fsAPI "github.com/matrix-org/dendrite/federationsender/api"
	"github.com/matrix-org/dendrite/roomserver/proto"
)

// RoomserverInputAPI is used to write events to the room server.
type RoomserverInternalAPI interface {
	// needed to avoid chicken and egg scenario when setting up the
	// interdependencies between the roomserver and other input APIs
	SetFederationSenderAPI(fsAPI fsAPI.FederationSenderInternalAPI)

	InputRoomEvents(
		ctx context.Context,
		request *InputRoomEventsRequest,
		response *InputRoomEventsResponse,
	)

	PerformInvite(
		ctx context.Context,
		req *PerformInviteRequest,
		res *PerformInviteResponse,
	) error

	PerformJoin(
		ctx context.Context,
		req *PerformJoinRequest,
		res *PerformJoinResponse,
	)

	PerformLeave(
		ctx context.Context,
		req *PerformLeaveRequest,
		res *PerformLeaveResponse,
	) error

	PerformPeek(
		ctx context.Context,
		req *PerformPeekRequest,
		res *PerformPeekResponse,
	)

	PerformPublish(
		ctx context.Context,
		req *PerformPublishRequest,
		res *PerformPublishResponse,
	)

	QueryPublishedRoomsGRPC(
		ctx context.Context,
		req *proto.PublishedRoomsRequest,
	) (*proto.PublishedRoomsResponse, error)

	// Query the latest events and state for a room from the room server.
	QueryLatestEventsAndState(
		ctx context.Context,
		request *QueryLatestEventsAndStateRequest,
		response *QueryLatestEventsAndStateResponse,
	) error

	// Query the state after a list of events in a room from the room server.
	QueryStateAfterEvents(
		ctx context.Context,
		request *QueryStateAfterEventsRequest,
		response *QueryStateAfterEventsResponse,
	) error

	// Query whether the roomserver is missing any auth or prev events.
	QueryMissingAuthPrevEvents(
		ctx context.Context,
		request *QueryMissingAuthPrevEventsRequest,
		response *QueryMissingAuthPrevEventsResponse,
	) error

	// Query a list of events by event ID.
	QueryEventsByID(
		ctx context.Context,
		request *QueryEventsByIDRequest,
		response *QueryEventsByIDResponse,
	) error

	// Query the membership event for an user for a room.
	QueryMembershipForUserGRPC(
		ctx context.Context,
		req *proto.MembershipForUserRequest,
	) (*proto.MembershipForUserResponse, error)

	// Query a list of membership events for a room
	QueryMembershipsForRoom(
		ctx context.Context,
		request *QueryMembershipsForRoomRequest,
		response *QueryMembershipsForRoomResponse,
	) error

	// Query if we think we're still in a room.
	QueryServerJoinedToRoomGRPC(
		ctx context.Context,
		req *proto.ServerJoinedToRoomRequest,
	) (*proto.ServerJoinedToRoomResponse, error)

	// Query whether a server is allowed to see an event
	QueryServerAllowedToSeeEventGRPC(
		ctx context.Context,
		req *proto.ServerAllowedToSeeEventRequest,
	) (*proto.ServerAllowedToSeeEventResponse, error)

	// Query missing events for a room from roomserver
	QueryMissingEvents(
		ctx context.Context,
		request *QueryMissingEventsRequest,
		response *QueryMissingEventsResponse,
	) error

	// Query to get state and auth chain for a (potentially hypothetical) event.
	// Takes lists of PrevEventIDs and AuthEventsIDs and uses them to calculate
	// the state and auth chain to return.
	QueryStateAndAuthChain(
		ctx context.Context,
		request *QueryStateAndAuthChainRequest,
		response *QueryStateAndAuthChainResponse,
	) error

	// QueryCurrentState retrieves the requested state events. If state events are not found, they will be missing from
	// the response.
	QueryCurrentState(ctx context.Context, req *QueryCurrentStateRequest, res *QueryCurrentStateResponse) error
	// QueryRoomsForUser retrieves a list of room IDs matching the given query.
	QueryRoomsForUserGRPC(ctx context.Context, req *proto.RoomsForUserRequest) (*proto.RoomsForUserResponse, error)
	// QueryBulkStateContent does a bulk query for state event content in the given rooms.
	QueryBulkStateContent(ctx context.Context, req *QueryBulkStateContentRequest, res *QueryBulkStateContentResponse) error
	// QuerySharedUsers returns a list of users who share at least 1 room in common with the given user.
	QuerySharedUsersGRPC(ctx context.Context, req *proto.SharedUsersRequest) (*proto.SharedUsersResponse, error)
	// QueryKnownUsers returns a list of users that we know about from our joined rooms.
	QueryKnownUsers(ctx context.Context, req *QueryKnownUsersRequest, res *QueryKnownUsersResponse) error
	// QueryServerBannedFromRoom returns whether a server is banned from a room by server ACLs.
	QueryServerBannedFromRoomGRPC(ctx context.Context, req *proto.ServerBannedFromRoomRequest) (*proto.ServerBannedFromRoomResponse, error)

	// Query a given amount (or less) of events prior to a given set of events.
	PerformBackfill(
		ctx context.Context,
		request *PerformBackfillRequest,
		response *PerformBackfillResponse,
	) error

	// PerformForget forgets a rooms history for a specific user
	PerformForget(ctx context.Context, req *PerformForgetRequest, resp *PerformForgetResponse) error

	// Asks for the default room version as preferred by the server.
	QueryRoomVersionCapabilitiesGRPC(
		ctx context.Context,
		request *proto.RoomVersionCapabilitiesRequest,
	) (*proto.RoomVersionCapabilitiesResponse, error)

	// Asks for the room version for a given room.
	QueryRoomVersionForRoomGRPC(
		ctx context.Context,
		request *proto.RoomVersionForRoomRequest,
	) (*proto.RoomVersionForRoomResponse, error)

	// Set a room alias
	SetRoomAlias(
		ctx context.Context,
		req *SetRoomAliasRequest,
		response *SetRoomAliasResponse,
	) error

	// Get the room ID for an alias
	GetRoomIDForAlias(
		ctx context.Context,
		req *GetRoomIDForAliasRequest,
		response *GetRoomIDForAliasResponse,
	) error

	// Get all known aliases for a room ID
	GetAliasesForRoomID(
		ctx context.Context,
		req *GetAliasesForRoomIDRequest,
		response *GetAliasesForRoomIDResponse,
	) error

	// Get the user ID of the creator of an alias
	GetCreatorIDForAlias(
		ctx context.Context,
		req *GetCreatorIDForAliasRequest,
		response *GetCreatorIDForAliasResponse,
	) error

	// Remove a room alias
	RemoveRoomAlias(
		ctx context.Context,
		req *RemoveRoomAliasRequest,
		response *RemoveRoomAliasResponse,
	) error
}
