package api

import (
	"OpenIM/pkg/common/config"
	"OpenIM/pkg/common/log"
	"OpenIM/pkg/common/mw"
	"OpenIM/pkg/common/prome"
	"OpenIM/pkg/discoveryregistry"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func NewGinRouter(zk discoveryregistry.SvcDiscoveryRegistry) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create("../logs/api.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//	gin.SetMode(gin.DebugMode)
	r := gin.New()
	log.Info("load config: ", config.Config)
	r.Use(gin.Recovery(), mw.CorsHandler(), mw.GinParseOperationID())
	if config.Config.Prometheus.Enable {
		prome.NewApiRequestCounter()
		prome.NewApiRequestFailedCounter()
		prome.NewApiRequestSuccessCounter()
		r.Use(prome.PrometheusMiddleware)
		r.GET("/metrics", prome.PrometheusHandler())
	}
	zk.AddOption(mw.GrpcClient()) // 默认RPC中间件
	userRouterGroup := r.Group("/user")
	{
		u := NewUser(zk)
		userRouterGroup.POST("/user_register", u.UserRegister)
		userRouterGroup.POST("/update_user_info", u.UpdateUserInfo) //1
		userRouterGroup.POST("/set_global_msg_recv_opt", u.SetGlobalRecvMessageOpt)
		userRouterGroup.POST("/get_users_info", u.GetUsersPublicInfo) //1
		userRouterGroup.POST("/get_all_users_uid", u.GetAllUsersID)   // todo
		userRouterGroup.POST("/account_check", u.AccountCheck)        // todo
		userRouterGroup.POST("/get_users", u.GetUsers)
	}
	////friend routing group
	friendRouterGroup := r.Group("/friend")
	{
		f := NewFriend(zk)
		friendRouterGroup.POST("/add_friend", f.ApplyToAddFriend)                 //1
		friendRouterGroup.POST("/delete_friend", f.DeleteFriend)                  //1
		friendRouterGroup.POST("/get_friend_apply_list", f.GetFriendApplyList)    //1
		friendRouterGroup.POST("/get_self_friend_apply_list", f.GetSelfApplyList) //1
		friendRouterGroup.POST("/get_friend_list", f.GetFriendList)               //1
		friendRouterGroup.POST("/add_friend_response", f.RespondFriendApply)      //1
		friendRouterGroup.POST("/set_friend_remark", f.SetFriendRemark)           //1
		friendRouterGroup.POST("/add_black", f.AddBlack)                          //1
		friendRouterGroup.POST("/get_black_list", f.GetPaginationBlacks)          //1
		friendRouterGroup.POST("/remove_black", f.RemoveBlack)                    //1
		friendRouterGroup.POST("/import_friend", f.ImportFriends)                 //1
		friendRouterGroup.POST("/is_friend", f.IsFriend)                          //1

	}
	groupRouterGroup := r.Group("/group")
	g := NewGroup(zk)
	{
		groupRouterGroup.POST("/create_group", g.NewCreateGroup)                                //1
		groupRouterGroup.POST("/set_group_info", g.NewSetGroupInfo)                             //1
		groupRouterGroup.POST("/join_group", g.JoinGroup)                                       //1
		groupRouterGroup.POST("/quit_group", g.QuitGroup)                                       //1
		groupRouterGroup.POST("/group_application_response", g.ApplicationGroupResponse)        //1
		groupRouterGroup.POST("/transfer_group", g.TransferGroupOwner)                          //1
		groupRouterGroup.POST("/get_recv_group_applicationList", g.GetRecvGroupApplicationList) //1
		groupRouterGroup.POST("/get_user_req_group_applicationList", g.GetUserReqGroupApplicationList)
		groupRouterGroup.POST("/get_groups_info", g.GetGroupsInfo) //1
		groupRouterGroup.POST("/kick_group", g.KickGroupMember)    //1
		//groupRouterGroup.POST("/get_group_all_member_list", g.GetGroupAllMemberList) //1
		groupRouterGroup.POST("/get_group_members_info", g.GetGroupMembersInfo) //1
		groupRouterGroup.POST("/invite_user_to_group", g.InviteUserToGroup)     //1
		groupRouterGroup.POST("/get_joined_group_list", g.GetJoinedGroupList)
		groupRouterGroup.POST("/dismiss_group", g.DismissGroup) //
		groupRouterGroup.POST("/mute_group_member", g.MuteGroupMember)
		groupRouterGroup.POST("/cancel_mute_group_member", g.CancelMuteGroupMember) //MuteGroup
		groupRouterGroup.POST("/mute_group", g.MuteGroup)
		groupRouterGroup.POST("/cancel_mute_group", g.CancelMuteGroup)
		//groupRouterGroup.POST("/set_group_member_nickname", g.SetGroupMemberNickname)
		groupRouterGroup.POST("/set_group_member_info", g.SetGroupMemberInfo)
		groupRouterGroup.POST("/get_group_abstract_info", g.GetGroupAbstractInfo)
	}
	superGroupRouterGroup := r.Group("/super_group")
	{
		superGroupRouterGroup.POST("/get_joined_group_list", g.GetJoinedSuperGroupList)
		superGroupRouterGroup.POST("/get_groups_info", g.GetSuperGroupsInfo)
	}
	////certificate
	authRouterGroup := r.Group("/auth")
	{
		a := NewAuth(zk)
		u := NewUser(zk)
		authRouterGroup.POST("/user_register", u.UserRegister) //1
		authRouterGroup.POST("/user_token", a.UserToken)       //1
		authRouterGroup.POST("/parse_token", a.ParseToken)     //1
		authRouterGroup.POST("/force_logout", a.ForceLogout)   //1
	}
	////Third service
	thirdGroup := r.Group("/third")
	{
		t := NewThird(zk)
		thirdGroup.POST("/get_rtc_invitation_info", t.GetSignalInvitationInfo)
		thirdGroup.POST("/get_rtc_invitation_start_app", t.GetSignalInvitationInfoStartApp)
		thirdGroup.POST("/fcm_update_token", t.FcmUpdateToken)
		thirdGroup.POST("/set_app_badge", t.SetAppBadge)

		thirdGroup.POST("/apply_put", t.ApplyPut)
		thirdGroup.POST("/get_put", t.GetPut)
		thirdGroup.POST("/confirm_put", t.ConfirmPut)
	}
	////Message
	chatGroup := r.Group("/msg")
	{
		m := NewMsg(zk)
		chatGroup.POST("/newest_seq", m.GetSeq)
		chatGroup.POST("/send_msg", m.SendMsg)
		chatGroup.POST("/pull_msg_by_seq", m.PullMsgBySeqs)
		chatGroup.POST("/del_msg", m.DelMsg)
		chatGroup.POST("/del_super_group_msg", m.DelSuperGroupMsg)
		chatGroup.POST("/clear_msg", m.ClearMsg)

		chatGroup.POST("/send_msg", m.ManagementSendMsg)
		chatGroup.POST("/batch_send_msg", m.ManagementBatchSendMsg)
		chatGroup.POST("/check_msg_is_send_success", m.CheckMsgIsSendSuccess)
		chatGroup.POST("/get_users_online_status", m.GetUsersOnlineStatus)
		chatGroup.POST("/account_check", m.AccountCheck)
		//chatGroup.POST("/set_message_reaction_extensions", msg.SetMessageReactionExtensions)
		//chatGroup.POST("/get_message_list_reaction_extensions", msg.GetMessageListReactionExtensions)
		//chatGroup.POST("/add_message_reaction_extensions", msg.AddMessageReactionExtensions)
		//chatGroup.POST("/delete_message_reaction_extensions", msg.DeleteMessageReactionExtensions)
	}
	////Conversation
	conversationGroup := r.Group("/conversation")
	{ //1
		c := NewConversation(zk)
		conversationGroup.POST("/get_all_conversations", c.GetAllConversations)
		conversationGroup.POST("/get_conversation", c.GetConversation)
		conversationGroup.POST("/get_conversations", c.GetConversations)
		conversationGroup.POST("/set_conversation", c.SetConversation)
		conversationGroup.POST("/batch_set_conversation", c.BatchSetConversations)
		conversationGroup.POST("/set_recv_msg_opt", c.SetRecvMsgOpt)
		conversationGroup.POST("/modify_conversation_field", c.ModifyConversationField)
	}
	return r
}
