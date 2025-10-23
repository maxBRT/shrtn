# Shrtn

[Shrtn](https://shrtn.it.com/) is a self-hosted URL shortener built with Go.
It provides a simple REST API for creating and managing short links, backed by a PostgreSQL database and secured behind a custom [Feather API Gateway](https://github.com/maxBRT/feather).

This project was designed to demonstrate a simple, yet full production setup:

A Go backend handling short URL generation and redirects

Feather as a custom reverse proxy performing TLS termination and routing

Deployment on a DigitalOcean VPS with HTTPS support

<img width="1206" height="523" alt="image" src="https://github.com/user-attachments/assets/cdce5094-ae12-4cd9-8ca0-e50f5883499d" />

## License

This project is licensed under the [MIT License](./LICENSE).





