## HTTPs and the way it works: 

1. https://ericchiang.github.io/post/go-tls/ - learn by examples in Go
2. https://security.stackexchange.com/questions/20803/how-does-ssl-tls-work - great theoretical reference
3. https://www.digicert.com/ssl-cryptography.htm

### Description:

TLS - is the new name for SSL. Last widely supported version SSL 3.0 

What SSL is for: 

1. Encryption (data protection)
2. Identification (client server know end trust each other)
3. Data integrity (data alterations are detectable)

### SSL layers:

#### Bottom layer: 

Record protocol, data sent in SSL tunnel is split into records. Record looks like this: 
`HH V1:V2 L1:L2 data`

`HH` - single byte indicating the type of data in the record. There are four types: 

- change_cipher_spec
- alert 
- handshake
- application_data

`V1:V2` - protocol version - two bytes

V1 is alwys 0x03. V2 varies:

- V2=0x00 for SSLv3
- V2=0x01 for TLS1.0
- V2=0x02 for TLS1.1
- V2=0x03 for TLS1.2

`L1:L2` - is the length of data in bytes in big-endian, i.e. total is `256*L1+ L2` - limited by 18432 bytes

#### Records data: 

Payload is written and compressed using agreed compression algorithm (either null or Deflate compression)
Compressed payload is protected against alterations and encrypted as follows: 

`Final Payload = Encrypt(Raw Payload + MAC + Padding)`

Therefore the final record looks as follows: `HH V1:V2 L1:L2 Final Payload`

MAC - more on that https://en.wikipedia.org/wiki/Hash-based_message_authentication_code and https://en.wikipedia.org/wiki/Message_authentication_code

long story short MAC is added to the payload and it is encrypted. HMAC is calculated based on the data hash, therefore it is tied to the payload content. At the same time MAC is passed unencrypted with the message.
Receiver will decrypt the message and extract the MAC attached to payload and compare it to the  attached MAC to verify that the sender has the same secret key and also verify that the data was not altered.

### Steps:

#### Handshake:

1. Computers agree how to encrypt: client says "hello" message containing: 
    a. Key Exchange method (RSA)
    b. Cipher (AES, RC4 or Triple DES)
    c. Hash (HMAC-MD5 or SHA)
2. Client sends a,b,c and sends SSL version it supports (3.3) and some random number - which is used to compute master secret
3. Server responds with its own a,b,c
4. Server sends certificate:
    a. Issuer
    b. Public certificate
5. 

TO BE FILLED LATER

### RSA

Key exchange method 

>RSA: the server's key is of type RSA. The client generates a random value (the "pre-master secret" of 48 bytes, out of which 46 are random) and encrypts it with the server's public key. There is no ServerKeyExchange

The idea of Server generating Public/Private key pair. Public for encryption and private for decryption. Then client generates a random key which is encrypted with Public key of server(which is verified with CA, whose signature is shared along with Public certificate). The server decrypts using its private key and then encrypts a message using decrypted client generated random key. This random key is then used for further encryption.

### Golang usages

1. Create RSA keys

2. Creating digital signatures

3. Creating self-signed certs

4. Making client trust server

5. Making server trust client

6. Bonus (more flags and options)