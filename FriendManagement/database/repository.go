package database

import (
	"Assignment/models"
)

//Create new user by email
func (db Database) CreateUserByEmail(email string) error {
	query := `INSERT INTO user_profile(email) VALUES ($1);`
	_, err := db.Conn.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}

//create a sender and reciever by requestor
func (db Database) CreateSubcribeFriendByRequestorAndTarget(sender, reciever string) error {
	query := `INSERT INTO subcription(sender, reciever) VALUES ($1,$2);`
	_, err := db.Conn.Exec(query, sender, reciever)
	if err != nil {
		return err
	}
	return nil
}

//create a requestor and target by order
func (db Database) CreateBlockFriendByRequestorAndTarget(requestor, target string) error {
	query := `INSERT INTO block_user(requestor, target) VALUES ($1,$2);`
	_, err := db.Conn.Exec(query, requestor, target)
	if err != nil {
		return err
	}
	return nil
}

// add friend connection between 2 emails 1 and 2
func (db Database) CreateUser(a *models.FConnectionrequest) error {
	firstemail := a.Friends[0]
	secondemail := a.Friends[1]
	query := `INSERT INTO friend (the_first_user,the_second_user) VALUES ($1,$2);`
	_, err := db.Conn.Exec(query, firstemail, secondemail)
	if err != nil {
		return err
	}
	return nil
}

//GET all User in Friend Email
func (db Database) GetAllUser() ([]models.User, error) {
	allUser := []models.User{}
	query := `SELECT * from user_profile;`
	row, err := db.Conn.Query(query)
	if err != nil {
		return allUser, err
	}
	//To print out all User(email)
	for row.Next() {
		var item models.User
		err := row.Scan(&item.Email)
		if err != nil {
			return allUser, err
		}
		allUser = append(allUser, item)
	}
	return allUser, nil
}

//execute to get all Subscriber
func (db Database) GetAllSubscriber(requester string) (*models.User, error) {
	target := &models.User{}
	query := `SELECT sender FROM subcription s WHERE reciever = $1;`
	row, err := db.Conn.Query(query, requester)
	if err != nil {
		return target, err
	}
	for row.Next() {
		var item models.SubscriptionRequest
		err := row.Scan(&item.Requester)
		if err != nil {
			return target, err
		}
		target.Subscription = append(target.Subscription, item.Requester)
	}
	return target, err
}

//execute to Get all Email has been blocked
func (db Database) GetAllBlockEmail(requestor string) (*models.User, error) {
	targetLst := &models.User{}
	query := `SELECT requestor FROM block_user WHERE target =$1;`
	row, err := db.Conn.Query(query, requestor)
	if err != nil {
		return targetLst, err
	}
	for row.Next() {
		var item models.BlockRequest
		err := row.Scan(&item.Requester)
		if err != nil {
			return targetLst, err
		}
		targetLst.Blocked = append(targetLst.Blocked, item.Requester)
	}
	return targetLst, nil
}

//execute to GET all Target(Email is blocked)
func (db Database) GetTargetBlockByRequest(requestor, target string) (int, error) {
	var countBlocked int
	query := `SELECT COUNT(*) FROM block_user WHERE requestor= $1 AND target= $2;`
	row := db.Conn.QueryRow(query, requestor, target)
	err := row.Scan(&countBlocked)
	if err != nil {
		return countBlocked, err
	}
	return countBlocked, nil
}

//execute to GET all Reciever when sender is subscribed
func (db Database) GetRecieverSubscribeBySender(sender, reciever string) (int, error) {
	var countSubscribe int
	query := `SELECT COUNT(*) FROM subcription WHERE sender =$1 AND reciever =$2;`
	row := db.Conn.QueryRow(query, sender, reciever)
	err := row.Scan(&countSubscribe)
	if err != nil {
		return countSubscribe, err
	}
	return countSubscribe, nil
}

//execute to GET Friends
func (db Database) GetFriendListByRequest(request *models.FConnectionrequest) (int, error) {
	var countFriend int
	query := `SELECT COUNT(*) FROM friend 
				WHERE the_first_user =$1 AND the_second_user = $2
				OR the_first_user = $2 AND the_second_user = $1;`
	row := db.Conn.QueryRow(query, request.Friends[0], request.Friends[1])
	err := row.Scan(&countFriend)
	if err != nil {
		return countFriend, err
	}
	return countFriend, nil
}

//execute to GET Friend Email has been blocked
func (db Database) GetBlockListByRequest(request *models.FConnectionrequest) (int, error) {
	var countEmailBlocked int
	query := `SELECT COUNT(*) FROM block_user
				WHERE requestor = $1 AND target = $2
				OR requestor = $2 AND target = $1;`
	row := db.Conn.QueryRow(query, request.Friends[0], request.Friends[1])
	err := row.Scan(&countEmailBlocked)
	if err != nil {
		return countEmailBlocked, err
	}
	return countEmailBlocked, nil
}

//execute to GET USER emails by Requester
func (db Database) GetUserByRequest(firstEmail, secondEmail string) (int, error) {
	var countUser int
	query := `SELECT COUNT(*) FROM user_profile 
			WHERE email in($1,$2);`
	row := db.Conn.QueryRow(query, firstEmail, secondEmail)
	err := row.Scan(&countUser)
	if err != nil {
		return countUser, err
	}
	return countUser, nil
}

//execute to GET User by Emai;
func (db Database) GetUserByEmail(email string) (int, error) {
	var countUserEmail int
	query := `SELECT COUNT(*) FROM user_profile WHERE email = $1;`
	row := db.Conn.QueryRow(query, email)
	err := row.Scan(&countUserEmail)
	if err != nil {
		return countUserEmail, err
	}
	return countUserEmail, nil
}

//execute to GET all Friend Email
func (db Database) GetFriendListByEmail(email string) (*models.User, error) {
	var friendLst = &models.User{}
	query := `SELECT the_second_user, id FROM friend WHERE the_first_user =$1 UNION 
			SELECT the_first_user, id FROM friend WHERE the_second_user = $1
			ORDER BY id;`
	row, err := db.Conn.Query(query, email)
	if err != nil {
		return friendLst, err
	}
	var id int
	for row.Next() {
		var item models.User
		err := row.Scan(&item.Email, &id)
		if err != nil {
			return friendLst, err
		}
		friendLst.Friends = append(friendLst.Friends, item.Email)
	}
	return friendLst, nil
}
