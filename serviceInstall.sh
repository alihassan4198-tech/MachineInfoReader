echo "Deleting"
systemctl stop machine_info_service.service
#rm /usr/bin/linux_machine_info_gatherer
rm /usr/bin/machine_info_gatherer
rm /lib/systemd/system/machine_info_service.service

echo "Coping"
cp ./machine_info_gatherer /usr/bin
cp ./machine_info_service.service /lib/systemd/system/

echo "Starting"
systemctl enable machine_info_service.service
systemctl start machine_info_service.service
systemctl status machine_info_service.service
