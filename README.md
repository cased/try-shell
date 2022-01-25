# Try Cased Shell

... in less than one minute.

## Local install with Docker Compose
### Usage

1. [Schedule a demo of Cased Shell](https://cased.com). Login to [Cased](https://app.cased.com) with the account created during the demo.
2. Click on the `localhost:NNNN` shell on your dashboard to view the demo shell.
3. Click 'Settings' and run the command listed. Alternatively, note the value for `CASED_SHELL_HOSTNAME`, `CASED_SHELL_SECRET`, and `AUTHORIZED_KEY` and then run:

```shell
curl https://raw.githubusercontent.com/cased/try-shell/main/run.sh | CASED_SHELL_HOSTNAME=<hostname> CASED_SHELL_SECRET=<secret> AUTHORIZED_KEY=<authorized key line> bash
```

Visit the URL output by that command to try out Cased Shell.

### Developing

If you'd prefer view the demo source before running it or if you'd like to contribute:

```shell
git clone https://github.com/cased/try-shell
cd try-shell
CASED_SHELL_HOSTNAME=<hostname> CASED_SHELL_SECRET=<secret> AUTHORIZED_KEY=<authorized key line> bash run.sh
```

### Notes

* Uploads and downloads are not currently available for local try-outs.

## Experimental: Heroku

1. Create a Cased Shell instance named `<something>.herokuapp.com`. On the Settings tab, enable Certificate Authentication.
2. Click this button (using Chrome):

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)