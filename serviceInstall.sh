echo "Deleting"
systemctl stop machine_info_service.service
rm /lib/systemd/system/machine_info_service.service


echo "Coping"
cp ./machine_info_service.service /lib/systemd/system/


echo "Starting"
systemctl enable machine_info_service.service
systemctl start machine_info_service.service
systemctl status machine_info_service.service
