package dockerbackend

// import (
// 	"context"
// 	"fmt"
// 	"text/tabwriter"
// 	"os"
// 	"time"

// 	"github.com/docker/docker/api/types/container"
// 	"github.com/docker/docker/client"
// )

// func Get_all_containers() {
// 	apiClient, err := client.NewClientWithOpts(
// 		client.FromEnv,
// 		client.WithAPIVersionNegotiation(),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer apiClient.Close()

// 	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{
// 		All: true,
// 		Size: true, // 获取容器大小信息
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 创建表格输出
// 	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
// 	fmt.Fprintln(w, "CONTAINER ID\tIMAGE\tCOMMAND\tCREATED\tSTATUS\tSIZE\tPORTS\tNAMES")

// 	for _, ctr := range containers {
// 		// 格式化容器ID (显示前12个字符)
// 		containerID := ctr.ID
// 		if len(containerID) > 12 {
// 			containerID = containerID[:12]
// 		}

// 		// 格式化创建时间
// 		created := time.Unix(ctr.Created, 0).Format("2006-01-02 15:04:05")

// 		// 格式化端口信息
// 		var ports string
// 		for _, p := range ctr.Ports {
// 			if p.PublicPort > 0 {
// 				ports += fmt.Sprintf("%d->%d/%s, ", p.PublicPort, p.PrivatePort, p.Type)
// 			} else {
// 				ports += fmt.Sprintf("%d/%s, ", p.PrivatePort, p.Type)
// 			}
// 		}
// 		if len(ports) > 2 {
// 			ports = ports[:len(ports)-2] // 移除最后的", "
// 		}

// 		// 格式化名称 (移除前面的斜杠)
// 		var names string
// 		for _, n := range ctr.Names {
// 			if len(n) > 1 {
// 				names += n[1:] + ", "
// 			}
// 		}
// 		if len(names) > 2 {
// 			names = names[:len(names)-2] // 移除最后的", "
// 		}

// 		// 格式化容器大小 (MB)
// 		sizeMB := float64(ctr.SizeRw) / (1024 * 1024)

// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%.2fMB\t%s\t%s\n",
// 			containerID,
// 			ctr.Image,
// 			ctr.Command,
// 			created,
// 			ctr.Status,
// 			sizeMB,
// 			ports,
// 			names)
// 	}

// 	w.Flush() // 确保所有输出被写入
// 	fmt.Printf("\nTotal containers: %d\n", len(containers))
// }