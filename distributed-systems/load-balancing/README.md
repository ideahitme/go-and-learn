## Load Balancing

### Terms

List of useful terms needed to understand how load-balancing works:

1. Framing - https://en.wikipedia.org/wiki/Frame_(networking) data-link layer. Frame is a group of 1000s of bytes and allows grouping 8-bits into bytes by a framing protocol, i.e. finding right byte boundaries.  Example protocols: 

a. HDLC - high level data-link control

Flag: 01111110 <- when receiver sees this, it knows that next 8 bits make up a byte. Example: 

01001111110001111100111 -> 010 (ignored) | 01111110 (flag) | 00111110 (first byte) 

To avoid mistakes when a byte consists of 6 consecutive 1's, HDLC uses bit stuffing, any consecutive sequence of 5 1's is followed by 0

b. Ethernet protocol:

uses preamble 

___
Point to point connection, for example major hubs between cities, use PPP or HDLC protocol. Do not support forwarding based on MAC address
Multipoint - for example local area network with a hub switch, forwards local packets with source/destination being device mac addresses


2. VIP (virtual IP) - 

3. Linux Virtual Server (LVS) - linux kernel load-balancing of TCP/UDP connections on layer 4. Three ways of packet forwarding: 

a. NAT - packets are received from end users and the destination port and IP address are changed to that of the chosen real server. Return packets pass through the linux director at which time the mapping is undone so the end user sees replies from the expected source.

b. Direct Routing: Packets from end users are forwarded directly to the real server. The IP packet is not modified, so the real servers must be configured to accept traffic for the virtual server's IP address. This can be done using a dummy interface or packet filtering to redirect traffic addressed to the virtual server's IP address to a local port. The real server may send replies directly back to the end user. Thus, the linux director does not need to be in the return path.

c. IP-IP Encapsulation (Tunnelling): Allows packets addressed to an IP address to be redirected to another address, possibly on a different network. In the context of layer 4 switching the behaviour is very similar to that of direct routing, except that when packets are forwarded they are encapsulated in an IP packet, rather than just manipulating the ethernet frame. The main advantage of using tunnelling is that real servers can be on a different networks.
