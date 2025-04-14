#!/bin/env bash

color_reset="\033[0m"
color_red="\033[31m"
color_green="\033[32m"
color_yellow="\033[33m"
color_blue="\033[34m"
color_cyan="\033[36m"
color_magenta="\033[35m"

usage() {
	printf "Usage: $0 [windows][run][release][clean] \n\n\twindows:  build for windows, leave empty for building for linux. If you are compiling for windows make sure to change the windows_user variable in this script\n\trun:      run the built file after building\n\trelease:  build with release flags\n\tclean:    clean the output directory\n"
	exit 1
}

log_info() {
	printf "${color_cyan}$1${color_reset}\n"
}

log_error() {
	printf "${color_red}$1${color_reset}\n"
}

log_success() {
	printf "${color_green}$1${color_reset}\n"
}

project_name="game_engine"
project_dir=$(realpath $(dirname $0))
out_dir="out"
out_name="$project_name"
linker_flags=""
tags="-tags=debug"

echo "Project Directory: $project_dir"
cd $project_dir

# If you are using windows, set the windows username bellow. As it appear under "C:/Users/XXX"
windows_user="ngkil"

windows_flag=false
release_flag=false
clean_flag=false
run_flag=false

# get parameters and set the flags

while [ "$1" != "" ]; do
	case $1 in
	windows)
		windows_flag=true
		;;
	release)
		release_flag=true
		;;
	clean)
		clean_flag=true
		;;
	run)
		run_flag=true
		;;
	*)
		usage
		;;
	esac
	shift
done

if [ "$release_flag" == "true" ]; then
	out_name="$out_name-release"
	linker_flags="$linker_flags -s -w"
	tags="-tags= "
	if [ "$windows_flag" == "true" ]; then
		linker_flags="$linker_flags -H windowsgui"
	fi
fi

# check if parameter one exist and is equal to "windows"
if [ "$windows_flag" == "true" ]; then
	# add windows to the tags
	if [ "$release_flag" == "true" ]; then
		tags="-tags=windows_os"
	else
		tags="$tags,windows_os"
	fi

	subdir="engine"
	windows_out_dir="/mnt/c/Users/$windows_user/Desktop/go_projects/$project_name/$subdir"
	if [ ! -d "$windows_out_dir" ]; then
		log_info "Creating directory $windows_out_dir"
		mkdir -p $windows_out_dir
	fi
	out_name="$out_name.exe"

	if [ "$clean_flag" == "true" ]; then
		log_info "Cleaning previous file"
		rm $out_dir/$out_name 2>/dev/null
	fi

	echo "tags: $tags"
	log_info "Packages to build:"
	GOOS=windows go list -f '{{.GoFiles}}' $tags ./src/...

	log_info "Building for windows"
	error_output=$(CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "$linker_flags" -o $out_dir/$out_name $tags ./src 2>&1)

	if [ $? -ne 0 ]; then
		log_error "Build failed with error:"
		echo "----------------------------------"
		echo "$error_output"
		exit 1
	else
		log_success "Build succeeded"
	fi

	log_info "Moving file to $windows_out_dir"
	mv $out_dir/$out_name $windows_out_dir
	if [ "$release_flag" == "false" ]; then
		printf "$out_name\npause\n" >$windows_out_dir/debug_run.bat
	fi

	# add color to the output

	printf "${color_green}Done Building${color_reset}\n---------------------------------\n"

	if [ "$run_flag" == "true" ]; then
		log_info "Running the built file"
		powershell.exe -command "C:/Users/$windows_user/Desktop/go_projects/$project_name/$subdir/$out_name"
	fi

else
	# add linux to the tags
	if [ "$release_flag" == "true" ]; then
		tags="-tags=linux_os"
	else
		tags="$tags,linux_os"
	fi

	if [ "$clean_flag" == "true" ]; then
		log_info "Cleaning previous file"
		rm $project_dir/$out_dir/$out_name #2>/dev/null
	fi

	echo "tags: $tags"

	log_info "Packages to build:"
	go list -f '{{.GoFiles}}' $tags $project_dir/src $project_dir/src/build

	log_info "Building for linux"
	error_output=$(go build -ldflags "$linker_flags" -o $project_dir/$out_dir/$out_name $tags $project_dir/src)

	if [ $? -ne 0 ]; then
		log_error "Build failed with error:"
		echo "----------------------------------"
		echo "$error_output"
		exit 1
	else
		log_success "Build succeeded"
	fi

	log_success "Done Building"
	echo "---------------------------------"

	if [ "$run_flag" == "true" ]; then
		log_info "Running the built file"
		$project_dir/$out_dir/$out_name
	fi
fi
