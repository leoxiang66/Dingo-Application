import { writable } from "svelte/store";
import { Get_all_containers } from "../wailsjs/go/main/App";

// Initialize stores
export const page_idx = writable(0);
export const containers = writable([]);

// Load containers asynchronously
(async () => {
    try {
        const containerData = await Get_all_containers();
        containers.set(containerData);
    } catch (error) {
        console.error("Failed to load containers:", error);
    }
})();