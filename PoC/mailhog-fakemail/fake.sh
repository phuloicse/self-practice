#!/bin/bash
set -e
set -x
# Create /etc/mailname
echo "fake-vegas.local" > /etc/mailname

# Configure Postfix
postconf -e "myhostname = fake-vegas.local"
postconf -e "myorigin = /etc/mailname"
postconf -e "inet_interfaces = loopback-only"
postconf -e "inet_protocols = all"
postconf -e "mydestination = localhost"

# Start postfix
service postfix start

# Give it time to start
sleep 3

# Send email
echo "Welcome to Vegas!\nThis is a spoofed test email." | mail -s "Vegas Deals" -a "From: VCB@myapp.com" VCB@myapp.com

# Keep container running to view logs
tail -f /var/log/mail.log
