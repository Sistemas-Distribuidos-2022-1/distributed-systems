from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def readjust(name, role, salary):
    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request for readjustment
    client.send(("RAJ " + name + " " + role + " " + salary).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    print("NAME: " + name + " | ROLE: " + role + " | SALARY: " + salary + " -> " + resp)
    client.close()

if __name__ == "__main__":
    print("-------------------- Employee Salary Readjustment --------------------")
    print("Element format: NAME ROLE SALARY")
    print("Multiple data input: NAME1 ROLE1 SALARY1 NAME2 ROLE2 SALARY2 ...")
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
        if len(data)%3 != 0:
            print("Invalid input format")
            continue
        
        # Counts the number o employees present in the input data.
        numEmployees = len(data)//3
        for i in range(numEmployees):
            name = data[i*3]
            role = data[i*3+1]
            salary = data[i*3+2]

            if float(salary):
                # Passes to the processing function
                readjust(name, role, salary)