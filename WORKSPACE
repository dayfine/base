load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ================================================================
# Repositories
# ================================================================
http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.0/rules_go-0.19.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.0/rules_go-0.19.0.tar.gz",
    ],
    sha256 = "9fb16af4d4836c8222142e54c9efa0bb5fc562ffc893ce2abeac3e25daead144",
)
http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
    ],
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
)

# ================================================================
# Go extensions
# ================================================================
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()


go_repository(
    name = "com_github_google_go-cmp",
    commit = "3af367b6b30c263d47e8895973edcca9a49cf029",
    importpath = "github.com/google/go-cmp",
)

go_repository(
    name = "com_github_d4l3k_messagediff",
    commit = "b9e99b2f9263a86c71c1ca4507f34502448c58a4",
    importpath = "github.com/d4l3k/messagediff",
)
