t-stack:
	go test ./stack -v
t-queue:
	go test ./queue -v
t-ll:
	go test ./linkedlist -v 

.PHONY: t-stack t-queue t-ll
