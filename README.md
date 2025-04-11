# Requirements

1. Take input of power readings for each container in a singular form
2. Save power readings to database after each entry noting time to database
3. export from database to excel file using a date and time picker for a range.

# Technologies
* [golang](https://go.dev/) - language
* [gorm](https://github.com/go-gorm/gorm) - package - DB Access
* [caddy](https://github.com/caddyserver/caddy) - package - Auto TLS and Cert Renewal
* [fiber v2](https://github.com/gofiber/fiber) - package - web framework
* [Sqlite](https://sqlite.org/) - Database - Sqlite (stores all records locally in a singular file)
* [github](https://github.com) - Github
* [excelize](https://github.com/qax-os/excelize)

# Plan
1. Create Repo using github
	1. Use Github
	2. Make Repo Private
2. Create web application
	1. Create web application using the web framework fiber - fiber v2 is stable
3. Link project to repo
4. Install Sqlite
5. Start DB Engine
	1. Start Sqlite
6. Install Caddy
7. Start Caddy
	1. in the command line start caddy
		`caddy`
8. Configure Caddy
	1. Configure Caddy via the Caddy file to specify endpoints
9. Init database
	1. Initialize the Database using gorm
10. Seed database
	1. Seed Database using gorm using gorm for initial data
11. Test Application
	1. Test application for basic functionality and any errors
12. Pentest application
	1. check for sanitized inputs
	2. check for any memory issues
	3. check for any unsecured connections in clear text

Branches
-mrwill
-mrchanman
