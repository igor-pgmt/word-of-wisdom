# Word of wisdom

Implemented PoW algorithm is a simple hash-based approach. 

- Challenge Generation: The server generates a random challenge string. This challenge is sent to the client.

- Proof Finding: The client must find a proof that, when concatenated with the challenge string and hashed using SHA-256, produces a hash with a specified number of leading zeros (defined by the difficulty level).

- Proof Validation: The server receives the proof from the client and validates it by hashing the challenge and the proof together. If the resulting hash has the required number of leading zeros, the proof is considered valid.

## Prerequisites
Install docker and docker-compose v2:
```bash
sudo apt-get install docker.io docker-compose-v2
```

## Build and run:
In project root directory:
```bash
make build-run
```

Server and client will be created, then the client will connect to the server every second to get a citate:
```bash
 ✔ Container word-of-wisdom-server-1  Created                                                            0.0s 
 ✔ Container word-of-wisdom-client-1  Created                                                            0.0s 
Attaching to client-1, server-1
server-1  | time="2024-06-29T14:19:13Z" level=info msg="Server listening on 0.0.0.0:8080"
client-1  | time="2024-06-29T14:19:14Z" level=info msg="Word of Wisdom: Have no fear of perfection you'll never reach it."
client-1  | time="2024-06-29T14:19:15Z" level=info msg="Word of Wisdom: See driving is like stabbing someone, it's very personal. While flying is like shooting someone, it's more distant."
```

