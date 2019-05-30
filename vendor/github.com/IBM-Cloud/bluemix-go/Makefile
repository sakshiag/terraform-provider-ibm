.PHONY : test
test: test_deps vet
<<<<<<< HEAD:vendor/github.com/IBM-Cloud/bluemix-go/Makefile
	go test ./...
=======
	go test ./... -timeout 120m
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe:vendor/github.com/IBM-Cloud/bluemix-go/Makefile

.PHONY : test_deps
test_deps: 
	go get -t ./...

.PHONY : vet

vet:
	@echo 'go vet $$(go list ./... | grep -v vendor)'
	@go vet $$(go list ./... | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi
