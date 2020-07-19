# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
