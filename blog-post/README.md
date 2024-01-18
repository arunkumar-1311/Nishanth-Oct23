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
register.go - helps to register the new user for our website
### admin
It contain all admin accessable features.
category.go - helps to perform all needed operations with categories basically a CRUD api's.
post.go - it contains all needed api's to manipulate the post, we're going to publish or published posts.

## logger
This package helps to record all needed logging details for our application.
logrusLog.go - it contains basic logging using logrus package.

## models
Contains all needed type struct to work in this project
application.go - holds all neccessory struct for api's
message.go - contains all needed response struct


