GOPATH=~/go

.PHONY: all
all: clean test build

update:
	gcloud components update

gae-local:
	GOPATH=$(GOPATH) dev_appserver.py app.yaml


gae-dev:
	cp config_gaeintrodev ~/.config/gcloud/configurations/config_gaeintrodev
	gcloud config configurations activate gaeintrodev
	GOPATH=$(GOPATH) gcloud app deploy default/app.yaml hello/app.yaml --stop-previous-version --promote --quiet
	gcloud app deploy cron.yaml --quiet
	gcloud app deploy --quiet default/dispatch.yaml
	gcloud app deploy --quiet hello/dispatch.yaml