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

func (db Database) CheckUserExist(userID int64) (bool, error) {
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

func (db Database) DeleteUser(userID int64) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", userID).Delete(&BaleBot{}).Error
	if err != nil {
		return fmt.Errorf("error in delete user, error: %v", err)
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

func (db Database) SetUserPhoneNumber(userID int64, userNumber string, newStat string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"phone_number": userNumber, "stat": newStat}).Error
	if err != nil {
		return fmt.Errorf("error in set user number, error: %v", err)
	}
	return nil
}

func (db Database) SetUserPassword(userID int64, password string, newState string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"password": password, "stat": newState}).Error
	if err != nil {
		return fmt.Errorf("error in set user password, error: %v", err)
	}
	return nil
}

func (db Database) SetUserCaptcha(userID int64, captcha string, newState string) error {
	model := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID))
	update := model.Updates(map[string]interface{}{"captcha": captcha, "stat": newState})
	if newState == "login" {
		update = model.Update("entered_time", time.Now())
	}
	if update.Error != nil {
		return fmt.Errorf("error in set user password, error: %v", update.Error)
	}
	return nil
}

func (db Database) GetEnteredTime(userID int64) (*time.Time, error) {
	var userInfo BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", userID).Find(&userInfo).Error
	if err != nil {
		return nil, fmt.Errorf("error in get user entered time,error: %v", err)
	}
	return userInfo.EnteredTime, nil
}

func (db Database) SetUserSiteCookie(userID int64, cookie string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"site_cookie": cookie}).Error
	if err != nil {
		return fmt.Errorf("error in set user password, error: %v", err)
	}
	return nil
}

func (db Database) GetUserSiteCookieAndVerificationToken(userID int64) (string, string, error) {
	var user BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Find(&user).Error
	if err != nil {
		return "", "", fmt.Errorf("error in get user site cookie, error: %v", err)
	}
	return user.SiteCookie, user.VerificationToken, nil
}

func (db Database) SetUserEnteredCodeByPhoneMessage(userID int64, enteredCode string) error {
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

func (db Database) SetUserNationalCode(userID int64, userNumber string, newState string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"national_code": userNumber, "stat": newState}).Error
	if err != nil {
		return fmt.Errorf("error in set national code, error: %v", err)
	}
	return nil
}

func (db Database) SetUserBirthDate(userID int64, BirthDate string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"birth_date": BirthDate, "stat": "enterJobTitle"}).Error
	if err != nil {
		return fmt.Errorf("error in set birth date, error: %v", err)
	}
	return nil
}

func (db Database) SetUserJobId(userID int64, jobId string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"job_id": jobId, "stat": "login"}).Error
	if err != nil {
		return fmt.Errorf("error in set job title, error: %v", err)
	}
	return nil
}

func (db Database) SetJobIdLoginCode(userID int64, jobId string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"job_id_login_code": jobId}).Error
	if err != nil {
		return fmt.Errorf("error in set job id, error: %v", err)
	}
	return nil
}

func (db Database) SetUserVerificationToken(requestVerificationToken string, userID int64, siteCookie string) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"verification_token": requestVerificationToken, "site_cookie": siteCookie}).Error
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

func (db Database) SetRepresentedChildForUser(userID int64, RepresentedChild []byte) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"represented_child": RepresentedChild}).Error
	if err != nil {
		return fmt.Errorf("error in set Represented Child, error: %v", err)
	}
	return nil
}

func (db Database) GetRepresentedChildForUser(userID int64) ([]byte, error) {
	var user BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error in get Represented Child, error: %v", err)
	}
	return user.RepresentedChild, nil
}

func (db Database) DeleteRepresentedChildForUser(userID int64) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Delete("represented_child").Error
	if err != nil {
		return fmt.Errorf("error in get Represented Child, error: %v", err)
	}
	return nil
}

func (db Database) SetCurrentChildForUser(userID int64, currentChild []byte) error {
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Updates(map[string]interface{}{"current_child_info": currentChild}).Error
	if err != nil {
		return fmt.Errorf("error in set Represented Child, error: %v", err)
	}
	return nil
}

func (db Database) GetCurrentChildForUser(userID int64) ([]byte, error) {
	var user BaleBot
	err := db.gorm.Model(&BaleBot{}).Where("id = ?", uint(userID)).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error in get Represented Child, error: %v", err)
	}
	return user.CurrentChildInfo, nil
}
