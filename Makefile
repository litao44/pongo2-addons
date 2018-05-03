all: build

export GOPATH 		:= $(CURDIR)/_project

CURRENT_GIT_GROUP 	:= github.com/litao44
CURRENT_GIT_REPO 	:= pongo2-addons

folder_dep:
	mkdir -p $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)
	test -d $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO) || ln -s $(CURDIR) $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)

deps: folder_dep
	mkdir -p $(CURDIR)/vendor
	glide install

test:
	go test -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)

clean:
	@rm -rf _project

.PHONY:  deps test clean docker
