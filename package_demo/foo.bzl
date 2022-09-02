load("@io_bazel_rules_go//go:def.bzl", "go_binary")
print("this is from remote!")

def _foo_binary_impl(ctx):
  out = ctx.actions.declare_file(ctx.label.name)
  ctx.actions.write(
    output = out,
    content = "Hello {}!\n".format(ctx.attr.username),
  )
  return [DefaultInfo(files = depset([out]))]

foo_binary = rule(
  implementation=_foo_binary_impl,
  attrs = {
    'username': attr.string(),
  },
)

def my_macro(name):
    go_binary(
        name = name,
        srcs = ["main.go"],
        deps = ["//package_demo/a_package:mymath"],
    )   
    
