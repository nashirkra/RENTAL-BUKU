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

You can use the files contained in the root folder with the extension (req _ *. Http) and make requests using the REST Client plugin on VS Code

-------
License
-------
LICENSE.