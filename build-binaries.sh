#!/bin/bash

GO_DETECT=1

go version || GO_DETECT=0


if [ $GO_DETECT -eq 0 ];
then 
    return
fi

echo "Deleting old files..."

rm ./machine_info_gatherer
rm ./server


echo "Building Binaries..."

cd ./machine_info_reader && go build . && mv machine_info_gatherer .. && cd .. || cd ..
cd ./machine_info_server && go build . && mv server .. && cd .. || cd ..
