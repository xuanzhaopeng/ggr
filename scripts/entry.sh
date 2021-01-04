#!/bin/sh

htpasswd -bc /etc/grid-router/users.htpasswd test $PASSWORD
echo "${DATA}" > /etc/grid-router/quota/test.xml

/usr/bin/ggr -listen :4444 -users /etc/grid-router/users.htpasswd -quotaDir /etc/grid-router/quota
