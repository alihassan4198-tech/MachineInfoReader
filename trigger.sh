if [[ "$OSTYPE" == "Linux"* ]];
then
    /bin/machine_info_gatherer /home/machineinfocsv/ > /home/machineinfolog/machineinfo.log
else
    # echo "$USER" > /Users/abdulrehman/machineinfolog/machineinfo.log
    # echo "$SHELL" >> /Users/abdulrehman/machineinfolog/machineinfo.log
    /Users/Shared/machine_info_gatherer /Users/Shared/machineinfocsv/ >> /Users/Shared/machineinfolog/machineinfo.log
fi 
# /Users/abdulrehman/machine_info_gatherer "/Users/abdulrehman/machineinfocsv/" > /Users/abdulrehman/machineinfolog/machineinfo.log