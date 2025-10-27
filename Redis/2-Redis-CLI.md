# Redis-CLI

## Host, port, password, and database 
By default, redis-cli connects to the server at the address 127.0.0.1 with port 6379. You can change the port using several command line options. To specify a different host name or an IP address, use the -h option. In order to set a different port, use -p.

```bash
$ redis-cli -h redis15.localnet.org -p 6390 PING
PONG
```

If your instance is password protected, the -a <password> option will perform authentication saving the need of explicitly using the AUTH command:

```bash
$ redis-cli -a myUnguessablePazzzzzword123 PING
PONG
```
**NOTE**: For security reasons, provide the password to redis-cli automatically via the `REDISCLI_AUTH` environment variable.

