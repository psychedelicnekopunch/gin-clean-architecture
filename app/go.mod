module github.com/psychedelicnekopunch/gin-clean-architecture

go 1.14

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.2.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.16
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
