package postman

// func TestAccWorkspaceDataSource(t *testing.T) {
//     resource.Test(t, resource.TestCase{
//         ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
//         Steps: []resource.TestStep{
//             // Read testing
//             {
//                 Config: providerConfig + `data "postman_workspace" "test" {}`,
//                 Check: resource.ComposeAggregateTestCheckFunc(
//                     // Verify number of workspace returned
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.#", "9"),
//                     // Verify the first coffee to ensure all attributes are set
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.description", ""),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.id", "1"),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.image", "/hashicorp.png"),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.ingredients.#", "1"),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.ingredients.0.id", "6"),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.name", "HCP Aeropress"),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.price", "200"),
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "workspace.0.teaser", "Automation in a cup"),
//                     // Verify placeholder id attribute
//                     resource.TestCheckResourceAttr("data.postman_workspace.test", "id", "placeholder"),
//                 ),
//             },
//         },
//     })
// }
