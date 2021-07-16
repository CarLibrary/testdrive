package model

import "gorm.io/gorm"

type TestDrive struct {
	gorm.Model
	//userid
	Userid uint `gorm:"not null"`
	//品牌
	Carband string `gorm:"not null;size:255"`
	//车系
	CarSeries string `gorm:"not null;size:255"`
	//城市
	City string `gorm:"not null;size:10"`
	//姓名
	Username string `gorm:"not null;size:5"`
	//手机号
	PhoneNumber string `gorm:"not null;size:11"`
	//status
	Status string `gorm:"not null;default:已提交;size:3"`
}

// TestDrive 试驾
func (t *TestDrive)TestDrive() (testdrive *TestDrive,err error) {
	if err:=db.Table("test_drives").Create(t).Error;err!=nil{
		return new(TestDrive),err
	}
	return t,nil
}

//查看试驾
func FindMytest(userid uint) (list *[]TestDrive,err error)  {
	if err=db.Table("test_drive").Where("userid = ?",userid).Find(list).Error;err!=nil{
		return new([]TestDrive), err
	}
	return list, nil
}

