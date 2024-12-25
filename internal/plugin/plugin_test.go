package plugin

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	//go:embed input_examples/project_root_input_dump
	projectRootInputDump []byte
	//go:embed input_examples/grpc_root_input_dump
	grpcRootInputDump []byte
)

//go:embed input_examples/expected/package.json
var expectedPackageJson string

var (
	//go:embed input_examples/expected/project_root/index_ts
	expectedRootRelativeIndexTs string
	//go:embed input_examples/expected/project_root/tsconfig.json
	expectedRootRelativeTsConfig string
)

var (
	//go:embed input_examples/expected/grpc_root/index_ts
	expectedGrpcRelativeIndexTs string
	//go:embed input_examples/expected/grpc_root/tsconfig.json
	expectedGrpcRelativeTsConfig string
)

const (
	npmPackage      = "@godverv/velez_api/"
	indexTsName     = npmPackage + "index.ts"
	packageJsonName = npmPackage + "package.json"
	tsconfigName    = npmPackage + "tsconfig.json"
)

func TestPlugin(t *testing.T) {
	t.Run("OK_ROOT_RELATIVE", func(t *testing.T) {
		pl, err := getPlugin(projectRootInputDump)
		require.NoError(t, err)

		err = Generate(pl)
		require.NoError(t, err)

		actual := pl.Response()

		require.Equal(t, indexTsName, actual.File[0].GetName())
		require.Equal(t, expectedRootRelativeIndexTs, actual.File[0].GetContent())

		require.Equal(t, tsconfigName, actual.File[1].GetName())
		require.JSONEq(t, expectedRootRelativeTsConfig, actual.File[1].GetContent())

		require.Equal(t, packageJsonName, actual.File[2].GetName())
		require.JSONEq(t, expectedPackageJson, actual.File[2].GetContent())

	})

	t.Run("OK_GRPC_PACKAGE_RELATIVE", func(t *testing.T) {
		pl, err := getPlugin(grpcRootInputDump)
		require.NoError(t, err)

		err = Generate(pl)
		require.NoError(t, err)

		actual := pl.Response()

		require.Equal(t, indexTsName, actual.File[0].GetName())
		require.Equal(t, expectedGrpcRelativeIndexTs, actual.File[0].GetContent())

		require.Equal(t, tsconfigName, actual.File[1].GetName())
		require.JSONEq(t, expectedGrpcRelativeTsConfig, actual.File[1].GetContent())

		require.Equal(t, packageJsonName, actual.File[2].GetName())
		require.JSONEq(t, expectedPackageJson, *actual.File[2].Content)

	})
}

func getPlugin(in []byte) (*protogen.Plugin, error) {
	req := &pluginpb.CodeGeneratorRequest{}

	err := proto.Unmarshal(in, req)
	if err != nil {
		return nil, err
	}

	opts := protogen.Options{}

	return opts.New(req)
}
