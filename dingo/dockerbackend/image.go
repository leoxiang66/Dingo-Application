package dockerbackend

import (
	"context"
	"fmt"
	"text/tabwriter"
	"os"
	"time"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// Image 类型封装单个镜像的所有信息
type Image struct {
	ID          string            // 镜像ID (完整)
	RepoTags    []string          // 仓库标签
	RepoDigests []string          // 仓库摘要
	Created     time.Time         // 创建时间
	Size        int64             // 镜像大小(bytes)
	SharedSize  int64             // 共享大小
	VirtualSize int64             // 虚拟大小
	Labels      map[string]string // 标签
}

// Images 是Image的切片类型
type Images []Image



func Get_all_images() Images{
	cli, err := client.NewClientWithOpts(client.FromEnv,client.WithAPIVersionNegotiation(),)
	if err != nil {
		panic(err)
	}

	dockerImages, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		panic(err)
	}

	// 转换为我们的Images类型
	var images Images
	for _, di := range dockerImages {
		images = append(images, Image{
			ID:          di.ID,
			RepoTags:    di.RepoTags,
			RepoDigests: di.RepoDigests,
			Created:     time.Unix(di.Created, 0),
			Size:        di.Size,
			SharedSize:  di.SharedSize,
			VirtualSize: di.VirtualSize,
			Labels:      di.Labels,
		})
	}

	return images
}

// PrintTable 以表格形式打印镜像列表
func (imgs Images) PrintTable() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "IMAGE ID\tREPOSITORY:TAG\tSIZE\tShared Size\tCREATED\tVIRTUAL SIZE")

	for _, img := range imgs {
		// 格式化镜像ID (显示前12个字符)
		imageID := img.ID
		if len(imageID) > 12 {
			imageID = imageID[7:19] // 去除"sha256:"前缀并取12位
		}

		// 格式化仓库标签
		var repoTag string
		if len(img.RepoTags) > 0 {
			repoTag = img.RepoTags[0]
		} else {
			repoTag = "<none>:<none>"
		}

		// 计算大小(MB)
		sizeMB := float64(img.Size) / (1024 * 1024)
		sharedSizeMB := float64(img.SharedSize) / (1024 * 1024)
		virtualSizeMB := float64(img.VirtualSize) / (1024 * 1024)

		fmt.Fprintf(w, "%s\t%s\t%.2fMB\t%.2fMB\t%s\t%.2fMB\n",
			imageID,
			repoTag,
			sizeMB,
			sharedSizeMB,
			img.Created.Format("2006-01-02 15:04:05"),
			virtualSizeMB)
	}

	w.Flush()
	fmt.Printf("\nTotal images: %d\n", len(imgs))
}
