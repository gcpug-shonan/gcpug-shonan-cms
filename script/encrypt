#!/bin/sh
root=`dirname $0`/..
secret_dir=$root/.secret
suffix='.enc'

for target in $(find $secret_dir ! -name "*$suffix" ! -name ".*" -type f); do
    if [ -e "${target}${suffix}" ]; then
        rm -f ${target}${suffix}
    fi
    key=$(echo $SECURE_KEY | openssl dgst -md5 |sed 's/^.* //')
    iv=$(echo ${key}${SECURE_KEY} | openssl dgst -md5 |sed 's/^.* //')
    openssl enc -aes-256-cbc -K $key -iv $iv -a -salt -in ${target} -out ${target}${suffix}
    if [ $? -ne 0 ]; then
        rm -rf ${target}${suffix}
    fi

done
