package service

import (
	_ "Assignment/common"
	"Assignment/database"
	"Assignment/models"
	"log"
)

type UserService interface {
	CreateUser(a *models.FListrespone) (*models.ReplyResult, error)
	CreateConnection(a *models.FConnectionrequest) (*models.ReplyResult, error)
	GetFriendList(a *models.FListRequest) (*models.FListrespone, error)
	GetCommonFriendList(a *models.CommonFriendRequest) (*models.FListrespone, error)
	CreateSubscribe(a *models.SubscriptionRequest) (*models.ReplyResult, error)
	CreateBlockFriend(a *models.BlockRequest) (*models.ReplyResult, error)
	CreateUpdateReceive(a *models.UpdateEmailRequest) (*models.RespondEmail, error)
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

func (st Storage) CreateSubscribe(a *models.SubscriptionRequest) (*models.ReplyResult, error) {
	response := &models.ReplyResult{}

}
