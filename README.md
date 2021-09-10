# Try Cased Shell

... in less than one minute.

## Usage

> :construction: :warning: This approach will work **once this repo is open sourced**. :construction: :warning:

1. Login to https://app.cased.com.
2. Click on the `localhost` shell on your dashboard to view the demo shell.
3. Click 'Settings' and run the command listed. Alternatively, note the value for `CASED_SHELL_HOSTNAME`, `CASED_SHELL_SECRET`, and `AUTHORIZED_KEY` and then run:

```shell
curl https://raw.githubusercontent.com/cased/try-shell/main/run.sh | CASED_SHELL_HOSTNAME=<hostname> CASED_SHELL_SECRET=<secret> AUTHORIZED_KEY=<authorized key line> bash
```

Visit the URL output by that command to try out Cased Shell.

> :construction: :warning: Use this approach for now. :construction: :warning:

If you'd prefer to view the demo source before running it:

```shell
git clone https://github.com/cased/try-shell
cd try-shell
CASED_SHELL_HOSTNAME=<hostname> CASED_SHELL_SECRET=<secret> AUTHORIZED_KEY=<authorized key line> bash run.sh
```