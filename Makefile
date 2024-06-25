in_memory:
	@echo "Setting storage to in-memory"
	@echo STORAGE=in_memory > .env
	docker-compose --profile no-db up --build

postgresql:
	@echo "Setting storage to PostgreSQL"
	@echo STORAGE=postgresql > .env
	docker-compose --profile with-db up --build

clean:
	@echo "Cleaning up"
	docker-compose down -v