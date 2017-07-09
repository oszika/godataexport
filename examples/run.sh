#!/bin/sh

generate(){
	mode=$1

	`MODE=$mode go generate`
	output=$(go run *.go)

	if [ "$output" = "My exported template" ]; then
		echo "Mode $mode success"
	else
		echo "Mode $mode fail"
	fi

	rm data.go
}

generate source
generate path
