import { writable } from "svelte/store";
import { Get_all_containers ,Get_all_images} from "../wailsjs/go/main/App";

// Initialize stores
export const page_idx = writable(0);
export const containers = writable([]);
export const images = writable([]);

// Load containers asynchronously
(async () => {
    try {
        const containerData = await Get_all_containers();
        const imageData = await Get_all_images();
        containers.set(containerData);
        images.set(imageData);
    } catch (error) {
        console.error("Failed to load containers:", error);
    }
})();