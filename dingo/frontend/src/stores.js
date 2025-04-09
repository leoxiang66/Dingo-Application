import { writable } from "svelte/store";
import { Get_all_containers ,Get_all_images} from "../wailsjs/go/main/App";

// Initialize stores
export const page_idx = writable(0);
export const containers = writable([]);
export const images = writable([]);



async function load_img_data() {
    try {
        const imageData = await Get_all_images();
        images.set(imageData);
    } catch (error) {
        console.error("Failed to load images:", error);
    }
}

async function load_container_data() {
    try {
        const containerData = await Get_all_containers();
        containers.set(containerData);
    } catch (error) {
        console.error("Failed to load containers:", error);
    }
}

load_img_data();
load_container_data();