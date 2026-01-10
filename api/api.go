package api

import (
	"net/http"

	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	querier repo.Querier
}

func NewMessageHandler(querier repo.Querier) *MessageHandler {
	return &MessageHandler{
		querier: querier,
	}
}

func (h *MessageHandler) WireHttpHandler() http.Handler {

	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, _ any) {
		c.String(http.StatusInternalServerError, "Internal Server Error: panic")
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.POST("/message", h.handleCreateMessage)
	r.POST("/user", h.handleAddUser)
	r.POST("/group", h.handleCreateGroup)
	r.POST("/thread", h.handleStartThread)
	r.POST("/group/member", h.handleAddGroupMember)
	r.POST("/thread/member", h.handleAddThreadMember)
	r.POST("/dm", h.handleStartDm)

	r.GET("/message/:message_id", h.handleGetMessage)
	r.GET("/message/dm/:dm_id", h.handleGetDmMessage)
	r.GET("/message/group/:group_id", h.handleGetGroupMessage)
	r.GET("/message/thread/:thread_id", h.handleGetThreadMessage)
	r.GET("/user/group/:user_id", h.handleGetUserGroup)
	r.GET("/user/thread/:user_id", h.handleGetUserThread)
	r.GET("/user/group/creator/:user_id", h.handleGetGroupsUserCreated)
	r.GET("/user/thread/creator/:user_id", h.handleGetThreadsUserStarted)
	r.GET("/group/thread/:group_id", h.handleGetGroupThread)
	r.GET("/group/group-list", h.handleGetListOfGroups)
	
	r.DELETE("/message/:message_id", h.handleDeleteMessageByID)
	r.DELETE("/thread/:thread_id/:created-by", h.handleDeleteThreadByID)
	r.DELETE("/group/:group_id/:created-by", h.handleDeleteGroupByID)


	return r
}

func (h *MessageHandler) handleCreateMessage(c *gin.Context) {
	var req repo.CreateMessageParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.CreateMessage(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleAddUser(c *gin.Context) {
	var req repo.AddUserParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.AddUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleCreateGroup(c *gin.Context) {
	var req repo.CreateGroupParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.CreateGroup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleStartThread(c *gin.Context) {
	var req repo.StartThreadParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.StartThread(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleAddGroupMember(c *gin.Context) {
	var req repo.AddGroupMemberParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.AddGroupMember(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleAddThreadMember(c *gin.Context) {
	var req repo.AddThreadMemberParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.AddThreadMember(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleStartDm(c *gin.Context) {
	var req repo.StartDmParams

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.querier.StartDm(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetMessage(c *gin.Context) {
	id := c.Param("message_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message id is required"})
		return
	}

	message, err := h.querier.GetMessageByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetDmMessage(c *gin.Context) {
	id := c.Param("dm_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dm id is required"})
		return
	}

	message, err := h.querier.GetDmMessages(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetGroupMessage(c *gin.Context) {
	id := c.Param("group_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group id is required"})
		return
	}

	message, err := h.querier.GetGroupMessages(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetThreadMessage(c *gin.Context) {
	id := c.Param("thread_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thread id is required"})
		return
	}

	message, err := h.querier.GetThreadMessages(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetUserGroup(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	message, err := h.querier.GetUserGroups(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetUserThread(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	message, err := h.querier.GetUserThreads(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetGroupsUserCreated(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	message, err := h.querier.GetGroupsUserCreated(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetThreadsUserStarted(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	message, err := h.querier.GetThreadsUserStarted(c, &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetGroupThread(c *gin.Context) {
	id := c.Param("group_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group id is required"})
		return
	}

	message, err := h.querier.GetGroupThreads(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetListOfGroups(c *gin.Context) {

	message, err := h.querier.GetListOfGroups(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleDeleteMessageByID(c *gin.Context) {
	id := c.Param("message_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message id is required"})
		return
	}
	err := h.querier.DeleteMessageByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


}

func (h *MessageHandler) handleDeleteThreadByID(c *gin.Context) {
	id := c.Param("thread_id")
	createdBy := c.Param("created-by")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thread id is required"})
		return
	}
	
	if createdBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "created by is required"})
		return
	}
	
	err := h.querier.DeleteThreadByID(c, repo.DeleteThreadByIDParams{
	ThreadID: id,
	CreatedBy: &createdBy,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


}

func (h *MessageHandler) handleDeleteGroupByID(c *gin.Context) {
	id := c.Param("group_id")
	createdBy := c.Param("created-by")
	
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group id is required"})
		return
	}

	if createdBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "created by is required"})
		return
	}

	err := h.querier.DeleteGroupByID(c, repo.DeleteGroupByIDParams{
		GpID: id,
		CreatedBy: createdBy,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


}
