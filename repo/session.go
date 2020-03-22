package repo

import "github.com/jinzhu/gorm"

type session struct {
	gorm.Model
	uuid string
	sessionType string
}

type SessionRepo interface {
	CreateSession(uuid string, sessionType string) error
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) (SessionRepo, error) {
	err := db.AutoMigrate().Error
	if err != nil {
		return nil, err
	}
	return &sessionRepo{
		db: db,
	}, nil
}

func (r sessionRepo) CreateSession(uuid string, sessionType string) error {
	return r.db.Create(
		session{
		uuid: uuid,
		sessionType: sessionType,
		}).Error
}

func (r sessionRepo) GetSession(uuid string) (*session, error) {
	retrievedSession := &session{}
	r.db.First(retrievedSession, "uuid = ?", uuid)
	return retrievedSession, r.db.Error
}

func (r sessionRepo) DeleteSession(uuid string) error {
	return r.db.Delete(&session{}, "uuid = ?", uuid).Error
}
