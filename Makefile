PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))

$(info name is $(NAME))

build:
	go build

clean:
	rm $(NAME)
