package factory

import (
	"math/rand"

	"github.com/linn221/go-blog/data"
	"github.com/linn221/go-blog/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ          ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MakePosts(count int) []models.Post {
	posts := make([]models.Post, count)
	categoryCount := len(data.CategoryNames)
	tagCount := len(data.TagNames)
	// each post having 3 tags
	for i := range posts {
		posts[i] = models.Post{
			Title:      randSeq(10),
			Content:    randSeq(20),
			CategoryID: uint(rand.Intn(categoryCount) + 1),
			Tags: []models.Tag{
				{
					ID: uint(rand.Intn(tagCount) + 1),
				},
				{
					ID: uint(rand.Intn(tagCount) + 1),
				},
				{
					ID: uint(rand.Intn(tagCount) + 1),
				},
			},
		}
	}
	return posts
}
