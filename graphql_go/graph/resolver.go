package graph

import (
	"errors"
	"graphql_go/graph/model"
	"graphql_go/models"

	"context"

	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*models.User
	err := r.DB.Preload("Posts").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:    string(rune(user.ID)),
			Name:  user.Name,
			Email: user.Email,
			Age:   int32(user.Age),
		})
	}
	return result, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user models.User
	if err := r.DB.Preload("Posts").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &model.User{
		ID:    string(rune(user.ID)),
		Name:  user.Name,
		Email: user.Email,
		Age:   int32(user.Age),
	}, nil

}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []*models.Post

	if err := r.DB.Preload("Author").Find(&posts).Error; err != nil {
		return nil, err
	}

	var result []*model.Post
	for _, post := range posts {
		result = append(result, &model.Post{
			ID:      string(rune(post.ID)),
			Title:   post.Title,
			Content: post.Content,
			Author: &model.User{
				ID:    string(rune(post.Author.ID)),
				Name:  post.Author.Name,
				Email: post.Author.Email,
				Age:   int32(post.Author.Age),
			},
		})
	}

	return result, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var post models.Post
	if err := r.DB.Preload("Author").First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	var result = &model.Post{
		ID:      string(rune(post.ID)),
		Title:   post.Title,
		Content: post.Content,
		Author: &model.User{
			ID:    string(rune(post.Author.ID)),
			Name:  post.Author.Name,
			Email: post.Author.Email,
			Age:   int32(post.Author.Age),
		},
	}

	return result, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	user := &models.User{
		Name:  input.Name,
		Age:   int(input.Age),
		Email: input.Email,
	}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	var result = &model.User{
		ID:    string(rune(user.ID)),
		Name:  user.Name,
		Email: user.Email,
		Age:   int32(user.Age),
	}

	return result, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	var author models.User
	err := r.DB.First(&author, input.AuthorID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	post := &models.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: author.ID,
	}

	if err := r.DB.Create(post).Error; err != nil {
		return nil, err
	}

	var result = &model.Post{
		ID:      string(rune(post.ID)),
		Title:   post.Title,
		Content: post.Content,
		Author: &model.User{
			ID:    string(rune(post.Author.ID)),
			Name:  post.Author.Name,
			Email: post.Author.Email,
			Age:   int32(post.Author.Age),
		},
	}

	return result, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.Age = int(input.Age)
	user.Name = input.Name
	user.Email = input.Email

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	var result = &model.User{
		ID:    string(rune(user.ID)),
		Name:  user.Name,
		Email: user.Email,
		Age:   int32(user.Age),
	}
	return result, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("user not found")
		}
		return false, err
	}

	if err := r.DB.Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
