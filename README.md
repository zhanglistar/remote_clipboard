# remote_clipboard
Remote clipboard, send text to remote mac machine from another machine.
Using UDP protocol written in Golang.

Compile
1. compile the three go files.

Usage
1. In remote server, start proxy
nohup ./proxy &

2. In local mac, start receiver,
nohup ./udp_server &

3. In remote server, send text to local mac, using
echo "test" | geter
