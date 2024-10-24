#!/bin/sh

systemctl daemon-reload
systemctl enable --now podman.socket
systemctl enable --now test.target
