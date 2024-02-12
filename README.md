# checking-automata

checking automata is a small golang script to automate checking in using 
sesame time. 

## requirements

to allow a full automation, it requires an already logged in session
in sesame time. otherwise corporate login with two-factor auth
is required

## set it up

### chrome

run a chrome instance  

 ```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --remote-debugging-port=9222 --no-first-run --no-default-browser-check --user-data-dir=$(mktemp -d -t chrome)
```

it will output its ws url, which is needed to attach to the running instance

### mac os

to set it as a cron job, give full disk access permissions to cron in 

/usr/sbin/cron

compile it 

```bash
make build
```

and then add it on the crontab

```bash
crontab -e
```

```bash
30 8 * * 1-5 /pth/to/your/file
```






