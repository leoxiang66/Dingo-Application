<script>
  import Modal from "./Modal.svelte";
</script>

<div
  class="shadow-lg overflow-auto h-3/4 w-3/4 bg-white block m-auto rounded-xl text-black"
>
  <table class="table">
    <!-- head -->
    <thead class="text-black bg-green-200">
      <tr>
        <th></th>
        <th>Name</th>
        <th>Job</th>
        <th>Favorite Color</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
      <!-- row 1 -->
      <tr class="hover:bg-green-100/30">
        <th>1</th>
        <td>Cy Ganderton</td>
        <td>Quality Control Specialist</td>
        <td>Blue</td>
        <td><Modal /></td>
      </tr>
      <!-- row 2 -->
      <tr class="hover:bg-green-100/30">
        <th>2</th>
        <td>Hart Hagerty</td>
        <td>Desktop Support Technician</td>
        <td>Purple</td>
        <td><Modal /></td>
      </tr>
      <!-- row 3 -->
      <tr class="hover:bg-green-100/30">
        <th>3</th>
        <td>Brice Swyre</td>
        <td>Tax Accountant</td>
        <td>Red</td>
        <td><Modal /></td>
      </tr>
    </tbody>
  </table>
</div>
