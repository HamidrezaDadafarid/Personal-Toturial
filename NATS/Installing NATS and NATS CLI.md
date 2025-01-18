## Installing NATS server
1. download the nats-server binary file:
```bash
curl -sf https://binaries.nats.dev/nats-io/nats-server/v2@latest | sh
```
2. move the binary to /usr/local/bin path:
```bash
sudo mv nats-server /usr/local/bin
```
3. You can start it using:
```bash
nats-server
```

## Installing NATS CLI
1. Download the .deb file from the NATS CLI [GitHub releases page](https://github.com/nats-io/natscli/releases):
```bash
wget https://github.com/nats-io/natscli/releases/download/v0.1.6/nats-X.Y.Z-386.deb
```
Replace `X.Y.Z` with the version number of the release you want.

2. Install the .deb package:
```bash
sudo dpkg -i nats_X.Y.Z_amd64.deb
```
3. Fix any missing dependencies (if required):
```bash
sudo apt-get install -f
```
4. Verify the installation:
```bash
nats --version
```

