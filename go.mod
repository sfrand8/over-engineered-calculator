module over-engineered-calculator

go 1.24.1

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/go-chi/chi/v5 v5.2.1
	github.com/sfrand8/over-engineered-calculator/pkg/http-middleware v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
	github.com/swaggo/http-swagger v1.3.4
	github.com/swaggo/swag v1.8.1
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/sfrand8/over-engineered-calculator/pkg/http-middleware => ./pkg/http-middleware
