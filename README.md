# RENTAL BUKU 20210415
-----
About
-----
This is a API client of RENTAL BUKU written in Go with MySQL + Gin + Gorm + JWT.

-----
Usage
-----
Don't forget to adjust ./conf/db-conf.go for Database Connection Configuration 
for the first run server, it will create table with AutoMigration feature
.. code-block:: go

    db.AutoMigrate(entity.User{}, entity.Category{}, entity.Book{}, entity.Loan{}, entity.FinePayment{})

## Usage API
You can use the files contained in the root folder with the extension (req_*.http) and make requests using the REST Client plugin on VS Code
User sample requests:
* **req_get_users.http** : get All User
* **req_profile_get.http** : get user profile (require Authorization)
* **req_profile_put.http** : update user profile (require Authorization)
* **req_user_login.http** : login user
* **req_user_register.http** : register new user
Category sample requests:
* **req_category_delete.http** : Delete Category by ID
* **req_category_get.http** : Get All Category and Get By ID
* **req_category_insert.http** : Entry Category
* **req_category_update.http** : Update Category
Book sample requests:
* **req_book_delete.http** : Delete Book by ID
* **req_book_get.http** : Get Books and get by ID
* **req_book_insert.http** : Entry Book
* **req_book_update.http** : Update Book
Loan sample requests:
* **req_loan_get.http** : Get All Loan and get by ID
* **req_loan_insert.http** : Entry Loan
* **req_loan_returnBook.http** : Return the Book that has been borrowed By LoanID

-------
License
-------
LICENSE.