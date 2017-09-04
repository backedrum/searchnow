BINARY=searchnow

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
	govendor sync
	go test -v ./display
	go test -v ./handlers
	go build -o ${BINARY} *.go

format:
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

$(BINARY)_no_tests:
	govendor sync
	go build -o ${BINARY} *.go
