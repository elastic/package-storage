#!/usr/bin/env bash
#
# Given the stack version this script will bump the version.
#
# This script is executed by the automation we are putting in place
# and it requires the git add/commit commands.
#
# Parameters:
#	$1 -> the version to be bumped. Mandatory.
#	$2 -> whether to create a branch where to commit the changes to.
#		  this is required when reusing an existing Pull Request.
#		  Optional. Default true.
#
set -euo pipefail
MSG="parameter missing."
VERSION=${1:?$MSG}
CREATE_BRANCH=${2:-true}

MAJOR_MINOR_PATCH_VERSION=$(echo "$VERSION" | sed -E -e "s#([0-9]+\.[0-9]+\.[0-9]+).*#\1#g")

OS=$(uname -s| tr '[:upper:]' '[:lower:]')

if [ "${OS}" == "darwin" ] ; then
	SED="sed -i .bck"
else
	SED="sed -i"
fi

if [ "$CREATE_BRANCH" = "true" ]; then
	git checkout -b "update-stack-version-$(date "+%Y%m%d%H%M%S")"
else
	echo "Branch creation disabled."
fi

echo "Update stack with version ${VERSION}"
${SED} -E -e "s#(image: docker\.elastic\.co/.*):[0-9]+\.[0-9]+\.[0-9]+(-[a-f0-9]{8})?#\1:${VERSION}#g" testing/environments/snapshot.yml
git add testing/environments/snapshot.yml
${SED} -E -e "s#(kibana\.version)=[0-9]+\.[0-9]+\.[0-9]+#\1=${MAJOR_MINOR_PATCH_VERSION}#g" testing/main_integration_test.go
git add testing/main_integration_test.go

echo "Commit changes"
git diff --staged --quiet || git commit -m "bump stack version ${VERSION}"
git --no-pager log -1

echo "You can now push and create a Pull Request"
