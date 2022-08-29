load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "hello",
    srcs = ["main.go"],
    deps = ["//package_demo/a_package:mymath"],
)