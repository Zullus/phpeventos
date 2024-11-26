# phpeventos

**Descrição:**

phpeventos é uma aplicação web PHP para gerenciamento de eventos. Permite criar, editar e listar eventos.

**Pré-requisitos:**

* PHP >= 7.4
* Composer
* Servidor web (Apache, Nginx)
* Banco de dados MySQL ou MariaDB

**Instalação:**

1º. **Clone o repositório:**

   ```bash
   git clone https://github.com/Zullus/phpeventos
   ```

2º. **Opte pela forma de executar o sistema**

### Usando seu servidor MySQL

Caso tenha um servidor MySQL configurado, execute em seu banco o arquivo

   ```bash
   /sql/baseinicial.sql
   ```

Esse arquivo irá criar a tabela _eventos_ e o usuário _eventos_ com senha _eventos_ em seu banco de dados.

Crie um arquivo ".env" na pasta "phpeventos/app/src" à partir do arquivo ".env-example" com as configurações do seu banco de dados.

Vá até a pasta "phpeventos/app/src" e execute o comando:

   ```bash
   composer install
   ```

Após, suba o servidor web do PHP com o comando:

   ```bash
   php -S localhost:8090
   ```

Abra o nagador em <http://localhost:8090>.

A tela inicial deve ser exibida.

#### Usando Docker

Caso tenha Docker em seu computador, use o Dockerfile na pasta "phpeventos/" e execute o comando:

   ```bash
   docker build -t alpine-apache-php:latest .  
   ```

Após o fim do _build_, execute o comando:

   ```bash
   cd docker  
   ```

Após, execute o Docker Compose com o servidor Web Apache, o MySQL e o PhpMyAdmin com o comando:

   ```bash
   docker compose up -d  
   ```

Abra o nagador em <http://localhost:8080>.

Caso precise de ajuda com o banco de dados, que está no docker-compose, você pode usar o PhpMyAdmin abrindo o navegador no endereço <http://localhost:8090>, sendo usuário e senha _eventos_

Qualquer dúvida, o podem me chamar pelo meus contatos.

---

#### TODO

* separar o app em dois microsserviços, onde o HTML poderia ficar em um servidor NGIX mais simples e o PHP em outro;
* melhorar o retorno da página web com reloads automáticos;
* achar um calendário que seja clicável
