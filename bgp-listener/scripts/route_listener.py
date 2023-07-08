from __future__ import print_function
from sys import stdout, stdin
from time import sleep

messages = [
    'announce route 100.10.0.0/24 next-hop self',
    'withdraw route 100.10.0.0/24 next-hop self'
]

sleep(5)

for message in messages: 
    stdout.write(message+'\n')
    stdout.flush()
    sleep(0.2)

while True:
    string = stdin.read()
    print(string, " read")
    sleep(1)