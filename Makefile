# Tutorial repo root. The Go module is heroes-service/.
.PHONY: test vet lint race build run

test vet lint race build run:
	$(MAKE) -C heroes-service $@
