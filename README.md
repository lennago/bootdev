# bootdev

A collection of projects completed as part of my coursework on [Boot.dev](https://www.boot.dev), a backend-development learning platform. Each folder is a self-contained project from a different course, mostly written in Go and Python.

## Projects

| Project | CI | Description |
| --- | --- | --- |
| [`chirpy`](./chirpy) | | HTTP web server built from scratch in Go, including routing, JSON handling, and a REST API backed by a database. |
| [`pokedexcli`](./pokedexcli) | | A CLI Pokedex in Go that consumes the PokeAPI REST API, with caching and a REPL-style interface. |
| [`gator`](./gator) | | RSS feed aggregator CLI in Go with persistent storage (PostgreSQL) and scheduled feed fetching. |
| [`http-from-tcp`](./http-from-tcp) | | Implements an HTTP server from raw TCP sockets in Go, without using the standard `net/http` package, to understand the protocol at a low level. |
| [`learn-pub-sub`](./learn-pub-sub) | | Pub/sub messaging system in Go using RabbitMQ, built around a multiplayer game simulation. |
| [`kubernetes`](./kubernetes) | | Deploying and orchestrating an application with Kubernetes. |
| [`static-site-generator`](./static-site-generator) | | A static site generator written in Python that converts Markdown files into HTML pages. |
| [`asteroids`](./asteroids) | | A recreation of the classic Asteroids arcade game in Python using Pygame. |
| [`bookbot`](./bookbot) | | A Python CLI tool that analyzes text files and reports word counts and character frequency statistics. |
| [`ai-agent`](./ai-agent) | | A toy AI coding agent in Python that uses an LLM (via function calling) to read, write, and execute files in a sandboxed directory. |
| [`learn-cicd`](./learn-cicd) | ![learn-cicd](https://github.com/lennago/bootdev/actions/workflows/ci.yml/badge.svg) | Notely — a Go web server covering building, testing, and deploying with GitHub Actions and Google Cloud Run. |

## About

These projects were built to practice backend fundamentals: HTTP, APIs, databases, Git, Linux, and basic AI tooling, primarily in Go and Python.
