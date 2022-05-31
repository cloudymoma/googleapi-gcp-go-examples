package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Kubernetes Engine API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/container
// 2. This sample uses Application Default Credentials for authentication.
//    If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run
//    `gcloud beta auth application-default login`.
//    For more information, see
//    https://developers.google.com/identity/protocols/application-default-credentials
// 3. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
        "fmt"
        "log"

        "golang.org/x/net/context"
        "golang.org/x/oauth2/google"
        container "google.golang.org/api/container/v1"
)

func main() {
        ctx := context.Background()

        c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
        if err != nil {
                log.Fatal(err)
        }

        containerService, err := container.New(c)
        if err != nil {
                log.Fatal(err)
        }

        // Deprecated. The Google Developers Console [project ID or project
        // number](https://developers.google.com/console/help/new/#projectnumber).
        // This field has been deprecated and replaced by the parent field.
        projectId := "flius-vpc-2" // TODO: Update placeholder value.

        // Deprecated. The name of the Google Compute Engine
        // [zone](/compute/docs/zones#available) in which the cluster
        // resides.
        // This field has been deprecated and replaced by the parent field.
        zone := "us-central1-c" // TODO: Update placeholder value.

        // Deprecated. The name of the cluster.
        // This field has been deprecated and replaced by the parent field.
        clusterId := "gke-sa-test" // TODO: Update placeholder value.

        rb := &container.CreateNodePoolRequest{
                // TODO: Add desired fields of the request body.
                NodePool: &container.NodePool{
                  Name: "pool-3",
                  InitialNodeCount: 1,
                  Config: &container.NodeConfig{
                    ServiceAccount: "sa-image-1@flius-vpc-2.iam.gserviceaccount.com",
                    OauthScopes: []string{
              					"https://www.googleapis.com/auth/cloud-platform",
              				},
                  },
                },
        }

        resp, err := containerService.Projects.Zones.Clusters.NodePools.Create(projectId, zone, clusterId, rb).Context(ctx).Do()
        if err != nil {
                log.Fatal(err)
        }

        // TODO: Change code below to process the `resp` object:
        fmt.Printf("%#v\n", resp)
}
