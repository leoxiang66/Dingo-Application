// import { writable } from "svelte/store";
import { Get_all_containers, Get_all_images } from "../wailsjs/go/main/App";

// Initialize stores
// export const page_idx = writable(0);
// export const containers = writable([]);
// export const images = writable([]);

export async function load_img_data() {
    try {
        return await Get_all_images();
    } catch (error) {
        console.error("Failed to load images:", error);
    }
}

export async function load_container_data() {
    try {
        return await Get_all_containers();
    } catch (error) {
        console.error("Failed to load containers:", error);
    }
}
