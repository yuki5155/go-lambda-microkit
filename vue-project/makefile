.PHONY: copy-frontend

copy-frontend:
	@if [ ! -d "src" ]; then \
		echo "Error: src directory not found."; \
		exit 1; \
	fi
	@if [ -d "front_files" ]; then \
		echo "Cleaning front_files directory..."; \
		rm -rf front_files/*; \
	else \
		echo "Creating front_files directory..."; \
		mkdir -p front_files; \
	fi
	@echo "Copying .ts and .vue files from src/ to front_files/..."
	@find src -type f \( -name "*.ts" -o -name "*.vue" \) -exec cp {} front_files/ \;
	@echo "Frontend files copied successfully."

	@echo "Copied files:"
	@ls -1 front_files/