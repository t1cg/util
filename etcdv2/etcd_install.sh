#!/bin/bash

ETCD=etcd
ETCD_VER=v3.1.0
DOWNLOAD_URL=https://github.com/coreos/etcd/releases/download
INSTALL_DIR=/opt/$ETCD/$ETCD_VER
DOWNLOAD_DIR=/opt/$ETCD/$ETCD_VER/download
UNZIP_DIR=/opt/$ETCD/$ETCD_VER/unzip

echo "starting..."
echo "-->creating folders..."
sudo mkdir -p $UNZIP_DIR;
sudo mkdir -p $DOWNLOAD_DIR;
echo "-->downloading..."
sudo curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-darwin-amd64.zip -o $DOWNLOAD_DIR/temp.zip;
echo "-->installing..."
sudo unzip $DOWNLOAD_DIR/temp.zip -d $UNZIP_DIR && sudo mv $UNZIP_DIR/etcd-${ETCD_VER}-darwin-amd64/* $INSTALL_DIR

echo "-->removing temp file..."
sudo rm -rf ${DOWNLOAD_DIR};
sudo rm -rf ${UNZIP_DIR};

echo "-->testing..."
$INSTALL_DIR/etcd --version;

echo "exiting"
