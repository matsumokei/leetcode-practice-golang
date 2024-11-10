#!/bin/bash

url=$1
num=$2

if [ -z "$url" ]; then
    echo "Error: No URL provided."
    exit 1
fi

if ! [[ $url =~ ^https://leetcode.com/problems/.+ ]]; then
    echo "Error: Invalid URL format."
    exit 1
fi

DIR="leetcode/"

problem_name=$(echo $url | sed -n 's|.*/problems/\(.*\)/.*|\1|p')

problem_name=${problem_name%/description}

snake_case_filename=$(echo $problem_name | tr '-' '_')
echo "${snake_case_filename}"

mkdir -p $DIR$num.$snake_case_filename
touch $DIR$num.$snake_case_filename/$snake_case_filename.go
touch $DIR$num.$snake_case_filename/README.md
touch $DIR$num.$snake_case_filename/$snake_case_filename\_test.go

cat << EOS > $DIR$num.$snake_case_filename/$snake_case_filename.go
package leetcode

// ${url}

/*
TODO: paste description from leetcode
*/

func dummy() {
    // TODO: implement
}

EOS

cat << EOS > $DIR$num.$snake_case_filename/$snake_case_filename\_test.go
package leetcode

import (
	"fmt"
	"testing"
)
EOS

cat << EOS > $DIR$num.$snake_case_filename/README.md
<script type="text/javascript" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/3.2.2/es5/tex-mml-chtml.min.js">
</script>
<script type="text/x-mathjax-config">
 MathJax.Hub.Config({
 tex2jax: {
 inlineMath: [['$', '$'] ],
 displayMath: [ ['$$','$$'], ["\\[","\\]"] ]
 }
 });
</script>


## Problems

## Explanation

## Requirements

EOS