#!/bin/sh

systemctl stop podman.socket
systemctl disable podman.socket

systemctl stop cockpit.socket
systemctl disable cockpit.socket

systemctl stop test.target
systemctl disable test.target

firewall-cmd --service cockpit --permanent --remove-port=3389/tcp
firewall-cmd --service cockpit --permanent --add-port=9090/tcp
firewall-cmd --remove-service=http --permanent
firewall-cmd --reload
semanage port -d -t websm_port_t -p tcp 3389
