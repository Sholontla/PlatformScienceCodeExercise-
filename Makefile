SERVICES := analytics file_server fleet_service

.PHONY: all $(SERVICES)

all: $(SERVICES)

$(SERVICES):
	cd $@ && docker-compose up -d