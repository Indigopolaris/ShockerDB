# Overview

	ShockerDB consists of an electron like UI app built from the Wails project with a fiber based web backend that transfers and retreives data from a (debian?) based server running an SQLite instance. Data is categorized by date, time, and client. The aim of the project is to introduce a higher level of data security, improve control of KwH data storage, and to minimize data errors.

# Requirements

 1. Take input of power readings for each container in a singular form.
 2. Save power readings to database after each entry noting time to database.
 3. export from database to excel file using a date and time picker for a range.

# Technologies

 * [golang](https://go.dev/) - language
 * [gorm](https://github.com/go-gorm/gorm) - package - DB Access
 * [caddy](https://github.com/caddyserver/caddy) - package - Auto TLS and Cert Renewal
 * [fiber v2](https://github.com/gofiber/fiber) - package - web framework
 * [Sqlite](https://sqlite.org/) - Database - Sqlite (stores all records locally in a singular file)
 * [github](https://github.com) - Github
 * [excelize](https://github.com/qax-os/excelize)
 * [Wails](https://wails.io/)  - [git] (https://github.com/wailsapp/wails)
 * Html (https://www.w3schools.com/html/)
 * Http (https://www.w3schools.com/whatis/whatis_http.asp), (https://www.w3.org/Protocols/)
 * Typescript (https://www.typescriptlang.org/)
 * CSS (https://www.w3schools.com/css/)
 * CSS bootstrap (https://getbootstrap.com/)
 * JSON

# Plan

	1. ~~Create Repo using github~~
  1. ~~Use Github~~
  2.~~ Make Repo Private~~
 2. Create web application
  1. Create web application using the web framework fiber - fiber v2 is stable
 3. Create wails app with react (electron)
  1. follow wails.io instructions, linux dev install might have issues depending on shell
 4. Link webapp and wails project to repo for central packaging
 6. Install Caddy
 7. Start Caddy
  1. in the command line start caddy
   `caddy`
 8. Configure Caddy
  1. Configure Caddy via the Caddy file to specify endpoints, transfer over http(s)
 9. Init database
  1. Initialize the Database using gorm, add connection string to project, add to .gitignore
 10. Seed database
  1. Seed Database using gorm using gorm for initial data
 11. Test Application
  1. Test application for basic functionality and any errors
 12. Pentest application
  1. check for sanitized inputs
  2. check for any memory issues
  3. check for any unsecured connections in clear text
  4. ensure code is anti-sql injection


## Branches

-mrwill
-mrchanman

# App layout
	App pages
	 1. power reading submission
		-select datetime, data entered into boxes for site readings -> auto divison between clients, amp readings auto calculated if applicable
		- edit data box for individual or range of readings by meter and datetime (should the replaced reading be saved?)
     2. reading export page
        -select client(s) and datetime range, exports to .csv
	 3. admin panel (requires admin pw)
		- edit clients, box owner, meter type (KwH, or amp), link to teams excel (onedrive), admin recovery pw listed in admin details .gitignore page

# structs/objects and classes TBD
 - file for each page
 - client
 - meter (amp,kwh)
 - PR HERE!! ---

# tech stack

 server side
  -sqllite DB
  -debian server hosting http ShockerDB web app

 client side
  - wails (electron) app
  - connects to shockerDB web page
  - app retrieves/stores data to/from DB

# app dev needs

	- site obj -> container obj -> breaker obj ->
 - switch breaker obj params between kw and amp

 -need DB server settings
 -settings needs
 - obj adding (new containers, etc)
 - future support for rs485 communications to server

# About wails

using the official Wails React-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: <https://wails.io/docs/reference/project-config>

## wails Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on <http://localhost:34115>. Connect
to this in your browser, and you can call your Go code from devtools.

## wailsapp Building

To build a redistributable, production mode package, use `wails build`.
