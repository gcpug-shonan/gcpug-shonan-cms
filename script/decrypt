#!/bin/sh
root=`dirname $0`/..
secret_dir=$root/.secret
suffix='.enc'

for target in $(find $secret_dir -name "*$suffix" ! -name ".*" -type f); do
    file=$(echo $target | sed -e "s/.enc$//")
    if [ -e $file ]; then
        echo "Skip $file"
        continue
    fi
    key=$(echo $SECURE_KEY | openssl dgst -md5 |sed 's/^.* //')
    iv=$(echo ${key}${SECURE_KEY} | openssl dgst -md5 |sed 's/^.* //')
    openssl enc -aes-256-cbc -d -K "$key" -iv "$iv" -a -salt -in ${target} -out ${file}
    if [ $? -ne 0 ]; then
        rm -rf $file
    fi
done
