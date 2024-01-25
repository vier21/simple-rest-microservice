mysql:
	sudo docker run -it --name db-mysql -v /opt/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_USER=xavier -e MYSQL_PASSWORD=xavier123 -e MYSQL_DATABASE=category -d -p 3306:3306 mysql:latest
migrate:
	 migrate -path db/migration/ -database "mysql://root:root123@tcp(127.0.0.1:3306)/category" -verbose up
db_test:
	sudo docker run -it --name db-mysql-test -v /opt/mysql/data_test:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_USER=xavier -e MYSQL_PASSWORD=xavier123 -e MYSQL_DATABASE=category_test -d -p 2109:3306 mysql:latest
migrate_test:
	migrate -path db/migration_test/ -database "mysql://root:root123@tcp(127.0.0.1:2109)/category_test" -verbose up

