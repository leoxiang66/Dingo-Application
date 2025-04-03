<script>
  import { images } from "../stores"; // Import your images store
  import RemoveImage from "./RemoveImage.svelte"; // Assuming you have an image removal component

  // Helper function to convert any value to string
  function toString(value) {
    if (value === null || value === undefined) return "-";
    if (typeof value === "object") return JSON.stringify(value);
    return String(value);
  }

  // Helper to format size in MB
  function toMB(bytes) {
    return bytes ? (bytes / (1024 * 1024)).toFixed(2) + " MB" : "-";
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
        <th>IMAGE ID</th>
        <th>REPOSITORY:TAG</th>
        <th>SIZE</th>
        <th>SHARED SIZE</th>
        <th>CREATED</th>
        <th>VIRTUAL SIZE</th>
        <th>ACTIONS</th>
      </tr>
    </thead>
    <tbody>
      {#each $images as image, index}
        <tr class="hover:bg-green-100/30 text-sm font-thin h-[100px]">
          <th>{index + 1}</th>
          <td class="font-mono">
            {#if image.ID && image.ID.length > 12}
              {image.ID.slice(7, 19)}
            {:else}
              {toString(image.ID)}
            {/if}
          </td>
          <td>
            {#if image.RepoTags && image.RepoTags.length > 0}
              {image.RepoTags[0]}
            {:else}
              &lt;none&gt;:&lt;none&gt;
            {/if}
          </td>
          <td>{toMB(image.Size)}</td>
          <td>{toMB(image.SharedSize)}</td>
          <td>
            {toString(image.Created)}
          </td>
          <td>{toMB(image.VirtualSize)}</td>
          <td>
            <div class="flex justify-center items-center">
              <RemoveImage {image} />
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>