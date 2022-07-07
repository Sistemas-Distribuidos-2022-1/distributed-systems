from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def retirement_check(age: str, service: str):
    # Checks if values are valid.
    try:
        int(age)
        int(service)
    except ValueError:
        print("Invalid input!")
        return

    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request for readjustment
    client.send(("RAC " + age + " " + service).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    if resp == "TRUE":
        print("Available!")
    elif resp == "FALSE":
        print("Unavailable!")
    else:
        print("It was not possible to check retirement availability for " + age + ", " + service)

    client.close()

if __name__ == "__main__":
    print("-------------------- Retirement Availability Calculator --------------------")
    print("Element format: AGE SERVICE_TIME")
    print("Multiple data input: NOT SUPPORTED!")
    print("type \'exit\' to close")

    while True:
        # Collect input data from terminal
        input_message = input("> ")
        input_message = input_message.replace("\n", "")
        input_message = input_message.replace("\r", "")

        if input_message == "exit":
            break

        # Check that the input string has the correct format
        data = input_message.split(" ")
        if len(data) != 2:
            print("Invalid input!")
            continue
        
        # Passes to the processing function
        retirement_check(data[0], data[1])