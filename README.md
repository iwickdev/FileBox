# FileBox
Simple but usable file upload service that is a easy and self hostable way to host files on a server yourself

## Usage
- Install GoLang from [https://go.dev/dl/](https://go.dev/dl/)
- Move into this folder and run `go build .` to compile an executable for your operating system
- Make sure the `accets` folder is located in the directory with the executable (this contains html files and images)
- Run the program and access your new FileBox instance!

## Adding Users
- Find the `Users.conf` and use this format `Username=Password`, one per line. Then restar the program for the new user to take affect

## Change Host IP/Port
- Find the `FileBox.conf` file and replace the `Host=127.0.0.1:80` to whatever you want to use to host the service. Use `IP:PORT` format

## Change Storage Location
- Find the `FileBox.conf` file and replace the `StoragePath=storage` to a location of your choosing after the `=` sign
