package dao

import (
	"../framework"
	"../model"
	"fmt"
)

type UserDao struct {

}

func (p *UserDao) Insert(user *model.User) int64 {
	result, err := framework.DB.Exec("INSERT INTO user(`username`,`password`,`create_time`) value(?,?,?)", user.Username, user.Password, user.CreateTime)
	if err!=nil {
		fmt.Println("Insert error")
	}
	id,err:=result.LastInsertId()
	if err!=nil {
		fmt.Println("Insert error")
		return 0
	}
	return id
}

func (p *UserDao) SelectUserByName(username string) []model.User{
	rows,err:=framework.DB.Query("Select * from user where username = ?",username)
	if err !=nil {
		fmt.Println("Select user by name error")
		return nil
	}
	var users []model.User
	for rows.Next()  {
		var user model.User
		err := rows.Scan(&user.ID,&user.Username,&user.Password,&user.CreateTime)
		if err!=nil {
			fmt.Println("Select user by name error")
			continue
		}
		users=append(users,user)
	}
	rows.Close()
	return users
}

func (p *UserDao) SelectAllUser() []model.User {
	rows,err:=framework.DB.Query("Select * from user")
	if err!=nil {
		fmt.Println("SelectAllUser error")
	}
	var users []model.User
	for rows.Next()  {
		var user model.User
		err :=rows.Scan(user.ID, &user.Username, &user.Password, &user.CreateTime)
		if err!=nil {
			fmt.Println("SelectAllUser error")
			continue
		}
		users=append(users,user)
	}
	rows.Close()
	return users
}