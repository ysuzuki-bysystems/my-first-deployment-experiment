#!/bin/sh

systemctl stop podman.socket
systemctl disable podman.socket

systemctl stop test.target
systemctl disable test.target
