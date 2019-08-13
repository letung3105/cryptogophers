CryptoGophers
===

Solving cryptopals challenges, found [here](https://cryptopals.com/), using Golang


Set 1: Basics
---

* [x] Convert hex to base64
* [x] Fixed XOR
* [x] Single-byte XOR cipher
* [x] Detect single-character XOR
* [x] Implement repeating-key XOR
* [x] Break repeating-key XOR
* [x] AES in ECB mode
* [x] Detect AES in ECB mode


Set 2: Block crypto
---

* [x] Implement PKCS#7 padding
* [x] Implement CBC mode
* [x] An ECB/CBC detection oracle
* [x] Byte-at-a-time ECB decryption (Simple)
* [x] ECB cut-and-paste
* [x] Byte-at-a-time ECB decryption (Harder)
* [ ] PKCS#7 padding validation
* [ ] CBC bitflipping attacks


Set 3: Block and stream crypto
---

* [ ] The CBC padding oracle
* [ ] Implement CTR, the stream cipher mode
* [ ] Break fixed-nonce CTR mode using substitutions
* [ ] Break fixed-nonce CTR statistically
* [ ] Implement the MT19937 Mersenne Twister RNG
* [ ] Crack an MT19937 seed
* [ ] Clone an MT19937 RNG from its output
* [ ] Create the MT19937 stream cipher and break it


Set 4: Stream crypto and randomness
---

* [ ] Break "random access read/write" AES CTR
* [ ] CTR bitflipping
* [ ] Recover the key from CBC with IV=Key
* [ ] Implement a SHA-1 keyed MAC
* [ ] Break a SHA-1 keyed MAC using length extension
* [ ] Break an MD4 keyed MAC using length extension
* [ ] Implement and break HMAC-SHA1 with an artificial timing leak
* [ ] Break HMAC-SHA1 with a slightly less artificial timing leak


Set 5: Diffie-Hellman and friends
---

* [ ] Implement Diffie-Hellman
* [ ] Implement a MITM key-fixing attack on Diffie-Hellman with parameter injection
* [ ] Implement DH with negotiated groups, and break with malicious "g" parameters
* [ ] Implement Secure Remote Password (SRP)
* [ ] Break SRP with a zero key
* [ ] Offline dictionary attack on simplified SRP
* [ ] Implement RSA
* [ ] Implement an E=3 RSA Broadcast attack


Set 6: RSA and DSA
---

* [ ] Implement unpadded message recovery oracle
* [ ] Bleichenbacher's e=3 RSA Attack
* [ ] DSA key recovery from nonce
* [ ] DSA nonce recovery from repeated nonce
* [ ] DSA parameter tampering
* [ ] RSA parity oracle
* [ ] Bleichenbacher's PKCS *5 Padding Oracle (Simple Case)
* [ ] Bleichenbacher's PKCS *5 Padding Oracle (Complete Case)


Set 7: Hashes
---

* [ ] CBC-MAC Message Forgery
* [ ] Hashing with CBC-MAC
* [ ] Compression Ratio Side-Channel Attacks
* [ ] Iterated Hash Function Multicollisions
* [ ] Kelsey and Schneier's Expandable Messages
* [ ] Kelsey and Kohno's Nostradamus Attack
* [ ] MD4 Collisions
* [ ] RC4 Single-Byte Biases


Set 8: Abstract Algebra
---

* [ ] Diffie-Hellman Revisited: Small Subgroup Confinement
* [ ] Pollard's Method for Catching Kangaroos
* [ ] Elliptic Curve Diffie-Hellman and Invalid-Curve Attacks
* [ ] Single-Coordinate Ladders and Insecure Twists
* [ ] Duplicate-Signature Key Selection in ECDSA (and RSA)
* [ ] Key-Recovery Attacks on ECDSA with Biased Nonces
* [ ] Key-Recovery Attacks on GCM with Repeated Nonces
* [ ] Key-Recovery Attacks on GCM with a Truncated MAC


TODOs
----

* Reduce verbosity? (panic everything)
