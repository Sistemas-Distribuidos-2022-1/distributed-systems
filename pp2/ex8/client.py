from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def special_credit_check(balance: str):
    # Checks if values are valid.
    try:
        float(balance)
    except ValueError:
        print("Invalid input!")
        return

    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request for readjustment
    client.send(("SCC " + balance).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    print("SPECIAL CREDIT AVAILABLE: " + resp)

    client.close()

if __name__ == "__main__":
    print("-------------------- Special Credit Checker --------------------")
    print("Element format: AVG_BALANCE")
    print("Multiple data input: NOT SUPPORTED!")
    print("type \'exit\' to close")

    while True:
        # Collect input data from terminal
        input_message = input("> ")
        input_message = input_message.replace("\n", "")
        input_message = input_message.replace("\r", "")

        if input_message == "exit":
            break

        # Passes to the processing function
        special_credit_check(input_message)