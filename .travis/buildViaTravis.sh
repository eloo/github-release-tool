#!/bin/bash
# This script will build the project.

if [ "$TRAVIS_PULL_REQUEST" != "false" ]; then
  echo -e "Build Pull Request #$TRAVIS_PULL_REQUEST => Branch [$TRAVIS_BRANCH]"
  make goBuild
elif [ "$TRAVIS_BRANCH" == "master" ] ; then
  echo -e 'Build Master Branch ['$TRAVIS_BRANCH']'
  make goBuild
elif [ "$TRAVIS_BRANCH" == "develop" ] ; then
  echo -e 'Build Develop Branch ['$TRAVIS_BRANCH']'
  make goBuild
else
  echo -e 'Build brand ['$TRAVIS_BRANCH']'
  make goBuild
fi

exit $EXIT
