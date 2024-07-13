#!/bin/sh

log_message_success() {
    local text="$1"
    echo -e "\e[32m$text\e[0m"
}

log_message_info() {
    local text="$1"
    echo -e "\e[34m$text\e[0m"
}

log_message_info "🔧 Fetching executable..."
wget https://github.com/brownhounds/migoro/releases/download/0.1.2/migoro-linux-arm64


log_message_info "🔧 Installing..."
chmod +x ./migoro-linux-arm64
bin_dir="$HOME/.local/bin"

# Check if the directory exists
if [ ! -d "$bin_dir" ]; then
    mkdir -p "$bin_dir"
fi

mv ./migoro-linux-arm64 $HOME/.local/bin/migoro

log_message_info 'Add folowing to .bashrc or .zshrc'
log_message_info 'export PATH="$HOME/.local/bin:$PATH"'
log_message_info 'Source the file or restart terminal session'

log_message_success '✨ All Done!'
