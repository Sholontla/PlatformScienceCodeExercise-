SERVICES :=  cloud_service/prometheus finance_server publisher_service ws_service  front-end-finance
.PHONY: all $(SERVICES)

all: $(SERVICES) create_topic

$(SERVICES):
	cd $@ && docker-compose up -d

create_topic:
	docker exec broker \
	kafka-topics --bootstrap-server broker:9092 \
				--create \
				--topic orderMessage