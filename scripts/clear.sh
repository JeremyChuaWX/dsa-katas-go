#!/bin/sh

while getopts 'n:' OPTION; do
    case "$OPTION" in
        n)
            NAME="$OPTARG"
            ;;
        ?)
            ;;
    esac
done

KATA_PATH="./katas/"

if [[ $NAME ]]; then
    rm $KATA_PATH$NAME.go
    echo "cleared $NAME"
else
    rm $KATA_PATH*.go
    echo "cleared all katas"
fi

exit 0
