# https://hub.docker.com/_/postgres
FROM postgres:15

# Install postgis extension
RUN apt-get update \
      && apt-get install wget -y \
      && apt-get install postgresql-15-postgis-3 -y \
      && apt-get install postgis -y


# Database initialization SQL
COPY ./0_init.sql /docker-entrypoint-initdb.d/0_init.sql