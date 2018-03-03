# What is kube-ingwatcher ?

kube-ingwatcher acts as the following:
  - ingressSender listens for new ingresses (kubernetes' in-cluster container)
  - ingressReceiver receives what ingressSender sends

Ingress sender json-marshals ingress information and sends it to a backend.
When received, the backend can generate a file from a given template and run a given script.

I use it to automatically configure an nginx config for that ingress and to generate a Let's Encrypt certificate that fits domains as soon as an ingress appears on a cluster (and have a certain ingress.class and/or specific labels).

# Usage

```
$ /tmp/kube-ingwatcher --help

Usage: kube-ingwatcher [OPTIONS] COMMAND [arg...]

Acts as ingress sender or receiver and makes actions

Options:
  -v, --version         Show the version and exit

Commands:
  ingressSender, is     Ingress sender mode
  ingressReceiver, ir   Ingress receiver mode

Run 'kube-ingwatcher COMMAND --help' for more information on a command.
```

```
$ /tmp/kube-ingwatcher ingressSender --help

Usage: kube-ingwatcher ingressSender [OPTIONS]

Ingress sender mode

Options:
  -a, --addr   Destination address to send ingress to
  -p, --port   Default port to connect to (default 6565)
```

```
$ /tmp/kube-ingwatcher ingressReceiver --help

Usage: kube-ingwatcher ingressReceiver [OPTIONS]

Ingress receiver mode

Options:
  -a, --addr          Interface address to bind (default "0.0.0.0")
  -p, --port          Port to bind (default 6565)
  -t, --template      Template to render when an ingress is received
  -d, --destination   Destination directory where to generate templates
  -p, --prefix        Prefix to use for filename when generating a template
  -s, --suffix        Suffix to use for filenme when generating a template
  -A, --addCmd        Command to execute after rendering a file from the template
  -D, --deleteCmd     Command to execute after deleting a file previously created with the template
```

# Scripts / examples

When receiving a new ingress, creates a certificate using certbot :
```
#!/bin/bash

DOMAINS="$1"
CMD_DOMS=""
for dom in $(echo ${DOMAINS} | tr -s "," " "); do
  CMD_DOMS="${CMD_DOMS} -d ${dom}"
done

sudo certbot certonly --test-cert --agree-tos --register-unsafely-without-email -n --standalone --expand --preferred-challenges http ${CMD_DOMS}
sudo systemctl reload nginx
```

When an ingress is deleted, just reload nginx :
```
#!/bin/bash

sudo systemctl reload nginx
```
