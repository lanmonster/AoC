#! /bin/sh
NC='\033[0m' # No Color
RED='\033[0;31m'
GREEN='\033[0;32m'
PURPLE='\033[0;35m'

fail() {
    # shellcheck disable=SC2059
    printf "${RED}FAIL${NC}\n"
    exit
}
pass() {
    # shellcheck disable=SC2059
    printf "${GREEN}PASS${NC}\n"
    exit
}
cleanup() {
    # shellcheck disable=SC2317
    tput cnorm
}
tput civis
trap cleanup EXIT

jakt aoc.jakt || exit 1
clear

if [ -n "$1" ]
then
    ds="$1"
else
    # shellcheck disable=SC2125
    ds=./day*/
fi

HAS_ANY_FAILURES=false
for d in $ds
do 
    num="$(echo "$d" | tr -d -c "[:digit:]")"
    printf "[ ⏳ ] Day %d\r" "$num"
    hyperfine -i --export-json "$TMPDIR/d$num" "build/aoc $num input" 1>/dev/null 2> /dev/null
    if build/aoc "$num" input 1> /dev/null 2> /dev/null
    then
        RUN_FAILED=false
    else
        RUN_FAILED=true
        HAS_ANY_FAILURES=true
    fi

    if [ $RUN_FAILED = true ]
    then
        EMOJI="❌"
    else
        EMOJI="✅"
    fi
    
    printf "\r[ %s ] Day %d" "$EMOJI" "$num" 
    printf " -> $PURPLE%s$NC\n" "$(jq '.results' < "$TMPDIR/d$num" | jq -r '(.[0].mean*1000|round/1000|tostring) + " ± " + (.[0].stddev*1000|round/1000|tostring) + " s"')"
done

if [ $HAS_ANY_FAILURES = true ]
then
    fail
fi

pass

