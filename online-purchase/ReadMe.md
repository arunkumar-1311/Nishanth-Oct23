# Online-Purchase Application
This repository contains all needed API's to create online-purchase RTE.<br />

## main.go
This is the main function of the application. Here we acquire a db connection and load all lookup files. Then call the route function to start the server<br />

## adapator
### gormConnection.go 
--- Helps to load the .env and make a new DB connection.<br />

## handler
### account.go
--- type Handlers struct - which is responsible for this whole routing functions.<br />
--- Register(ctx *context.Context) - helps to register new profile to the application.<br />
--- Login(ctx *context.Context) - helps to log-in to the existing profile.<br />

### brand.go
--- CreateBrand(ctx *context.Context) - helps to create a new brand with its needed specification. This API can only hit by Admin.<br />
--- GetBrands(ctx *context.Context) - helps to fetch all the brands available in the website.<br />
--- GetBrandByID(ctx *context.Context) - helps to get the specific brand by its id.<br />
--- UpdateBrand(ctx *context.Context) - helps to update the existing brand and its specification. This API can only hit by Admin.<br />
--- DeleteBrand(ctx *context.Context) - helps to delete the in-active brand. This API can only hit by Admin.<br />

### order.go
--- CreateOrder(ctx *context.Context) - helps to place a new order to our application.<br />
--- GetOrderByID(ctx *context.Context) - helps to fetch the order by its unique ID. This feature can be accessed by either Admin or a User who placed a order.<br />
--- GetAllOrders(ctx *context.Context) - helps to fetch all the active orders. This API can only hit by Admin.<br />
--- GetAllOrderStatus(ctx *context.Context) - helps to fetch all order status. This API can only hit by Admin.<br />
--- CancelOrder(ctx *context.Context) - helps to cancel the existing order status. This can be done by the user who placed a order.<br />
--- UpdateStatus(ctx *context.Context) - helps to update the order status of the order. This API can only hit by Admin. If the order status is turned to Delivered then it active status will be turned false.<br />
--- DeleteOrder(ctx *context.Context) - helps to soft delete the in-active order.<br />

### ram.go
--- CreateRAM(ctx *context.Context) - helps to create the RAM product with its all specification. This API can be only hit by Admin.<br />
--- GetAllRAMs(ctx *context.Context) - helps to get all RAM's listed in the website.<br />
--- GetRamByID(ctx *context.Context) - helps to fetch the RAM by its unique ID. <br />
--- UpdateRAM(ctx *context.Context) - helps to change the RAM status and its specification. This API can be only hit by Admin.<br />
--- DeleteRAM(ctx *context.Context) - helps to delete the in-active RAM product. This API can be only hit by Admin.<br />


## logger
### zap-log.go
--- ZapLog() *zap.Logger - helps to track the log and append to the log.log file.<br />

## lookup
### Lookup_000.go
--- Lookup_000(db *gorm.DB) error - helps to create and auto migrate the lookup table.<br />
### Lookup_001.go
--- Lookup_001(db *gorm.DB) error - helps to create and auto migrate the application tables.<br />
### Lookup_002.go
--- Lookup_002(db *gorm.DB) error - helps to insert data into roles assert table.<br />
### Lookup_003.go
--- Lookup_003(db *gorm.DB) error - helps to insert data into order_status assert table.<br />
### master.go

## middleware
### authorization.go
--- Authorization(ctx *context.Context) - helps to check the authentication and authorization of the user for that API.<br />

## models
### application.go
--- type Roles struct - helps to create roles table with its needed column.<br />
--- type OrderStatus struct - helps to create order_status table with its needed column.<br />
--- type Brand struct - helps to create brand table with its needed column.<br />
--- type Ram struct - helps to create ram table with its needed column.<br />
--- type Users struct - helps to create user table with its needed column.<br />
--- type Orders struct - helps to create orders table with its needed column.<br />
--- type Login struct - It assist the login API.<br />
--- type Claims struct - helps to get the claims from the bearer token.<br />
--- type OrderResponse struct - helps to send the order response to the client.<br />
### response.go
--- type Message struct - helps to send the organized message to the client.<br />

## repository
### account.go
--- CreateUser(models.Users) error - helps to create a new profile with needed fields. <br />
--- FindUserAndEmail(string, string) (models.Users, error) - helps to check wheather the user name or email id is exist. <br />
--- User(string, *models.Users) error - helps to fetch the existing user profile. <br />
### brand.go
--- CreateNewBrand(models.Brand) error - helps to create new brand for the application. <br />
--- ReadBrands(*[]models.Brand) error - helps to read all brands in the website. <br />
--- ReadBrandByID(string, *models.Brand) error - helps to read the brand by its unique ID. <br />
--- UpdateBrandByID(string, models.Brand) error - helps to update the existing brand details. <br />
--- DeleteBrandByID(string) error - helps to delete the in-active brand of the application. <br />
### db_connection.go
--- type GORM_Connection struct - helps to hold the gorm database connection. <br />
### order.go
--- CreateNewOrder(models.Orders) error - helps to create a order. <br />
--- ReadOrders(*[]models.Orders) error - helps to read all orders placed in the application. <br />
--- ReadOrderByID(string, *models.Orders) error -  helps to read the order by its unique ID. <br />
--- UpdateOrderByID(string, models.Orders) error - helps to update the order by its unique ID. <br />
--- DeleteOrderByID(string) error - helps to delete the order by its unique ID. <br />
--- CancelOrderByID(string, bool) error - helps to cancel the order by its unique ID. <br />
--- ReadAllOrderStatus(*[]models.OrderStatus) error - helps to read order status. <br />
--- UpdateOrderStatusByID(string, string) error - helps to update the order status ID by order ID. <br />
### ram.go
--- CreateNewRam(models.Ram) error - helps to create the new RAM product. <br />
--- ReadRAMs(*[]models.Ram) error - helps to read all RAM product in th Database. <br />
--- ReadRAMByID(string, *models.Ram) error - helps to read the RAM specification with its ID. <br />
--- UpdateRAMByID(string, models.Ram) error - helps to update the RAM specification with its ID. <br />
--- DeleteRAMByID(string) error - helps to delete the RAM product with its ID. <br />

## router
### routes.go
--- Routes(db adaptor.Database) - contains all handlers and path. <br />
GET : http://localhost:8000/brands <br />
GET : http://localhost:8000/rams <br />
GET : http://localhost:8000/brand/:id <br />
GET : http://localhost:8000/ram/:id <br />
GET : http://localhost:8000/admin/orders <br />
GET : http://localhost:8000/admin/orderstatus <br />
GET : http://localhost:8000/order/:id <br />

POST : http://localhost:8000/user/new <br />
POST : http://localhost:8000/user <br />
POST : http://localhost:8000/admin/brand <br />
POST : http://localhost:8000/admin/ram <br />
POST : http://localhost:8000/order <br />

PATCH : http://localhost:8000/admin/brand/:id <br />
PATCH : http://localhost:8000/admin/ram/:id <br />
PATCH : http://localhost:8000/admin/order/:id <br />
PATCH : http://localhost:8000/order/:id <br />

DELETE : http://localhost:8000/admin/brand/:id <br />
DELETE : http://localhost:8000/admin/ram/:id <br />
DELETE : http://localhost:8000/admin/order/:id <br />


## service
### response.go
--- SendResponse(c *context.Context, status int, err string, message string, data ...interface{}) error - helps to send the response to the client. <br />
### adminAccess.go
--- AdminAccess(tokenString string) error - check it the given token is token of Admin. <br />
### emailVerification.go
--- EmailAndNameValidation(user models.Users, db adaptor.Database) (result error) - helps to check the user name and email ID is already exist. <br />
### password.go
--- GenerateHash(password *string) error - helps to hash the given string password. <br />
--- CompareHashPassword(password, hash string) bool - helps to compare the hash and password and return true or false. <br />
### token.go
--- CreateToken(username, email, role, userID string) (string, error) - helps to create the token with given fields. <br />
--- VerifyToken(tokenString string) (tokenData []byte, err error) - check the given token is valid or not. <br />
### UUID.go
--- UniqueID() string - helps to generate the unique UUID. <br />

## .env
Contains the database credentials to acquire postgres Database connection. <br />
--- host <br />
--- port <br />
--- user <br />
--- password <br />
--- dbname <br />