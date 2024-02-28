import "vite/modulepreload-polyfill";
import "./styles.css";
import "flowbite";
import Alpine from "alpinejs";
import htmx from "htmx.org";

window.Alpine = Alpine;
window.htmx = htmx;

Alpine.start();
