from __future__ import print_function
from sys import stdout
from time import sleep

messages = [
    'announce route 100.0.0.1/33 next-hop self'
    'announce route '
]

sleep(5)

for message in messages: 
    stdout.write(message+'\n')
    stdout.flush()
    sleep(0.2)

while True:
    sleep(1)