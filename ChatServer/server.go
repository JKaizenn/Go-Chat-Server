package main

import (
    // TODO: Add your imports here
    // Think about what you need: printing, networking, error handling
)

func startServer() {
    // TODO: Create TCP listener on port 8080
    // Hint: listener, err := net.Listen("tcp", ":8080")
    
    // TODO: Check for errors
    
    // TODO: Print that server is starting
    
    // TODO: Loop to accept connections
    // Hint: Use for loop with listener.Accept()
}

func handleConnection(/* TODO: what type for connection parameter? */) {
    // TODO: For now, just print that someone connected
    // Later we'll read messages from this connection
    
    // TODO: Don't forget to close the connection when done
    // Hint: defer conn.Close()
}