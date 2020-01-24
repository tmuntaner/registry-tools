[![Build Status](https://travis-ci.org/tmuntaner/registry-tools.svg?branch=master)](https://travis-ci.org/tmuntaner/registry-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/tmuntaner/registry-tools)](https://goreportcard.com/report/github.com/tmuntaner/registry-tools)

# Registry Tools

Registry-Tools is a collection of tools to help work with docker registries.

## Usage

### docker-ls

`docker-ls` is a tool to browse a docker registry's repositories and tags.

#### docker-ls repositories

`docker-ls repositories [registry]` can be used to list all the repositories for a given registry.

**Note:**

* The official docker hub doesn't support the endpoint necessary to make this command work.

**Example:**

`docker-ls repositories registry.suse.com`

#### docker-ls tags

`docker-ls tags [repository]` can be used to browse the tags for a specific repository.

**Examples:**

* `docker-ls tags nginx`
* `docker-ls tags opensuse/leap`
* `docker-ls tags registry.suse.com/suse/sle15`
* `docker-ls tags suse/sle15 -r registry.suse.com`
