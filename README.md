This code demonstrates the Command pattern in the Go programming language. The Command pattern is a behavioral design pattern that encapsulates a request as 
an object, thereby allowing for parameterization of clients with queues, requests, and operations.

Key Components:

Command Interface: The Command interface defines the Execute method, which is responsible for executing a specific command and returning an integer result.

Concrete Commands: There are two concrete command implementations, AddCommand and SubtractCommand. These classes encapsulate the logic for adding and subtracting numbers.

Calculator: The Calculator is an object that acts as the invoker and client. It has methods to set the command and execute it.
