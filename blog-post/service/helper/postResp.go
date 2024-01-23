package helper

import "blog_post/models"

func PostResp(allPosts models.AllPost, postResp *[]models.PostResponse) error {

	for index, value := range allPosts.Post {

		data := models.PostResponse{
			CreatedAt: value.CreatedAt, PostID: value.PostID, Title: value.Title, Content: value.Content,
			Excerpt: value.Excerpt, Status: value.Status, CategoryID: value.CategoryID, Categories: value.Categories,
			Comments: value.Comments,
		}
		*postResp = append(*postResp, data)
		for _, comment := range value.PostComments {
			comments := models.CommentsResponse{
				CreatedAt: comment.CreatedAt, CommentID: comment.CommentID, Content: comment.Content, Website: comment.Website,
				UserID: comment.Users.UserID, Email: comment.Users.Email, Name: comment.Users.Name,
			}
			
			(*postResp)[index].PostComments = append((*postResp)[index].PostComments, comments)
		}

	}
	return nil
}
