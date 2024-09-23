# デフォルトのターゲット
.PHONY: all
all: create-env-file

# 環境変数を.envファイルにエクスポート
.PHONY: create-env-file
create-env-file:
	@echo "Creating .env file from current environment variables..."
	@printenv | grep -v '^_' | sed 's/=.*//g' | sed 's/^/export /' > .env.temp
	@while read line; do \
		var_name=$$(echo $$line | cut -d' ' -f2); \
		var_value=$$(printenv $$var_name); \
		echo "$$var_name=$$var_value" >> .env; \
	done < .env.temp
	@rm .env.temp
	@echo "Created .env file with $(shell wc -l < .env) environment variables."

# .envファイルの内容を表示（セキュリティ上重要な情報は表示しないよう注意）
.PHONY: show-env-file
show-env-file:
	@echo "Contents of .env file (excluding sensitive information):"
	@grep -vE 'KEY|TOKEN|PASSWORD|SECRET' .env || true

# .envファイルを削除
.PHONY: clean
clean:
	@echo "Removing .env file..."
	@rm -f .env
	@echo ".env file removed."

# ヘルプメッセージを表示
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  create-env-file : Create .env file from current environment variables"
	@echo "  show-env-file   : Show contents of .env file (excluding sensitive information)"
	@echo "  clean           : Remove .env file"
	@echo "  help            : Show this help message"