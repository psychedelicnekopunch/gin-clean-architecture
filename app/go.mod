module github.com/psychedelicnekopunch/gin-clean-architecture

go 1.14

require (
	github.com/gin-gonic/gin v1.7.0
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.2.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
