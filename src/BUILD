load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

load("//src:build_defs.bzl", "go_generated_test")

go_library(
    name = "efilfoemag_lib",
    srcs = ["efilfoemag.go"],
    importpath = "github.com/pawelz/efilfoemag/src",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "efilfoemag",
    embed = [":efilfoemag_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "state",
    srcs = ["state.go"],
    importpath = "github.com/pawelz/efilfoemag/src/state",
    visibility = ["//visibility:public"],
)

go_library(
    name = "bits",
    srcs = ["bits.go"],
    importpath = "github.com/pawelz/efilfoemag/src/bits",
    visibility = ["//visibility:public"],
)

go_test(
    name = "bits_test",
    srcs = ["bits_test.go"],
    embed = [":bits"],
)

go_library(
    name = "neighborhood",
    srcs = ["neighborhood.go"],
    deps = [
        ":bits",
        ":state",
    ],
    importpath = "github.com/pawelz/efilfoemag/src/neighborhood",
    visibility = ["//visibility:public"],
)

go_test(
    name = "neighborhood_test",
    srcs = ["neighborhood_test.go"],
    embed = [":neighborhood"],
    deps = [
        ":state",
    ],
)

go_library(
    name = "grid",
    srcs = ["grid.go"],
    deps = [
        ":bits",
        ":state",
    ],
    importpath = "github.com/pawelz/efilfoemag/src/grid",
    visibility = ["//visibility:public"],
)

go_test(
    name = "grid_test",
    srcs = ["grid_test.go"],
    deps = [
        ":bits",
    ],
    embed = [":grid"],
)

go_generated_test(
    name = "grid_gentest",
    src = "grid_testgen.py",
    deps = [
        ":state",
    ],
    embed = [":grid"],
)
