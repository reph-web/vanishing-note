# VANISHING NOTE

![image](https://github.com/user-attachments/assets/118d6585-6210-4db8-afde-486196a9bd40)


This quick project is an E2EE (End-to-End Encryption) Pastebin-like website built with Go using the Fiber framework. 
Each note created is encrypted with AES-128 and decrypted on the client side by a JavaScript script in the HTML. The key is stored in the URL fragment shared by the sender, so the server never has access to the key or the note's data, which is stored in an encrypted format in a MySql database.

Users can choose how long (from 5 minutes to 1 week) the note should be accessible. Once the note is viewed, it is automatically deleted from the database, or if an expired note is accessed, it is also deleted. To enhance security, a goroutine runs every 30 minutes to delete expired notes from the database that haven't been viewed.

To set up the project:
- Install Go 1.24.0 or newer.
- Clone the repository.
- In the project's folder, run `go mod tidy` to install the dependencies.
- In the `database/database.go` file, don't forget to modify the DSN string to match your own MySQL connection string.
- In the `main.go` file, the default listening port is `9000`. You can change it to any port of your choice.
- Run `go run main.go` to start the server.
