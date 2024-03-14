package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	gorm *gorm.DB
}

func NewDatabase(db *gorm.DB) Database {
	db.AutoMigrate(&BaleBot{})

	return Database{
		gorm: db,
	}
}

func (db Database) CheckUser(userID int64) (bool, error) {
	var check bool
	err := db.gorm.Raw("SELECT EXISTS(SELECT id FROM bale_bots WHERE id = ?)", userID).Scan(&check).Error
	if err != nil {
		return false, err
	}
	return check, nil
}

func (db Database) GetUserInfo(userID int64) (*BaleBot, error) {
	var user BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db Database) AddUser(userID int64) error {
	var user = BaleBot{
		ID:   uint(userID),
		Stat: "enterPhoneNumber",
	}

	err := db.gorm.Model(&BaleBot{}).Create(&user).Error
	if err != nil {
		return fmt.Errorf("error in add user, error: %v", err)
	}
	return nil
}

func (db Database) GetUserState(userID int64) (string, error) {
	var user BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Find(&user).Error
	if err != nil {
		return "", fmt.Errorf("error in get the user id, error: %v", err)
	}
	return user.Stat, nil
}

func (db Database) UpdateStat(userID int64, newStat string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Update("stat", newStat).Error
	if err != nil {
		return fmt.Errorf("error in updade stat, error: %v", err)
	}
	return nil
}

func (db Database) SetUserPhoneNumber(userID int64, userNumber int64) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"phone_number": userNumber, "stat": "enterVerificationCode"}).Error
	if err != nil {
		return fmt.Errorf("error in set user number, error: %v", err)
	}
	return nil
}

func (db Database) SetUserEnteredCodeByPhoneMessage(userID int64, enteredCode int64) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"verification_code": enteredCode, "stat": "enterFirstName"}).Error
	if err != nil {
		return fmt.Errorf("error in set user number, error: %v", err)
	}
	return nil
}

func (db Database) SetUserFirstName(userID int64, firstName string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"first_name": firstName, "stat": "enterLastName"}).Error
	if err != nil {
		return fmt.Errorf("error in set first name, error: %v", err)
	}
	return nil
}

func (db Database) SetUserLastName(userID int64, lastName string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"last_name": lastName, "stat": "Code"}).Error
	if err != nil {
		return fmt.Errorf("error in set last name, error: %v", err)
	}
	return nil
}

func (db Database) SetUserNationalCode(userID int64, userNumber int64) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"national_code": userNumber, "stat": "enterBirthDay"}).Error
	if err != nil {
		return fmt.Errorf("error in set national code, error: %v", err)
	}
	return nil
}

func (db Database) SetUserBirthDate(userID int64, BirthDate time.Time) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"birth_date": BirthDate, "stat": "enterJobTitle"}).Error
	if err != nil {
		return fmt.Errorf("error in set birth date, error: %v", err)
	}
	return nil
}

func (db Database) SetUserJobTitle(userID int64, jobTitle string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"job_title": jobTitle, "stat": "login"}).Error
	if err != nil {
		return fmt.Errorf("error in set job title, error: %v", err)
	}
	return nil
}

func (db Database) SetUserVerificationToken(requestVerificationToken string, userID int64) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"verification_token": requestVerificationToken}).Error
	if err != nil {
		return fmt.Errorf("error in set verification token, error: %v", err)
	}
	return nil
}

func (db Database) GetUserVerificationToken(userID int64) (string, error) {
	var user BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Find(&user).Error
	if err != nil {
		return "", fmt.Errorf("error in set verification token, error: %v", err)
	}
	return user.VerificationToken, nil
}
