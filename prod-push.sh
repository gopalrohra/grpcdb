#!/bin/bash
set -e
cd ~/projects/grpcdb/
mkdir -p build
echo "building golang app for production .................."
cp .env* build/
GOOS=linux GOARCH=amd64 go build -o build/ .
ts=`date +%Y%m%d%H%M%S`
PROD_APP_DIR="~/deployments/grpcdb"
echo "Creating initial directories on production................"
ssh caretaker@appshome.in << EOF
echo "Connected successfully to appshome.in."
set -e
echo "Creating directory structure for artifacts................"
mkdir -p $PROD_APP_DIR/build
EOF
echo "Shipping artifacts to production."
scp build/grpcdb caretaker@appshome.in:$PROD_APP_DIR/build/

echo "Connecting to production server."
ssh caretaker@appshome.in << EOF
echo "Connected successfully to appshome.in."
set -e
echo "Coppying the artifacts."
cp $PROD_APP_DIR/build/grpcdb $PROD_APP_DIR/grpcdb-$ts
echo "Stopping the service.."
systemctl --user stop grpcdb.service
echo "Configuring the deployed directory for production ..................."
rm -f $PROD_APP_DIR/current
ln -s $PROD_APP_DIR/grpcdb-$ts $PROD_APP_DIR/current
echo "Starting the newly deployed artifact."
systemctl --user start grpcdb.service
echo "Removing the artifacts."
rm -rf $PROD_APP_DIR/build
echo "Deployment successful ............."
echo "Terminating the connection to appshome.in."
EOF
echo "Cleaning up the artifacts locally."
rm -rf build
