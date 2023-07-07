Setup da aplicação:

1) Checkout na branch em seu computador
2) No terminal, executar o comando `docker-compose up --build -d`

O Endereço da aplicação é http://localhost:8080/.

As tabelas do banco de dados são criadas automaticamente durante o build do container "mysql" (os comandos estão em /docker/provision/mysql/init/01-databases.sql).