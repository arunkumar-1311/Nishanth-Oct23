package helper

import (
	"blog_post/models"
	"blog_post/repository"
)

// Helps to find the total comments and categories in the post
func CommentsAndCategory(Post []models.Post) error {

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

		for _, value := range Post[i].CategoryID {
			category, err := repository.ReadCategoryByID(value)
			if err != nil {
				return err
			}
			Post[i].Categories = append(Post[i].Categories, category)
		}
	}

	return nil
}

// Helps to find the all categories
func Categories(Categories []models.CategoryResponse) error {
	var totalPost []models.Post

	category := make(map[string]int)

	if err := repository.ReadPosts(&totalPost); err != nil {
		return err
	}

	for i := 0; i < len(totalPost); i++ {
		for key, value := range totalPost[i].CategoryID {
			name, err := repository.ReadCategoryByID(value)
			if err != nil {
				return err
			}
			category[name] = category[name] + 1
			totalPost[i].CategoryID[key] = name
		}

	}

	for i := 0; i < len(Categories); i++ {
		for key := range category {
			if Categories[i].Category.Name == key {
				Categories[i].Total = category[key]
			}
		}
	}

	return nil
}
