#!/bin/sh

sudo crontab -l > cron_bkp
cp ./machine_info_gatherer /usr/bin
# sudo crontab cron_bkp
# sudo rm cron_bkp