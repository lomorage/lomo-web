#!/bin/bash

# MIT License
#
# Copyright (c) 2017 Mauro Bringolf
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

CAPITALIZE() {
  first_capital="$(echo "$1"|cut -c1|tr '[:lower:]' '[:upper:]' )"
  rest_regular="$(echo "$1"|cut -c2-)"
  echo "$first_capital$rest_regular"
}

CREATE_CHANGELOG() {

  # Definition of default commit types
  COMMIT_TYPES=( breaking bugfix feature frontend backend workflow testing documentation internal )

  # String to accumulate changelog
  CONTENT=""

  # Get all commits with type annotations and make them paragraphs.
  for TYPE in "${COMMIT_TYPES[@]}"
    do
      if [ -z "$1" ]
        then
          PARAGRAPH="$(git log --format="* %s (%h)" --grep="^\[${TYPE}\]")"
        else
          PARAGRAPH="$(git log "$1"..HEAD --format="* %s (%h)" --grep="^\[${TYPE}\]")"
      fi
      if [ ! -z "$PARAGRAPH" ]
        then
          TITLE="$(CAPITALIZE "$TYPE")"
          PARAGRAPH="${PARAGRAPH//\[$TYPE\] /}"
          CONTENT="$CONTENT## $TITLE\n\n$PARAGRAPH\n\n"
      fi
    done

  # Regex used to find commits without types
  TYPES_REGEX=""
  for TYPE in "${COMMIT_TYPES[@]}"
    do
      TYPES_REGEX="$TYPES_REGEX\[$TYPE\]\|"
  done
  TYPES_REGEX="$TYPES_REGEX\[skip-changelog\]"

  # Get all commit without type annotation and make them another paragraph.
  if [ -z "$1" ]
    then
      PARAGRAPH="$(git log --format=";* %s (%h);")"
    else
      PARAGRAPH="$(git log "$1"..HEAD --format=";* %s (%h);")"
  fi
  OIFS="$IFS"
  IFS=";"
  FILTERED_PARAGRAPH=""
  for COMMIT in $PARAGRAPH
   do
     TRIMMED_COMMIT="$(echo "$COMMIT" | xargs)"
    if [ ! -z "$TRIMMED_COMMIT" ] && ! echo "$TRIMMED_COMMIT" | grep -q "$TYPES_REGEX"
      then
        FILTERED_PARAGRAPH="$FILTERED_PARAGRAPH$TRIMMED_COMMIT\n"
    fi
  done
  IFS="$OIFS"

  # Only add to content if there are commits without type annotations.
  if [ ! -z "$FILTERED_PARAGRAPH" ]
   then
     CONTENT="\n\n## CHANGE LOGS\n\n$FILTERED_PARAGRAPH\n\n$CONTENT"
  fi

  # Output changelog
  echo -e "$CONTENT"
}

# Generate changelog either from last tag or from beginning
if [ -z "$(git tag)" ]
  then
    CONTENT=$(CREATE_CHANGELOG)
  else
    LATEST_RELEASE="$(git describe --tags --abbrev=0)"
    CONTENT=$(CREATE_CHANGELOG "${LATEST_RELEASE}")
fi


if [ -z "$CONTENT" ]
  then
    echo -e "No changes made since last release $LATEST_RELEASE."
	exit
fi

RELEASE_VERSION=$(awk -F "\"" '/[[:digit:]]{4}-[[:digit:]]{2}-[[:digit:]]{2}\.[[:digit:]]{2}-[[:digit:]]{2}-[[:digit:]]{2}\.0\.[a-zA-Z0-9]{7}/{ print $2 }' main.go)
CONTENT="$RELEASE_VERSION$CONTENT"
echo "$CONTENT"

hub release create -a LomoWebOSX.zip -m "$CONTENT" $RELEASE_VERSION

