# Makefile
# project : KIS_HTS

APP_NAME=KIS_HTS
MAIN=cmd/app/main.go

BUILD_DIR=build

GO=go

# 환경
ENV_DEV=APP_ENV=dev
ENV_PAPER=APP_ENV=paper
ENV_PROD=APP_ENV=prod

.PHONY: help
help:
	@echo "==== KIS_HTS Commands ===="
	@echo "make run-dev      # 개발 환경 실행"
	@echo "make run-paper    # 모의투자 실행"
	@echo "make run-prod     # 실거래 실행"
	@echo "make build        # 현재 OS 빌드"
	@echo "make build-mac    # macOS 빌드"
	@echo "make build-win    # Windows 빌드"
	@echo "make clean        # build 삭제"

#################################################
# RUN
#################################################

run-dev:
	$(ENV_DEV) $(GO) run $(MAIN)

run-paper:
	$(ENV_PAPER) $(GO) run $(MAIN)

run-prod:
	$(ENV_PROD) $(GO) run $(MAIN)

#################################################
# BUILD
#################################################

build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN)

#################################################
# MAC BUILD
#################################################

build-mac:
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)-mac-amd64 $(MAIN)
	GOOS=darwin GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)-mac-arm64 $(MAIN)

#################################################
# WINDOWS BUILD
#################################################

build-win:
	mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)-win-amd64.exe $(MAIN)

#################################################
# CLEAN
#################################################

clean:
	rm -rf $(BUILD_DIR)