build:
	go build -o inline-sample main.go

run:
	./inline-sample

clean:
	rm inline-sample

start: build run clean

showin:
	go run -gcflags="-m" main.go

showinde:
	go run -gcflags="-d pctab=pctoinline" main.go

bench:
	go test -benchmem -run=^$ -bench .