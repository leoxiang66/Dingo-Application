<script>
  import Demo from "../components/Demo.svelte";
  import Sidebar from "../components/Sidebar.svelte";
</script>

<div class="flex w-full h-full">
  <Sidebar />

  <div class="block m-auto">
    <div class="flex flex-col">
      <div class="flex justify-center items-center">
        <div class="mx-10 card shadow-lg">
          <div class="card-body"><Demo /></div>
        </div>

        <div class="mx-10 card shadow-lg">
          <div class="card-body"><Demo /></div>
        </div>
      </div>

      <div class="flex items-center mt-10">
        <div class="inline-grid *:[grid-area:1/1]">
          <div class="status status-error animate-ping"></div>
          <div class="status status-error"></div>
        </div>

        <p class="mx-3 text-black font-mono text-sm">Docker backend is down</p>
      </div>
    </div>
  </div>
</div>
