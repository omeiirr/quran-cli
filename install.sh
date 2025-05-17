#!/bin/sh
# install-yourapp.sh

VERSION="v1.0.0"
OS=$(uname | tr '[:upper:]' '[:lower:]')

echo $OS

DL_LINK="https://github.com/omeiirr/quran-cli/releases/download/$VERSION/quran-$OS.zip"
echo $DL_LINK
curl -LO $DL_LINK

unzip quran-$OS.zip
sudo mv quran /usr/local/bin/
sudo chmod +x /usr/local/bin/quran

echo "Successfully installed! Run with: quran"
