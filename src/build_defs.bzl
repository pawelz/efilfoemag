load("@io_bazel_rules_go//go:def.bzl", "go_test")
load("@rules_python//python:defs.bzl", "py_binary")

def go_generated_test(
    name,
    src,
    deps = [],
    embed = [],
):
  """Go tests generated with Python scripts.

  Args:
    name the name of the test
    src a Python scripts that spits out a Go test
    deps passed to the generated go_test rule
    embed passed to the generated go_test rule
  """
  generator_name = src[:-len(".py")]
  generated_test_file = name + "_gen_test.go"
  py_binary(
      name = generator_name,
      srcs = [src],
  )
  native.genrule(
      name = name + "_gen_test",
      srcs = [generator_name],
      cmd = "./$(locations " + generator_name + ") > $@",
      outs = [generated_test_file],
  )
  go_test(
      name = name,
      srcs = [generated_test_file],
      deps = deps,
      embed = embed,
  )
