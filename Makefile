in_memory:
	@echo "Setting storage to in-memory"
	@echo STORAGE=in_memory > .env
	docker-compose up --build

postgresql:
	@echo "Setting storage to PostgreSQL"
	@echo STORAGE=postgresql > .env
	docker-compose up --build

clean:
	@echo "Cleaning up"
	docker-compose down -v