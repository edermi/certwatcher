module github.com/drfabiocastro/certwatcher

go 1.19

replace internal/runner => ./internal/runner

replace pkg/config => ./pkg/config

require (
	github.com/projectdiscovery/gologger v1.1.7 // indirect
	internal/runner v0.0.0-00010101000000-000000000000
)

require (
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mholt/archiver v3.1.1+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	gopkg.in/djherbis/times.v1 v1.3.0 // indirect
	pkg/config v0.0.0-00010101000000-000000000000 // indirect
)
