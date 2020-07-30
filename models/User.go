package models

import (
	"errors"
	"html"
	"time"
	"strings"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User type
type User struct {
	ID			string		`gorm:"primary_key;" json:"id"`
	Username 	string		`gorm:"size:32;not null;unique" json:"username"`
	Email		string		`gorm:"size:100;not null;unique" json:"email"`
	Password	string		`gorm:"size:100;not null;" json:"password"`
	Role		int8		`gorm:"not null;" json:"role"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Hash password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword when logging in
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave Hash password before saving data
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Validate user data
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil
	}
}

// Prepare new data 
func (u *User) Prepare() {
	u.ID = uuid.New().String()
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// SaveUser to database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// FindUsers in DB limit by 10
func (u *User) FindUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Order("updated_at DESC").Limit(10).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindUserByID by passing id through param
func (u *User) FindUserByID(db *gorm.DB, uid string) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User tidak ditemukan")
	}
	return u, err
}

// UpdateUserByID by passing id through param
func (u *User) UpdateUserByID(db *gorm.DB, uid string) (*User, error) {
	var err error

	// Hash the password first
	bsError := u.BeforeSave()
	if bsError != nil {
		log.Fatal(bsError)
	}

	// Update the user
	db = db.Debug().Model(&User{}).Where("id =? ", uid).Take(&User{}).UpdateColumns(
		map[string]interface{} {
			"password":		u.Password,
			"username": 	u.Username,
			"email":		u.Email,
			"updated_at":	time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// Display updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// DeleteUserByID by passing id through param
func (u *User) DeleteUserByID(db *gorm.DB, uid string) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}