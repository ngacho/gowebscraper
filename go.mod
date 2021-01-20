module webscraper

go 1.15

require (
// First of all, we need to install the colly dependency
// to do this I highly recommend to use go module
// just run go mod init <project-name> this will generate the go.mod file
// where all dependencies used in the project will be.

	github.com/PuerkitoBio/goquery v1.6.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)
