package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model `json:"-"`
	Username string `gorm:not null;unique" json:"username"`
	Password string `gorm:not null" json:"-"`
	UUID 	 string `gorm:not null;unique" json:"uuid"`
}

type UserManager struct {
	db *DB
}

func NewUserManager(db *DB) (*UserManager, error) {
	db.AutoMigrate(&User{})
	usermgr := UserManager{}
	usermgr.db = db
	return &usermgr, nil
}

func (state *UserManager) HasUser(username string) bool {
	if err := state.db.Where("username=?", username).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

func (state *UserManager) FindUser(username string) *User {
	user := User{}
	state.db.Where("username=?", username).Find(&user)
	return &user
}

func (state *UserManager) FindUserByUUID(uuid string) *User {
	user := User{}
	state.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

func (state *UserManager) AddUser(username, password string) *User {
	passwordHash := state.HashPassword(username, password)
	guid, _ := uuid.NewV4()
	user := &User{
		Username: username,
		Password: passwordHash,
		UUID: guid.String(),
	}
	state.db.Create(&user)
	return user
}

func (state *UserManager) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unseccessful")
	}
	reuturn string(hash)
}

func (state *UserManager) CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}