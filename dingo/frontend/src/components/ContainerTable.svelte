<script>
  import { containers } from "../stores";
  import StartCtn from "./StartCtn.svelte";
  import StopCtn from "./StopCtn.svelte";

  // Helper function to convert any value to string
  function toString(value) {
    if (value === null || value === undefined) return "-";
    if (typeof value === "object") return JSON.stringify(value);
    return String(value);
  }
</script>

<div
  class="shadow-lg overflow-auto h-4/5 w-[85%] bg-white block m-auto rounded-xl text-black"
>
  <table class="table">
    <!-- Table header -->
    <thead class="text-black bg-green-200">
      <tr>
        <th>#</th>
        <th>ID</th>
        <th>Image</th>
        <!-- <th>Command</th> -->
        <th>Created</th>
        <th>Status</th>
        <!-- <th>Size</th> -->
        <th>Ports</th>
        <th>Names</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      {#each $containers as container, index}
        <tr class="hover:bg-green-100/30 text-sm font-thin h-[100px]">
          <th>{index + 1}</th>
          <td class="font-mono">{toString(container.ID).slice(0, 12)}</td>
          <td>{toString(container.Image)}</td>
          <!-- <td>{toString(container.Command)}</td> -->
          <td>{toString(container.Created)}</td>
          <td>{toString(container.Status)}</td>
          <!-- <td>{toString(container.SizeRw)} bytes</td> -->
          <td>{toString(container.Ports)}</td>
          <td>{toString(container.Names).slice(2, -2)}</td>
          <td>
            <div class="flex justify-center items-center"><StartCtn {container} /> <StopCtn /></div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
