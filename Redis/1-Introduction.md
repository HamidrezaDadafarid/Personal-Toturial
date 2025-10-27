# Redis

##  What is Redis?
Redis is an in-memory data store used by millions of developers as a cache, vector database, document database, streaming engine, and message broker. Redis has built-in replication and different levels of on-disk persistence. It supports complex data types (for example, strings, hashes, lists, sets, sorted sets, and JSON), with atomic operations defined on those data types.

## Installation
Add the repository to the APT index, update it, and install Redis Open Source:

```bash
sudo apt-get install lsb-release curl gpg
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
sudo chmod 644 /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
sudo apt-get update
sudo apt-get install redis
```

The most recent version of Redis Open Source will be installed, along with the redis-tools package (redis-cli, etc.). If you need to install an earlier version, run the following command to list the available versions:

```bash
apt policy redis

redis:
  Installed: (none)
  Candidate: 6:8.0.0-1rl1~bookworm1
  Version table:
     6:8.0.0-1rl1~bookworm1 500
        500 https://packages.redis.io/deb bookworm/main arm64 Packages
        500 https://packages.redis.io/deb bookworm/main all Packages
     6:7.4.3-1rl1~bookworm1 500
        500 https://packages.redis.io/deb bookworm/main arm64 Packages
        500 https://packages.redis.io/deb bookworm/main all Packages
     6:7.4.2-1rl1~bookworm1 500
        500 https://packages.redis.io/deb bookworm/main arm64 Packages
        500 https://packages.redis.io/deb bookworm/main all Packages
```

To install an earlier version, say 7.4.2, run the following command:

```bash
sudo apt-get install redis=6:7.4.2-1rl1~jammy1
```

Redis should start automatically after the initial installation and also at boot time. Should that not be the case on your system, run the following commands:

```bash
sudo systemctl enable redis-server
sudo systemctl start redis-server
```

