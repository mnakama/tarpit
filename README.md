# tarpit
A simple HTTP [tarpit](https://wikipedia.org/wiki/Tarpit_(networking)) for catching PHP exploit scanners.
I built it when I saw bots constantly scanning my site for PHP exploits. It keeps the HTTP connection open
and sends a message every second to the scanner. Poorly programmed scanning bots will hang for a while.

## Usage
Configure the constants as you see fit, then `go build` it.

```
maxTar      Maximum tar count before we let them go
message     Message that repeats while we tarpit them
messageEnd  Message we show when we free them from the tarpit
```

### Nginx config
Add a section like this to your nginx.conf in the virtual servers where you want to tarpit:
```
location ~ \.php$ {
  proxy_pass http://localhost:8081;
  proxy_buffering off;
  proxy_set_header X-Forwarded-For $remote_addr;
}
```

You can change the location to tarpit anything you like. Note that tarpitting will use up a bit of resources on your server.
If it's a concern, you should do an sshguard-style firewall instead.
