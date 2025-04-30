# this project

Testing sonic using 1.24 sdk in bazel will cause errors

# my local environment
bazelisk
```
> bazelisk version
Bazelisk version: development
Build label: 8.2.1
Build target: @@//src/main/java/com/google/devtools/build/lib/bazel:BazelServer
Build time: Thu Apr 17 18:31:45 2025 (1744914705)
Build timestamp: 1744914705
Build timestamp as int: 1744914705
```

golang 
```bash
> go version
go version go1.24.0 darwin/arm64
```

go env (desensitized)
```
 go env
AR='ar'
CC='cc'
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_ENABLED='1'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
CXX='c++'
GCCGO='gccgo'
GO111MODULE='on'
GOARCH='arm64'
GOARM64='v8.0'
GOAUTH='netrc'
GOBIN=''
GOCACHE='/Users/jettcc/Library/Caches/go-build'
GOCACHEPROG=''
GODEBUG=''
GOENV='/Users/jettcc/Library/Application Support/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFIPS140='off'
GOFLAGS=''
GOGCCFLAGS='-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -ffile-prefix-map=/var/folders/d3/xm_9n5tn33d1mhmvwr5stw7c0000gn/T/go-build4206925790=/tmp/go-build -gno-record-gcc-switches -fno-common'
GOHOSTARCH='arm64'
GOHOSTOS='darwin'
GOINSECURE=''
GOMODCACHE='/Users/jettcc/go/pkg/mod'
GOOS='darwin'
GOPATH='/Users/jettcc/go'
GOPROXY='https://goproxy.cn,direct'
GOROOT='/opt/homebrew/Cellar/go/1.24.0/libexec'
GOSUMDB='sum.golang.org'
GOTELEMETRY='on'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/opt/homebrew/Cellar/go/1.24.0/libexec/pkg/tool/darwin_arm64'
GOVCS=''
GOVERSION='go1.24.0'
GOWORK=''
PKG_CONFIG='pkg-config'
```


# error recurrence method
the root directory
```bash
bazlisk build //...
```

u will see
```
> bazelisk build //...
Starting local Bazel server (8.2.1) and connecting to it...
INFO: Analyzed 3 targets (185 packages loaded, 16082 targets configured).
ERROR: /private/var/tmp/_bazel_jettcc/f959f436ee736a058c22899accf17a98/external/gazelle++go_deps+com_github_bytedance_sonic/internal/rt/BUILD.bazel:3:11: GoCompilePkg external/gazelle++go_deps+com_github_bytedance_sonic/internal/rt/rt.a failed: (Exit 1): builder failed: error executing GoCompilePkg command (from target @@gazelle++go_deps+com_github_bytedance_sonic//internal/rt:rt) bazel-out/darwin_arm64-opt-exec-ST-d57f47055a04/bin/external/rules_go++go_sdk+sonic-bazel__download_0/builder_reset/builder compilepkg -sdk external/rules_go++go_sdk+sonic-bazel__download_0 -goroot ... (remaining 63 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
external/gazelle++go_deps+com_github_bytedance_sonic/internal/rt/stubs.go:33:22: undefined: GoMapIterator
external/gazelle++go_deps+com_github_bytedance_sonic/internal/rt/stubs.go:36:54: undefined: GoMapIterator
compilepkg: error running subcommand external/rules_go++go_sdk+sonic-bazel__download_0/pkg/tool/darwin_arm64/compile: exit status 2
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 4.522s, Critical Path: 0.47s
INFO: 10 processes: 47 action cache hit, 8 internal, 2 darwin-sandbox.
ERROR: Build did NOT complete successfully
```