package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/marcsj/ocaptchas/util"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type ChallengeImagesRepo interface {
	GetChallengeImages(
		number int, label string) (images [][]byte, answer string, err error)
	ScanForChallenges() error
}

type challengeImagesRepo struct {
	db     *gorm.DB
	folder string
}

func NewChallengeImagesRepo(db *gorm.DB, imageFolder string) (ChallengeImagesRepo, error) {
	err := db.AutoMigrate(ImageChallenge{}).Error
	if err != nil {
		return nil, err
	}
	return &challengeImagesRepo{
		db:     db,
		folder: imageFolder,
	}, nil
}

type ImageChallenge struct {
	gorm.Model
	Label string
	Path  string
}

func (r challengeImagesRepo) GetChallengeImages(
	number int, label string) (images [][]byte, answer string, err error) {
	challenges, answer, err := r.getChallenges(number, label)
	if err != nil {
		return
	}
	images = make([][]byte, 0)
	for _, challenge := range challenges {
		img, _, err := util.ReadImage(challenge.Path)
		if err != nil {
			return
		}
		imgData, err := util.ConvertImage(img)
		if err != nil {
			return
		}
		images = append(images, imgData)
	}
	return
}

func (r challengeImagesRepo) getChallenges(
	number int, label string) (
	challenges []*ImageChallenge, answer string, err error) {
	correct, incorrect, err := r.getLabelChallenges(number, label)
	if err != nil {
		return
	}
	correctIDs := make([]uint, 0)
	var answers []string
	for _, correct := range correct {
		correctIDs = append(correctIDs, correct.ID)
	}
	challenges = append(correct, incorrect...)
	rand.Shuffle(len(challenges), func(i, j int) {
		challenges[i], challenges[j] = challenges[j], challenges[i]
	})
	for i, ch := range challenges {
		util.ContainsUInt(correctIDs, ch.ID)
		answers = append(answers, strconv.Itoa(i))
	}
	answer = strings.Join(answers, ",")
	return
}

func (r challengeImagesRepo) getLabelChallenges(
	number int, label string) (
	correct []*ImageChallenge, incorrect []*ImageChallenge, err error) {
	numberCorrect := rand.Intn(number - 2)
	if numberCorrect < 0 {
		numberCorrect = 1
	}
	correct = make([]*ImageChallenge, numberCorrect)
	incorrect = make([]*ImageChallenge, number-numberCorrect)
	err = r.db.
		Limit(numberCorrect).
		Find(correct, "label = ?", label).
		Order(gorm.Expr("random()")).
		Error
	if err != nil {
		return
	}
	err = r.db.
		Limit(number-numberCorrect).
		Find(incorrect, "label <> ?", label).
		Order(gorm.Expr("random()")).
		Error
	return
}

func (r challengeImagesRepo) ScanForChallenges() error {
	labels, err := r.getLabels()
	if err != nil {
		return err
	}
	for _, label := range labels {
		err = r.DeleteWithLabel(label)
		if err != nil {
			return err
		}
		images, err := r.getImages(label)
		if err != nil {
			return err
		}
		for _, img := range images {
			err = r.NewImageEntry(label, img)
			if err != nil {
				log.Println(err)
			}
			continue
		}
	}
	return nil
}

func (r challengeImagesRepo) getLabels() ([]string, error) {
	labels := make([]string, 0)
	f, err := os.Open(r.folder)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			labels = append(labels, file.Name())
		}
	}
	return labels, nil
}

func (r challengeImagesRepo) getImages(label string) ([]string, error) {
	path := fmt.Sprintf("%s/%s", r.folder, label)
	labels := make([]string, 0)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	images := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			filePath := fmt.Sprintf("%s/%s", path, file.Name())
			_, _, err := util.ReadImage(filePath)
			if err != nil {
				log.Println(err)
				continue
			}
			images = append(images, filePath)
		}
	}
	return labels, nil
}

func (r challengeImagesRepo) NewImageEntry(label string, path string) error {
	return r.db.Create(
		&ImageChallenge{
			Label: label,
			Path:  path,
		}).
		Error
}

func (r challengeImagesRepo) DeleteWithLabel(label string) error {
	return r.db.
		Unscoped().
		Delete(&ImageChallenge{}, "label = ?", label).
		Error
}
