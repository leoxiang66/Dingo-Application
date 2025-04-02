<script>
  import { page_idx } from "../stores";
  import { Layers, Cuboid, ChartLine, GithubIcon } from "@lucide/svelte";

</script>

<div class="h-full flex">
  <div class="ml-2 flex flex-col items-center h-full w-[70px]">
    <div class="p-6 pt-10">
      <div class="pl-1 w-16">
        <img src="/box.png" alt="avatar" class="w-14" />
      </div>
    </div>

    <div
      class="w-full m-1 flex justify-center tooltip tooltip-right"
      data-tip="Status"
    >
      <button
        class="sidebar-item hover:bg-green-200 {$page_idx === 0
          ? 'bg-green-300 hover:bg-green-300'
          : ''}"
        on:click={() => page_idx.set(0)}
      >
        <ChartLine />
      </button>
    </div>

    <div
      class="w-full m-1 flex justify-center tooltip tooltip-right"
      data-tip="Images"
    >
      <button
        class="sidebar-item hover:bg-green-200 {$page_idx === 1
          ? 'bg-green-300 hover:bg-green-300'
          : ''}"
        on:click={() => page_idx.set(1)}
      >
        <Layers />
      </button>
    </div>

    <div
      class="w-full m-1 flex justify-center tooltip tooltip-right"
      data-tip="Containers"
    >
      <button
        class="sidebar-item hover:bg-green-200 {$page_idx === 2
          ? 'bg-green-300 hover:bg-green-300'
          : ''}"
        on:click={() => page_idx.set(2)}
      >
        <Cuboid />
      </button>
    </div>

    <a
      href="#"
      class="w-full mb-6 mt-auto flex justify-center tooltip tooltip-right"
      data-tip="github.com/leoxiang66/Dingo-Application"
      on:click={() => window.runtime.BrowserOpenURL("https://github.com/leoxiang66/Dingo-Application")}
    >
      <button class="sidebar-item">
        <GithubIcon class="w-[20px]" />
      </button>
    </a>
  </div>

  <div
    class="ml-[-1px] divider divider-horizontal before:w-[0.1px] after:w-[0.1px] before:bg-gray-300/30 after:bg-gray-300/30"
  ></div>
</div>
