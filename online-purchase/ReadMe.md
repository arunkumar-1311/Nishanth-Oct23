# Online-Purchase Application
This repository contains all needed API's to create online-purchase RTE.

## main.go
This is the main function of the application. Here we acquire a db connection and load all lookup files. Then call the route function to start the server

## adapator
### gormConnection.go 
--- Helps to load the .env and make a new DB connection.

## handler
### account.go
--- type Handlers struct - which is responsible for this whole routing functions.
--- Register(ctx *context.Context) - helps to register new profile to the application.
--- Login(ctx *context.Context) - helps to log-in to the existing profile.

### brand.go
--- CreateBrand(ctx *context.Context) - helps to create a new brand with its needed specification. This API can only hit by Admin.
--- GetBrands(ctx *context.Context) - helps to fetch all the brands available in the website.
--- GetBrandByID(ctx *context.Context) - helps to get the specific brand by its id.
--- UpdateBrand(ctx *context.Context) - helps to update the existing brand and its specification. This API can only hit by Admin.
--- DeleteBrand(ctx *context.Context) - helps to delete the in-active brand. This API can only hit by Admin.

### order.go
--- CreateOrder(ctx *context.Context) - helps to place a new order to our application
--- GetOrderByID(ctx *context.Context) - helps to fetch the order by its unique ID. This feature can be accessed by either Admin or a User who placed a order.
--- GetAllOrders(ctx *context.Context) - helps to fetch all the active orders. This API can only hit by Admin.
--- GetAllOrderStatus(ctx *context.Context) - helps to fetch all order status. This API can only hit by Admin.
--- CancelOrder(ctx *context.Context) - helps to cancel the existing order status. This can be done by the user who placed a order.
--- UpdateStatus(ctx *context.Context) - helps to update the order status of the order. This API can only hit by Admin. If the order status is turned to Delivered then it active status will be turned false
--- DeleteOrder(ctx *context.Context) - helps to soft delete the in-active order.

### ram.go
--- CreateRAM(ctx *context.Context) - helps to create the RAM product with its all specification. This API can be only hit by Admin.
--- GetAllRAMs(ctx *context.Context) - helps to get all RAM's listed in the website.
--- GetRamByID(ctx *context.Context) - helps to fetch the RAM by its unique ID. 
--- UpdateRAM(ctx *context.Context) - helps to change the RAM status and its specification. This API can be only hit by Admin.
--- DeleteRAM(ctx *context.Context) - helps to delete the in-active RAM product. This API can be only hit by Admin.


## logger
### zap-log.go
--- ZapLog() *zap.Logger - helps to track the log and append to the log.log file.

## lookup
### Lookup_000.go
--- Lookup_000(db *gorm.DB) error - helps to create and auto migrate the lookup table.
### Lookup_001.go
--- Lookup_001(db *gorm.DB) error - helps to create and auto migrate the application tables.
### Lookup_002.go
--- Lookup_002(db *gorm.DB) error - helps to insert data into roles assert table.
### Lookup_003.go
--- Lookup_003(db *gorm.DB) error - helps to insert data into order_status assert table.
### master.go

## middleware
### authorization.go
--- Authorization(ctx *context.Context) - helps to check the authentication and authorization of the user for that API.

## models
### application.go
--- type Roles struct - helps to create roles table with its needed column.
--- type OrderStatus struct - helps to create order_status table with its needed column.
--- type Brand struct - helps to create brand table with its needed column.
--- type Ram struct - helps to create ram table with its needed column.
--- type Users struct - helps to create user table with its needed column.
--- type Orders struct - helps to create orders table with its needed column.
--- type Login struct - It assist the login API.
--- type Claims struct - helps to get the claims from the bearer token.
--- type OrderResponse struct - helps to send the order response to the client.
### response.go
--- type Message struct - helps to send the organized message to the client.

## repository
### account.go
--- CreateUser(models.Users) error - 
--- FindUserAndEmail(string, string) (models.Users, error) - 
--- User(string, *models.Users) error - 
### brand.go
--- CreateNewBrand(models.Brand) error
--- ReadBrands(*[]models.Brand) error
--- ReadBrandByID(string, *models.Brand) error
--- UpdateBrandByID(string, models.Brand) error
--- DeleteBrandByID(string) error
