package models

import (
	"fmt"
	"log"
	"net/http"
)

//F: Friend
type User struct {
	Email        string   `json: email`
	Friends      []string `json: friends`
	Subscription []string `json: subscription`
	Blocked      []string `json: blocked`
}

type ListFriend struct {
	Users []User `json: users`
}

type ReplyResult struct {
	Success bool `json: success`
}

type FConnectionrequest struct {
	Friends []string `json: friends`
}

type FListRequest struct {
	Email string `json:email`
}

type EmailRequest struct {
	Email string `json:email`
}

type CommonFriendRequest struct {
	Friends []string `json:friends`
}

type FListrespone struct {
	Success bool     `json: success`
	Friends []string `json: friends`
	Count   int      `json: count`
}

type SubscriptionRequest struct {
	Requester string `json: requester`
	Target    string `json:target`
}

type BlockRequest struct {
	Requester string `json: requester`
	Target    string `json: target`
}

type UpdateEmailRequest struct {
	Sender string `json: sender`
	Text   string `json: text`
}

type RespondEmail struct {
	Success    bool     `json: success`
	Recipients []string `json: recipients`
}

//check for an empty list of friend. if it is empty, the program will be warning
func (friend *FConnectionrequest) Bind(r *http.Request) error {
	userEmailOne := friend.Friends[0]
	userEmailTwo := friend.Friends[1]
	if userEmailOne == "" || userEmailTwo == "" {
		return fmt.Errorf("friend is essential part")
	}
	if userEmailOne == userEmailTwo {
		log.Print("can't connect by itself")
		return fmt.Errorf("can't connect itself")
	}
	return nil
}

//check email for non-empty
func (email *EmailRequest) Bind(r *http.Request) error {
	if email.Email == "" {
		log.Print("email is essential part")
		return fmt.Errorf("email is essential part")
	}
	return nil
}

//check for non-email and target
func (sub *SubscriptionRequest) Bind(r *http.Request) error {
	requestor := sub.Requester
	target := sub.Target
	if requestor == "" || target == "" {
		log.Print("email cannot empty.")
		return fmt.Errorf("email cannot empty.")
	}
	if requestor == target {
		log.Print("can't not subcribe by itself")
		return fmt.Errorf("can't not subcribe by itself")
	}
	return nil
}
