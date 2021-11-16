# Cased Shell on Heroku

1. Create a Heroku app. Its name will be referenced as `<app_name>` throughout this tutorial.
2. Create a Cased Shell instance named `<app_name>.herokuapp.com`. On the Settings tab, enable Certificate Authentication. Configure a host to allow access by connections signed by your certificate, then add the host to Cased Shell.
3. Add a `CASED_SHELL_SECRET` Config Var to your Heroku app using the value from the Cased Shell Settings tab.
4. Clone this repo.
5. Run the following commands to deploy Cased Shell to Heroku:

```
heroku git:remote -a <app_name>
heroku stack:set container
heroku labs:enable runtime-dyno-metadata
git push heroku master
```