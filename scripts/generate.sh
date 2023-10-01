#!/bin/sh

FORCE=false

while getopts 'n:f' OPTION; do
    case "$OPTION" in
        n)
            NAME="$OPTARG"
            ;;
        f)
            FORCE=true
            ;;
        ?)
            ;;
    esac
done

if [[ -z $NAME ]]; then
    echo "missing kata name"
    exit 1
fi

TEMPLATE_PATH="./templates/$NAME.template"
KATA_PATH="./katas/$NAME.go"

if [[ -f $KATA_PATH && $FORCE == false ]]; then
    echo "$KATA_PATH already exists"
    exit 1
fi

cp -f $TEMPLATE_PATH $KATA_PATH
echo "generated $KATA_PATH from $TEMPLATE_PATH"

exit 0
