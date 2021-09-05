import socket

HOST = 'localhost'  # The server's hostname or IP address
PORT = 8080         # The port used by the server

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall(b'Hello, world\n')
    data = s.recv(1024)

print('Received', repr(data))