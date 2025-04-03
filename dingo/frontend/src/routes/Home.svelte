<script>
  import ContainerChart from "../components/ContainerChart.svelte";
  import ImageChart from "../components/ImageChart.svelte";
  import Sidebar from "../components/Sidebar.svelte";
</script>

<div class="flex w-full h-full overflow-y-auto">
  <Sidebar />

  <div class="block mx-auto mt-[10%]">
    <div class="flex flex-col">
      <div class="flex items-center mb-10 ml-auto">
        <div class="inline-grid *:[grid-area:1/1]">
          <div class="status status-success animate-ping"></div>
          <div class="status status-success"></div>
        </div>

        <p class="ml-3 mr-20 text-black font-mono text-sm">Docker backend is up</p>
      </div>
      <div class="flex justify-center items-start gap-10">
        <div class="mx-5 card shadow-lg flex-1 min-h-[500px] w-[500px]">
          <div class="card-body h-full flex justify-center items-center"><ImageChart /></div>
        </div>

        <div class="mx-5 card shadow-lg flex-1 min-h-[500px] w-[500px]">
          <div class="card-body h-full flex justify-center items-center"><ContainerChart /></div>
        </div>
      </div>

      
    </div>
  </div>
</div>
