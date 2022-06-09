package service

import (
	"Assignment/common"
	"Assignment/database"
	"Assignment/models"
	"log"
	"strings"
)

type UserService interface {
	CreateUser(a *models.FListRequest) (*models.ReplyResult, error)
	CreateConnection(a *models.FConnectionrequest) (*models.ReplyResult, error)
	GetFriendList(a *models.FListRequest) (*models.FListrespone, error)
	GetCommonFriendList(a *models.CommonFriendRequest) (*models.FListrespone, error)
	CreateSubscribe(a *models.SubscriptionRequest) (*models.ReplyResult, error)
	CreateBlockFriend(a *models.BlockRequest) (*models.ReplyResult, error)
	CreateUpdateReceive(a *models.EmailRequest) (*models.RespondEmail, error)
}

type Storage struct {
	Db database.Database
}

// CreateUser Use for dependency injection it is call high-level controller
func (st Storage) CreateUser(a *models.FListRequest) (*models.ReplyResult, error) {
	response := &models.ReplyResult{}
	err := st.Db.CreateUserByEmail(a.Email)
	if err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

// CreateConnection Connection of User by request, friend by request, friend by 2 emails
func (st Storage) CreateConnection(a *models.FConnectionrequest) (*models.ReplyResult, error) {
	response := &models.ReplyResult{}
	countUser, err := st.Db.GetUserByRequest(a.Friends[0], a.Friends[1])
	if err != nil {
		return response, err
	}
	if countUser <= 1 {
		log.Printf("Email does not exsist")
		return response, err
	}
	countFriend, err := st.Db.GetFriendListByRequest(a)
	if err != nil {
		return response, err
	}
	if countFriend == 1 {
		log.Printf("Cannot add Friend by Request. Because it is exsting an email ")
		return response, nil
	}
	countBlocked, err := st.Db.GetBlockListByRequest(a)
	if countBlocked == 1 || countBlocked == 2 {
		log.Printf("Cannot block by request. Because it will be empty when it blocked")
		return response, err
	}

	//add Friend by 2 emails
	if err = st.Db.CreateUser(a); err != nil {
		if err != nil {
			response.Success = false
			return response, err
		}
	}
	response.Success = true
	return response, nil
}

// GetFriendList acquire list of friends
func (st Storage) GetFriendList(a *models.FListRequest) (*models.FListrespone, error) {
	response := &models.FListrespone{}
	rep, err := st.Db.GetFriendListByEmail(a.Email)
	if err != nil {
		return response, err
	}
	response.Success = true
	response.Friends = rep.Friends
	response.Count = len(rep.Friends)
	return response, nil
}

// GetCommonFriendList execute to get friend who is common with everyone
func (st Storage) GetCommonFriendList(a *models.CommonFriendRequest) (*models.FListrespone, error) {
	response := &models.FListrespone{}
	Emailone, err := st.Db.GetFriendListByEmail(a.Friends[0])
	if err != nil {
		return response, err
	}
	Emailtwo, err := st.Db.GetFriendListByEmail(a.Friends[1])
	if err != nil {
		return response, err
	}
	var commonF []string
	for _, friendA := range Emailone.Friends {
		for _, friendB := range Emailtwo.Friends {
			if friendA == friendB {
				commonF = append(commonF, friendA)
			}
		}
	}
	response.Success = true
	response.Friends = commonF
	response.Count = len(commonF)
	return response, nil
}

// CreateSubscribe to get email by users. we check if subscribe is existed or not
func (st Storage) CreateSubscribe(a *models.SubscriptionRequest) (*models.ReplyResult, error) {
	response := &models.ReplyResult{}
	countUser, err := st.Db.GetUserByRequest(a.Requester, a.Target)
	if err != nil {
		return response, err
	}
	if countUser <= 1 {
		log.Printf("Email is not enough for subcribbing")
		return response, err
	}
	countSubscribe, err := st.Db.GetRecieverSubscribeBySender(a.Requester, a.Target)
	if err != nil {
		return response, err
	}
	if countSubscribe == 1 {
		log.Printf("%s Added %s to subcribe", a.Target, a.Requester)
		return response, err
	}
	if err := st.Db.CreateSubcribeFriendByRequestorAndTarget(a.Target, a.Requester); err != nil {
		return response, err
	}
	response.Success = true
	return response, err
}

// CreateBlockFriend setting up and checking conditions to block friends
func (st Storage) CreateBlockFriend(a *models.BlockRequest) (*models.ReplyResult, error) {
	response := &models.ReplyResult{}
	countUser, err := st.Db.GetUserByRequest(a.Requester, a.Target)
	if err != nil {
		return response, err
	}
	if countUser <= 1 {
		log.Printf("Email is not enough to get block")
		return response, err
	}
	countBlocked, err := st.Db.GetTargetBlockByRequest(a.Requester, a.Target)
	if err != nil {
		return response, err
	}
	if countBlocked == 1 {
		log.Printf("can't block, because %s block %s ", a.Requester, a.Target)
		return response, err
	}
	if err := st.Db.CreateBlockFriendByRequestorAndTarget(a.Target, a.Requester); err != nil {
		return response, err
	}
	response.Success = true
	return response, err
}

// CreateUpdateReceive to update data when they blocked and added friends
func (st Storage) CreateUpdateReceive(a *models.EmailRequest) (*models.RespondEmail, error) {
	response := &models.RespondEmail{}
	countUser, err := st.Db.GetUserByEmail(a.Sender)
	if err != nil {
		return response, err
	}
	if countUser == 0 {
		log.Printf("Email is not existing")
		return response, err
	}
	blockLst, err := st.Db.GetAllBlockEmail(a.Sender)
	if err != nil {
		return response, err
	}
	allUser, err := st.Db.GetAllUser()
	if err != nil {
		return response, err
	}
	friendlst, err := st.Db.GetFriendListByEmail(a.Sender)
	if err != nil {
		return response, err
	}
	subscribeLst, err := st.Db.GetAllSubscriber(a.Sender)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, user := range allUser {
		boolBlock := common.CheckListExisting(blockLst.Blocked, user.Email)
		if !boolBlock {
			boolFriend := common.CheckListExisting(friendlst.Friends, user.Email)
			boolSubscribe := common.CheckListExisting(subscribeLst.Subscription, user.Email)
			boolMention := strings.Contains(a.Text, user.Email)
			if boolFriend || boolSubscribe || boolMention {
				list = append(list, user.Email)
			}
		}
	}
	response.Success = true
	response.Recipients = list
	return response, nil
}
