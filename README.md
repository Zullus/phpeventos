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

2º. **Opte pela forma de executar o sistema com Docker**

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

* melhorar o retorno da página web com reloads automáticos;
* achar um calendário que seja clicável
* adicionar o docker para rodar o sistema com Go
* refazer o código da API em C#
* refazer o código da API em Python
* refazer o código da API em Java