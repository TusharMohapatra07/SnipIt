# Snippetbox  

Snippetbox is a lightweight web application for managing code snippets, built in Go. <!-- It features user authentication, database integration, secure web practices, and demonstrates best practices for building modern web applications. -->

## Features  

- Create, view, and manage code snippets.
- HTML templating for dynamic web pages.
- PostgreSQL database integration for persistent storage.
- Modular code structure for scalability and maintainability.
- Secure web development practices, including CSRF protection and input validation.
- User authentication with session-based login.  

## Requirements

- Go 1.20+
- PostgreSQL 13+

## Installation

1. **Clone the repository**
```bash
   git clone https://github.com/TusharMohapatra07/snippetbox.git
   cd snippetbox
```

2. **Install Dependencies**
   Ensure you have Go installed. Then run:
```bash
  go mod tidy
```

3. **Copy `.env.example` to a new file named `.env` in the same directory**
  You can use the following command:
```bash
  cp .env.example .env
```

4. **Setup the database**
  Make sure PostgreSQL is installed and running. Create the database and schema:
```bash
  psql -U postgres
  CREATE DATABASE snippetbox;
  \c snippetbox
  \i ./db-table-create.sql
```

5. **Set connection string**
     Replace the `CONNSTR` placeholder with an actual `PostgreSQL` connection string. For Example:
```make
  CONNSTR=postgres://username:password@localhost:5432/snippetbox?sslmode=disable
```

6. **Generate TLS certificates**  
The application requires TLS certificates for secure connections.  

- First, create a `tls` directory in the project root:  
  ```bash
  mkdir tls
  ```
- Change into the `tls` directory:  
  ```bash
  cd tls
  ```
- Run Go’s built-in certificate generator:  
  ```bash
  go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
  ```
  This will generate `cert.pem` and `key.pem` inside the `tls` directory.


7. **Run the application**
    *Optional: Change the port by modifying the `ADDRESS` variable in `.env` file (default is :4040)*
```bash
  make 
```
  Open your browser and go to `http://localhost:<PORT>`.

## License  

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.  
