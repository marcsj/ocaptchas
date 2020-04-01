package repo

import "github.com/jinzhu/gorm"

type SessionRepo interface {
	CreateSession(uuid string, sessionType SessionType, answer string) error
	GetSession(uuid string) (*Session, error)
	DeleteSession(uuid string) error
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) (SessionRepo, error) {
	err := db.AutoMigrate(Session{}).Error
	if err != nil {
		return nil, err
	}
	return &sessionRepo{
		db: db,
	}, nil
}

type Session struct {
	gorm.Model
	UUID        string
	SessionType SessionType
	Answer      string
}

type SessionType int

var (
	SessionType_Images       SessionType = 0
	SessionType_Alphanumeric SessionType = 1
	SessionType_Questions    SessionType = 2
)

func (r sessionRepo) CreateSession(
	uuid string, sessionType SessionType, answer string) error {
	return r.db.Create(
		Session{
			UUID:        uuid,
			SessionType: sessionType,
			Answer:      answer,
		}).Error
}

func (r sessionRepo) GetSession(uuid string) (*Session, error) {
	retrievedSession := &Session{}
	r.db.First(retrievedSession, "uuid = ?", uuid)
	return retrievedSession, r.db.Error
}

func (r sessionRepo) DeleteSession(uuid string) error {
	return r.db.Delete(&Session{}, "uuid = ?", uuid).Error
}
