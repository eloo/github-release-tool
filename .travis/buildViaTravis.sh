#!/bin/bash
# This script will build the project.
echo -e "Fetching depencies and test"
if [ "$TRAVIS_PULL_REQUEST" != "false" ]; then
  echo -e "Build Pull Request #$TRAVIS_PULL_REQUEST => Branch [$TRAVIS_BRANCH]"
  make test build
elif [ "$TRAVIS_BRANCH" == "master" ] ; then
  echo -e 'Build Master Branch ['$TRAVIS_BRANCH']'
  make test golang-coverage-report release
elif [ "$TRAVIS_BRANCH" == "develop" ] ; then
  echo -e 'Build Develop Branch ['$TRAVIS_BRANCH']'
  make test golang-coverage-report build
else
  echo -e 'Build brand ['$TRAVIS_BRANCH']'
  make test build
fi

exit $EXIT
