package helper

import (
	"blog_post/models"
	"blog_post/repository"
	"fmt"
)

// Helps to find the total comments and categories in the post
func CommentsAndCategory(Post []models.Post, CategoriesCount *[]models.CategoriesCount, Archieves *[]string) error {

	// Helps to find the number of comments in that particular posts
	for i := 0; i < len(Post); i++ {
		var count int64

		if err := repository.NumberOfComments(Post[i].PostID, &count); err != nil {
			return err
		}
		Post[i].Comments = int(count)
		if err := repository.PostComments(Post[i].PostID, &Post[i].PostComments); err != nil {
			return err
		}
	}
	// Helps to find the all categories and archieves data
	var totalPost []models.Post
	category := make(map[string]int)
	archieve := make(map[string]int)
	if err := repository.ReadPosts(&totalPost); err != nil {
		return err
	}

	for i := 0; i < len(totalPost); i++ {
		year, month, _ := totalPost[i].CreatedAt.Date()
		archieve[fmt.Sprint(month.String(), "-", year)] = 1

		for key, value := range totalPost[i].CategoryID {
			name, err := repository.ReadCategoryByID(value)
			if err != nil {
				return err
			}
			category[name] = category[name] + 1
			totalPost[i].CategoryID[key] = name
		}

	}

	for key := range category {
		data := models.CategoriesCount{CategoryName: key, Total: category[key]}
		*CategoriesCount = append(*CategoriesCount, data)
	}

	for key := range archieve {
		*Archieves = append(*Archieves, key)
	}
	return nil
}
