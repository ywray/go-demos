# go-demos
Improving the development efficiency of Go code

This project contains quick start of using the standard library, third-party open source libraries, and some personally written utility code snippets (such as calculating the edit distance between two strings, read files and so on).

Some code needs to be run to see the effect, and the specific test methods will be written in the main.go file in each subpath.

## Code Structure

### personal : Some utility code snippets written by myself
* read_file: Traverse all files under a certain path
* ldflags: Cooperate with the makefile file to assign variables in the makefile to the Go internals

### stl : Contains demos related to the standard library
* sort_map: Sort map by value
* errgroup: Implement sub-coroutine error propagation based on WaitGroup, and use context to control the coroutine's lifecycle at the same time
* format: Compare Sprintf and math.Round
* use_regexp: Whether the string meets the regular expression, if not, customize the error message
* use_url: Parse addresses/parameters in URLs
* md5
* use_string/base64: Encoding and decoding
* use_string/index: The difference between string and []rune indexing

### third :  Contains demos related to third-party libraries
* use_rollingwriter: Write log files locally, configurable for single log file name, file size, rolling strategy, split rule, maximum retention, etc.
* use_redis/redislock: go-redis lib
* use_zerolog: Logging library
* use_uuid: Generate pure numeric IDs by intercepting UUIDs
* use_yaml: Read and parse YAML configuration files
* use_gin: Gin implements a simple ping pong server
* fuzzysearch: String fuzzy matching
