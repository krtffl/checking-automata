# checking-automata

checking automata is a small golang script to automate checking in using 
sesame time. 

## requirements

to allow a full automation, it requires an already logged in session
in sesame time. otherwise corporate login with two-factor auth
is required. this process is not considered yet, but i might add it

## set it up

### chrome

run a chrome instance  

 ```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --remote-debugging-port=9222 --no-first-run --no-default-browser-check --user-data-dir=$(mktemp -d -t chrome)
```

it will output its ws url, which is needed to attach to the running instance. perform the corporate login
and leave the window open (it might work if window is closed but process is not killed)

### mac os

to set it as a cron job, give full disk access permissions to cron in `System Preferences` > `Security & Privacy` > `Full Disk Access` > `+`

```bash
/usr/sbin/cron
```

compile the application

```bash
make build
```

and then add it on the crontab via

```bash
crontab -e
```

make sure to specify the path to the config file. also it is advised to generate the output into a log file


```bash
30 8 * * 1-5 /pth/to/your/file --config /pth/to/your/config >> /pth/to/logfile/ 2>&1
```





