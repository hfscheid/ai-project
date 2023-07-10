from sys import stdin, stdout
import json
import os


PORT = 5555


def message_parser(line):
    temp_message = json.loads(line)

    file = open("/opt/bgp-listener/logs/log.txt", "a+")
    json.dump(temp_message, file)

    file.close()


while True:
    try:
        line = stdin.readline().strip()

        #
        if line == "":
            counter += 1
            if counter == 100:
                break
            continue
        counter = 0

        message = message_parser(line)

    except KeyboardInterrupt:
        pass
    except IOError:
        pass
