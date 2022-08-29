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