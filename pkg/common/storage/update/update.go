package update

import (
	"time"

	"github.com/openimsdk/open-im-server/v3/pkg/common/storage/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlackUpdate interface {
	Len() int
	Map() map[string]any
	WithOwnerUserID(value string) BlackUpdate
	WithPtrOwnerUserID(value *string) BlackUpdate
	DelOwnerUserID() BlackUpdate
	HasOwnerUserID() bool
	GetOwnerUserID() (string, bool)
	GetPtrOwnerUserID() *string
	WithBlockUserID(value string) BlackUpdate
	WithPtrBlockUserID(value *string) BlackUpdate
	DelBlockUserID() BlackUpdate
	HasBlockUserID() bool
	GetBlockUserID() (string, bool)
	GetPtrBlockUserID() *string
	WithCreateTime(value time.Time) BlackUpdate
	WithPtrCreateTime(value *time.Time) BlackUpdate
	DelCreateTime() BlackUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
	WithAddSource(value int32) BlackUpdate
	WithPtrAddSource(value *int32) BlackUpdate
	DelAddSource() BlackUpdate
	HasAddSource() bool
	GetAddSource() (int32, bool)
	GetPtrAddSource() *int32
	WithOperatorUserID(value string) BlackUpdate
	WithPtrOperatorUserID(value *string) BlackUpdate
	DelOperatorUserID() BlackUpdate
	HasOperatorUserID() bool
	GetOperatorUserID() (string, bool)
	GetPtrOperatorUserID() *string
	WithEx(value string) BlackUpdate
	WithPtrEx(value *string) BlackUpdate
	DelEx() BlackUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
}

type CacheUpdate interface {
	Len() int
	Map() map[string]any
	WithKey(value string) CacheUpdate
	WithPtrKey(value *string) CacheUpdate
	DelKey() CacheUpdate
	HasKey() bool
	GetKey() (string, bool)
	GetPtrKey() *string
	WithValue(value string) CacheUpdate
	WithPtrValue(value *string) CacheUpdate
	DelValue() CacheUpdate
	HasValue() bool
	GetValue() (string, bool)
	GetPtrValue() *string
	WithExpireAt(value *time.Time) CacheUpdate
	WithPtrExpireAt(value **time.Time) CacheUpdate
	DelExpireAt() CacheUpdate
	HasExpireAt() bool
	GetExpireAt() (*time.Time, bool)
	GetPtrExpireAt() **time.Time
}

type ConversationUpdate interface {
	Len() int
	Map() map[string]any
	WithOwnerUserID(value string) ConversationUpdate
	WithPtrOwnerUserID(value *string) ConversationUpdate
	DelOwnerUserID() ConversationUpdate
	HasOwnerUserID() bool
	GetOwnerUserID() (string, bool)
	GetPtrOwnerUserID() *string
	WithConversationID(value string) ConversationUpdate
	WithPtrConversationID(value *string) ConversationUpdate
	DelConversationID() ConversationUpdate
	HasConversationID() bool
	GetConversationID() (string, bool)
	GetPtrConversationID() *string
	WithConversationType(value int32) ConversationUpdate
	WithPtrConversationType(value *int32) ConversationUpdate
	DelConversationType() ConversationUpdate
	HasConversationType() bool
	GetConversationType() (int32, bool)
	GetPtrConversationType() *int32
	WithUserID(value string) ConversationUpdate
	WithPtrUserID(value *string) ConversationUpdate
	DelUserID() ConversationUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithGroupID(value string) ConversationUpdate
	WithPtrGroupID(value *string) ConversationUpdate
	DelGroupID() ConversationUpdate
	HasGroupID() bool
	GetGroupID() (string, bool)
	GetPtrGroupID() *string
	WithRecvMsgOpt(value int32) ConversationUpdate
	WithPtrRecvMsgOpt(value *int32) ConversationUpdate
	DelRecvMsgOpt() ConversationUpdate
	HasRecvMsgOpt() bool
	GetRecvMsgOpt() (int32, bool)
	GetPtrRecvMsgOpt() *int32
	WithIsPinned(value bool) ConversationUpdate
	WithPtrIsPinned(value *bool) ConversationUpdate
	DelIsPinned() ConversationUpdate
	HasIsPinned() bool
	GetIsPinned() (bool, bool)
	GetPtrIsPinned() *bool
	WithIsPrivateChat(value bool) ConversationUpdate
	WithPtrIsPrivateChat(value *bool) ConversationUpdate
	DelIsPrivateChat() ConversationUpdate
	HasIsPrivateChat() bool
	GetIsPrivateChat() (bool, bool)
	GetPtrIsPrivateChat() *bool
	WithBurnDuration(value int32) ConversationUpdate
	WithPtrBurnDuration(value *int32) ConversationUpdate
	DelBurnDuration() ConversationUpdate
	HasBurnDuration() bool
	GetBurnDuration() (int32, bool)
	GetPtrBurnDuration() *int32
	WithGroupAtType(value int32) ConversationUpdate
	WithPtrGroupAtType(value *int32) ConversationUpdate
	DelGroupAtType() ConversationUpdate
	HasGroupAtType() bool
	GetGroupAtType() (int32, bool)
	GetPtrGroupAtType() *int32
	WithAttachedInfo(value string) ConversationUpdate
	WithPtrAttachedInfo(value *string) ConversationUpdate
	DelAttachedInfo() ConversationUpdate
	HasAttachedInfo() bool
	GetAttachedInfo() (string, bool)
	GetPtrAttachedInfo() *string
	WithEx(value string) ConversationUpdate
	WithPtrEx(value *string) ConversationUpdate
	DelEx() ConversationUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
	WithMaxSeq(value int64) ConversationUpdate
	WithPtrMaxSeq(value *int64) ConversationUpdate
	DelMaxSeq() ConversationUpdate
	HasMaxSeq() bool
	GetMaxSeq() (int64, bool)
	GetPtrMaxSeq() *int64
	WithMinSeq(value int64) ConversationUpdate
	WithPtrMinSeq(value *int64) ConversationUpdate
	DelMinSeq() ConversationUpdate
	HasMinSeq() bool
	GetMinSeq() (int64, bool)
	GetPtrMinSeq() *int64
	WithCreateTime(value time.Time) ConversationUpdate
	WithPtrCreateTime(value *time.Time) ConversationUpdate
	DelCreateTime() ConversationUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
	WithIsMsgDestruct(value bool) ConversationUpdate
	WithPtrIsMsgDestruct(value *bool) ConversationUpdate
	DelIsMsgDestruct() ConversationUpdate
	HasIsMsgDestruct() bool
	GetIsMsgDestruct() (bool, bool)
	GetPtrIsMsgDestruct() *bool
	WithMsgDestructTime(value int64) ConversationUpdate
	WithPtrMsgDestructTime(value *int64) ConversationUpdate
	DelMsgDestructTime() ConversationUpdate
	HasMsgDestructTime() bool
	GetMsgDestructTime() (int64, bool)
	GetPtrMsgDestructTime() *int64
	WithLatestMsgDestructTime(value time.Time) ConversationUpdate
	WithPtrLatestMsgDestructTime(value *time.Time) ConversationUpdate
	DelLatestMsgDestructTime() ConversationUpdate
	HasLatestMsgDestructTime() bool
	GetLatestMsgDestructTime() (time.Time, bool)
	GetPtrLatestMsgDestructTime() *time.Time
}

type FriendUpdate interface {
	Len() int
	Map() map[string]any
	WithID(value primitive.ObjectID) FriendUpdate
	WithPtrID(value *primitive.ObjectID) FriendUpdate
	DelID() FriendUpdate
	HasID() bool
	GetID() (primitive.ObjectID, bool)
	GetPtrID() *primitive.ObjectID
	WithOwnerUserID(value string) FriendUpdate
	WithPtrOwnerUserID(value *string) FriendUpdate
	DelOwnerUserID() FriendUpdate
	HasOwnerUserID() bool
	GetOwnerUserID() (string, bool)
	GetPtrOwnerUserID() *string
	WithFriendUserID(value string) FriendUpdate
	WithPtrFriendUserID(value *string) FriendUpdate
	DelFriendUserID() FriendUpdate
	HasFriendUserID() bool
	GetFriendUserID() (string, bool)
	GetPtrFriendUserID() *string
	WithRemark(value string) FriendUpdate
	WithPtrRemark(value *string) FriendUpdate
	DelRemark() FriendUpdate
	HasRemark() bool
	GetRemark() (string, bool)
	GetPtrRemark() *string
	WithCreateTime(value time.Time) FriendUpdate
	WithPtrCreateTime(value *time.Time) FriendUpdate
	DelCreateTime() FriendUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
	WithAddSource(value int32) FriendUpdate
	WithPtrAddSource(value *int32) FriendUpdate
	DelAddSource() FriendUpdate
	HasAddSource() bool
	GetAddSource() (int32, bool)
	GetPtrAddSource() *int32
	WithOperatorUserID(value string) FriendUpdate
	WithPtrOperatorUserID(value *string) FriendUpdate
	DelOperatorUserID() FriendUpdate
	HasOperatorUserID() bool
	GetOperatorUserID() (string, bool)
	GetPtrOperatorUserID() *string
	WithEx(value string) FriendUpdate
	WithPtrEx(value *string) FriendUpdate
	DelEx() FriendUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
	WithIsPinned(value bool) FriendUpdate
	WithPtrIsPinned(value *bool) FriendUpdate
	DelIsPinned() FriendUpdate
	HasIsPinned() bool
	GetIsPinned() (bool, bool)
	GetPtrIsPinned() *bool
}

type FriendRequestUpdate interface {
	Len() int
	Map() map[string]any
	WithFromUserID(value string) FriendRequestUpdate
	WithPtrFromUserID(value *string) FriendRequestUpdate
	DelFromUserID() FriendRequestUpdate
	HasFromUserID() bool
	GetFromUserID() (string, bool)
	GetPtrFromUserID() *string
	WithToUserID(value string) FriendRequestUpdate
	WithPtrToUserID(value *string) FriendRequestUpdate
	DelToUserID() FriendRequestUpdate
	HasToUserID() bool
	GetToUserID() (string, bool)
	GetPtrToUserID() *string
	WithHandleResult(value int32) FriendRequestUpdate
	WithPtrHandleResult(value *int32) FriendRequestUpdate
	DelHandleResult() FriendRequestUpdate
	HasHandleResult() bool
	GetHandleResult() (int32, bool)
	GetPtrHandleResult() *int32
	WithReqMsg(value string) FriendRequestUpdate
	WithPtrReqMsg(value *string) FriendRequestUpdate
	DelReqMsg() FriendRequestUpdate
	HasReqMsg() bool
	GetReqMsg() (string, bool)
	GetPtrReqMsg() *string
	WithCreateTime(value time.Time) FriendRequestUpdate
	WithPtrCreateTime(value *time.Time) FriendRequestUpdate
	DelCreateTime() FriendRequestUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
	WithHandlerUserID(value string) FriendRequestUpdate
	WithPtrHandlerUserID(value *string) FriendRequestUpdate
	DelHandlerUserID() FriendRequestUpdate
	HasHandlerUserID() bool
	GetHandlerUserID() (string, bool)
	GetPtrHandlerUserID() *string
	WithHandleMsg(value string) FriendRequestUpdate
	WithPtrHandleMsg(value *string) FriendRequestUpdate
	DelHandleMsg() FriendRequestUpdate
	HasHandleMsg() bool
	GetHandleMsg() (string, bool)
	GetPtrHandleMsg() *string
	WithHandleTime(value time.Time) FriendRequestUpdate
	WithPtrHandleTime(value *time.Time) FriendRequestUpdate
	DelHandleTime() FriendRequestUpdate
	HasHandleTime() bool
	GetHandleTime() (time.Time, bool)
	GetPtrHandleTime() *time.Time
	WithEx(value string) FriendRequestUpdate
	WithPtrEx(value *string) FriendRequestUpdate
	DelEx() FriendRequestUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
}

type GroupUpdate interface {
	Len() int
	Map() map[string]any
	WithGroupID(value string) GroupUpdate
	WithPtrGroupID(value *string) GroupUpdate
	DelGroupID() GroupUpdate
	HasGroupID() bool
	GetGroupID() (string, bool)
	GetPtrGroupID() *string
	WithGroupName(value string) GroupUpdate
	WithPtrGroupName(value *string) GroupUpdate
	DelGroupName() GroupUpdate
	HasGroupName() bool
	GetGroupName() (string, bool)
	GetPtrGroupName() *string
	WithNotification(value string) GroupUpdate
	WithPtrNotification(value *string) GroupUpdate
	DelNotification() GroupUpdate
	HasNotification() bool
	GetNotification() (string, bool)
	GetPtrNotification() *string
	WithIntroduction(value string) GroupUpdate
	WithPtrIntroduction(value *string) GroupUpdate
	DelIntroduction() GroupUpdate
	HasIntroduction() bool
	GetIntroduction() (string, bool)
	GetPtrIntroduction() *string
	WithFaceURL(value string) GroupUpdate
	WithPtrFaceURL(value *string) GroupUpdate
	DelFaceURL() GroupUpdate
	HasFaceURL() bool
	GetFaceURL() (string, bool)
	GetPtrFaceURL() *string
	WithCreateTime(value time.Time) GroupUpdate
	WithPtrCreateTime(value *time.Time) GroupUpdate
	DelCreateTime() GroupUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
	WithEx(value string) GroupUpdate
	WithPtrEx(value *string) GroupUpdate
	DelEx() GroupUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
	WithStatus(value int32) GroupUpdate
	WithPtrStatus(value *int32) GroupUpdate
	DelStatus() GroupUpdate
	HasStatus() bool
	GetStatus() (int32, bool)
	GetPtrStatus() *int32
	WithCreatorUserID(value string) GroupUpdate
	WithPtrCreatorUserID(value *string) GroupUpdate
	DelCreatorUserID() GroupUpdate
	HasCreatorUserID() bool
	GetCreatorUserID() (string, bool)
	GetPtrCreatorUserID() *string
	WithGroupType(value int32) GroupUpdate
	WithPtrGroupType(value *int32) GroupUpdate
	DelGroupType() GroupUpdate
	HasGroupType() bool
	GetGroupType() (int32, bool)
	GetPtrGroupType() *int32
	WithNeedVerification(value int32) GroupUpdate
	WithPtrNeedVerification(value *int32) GroupUpdate
	DelNeedVerification() GroupUpdate
	HasNeedVerification() bool
	GetNeedVerification() (int32, bool)
	GetPtrNeedVerification() *int32
	WithLookMemberInfo(value int32) GroupUpdate
	WithPtrLookMemberInfo(value *int32) GroupUpdate
	DelLookMemberInfo() GroupUpdate
	HasLookMemberInfo() bool
	GetLookMemberInfo() (int32, bool)
	GetPtrLookMemberInfo() *int32
	WithApplyMemberFriend(value int32) GroupUpdate
	WithPtrApplyMemberFriend(value *int32) GroupUpdate
	DelApplyMemberFriend() GroupUpdate
	HasApplyMemberFriend() bool
	GetApplyMemberFriend() (int32, bool)
	GetPtrApplyMemberFriend() *int32
	WithNotificationUpdateTime(value time.Time) GroupUpdate
	WithPtrNotificationUpdateTime(value *time.Time) GroupUpdate
	DelNotificationUpdateTime() GroupUpdate
	HasNotificationUpdateTime() bool
	GetNotificationUpdateTime() (time.Time, bool)
	GetPtrNotificationUpdateTime() *time.Time
	WithNotificationUserID(value string) GroupUpdate
	WithPtrNotificationUserID(value *string) GroupUpdate
	DelNotificationUserID() GroupUpdate
	HasNotificationUserID() bool
	GetNotificationUserID() (string, bool)
	GetPtrNotificationUserID() *string
}

type GroupMemberUpdate interface {
	Len() int
	Map() map[string]any
	WithGroupID(value string) GroupMemberUpdate
	WithPtrGroupID(value *string) GroupMemberUpdate
	DelGroupID() GroupMemberUpdate
	HasGroupID() bool
	GetGroupID() (string, bool)
	GetPtrGroupID() *string
	WithUserID(value string) GroupMemberUpdate
	WithPtrUserID(value *string) GroupMemberUpdate
	DelUserID() GroupMemberUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithNickname(value string) GroupMemberUpdate
	WithPtrNickname(value *string) GroupMemberUpdate
	DelNickname() GroupMemberUpdate
	HasNickname() bool
	GetNickname() (string, bool)
	GetPtrNickname() *string
	WithFaceURL(value string) GroupMemberUpdate
	WithPtrFaceURL(value *string) GroupMemberUpdate
	DelFaceURL() GroupMemberUpdate
	HasFaceURL() bool
	GetFaceURL() (string, bool)
	GetPtrFaceURL() *string
	WithRoleLevel(value int32) GroupMemberUpdate
	WithPtrRoleLevel(value *int32) GroupMemberUpdate
	DelRoleLevel() GroupMemberUpdate
	HasRoleLevel() bool
	GetRoleLevel() (int32, bool)
	GetPtrRoleLevel() *int32
	WithJoinTime(value time.Time) GroupMemberUpdate
	WithPtrJoinTime(value *time.Time) GroupMemberUpdate
	DelJoinTime() GroupMemberUpdate
	HasJoinTime() bool
	GetJoinTime() (time.Time, bool)
	GetPtrJoinTime() *time.Time
	WithJoinSource(value int32) GroupMemberUpdate
	WithPtrJoinSource(value *int32) GroupMemberUpdate
	DelJoinSource() GroupMemberUpdate
	HasJoinSource() bool
	GetJoinSource() (int32, bool)
	GetPtrJoinSource() *int32
	WithInviterUserID(value string) GroupMemberUpdate
	WithPtrInviterUserID(value *string) GroupMemberUpdate
	DelInviterUserID() GroupMemberUpdate
	HasInviterUserID() bool
	GetInviterUserID() (string, bool)
	GetPtrInviterUserID() *string
	WithOperatorUserID(value string) GroupMemberUpdate
	WithPtrOperatorUserID(value *string) GroupMemberUpdate
	DelOperatorUserID() GroupMemberUpdate
	HasOperatorUserID() bool
	GetOperatorUserID() (string, bool)
	GetPtrOperatorUserID() *string
	WithMuteEndTime(value time.Time) GroupMemberUpdate
	WithPtrMuteEndTime(value *time.Time) GroupMemberUpdate
	DelMuteEndTime() GroupMemberUpdate
	HasMuteEndTime() bool
	GetMuteEndTime() (time.Time, bool)
	GetPtrMuteEndTime() *time.Time
	WithEx(value string) GroupMemberUpdate
	WithPtrEx(value *string) GroupMemberUpdate
	DelEx() GroupMemberUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
}

type GroupRequestUpdate interface {
	Len() int
	Map() map[string]any
	WithUserID(value string) GroupRequestUpdate
	WithPtrUserID(value *string) GroupRequestUpdate
	DelUserID() GroupRequestUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithGroupID(value string) GroupRequestUpdate
	WithPtrGroupID(value *string) GroupRequestUpdate
	DelGroupID() GroupRequestUpdate
	HasGroupID() bool
	GetGroupID() (string, bool)
	GetPtrGroupID() *string
	WithHandleResult(value int32) GroupRequestUpdate
	WithPtrHandleResult(value *int32) GroupRequestUpdate
	DelHandleResult() GroupRequestUpdate
	HasHandleResult() bool
	GetHandleResult() (int32, bool)
	GetPtrHandleResult() *int32
	WithReqMsg(value string) GroupRequestUpdate
	WithPtrReqMsg(value *string) GroupRequestUpdate
	DelReqMsg() GroupRequestUpdate
	HasReqMsg() bool
	GetReqMsg() (string, bool)
	GetPtrReqMsg() *string
	WithHandledMsg(value string) GroupRequestUpdate
	WithPtrHandledMsg(value *string) GroupRequestUpdate
	DelHandledMsg() GroupRequestUpdate
	HasHandledMsg() bool
	GetHandledMsg() (string, bool)
	GetPtrHandledMsg() *string
	WithReqTime(value time.Time) GroupRequestUpdate
	WithPtrReqTime(value *time.Time) GroupRequestUpdate
	DelReqTime() GroupRequestUpdate
	HasReqTime() bool
	GetReqTime() (time.Time, bool)
	GetPtrReqTime() *time.Time
	WithHandleUserID(value string) GroupRequestUpdate
	WithPtrHandleUserID(value *string) GroupRequestUpdate
	DelHandleUserID() GroupRequestUpdate
	HasHandleUserID() bool
	GetHandleUserID() (string, bool)
	GetPtrHandleUserID() *string
	WithHandledTime(value time.Time) GroupRequestUpdate
	WithPtrHandledTime(value *time.Time) GroupRequestUpdate
	DelHandledTime() GroupRequestUpdate
	HasHandledTime() bool
	GetHandledTime() (time.Time, bool)
	GetPtrHandledTime() *time.Time
	WithJoinSource(value int32) GroupRequestUpdate
	WithPtrJoinSource(value *int32) GroupRequestUpdate
	DelJoinSource() GroupRequestUpdate
	HasJoinSource() bool
	GetJoinSource() (int32, bool)
	GetPtrJoinSource() *int32
	WithInviterUserID(value string) GroupRequestUpdate
	WithPtrInviterUserID(value *string) GroupRequestUpdate
	DelInviterUserID() GroupRequestUpdate
	HasInviterUserID() bool
	GetInviterUserID() (string, bool)
	GetPtrInviterUserID() *string
	WithEx(value string) GroupRequestUpdate
	WithPtrEx(value *string) GroupRequestUpdate
	DelEx() GroupRequestUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
}

type LogUpdate interface {
	Len() int
	Map() map[string]any
	WithLogID(value string) LogUpdate
	WithPtrLogID(value *string) LogUpdate
	DelLogID() LogUpdate
	HasLogID() bool
	GetLogID() (string, bool)
	GetPtrLogID() *string
	WithPlatform(value string) LogUpdate
	WithPtrPlatform(value *string) LogUpdate
	DelPlatform() LogUpdate
	HasPlatform() bool
	GetPlatform() (string, bool)
	GetPtrPlatform() *string
	WithUserID(value string) LogUpdate
	WithPtrUserID(value *string) LogUpdate
	DelUserID() LogUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithCreateTime(value time.Time) LogUpdate
	WithPtrCreateTime(value *time.Time) LogUpdate
	DelCreateTime() LogUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
	WithUrl(value string) LogUpdate
	WithPtrUrl(value *string) LogUpdate
	DelUrl() LogUpdate
	HasUrl() bool
	GetUrl() (string, bool)
	GetPtrUrl() *string
	WithFileName(value string) LogUpdate
	WithPtrFileName(value *string) LogUpdate
	DelFileName() LogUpdate
	HasFileName() bool
	GetFileName() (string, bool)
	GetPtrFileName() *string
	WithSystemType(value string) LogUpdate
	WithPtrSystemType(value *string) LogUpdate
	DelSystemType() LogUpdate
	HasSystemType() bool
	GetSystemType() (string, bool)
	GetPtrSystemType() *string
	WithAppFramework(value string) LogUpdate
	WithPtrAppFramework(value *string) LogUpdate
	DelAppFramework() LogUpdate
	HasAppFramework() bool
	GetAppFramework() (string, bool)
	GetPtrAppFramework() *string
	WithVersion(value string) LogUpdate
	WithPtrVersion(value *string) LogUpdate
	DelVersion() LogUpdate
	HasVersion() bool
	GetVersion() (string, bool)
	GetPtrVersion() *string
	WithEx(value string) LogUpdate
	WithPtrEx(value *string) LogUpdate
	DelEx() LogUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
}

type MsgDocModelUpdate interface {
	Len() int
	Map() map[string]any
	WithDocID(value string) MsgDocModelUpdate
	WithPtrDocID(value *string) MsgDocModelUpdate
	DelDocID() MsgDocModelUpdate
	HasDocID() bool
	GetDocID() (string, bool)
	GetPtrDocID() *string
	WithMsg(value []*model.MsgInfoModel) MsgDocModelUpdate
	WithPtrMsg(value *[]*model.MsgInfoModel) MsgDocModelUpdate
	DelMsg() MsgDocModelUpdate
	HasMsg() bool
	GetMsg() ([]*model.MsgInfoModel, bool)
	GetPtrMsg() *[]*model.MsgInfoModel
}

type ObjectUpdate interface {
	Len() int
	Map() map[string]any
	WithName(value string) ObjectUpdate
	WithPtrName(value *string) ObjectUpdate
	DelName() ObjectUpdate
	HasName() bool
	GetName() (string, bool)
	GetPtrName() *string
	WithUserID(value string) ObjectUpdate
	WithPtrUserID(value *string) ObjectUpdate
	DelUserID() ObjectUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithHash(value string) ObjectUpdate
	WithPtrHash(value *string) ObjectUpdate
	DelHash() ObjectUpdate
	HasHash() bool
	GetHash() (string, bool)
	GetPtrHash() *string
	WithEngine(value string) ObjectUpdate
	WithPtrEngine(value *string) ObjectUpdate
	DelEngine() ObjectUpdate
	HasEngine() bool
	GetEngine() (string, bool)
	GetPtrEngine() *string
	WithKey(value string) ObjectUpdate
	WithPtrKey(value *string) ObjectUpdate
	DelKey() ObjectUpdate
	HasKey() bool
	GetKey() (string, bool)
	GetPtrKey() *string
	WithSize(value int64) ObjectUpdate
	WithPtrSize(value *int64) ObjectUpdate
	DelSize() ObjectUpdate
	HasSize() bool
	GetSize() (int64, bool)
	GetPtrSize() *int64
	WithContentType(value string) ObjectUpdate
	WithPtrContentType(value *string) ObjectUpdate
	DelContentType() ObjectUpdate
	HasContentType() bool
	GetContentType() (string, bool)
	GetPtrContentType() *string
	WithGroup(value string) ObjectUpdate
	WithPtrGroup(value *string) ObjectUpdate
	DelGroup() ObjectUpdate
	HasGroup() bool
	GetGroup() (string, bool)
	GetPtrGroup() *string
	WithCreateTime(value time.Time) ObjectUpdate
	WithPtrCreateTime(value *time.Time) ObjectUpdate
	DelCreateTime() ObjectUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
}

type SeqConversationUpdate interface {
	Len() int
	Map() map[string]any
	WithConversationID(value string) SeqConversationUpdate
	WithPtrConversationID(value *string) SeqConversationUpdate
	DelConversationID() SeqConversationUpdate
	HasConversationID() bool
	GetConversationID() (string, bool)
	GetPtrConversationID() *string
	WithMaxSeq(value int64) SeqConversationUpdate
	WithPtrMaxSeq(value *int64) SeqConversationUpdate
	DelMaxSeq() SeqConversationUpdate
	HasMaxSeq() bool
	GetMaxSeq() (int64, bool)
	GetPtrMaxSeq() *int64
	WithMinSeq(value int64) SeqConversationUpdate
	WithPtrMinSeq(value *int64) SeqConversationUpdate
	DelMinSeq() SeqConversationUpdate
	HasMinSeq() bool
	GetMinSeq() (int64, bool)
	GetPtrMinSeq() *int64
}

type SeqUserUpdate interface {
	Len() int
	Map() map[string]any
	WithUserID(value string) SeqUserUpdate
	WithPtrUserID(value *string) SeqUserUpdate
	DelUserID() SeqUserUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithConversationID(value string) SeqUserUpdate
	WithPtrConversationID(value *string) SeqUserUpdate
	DelConversationID() SeqUserUpdate
	HasConversationID() bool
	GetConversationID() (string, bool)
	GetPtrConversationID() *string
	WithMinSeq(value int64) SeqUserUpdate
	WithPtrMinSeq(value *int64) SeqUserUpdate
	DelMinSeq() SeqUserUpdate
	HasMinSeq() bool
	GetMinSeq() (int64, bool)
	GetPtrMinSeq() *int64
	WithMaxSeq(value int64) SeqUserUpdate
	WithPtrMaxSeq(value *int64) SeqUserUpdate
	DelMaxSeq() SeqUserUpdate
	HasMaxSeq() bool
	GetMaxSeq() (int64, bool)
	GetPtrMaxSeq() *int64
	WithReadSeq(value int64) SeqUserUpdate
	WithPtrReadSeq(value *int64) SeqUserUpdate
	DelReadSeq() SeqUserUpdate
	HasReadSeq() bool
	GetReadSeq() (int64, bool)
	GetPtrReadSeq() *int64
}

type UserUpdate interface {
	Len() int
	Map() map[string]any
	WithUserID(value string) UserUpdate
	WithPtrUserID(value *string) UserUpdate
	DelUserID() UserUpdate
	HasUserID() bool
	GetUserID() (string, bool)
	GetPtrUserID() *string
	WithNickname(value string) UserUpdate
	WithPtrNickname(value *string) UserUpdate
	DelNickname() UserUpdate
	HasNickname() bool
	GetNickname() (string, bool)
	GetPtrNickname() *string
	WithFaceURL(value string) UserUpdate
	WithPtrFaceURL(value *string) UserUpdate
	DelFaceURL() UserUpdate
	HasFaceURL() bool
	GetFaceURL() (string, bool)
	GetPtrFaceURL() *string
	WithEx(value string) UserUpdate
	WithPtrEx(value *string) UserUpdate
	DelEx() UserUpdate
	HasEx() bool
	GetEx() (string, bool)
	GetPtrEx() *string
	WithAppMangerLevel(value int32) UserUpdate
	WithPtrAppMangerLevel(value *int32) UserUpdate
	DelAppMangerLevel() UserUpdate
	HasAppMangerLevel() bool
	GetAppMangerLevel() (int32, bool)
	GetPtrAppMangerLevel() *int32
	WithGlobalRecvMsgOpt(value int32) UserUpdate
	WithPtrGlobalRecvMsgOpt(value *int32) UserUpdate
	DelGlobalRecvMsgOpt() UserUpdate
	HasGlobalRecvMsgOpt() bool
	GetGlobalRecvMsgOpt() (int32, bool)
	GetPtrGlobalRecvMsgOpt() *int32
	WithCreateTime(value time.Time) UserUpdate
	WithPtrCreateTime(value *time.Time) UserUpdate
	DelCreateTime() UserUpdate
	HasCreateTime() bool
	GetCreateTime() (time.Time, bool)
	GetPtrCreateTime() *time.Time
}

type VersionLogUpdate interface {
	Len() int
	Map() map[string]any
	WithID(value primitive.ObjectID) VersionLogUpdate
	WithPtrID(value *primitive.ObjectID) VersionLogUpdate
	DelID() VersionLogUpdate
	HasID() bool
	GetID() (primitive.ObjectID, bool)
	GetPtrID() *primitive.ObjectID
	WithDID(value string) VersionLogUpdate
	WithPtrDID(value *string) VersionLogUpdate
	DelDID() VersionLogUpdate
	HasDID() bool
	GetDID() (string, bool)
	GetPtrDID() *string
	WithLogs(value []model.VersionLogElem) VersionLogUpdate
	WithPtrLogs(value *[]model.VersionLogElem) VersionLogUpdate
	DelLogs() VersionLogUpdate
	HasLogs() bool
	GetLogs() ([]model.VersionLogElem, bool)
	GetPtrLogs() *[]model.VersionLogElem
	WithVersion(value uint) VersionLogUpdate
	WithPtrVersion(value *uint) VersionLogUpdate
	DelVersion() VersionLogUpdate
	HasVersion() bool
	GetVersion() (uint, bool)
	GetPtrVersion() *uint
	WithDeleted(value uint) VersionLogUpdate
	WithPtrDeleted(value *uint) VersionLogUpdate
	DelDeleted() VersionLogUpdate
	HasDeleted() bool
	GetDeleted() (uint, bool)
	GetPtrDeleted() *uint
	WithLastUpdate(value time.Time) VersionLogUpdate
	WithPtrLastUpdate(value *time.Time) VersionLogUpdate
	DelLastUpdate() VersionLogUpdate
	HasLastUpdate() bool
	GetLastUpdate() (time.Time, bool)
	GetPtrLastUpdate() *time.Time
	WithLogLen(value int) VersionLogUpdate
	WithPtrLogLen(value *int) VersionLogUpdate
	DelLogLen() VersionLogUpdate
	HasLogLen() bool
	GetLogLen() (int, bool)
	GetPtrLogLen() *int
}

type Update interface {
	NewBlackUpdate() BlackUpdate
	NewCacheUpdate() CacheUpdate
	NewConversationUpdate() ConversationUpdate
	NewFriendUpdate() FriendUpdate
	NewFriendRequestUpdate() FriendRequestUpdate
	NewGroupUpdate() GroupUpdate
	NewGroupMemberUpdate() GroupMemberUpdate
	NewGroupRequestUpdate() GroupRequestUpdate
	NewLogUpdate() LogUpdate
	NewMsgDocModelUpdate() MsgDocModelUpdate
	NewObjectUpdate() ObjectUpdate
	NewSeqConversationUpdate() SeqConversationUpdate
	NewSeqUserUpdate() SeqUserUpdate
	NewUserUpdate() UserUpdate
	NewVersionLogUpdate() VersionLogUpdate
}
