<script>
    import { Chart, registerables } from "chart.js";
    import { onMount } from "svelte";

    let chartId = "chart-" + Math.random().toString(36).substring(2, 9); // 生成随机 ID
  
    // 注册 Chart.js 的所有组件
    Chart.register(...registerables);
  
    let chartInstance;
  
    onMount(() => {
      const ctx = document.getElementById(chartId).getContext("2d");
      chartInstance = new Chart(ctx, {
        type: "bar", // 修改为饼图类型
        data: {
          labels: ["苹果", "香蕉", "橘子"],
          datasets: [
            {
              label: "销量占比",
              data: [12, 19, 3],
              backgroundColor: ["#FF6384", "#36A2EB", "#FFCE56"],
              borderColor: ["#fff"], // 可选：设置边框颜色
              borderWidth: 1 // 可选：边框宽度
            }
          ]
        },
        options: {
          indexAxis: 'y',
          responsive: false,
          scales: {
            x: {
              grid: {
                display: false // 关闭x轴网格线
              }
            },
            y: {
              grid: {
                display: false // 关闭y轴网格线
              }
            }
          },
          plugins: {
            legend: {
              position: "top" // 图例位置（可选）
            },
            tooltip: {
              callbacks: {
                label: (context) => 
                  `${context.label}: ${context.raw}` // 显示百分比
              }
            }
          }
        }
      });
  
      return () => chartInstance.destroy();
    });
  </script>
  
  <canvas id={chartId}></canvas>