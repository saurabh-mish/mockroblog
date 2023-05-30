[![Tests](https://github.com/saurabh-mish/mockroblog/actions/workflows/test.yml/badge.svg)](https://github.com/saurabh-mish/mockroblog/actions/workflows/test.yml) &emsp;&emsp; [![Audit](https://github.com/saurabh-mish/mockroblog/actions/workflows/audit.yaml/badge.svg)](https://github.com/saurabh-mish/mockroblog/actions/workflows/audit.yaml) &emsp;&emsp; [![codecov](https://codecov.io/gh/saurabh-mish/mockroblog/branch/main/graph/badge.svg?token=BHAOSMITWR)](https://codecov.io/gh/saurabh-mish/mockroblog)

# Mockroblog

This project has services for a web application similar to Reddit.

It was a group project for [CPSC 449][1] in Spring 2020. A group of three students, were assigned a role of either *Dev*, *Test*, or *Ops*. After each major release, we took on a different role.

This project was initially implemented using Python’s Flask framework. In this re-visit I'm implementing it in Go.

### Whats New?

+ This project is in vanilla Go, except database drivers and experimental Golang packages

+ Production-oriented stack:

  1. NGINX as the web server instead of Caddy and Gunicorn
  2. Postgres as the database server instead of SQLite

+ All deployments will be on a local multi-node kubernetes cluster instead of a Linux VM / server

+ Impletation of application security and cluster-level security

+ Monitoring with Prometheus and Grafana

+ Improved testing and documentation

[1]: https://catalog.fullerton.edu/preview_course_nopop.php?catoid=61&coid=447756
