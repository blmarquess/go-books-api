run:
		docker-compose up -d --build

stop:
		docker-compose down --remove-orphans

test-all:
		curl --location --request GET 'http://localhost:8000/books' \
		--data-raw ''

test-one:
		curl --location --request GET 'http://localhost:8000/books/1' \
		--data-raw ''