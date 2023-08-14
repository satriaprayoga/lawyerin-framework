BINARY_NAME=menundaApp.exe

build:
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} .
	@echo Menunda built!


run: build
	@echo Staring Menunda...
	@.\tmp\${BINARY_NAME} &
	@echo Menunda started!

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo Cleaned!

start_compose:
	@docker-compose up -d

stop_compose:
	@docker-compose down

test:
	@echo Testing...
	@go test ./...
	@echo Done!


start: run

stop:
	@echo Ending services...
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Menunda

restart: stop start
