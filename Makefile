GOCMD = go
GOBUILD = $(GOCMD) build
GOFLAGS = -ldflags="-s -w" -trimpath
TARGET = uploader
BUILD_DIR = build

all: build

build:
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(GOFLAGS) -o $(BUILD_DIR)/$(TARGET)

clean:
	rm -rf $(BUILD_DIR)

# Version & Release (CI builds on tag push)
patch:
	@new_tag=$$(svu patch) && \
	git commit --allow-empty -m "chore: bump to $$new_tag" && \
	git tag "$$new_tag" && \
	echo "Tagged $$new_tag"

minor:
	@new_tag=$$(svu minor) && \
	git commit --allow-empty -m "chore: bump to $$new_tag" && \
	git tag "$$new_tag" && \
	echo "Tagged $$new_tag"

major:
	@new_tag=$$(svu major) && \
	git commit --allow-empty -m "chore: bump to $$new_tag" && \
	git tag "$$new_tag" && \
	echo "Tagged $$new_tag"

release:
	@git push origin HEAD $$(git describe --tags --abbrev=0) && \
	echo "Pushed $$(git describe --tags --abbrev=0)"

.PHONY: all build clean patch minor major release
