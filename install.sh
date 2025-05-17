#!/bin/sh
# install-yourapp.sh

VERSION="v1.0.0"
OS=$(uname | tr '[:upper:]' '[:lower:]')

echo $OS
curl -LO "https://github.com/omeiirr/quran-cli/releases/download/$VERSION/quran-$OS.zip"

unzip quran-$OS.zip
sudo mv quran /usr/local/bin/
sudo chmod +x /usr/local/bin/quran

echo "✅ Installed! Run with: quran"

