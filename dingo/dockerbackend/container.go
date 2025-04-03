package dockerbackend

import (
	"context"
	"fmt"
	"text/tabwriter"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Container 类型封装单个容器的所有信息
type Container struct {
	ID      string    // 容器ID (完整)
	Image   string    // 镜像名称
	Command string    // 启动命令
	Created time.Time // 创建时间
	Status  string    // 运行状态
	SizeMB  float64   // 容器大小(MB)
	Ports   []Port    // 端口映射
	Names   []string  // 容器名称(不带斜杠)
}

// Port 表示端口映射信息
type Port struct {
	IP          string
	PublicPort  uint16
	PrivatePort uint16
	Type        string
}

// Containers 是Container的切片类型
type Containers []Container

func Get_all_containers() Containers {
	apiClient, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	// 获取原始容器列表
	dockerContainers, err := apiClient.ContainerList(context.Background(), container.ListOptions{
		All:  true,
		Size: true,
	})
	if err != nil {
		panic(err)
	}

	// 转换为我们的Containers类型
	var containers Containers
	for _, dc := range dockerContainers {
		// 转换端口信息
		var ports []Port
		for _, p := range dc.Ports {
			ports = append(ports, Port{
				IP:          p.IP,
				PublicPort:  p.PublicPort,
				PrivatePort: p.PrivatePort,
				Type:        p.Type,
			})
		}

		// 处理容器名称(移除斜杠)
		var names []string
		for _, n := range dc.Names {
			if len(n) > 0 && n[0] == '/' {
				names = append(names, n[1:])
			} else {
				names = append(names, n)
			}
		}

		containers = append(containers, Container{
			ID:      dc.ID,
			Image:   dc.Image,
			Command: dc.Command,
			Created: time.Unix(dc.Created, 0),
			Status:  dc.Status,
			SizeMB:  float64(dc.SizeRw) / (1024 * 1024),
			Ports:   ports,
			Names:   names,
		})
	}

	return containers
}

// PrintTable 以表格形式打印容器列表
func (cs Containers) PrintTable() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "CONTAINER ID\tIMAGE\tCOMMAND\tCREATED\tSTATUS\tSIZE\tPORTS\tNAMES")

	for _, c := range cs {
		// 格式化端口
		var portsStr string
		for _, p := range c.Ports {
			if p.PublicPort > 0 {
				portsStr += fmt.Sprintf("%d->%d/%s, ", p.PublicPort, p.PrivatePort, p.Type)
			} else {
				portsStr += fmt.Sprintf("%d/%s, ", p.PrivatePort, p.Type)
			}
		}
		if len(portsStr) > 2 {
			portsStr = portsStr[:len(portsStr)-2]
		}

		// 格式化名称
		var namesStr string
		for _, n := range c.Names {
			namesStr += n + ", "
		}
		if len(namesStr) > 2 {
			namesStr = namesStr[:len(namesStr)-2]
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%.2fMB\t%s\t%s\n",
			c.ID[:12],
			c.Image,
			c.Command,
			c.Created.Format("2006-01-02 15:04:05"),
			c.Status,
			c.SizeMB,
			portsStr,
			namesStr)
	}

	w.Flush()
	fmt.Printf("\nTotal containers: %d\n", len(cs))
}

// 使用示例
func ExampleUsage() {
	containers := Get_all_containers()
	containers.PrintTable()
}