package services

import (
	"database/sql"
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/user-service/internal/pkg"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type UserDataService struct {
	repo iUserData
}

type iUserData interface {
	AddData(u *entities.UserData) (*entities.UserData, error)
	GetUserDataById(id int) (*entities.UserData, error)
}

func NewUserDataService(repo iUserData) *UserDataService {
	return &UserDataService{
		repo: repo,
	}
}

func (s *UserDataService) AddData(img io.Reader, userData *entities.UserData) (*entities.UserData, error) {
	if img != nil {
		user, _ := s.repo.GetUserDataById(userData.Userid)

		//upload image to s3
		sess := &session.Session{}

		if os.Getenv("AWS_ACCESS_KEY") != "" && os.Getenv("AWS_SECRET") != "" {
			config := &aws.Config{
				Region: aws.String("eu-central-1"),
				Credentials: credentials.NewStaticCredentials(
					os.Getenv("AWS_ACCESS_KEY"),
					os.Getenv("AWS_SECRET"),
					"",
				),
			}

			sess = session.Must(session.NewSession(config))
		} else {
			sess = session.Must(session.NewSession())
		}

		uploader := s3manager.NewUploader(sess)

		// if user already has avatar uploaded we need to delete it
		if user != nil && user.AvatarName != "" {
			batcher := s3manager.NewBatchDelete(sess)
			object := []s3manager.BatchDeleteObject{
				{
					Object: &s3.DeleteObjectInput{
						Key:    aws.String(user.AvatarName),
						Bucket: aws.String(os.Getenv("AWS_BUCKET")),
					},
				},
			}
			err := batcher.Delete(aws.BackgroundContext(), &s3manager.DeleteObjectsIterator{
				Objects: object,
			})
			if err != nil {
				logrus.Error(err.Error())
			}
		}

		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("AWS_BUCKET")),
			Key:    aws.String(userData.AvatarName),
			Body:   img,
		})

		if err != nil {
			logrus.Error(err)
			return nil, pkg.NewError("couldn't upload image to s3", http.StatusInternalServerError)
		}
	}

	_, err := s.repo.AddData(userData)
	if err != nil {
		logrus.Info(err)
		return nil, pkg.NewError("couldn't write data to db", http.StatusInternalServerError)
	}

	return nil, nil
}

func (s *UserDataService) GetData(userId int) (*entities.UserData, error) {
	data, err := s.repo.GetUserDataById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewError("no data was registered", http.StatusNotFound)
		}
	}

	return data, nil
}
