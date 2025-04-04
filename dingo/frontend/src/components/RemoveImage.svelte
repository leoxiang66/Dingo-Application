<script>
  let modalRef;

  function openModal() {
    modalRef?.showModal();
  }
</script>

<button class="btn mx-1  btn-neutral btn-soft btn-sm" on:click={openModal}>delete</button>

<dialog bind:this={modalRef} class="modal">
  <div class="modal-box">
    <h3 class="text-lg font-bold">Hello!</h3>
    <p class="py-4">Press ESC key or click the button below to close</p>
    <div class="modal-action">
      <form method="dialog">
        <button class="btn">Close</button>
      </form>
    </div>
  </div>
</dialog>

<style>
  /* 可选：添加 modal 样式 */
  .modal-box {
    @apply w-11/12 max-w-md bg-white text-black;
  }
</style>
