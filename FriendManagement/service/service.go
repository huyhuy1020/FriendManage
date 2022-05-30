package service

import (
	"Assignment/database"
	"Assignment/models"
)

type UserService interface {
	CreateUser(a *models.FListrespone) (*models.ReplyResult, error)
	CreateConnection(a *models.FConnectionrequest) (*models.ReplyResult, error)
	GetFriendList(a *models.FListRequest) (*models.FListrespone, error)
	GetCommonFriendList(a *models.CommonFriendRequest) (*models.FListrespone, error)
	CreateSubscribe(a *models.SubscriptionRequest) (*models.ReplyResult, error)
	CreateBlockFriend(a *models.BlockRequest) (*models.ReplyResult, error)
	CreateUpdateRecieve(a *models.UpdateEmailRequest) (*models.RespondEmail, error)
}

type Storage struct {
	Db database.Database
}

func (st Storage) CreateUser(a *models.FListrespone) (*models.ReplyResult, error) {
	respone := &models.ReplyResult{}
	if err := st.Db.Cr
}
