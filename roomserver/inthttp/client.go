package inthttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	fsInputAPI "github.com/matrix-org/dendrite/federationsender/api"
	"github.com/matrix-org/dendrite/internal/caching"
	"github.com/matrix-org/dendrite/internal/httputil"
	"github.com/matrix-org/dendrite/roomserver/api"
	"github.com/matrix-org/dendrite/roomserver/intgrpc"
	"github.com/matrix-org/dendrite/roomserver/proto"
	"github.com/opentracing/opentracing-go"
)

const (
	// Alias operations
	RoomserverSetRoomAliasPath         = "/roomserver/setRoomAlias"
	RoomserverGetRoomIDForAliasPath    = "/roomserver/GetRoomIDForAlias"
	RoomserverGetAliasesForRoomIDPath  = "/roomserver/GetAliasesForRoomID"
	RoomserverGetCreatorIDForAliasPath = "/roomserver/GetCreatorIDForAlias"
	RoomserverRemoveRoomAliasPath      = "/roomserver/removeRoomAlias"

	// Input operations
	RoomserverInputRoomEventsPath = "/roomserver/inputRoomEvents"

	// Perform operations
	RoomserverPerformInvitePath   = "/roomserver/performInvite"
	RoomserverPerformPeekPath     = "/roomserver/performPeek"
	RoomserverPerformJoinPath     = "/roomserver/performJoin"
	RoomserverPerformLeavePath    = "/roomserver/performLeave"
	RoomserverPerformBackfillPath = "/roomserver/performBackfill"
	RoomserverPerformPublishPath  = "/roomserver/performPublish"
	RoomserverPerformForgetPath   = "/roomserver/performForget"

	// Query operations
	RoomserverQueryLatestEventsAndStatePath  = "/roomserver/queryLatestEventsAndState"
	RoomserverQueryStateAfterEventsPath      = "/roomserver/queryStateAfterEvents"
	RoomserverQueryMissingAuthPrevEventsPath = "/roomserver/queryMissingAuthPrevEvents"
	RoomserverQueryEventsByIDPath            = "/roomserver/queryEventsByID"
	RoomserverQueryMembershipForUserPath     = "/roomserver/queryMembershipForUser"
	RoomserverQueryMembershipsForRoomPath    = "/roomserver/queryMembershipsForRoom"
	RoomserverQueryMissingEventsPath         = "/roomserver/queryMissingEvents"
	RoomserverQueryStateAndAuthChainPath     = "/roomserver/queryStateAndAuthChain"
	RoomserverQueryCurrentStatePath          = "/roomserver/queryCurrentState"
	RoomserverQueryBulkStateContentPath      = "/roomserver/queryBulkStateContent"
	RoomserverQueryKnownUsersPath            = "/roomserver/queryKnownUsers"
)

type httpRoomserverInternalAPI struct {
	roomserverURL string
	httpClient    *http.Client
	cache         caching.RoomVersionCache
	grpcClient    *intgrpc.RoomServiceClient
}

// NewRoomserverClient creates a RoomserverInputAPI implemented by talking to a HTTP POST API.
// If httpClient is nil an error is returned
func NewRoomserverClient(
	roomserverURL string,
	httpClient *http.Client,
	cache caching.RoomVersionCache,
	grpcClient *intgrpc.RoomServiceClient,
) (api.RoomserverInternalAPI, error) {
	if httpClient == nil {
		return nil, errors.New("NewRoomserverInternalAPIHTTP: httpClient is <nil>")
	}
	return &httpRoomserverInternalAPI{
		roomserverURL: roomserverURL,
		httpClient:    httpClient,
		cache:         cache,
		grpcClient:    grpcClient,
	}, nil
}

// SetFederationSenderInputAPI no-ops in HTTP client mode as there is no chicken/egg scenario
func (h *httpRoomserverInternalAPI) SetFederationSenderAPI(fsAPI fsInputAPI.FederationSenderInternalAPI) {
}

// SetRoomAlias implements RoomserverAliasAPI
func (h *httpRoomserverInternalAPI) SetRoomAlias(
	ctx context.Context,
	request *api.SetRoomAliasRequest,
	response *api.SetRoomAliasResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SetRoomAlias")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverSetRoomAliasPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// GetRoomIDForAlias implements RoomserverAliasAPI
func (h *httpRoomserverInternalAPI) GetRoomIDForAlias(
	ctx context.Context,
	request *api.GetRoomIDForAliasRequest,
	response *api.GetRoomIDForAliasResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetRoomIDForAlias")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverGetRoomIDForAliasPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// GetAliasesForRoomID implements RoomserverAliasAPI
func (h *httpRoomserverInternalAPI) GetAliasesForRoomID(
	ctx context.Context,
	request *api.GetAliasesForRoomIDRequest,
	response *api.GetAliasesForRoomIDResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetAliasesForRoomID")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverGetAliasesForRoomIDPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// GetCreatorIDForAlias implements RoomserverAliasAPI
func (h *httpRoomserverInternalAPI) GetCreatorIDForAlias(
	ctx context.Context,
	request *api.GetCreatorIDForAliasRequest,
	response *api.GetCreatorIDForAliasResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetCreatorIDForAlias")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverGetCreatorIDForAliasPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// RemoveRoomAlias implements RoomserverAliasAPI
func (h *httpRoomserverInternalAPI) RemoveRoomAlias(
	ctx context.Context,
	request *api.RemoveRoomAliasRequest,
	response *api.RemoveRoomAliasResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RemoveRoomAlias")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverRemoveRoomAliasPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// InputRoomEvents implements RoomserverInputAPI
func (h *httpRoomserverInternalAPI) InputRoomEvents(
	ctx context.Context,
	request *api.InputRoomEventsRequest,
	response *api.InputRoomEventsResponse,
) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InputRoomEvents")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverInputRoomEventsPath
	err := httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
	if err != nil {
		response.ErrMsg = err.Error()
	}
}

func (h *httpRoomserverInternalAPI) PerformInvite(
	ctx context.Context,
	request *api.PerformInviteRequest,
	response *api.PerformInviteResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformInvite")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformInvitePath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

func (h *httpRoomserverInternalAPI) PerformJoin(
	ctx context.Context,
	request *api.PerformJoinRequest,
	response *api.PerformJoinResponse,
) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformJoin")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformJoinPath
	err := httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
	if err != nil {
		response.Error = &api.PerformError{
			Msg: fmt.Sprintf("failed to communicate with roomserver: %s", err),
		}
	}
}

func (h *httpRoomserverInternalAPI) PerformPeek(
	ctx context.Context,
	request *api.PerformPeekRequest,
	response *api.PerformPeekResponse,
) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformPeek")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformPeekPath
	err := httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
	if err != nil {
		response.Error = &api.PerformError{
			Msg: fmt.Sprintf("failed to communicate with roomserver: %s", err),
		}
	}
}

func (h *httpRoomserverInternalAPI) PerformLeave(
	ctx context.Context,
	request *api.PerformLeaveRequest,
	response *api.PerformLeaveResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformLeave")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformLeavePath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

func (h *httpRoomserverInternalAPI) PerformPublish(
	ctx context.Context,
	req *api.PerformPublishRequest,
	res *api.PerformPublishResponse,
) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformPublish")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformPublishPath
	err := httputil.PostJSON(ctx, span, h.httpClient, apiURL, req, res)
	if err != nil {
		res.Error = &api.PerformError{
			Msg: fmt.Sprintf("failed to communicate with roomserver: %s", err),
		}
	}
}

// QueryLatestEventsAndState implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryLatestEventsAndState(
	ctx context.Context,
	request *api.QueryLatestEventsAndStateRequest,
	response *api.QueryLatestEventsAndStateResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryLatestEventsAndState")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryLatestEventsAndStatePath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryStateAfterEvents implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryStateAfterEvents(
	ctx context.Context,
	request *api.QueryStateAfterEventsRequest,
	response *api.QueryStateAfterEventsResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryStateAfterEvents")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryStateAfterEventsPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryStateAfterEvents implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryMissingAuthPrevEvents(
	ctx context.Context,
	request *api.QueryMissingAuthPrevEventsRequest,
	response *api.QueryMissingAuthPrevEventsResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryMissingAuthPrevEvents")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryMissingAuthPrevEventsPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryEventsByID implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryEventsByID(
	ctx context.Context,
	request *api.QueryEventsByIDRequest,
	response *api.QueryEventsByIDResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryEventsByID")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryEventsByIDPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryMembershipForUser implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryMembershipForUser(
	ctx context.Context,
	request *api.QueryMembershipForUserRequest,
	response *api.QueryMembershipForUserResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryMembershipForUser")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryMembershipForUserPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryMembershipsForRoom implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryMembershipsForRoom(
	ctx context.Context,
	request *api.QueryMembershipsForRoomRequest,
	response *api.QueryMembershipsForRoomResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryMembershipsForRoom")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryMembershipsForRoomPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryMissingEvents implements RoomServerQueryAPI
func (h *httpRoomserverInternalAPI) QueryMissingEvents(
	ctx context.Context,
	request *api.QueryMissingEventsRequest,
	response *api.QueryMissingEventsResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryMissingEvents")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryMissingEventsPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// QueryStateAndAuthChain implements RoomserverQueryAPI
func (h *httpRoomserverInternalAPI) QueryStateAndAuthChain(
	ctx context.Context,
	request *api.QueryStateAndAuthChainRequest,
	response *api.QueryStateAndAuthChainResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryStateAndAuthChain")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryStateAndAuthChainPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

// PerformBackfill implements RoomServerQueryAPI
func (h *httpRoomserverInternalAPI) PerformBackfill(
	ctx context.Context,
	request *api.PerformBackfillRequest,
	response *api.PerformBackfillResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformBackfill")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformBackfillPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

func (h *httpRoomserverInternalAPI) QueryCurrentState(
	ctx context.Context,
	request *api.QueryCurrentStateRequest,
	response *api.QueryCurrentStateResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryCurrentState")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryCurrentStatePath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

func (h *httpRoomserverInternalAPI) QueryBulkStateContent(
	ctx context.Context,
	request *api.QueryBulkStateContentRequest,
	response *api.QueryBulkStateContentResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryBulkStateContent")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryBulkStateContentPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, request, response)
}

func (h *httpRoomserverInternalAPI) QueryKnownUsers(
	ctx context.Context, req *api.QueryKnownUsersRequest, res *api.QueryKnownUsersResponse,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "QueryKnownUsers")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverQueryKnownUsersPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, req, res)
}

func (h *httpRoomserverInternalAPI) QueryServerBannedFromRoomGRPC(ctx context.Context, req *proto.ServerBannedFromRoomRequest) (*proto.ServerBannedFromRoomResponse, error) {
	return h.grpcClient.QueryServerBannedFromRoom(ctx, req)
}

// QuerySharedUsersGRPC returns a list of users who share at least 1 room in common with the given user.
func (h *httpRoomserverInternalAPI) QuerySharedUsersGRPC(ctx context.Context, req *proto.SharedUsersRequest) (*proto.SharedUsersResponse, error) {
	return h.grpcClient.QuerySharedUsers(ctx, req)
}

func (h *httpRoomserverInternalAPI) QueryRoomsForUserGRPC(ctx context.Context, req *proto.RoomsForUserRequest) (*proto.RoomsForUserResponse, error) {
	return h.grpcClient.QueryRoomsForUser(ctx, req)
}

func (h *httpRoomserverInternalAPI) QueryPublishedRoomsGRPC(ctx context.Context, req *proto.PublishedRoomsRequest) (*proto.PublishedRoomsResponse, error) {
	return h.grpcClient.QueryPublishedRooms(ctx, req)
}

func (h *httpRoomserverInternalAPI) QueryRoomVersionForRoomGRPC(ctx context.Context, req *proto.RoomVersionForRoomRequest) (*proto.RoomVersionForRoomResponse, error) {
	return h.grpcClient.QueryRoomVersionForRoom(ctx, req)
}

func (h *httpRoomserverInternalAPI) QueryRoomVersionCapabilitiesGRPC(ctx context.Context, req *proto.RoomVersionCapabilitiesRequest) (*proto.RoomVersionCapabilitiesResponse, error) {
	return h.grpcClient.QueryRoomVersionCapabilities(ctx, req)
}

func (h *httpRoomserverInternalAPI) QueryServerAllowedToSeeEventGRPC(ctx context.Context, req *proto.ServerAllowedToSeeEventRequest) (*proto.ServerAllowedToSeeEventResponse, error) {
	return h.grpcClient.QueryServerAllowedToSeeEvent(ctx, req)
}

func (h *httpRoomserverInternalAPI) QueryServerJoinedToRoomGRPC(ctx context.Context, req *proto.ServerJoinedToRoomRequest) (*proto.ServerJoinedToRoomResponse, error) {
	return h.grpcClient.QueryServerJoinedToRoomGRPC(ctx, req)
}

func (h *httpRoomserverInternalAPI) PerformForget(ctx context.Context, req *api.PerformForgetRequest, res *api.PerformForgetResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PerformForget")
	defer span.Finish()

	apiURL := h.roomserverURL + RoomserverPerformForgetPath
	return httputil.PostJSON(ctx, span, h.httpClient, apiURL, req, res)

}
