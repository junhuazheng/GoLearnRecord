package model

import (
	"encoding/json"
	"fmt"

	"github.com.garyburd/redigo/redis"
)

//UserDao instance
var CurrentUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

//initialize an instance of the UserDao structure
func InitUserDao(pool *redis.Pool) (currentUserDao *UserDao) {
	currentUserDao = &UserDao{pool: pool}
	return
}

func idIncr(conn redis.Conn) (id int, err error) {
	res, err := conn.Do("incr", "users_id")
	id = int(res.(int64))
	if err != nil {
		fmt.Printf("id incr error: %v\n", err)
		return
	}
	return
}

//Get user infomation based on id
//Get success return user infomation, err nil
//Get failed return err , user is nil
func (this *UserDao) GetUserById(id int) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		err = ERROR_USER_DOES_NOT_EXIST
		return
	}

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Printf("Unmarshal user info error: %v\n", err)
		return
	}
	return
}

//Get user infomation based on username
func (this *UserDao) GetUserByUserName(userName string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("hget", "users", userName))
	if err != nil {
		err = ERROR_USER_DOES_NOT_EXIST
		return
	}
	
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Printf("Unmarshal user info error: %v\n", err)
		return
	}
	return
}

//Registered users
//the user name cannot be repeated
func (this *UserDao) Register(userNaem, password, passwordConfirm string) (user User, err error) {
	//determine if the password is correct
	if password != passwordConfirm {
		err = ERROR_PASSWORD_DOES_NOT_MATCH
		return
	}

	//make sure the user name does not duplicate
	user, err = this.GetUserByUserName(userName)
	if err != nil {
		fmt.Println("User already exists!\n")
		err = ERROR_USER_ALEADY_EXISTS
		return
	}

	conn := this.pool.Get()
	defer conn.Close()

	//id auto-increment 1, id as the next user
	id, err := idIncr(conn)
	if err != nil {
		return
	}

	user = User{ID: id, Name: userName, Password: password}
	info, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json marshal error: %v\n", err)
		return
	}
	_, err = conn.Do("hset", "users", userNaem, info)
	if err != nil {
		fmt.Printf("ser user to redis error: %v\n", err)
		return
	}
	return
}

func (this *UserDao) Login(userName, password string) (user User, err error) {
	user, err = this.GetUserByUserName(userNaem)
	if err != nil {
		fmt.Printf("get user by name error: %v\n", err)
		return
	}

	if user.Password != password {
		err = ERROR_USER_PWD
		return
	}
	return
}

