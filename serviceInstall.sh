echo "Deleting"
systemctl stop machine_info_server.service
rm /lib/systemd/system/machine_info_server.service


echo "Coping"
cp ./machine_info_server.service /lib/systemd/system/


echo "Starting"
systemctl enable machine_info_server.service
systemctl start machine_info_server.service
systemctl status machine_info_server.service
