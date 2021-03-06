package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"crypto/md5"
	"encoding/hex"
)

type UserModel struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Pass       string `json:"pass"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	UserId      string `json:"userId"`
}

func (self *UserModel) UserIsExist() bool {

	mORM := orm.NewOrm()

	return mORM.QueryTable(self).Filter("name", self.Name).Exist()
}

func (self *UserModel) UserIdIsExist() bool {

	mORM := orm.NewOrm()

	return mORM.QueryTable(self).Filter("userId", self.UserId).Exist()
}

func (self *UserModel) InsetUser() bool {

	self.CreateTime = time.Now().Unix()
	self.UpdateTime = self.CreateTime
	mORM := orm.NewOrm()
	self.UserId = getMD5String(self.Name)
	_, err := mORM.Insert(self)
	if err != nil {
		beego.Error(err)
	}
	return err == nil
}

func (self *UserModel) GetUser() bool {

	mORM := orm.NewOrm()

	return mORM.Read(self, "name") == nil
}

func (self *UserModel) ResetPassword() bool {

	fmt.Println(self)
	mORM := orm.NewOrm()
	_, err := mORM.Raw("UPDATE user_model SET pass = ?, update_time = ? WHERE user_id = ?", self.Pass, time.Now().Unix(), self.UserId).Exec()
	return err == nil
}

func getMD5String (rawString string) string {

	h := md5.New()
	h.Write([]byte(rawString))
	return hex.EncodeToString(h.Sum(nil))
}

