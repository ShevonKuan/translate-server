# for linux x86_64
install_deepl(){
    last_version=$(curl -Ls "https://api.github.com/repos/ShevonKuan/deepl-server/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    if [[ ! -n "$last_version" ]]; then
        echo -e "${red}Failed to detect DeepL version, probably due to exceeding Github API limitations.${plain}"
        exit 1
    fi
    echo -e "DeepL latest version: ${last_version}, Start install..."
    wget -N --no-check-certificate -O /usr/bin/deepl https://github.com/ShevonKuan/deepl-server/releases/download/${last_version}/DeepLServer-linux-amd64


    chmod +x /usr/bin/deepl
    wget -N --no-check-certificate -O /etc/systemd/system/deepl.service https://raw.githubusercontent.com/ShevonKuan/deepl-server/main/deepl.service
    systemctl daemon-reload
    systemctl enable deepl
    systemctl start deepl
    echo -e "Installed successfully, listening at 0.0.0.0:1188"
}
install_deepl