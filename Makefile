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

.PHONY: all build clean
