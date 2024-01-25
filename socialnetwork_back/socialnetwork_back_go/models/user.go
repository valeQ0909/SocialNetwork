package models

import "log"

// 映射字段名
type TableUser struct {
	Id        int64
	Username  string
	Password  string
	Avatar    string
	Signature string
}

// 表名映射
func (u TableUser) TableName() string {
	return "users"
}

// GetTableUserList 获取全部TableUser对象
func GetTableUserList() ([]TableUser, error) {
	tableUser := []TableUser{}
	if err := DB.Find(&tableUser).Error; err != nil {
		log.Println(err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// GetTableUserByUsername 根据username获得TableUser对象
func GetTableUserByUsername(username string) (TableUser, error) {
	tableUser := TableUser{}
	if err := DB.Where("username = ?", username).First(&tableUser).Error; err != nil {
		log.Println("GetTableUserByUsername未查询到用户", err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// GetTableUserById 根据id获得TableUser对象
func GetTableUserById(id int64) (TableUser, error) {
	tableUser := TableUser{}
	if err := DB.Where("id = ?", id).First(&tableUser).Error; err != nil {
		log.Println(err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// InsertTableUser 将tableUser插入表内
func InsertTableUser(tableUser *TableUser) bool {
	if err := DB.Create(&tableUser).Error; err != nil {
		log.Println("InsertTableUser插入用户失败", err.Error())
		return false
	}
	return true
}

func UpdateAvatar(userid int64, filePath string) error {
	err := DB.Model(&TableUser{}).Where("id = ?", userid).Update("avatar", filePath).Error
	if err != nil {
		log.Println("UpdateAvatar更新用户头像失败", err.Error())
		return err
	}
	return nil
}

func UpdatePersonalInfo(userid int64, username string, signature string) error {
	err := DB.Model(&TableUser{}).Where("id = ?", userid).Updates(TableUser{Username: username, Signature: signature}).Error
	if err != nil {
		log.Println("UpdatePersonalInfo更新用户名称和签名失败", err.Error())
		return err
	}
	return nil
}
