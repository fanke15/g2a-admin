APP_NAME = "g2aAdmin"

build:
	chmod 777 build/build.sh
	./build/build.sh $(APP_NAME)

.PHONY:  build
