#!/bin/bash

# 设置远程服务器和路径
remote_path="/data/program/app/galactus/SignIos"
# 建立SSH连接并执行远程命令
sshpass -p "$blade_password" ssh -o StrictHostKeyChecking=no -T "$blade_remote_server" << EOF
  mkdir -p $remote_path
  rm -rf $remote_path/*.sh
EOF

sshpass -p "$blade_password" scp -p ./start.sh "$blade_remote_server:$remote_path"
sshpass -p "$blade_password" scp -p ./stop.sh "$blade_remote_server:$remote_path"
