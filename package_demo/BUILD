load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load(":foo.bzl", "foo_binary")

go_binary(
    name = "hello",
    srcs = ["main.go"],
    deps = [
         "//package_demo/a_package:mymath",
    ],
)

foo_binary(
    name = "bin",
    username = "Alice",
)