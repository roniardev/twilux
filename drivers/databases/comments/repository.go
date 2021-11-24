package comments

import (
	"context"
	"errors"
	"fmt"
	"twilux/business/comments"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentsRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentsRepository {
	return &CommentsRepository{db}
}

func (repo *CommentsRepository) GetAll(snippetId string, ctx context.Context) ([]comments.Domain, error) {
	sav := []Comment{}
	result := repo.db.Where("snippet_id = ?", snippetId).Find(&sav)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}
	fmt.Println("GetAll Repo db")
	fmt.Println(result)
	return ToListDomain(sav), nil

}

func (repo *CommentsRepository) Create(domain comments.Domain, ctx context.Context) (comments.Domain, error) {
	commentsDb := FromDomain(domain)
	commentsDb.Id, _ = nanoid.Nanoid(10)

	err := repo.db.Create(&commentsDb).Error
	if err != nil {
		return comments.Domain{}, err
	}
	errs := repo.db.Preload(clause.Associations).Preload("Comments."+clause.Associations).First(&commentsDb, commentsDb).Error
	if errs != nil {
		return commentsDb.ToDomain(), nil
	}
	return commentsDb.ToDomain(), nil
}

func (repo *CommentsRepository) Update(domain comments.Domain, ctx context.Context) (comments.Domain, error) {
	commentDb := FromDomain(domain)
	// username check
	res := repo.db.Where("username = ? AND snippet_id = ?", commentDb.Username, commentDb.SnippetId).Updates(&commentDb)
	if res.Error != nil {
		return comments.Domain{}, res.Error
	}
	if res.Error == nil {
		fmt.Print("harusnya bisa ")
	}

	return commentDb.ToDomain(), nil
}

// Update deleted_at field to specific snippet by id
func (repo *CommentsRepository) Delete(domain comments.Domain, ctx context.Context) (comments.Domain, error) {
	commentDb := FromDomain(domain)
	isEligible, _ := repo.CheckAuth(commentDb.Username, commentDb.SnippetId, ctx)

	if !isEligible {
		return comments.Domain{}, errors.New("you are not eligible to delete this comment")
	}
	res := repo.db.Where("username = ? AND snippet_id = ?", commentDb.Username, commentDb.SnippetId).Delete(&commentDb)
	if res.Error != nil {
		return comments.Domain{}, res.Error
	}

	return commentDb.ToDomain(), nil
}

//check username is auth or not
func (repo *CommentsRepository) CheckAuth(username string, snippetId string, ctx context.Context) (bool, error) {
	commentDb := Comment{}
	res := repo.db.Where("username = ? AND snippet_id = ?", username, snippetId).First(&commentDb)
	if res.Error != nil {
		return false, res.Error
	}
	if res.Error == nil {
		return true, nil
	}
	return false, nil
}
