<script>
  import { Chart, registerables } from "chart.js";
  import { onMount } from "svelte";
  import { containers } from "../stores"; // Import your containers store

  let chartId = "chart-" + Math.random().toString(36).substring(2, 9);
  
  // 注册 Chart.js 的所有组件
  Chart.register(...registerables);
  
  let chartInstance;
  
  onMount(() => {
    // Prepare data from containers
    const containerData = $containers.map(container => {

      console.log(`Container ${container.ID} size: ${container.SizeMB}MB`); // 调试日志
      return {
        name: container.Names[0]?.replace('/', '') || container.ID.slice(0, 12),
        size: container.SizeMB
      };
    });

    const ctx = document.getElementById(chartId).getContext("2d");
    chartInstance = new Chart(ctx, {
      type: "bar", // Keep as bar chart
      data: {
        labels: containerData.map(item => item.name),
        datasets: [
          {
            label: "Size (MB)",
            data: containerData.map(item => item.size),
            backgroundColor: ["#FF6384", "#36A2EB", "#FFCE56"], // Keep original colors
            borderColor: ["#fff"], // 保持原边框颜色
            borderWidth: 1 // 保持原边框宽度
          }
        ]
      },
      options: {
        indexAxis: 'y',
        responsive: false, // Keep original responsive setting
        scales: {
          x: {
            grid: {
              display: false // 保持关闭x轴网格线
            }
          },
          y: {
            grid: {
              display: false // 保持关闭y轴网格线
            }
          }
        },
        plugins: {
          legend: {
            position: "top" // 保持图例位置
          },
          tooltip: {
            callbacks: {
              label: (context) => 
                `${context.label}: ${context.raw} MB` // 修改为显示MB单位
            }
          }
        }
      }
    });
  
    return () => chartInstance.destroy();
  });
</script>
  
<canvas id={chartId}></canvas>