# Handin3DSYS
# How to run the program
It is necesarry to start the server before running a client. This is done by entering the repository with the server.go - file, then use the 'go run' command

Then you can start launching clients by using the same 'go run' command on client.go.

Once the client is up and running, it will ask you to write your username. Type it and then press enter

Now the client is a part of the server. A list of commands will be provided in the terminal. You can now choose to write a message or use one of the following commands:

/profile - shows your name and your userID

/quit - leaves the server and the client termninates

/help - shows the list of the possible commands

/lt - shows the current Lamport Time Stamp for the client.

To use the commands just type it in the terminal and then press enter

If you just want to send a message type anything you want in the terminal and press enter.
Now the message will be send to the server distributed to all currently active clients.

In the terminal you will recieve messages for when people join, leave or message that they just sent. This happens basically as soon as the action occurs.

Once you are done, use the aformentioned /quit command to leave and terminate the operation.





 
