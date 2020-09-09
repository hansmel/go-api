#!/bin/bash

# verify input argument
if [ $# != 1 ]; then
printf "\nUsage: build.sh IMAGE NAME [example: build.sh hanmel/webservice]\n\n"
exit -1
fi

# set image name from input argument
imagename=$1

# read VERSION file
versionstr=$(cat VERSION)

# remove any white spaces
version="$(echo -e "${versionstr}" | tr -d '[:space:]')"

# verify correct version format
# capture all dots in the version variable
dotstr="${version//[^.]}"
if [ ${#dotstr} != 2 ]; then 
printf "\nBuild aborted - VERSION file has wrong format [example: 1.2.3]\n\n"
exit -1
fi

# segment the version string into its components
IFS='.'
read -a versionarr <<< "$version"
unset IFS

# populate version information
major="${versionarr[0]}"
minor="${versionarr[1]}"
patch="${versionarr[2]}"

printf "\nBuilding $imagename:$version\n\n"
docker build -t $imagename:$version .
printf "\nBuild completed\n\n"

printf "\nPushing $imagename:$version to docker hub\n\n"
docker push $imagename:$version
printf "\nPush completed\n\n"

exit 0
