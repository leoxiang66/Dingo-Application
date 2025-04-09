<script>
  import { Chart, registerables } from "chart.js";
  import { onMount } from "svelte";
  import { load_img_data } from "../stores";

  let chartId = "chart-" + Math.random().toString(36).substring(2, 9);
  
  Chart.register(...registerables);
  
  let chartInstance;
  
  onMount(async() => {
    // Generate distinct colors for each image
    const generateColors = (count) => {
      const colors = [];
      const goldenRatio = 0.618033988749895;
      let h = Math.random(); // Start at random hue
      
      for (let i = 0; i < count; i++) {
        h += goldenRatio;
        h %= 1;
        colors.push(`hsl(${Math.floor(h * 360)}, 70%, 60%)`);
      }
      return colors;
    };

    let images = await load_img_data()
    const imageData = images.map(image => {
      return {
        name: image.ID.slice(7, 19),
        size: image.Size / (1024 * 1024) // Convert to MB
      };
    });

    const imageColors = generateColors(imageData.length);

    const ctx = document.getElementById(chartId).getContext("2d");
    chartInstance = new Chart(ctx, {
      type: "bar",
      data: {
        labels: imageData.map(item => item.name),
        datasets: [
          {
            label: "Size (MB)",
            data: imageData.map(item => item.size),
            backgroundColor: imageColors, // Different color for each image
            borderColor: "#fff",
            borderWidth: 1
          }
        ]
      },
      options: {
        indexAxis: 'y',
        responsive: false,
        scales: {
          x: {
            grid: {
              display: false
            },
            title: {
              display: true,
              text: 'Size (MB)'
            }
          },
          y: {
            grid: {
              display: false
            }
          }
        },
        plugins: {
          legend: {
            position: "top"
          },
          tooltip: {
            callbacks: {
              label: (context) => 
                `${context.raw.toFixed(2)} MB`
            }
          }
        }
      }
    });
  
    return () => chartInstance.destroy();
  });
</script>


<canvas id={chartId}></canvas>
