# Use an official MySQL Docker image as a base
FROM mysql:latest

# Environment variables
ENV MYSQL_DATABASE=classicmodels
ENV MYSQL_USER=admin
ENV MYSQL_PASSWORD=secret
ENV MYSQL_ROOT_PASSWORD=rootpassword

ADD https://raw.githubusercontent.com/hhorak/mysql-sample-db/master/mysqlsampledatabase.sql /docker-entrypoint-initdb.d/init.sql

# Grant remote access and create user
RUN echo "CREATE DATABASE IF NOT EXISTS \`$MYSQL_DATABASE\` ;" >> /docker-entrypoint-initdb.d/init.sql && \
    echo "CREATE USER '$MYSQL_USER'@'%' IDENTIFIED BY '$MYSQL_PASSWORD' ;" >> /docker-entrypoint-initdb.d/init.sql && \
    echo "GRANT ALL ON \`$MYSQL_DATABASE\`.* TO '$MYSQL_USER'@'%' ;" >> /docker-entrypoint-initdb.d/init.sql && \
    echo "FLUSH PRIVILEGES ;" >> /docker-entrypoint-initdb.d/init.sql && \
    chmod 644 /docker-entrypoint-initdb.d/init.sql \
