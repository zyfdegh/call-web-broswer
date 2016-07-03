# call-broswer
call system default web broswer to open a link.

# Run

check your go version by `go version`

if go version < 1.5.0
	* go get github.com/skratchdot/open-golang/open
	* set GOPATH
	* go build main.go
	* run

else if go version == 1.5.x
	* go get github.com/Masterminds/glide
	* copy this project under GOPATH
	* glide init && glide install
	* export GO15VENDOREXPERIMENT=1
	* go build src/main.go

else if go version >=1.6.0
	* go get github.com/Masterminds/glide
	* copy this project under GOPATH
	* glide init && glide install
	* go build src/main.go

# Related

github.com/zyfdegh/mdviewer
