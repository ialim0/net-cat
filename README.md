# net-cat
This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.
# TCPChat - NetCat-Like Chat Application in Go

TCPChat is a versatile chat application that reimagines the classic NetCat utility in a Server-Client architecture. This project allows you to run a server in listening mode on a specified port, accepting incoming connections from multiple clients. Clients can join the chat, exchange messages, and disconnect gracefully, all while maintaining a seamless group chat experience. TCPChat is designed to be an educational project, helping you explore various aspects of Go programming, network communication, and concurrency.

## Features

1. **TCP Connection**: TCPChat facilitates communication between a central server and multiple clients, establishing a one-to-many relationship.

2. **Client Naming**: Each client must provide a unique name upon connection, ensuring proper identification within the chat.

3. **Controlled Connections**: The server can handle a maximum of 10 connections, preventing overload and ensuring smooth operation.

4. **Message Exchange**: Clients can send text messages to the chat room. Empty messages are ignored.

5. **Message Timestamping**: Sent messages are identified by their timestamp and the sender's name in the format `[YYYY-MM-DD HH:MM:SS][client.name]: [client.message]`.

6. **Message History**: New clients joining the chat receive all previously sent messages, enabling them to catch up on the conversation.

7. **Client Join Notification**: When a client connects to the server, all other clients are informed that a new member has joined the group.

8. **Client Exit Notification**: When a client disconnects, the server notifies the remaining clients about the departure.

9. **Message Broadcast**: Messages sent by one client are received by all other clients in the chat.

10. **Graceful Client Exit**: If a client leaves the chat, the remaining clients continue without disruption.

11. **Default Port**: If no port is specified, the default port is set to 8989.

## Project Structure

The TCPChat project is written in the Go programming language and follows best practices for Go development. It incorporates Go-routines, channels, and mutexes to handle concurrency and ensure thread safety.
## Learning Opportunities
This project offers valuable insights into various aspects of programming, networking, and Go development:

Manipulation of data structures.
Understanding and implementing NetCat-like functionality.
TCP/UDP communication and socket handling.
Concurrency in Go, including goroutines and channels.
Ensuring thread safety with mutexes.
Handling errors gracefully on both the server and client sides.

Author
This project was created as an educational exercise by [ialimoud & ader]. Feel free to explore the code, experiment with the features, and contribute to its development.
## Learning Opportunities
This project offers valuable insights into various aspects of programming, networking, and Go development:

Manipulation of data structures.
Understanding and implementing NetCat-like functionality.
TCP/UDP communication and socket handling.
Concurrency in Go, including goroutines and channels.
Ensuring thread safety with mutexes.
Handling errors gracefully on both the server and client sides.

Author
This project was created as an educational exercise by [ialimoud & ader]. Feel free to explore the code, experiment with the features, and contribute to its development.


## Usage

To start the TCPChat server, use the following commands:

```sh
$ go run .
Listening on port :8989
$ go run . 2525
Listening on port :2525
$ nc $IP $port
$ nc localhost 2525
Welcome to TCP-Chat!
[ENTER YOUR NAME]: Yenlik
[2020-01-20 16:03:43][Yenlik]: hello
[2020-01-20 16:03:46][Yenlik]: How are you?

