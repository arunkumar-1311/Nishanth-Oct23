package helper

import (
	"blog_post/adaptor"
	"blog_post/models"
	
)

// Helps to find the total comments and categories in the post
func CommentsAndCategory(Post []models.Post, db adaptor.Database) error {

	// Helps to find the number of comments in that particular posts
	for i := 0; i < len(Post); i++ {
		var count int64

		if err := db.NumberOfComments(Post[i].PostID, &count); err != nil {
			return err
		}
		Post[i].Comments = int(count)
		if err := db.PostComments(Post[i].PostID, &Post[i].PostComments); err != nil {
			return err
		}

		for _, value := range Post[i].CategoryID {
			category, err := db.ReadCategoryByID(value)
			if err != nil {
				return err
			}
			Post[i].Categories = append(Post[i].Categories, category)
		}
	}

	return nil
}

// Helps to find the all categories
func Categories(Categories []models.CategoryResponse, db adaptor.Database) error {
	var totalPost []models.Post

	category := make(map[string]int)

	if err := db.ReadPosts(&totalPost); err != nil {
		return err
	}

	for i := 0; i < len(totalPost); i++ {
		for key, value := range totalPost[i].CategoryID {
			name, err := db.ReadCategoryByID(value)
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
