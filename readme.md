# IP over Doggo Carrier
IP over Doggo Carrier (IPoDC) is an innovative new technique based on [RFC1149](https://tools.ietf.org/html/rfc1149) (A Standard for the Transmission of IP Datagrams on Avian Carriers).
IPoDC has been tested and succesfully deployed within the Moore household.
The current implementation is simple, with packets being displayed on the screen in base64, at which point a human operator writes the data onto a sheet of paper attached to a dog, then alerts the other operator to call the dog through a verbal message like "Sending".
The dog will then run to the other operator after being called who can then type the packet into their computer to complete the transfer.
Using this method an ICMP ping was sent across the network over the subnet 10.1.0.0/24 with a device at 10.1.0.1 sending a packet to 10.1.0.2.
```
➜  ~ ping 10.1.0.2 -c 1 -W 100000 -s 16
PING 10.1.0.2 (10.1.0.2) 16(44) bytes of data.

64 bytes from 10.1.0.2: icmp_seq=1 ttl=64 time=67088 ms

--- 10.1.0.2 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 67087.569/67087.569/67087.569/0.000 ms
```
Packetloss was an non-existent with the packet only having an RTT of slightly over one minute (67087.569ms).
This RTT could be improved by a faster human operator and the use of a more efficient serialization method for transcription by the operator.

## Notes
* Can become costly rather if doggo requires treats for each transmission
  * Even with treats doggos can get annoyed with running between operators
* For fastest transcription during ICMP tests use a packetsize of 0 to decrease the packet length significantly
* Multicast traffic while possible, certainly does not sound fun and has not yet been tested
