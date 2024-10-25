#!/bin/sh

semanage port -a -t websm_port_t -p tcp 3389
firewall-cmd --service cockpit --permanent --add-port=3389/tcp
firewall-cmd --service cockpit --permanent --remove-port=9090/tcp
firewall-cmd --add-service=http --permanent
firewall-cmd --reload

systemctl daemon-reload

systemctl enable --now podman.socket
systemctl enable --now cockpit.socket
systemctl enable --now test.target
