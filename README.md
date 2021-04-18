# RENTAL BUKU 20210415
==============
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
* **req_register.http** : register new user
* **req_login.http** : login user
* **req_get_profile.http** : get user profile (require Authorization)
* **req_put_profile.http** : update user profile (require Authorization)

-------
License
-------
LICENSE.