<script>
  import { Chart, registerables } from "chart.js";
  import { onMount } from "svelte";
  import { images } from "../stores";

  let chartId = "chart-" + Math.random().toString(36).substring(2, 9);

  Chart.register(...registerables);

  let chartInstance;

  onMount(() => {
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

    const imageData = $images.map((image) => {
      return {
        name: image.RepoTags[0] || image.ID.slice(7, 19),
        size: image.Size / (1024 * 1024),
        virtualSize: image.VirtualSize / (1024 * 1024),
      };
    });

    const imageColors = generateColors(imageData.length);
    const virtualColors = imageColors.map((color) =>
      color.replace("60%)", "40%)")
    ); // Darker version

    const ctx = document.getElementById(chartId).getContext("2d");
    chartInstance = new Chart(ctx, {
      type: "bar",
      data: {
        labels: imageData.map((item) => item.name),
        datasets: [
          {
            label: "Actual Size (MB)",
            data: imageData.map((item) => item.size),
            backgroundColor: imageColors, // Different color for each image
            borderColor: "#fff",
            borderWidth: 1,
          },
          {
            label: "Virtual Size (MB)",
            data: imageData.map((item) => item.virtualSize),
            backgroundColor: virtualColors, // Matching but darker colors
            borderColor: "#fff",
            borderWidth: 1,
          },
        ],
      },
      options: {
        indexAxis: "y",
        responsive: false,
        scales: {
          x: {
            grid: {
              display: false,
            },
            title: {
              display: true,
              text: "Size (MB)",
            },
          },
          y: {
            grid: {
              display: false,
            },
          },
        },
        plugins: {
          legend: {
            position: "top",
            onClick: (e, legendItem, legend) => {
              // Toggle visibility when clicking legend items
              const index = legendItem.datasetIndex;
              const meta = chartInstance.getDatasetMeta(index);
              meta.hidden = !meta.hidden;
              chartInstance.update();
            },
          },
          tooltip: {
            callbacks: {
              label: (context) => {
                const datasetLabel = context.dataset.label || "";
                return `${datasetLabel}: ${context.raw.toFixed(2)} MB`;
              },
              afterLabel: (context) => {
                // Show both sizes in tooltip
                if (context.datasetIndex === 0) {
                  const virtualSize = imageData[context.dataIndex].virtualSize;
                  return `Virtual Size: ${virtualSize.toFixed(2)} MB`;
                }
              },
            },
          },
        },
      },
    });

    return () => chartInstance.destroy();
  });
</script>

<canvas id={chartId}></canvas>
