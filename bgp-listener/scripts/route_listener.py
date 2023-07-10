from sys import stdin, stdout
import json
import os


default_message = {
    "type": "N/A",
    "peer": "N/A",
    "atrORroute": "N/A",
    "next-hop": "N/A",
    "nlri": "N/A",
    "as-path": "N/A",
    "community": "N/A",
    "peerASN": "N/A"
}


def message_parser(line):
    temp_message = json.loads(line)

    # means its a message
    if "type" in temp_message:
        # means its an update message (announce or withdraw)
        if temp_message["type"] == "update":
            # if its an announce update
            if "update" in temp_message["neighbor"]["message"]:
                if "announce" in temp_message["neighbor"]["message"]["update"]:

                    msg = default_message.copy()
                    
                    peerIP = temp_message["neighbor"]["address"]["peer"]
                    peerASN = temp_message["neighbor"]["asn"]["peer"]
                    aspath = temp_message["neighbor"]["message"]["update"]["attribute"]["as-path"]["0"]["value"]
                    nlri_list = [i["nlri"] for i in temp_message["neighbor"]["message"]["update"]["announce"]["ipv4 unicast"][peerIP]]


                    if "community" in temp_message["neighbor"]["message"]["update"]["attribute"]:
                        community = temp_message["neighbor"]["message"]["update"]["attribute"]["community"][0]
                        msg["community"] = ":".join(str(e) for e in community)


                    msg["type"] = "Announcement"
                    msg["peer"] = peerIP
                    msg["peerASN"] = peerASN
                    msg["as-path"] = ", ".join(str(e) for e in aspath)
                    msg["next-hop"] = aspath[0]
                    msg["nlri"] = ", ".join(nlri_list)

                    file = open("/opt/bgp-listener/logs/log.txt", "a+")
                    file.write("\n")
                    file.write(f"{msg['type']}:\n\tPeer:\n\t\tAddr:\t{msg['peer']}\n\t\tAS:\t\t{msg['peerASN']}\n\tPrefix(es):\t{msg['nlri']}\n\tNext-hop:\t{msg['next-hop']}\n\tAS-path\t\t{msg['as-path']}\n\tCommunity:\t{msg['community']}\n\n")
                    file.close()

                # if its an withdraw update
                if "withdraw" in temp_message["neighbor"]["message"]["update"]:

                    msg = default_message.copy()

                    peerIP = temp_message["neighbor"]["address"]["peer"]
                    peerASN = temp_message["neighbor"]["asn"]["peer"]
                    nlri_list = [i["nlri"] for i in temp_message["neighbor"]["message"]["update"]["withdraw"]["ipv4 unicast"]]

                    msg["type"] = "Withdrawal"
                    msg["peer"] = peerIP
                    msg["peerASN"] = peerASN
                    msg["nlri"] = ", ".join(nlri_list)

                    file = open("/opt/bgp-listener/logs/log.txt", "a+")
                    file.write("\n")
                    file.write(f"{msg['type']}:\n\tPeer:\n\t\tAddr:\t{msg['peer']}\n\t\tAS:\t\t{msg['peerASN']}\n\tPrefix(es):\t{msg['nlri']}\n\n")
                    file.close()

file = open("/opt/bgp-listener/logs/log.txt", "w+")
file.write("")
file.close()

while True:
    try:
        line = stdin.readline().strip()

        # avoid loops
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
