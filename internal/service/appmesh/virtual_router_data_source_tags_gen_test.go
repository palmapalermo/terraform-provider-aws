// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package appmesh_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	tfappmesh "github.com/hashicorp/terraform-provider-aws/internal/service/appmesh"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testAccAppMeshVirtualRouterDataSource_tagsSerial(t *testing.T) {
	t.Helper()

	testCases := map[string]func(t *testing.T){
		acctest.CtBasic:                  testAccAppMeshVirtualRouterDataSource_tags,
		"NullMap":                        testAccAppMeshVirtualRouterDataSource_tags_NullMap,
		"EmptyMap":                       testAccAppMeshVirtualRouterDataSource_tags_EmptyMap,
		"DefaultTags_nonOverlapping":     testAccAppMeshVirtualRouterDataSource_tags_DefaultTags_nonOverlapping,
		"IgnoreTags_Overlap_DefaultTag":  testAccAppMeshVirtualRouterDataSource_tags_IgnoreTags_Overlap_DefaultTag,
		"IgnoreTags_Overlap_ResourceTag": testAccAppMeshVirtualRouterDataSource_tags_IgnoreTags_Overlap_ResourceTag,
	}

	acctest.RunSerialTests1Level(t, testCases, 0)
}

func testAccAppMeshVirtualRouterDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_appmesh_virtual_router.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/VirtualRouter/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtKey1: config.StringVariable(acctest.CtValue1),
					}),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
		},
	})
}

func testAccAppMeshVirtualRouterDataSource_tags_NullMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_appmesh_virtual_router.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/VirtualRouter/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:        config.StringVariable(rName),
					acctest.CtResourceTags: nil,
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func testAccAppMeshVirtualRouterDataSource_tags_EmptyMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_appmesh_virtual_router.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/VirtualRouter/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:        config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{}),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func testAccAppMeshVirtualRouterDataSource_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_appmesh_virtual_router.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.AppMeshServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/VirtualRouter/data.tags_defaults/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func testAccAppMeshVirtualRouterDataSource_tags_IgnoreTags_Overlap_DefaultTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_appmesh_virtual_router.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.AppMeshServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/VirtualRouter/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtProviderKey1),
					),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
					expectFullVirtualRouterDataSourceTags(dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func testAccAppMeshVirtualRouterDataSource_tags_IgnoreTags_Overlap_ResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_appmesh_virtual_router.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.AppMeshServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/VirtualRouter/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtResourceKey1),
					),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
					expectFullVirtualRouterDataSourceTags(dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func expectFullVirtualRouterDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTagsSpecTags(tfappmesh.ServicePackage(context.Background()), resourceAddress, &types.ServicePackageResourceTags{
		IdentifierAttribute: "arn",
	}, knownValue)
}
