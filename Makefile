.PHONY: dev

dev:
ifeq ($(OS),Windows_NT)
	air -c .air.windows.toml
else
	air
endif 