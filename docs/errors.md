# Popular errors
## protocol version not supported
```
Post https://paste.lcomrade.su/api/new: remote error: tls: protocol version not supported
```

You are probably using a GO 1.12 or lower.
This GO version does not support the latest versions of the TLS protocol.

There are several solutions to the problem:
1. Update your GO.
2. Switch the server to use the HTTP instead of HTTPS.
3. Compile the server using an older version of your GO.
4. Set up a reverse proxy and make requests to it.
