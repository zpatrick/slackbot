package slackbot

import (
	"context"

	"github.com/nlopes/slack"
)

// SlackClient is used to mock a *slack.Client
type SlackClient interface {
	AddPin(channel string, item slack.ItemRef) error
	AddPinContext(ctx context.Context, channel string, item slack.ItemRef) error
	AddReaction(name string, item slack.ItemRef) error
	AddReactionContext(ctx context.Context, name string, item slack.ItemRef) error
	AddStar(channel string, item slack.ItemRef) error
	AddStarContext(ctx context.Context, channel string, item slack.ItemRef) error
	ArchiveChannel(channelID string) error
	ArchiveChannelContext(ctx context.Context, channelID string) (err error)
	ArchiveConversation(channelID string) error
	ArchiveConversationContext(ctx context.Context, channelID string) error
	ArchiveGroup(group string) error
	ArchiveGroupContext(ctx context.Context, group string) error
	AuthTest() (response *slack.AuthTestResponse, error error)
	AuthTestContext(ctx context.Context) (response *slack.AuthTestResponse, error error)
	CloseConversation(channelID string) (noOp bool, alreadyClosed bool, err error)
	CloseConversationContext(ctx context.Context, channelID string) (noOp bool, alreadyClosed bool, err error)
	CloseGroup(group string) (bool, bool, error)
	CloseGroupContext(ctx context.Context, group string) (bool, bool, error)
	CloseIMChannel(channel string) (bool, bool, error)
	CloseIMChannelContext(ctx context.Context, channel string) (bool, bool, error)
	ConnectRTM() (info *slack.Info, websocketURL string, err error)
	ConnectRTMContext(ctx context.Context) (info *slack.Info, websocketURL string, err error)
	CreateChannel(channelName string) (*slack.Channel, error)
	CreateChannelContext(ctx context.Context, channelName string) (*slack.Channel, error)
	CreateChildGroup(group string) (*slack.Group, error)
	CreateChildGroupContext(ctx context.Context, group string) (*slack.Group, error)
	CreateConversation(channelName string, isPrivate bool) (*slack.Channel, error)
	CreateConversationContext(ctx context.Context, channelName string, isPrivate bool) (*slack.Channel, error)
	CreateGroup(group string) (*slack.Group, error)
	CreateGroupContext(ctx context.Context, group string) (*slack.Group, error)
	CreateUserGroup(userGroup slack.UserGroup) (slack.UserGroup, error)
	CreateUserGroupContext(ctx context.Context, userGroup slack.UserGroup) (slack.UserGroup, error)
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	DeleteFile(fileID string) error
	DeleteFileComment(commentID, fileID string) error
	DeleteFileCommentContext(ctx context.Context, fileID, commentID string) (err error)
	DeleteFileContext(ctx context.Context, fileID string) (err error)
	DeleteMessage(channel, messageTimestamp string) (string, string, error)
	DeleteMessageContext(ctx context.Context, channel, messageTimestamp string) (string, string, error)
	DeleteUserPhoto() error
	DeleteUserPhotoContext(ctx context.Context) error
	DisableUser(teamName string, uid string) error
	DisableUserContext(ctx context.Context, teamName string, uid string) error
	DisableUserGroup(userGroup string) (slack.UserGroup, error)
	DisableUserGroupContext(ctx context.Context, userGroup string) (slack.UserGroup, error)
	EnableUserGroup(userGroup string) (slack.UserGroup, error)
	EnableUserGroupContext(ctx context.Context, userGroup string) (slack.UserGroup, error)
	EndDND() error
	EndDNDContext(ctx context.Context) error
	EndSnooze() (*slack.DNDStatus, error)
	EndSnoozeContext(ctx context.Context) (*slack.DNDStatus, error)
	GetAccessLogs(params slack.AccessLogParameters) ([]slack.Login, *slack.Paging, error)
	GetAccessLogsContext(ctx context.Context, params slack.AccessLogParameters) ([]slack.Login, *slack.Paging, error)
	GetBillableInfo(user string) (map[string]slack.BillingActive, error)
	GetBillableInfoContext(ctx context.Context, user string) (map[string]slack.BillingActive, error)
	GetBillableInfoForTeam() (map[string]slack.BillingActive, error)
	GetBillableInfoForTeamContext(ctx context.Context) (map[string]slack.BillingActive, error)
	GetBotInfo(bot string) (*slack.Bot, error)
	GetBotInfoContext(ctx context.Context, bot string) (*slack.Bot, error)
	GetChannelHistory(channelID string, params slack.HistoryParameters) (*slack.History, error)
	GetChannelHistoryContext(ctx context.Context, channelID string, params slack.HistoryParameters) (*slack.History, error)
	GetChannelInfo(channelID string) (*slack.Channel, error)
	GetChannelInfoContext(ctx context.Context, channelID string) (*slack.Channel, error)
	GetChannelReplies(channelID, thread_ts string) ([]slack.Message, error)
	GetChannelRepliesContext(ctx context.Context, channelID, thread_ts string) ([]slack.Message, error)
	GetChannels(excludeArchived bool) ([]slack.Channel, error)
	GetChannelsContext(ctx context.Context, excludeArchived bool) ([]slack.Channel, error)
	GetConversationHistory(params *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error)
	GetConversationHistoryContext(ctx context.Context, params *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error)
	GetConversationInfo(channelID string, includeLocale bool) (*slack.Channel, error)
	GetConversationInfoContext(ctx context.Context, channelID string, includeLocale bool) (*slack.Channel, error)
	GetConversationReplies(params *slack.GetConversationRepliesParameters) (msgs []slack.Message, hasMore bool, nextCursor string, err error)
	GetConversationRepliesContext(ctx context.Context, params *slack.GetConversationRepliesParameters) (msgs []slack.Message, hasMore bool, nextCursor string, err error)
	GetConversations(params *slack.GetConversationsParameters) (channels []slack.Channel, nextCursor string, err error)
	GetConversationsContext(ctx context.Context, params *slack.GetConversationsParameters) (channels []slack.Channel, nextCursor string, err error)
	GetDNDInfo(user *string) (*slack.DNDStatus, error)
	GetDNDInfoContext(ctx context.Context, user *string) (*slack.DNDStatus, error)
	GetDNDTeamInfo(users []string) (map[string]slack.DNDStatus, error)
	GetDNDTeamInfoContext(ctx context.Context, users []string) (map[string]slack.DNDStatus, error)
	GetEmoji() (map[string]string, error)
	GetEmojiContext(ctx context.Context) (map[string]string, error)
	GetFileInfo(fileID string, count, page int) (*slack.File, []slack.Comment, *slack.Paging, error)
	GetFileInfoContext(ctx context.Context, fileID string, count, page int) (*slack.File, []slack.Comment, *slack.Paging, error)
	GetFiles(params slack.GetFilesParameters) ([]slack.File, *slack.Paging, error)
	GetFilesContext(ctx context.Context, params slack.GetFilesParameters) ([]slack.File, *slack.Paging, error)
	GetGroupHistory(group string, params slack.HistoryParameters) (*slack.History, error)
	GetGroupHistoryContext(ctx context.Context, group string, params slack.HistoryParameters) (*slack.History, error)
	GetGroupInfo(group string) (*slack.Group, error)
	GetGroupInfoContext(ctx context.Context, group string) (*slack.Group, error)
	GetGroups(excludeArchived bool) ([]slack.Group, error)
	GetGroupsContext(ctx context.Context, excludeArchived bool) ([]slack.Group, error)
	GetIMChannels() ([]slack.IM, error)
	GetIMChannelsContext(ctx context.Context) ([]slack.IM, error)
	GetIMHistory(channel string, params slack.HistoryParameters) (*slack.History, error)
	GetIMHistoryContext(ctx context.Context, channel string, params slack.HistoryParameters) (*slack.History, error)
	GetReactions(item slack.ItemRef, params slack.GetReactionsParameters) ([]slack.ItemReaction, error)
	GetReactionsContext(ctx context.Context, item slack.ItemRef, params slack.GetReactionsParameters) ([]slack.ItemReaction, error)
	GetStarred(params slack.StarsParameters) ([]slack.StarredItem, *slack.Paging, error)
	GetStarredContext(ctx context.Context, params slack.StarsParameters) ([]slack.StarredItem, *slack.Paging, error)
	GetTeamInfo() (*slack.TeamInfo, error)
	GetTeamInfoContext(ctx context.Context) (*slack.TeamInfo, error)
	GetUserByEmail(email string) (*slack.User, error)
	GetUserByEmailContext(ctx context.Context, email string) (*slack.User, error)
	GetUserGroupMembers(userGroup string) ([]string, error)
	GetUserGroupMembersContext(ctx context.Context, userGroup string) ([]string, error)
	GetUserGroups() ([]slack.UserGroup, error)
	GetUserGroupsContext(ctx context.Context) ([]slack.UserGroup, error)
	GetUserIdentity() (*slack.UserIdentityResponse, error)
	GetUserIdentityContext(ctx context.Context) (*slack.UserIdentityResponse, error)
	GetUserInfo(user string) (*slack.User, error)
	GetUserInfoContext(ctx context.Context, user string) (*slack.User, error)
	GetUserPresence(user string) (*slack.UserPresence, error)
	GetUserPresenceContext(ctx context.Context, user string) (*slack.UserPresence, error)
	GetUsers() ([]slack.User, error)
	GetUsersContext(ctx context.Context) ([]slack.User, error)
	GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error)
	GetUsersInConversationContext(ctx context.Context, params *slack.GetUsersInConversationParameters) ([]string, string, error)
	InviteGuest(teamName, channel, firstName, lastName, emailAddress string) error
	InviteGuestContext(ctx context.Context, teamName, channel, firstName, lastName, emailAddress string) error
	InviteRestricted(teamName, channel, firstName, lastName, emailAddress string) error
	InviteRestrictedContext(ctx context.Context, teamName, channel, firstName, lastName, emailAddress string) error
	InviteToTeam(teamName, firstName, lastName, emailAddress string) error
	InviteToTeamContext(ctx context.Context, teamName, firstName, lastName, emailAddress string) error
	InviteUserToChannel(channelID, user string) (*slack.Channel, error)
	InviteUserToChannelContext(ctx context.Context, channelID, user string) (*slack.Channel, error)
	InviteUserToGroup(group, user string) (*slack.Group, bool, error)
	InviteUserToGroupContext(ctx context.Context, group, user string) (*slack.Group, bool, error)
	InviteUsersToConversation(channelID string, users ...string) (*slack.Channel, error)
	InviteUsersToConversationContext(ctx context.Context, channelID string, users ...string) (*slack.Channel, error)
	JoinChannel(channelName string) (*slack.Channel, error)
	JoinChannelContext(ctx context.Context, channelName string) (*slack.Channel, error)
	JoinConversation(channelID string) (*slack.Channel, string, []string, error)
	JoinConversationContext(ctx context.Context, channelID string) (*slack.Channel, string, []string, error)
	KickUserFromChannel(channelID, user string) error
	KickUserFromChannelContext(ctx context.Context, channelID, user string) (err error)
	KickUserFromConversation(channelID string, user string) error
	KickUserFromConversationContext(ctx context.Context, channelID string, user string) error
	KickUserFromGroup(group, user string) error
	KickUserFromGroupContext(ctx context.Context, group, user string) (err error)
	LeaveChannel(channelID string) (bool, error)
	LeaveChannelContext(ctx context.Context, channelID string) (bool, error)
	LeaveConversation(channelID string) (bool, error)
	LeaveConversationContext(ctx context.Context, channelID string) (bool, error)
	LeaveGroup(group string) error
	LeaveGroupContext(ctx context.Context, group string) (err error)
	ListPins(channel string) ([]slack.Item, *slack.Paging, error)
	ListPinsContext(ctx context.Context, channel string) ([]slack.Item, *slack.Paging, error)
	ListReactions(params slack.ListReactionsParameters) ([]slack.ReactedItem, *slack.Paging, error)
	ListReactionsContext(ctx context.Context, params slack.ListReactionsParameters) ([]slack.ReactedItem, *slack.Paging, error)
	ListStars(params slack.StarsParameters) ([]slack.Item, *slack.Paging, error)
	ListStarsContext(ctx context.Context, params slack.StarsParameters) ([]slack.Item, *slack.Paging, error)
	MarkIMChannel(channel, ts string) (err error)
	MarkIMChannelContext(ctx context.Context, channel, ts string) (err error)
	NewRTM(options ...slack.RTMOption) *slack.RTM
	NewRTMWithOptions(options *slack.RTMOptions) *slack.RTM
	OpenConversation(params *slack.OpenConversationParameters) (*slack.Channel, bool, bool, error)
	OpenConversationContext(ctx context.Context, params *slack.OpenConversationParameters) (*slack.Channel, bool, bool, error)
	OpenGroup(group string) (bool, bool, error)
	OpenGroupContext(ctx context.Context, group string) (bool, bool, error)
	OpenIMChannel(user string) (bool, bool, string, error)
	OpenIMChannelContext(ctx context.Context, user string) (bool, bool, string, error)
	PostEphemeral(channelID, userID string, options ...slack.MsgOption) (string, error)
	PostEphemeralContext(ctx context.Context, channelID, userID string, options ...slack.MsgOption) (timestamp string, err error)
	PostMessage(channel, text string, params slack.PostMessageParameters) (string, string, error)
	PostMessageContext(ctx context.Context, channel, text string, params slack.PostMessageParameters) (string, string, error)
	RemovePin(channel string, item slack.ItemRef) error
	RemovePinContext(ctx context.Context, channel string, item slack.ItemRef) error
	RemoveReaction(name string, item slack.ItemRef) error
	RemoveReactionContext(ctx context.Context, name string, item slack.ItemRef) error
	RemoveStar(channel string, item slack.ItemRef) error
	RemoveStarContext(ctx context.Context, channel string, item slack.ItemRef) error
	RenameChannel(channelID, name string) (*slack.Channel, error)
	RenameChannelContext(ctx context.Context, channelID, name string) (*slack.Channel, error)
	RenameConversation(channelID, channelName string) (*slack.Channel, error)
	RenameConversationContext(ctx context.Context, channelID, channelName string) (*slack.Channel, error)
	RenameGroup(group, name string) (*slack.Channel, error)
	RenameGroupContext(ctx context.Context, group, name string) (*slack.Channel, error)
	RevokeFilePublicURL(fileID string) (*slack.File, error)
	RevokeFilePublicURLContext(ctx context.Context, fileID string) (*slack.File, error)
	Search(query string, params slack.SearchParameters) (*slack.SearchMessages, *slack.SearchFiles, error)
	SearchContext(ctx context.Context, query string, params slack.SearchParameters) (*slack.SearchMessages, *slack.SearchFiles, error)
	SearchFiles(query string, params slack.SearchParameters) (*slack.SearchFiles, error)
	SearchFilesContext(ctx context.Context, query string, params slack.SearchParameters) (*slack.SearchFiles, error)
	SearchMessages(query string, params slack.SearchParameters) (*slack.SearchMessages, error)
	SearchMessagesContext(ctx context.Context, query string, params slack.SearchParameters) (*slack.SearchMessages, error)
	SendMessage(channel string, options ...slack.MsgOption) (string, string, string, error)
	SendMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (channel string, timestamp string, text string, err error)
	SendSSOBindingEmail(teamName, user string) error
	SendSSOBindingEmailContext(ctx context.Context, teamName, user string) error
	SetChannelPurpose(channelID, purpose string) (string, error)
	SetChannelPurposeContext(ctx context.Context, channelID, purpose string) (string, error)
	SetChannelReadMark(channelID, ts string) error
	SetChannelReadMarkContext(ctx context.Context, channelID, ts string) (err error)
	SetChannelTopic(channelID, topic string) (string, error)
	SetChannelTopicContext(ctx context.Context, channelID, topic string) (string, error)
	SetDebug(debug bool)
	SetGroupPurpose(group, purpose string) (string, error)
	SetGroupPurposeContext(ctx context.Context, group, purpose string) (string, error)
	SetGroupReadMark(group, ts string) error
	SetGroupReadMarkContext(ctx context.Context, group, ts string) (err error)
	SetGroupTopic(group, topic string) (string, error)
	SetGroupTopicContext(ctx context.Context, group, topic string) (string, error)
	SetPurposeOfConversation(channelID, purpose string) (*slack.Channel, error)
	SetPurposeOfConversationContext(ctx context.Context, channelID, purpose string) (*slack.Channel, error)
	SetRegular(teamName, user string) error
	SetRegularContext(ctx context.Context, teamName, user string) error
	SetRestricted(teamName, uid string) error
	SetRestrictedContext(ctx context.Context, teamName, uid string) error
	SetSnooze(minutes int) (*slack.DNDStatus, error)
	SetSnoozeContext(ctx context.Context, minutes int) (*slack.DNDStatus, error)
	SetTopicOfConversation(channelID, topic string) (*slack.Channel, error)
	SetTopicOfConversationContext(ctx context.Context, channelID, topic string) (*slack.Channel, error)
	SetUltraRestricted(teamName, uid, channel string) error
	SetUltraRestrictedContext(ctx context.Context, teamName, uid, channel string) error
	SetUserAsActive() error
	SetUserAsActiveContext(ctx context.Context) (err error)
	SetUserCustomStatus(statusText, statusEmoji string) error
	SetUserCustomStatusContext(ctx context.Context, statusText, statusEmoji string) error
	SetUserPhoto(image string, params slack.UserSetPhotoParams) error
	SetUserPhotoContext(ctx context.Context, image string, params slack.UserSetPhotoParams) error
	SetUserPresence(presence string) error
	SetUserPresenceContext(ctx context.Context, presence string) error
	ShareFilePublicURL(fileID string) (*slack.File, []slack.Comment, *slack.Paging, error)
	ShareFilePublicURLContext(ctx context.Context, fileID string) (*slack.File, []slack.Comment, *slack.Paging, error)
	StartRTM() (info *slack.Info, websocketURL string, err error)
	StartRTMContext(ctx context.Context) (info *slack.Info, websocketURL string, err error)
	UnArchiveConversation(channelID string) error
	UnArchiveConversationContext(ctx context.Context, channelID string) error
	UnarchiveChannel(channelID string) error
	UnarchiveChannelContext(ctx context.Context, channelID string) (err error)
	UnarchiveGroup(group string) error
	UnarchiveGroupContext(ctx context.Context, group string) error
	UnsetUserCustomStatus() error
	UnsetUserCustomStatusContext(ctx context.Context) error
	UpdateMessage(channelID, timestamp, text string) (string, string, string, error)
	UpdateMessageContext(ctx context.Context, channelID, timestamp, text string) (string, string, string, error)
	UpdateUserGroup(userGroup slack.UserGroup) (slack.UserGroup, error)
	UpdateUserGroupContext(ctx context.Context, userGroup slack.UserGroup) (slack.UserGroup, error)
	UpdateUserGroupMembers(userGroup string, members string) (slack.UserGroup, error)
	UpdateUserGroupMembersContext(ctx context.Context, userGroup string, members string) (slack.UserGroup, error)
	UploadFile(params slack.FileUploadParameters) (file *slack.File, err error)
	UploadFileContext(ctx context.Context, params slack.FileUploadParameters) (file *slack.File, err error)
}

// The DualSlackClient contains bot authentication and app authentication for Slack.
// By default, the client executes slack api methods with bot authentication.
// API methods that required user authentication execute with user authentication.
// The DualSlackClient satisfies the SlackClient interface.
type DualSlackClient struct {
	*slack.Client
	AppClient *slack.Client
}

// NewDualSlackClient creates a new instance of a DualSlackClient.
// The appToken corresponds to the "OAuth Access Token" for you Slack application.
// The botToken corresponds to the "Bot User OAuth Access Token" for you Slack application.
func NewDualSlackClient(appToken, botToken string) *DualSlackClient {
	return &DualSlackClient{
		Client:    slack.New(botToken),
		AppClient: slack.New(appToken),
	}
}

func (s *DualSlackClient) GetChannelHistory(channelID string, params slack.HistoryParameters) (*slack.History, error) {
	return s.AppClient.GetChannelHistory(channelID, params)
}

func (s *DualSlackClient) GetConversationHistory(params *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error) {
	return s.AppClient.GetConversationHistory(params)
}

func (s *DualSlackClient) GetIMHistory(channelID string, params slack.HistoryParameters) (*slack.History, error) {
	return s.AppClient.GetIMHistory(channelID, params)
}

func (s *DualSlackClient) GetGroupHistory(channelID string, params slack.HistoryParameters) (*slack.History, error) {
	return s.AppClient.GetGroupHistory(channelID, params)
}
