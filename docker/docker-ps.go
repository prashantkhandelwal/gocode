package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ImageInfo struct {
	Labels   map[string]string
	RepoTags []string
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// containers, err := cli.ContainerList(context.Background(),
	// 	types.ContainerListOptions{})

	// if err != nil {
	// 	panic(err)
	// }

	// for _, container := range containers {
	// 	//fmt.Printf("%s %s\n", container.ID[:10], container.ImageID)
	// 	//fmt.Println(container.Image)
	// 	// image, _, err := cli.ImageInspectWithRaw(context.Background(), container.ImageID)
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }

	// 	// fmt.Println(image.Container)
	// 	// fmt.Printf("%s", image.Parent)

	// }

	imageSummary, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	//fmt.Println(imageSummary.RepoTags)

	for _, imageinfo := range imageSummary {

		fmt.Printf("Image Id: %s", imageinfo.ID+"\n")
		i := &ImageInfo{
			RepoTags: imageinfo.RepoTags,
		}

		for _, t := range i.RepoTags {
			fmt.Println(t)
			//fmt.Printf("Repository Name: %s", t.RepoTags[2])
		}

		// i := &ImageInfo{imageinfo.Labels}

		// for k, v := range i.Labels {
		// 	fmt.Printf("Image Label %s with value %s\n", k, v)
		// }
	}

}
