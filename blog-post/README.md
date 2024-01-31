# Blog-post-Nishanth
Add Go lang environment to your system.

## main.go
main.go is the base file which run while compiling.

## adaptor
Contains all needed database connection in our project we have postgres connection with a help of GORM.
### gormConnection
It helps to attain the basic configuration for database operation.

## handler
This package holds all needed fiber handler function.
comments.go - needed api's for perform comment feature in the blog post.
filter.go - filter the content as per the user wish. By its category or its date of publish.
middleware.go - contains all needed middleware's like authentication, authorization and validation.
register.go - helps to register the new user for our website.
login.go - It holds the handler which helps to log-in to the existing account.
 
 ### admin
This package contains all admin operation handlers.
category.go - Contains all needed API to perform crud with categories
post.go - Contains all needed API to perform crud with post

### admin
It contain all admin accessable features.
category.go - helps to perform all needed operations with categories basically a CRUD api's.
post.go - it contains all needed api's to manipulate the post, we're going to publish or published posts.

## logger
This package helps to record all needed logging details for our application.
logrusLog.go - it contains basic logging using logrus package.

## models
Contains all needed type struct to work in this project.
application.go - holds all neccessory struct for api's.
message.go - contains all needed response struct.

## repository
It contains all db operation functions.
category.go - It carries needed queries to manipulate blog post categories.
comments.go - It carries needed queries to manipulate blog post comments.
filter.go - Contains all queries to handle blog post filter feature.
login.go - Holds needed queries to handle login operation of blog post.
overview.go - Carries all the needed queries to showcase the overview of the profile.
post.go - It carries needed queries to manipulate the posts of this application.
register.go - Contains all needed queries to create a new user for blog post application

## router
Contains all routers and paths to hit the api's.

## service
This package consist all assisting functions for the flow of application.
firstPost.go - helps to calculate the time difference between current date and the first post realised.
response.go - It holds the function with send the json response to the client.
### helper
adminAccess.go - check wheather the user is admin or user.
categoryAndComments.go - helps to load the category and the comments of certain post.
emailVerification.go - helps to verify the given email id is already exist or not.
password.go -  helps to generate hash of the given password and compare the hash with user login password.
token.go - It holds the business logic to create a jwt token and verify the token is valid or not.
tokenClaims.go - It gives the cliams of the given token.


