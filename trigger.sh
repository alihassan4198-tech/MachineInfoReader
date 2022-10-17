#!/bin/bash
date

os_type=$(echo $OSTYPE | tr '[:upper:]' '[:lower:]')


if [[ "$os_type" == *"linux"* ]];
then
    /bin/machine_info_gatherer /home/machineinfocsv/ > /home/machineinfolog/machineinfo.log
else
    /Users/Shared/machine_info_gatherer /Users/Shared/machineinfocsv/ > /Users/Shared/machineinfolog/machineinfo.log
fi 