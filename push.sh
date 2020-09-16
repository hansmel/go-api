#!/bin/bash

# verify input argument
if [ $# != 1 ]; then
printf "\nUsage: push.sh IMAGE NAME [example: push.sh hanmel/webservice]\n\n"
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
printf "\nPush aborted - VERSION file has wrong format [example: 1.2.3]\n\n"
exit -1
fi

# verify that the image exist in local image repository
imageid=$(docker images -q $imagename:$version)
if [ ${#imageid} != 12 ]; then 
  printf "\nPush aborted - image $imagename:$version does not exist in local image repository [must build image before push]\n\n"
  exit -1
fi

# segment the version string into its components
# IFS='.'
# read -a versionarr <<< "$version"
# unset IFS

# populate version information
# major="${versionarr[0]}"
# minor="${versionarr[1]}"
# patch="${versionarr[2]}"

printf "\nPushing $imagename:$version to docker hub\n\n"
docker push $imagename:$version
printf "\nPush completed\n\n"

exit 0